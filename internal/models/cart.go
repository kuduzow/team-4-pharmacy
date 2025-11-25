package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID     uint `json:"user_id"`
	Items      []CartItem
	TotalPrice int64 `json:"total_price"`
}

type CartItem struct {
	gorm.Model
	ItemID       uint   `json:"item_id"`
	MedicineID   uint   `json:"medicine_id"`
	Name         string `json:"name"`
	Quantity     int64  `json:"quantity"`
	PricePerUnit int64  `json:"price_per_unit"`
	LineTotal    int64  `json:"line_total"`
}

type CartCreateItemRequest struct {
	MedicineID uint  `json:"medicine_id"`
	Quantity   int64 `json:"quantity"`
}

type UpdateCartItemRequest struct {
	Quantity int64 `json:"quantity"`
}
