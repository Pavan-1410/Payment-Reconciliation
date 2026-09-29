package repository

import (
	"payment_reconciliation/models"

	"gorm.io/gorm"
)

type ProviderReportRepository struct {
	DB *gorm.DB
}

// this fills the transaction in provicer report tbl
func (r *ProviderReportRepository) CreateTransaction(transaction *models.ProviderReportTransaction) error {
	return r.DB.Create(transaction).Error
}
func (r *ProviderReportRepository) CreateTransactions(transaction *[]models.ProviderReportTransaction) error {
	return r.DB.Create(transaction).Error
}
func (r *ProviderReportRepository) FindbyReportID(reportID string) ([]models.ProviderReportTransaction, error) {
	var report []models.ProviderReportTransaction
	err := r.DB.Where("report_id = ?", reportID).Find(&report).Error
	if err != nil {
		return nil, err
	}
	return report, nil
}
