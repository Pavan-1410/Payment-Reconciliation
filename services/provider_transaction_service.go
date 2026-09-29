package services

import (
	"errors"
	"fmt"
	"payment_reconciliation/dto"
	"payment_reconciliation/models"
	"payment_reconciliation/repository"
)

type ProviderTransactionServices struct {
	ProviderTransactionRepo *repository.ProviderTransactionRepository
	PaymentRepo *repository.PaymentRepository
}

func (s * ProviderTransactionServices) 	ProcessPayment (paymentID int,result string) (*dto.ProviderTransactionResponse,error){
	// find the payment
	payment, err := s.PaymentRepo.FindByID(paymentID)
	if err != nil {
		return nil, err
	}
	// 2. Don't process an already successful payment
	if payment.Status == "Successful" {
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
		PaymentID: paymentID,
		ProviderRef: providerRef,
		AttemptNumber: attemptNumber,
		Status: result,
		FailureReason: "",
	}

	// transfer the payment
	err = s.ProviderTransactionRepo.CreateTransaction(transaction)

	if err != nil{
		return nil,err
	}

	// Update overall payment status

	if result == "success"{
		payment.Status = "Successful"
		err = s.PaymentRepo.UpdatePayment(payment)
	}

	if result == "fail"{
		payment.Status = "Failed"
		transaction.FailureReason = "Failed by force"
		err = s.PaymentRepo.UpdatePayment(payment)
	}


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
