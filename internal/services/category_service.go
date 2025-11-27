package services

import (
	"errors"
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
}

func NewCategoryService(categories repository.CategoryRepository) CategoryService {
	return &categoryService{
		categories: categories,
	}
}

func (s *categoryService) CreateCategory(req models.CreateCategory) (*models.Category, error) {
	if err := s.validateCategoryCreate(req); err != nil {
		return nil, err
	}
	category := &models.Category{
		Name: strings.TrimSpace(req.Name),
	}
	if err := s.categories.Create(category); err != nil {
		return nil, err
	}
	return category, nil
}

func (s *categoryService) GetAll() ([]models.Category, error) {
	return s.categories.GetAll()
}

func (s *categoryService) CreateSubcategory(req models.CreateSubcategory) (*models.Subcategory, error) {
	if err := s.validateSubcategoryCreate(req); err != nil {
		return nil, err
	}

	categories, err := s.categories.GetAll()
	if err != nil {
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
		return nil, ErrCategoryNotFound
	}

	sub := &models.Subcategory{
		Name:       strings.TrimSpace(req.Name),
		CategoryID: req.CategoryID,
	}

	if err := s.categories.CreateSubcategory(sub); err != nil {
		return nil, err
	}

	return sub, nil
}

func (s *categoryService) GetSubcategoriesByCategoryID(CategoryID uint) ([]models.Subcategory, error) {
	subs, err := s.categories.GetSubcategoriesByCategoryID(CategoryID)
	if err != nil {
		return nil, err
	}

	return subs, nil
}

func (s *categoryService) validateCategoryCreate(req models.CreateCategory) error {
	if strings.TrimSpace(req.Name) == "" {
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
