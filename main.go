package main

import (
	"context"
	"log"
	"payment_reconciliation/config"
	"payment_reconciliation/handlers"
	"payment_reconciliation/models"
	"payment_reconciliation/repository"
	"payment_reconciliation/routes"
	"payment_reconciliation/services"
	"payment_reconciliation/workers"

	"github.com/gin-gonic/gin"
)


func main() {

	config.ConnectDatabase()

	err := config.DB.AutoMigrate(
		&models.User{},
		&models.Payment{},
		&models.ProviderTransaction{},
		&models.ProviderReportTransaction{},
		&models.ReconciliationJob{},
		&models.ReconciliationResult{},
	
	)

	if err != nil {
		log.Fatal("Database migration failed:", err)
	}
	log.Println("Database connected and migration completed")

	// dependency wiring

	// call the DB repo constructor of user
	userRepo := repository.NewUserRepository(config.DB)

	//we are creating an object (instance) of the AuthService struct. and passing the and giving him user repository
	authService := &services.AuthService{	
    	UserRepo: userRepo,	// and giving him user repository
	}

	authHandler := &handlers.AuthHandler{
		AuthService: authService,
	}

	// payment dependency wiring
	paymentRepo := &repository.PaymentRepository{
		DB : config.DB,
	}

	paymentService:= &services.PaymentService{
		PaymentRepo: paymentRepo,
	}
	
	paymentHandler:= &handlers.PaymentHandler{
		PaymentService: paymentService,
	}

	// transaction dependency wiring
	providerTransactionrepo := &repository.ProviderTransactionRepository{
		DB: config.DB,
	}

	providerTransactionService := &services.ProviderTransactionServices{
		ProviderTransactionRepo: providerTransactionrepo,
		PaymentRepo: paymentRepo,
	}

	providerTransactionHandler := &handlers.ProviderTransactionHandler{
		ProviderTransactionService: providerTransactionService,
	}

	// provider report transaction wiring

	providerReportrepo := &repository.ProviderReportRepository{
		DB : config.DB,
	}

	providerReportService := &services.ProviderReportService{
		ProviderReportRepo: providerReportrepo,
	}

	providerReportHandler := &handlers.ProviderReportHandler{
		ProviderReportService: providerReportService,
	}

	// reconcialition wiring

	reconciliationJobRepo := &repository.ReconciliationJobRepository{
	DB: config.DB,
	}

	reconcialitionResultrepo := &repository.ReconciliationResultRepository{
		DB : config.DB,
	}

	reconcialitionService :=&services.ReconciliationService{
		JobRepo:reconciliationJobRepo,
		ResultRepo: reconcialitionResultrepo,
		ProviderReportRepo:providerReportrepo,
		ProviderTransactionRepo: providerTransactionrepo,
	}


	// workpool dependency

	workerPool := &workers.ReconciliationWorkerPool{
	JobQueue:              make(chan workers.ReconciliationJobData, 100),
	WorkerCount:           3,
	ReconciliationService: reconcialitionService,
	}

	reconcialitionHandlers := &handlers.ReconciliationHandler{
		ReconciliationService : reconcialitionService,
		 WorkerPool:            workerPool,
	}
	// satrting the work pool
	ctx, cancel := context.WithCancel(context.Background())
defer cancel()

	workerPool.Start(ctx)
	r := gin.Default()

	routes.AuthRouter(r,authHandler)
	routes.PaymentRouter(r,paymentHandler)
	routes.ProviderRouter(r,providerTransactionHandler)
	routes.ProviderReportRoutes(r,providerReportHandler)
	routes.ReconciliationRoutes(r,reconcialitionHandlers)

	r.Run(":8080")
}
// this is the main changes that are goin to staging