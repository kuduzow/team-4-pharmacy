package repository

import (
	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *models.Order) error

	GetByID(id uint) (*models.Order, error)

	GetByUserID(userID uint) (*models.Order, error)

	Update(order *models.Order) error

	Delete(id uint) error
}

type gormOrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &gormOrderRepository{db: db}
}

func (r *gormOrderRepository) Create(order *models.Order) error {
	if order == nil {
		return nil
	}
	return r.db.Create(order).Error
}
func (r *gormOrderRepository) GetByID(id uint) (*models.Order, error) {
	var order models.Order
	if err := r.db.First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
func (r *gormOrderRepository) GetByUserID(userID uint) (*models.Order, error) {
	var order models.Order

	if err := r.db.Where("user_id = ?", userID).Find(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
func (r *gormOrderRepository) Update(order *models.Order) error {
	if order == nil {
		return nil
	}
	return r.db.Save(order).Error
}
func (r *gormOrderRepository) Delete(id uint) error {
	return r.db.Delete(&models.Order{}).Error
}
