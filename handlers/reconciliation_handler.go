package handlers

import (
	"net/http"

	"payment_reconciliation/models"
	"payment_reconciliation/services"
	"payment_reconciliation/workers"

	"github.com/gin-gonic/gin"
)

type ReconciliationHandler struct {
	ReconciliationService *services.ReconciliationService
	WorkerPool            *workers.ReconciliationWorkerPool
}

func (h *ReconciliationHandler) Reconcile(c *gin.Context) {

	reportID := c.Param("report_id")

	if reportID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "report id is required",
		})
		return
	}

	job, err := h.ReconciliationService.StartReconciliation(reportID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	h.WorkerPool.Submit(
		workers.ReconciliationJobData{
			Job:              job,
			ReportTransactions:[]models.ProviderReportTransaction{},
		},
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "reconciliation completed successfully",
		"job":     job,
	})
}
