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

func (s *ReconciliationService) StartReconciliation(reportID string) (*models.ReconciliationJob, error) {
	// Check that the report exists
	reportTransactions, err := s.ProviderReportRepo.FindbyReportID(reportID)
	if err != nil {
		return nil, err
	}
	if len(reportTransactions) == 0 {
		return nil, fmt.Errorf("no transactions found for report %q", reportID)
	}

	startTime := time.Now()
	job := &models.ReconciliationJob{
		Status:    "processing",
		StartedAt: &startTime,
	}
	if err := s.JobRepo.CreateJob(job); err != nil {
		return nil, err
	}
	    
	ctx, cancel := context.WithTimeout(
        context.Background(),
        30*time.Second,
    )

	go func() {
    defer cancel()
		s.ProcessReconciliation( ctx, job, reportTransactions )
    }()


	return job,nil

}

func(s *ReconciliationService) ProcessReconciliation(ctx context.Context, job *models.ReconciliationJob,reportTransactions []models.ProviderReportTransaction){
	// Job has now started processing
	job.Status = "processing"

	if err := s.JobRepo.UpdateJob(job); err != nil {
		return
	}
	results := make([]models.ReconciliationResult, 0, len(reportTransactions))

		for _, reportTransaction := range reportTransactions {
			    select {
   		 			case <-ctx.Done():
        			s.failJob(job, ctx.Err())
        				return
					default:	
						eachDbTransaction, err := s.ProviderTransactionRepo.FindByProviderRef(reportTransaction.ProviderRef)
						if errors.Is(err, gorm.ErrRecordNotFound) {
							results = append(results, models.ReconciliationResult{
								ReconciliationJobID:         job.ID,
								ProviderReportTransactionID: reportTransaction.ID,
								ProviderTransactionID:       nil,
								Result:                      "missing internal payment",
								Details:                     "provider transaction not found internally",
							})
							continue
						}

						if !reportTransaction.Amount.Equal(eachDbTransaction.Payment.Amount) {
							results = append(results, models.ReconciliationResult{
								ReconciliationJobID:         job.ID,
								ProviderReportTransactionID: reportTransaction.ID,
								ProviderTransactionID:       &eachDbTransaction.ID,
								Result:                      "amount_mismatched",
								Details:                     "provider amount does not match internal amount",
							})
							continue
						}

						if reportTransaction.Status != eachDbTransaction.Status {
							results = append(results, models.ReconciliationResult{
								ReconciliationJobID:         job.ID,
								ProviderReportTransactionID: reportTransaction.ID,
								ProviderTransactionID:       &eachDbTransaction.ID,
								Result:                      "status_mismatch",
								Details:                     "provider status does not match internal status",
							})
							continue
						}

						results = append(results, models.ReconciliationResult{
							ReconciliationJobID:         job.ID,
							ProviderReportTransactionID: reportTransaction.ID,
							ProviderTransactionID:       &eachDbTransaction.ID,
							Result:                      "matched",
							Details:                     "provider transaction matches internal transaction",
						})
											// Save results
					if err := s.ResultRepo.CreateResults(results); err != nil {
						s.failJob(job, err)
						return
					}

					// Complete job
					completedAt := time.Now()

					job.Status = "completed"
					job.CompletedAt = &completedAt

					if err := s.JobRepo.UpdateJob(job); err != nil {
						return
					}
				}
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
