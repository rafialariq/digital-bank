package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Username    string    `gorm:"unique;not null" json:"username"`
	PhoneNumber string    `gorm:"column:phone_number;unique;not null" json:"phone_number"`
	Email       string    `gorm:"not null" json:"email"`
	Password    string    `gorm:"not null" json:"password"`
	Balanced    float64   `gorm:"default:0" json:"balanced"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}
