package repository

import (
	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"gorm.io/gorm"
)

type CartItemRepository interface {
	Create(item *models.CartItem) error

	GetByID(itemID uint) (*models.CartItem, error)

	Update(item *models.CartItem) error

	Delete(itemID uint) error
}

type gormCartItemRepository struct {
	db *gorm.DB
}

func NewCartItemRepository(db *gorm.DB) CartItemRepository {
	return &gormCartItemRepository{db: db}
}

func (r *gormCartItemRepository) Create(item *models.CartItem) error {
	if item == nil {
		return nil
	}
	return r.db.Create(item).Error
}

func (r *gormCartItemRepository) GetByID(itemID uint) (*models.CartItem, error) {

	var item models.CartItem

	if err := r.db.First(&item, itemID).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *gormCartItemRepository) Update(item *models.CartItem) error {

	if item == nil {
		return nil
	}
	return r.db.Save(item).Error
}

func (r *gormCartItemRepository) Delete(itemID uint) error {
	return r.db.Delete(&models.CartItem{}, itemID).Error
}
