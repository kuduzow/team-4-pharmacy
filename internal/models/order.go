package models

import "gorm.io/gorm"

type OrderStatus string

const (
	Draft          OrderStatus = "draft"
	PendingPayment OrderStatus = "pending_payment"
	Paid           OrderStatus = "paid"
	Canceled       OrderStatus = "canceled"
	Shipped        OrderStatus = "shipped"
	Completed      OrderStatus = "completed"
)

type Order struct {
	gorm.Model
	UserID          uint        `json:"user_id"`
	OrderStatus     OrderStatus `json:"order_status"`
	TotalPrice      int         `json:"total_price"`
	DiscountTotal   int         `json:"discountTotal"`
	FinalPrice      int         `json:"final_price"`
	DeliveryAddress string      `json:"delivery_address"`
	Comment         string      `json:"comment"`
}
type OrderItem struct {
	OrderID      uint   `json:"order_id"`
	MedicineID   uint   `json:"medicine_id"`
	MedicineName string `json:"medicine_name"`
	Quantity     int    `json:"quantity"`
	PricePerUnit uint   `json:"price_per_uint"`
	LineTotal    string `json:"line_total"`
}

type OrderCreate struct {
	UserID          uint        `json:"user_id"`
	OrderStatus     OrderStatus `json:"order_status"`
	TotalPrice      int         `json:"total_price"`
	DiscountTotal   int         `json:"discountTotal"`
	FinalPrice      int         `json:"final_price"`
	DeliveryAddress string      `json:"delivery_address"`
	Comment         string      `json:"comment"`
}
type OrderUpdate struct {
	UserID          *uint        `json:"user_id"`
	OrderStatus     *OrderStatus `json:"order_status"`
	TotalPrice      *int         `json:"total_price"`
	DiscountTotal   *int         `json:"discountTotal"`
	FinalPrice      *int         `json:"final_price"`
	DeliveryAddress *string      `json:"delivery_address"`
	Comment         *string      `json:"comment"`
}
