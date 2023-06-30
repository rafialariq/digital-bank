package service

import (
	"github.com/rafialariq/digital-bank/models/dto"
	"github.com/rafialariq/digital-bank/repository"
)

type PaymentService interface {
	MakePayment(dto.PaymentDTO) error
}

type paymentService struct {
	paymentRepo repository.PaymentRepository
}

func NewPaymentService(payment repository.PaymentRepository) PaymentService {
	return &paymentService{
		paymentRepo: payment,
	}
}

func (p *paymentService) MakePayment(payment dto.PaymentDTO) error {
	return p.paymentRepo.MakePayment(payment)
}
