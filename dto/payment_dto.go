package dto

type CreatePaymentRequest struct {
	Amount         string `json:"amount" binding:"required" example:"500.00"`
	IdempotencyKey string `json:"idempotency_key" binding:"required" example:"PAY-TEST-001"`
}

type PaymentResponce struct{
	ID             int    `json:"id"`
	UserID         int    `json:"user_id"`
	Amount         string `json:"amount"`
	Status         string `json:"status"`
	IdempotencyKey string `json:"idempotency_key"`
}