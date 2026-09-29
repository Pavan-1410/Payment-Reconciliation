package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`	// as we are using DTO for request so we not need binding required here
	Name      string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	PasswordHash string    `gorm:"not null"`
	IsAdmin   bool   `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}