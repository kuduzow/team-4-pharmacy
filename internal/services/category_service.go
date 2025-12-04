package services

import (
	"errors"
	"log/slog"
	"strings"

	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"github.com/kuduzow/team-4-pharmacy/internal/repository"
)

var ErrCategoryNotFound = errors.New("категория не найдена")
var ErrSubcategoryNotFound = errors.New("подкатегория не найдена")

type CategoryService interface {
	CreateCategory(req models.CreateCategory) (*models.Category, error)

	GetAll() ([]models.Category, error)

	CreateSubcategory(req models.CreateSubcategory) (*models.Subcategory, error)

	GetSubcategoriesByCategoryID(CategoryID uint) ([]models.Subcategory, error)
}

type categoryService struct {
	categories repository.CategoryRepository
	logger     *slog.Logger
}

func NewCategoryService(categories repository.CategoryRepository) CategoryService {
	logger := slog.Default()
	return &categoryService{
		categories: categories,
		logger:     logger,
	}
}

func (s *categoryService) CreateCategory(req models.CreateCategory) (*models.Category, error) {
	s.logger.Info("category_service.CreateCategory: creating category", slog.String("name", req.Name))

	if err := s.validateCategoryCreate(req); err != nil {
		s.logger.Error("category_service.CreateCategory: validation failed", slog.String("error", err.Error()))
		return nil, err
	}

	category := &models.Category{
		Name: strings.TrimSpace(req.Name),
	}

	if err := s.categories.Create(category); err != nil {
		s.logger.Error("category_service.CreateCategory: failed to create category", slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("category_service.CreateCategory: category created successfully", slog.Uint64("id", uint64(category.ID)))
	return category, nil
}

func (s *categoryService) GetAll() ([]models.Category, error) {
	s.logger.Info("category_service.GetAll: fetching all categories")
	categories, err := s.categories.GetAll()
	if err != nil {
		s.logger.Error("category_service.GetAll: failed to fetch categories", slog.String("error", err.Error()))
		return nil, err
	}
	s.logger.Info("category_service.GetAll: categories fetched successfully", slog.Int("count", len(categories)))
	return categories, nil
}

func (s *categoryService) CreateSubcategory(req models.CreateSubcategory) (*models.Subcategory, error) {
	s.logger.Info("category_service.CreateSubcategory: creating subcategory", slog.String("name", req.Name), slog.Uint64("category_id", uint64(req.CategoryID)))

	if err := s.validateSubcategoryCreate(req); err != nil {
		s.logger.Error("category_service.CreateSubcategory: validation failed", slog.String("error", err.Error()))
		return nil, err
	}

	categories, err := s.categories.GetAll()
	if err != nil {
		s.logger.Error("category_service.CreateSubcategory: failed to fetch categories", slog.String("error", err.Error()))
		return nil, err
	}

	found := false
	for _, c := range categories {
		if c.ID == req.CategoryID {
			found = true
			break
		}
	}

	if !found {
		s.logger.Error("category_service.CreateSubcategory: category not found", slog.Uint64("category_id", uint64(req.CategoryID)))
		return nil, ErrCategoryNotFound
	}

	sub := &models.Subcategory{
		Name:       strings.TrimSpace(req.Name),
		CategoryID: req.CategoryID,
	}

	if err := s.categories.CreateSubcategory(sub); err != nil {
		s.logger.Error("category_service.CreateSubcategory: failed to create subcategory", slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("category_service.CreateSubcategory: subcategory created successfully", slog.Uint64("id", uint64(sub.ID)))
	return sub, nil
}

func (s *categoryService) GetSubcategoriesByCategoryID(CategoryID uint) ([]models.Subcategory, error) {
	s.logger.Info("category_service.GetSubcategoriesByCategoryID: fetching subcategories", slog.Uint64("category_id", uint64(CategoryID)))

	subs, err := s.categories.GetSubcategoriesByCategoryID(CategoryID)
	if err != nil {
		s.logger.Error("category_service.GetSubcategoriesByCategoryID: failed to fetch subcategories", slog.String("error", err.Error()))
		return nil, err
	}

	s.logger.Info("category_service.GetSubcategoriesByCategoryID: subcategories fetched successfully", slog.Uint64("category_id", uint64(CategoryID)), slog.Int("count", len(subs)))
	return subs, nil
}

func (s *categoryService) validateCategoryCreate(req models.CreateCategory) error {
	if strings.TrimSpace(req.Name) == "" {
		s.logger.Warn("category_service.validateCategoryCreate: empty name provided")
		return errors.New("поле name не должно быть пустым")
	}
	return nil
}

func (s *categoryService) validateSubcategoryCreate(req models.CreateSubcategory) error {
	if strings.TrimSpace(req.Name) == "" {
		return errors.New("поле name не должно быть пустым")
	}

	if req.CategoryID == 0 {
		return errors.New("category_id не может быть пустым")
	}

	return nil
}
