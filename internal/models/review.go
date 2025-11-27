package models

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	UserID     uint   `json:"user_id"`
	MedicineID uint   `json:"medicine_id"`
	Rating     uint   `json:"rating"`
	Text       string `json:"text"`
}

type ReviewForPost struct {
	UserID     uint   `json:"user_id" binding:"required"`
	MedicineID uint   `json:"medicine_id" binding:"required"`
	Rating     uint   `json:"rating" binding:"required,min=1,max=5"`
	Text       string `json:"text" binding:"required"`
}

type ReviewForUpdate struct {
	Rating *uint   `json:"rating" binding:"omitempty,min=1,max=5"`
	Text   *string `json:"text" binding:"omitempty"`
}
