package handlers

import (
	"net/http"

	"payment_reconciliation/services"

	"github.com/gin-gonic/gin"
)

type ReconciliationHandler struct {
	ReconciliationService *services.ReconciliationService
}

func (h *ReconciliationHandler) Reconcile(c *gin.Context) {

	reportID := c.Param("report_id")

	if reportID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "report id is required",
		})
		return
	}

	job, err := h.ReconciliationService.Reconcile(reportID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "reconciliation completed successfully",
		"job":     job,
	})
}