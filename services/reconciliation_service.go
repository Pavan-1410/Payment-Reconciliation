package services

import (
	"context"
	"errors"
	"fmt"
	"payment_reconciliation/models"
	"payment_reconciliation/repository"
	"time"

	"gorm.io/gorm"
)

type ReconciliationService struct {
	JobRepo                 *repository.ReconciliationJobRepository
	ResultRepo              *repository.ReconciliationResultRepository
	ProviderReportRepo      *repository.ProviderReportRepository
	ProviderTransactionRepo *repository.ProviderTransactionRepository
}

func (s *ReconciliationService) StartReconciliation(reportID string) (*models.ReconciliationJob, []models.ProviderReportTransaction, error) {
	reportTransactions, err := s.ProviderReportRepo.FindbyReportID(reportID)
	if err != nil {
		return nil, nil, err
	}
	if len(reportTransactions) == 0 {
		return nil, nil, fmt.Errorf("no transactions found for report %q", reportID)
	}

	startTime := time.Now()
	job := &models.ReconciliationJob{
		Status:    "processing",
		StartedAt: &startTime,
	}
	if err := s.JobRepo.CreateJob(job); err != nil {
		return nil, nil, err
	}
	return job, reportTransactions, nil
}

func (s *ReconciliationService) ProcessReconciliation(ctx context.Context, job *models.ReconciliationJob, reportTransactions []models.ProviderReportTransaction) {
	job.Status = "processing"
	if err := s.JobRepo.UpdateJob(job); err != nil {
		s.failJob(job, err)
		return
	}

	results := make([]models.ReconciliationResult, 0, len(reportTransactions))
	for _, reportTransaction := range reportTransactions {
		select {
		case <-ctx.Done():
			s.failJob(job, ctx.Err())
			return
		default:
		}

		internalTransaction, err := s.ProviderTransactionRepo.FindByProviderRef(reportTransaction.ProviderRef)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			results = append(results, models.ReconciliationResult{
				ReconciliationJobID:         job.ID,
				ProviderReportTransactionID: reportTransaction.ID,
				Result:                      "missing internal payment",
				Details:                     "provider transaction not found internally",
			})
			continue
		}
		if err != nil {
			s.failJob(job, err)
			return
		}

		result := models.ReconciliationResult{
			ReconciliationJobID:         job.ID,
			ProviderReportTransactionID: reportTransaction.ID,
			ProviderTransactionID:       &internalTransaction.ID,
		}
		switch {
		case !reportTransaction.Amount.Equal(internalTransaction.Payment.Amount):
			result.Result = "amount_mismatched"
			result.Details = "provider amount does not match internal amount"
		case reportTransaction.Status != internalTransaction.Status:
			result.Result = "status_mismatch"
			result.Details = "provider status does not match internal status"
		default:
			result.Result = "matched"
			result.Details = "provider transaction matches internal transaction"
		}
		results = append(results, result)
	}

	if err := s.ResultRepo.CreateResults(results); err != nil {
		s.failJob(job, err)
		return
	}

	completedAt := time.Now()
	job.Status = "completed"
	job.CompletedAt = &completedAt
	if err := s.JobRepo.UpdateJob(job); err != nil {
		s.failJob(job, err)
	}
}

func (s *ReconciliationService) failJob(job *models.ReconciliationJob, cause error) error {
	job.Status = "failed"
	job.Error = cause.Error()
	if err := s.JobRepo.UpdateJob(job); err != nil {
		return errors.Join(cause, fmt.Errorf("updating failed reconciliation job: %w", err))
	}
	return cause
}
