package handlers

import (
	"net/http"

	"payment_reconciliation/services"
	"payment_reconciliation/workers"

	"github.com/gin-gonic/gin"
)

type ReconciliationHandler struct {
	ReconciliationService *services.ReconciliationService
	WorkerPool            *workers.ReconciliationWorkerPool
}

// Reconcile godoc
// @Summary Reconcile provider report
// @Description Compares provider report transactions with internal payment records and generates reconciliation results
// @Tags Reconciliation
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param report_id path string true "Provider Report ID"  default(REPORT-001)
// @Success 200 {object} models.ReconciliationJob
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/reconcil/{report_id} [post]
func (h *ReconciliationHandler) Reconcile(c *gin.Context) {

	reportID := c.Param("report_id")

	if reportID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "report id is required",
		})
		return
	}

	job, reportTransactions, err := h.ReconciliationService.StartReconciliation(reportID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	h.WorkerPool.Submit(
		workers.ReconciliationJobData{
			Job:                job,
			ReportTransactions: reportTransactions,
		},
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "reconciliation started",
		"job":     job,
	})
}
