package repository

import (
	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"gorm.io/gorm"
)

type MedicineRepository interface {
	Create(medicine *models.Medicine) error

	GetByID(id uint) (*models.Medicine, error)

	Delete(id uint) error

	Update(medicine *models.Medicine) error

	GetAll() ([]models.Medicine, error)

	GetInStock() ([]models.Medicine, error)
}

type gormMedecineRepository struct {
	db *gorm.DB
}

func NewMedicineRepository(db *gorm.DB) MedicineRepository {
	return &gormMedecineRepository{db: db}
}

func (r *gormMedecineRepository) Create(medicine *models.Medicine) error {
	if medicine == nil {
		return nil
	}

	return r.db.Create(&medicine).Error
}

func (r *gormMedecineRepository) GetByID(id uint) (*models.Medicine, error) {
	var medicine models.Medicine

	if err := r.db.First(&medicine, id).Error; err != nil {
		return nil, err
	}

	return &medicine, nil
}

func (r *gormMedecineRepository) Delete(id uint) error {
	var medicine models.Medicine

	if err := r.db.Delete(&medicine, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *gormMedecineRepository) Update(medicine *models.Medicine) error {
	if medicine == nil {
		return nil
	}

	return r.db.Save(&medicine).Error
}

func (r *gormMedecineRepository) GetAll() ([]models.Medicine, error) {
	var medicines []models.Medicine

	if err := r.db.Find(&medicines).Error; err != nil {
		return nil, err
	}
	return medicines, nil
}

func (r *gormMedecineRepository) GetInStock() ([]models.Medicine, error) {
	var medicines []models.Medicine

	if err := r.db.Where("in_stock = ?", true).Find(&medicines).Error; err != nil {
		return nil, err
	}

	return medicines, nil
}
