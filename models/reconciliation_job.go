package models

import "time"
// it stores the information of reconciliation job
type ReconciliationJob struct {
	ID          int        `gorm:"primaryKey;autoIncrement"`
	Status      string     `gorm:"type:varchar(20);not null;default:'pending';index"`
	StartedAt   *time.Time
	CompletedAt *time.Time
	Error       string     `gorm:"type:text"`	// failed to load provider transactions	Technical failure of the reconciliation process
	CreatedAt   time.Time
	UpdatedAt   time.Time
}