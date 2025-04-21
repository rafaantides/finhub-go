package handler

import (
	"errors"
	"finhub-go/internal/config"
	"finhub-go/internal/core/dto"
	appError "finhub-go/internal/core/errors"
	"finhub-go/internal/core/ports/inbound"
	"finhub-go/internal/utils"
	"finhub-go/internal/utils/pagination"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InvoiceHandler struct {
	service inbound.InvoiceService
}

func NewInvoiceHandler(service inbound.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{service: service}
}

func (h *InvoiceHandler) CreateInvoiceHandler(c *gin.Context) {
	ctx := c.Request.Context()
	var req dto.InvoiceRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(appError.NewAppError(http.StatusBadRequest, err))
		return
	}

	input, err := req.ToDomain()
	if err != nil {
		c.Error(appError.NewAppError(http.StatusBadRequest, err))
		return
	}

	data, err := h.service.CreateInvoice(ctx, *input)
	if err != nil {
		c.Error(appError.NewAppError(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusCreated, data)
}

func (h *InvoiceHandler) GetInvoiceByIDHandler(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := utils.ToUUID(c.Param("id"))
	if err != nil {
		c.Error(appError.NewAppError(http.StatusBadRequest, err))
		return
	}

	data, err := h.service.GetInvoiceByID(ctx, id)
	if err != nil {
		if errors.Is(err, appError.ErrNotFound) {
			c.Error(appError.NewAppError(http.StatusNotFound, err))
			return
		}
		c.Error(appError.NewAppError(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *InvoiceHandler) ListInvoicesHandler(c *gin.Context) {
	ctx := c.Request.Context()
	var flt dto.InvoiceFilters
	if err := c.ShouldBindQuery(&flt); err != nil {
		c.Error(appError.NewAppError(http.StatusBadRequest, err))
		return
	}

	pgn, err := pagination.NewPagination(c)

	if err != nil {
		c.Error(appError.NewAppError(http.StatusBadRequest, err))
		return
	}

	validColumns := map[string]bool{
		"id":         true,
		"title":      true,
		"amount":     true,
		"issue_date": true,
		"due_date":   true,
		"status_id":  true,
		"status":     true,
		"created_at": true,
		"updated_at": true,
	}

	if err := pgn.ValidateOrderBy("issue_date", config.OrderAsc, validColumns); err != nil {
		c.Error(appError.NewAppError(http.StatusBadRequest, err))
		return
	}

	response, total, err := h.service.ListInvoices(ctx, flt, pgn)

	if err != nil {
		c.Error(appError.NewAppError(http.StatusInternalServerError, err))
		return
	}

	pgn.SetPaginationHeaders(c, total)

	c.JSON(http.StatusOK, response)
}

func (h *InvoiceHandler) UpdateInvoiceHandler(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := utils.ToUUID(c.Param("id"))
	if err != nil {
		c.Error(appError.NewAppError(http.StatusBadRequest, err))
		return
	}

	var req dto.InvoiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(appError.NewAppError(http.StatusBadRequest, err))
		return
	}

	input, err := req.ToDomain()
	if err != nil {
		c.Error(appError.NewAppError(http.StatusBadRequest, err))
		return
	}

	data, err := h.service.UpdateInvoice(ctx, id, *input)
	if err != nil {
		c.Error(appError.NewAppError(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *InvoiceHandler) DeleteInvoiceHandler(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := utils.ToUUID(c.Param("id"))
	if err != nil {
		c.Error(appError.NewAppError(http.StatusBadRequest, err))
		return
	}

	err = h.service.DeleteInvoiceByID(ctx, id)
	if err != nil {
		c.Error(appError.NewAppError(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *InvoiceHandler) ListInvoiceDebtsHandler(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := utils.ToUUID(c.Param("id"))
	if err != nil {
		c.Error(appError.NewAppError(http.StatusBadRequest, err))
		return
	}
	var flt dto.DebtFilters
	if err := c.ShouldBindQuery(&flt); err != nil {
		c.Error(appError.NewAppError(http.StatusBadRequest, err))
		return
	}

	pgn, err := pagination.NewPagination(c)

	if err != nil {
		c.Error(appError.NewAppError(http.StatusBadRequest, err))
		return
	}

	validColumns := map[string]bool{
		"id":            true,
		"invoice":       true,
		"title":         true,
		"category_id":   true,
		"category":      true,
		"amount":        true,
		"purchase_date": true,
		"due_date":      true,
		"status_id":     true,
		"status":        true,
		"created_at":    true,
		"updated_at":    true,
	}

	if err := pgn.ValidateOrderBy("purchase_date", config.OrderAsc, validColumns); err != nil {
		c.Error(appError.NewAppError(http.StatusBadRequest, err))
		return
	}

	response, total, err := h.service.ListInvoiceDebts(ctx, id, flt, pgn)

	if err != nil {
		c.Error(appError.NewAppError(http.StatusInternalServerError, err))
		return
	}

	pgn.SetPaginationHeaders(c, total)

	c.JSON(http.StatusOK, response)
}
