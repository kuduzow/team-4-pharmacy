package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	Phone          int    `json:"phone"`
	DefaultAddress string `json:"default_address"`
	Cart           Cart
	Orders         []Order
	Reviews        []Review
}

type UserCreateRequest struct {
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	Phone          int    `json:"phone"`
	DefaultAddress string `json:"default_address"`
}

type UserUpdateRequest struct {
	FullName       *string `json:"full_name"`
	Email          *string `json:"email"`
	Phone          *int    `json:"phone"`
	DefaultAddress *string `json:"default_address"`
}
