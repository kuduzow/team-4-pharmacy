package repository

import (
	"log/slog"

	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *models.Category) error

	GetAll() ([]models.Category, error)

	CreateSubcategory(subcategory *models.Subcategory) error

	GetSubcategoriesByCategoryID(categoryID uint) ([]models.Subcategory, error)

	GetCategoryByID(id uint) (*models.Category, error)

	GetSubcategoryByID(id uint) (*models.Subcategory, error)
}

type gormCategoryRepository struct {
	db     *gorm.DB
	logger *slog.Logger
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	logger := slog.Default()
	return &gormCategoryRepository{
		db:     db,
		logger: logger,
	}
}

func (r *gormCategoryRepository) Create(category *models.Category) error {
	if category == nil {
		r.logger.Error("category_repository.Create: category is nil")
		return gorm.ErrInvalidData
	}

	r.logger.Info("category_repository.Create: creating category", slog.String("name", category.Name))
	if err := r.db.Create(category).Error; err != nil {
		r.logger.Error("category_repository.Create: failed to create category", slog.String("error", err.Error()), slog.String("name", category.Name))
		return err
	}

	r.logger.Info("category_repository.Create: category created successfully", slog.Uint64("id", uint64(category.ID)), slog.String("name", category.Name))
	return nil
}

func (r *gormCategoryRepository) GetAll() ([]models.Category, error) {
	var categories []models.Category

	r.logger.Info("category_repository.GetAll: fetching all categories")
	if err := r.db.Find(&categories).Error; err != nil {
		r.logger.Error("category_repository.GetAll: failed to fetch categories", slog.String("error", err.Error()))
		return nil, err
	}

	r.logger.Info("category_repository.GetAll: categories fetched successfully", slog.Int("count", len(categories)))
	return categories, nil
}

func (r *gormCategoryRepository) CreateSubcategory(subcategory *models.Subcategory) error {
	if subcategory == nil {
		r.logger.Error("category_repository.CreateSubcategory: subcategory is nil")
		return gorm.ErrInvalidData
	}

	r.logger.Info("category_repository.CreateSubcategory: creating subcategory", slog.String("name", subcategory.Name), slog.Uint64("category_id", uint64(subcategory.CategoryID)))
	if err := r.db.Create(subcategory).Error; err != nil {
		r.logger.Error("category_repository.CreateSubcategory: failed to create subcategory", slog.String("error", err.Error()), slog.String("name", subcategory.Name))
		return err
	}

	r.logger.Info("category_repository.CreateSubcategory: subcategory created successfully", slog.Uint64("id", uint64(subcategory.ID)), slog.String("name", subcategory.Name))
	return nil
}

func (r *gormCategoryRepository) GetSubcategoriesByCategoryID(categoryID uint) ([]models.Subcategory, error) {
	var subcategories []models.Subcategory

	r.logger.Info("category_repository.GetSubcategoriesByCategoryID: fetching subcategories", slog.Uint64("category_id", uint64(categoryID)))
	if err := r.db.Where("category_id = ?", categoryID).Find(&subcategories).Error; err != nil {
		r.logger.Error("category_repository.GetSubcategoriesByCategoryID: failed to fetch subcategories", slog.String("error", err.Error()), slog.Uint64("category_id", uint64(categoryID)))
		return nil, err
	}

	r.logger.Info("category_repository.GetSubcategoriesByCategoryID: subcategories fetched successfully", slog.Uint64("category_id", uint64(categoryID)), slog.Int("count", len(subcategories)))
	return subcategories, nil
}

func (r *gormCategoryRepository) GetCategoryByID(id uint) (*models.Category, error) {
	var category models.Category

	r.logger.Info("category_repository.GetCategoryByID: fetching category", slog.Uint64("id", uint64(id)))
	if err := r.db.First(&category, id).Error; err != nil {
		r.logger.Error("category_repository.GetCategoryByID: failed to fetch category", slog.String("error", err.Error()), slog.Uint64("id", uint64(id)))
		return nil, err
	}

	r.logger.Info("category_repository.GetCategoryByID: category fetched successfully", slog.Uint64("id", uint64(id)), slog.String("name", category.Name))
	return &category, nil
}

func (r *gormCategoryRepository) GetSubcategoryByID(id uint) (*models.Subcategory, error) {
	var subcategory models.Subcategory

	r.logger.Info("category_repository.GetSubcategoryByID: fetching subcategory", slog.Uint64("id", uint64(id)))
	if err := r.db.First(&subcategory, id).Error; err != nil {
		r.logger.Error("category_repository.GetSubcategoryByID: failed to fetch subcategory", slog.String("error", err.Error()), slog.Uint64("id", uint64(id)))
		return nil, err
	}

	r.logger.Info("category_repository.GetSubcategoryByID: subcategory fetched successfully", slog.Uint64("id", uint64(id)), slog.String("name", subcategory.Name))
	return &subcategory, nil
}
