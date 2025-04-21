package routes

import (
	"finhub-go/internal/core/ports/outbound/messagebus"
	"finhub-go/internal/core/ports/outbound/repository"
	"finhub-go/internal/core/service"
	"finhub-go/internal/core/service/upload"
	"finhub-go/internal/http/handler"
	"finhub-go/internal/http/middlewares"
	"finhub-go/internal/utils/logger"

	"github.com/gin-gonic/gin"
)

type Router struct {
	log  *logger.Logger
	repo repository.Repository
	mbus messagebus.MessageBus
}

func NewRouter(log *logger.Logger, repo repository.Repository, mbus messagebus.MessageBus) *Router {
	return &Router{
		log:  log,
		repo: repo,
		mbus: mbus,
	}
}

func (r *Router) Setup(enableDebug bool) *gin.Engine {
	if enableDebug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.Default()
	v1 := engine.Group("/api/v1")

	v1.Use(middlewares.ErrorMiddleware(r.log))
	// v1.Use(middlewares.CORSMiddleware())
	// TODO: rever esse middlewere aqui
	// v1.Use(middlewares.UUIDMiddleware(r.log))

	engine.StaticFile("/favicon.ico", "./static/favicon.ico")
	registerDocsRoutes(engine.Group("/docs/v1"))

	debtService := service.NewDebtService(r.repo)
	debtHandler := handler.NewDebtHandler(debtService)
	registerDebtRoutes(v1.Group("/debts"), debtHandler)

	invoiceService := service.NewInvoiceService(r.repo)
	invoiceHandler := handler.NewInvoiceHandler(invoiceService)
	registerInvoiceRoutes(v1.Group("/invoices"), invoiceHandler)

	categoryService := service.NewCategoryService(r.repo)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	registerCategoryRoutes(v1.Group("/categories"), categoryHandler)

	paymentStatusService := service.NewPaymentStatusService(r.repo)
	paymentStatusHandler := handler.NewPaymentStatusHandler(paymentStatusService)
	registerPaymentStatusRoutes(v1.Group("/payment_status"), paymentStatusHandler)

	uploadService := upload.NewUploadService(r.mbus)
	uploadHander := handler.NewUploadHandler(uploadService)
	registerUploadRoutes(v1.Group("/upload"), uploadHander)

	return engine
}

func registerDocsRoutes(router *gin.RouterGroup) {
	router.StaticFile("/swagger.json", "./docs/v1/swagger.json")

	router.GET("", func(c *gin.Context) {
		c.File("./docs/v1/docs.html")
	})
	router.GET("/redoc", func(c *gin.Context) {
		c.File("./docs/v1/redoc.html")
	})
	router.GET("/scalar", func(c *gin.Context) {
		c.File("./docs/v1/scalar.html")
	})
	router.GET("/swagger", func(c *gin.Context) {
		c.File("./docs/v1/swagger.html")
	})

}

func registerDebtRoutes(router *gin.RouterGroup, handler *handler.DebtHandler) {
	router.POST("", handler.CreateDebtHandler)
	router.GET("", handler.ListDebtsHandler)
	router.GET("/:id", handler.GetDebtByIDHandler)
	router.PUT("/:id", handler.UpdateDebtHandler)
	router.DELETE("/:id", handler.DeleteDebtHandler)
}

func registerInvoiceRoutes(router *gin.RouterGroup, handler *handler.InvoiceHandler) {
	router.POST("", handler.CreateInvoiceHandler)
	router.GET("", handler.ListInvoicesHandler)
	router.GET("/:id", handler.GetInvoiceByIDHandler)
	router.PUT("/:id", handler.UpdateInvoiceHandler)
	router.DELETE("/:id", handler.DeleteInvoiceHandler)
	router.GET("/:id/debts", handler.ListInvoiceDebtsHandler)
}

func registerCategoryRoutes(router *gin.RouterGroup, handler *handler.CategoryHandler) {
	router.POST("", handler.CreateCategoryHandler)
	router.GET("", handler.ListCategorysHandler)
	router.GET("/:id", handler.GetCategoryByIDHandler)
	router.PUT("/:id", handler.UpdateCategoryHandler)
	router.DELETE("/:id", handler.DeleteCategoryHandler)
}

func registerPaymentStatusRoutes(router *gin.RouterGroup, handler *handler.PaymentStatusHandler) {
	router.POST("", handler.CreatePaymentStatusHandler)
	router.GET("", handler.ListPaymentStatussHandler)
	router.GET("/:id", handler.GetPaymentStatusByIDHandler)
	router.PUT("/:id", handler.UpdatePaymentStatusHandler)
	router.DELETE("/:id", handler.DeletePaymentStatusHandler)
}

func registerUploadRoutes(router *gin.RouterGroup, handler *handler.UploadHandler) {
	router.POST("", handler.ProcessFileHandler)
}
