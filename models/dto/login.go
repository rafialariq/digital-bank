package dto

type LoginDTO struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
