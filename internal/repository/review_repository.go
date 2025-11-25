package repository

import (
	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"gorm.io/gorm"
)

type ReviewRepository interface {
	Create(review *models.Review) error
	Update(review *models.Review) error
	Delete(id uint) error
	ListByMedicineID(medicineID uint) ([]models.Review, error)
	GetByID(id uint) (*models.Review, error)
}

type gormReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &gormReviewRepository{db: db}
}

func (r *gormReviewRepository) Create(review *models.Review) error {
	if review == nil {
		return nil
	}
	return r.db.Create(review).Error
}

func (r *gormReviewRepository) Update(review *models.Review) error {
	if review == nil {
		return nil
	}
	return r.db.Save(review).Error
}

func (r *gormReviewRepository) Delete(id uint) error {

	return r.db.Delete(&models.Review{}, id).Error
}

func (r *gormReviewRepository) ListByMedicineID(medicineID uint) ([]models.Review, error) {
	var review []models.Review
	if err := r.db.Where("medicine_id = ?", medicineID).Find(&review).Error; err != nil {
		return nil, err
	}

	return review, nil
}

func (r *gormReviewRepository) GetByID(id uint) (*models.Review, error) {
	var review models.Review

	err := r.db.First(&review, id).Error
	if err != nil {
		return nil, err
	}
	return &review, err

}
