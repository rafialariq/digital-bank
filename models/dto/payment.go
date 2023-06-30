package dto

type PaymentDTO struct {
	SenderCode    string  `json:"sender_code"`
	RecepientCode int     `json:"recepient_code"`
	Amount        float64 `json:"amount"`
}
