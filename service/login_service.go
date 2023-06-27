package service

import (
	"errors"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
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
		return "", errors.New("invalid phone number format")
	}

	// check existing user
	existUser, err := l.loginRepo.GetUser(user)
	if err != nil {
		return "", err
	}

	authDuration, _ := strconv.Atoi(utility.DotEnv("AUTH_DURATION", ".env"))

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = existUser.Username
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(authDuration)).Unix()

	signedToken, err := token.SignedString([]byte(utility.DotEnv("TOKEN_KEY", ".env")))
	if err != nil {
		return "", err
	}

	return signedToken, nil

}
