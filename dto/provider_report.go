package dto

type ProviderReportTransactionRequest struct {
	ProviderRef string `json:"provider_ref" binding:"required" example:"TXN-1-1"`
	Amount      string `json:"amount" binding:"required" example:"500.00"`
	Status      string `json:"status" binding:"required" example:"successful"`
}

type ProviderReportRequest struct {
	ReportId     string                             `json:"report_id" binding:"required" example: "REPORT-001"`
	Transactions []ProviderReportTransactionRequest `json:"transaction" binding:"required" :`
}
