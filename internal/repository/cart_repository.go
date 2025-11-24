package repository

import (
	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"gorm.io/gorm"
)

type CartRepository interface {
	Create(item *models.CartItem) error

	GetByID(itemID uint) (*models.CartItem, error)

	Update(item *models.CartItem) error

	Delete(itemID uint) error
}

type gormCartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &gormCartRepository{db: db}
}

func (r *gormCartRepository) Create(item *models.CartItem) error {
	if item == nil {
		return nil
	}
	return r.db.Create(item).Error
}

func (r *gormCartRepository) GetByID(itemID uint) (*models.CartItem, error) {
	var item models.CartItem
	if err := r.db.First(&item, itemID).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *gormCartRepository) Update(item *models.CartItem) error {
	if item == nil {
		return nil
	}
	return r.db.Save(item).Error
}

func (r *gormCartRepository) Delete(itemID uint) error {
	return r.db.Delete(&models.CartItem{}, itemID).Error
}
