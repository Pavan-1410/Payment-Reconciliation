package routes

import (
	"payment_reconciliation/handlers"
	"payment_reconciliation/middelware"

	"github.com/gin-gonic/gin"
)

func AuthRouter(routes *gin.Engine, authhandler *handlers.AuthHandler){
	auth := routes.Group("/api/auth")
	
	auth.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to New Project API",
		})
	})

	auth.POST("/register",authhandler.Register)	// we need to call using struct instance
	auth.POST("/login",authhandler.Login)

	protected := routes.Group("/api/protected")
	protected.Use(middelware.AuthMiddleware())


	protected.GET("/", func(c *gin.Context) {
		c.JSON(200, "welcome to protected routes")
	})
}
func PaymentRouter (routes *gin.Engine, paymenthandler *handlers.PaymentHandler){
	payment := routes.Group("/api/payment")
	payment.Use(middelware.AuthMiddleware())
	payment.POST("/create",paymenthandler.CreatePayment)
}
func ProviderRouter (routes *gin.Engine, providerhandler *handlers.ProviderTransactionHandler ){
	transaction :=routes.Group("/api/transaction/:id/:change")
	transaction.Use(middelware.AuthMiddleware())
	transaction.POST("/",providerhandler.ProcessPayment)
}
func ProviderReportRoutes (routes *gin.Engine, providerhandler *handlers.ProviderReportHandler){
	report :=routes.Group("/api/provider")
	report.Use(middelware.AuthMiddleware())
	report.POST("/report",providerhandler.ImportReport)
}
func ReconciliationRoutes (routes *gin.Engine, providerhandler *handlers.ReconciliationHandler){
	reconcil :=routes.Group("/api/reconcil")
	reconcil.Use(middelware.AuthMiddleware())
	reconcil.POST("/:report_id",providerhandler.Reconcile)
}


