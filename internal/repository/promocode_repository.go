package repository

import (
	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"gorm.io/gorm"
)

type PromocodeRepository interface {
	Create(promocode *models.Promocode) error

	GetAll() ([]models.Promocode, error)

	Delete(id uint) error

	Update(promocode *models.Promocode) error
}

type gormPromocodeRepository struct {
	db *gorm.DB
}

func NewPromocodeRepository(db *gorm.DB) PromocodeRepository {
	return &gormPromocodeRepository{db: db}
}

func (r *gormPromocodeRepository) Create(promocode *models.Promocode) error {
	if promocode == nil {
		return nil
	}

	return r.db.Create(&promocode).Error
}

func (r *gormPromocodeRepository) GetAll() ([]models.Promocode, error) {
	var promocode []models.Promocode

	if err := r.db.Find(&promocode).Error; err != nil {
		return nil, err
	}

	return promocode, nil
}

func (r *gormPromocodeRepository) Delete(id uint) error {
	var promocode models.Promocode

	if err := r.db.Delete(&promocode, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *gormPromocodeRepository) Update(promocode *models.Promocode) error {
	if promocode == nil {
		return nil
	}

	return r.db.Save(&promocode).Error
}
