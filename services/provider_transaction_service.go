package services

import (
	"errors"
	"fmt"
	"payment_reconciliation/dto"
	"payment_reconciliation/models"
	"payment_reconciliation/repository"

	"gorm.io/gorm"
)

type ProviderTransactionServices struct {
	DB                      *gorm.DB
	ProviderTransactionRepo *repository.ProviderTransactionRepository
	PaymentRepo             *repository.PaymentRepository
}

func (s *ProviderTransactionServices) ProcessPayment(paymentID int, result string) (*dto.ProviderTransactionResponse, error) {

	// find the payment
	payment, err := s.PaymentRepo.FindByID(paymentID)
	if err != nil {
		return nil, err
	}
	// 2. Don't process an already successful payment
	if payment.Status == "successful" {
		return nil, errors.New("payment already successful")
	}

	// 3. Find previous attempts
	attempts, err := s.ProviderTransactionRepo.FindByPaymentId(paymentID)
	if err != nil {
		return nil, err
	}
	// fail the retry if attempt is >= 3
	if len(attempts) >= 3 {
		return nil, errors.New("maximum payment attempts reached")
	}

	attemptNumber := len(attempts) + 1

	// 4. Simulate provider processing
	providerRef := fmt.Sprintf("TXN-%d-%d", paymentID, attemptNumber)

	transaction := &models.ProviderTransaction{
		PaymentID:     paymentID,
		ProviderRef:   providerRef,
		AttemptNumber: attemptNumber,
		Status:        result,
		FailureReason: "",
	}

	if result == "fail" {
		transaction.FailureReason = "Failed by force"
	}

	err = s.DB.Transaction(func(tx *gorm.DB) error {
		if err := s.ProviderTransactionRepo.CreateTransaction(tx, transaction); err != nil {
			return err
		}

		if result == "success" {
			payment.Status = "successful"
			return s.PaymentRepo.UpdatePayment(tx, payment)
		}

		if result == "fail" {
			payment.Status = "Failed"
			return s.PaymentRepo.UpdatePayment(tx, payment)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &dto.ProviderTransactionResponse{
		ID:            transaction.ID,
		PaymentID:     transaction.PaymentID,
		ProviderRef:   transaction.ProviderRef,
		AttemptNumber: transaction.AttemptNumber,
		Status:        transaction.Status,
		FailureReason: transaction.FailureReason,
	}, nil

}
