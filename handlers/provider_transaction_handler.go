package handlers

import (
	"net/http"
	"payment_reconciliation/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProviderTransactionHandler struct {
	ProviderTransactionService *services.ProviderTransactionServices
}

// ProcessPayment godoc
// @Summary Process a payment
// @Description Processes a pending payment and updates its status
// @Tags 4. Transaction
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path uint true "Payment ID" default(3)
// @Param change path string true "Change Status to" Enums(success, fail) default(success)
// @Success 200 {object} dto.ProviderTransactionResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/transaction/{id}/{change}/ [post]
func (h *ProviderTransactionHandler) ProcessPayment (c * gin.Context){
	paymentId,err := strconv.Atoi(c.Param("id"))

	result := c.Param("change")

	if result != "success" && result != "fail" {
    c.JSON(http.StatusBadRequest, gin.H{
        "error": "invalid result. Use success or fail",
    })
    return
}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid payment id",
		})
		return
	}

	transaction,err := h.ProviderTransactionService.ProcessPayment(paymentId,result)
		
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message":     "payment processed successfully",
		"transaction": transaction,
	})

}
