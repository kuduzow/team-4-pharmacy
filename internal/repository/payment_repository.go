package repository

import (
	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(payment *models.Payment) error

	GetByID(id uint) error

	Update(payment *models.Payment) error

	Delete(id uint) error
}

type gormPaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &gormPaymentRepository{db: db}
}

func (r *gormPaymentRepository) Create(payment *models.Payment) error {
	if payment == nil {
		return nil
	}
	return r.db.Create(payment).Error
}
func (r *gormPaymentRepository) GetByID(id uint) error {
	var payment models.Payment

	if err := r.db.First(&payment, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *gormPaymentRepository) Update(payment *models.Payment) error {
	if payment == nil {
		return nil
	}
	return r.db.Save(payment).Error
}
func (r *gormPaymentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Payment{}, id).Error
	
}
