package models

type Status string

type Method string

const (
	Card          Method = "card"
	Cash          Method = "cash"
	Online_Wallet Method = "onlineWallet"
)

const (
	Pending Status = "pending"
	Succes  Status = "succes"
	Failed  Status = "failed"
)

type Payment struct {
	gorm.models
	Order_ID int    `json:"order_ID"`
	Amount   int    `json:"amount"`
	Status   Status `json:"-"`
	Method   Method `json:"-"`
	Paid_at  string `json:"paid_at"`
}

type PaymentCreate struct {
	Order_ID int    `json:"order_ID"`
	Amount   int    `json:"amount"`
	Status   Status `json:"status"`
	Method   Method `json:"method"`
	Paid_at  string `json:"paid_at"`
}
type PaymentUpdate struct {
	Amount  *int    `json:"amount"`
	Status  *Status `json:"status"`
	Method  *Method `json:"method"`
	Paid_at *string `json:"paid_at"`
}
