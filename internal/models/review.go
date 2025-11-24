<<<<<<< HEAD
=======
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
	UserID     uint   `json:"user_id"`
	MedicineID uint   `json:"medicine_id"`
	Rating     uint   `json:"rating"`
	Text       string `json:"text"`
}

type ReviewForUpdate struct {
	Rating *uint   `json:"rating"`
	Text   *string `json:"text"`
}
>>>>>>> 6ef08ef05f598c9ed5de0f70de6984d0a7880013
