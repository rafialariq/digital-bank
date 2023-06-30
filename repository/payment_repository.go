package repository

import (
	"fmt"

	"github.com/rafialariq/digital-bank/models"
	"github.com/rafialariq/digital-bank/models/dto"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	MakePayment(dto.PaymentDTO) error
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db}
}

func (p *paymentRepository) MakePayment(payment dto.PaymentDTO) error {
	tx := p.db.Begin()

	var user models.User
	if err := tx.Where("phone_number = ?", payment.SenderCode).First(&user).Error; err != nil {
		tx.Rollback()
		fmt.Println("eroor 1")
		return err
	}

	if user.Balanced < payment.Amount {
		tx.Rollback()
		fmt.Println("eroor 2")
		return fmt.Errorf("insufficient balance")
	}

	user.Balanced -= payment.Amount
	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		fmt.Println("eroor 3")
		return err
	}

	var merchant models.Merchant
	if err := tx.Where("merchant_code = ?", payment.RecepientCode).First(&merchant).Error; err != nil {
		tx.Rollback()
		fmt.Println("eroor 4")
		return err
	}

	merchant.Balanced += payment.Amount
	if err := tx.Model(&merchant).Where("merchant_code = ?", merchant.MerchantCode).Update("balanced", merchant.Balanced).Error; err != nil {
		tx.Rollback()
		fmt.Println("eroor 5")
		fmt.Println(err)
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		fmt.Println("eroor 6")
		return err
	}

	return nil
}
