package models

import (
	"time"

	"github.com/shopspring/decimal"
)

//  here we store report of payments that we get from payment provider
type ProviderReportTransaction struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	ProviderRef string    `gorm:"type:varchar(255);not null;index"`
	Amount decimal.Decimal `gorm:"type:numeric(12,2);not null"`
	Status      string    `gorm:"type:varchar(20);not null"`
	ReportID    string    `gorm:"type:varchar(255);not null;index"`
	CreatedAt   time.Time
}