package repository

import (
	"errors"
	"log/slog"

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
	db     *gorm.DB
	logger *slog.Logger
}

func NewUserRepository(logger *slog.Logger, db *gorm.DB) UserRepository {
	return &gormUserRepository{db: db, logger: logger}
}

func (r *gormUserRepository) Create(user *models.User) error {
	r.logger.Debug("Создание пользователя", "user", user)
	if user == nil {
		r.logger.Warn("Не удалось создать пользователя: пользователь равен нулю")
		return errors.New("user is nil")
	}
	r.logger.Info("Пользователь успешно создан", "user", user)
	return r.db.Create(user).Error
}

func (r *gormUserRepository) GetByID(id uint) (*models.User, error) {
	r.logger.Debug("Получение пользователя по ID", "id", id)
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		r.logger.Error("Не удалось получить пользователя", "id", id, "error", err)
		return nil, err
	}
	r.logger.Info("Пользователь успешно получен", "user", user)
	return &user, nil
}

func (r *gormUserRepository) Delete(id uint) error {
	r.logger.Debug("Удаление пользователя по ID", "id", id)
	var user models.User
	if err := r.db.Delete(&user, id).Error; err != nil {
		r.logger.Error("Не удалось удалить пользователя", "id", id, "error", err)
		return err
	}
	r.logger.Debug("Пользователь успешно удален", "id", id)
	return nil
}

func (r *gormUserRepository) Update(user *models.User) error {
	r.logger.Debug("Обновление пользователя", "user", user)
	if user == nil {
		r.logger.Warn("Не удалось обновить пользователя: пользователь равен нулю")
		return errors.New("user is nil")
	}
	r.logger.Info("Пользователь успешно обновлен", "user", user)
	return r.db.Save(user).Error
}

func (r *gormUserRepository) GetAll() ([]models.User, error) {

	r.logger.Debug("Получение всех пользователей")

	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		r.logger.Error("Не удалось получить пользователей", "error", err)
		return nil, err
	}

	r.logger.Debug("Пользователи успешно получены", "count", len(users))
	return users, nil
}
