package models

import (
	"time"

	"gorm.io/gorm"
)

type DiscountType string

const (
	DiscountTypeFixed   DiscountType = "fixed"
	DiscountTypePercent DiscountType = "percent"
)

type Promocode struct {
	gorm.Model
	Code           string       `json:"code"`
	Description    string       `json:"description"`
	DiscountType   DiscountType `json:"discount_type"`
	DiscountValue  float64      `json:"discount_value"`
	ValidFrom      time.Time    `json:"valid_from"`
	ValidTo        time.Time    `json:"valid_to"`
	MaxUses        *int         `json:"max_uses"`
	MaxUsesPerUser *int         `json:"max_uses_per_user"`
	IsActive       bool         `json:"is_active"`
}

type PromocodeCreateRequest struct {
	Code           string       `json:"code"`
	Description    string       `json:"description"`
	DiscountType   DiscountType `json:"discount_type"`
	DiscountValue  float64      `json:"discount_value"`
	ValidFrom      time.Time    `json:"valid_from"`
	ValidTo        time.Time    `json:"valid_to"`
	MaxUses        *int         `json:"max_uses"`
	MaxUsesPerUser *int         `json:"max_uses_per_user"`
	IsActive       bool         `json:"is_active"`
}

type PromocodeUpdateRequest struct {
	Code           *string       `json:"code"`
	Description    *string       `json:"description"`
	DiscountType   *DiscountType `json:"discount_type"`
	DiscountValue  *float64      `json:"discount_value"`
	ValidFrom      *time.Time    `json:"valid_from"`
	ValidTo        *time.Time    `json:"valid_to"`
	MaxUses        *int          `json:"max_uses"`
	MaxUsesPerUser *int          `json:"max_uses_per_user"`
	IsActive       *bool         `json:"is_active"`
}
