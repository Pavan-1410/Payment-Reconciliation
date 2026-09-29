package repository

import (
	"payment_reconciliation/models"

	"gorm.io/gorm"
)

type ReconciliationResultRepository struct {
	DB *gorm.DB
}

func (r *ReconciliationResultRepository) CreateResults(results []models.ReconciliationResult,) error {
	return r.DB.Create(&results).Error
}