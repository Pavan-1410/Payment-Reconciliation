package services

import (
	"errors"

	"payment_reconciliation/dto"
	"payment_reconciliation/models"
	"payment_reconciliation/repository"
	"payment_reconciliation/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	UserRepo *repository.UserRepository
}

func (s *AuthService) Register(req dto.RegisterRequest) (*dto.UserResponse, error) {	// req is just a varaibal

	// Check whether email already exists
	_, err := s.UserRepo.FindByEmail(req.Email)

	if err == nil {
		return nil, errors.New("email already registered")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, err
	}

	// Create user
	user := models.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
	}

	err = s.UserRepo.CreateUser(&user)

	if err != nil {
		return nil, err
	}

	// Return response without password
	return &dto.UserResponse{
		ID:      int(user.ID),
		Name:    user.Name,
		Email:   user.Email,
		IsAdmin: user.IsAdmin,
	}, nil
}

func (s * AuthService) Login(req dto.LoginRequest) (*dto.LoginResponse,error){

	// 1. Find user by email
	user, err := s.UserRepo.FindByEmail(req.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("User not exists")
		}

		return nil, err
	}

	// 2. Compare password with stored hash
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	)
		
	if err != nil {
		return nil, errors.New("invalid password")
	}

	// generate token

	token, err := utils.GenerateToken(int(user.ID),user.IsAdmin)

	if err != nil{
		return nil, err
	}

	// create login responce

	return &dto.LoginResponse{
		Token: token,
		User: dto.UserResponse{
			ID: int(user.ID),
			Name: user.Name,
			Email: user.Email,
			IsAdmin: user.IsAdmin,
		},
	},nil

}