package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB	//// DB can hold a pointer to a GORM database object
func ConnectDatabase(){
	if err := godotenv.Load();err != nil{
		log.Println("cannot find .env file")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	if dsn == ""{
		log.Fatal("Failed to create DSN")
	}

	db , err := gorm.Open(postgres.Open(dsn), &gorm.Config{})	//&gorm.Config{} is a struct type provided by gorm // 
				// gorm.Open returns the ptr
	if err != nil{
		log.Fatal("Failed to connect database")
	}
	DB = db
	log.Println("Database connected successfully")
}