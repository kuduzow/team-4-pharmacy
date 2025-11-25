package repository

import (
	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error

	GetByID(id uint) (*models.User, error)

	Delete(id uint) error

	Update(user *models.User) error

	GetAll() ([]models.User, error)
}

type gormUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &gormUserRepository{db: db}
}

func (r *gormUserRepository) Create(user *models.User) error {
	if user == nil {
		return nil
	}
	return r.db.Create(&user).Error
}

func (r *gormUserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *gormUserRepository) Delete(id uint) error {
	var user models.User
	if err := r.db.Delete(&user, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *gormUserRepository) Update(user *models.User) error {
	if user == nil {
		return nil
	}

	return r.db.Save(&user).Error
}

func (r *gormUserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
