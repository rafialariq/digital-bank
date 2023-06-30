package repository

import (
	"errors"

	"github.com/rafialariq/digital-bank/models"
	"github.com/rafialariq/digital-bank/models/dto"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginRepository interface {
	GetUser(*dto.LoginDTO) (*models.User, error)
}

type loginRepository struct {
	db *gorm.DB
}

func NewLoginRepository(db *gorm.DB) LoginRepository {
	return &loginRepository{db}
}

func (l *loginRepository) GetUser(user *dto.LoginDTO) (*models.User, error) {
	var existUser models.User

	err := l.db.Where("phone_number = ?", user.PhoneNumber).First(&existUser).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &existUser, errors.New("user not found")
		}

		return &existUser, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(user.Password))
	if err != nil {
		return &existUser, errors.New("username or password is not valid")
	}

	return &existUser, nil
}
