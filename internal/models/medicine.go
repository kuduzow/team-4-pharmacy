package models

import "gorm.io/gorm"

type Medicine struct {
	gorm.Model
	Name                 string       `json:"name"`
	Description          string       `json:"description"`
	Price                float64      `json:"price"`
	InStock              bool         `json:"in_stock"`
	StockQuantity        int          `json:"stock_quantity"`
	CategoryID           uint         `json:"category_id" gorm:"not null;index"`
	Category             *Category    `json:"-"`
	SubcategoryID        uint         `json:"subcategory_id" gorm:"not null;index"`
	Subcategory          *Subcategory `json:"-"`
	Manufacturer         string       `json:"manufacturer"`
	PrescriptionRequired bool         `json:"prescription_required"`
	AvgRating            float64      `json:"avg_rating"`
}

type MedicineCreateRequest struct {
	Name                 string  `json:"name"`
	Description          string  `json:"description"`
	Price                float64 `json:"price"`
	InStock              bool    `json:"in_stock"`
	StockQuantity        int     `json:"stock_quantity"`
	CategoryID           uint    `json:"category_id"`
	SubcategoryID        uint    `json:"subcategory_id"`
	Manufacturer         string  `json:"manufacturer"`
	PrescriptionRequired bool    `json:"prescription_required"`
	AvgRating            float64 `json:"avg_rating"`
}

type MedicineUpdateRequest struct {
	Name                 *string  `json:"name"`
	Description          *string  `json:"description"`
	Price                *float64 `json:"price"`
	InStock              *bool    `json:"in_stock"`
	StockQuantity        *int     `json:"stock_quantity"`
	Manufacturer         *string  `json:"manufacturer"`
	PrescriptionRequired *bool    `json:"prescription_required"`
	AvgRating            *float64 `json:"avg_rating"`
}
