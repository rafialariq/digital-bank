package service

import (
	"errors"

	"github.com/rafialariq/digital-bank/models/dto"
	"github.com/rafialariq/digital-bank/repository"
	"github.com/rafialariq/digital-bank/utility"
)

type LoginService interface {
	FindUser(*dto.LoginDTO) (string, error)
}

type loginService struct {
	loginRepo repository.LoginRepository
}

func NewLoginService(loginRepo repository.LoginRepository) LoginService {
	return &loginService{
		loginRepo: loginRepo,
	}
}

func (l *loginService) FindUser(user *dto.LoginDTO) (string, error) {

	// check phone number format
	if utility.IsPhoneInvalid(user.PhoneNumber) {
		return "invalid phone number", errors.New("invalid phone number format")
	}

	// check existing user
	existUser, err := l.loginRepo.GetUser(user)
	if err != nil {
		return "failed to check existing user", err
	}

	// generate token
	signedToken, err := utility.GenerateJWTToken(existUser.Id)
	if err != nil {
		return "failed to generate token", err
	}

	return signedToken, nil
}
