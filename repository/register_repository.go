package repository

import (
	"fmt"

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

func NewRegisterRepository(db *gorm.DB) RegisterRepository {
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
		if err == gorm.ErrRecordNotFound {
			return false
		}

		fmt.Println(err)
	}

	return true
}
