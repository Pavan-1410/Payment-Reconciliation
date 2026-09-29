package dto

type ProviderTransactionResponse struct {
	ID            int    `json:"id"`
	PaymentID     int    `json:"payment_id"`
	ProviderRef   string `json:"provider_ref"`
	AttemptNumber int    `json:"attempt_number"`
	Status        string `json:"status"`
	FailureReason string `json:"failure_reason"`
}