package repository

import (
	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"gorm.io/gorm"
)

type MedicineFilter struct {
	CategoryID    *uint
	SubcategoryID *uint
	InStock       *bool
}

type MedicineRepository interface {
	Create(medicine *models.Medicine) error

	GetByID(id uint) (*models.Medicine, error)

	Delete(id uint) error

	Update(medicine *models.Medicine) error

	GetAll() ([]models.Medicine, error)

	List(filter MedicineFilter) ([]models.Medicine, error)
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

func (r *gormMedecineRepository) List(filter MedicineFilter) ([]models.Medicine, error) {
	var medicines []models.Medicine

	query := r.db.Model(&models.Medicine{})

	if filter.CategoryID != nil {
		query = query.Where("category_id = ?", *filter.CategoryID)
	}

	if filter.SubcategoryID != nil {
		query = query.Where("subcategory_id = ?", *filter.SubcategoryID)
	}

	if filter.InStock != nil {
		query = query.Where("in_stock = ?", *filter.InStock)
	}

	if err := query.Find(&medicines).Error; err != nil {
		return nil, err
	}

	return medicines, nil
}
