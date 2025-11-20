package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `json:"name"`
}

type Subcategory struct {
	gorm.Model
	CategoryID uint   `json:"category_id"`
	Name       string `json:"name"`
}

type CreateCategory struct {
	Name string `json:"name"`
}

type CreateSubcategory struct {
	CategoryID uint   `json:"category_id"`
	Name       string `json:"name"`
}

type UpdateCategory struct {
	Name *string `json:"name"`
}

type UpdateSubcategory struct {
	CategoryID *uint   `json:"category_id"`
	Name       *string `json:"name"`
}
