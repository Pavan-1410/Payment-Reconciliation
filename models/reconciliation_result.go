package models

import "time"
// it stores the comparison of our DB payment data and provider report individual result of each comparison
type ReconciliationResult struct {
	ID                      int       `gorm:"primaryKey;autoIncrement"`
	ReconciliationJobID     int       `gorm:"not null;index"`
	ProviderReportTransactionID int `gorm:"not null;index"`
	ProviderTransactionID   *int      `gorm:"index"`
	Result                  string    `gorm:"type:varchar(50);not null"`
	Details                 string    `gorm:"type:text"`
	CreatedAt               time.Time

	Job ReconciliationJob `gorm:"foreignKey:ReconciliationJobID;references:ID"`
}