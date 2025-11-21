package models

import "gorm.io/gorm"

type Status string

type Method string

const (
	Card         Method = "card"
	Cash         Method = "cash"
	OnlineWallet Method = "online_wallet"
)


const (
	Pending Status = "pending"
	Succes  Status = "succes"
	Failed  Status = "failed"
)

type Payment struct {
	gorm.Model
	OrderID int    `json:"order_ID"`
	Amount  int    `json:"amount"`
	Status  Status `json:"-"`
	Method  Method `json:"-"`
	PaidAt  string `json:"paid_at"`
}

type PaymentCreate struct {
	OrderID int    `json:"order_ID"`
	Amount  int    `json:"amount"`
	Status  Status `json:"status"`
	Method  Method `json:"method"`
	PaidAt  string `json:"paid_at"`
}
type PaymentUpdate struct {
	Amount *int    `json:"amount"`
	Status *Status `json:"status"`
	Method *Method `json:"method"`
	PaidAt *string `json:"paid_at"`
}
