package repository

import (
	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"gorm.io/gorm"
)

type CartItemRepository interface {
	Create(item *models.CartItem) error

	GetItemsByCartID(cartID uint) ([]models.CartItem, error)

	GetCartItemByMedID(id uint) (*models.CartItem, error)

	Update(item *models.CartItem) error

	Delete(id uint) error
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
func (r *gormCartItemRepository) GetCartItemByMedID(medicineID uint) (*models.CartItem, error) {
	var cartItem models.CartItem

	err := r.db.Where("medicine_id = ?", medicineID).First(&cartItem).Error
	if err != nil {
		return nil, err
	}

	return &cartItem, nil
}
func (r *gormCartItemRepository) GetItemsByCartID(cartID uint) ([]models.CartItem, error) {

	var items []models.CartItem

	if err := r.db.Where("cart_id = ?", cartID).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *gormCartItemRepository) Update(item *models.CartItem) error {

	if item == nil {
		return nil
	}
	return r.db.Save(item).Error
}

func (r *gormCartItemRepository) Delete(id uint) error {
	return r.db.Delete(&models.CartItem{}, id).Error
}
