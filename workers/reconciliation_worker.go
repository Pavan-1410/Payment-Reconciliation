package workers

import (
	"context"
	"fmt"
	"sync"

	"payment_reconciliation/models"
	"payment_reconciliation/services"
)

type ReconciliationJobData struct {
    Job              *models.ReconciliationJob
    ReportTransactions []models.ProviderReportTransaction
}

type ReconciliationWorkerPool struct {
	JobQueue chan ReconciliationJobData	// channel it stores reconciliation jobIDs
	WorkerCount         int
	ReconciliationService *services.ReconciliationService	//workers use this to process jobs
	wg                  sync.WaitGroup
}

// stary the worker

func (p *ReconciliationWorkerPool) Start(ctx context.Context) {

	for i := 1; i <= p.WorkerCount; i++ {

		p.wg.Add(1)

		go p.worker(ctx, i)
	}
}

func (p *ReconciliationWorkerPool) worker(ctx context.Context,workerID int) {
	defer p.wg.Done()

	for {
		select {

		case jobData := <-p.JobQueue:

			fmt.Printf(
				"Worker %d processing job %d\n",
				workerID,
				jobData.Job.ID,
			)

			
			p.ReconciliationService.ProcessReconciliation(
				ctx,
                jobData.Job,
                jobData.ReportTransactions,
			)

		case <-ctx.Done():

			fmt.Printf(
				"Worker %d stopping\n",
				workerID,
			)

			return
		}
	}
}

func (p *ReconciliationWorkerPool) Submit(
	job ReconciliationJobData	,
) {
	p.JobQueue <- job
}

