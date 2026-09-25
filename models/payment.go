package models
import (
	"time"

	"github.com/shopspring/decimal"
)
type Payment struct {
	ID             int             `gorm:"primaryKey;autoIncrement"`
	UserID         int             `gorm:"not null;index"`
	Amount         decimal.Decimal `gorm:"type:numeric(12,2);not null"`
	Status         string          `gorm:"type:varchar(20);not null;default:'pending'"`
	IdempotencyKey string          `gorm:"type:varchar(255);uniqueIndex;not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time

	User User `gorm:"foreignKey:UserID;references:ID"`
}
