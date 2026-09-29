package repository

import (
	"payment_reconciliation/models"

	"gorm.io/gorm"
)

type ProviderTransactionRepository struct {
	DB *gorm.DB
}

func (r *ProviderTransactionRepository) CreateTransaction (transaction *models.ProviderTransaction) error{
	return r.DB.Create(transaction).Error
}

func (r *ProviderTransactionRepository) FindByPaymentId (paymentId int) ([]models.ProviderTransaction, error)	{//we not use * ptrr for []model.slice because A slice already behaves like a descriptor pointing to an underlying array:
	var transaction []models.ProviderTransaction
	
	err := r.DB.Where("payment_id=?",paymentId).
	Order("attempt_number ASC").
	Find(&transaction).Error

	if err != nil{
		return nil,err
	}

	return transaction,nil
	}	
	
func (r *ProviderTransactionRepository) FindByProviderRef(providerRef string) (*models.ProviderTransaction, error) {

	var transaction models.ProviderTransaction

	err := r.DB.
		Preload("Payment").
		Where("provider_ref = ?", providerRef).
		First(&transaction).Error

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}