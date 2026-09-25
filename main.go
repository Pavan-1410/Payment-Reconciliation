package main

import (
	"log"
	"payment_reconciliation/config"
	"payment_reconciliation/models"
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

	
	r := gin.Default()


	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to New Project API",
		})
	})

	r.Run(":8080")
}