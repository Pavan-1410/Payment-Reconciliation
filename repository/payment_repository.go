package repository

import (
	"payment_reconciliation/models"

	"gorm.io/gorm"
)

type PaymentRepository struct {
	DB *gorm.DB
}

func (r *PaymentRepository) CreatePayment (payment *models.Payment) error {
	return  r.DB.Create(payment).Error
}

// by idempotency key
func (r *PaymentRepository) FindPayment (idempotency_key string) (*models.Payment, error){	//by idempotency key
	var payment models.Payment

	err := r.DB.Where("idempotency_key = ?", idempotency_key).First(&payment).Error

	if err != nil{
		return nil,err
	}

	return &payment,nil
}

// by payment ID
func (r *PaymentRepository) FindByID (paymentID int) (*models.Payment, error) {
	var payment models.Payment
	err := r.DB.
		First(&payment, paymentID).
		Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *PaymentRepository) UpdatePayment (tx *gorm.DB,payment *models.Payment) error {
	return tx.Save(payment).Error
}