package models

type Merchant struct {
	MerchantCode int     `gorm:"" json:"merchant_code"`
	MerchantName string  `gorm:"" json:"merchant_name"`
	Balanced     float64 `gorm:"" json:"balanced"`
}
