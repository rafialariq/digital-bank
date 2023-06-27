package repository

import (
	"errors"

	"github.com/rafialariq/digital-bank/models"
	"github.com/rafialariq/digital-bank/models/dto"
	"gorm.io/gorm"
)

type RegisterRepository interface {
	InsertUser(*models.User) error
	FindExistingUser(*dto.RegisterDTO) bool
}

type registerRepository struct {
	db *gorm.DB
}

func NewRegisterRepo(db *gorm.DB) RegisterRepository {
	return &registerRepository{db}
}

func (r *registerRepository) InsertUser(user *models.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *registerRepository) FindExistingUser(user *dto.RegisterDTO) bool {
	var existUser models.User
	err := r.db.Where("username = ? OR phone_number = ?", user.Username, user.PhoneNumber).First(&existUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true
		}

		// logging here
		// not completed
	}

	return false
}
