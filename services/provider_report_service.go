package services

import (
	"errors"
	"payment_reconciliation/dto"
	"payment_reconciliation/models"
	"payment_reconciliation/repository"

	"github.com/shopspring/decimal"
)

type ProviderReportService struct {
	ProviderReportRepo *repository.ProviderReportRepository
}

func (s *ProviderReportService) ImportTransaction(req dto.ProviderReportRequest) error {
	if len(req.Transactions) == 0 {
		return errors.New("report must contain at least one transaction")
	}

	existingTransactions, err := s.ProviderReportRepo.FindbyReportID(req.ReportId)
	if err != nil {
		return err
	}
	if len(existingTransactions) > 0 {
		return errors.New("this report is already imported")
	}
	transactions := make([]models.ProviderReportTransaction, 0, len(req.Transactions))

	for _, item := range req.Transactions {
		// str to decimal
		amount, err := decimal.NewFromString(item.Amount)
		if err != nil {
			return errors.New("invalid transaction amount")
		}

		if amount.LessThanOrEqual(decimal.Zero) {
			return errors.New("transaction amount must be greater than zero")
		}

		transaction := models.ProviderReportTransaction{
			ProviderRef: item.ProviderRef,
			Amount:      amount,
			Status:      item.Status,
			ReportID:    req.ReportId,
		}

		transactions = append(transactions, transaction)
	}

	return s.ProviderReportRepo.CreateTransactions(&transactions)
}
