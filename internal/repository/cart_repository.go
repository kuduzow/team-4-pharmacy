package repository

import (
	"errors"

	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"gorm.io/gorm"
)

type CartRepository interface {
	Create(cart *models.Cart) error

	GetByUserID(userID uint) (*models.Cart, error)

	ClearByUserID(userID uint) error
}

type gormCartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &gormCartRepository{db: db}
}

func (r *gormCartRepository) Create(cart *models.Cart) error {
	return r.db.Create(cart).Error
}

func (r *gormCartRepository) GetByUserID(userID uint) (*models.Cart, error) {

	var cart models.Cart

	if err := r.db.Preload("Items").Where("user_id = ?", userID).First(&cart).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}

func (r *gormCartRepository) ClearByUserID(userID uint) error {

	var cart models.Cart

	if err := r.db.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return r.db.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{}).Error
}
