package repository

import (
	"payment_reconciliation/models"

	"gorm.io/gorm"
)

type ReconciliationJobRepository struct {
	DB *gorm.DB
}

func (r *ReconciliationJobRepository) CreateJob(job *models.ReconciliationJob) error {
	return r.DB.Create(job).Error
}
// to update status of the job
func (r *ReconciliationJobRepository) UpdateJob(job *models.ReconciliationJob) error {
	return r.DB.Save(job).Error
}