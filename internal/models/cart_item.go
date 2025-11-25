package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	CartID       uint   `json:"cart_id"`
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
