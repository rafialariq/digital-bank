package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rafialariq/digital-bank/models"
	"github.com/rafialariq/digital-bank/models/dto"
	"github.com/rafialariq/digital-bank/repository"
	"github.com/rafialariq/digital-bank/utility"
)

type RegisterService interface {
	CreateUser(*dto.RegisterDTO) (string, error)
}

type registerService struct {
	registerRepo repository.RegisterRepository
}

func NewRegisterService(registerRepo repository.RegisterRepository) RegisterService {
	return &registerService{
		registerRepo: registerRepo,
	}
}

func (r *registerService) CreateUser(user *dto.RegisterDTO) (string, error) {

	// password matching
	if user.Password != user.PasswordConfirm {
		return "password does not match", errors.New("password does not match")
	}

	// check if user already exist
	if r.registerRepo.FindExistingUser(user) {
		return "user already exist", errors.New("user already exist")
	}

	// field validation
	if utility.IsUsernameInvalid(user.Username) {
		return "invalid username", errors.New("invalid username")
	} else if utility.IsPhoneInvalid(user.PhoneNumber) {
		return "invalid phone number", errors.New("invalid phone number")
	} else if utility.IsEmailInvalid(user.Email) {
		return "invalid email", errors.New("invalid email")
	} else if utility.IsPasswordInvalid(user.Password) {
		return "invalid password", errors.New("invalid password")
	}

	// generate hashed password
	hashedPass := utility.PasswordHashing(user.Password)

	// assign dto.RegisterDTO to model.User
	newUser := models.User{
		Id:          uuid.New(),
		Username:    user.Username,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Password:    hashedPass,
		Balanced:    0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := r.registerRepo.InsertUser(&newUser)
	if err != nil {
		return "", err
	}

	return "new user created successfully", nil
}
