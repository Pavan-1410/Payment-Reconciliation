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

// ImportReport godoc
// @Summary Import provider report
// @Description Imports simulated payment provider transaction records for reconciliation
// @Tags Provider Reports
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.ProviderReportRequest true "Provider report details"
// @Success 201 {object} dto.ProviderReportRequest
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/provider/report [post]
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