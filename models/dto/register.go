package dto

type RegisterDTO struct {
	Username        string `json:"username"`
	PhoneNumber     string `json:"phone_number"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}
