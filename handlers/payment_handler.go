package handlers

import (
	"net/http"
	"payment_reconciliation/dto"
	"payment_reconciliation/services"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	PaymentService *services.PaymentService
}

// CreatePayment godoc
// @Summary Create a payment
// @Description Creates a new payment request for the authenticated user
// @Tags Payments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreatePaymentRequest true "Payment details"
// @Success 201 {object} dto.PaymentResponce
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 409 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/payment/create [post]
func (h *PaymentHandler) CreatePayment (c *gin.Context){

	userIDValue, exists := c.Get("user_id")

	if !exists{ 
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}
		
	userID, ok := userIDValue.(int)

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid user id",
		})
		return
	}

	var request dto.CreatePaymentRequest

	if err := c.ShouldBindJSON(&request); err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request data",
		})
		return
	}

	// call service
	payment, err := h.PaymentService.CreatePayment(userID, request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 4. Return response
	c.JSON(http.StatusCreated, gin.H{
		"message": "payment created successfully",
		"payment": payment,
	})

}