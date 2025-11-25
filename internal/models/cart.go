package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID     uint `json:"user_id"`
	Items      []CartItem
	TotalPrice int64 `json:"total_price"`
}

type UpdateCart struct {
	UserID     uint       `json:"user_id"`
	Items      []CartItem `json:"items"`
	TotalPrice int64      `json:"total_price"`
}
