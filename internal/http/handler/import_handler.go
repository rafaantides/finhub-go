package handler

import (
	"finhub-go/internal/core/ports/inbound"

	"github.com/gin-gonic/gin"
)

type ImporterHandler struct {
	service inbound.ImporterService
}

func NewImporterHandler(service inbound.ImporterService) *ImporterHandler {
	return &ImporterHandler{service: service}
}

func (h *ImporterHandler) ProcessFileHandler(c *gin.Context) {
	resource := c.PostForm("resource")
	action := c.PostForm("action")
	model := c.PostForm("model")

	if resource == "" || action == "" || model == "" {
		c.JSON(400, gin.H{"error": "Parâmetros 'resource' , 'model e 'action' são obrigatórios"})
		return
	}

	file, fileHeader, err := c.Request.FormFile("file")

	if err != nil {
		c.JSON(400, gin.H{"error": "Falha ao obter o arquivo"})
		return
	}
	defer file.Close()

	h.service.ImportFile(resource, model, action, file, fileHeader)
}
