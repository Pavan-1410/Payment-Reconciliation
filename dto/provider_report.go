package dto

type ProviderReportTransactionRequest struct {
	ProviderRef string `json:"provider_ref" binding:"required"`
	Amount      string `json:"amount" binding:"required"`
	Status      string `json:"status" binding:"required"`
}

type ProviderReportRequest struct {
	ReportId     string                             `json:"report_id" binding:"required"`
	Transactions []ProviderReportTransactionRequest `json:"transaction" binding:"required"`
}
