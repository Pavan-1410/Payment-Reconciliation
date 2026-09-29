package handlers

import (
	"net/http"

	"payment_reconciliation/dto"
	"payment_reconciliation/services"

	"github.com/gin-gonic/gin"
)

type ProviderReportHandler struct {
	ProviderReportService *services.ProviderReportService
}

func (h *ProviderReportHandler) ImportReport (c *gin.Context) {

	var req dto.ProviderReportRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request data",
		})
		return
	}

	err := h.ProviderReportService.ImportTransaction(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "provider report imported successfully",
	})
}