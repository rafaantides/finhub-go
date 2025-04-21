package handler

import (
	"finhub-go/internal/core/ports/inbound"

	"github.com/gin-gonic/gin"
)

type UploadHandler struct {
	service inbound.UploadService
}

func NewUploadHandler(service inbound.UploadService) *UploadHandler {
	return &UploadHandler{service: service}
}

func (h *UploadHandler) ProcessFileHandler(c *gin.Context) {
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
