package repository

import (
	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *models.Category) error

	GetAll() ([]models.Category, error)

	CreateSubcategory(subcategory *models.Subcategory) error

	GetSubcategoriesByCategoryID(categoryID uint) ([]models.Subcategory, error)
}

type gormCategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &gormCategoryRepository{db: db}

}

func (r *gormCategoryRepository) Create(category *models.Category) error {
	if category == nil {
		return gorm.ErrInvalidData
	}

	return r.db.Create(category).Error
}

func (r *gormCategoryRepository) GetAll() ([]models.Category, error) {
	var categories []models.Category

	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *gormCategoryRepository) CreateSubcategory(subcategory *models.Subcategory) error {
	if subcategory == nil {
		return gorm.ErrInvalidData
	}

	return r.db.Create(subcategory).Error
}

func (r *gormCategoryRepository) GetSubcategoriesByCategoryID(categoryID uint) ([]models.Subcategory, error) {
	var subcategories []models.Subcategory

	if err := r.db.Where("category_id = ?", categoryID).Find(&subcategories).Error; err != nil {
		return nil, err
	}

	return subcategories, nil
}
