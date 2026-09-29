package models

import "time"

type ProviderTransaction struct {
	ID            int    `gorm:"primaryKey;autoIncrement"`
	PaymentID     int    `gorm:"not null;index"`
	ProviderRef   string `gorm:"type:varchar(255);uniqueIndex"`
	AttemptNumber int    `gorm:"not null"`
	Status        string `gorm:"type:varchar(20);not null"`
	FailureReason string `gorm:"type:text"`
	CreatedAt     time.Time

	Payment Payment `gorm:"foreignKey:PaymentID;references:ID"`
}