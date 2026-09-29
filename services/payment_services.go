package services

import (
	"errors"
	"payment_reconciliation/dto"
	"payment_reconciliation/models"
	"payment_reconciliation/repository"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type PaymentService struct {	
	PaymentRepo *repository.PaymentRepository
}

func (s *PaymentService) CreatePayment (userID int,req dto.CreatePaymentRequest)( *dto.PaymentResponce, error){
		
	// 1. Convert amount string to decimal
	amount, err := decimal.NewFromString(req.Amount)	// here we are convertign to decimal

	if err != nil {
		return nil, errors.New("invalid amount")
	}
		
	// 2. Amount must be greater than zero
	if amount.LessThanOrEqual(decimal.Zero) {		// decimal is a struct type not int of float
		return nil, errors.New("amount must be greater than zero")
	}
	// check that payment by idempotency key

	existingPayment, err := s.PaymentRepo.FindPayment(req.IdempotencyKey)

	if err == nil {
		// payment alrady exists
		return &dto.PaymentResponce{
			ID: existingPayment.ID,
			UserID: existingPayment.UserID,
			Amount: existingPayment.Amount.StringFixed(2),
			Status: existingPayment.Status,
			IdempotencyKey: existingPayment.IdempotencyKey,
		},nil
	}
	
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// create new payment
	payment := models.Payment{
		UserID: userID,
		Amount: amount,
		Status: "pending",
		IdempotencyKey: req.IdempotencyKey,
	}

	err = s.PaymentRepo.CreatePayment(&payment)


	
	
	if err != nil {
		// Another concurrent request may have created the payment with the same idempotency key.
		existingPayment, findErr := s.PaymentRepo.FindPayment(req.IdempotencyKey)
			
		if findErr == nil {
		return &dto.PaymentResponce{
			ID:             existingPayment.ID,
			UserID:         existingPayment.UserID,
			Amount:         existingPayment.Amount.StringFixed(2),
			Status:         existingPayment.Status,
			IdempotencyKey: existingPayment.IdempotencyKey,
		}, nil
	}

		return nil, err
	}
	// 5. Convert model to response DTO
	return &dto.PaymentResponce{
		ID:             payment.ID,
		UserID:         payment.UserID,
		Amount:         payment.Amount.StringFixed(2),
		Status:         payment.Status,
		IdempotencyKey: payment.IdempotencyKey,
	}, nil

}