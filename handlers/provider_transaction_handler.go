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
