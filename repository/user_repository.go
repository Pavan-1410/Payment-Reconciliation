package repository

import (
	"payment_reconciliation/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {	// constructor function it  create and initialize a UserRepository.
	return &UserRepository{		// it returns the address of user repository struct 
		DB: db,
	}
}
// create user method	-- returns error
func (r *UserRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

// find the user method	-- returns user,error
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}