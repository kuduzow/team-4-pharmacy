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
	MedicineID uint   `json:"medicine_id"`
	Rating     uint   `json:"rating"`
	Text       string `json:"text"`
}

type ReviewForUpdate struct {
	Rating *uint   `json:"rating"`
	Text   *string `json:"text"`
}
