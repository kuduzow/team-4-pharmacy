package services

import (
	"errors"
	"log/slog"
	"strings"

	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"github.com/kuduzow/team-4-pharmacy/internal/repository"
	"gorm.io/gorm"
)

var ErrUserNotFound = errors.New("пользователь не найден")

type UserService interface {
	CreateUser(req models.UserCreateRequest) (*models.User, error)

	GetUserByID(id uint) (*models.User, error)

	UpdateUser(id uint, req models.UserUpdateRequest) (*models.User, error)

	DeleteUser(id uint) error

	GetAllUsers() ([]models.User, error)
}

type userService struct {
	users  repository.UserRepository
	logger *slog.Logger
}

func NewUserService(logger *slog.Logger, users repository.UserRepository) UserService {
	return &userService{
		users:  users,
		logger: logger,
	}
}

func (s *userService) CreateUser(req models.UserCreateRequest) (*models.User, error) {
	s.logger.Debug("Валидация данных для создания пользователя", "request", req)
	if err := s.validateUserCreate(req); err != nil {
		s.logger.Warn("Ошибка валидации при создании пользователя", "error", err)
		return nil, err
	}

	s.logger.Debug("Создание пользователя", "request", req)
	user := &models.User{
		FullName:       strings.TrimSpace(req.FullName),
		Email:          strings.TrimSpace(req.Email),
		Phone:          req.Phone,
		DefaultAddress: strings.TrimSpace(req.DefaultAddress),
	}

	if err := s.users.Create(user); err != nil {
		s.logger.Error("Не удалось создать пользователя", "error", err)
		return nil, err
	}
	s.logger.Info("Пользователь успешно создан", "user", user)
	return user, nil
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	s.logger.Debug("Получение пользователя по ID", "id", id)
	user, err := s.users.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.logger.Error("Пользователь не найден", "id", id)
			return nil, ErrUserNotFound
		}
		s.logger.Warn("Не удалось получить пользователя", "id", id, "error", err)
		return nil, err
	}
	s.logger.Info("Пользователь успешно получен", "user", user)
	return user, nil
}

func (s *userService) UpdateUser(id uint, req models.UserUpdateRequest) (*models.User, error) {
	s.logger.Debug("Обновление пользователя", "id", id, "request", req)
	user, err := s.users.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.logger.Error("Пользователь не найден", "id", id)
			return nil, ErrUserNotFound
		}
		s.logger.Error("Не удалось получить пользователя для обновления", "id", id, "error", err)
		return nil, err
	}

	if err := s.applyUserUpdate(user, req); err != nil {
		s.logger.Error("Ошибка при применении обновлений к пользователю", "id", id, "error", err)
		return nil, err
	}

	if err := s.users.Update(user); err != nil {
		s.logger.Warn("Не удалось обновить пользователя", "id", id, "error", err)
		return nil, err
	}
	s.logger.Info("Пользователь успешно обновлен", "user", user)
	return user, nil
}

func (s *userService) DeleteUser(id uint) error {
	s.logger.Debug("Удаление пользователя", "id", id)
	if _, err := s.users.GetByID(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.logger.Error("Пользователь не найден для удаления", "id", id)
			return ErrUserNotFound
		}
		s.logger.Error("Не удалось получить пользователя для удаления", "id", id, "error", err)
		return err
	}
	s.logger.Debug("Пользователь найден, приступаем к удалению", "id", id)
	return s.users.Delete(id)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	s.logger.Debug("Получение всех пользователей")
	return s.users.GetAll()
}

func (s *userService) validateUserCreate(req models.UserCreateRequest) error {
	if strings.TrimSpace(req.FullName) == "" {
		return errors.New("поле full_name не должно быть пустым")
	}

	if strings.TrimSpace(req.Email) == "" {
		return errors.New("поле email не должно быть пустым")
	}

	return nil
}

func (s *userService) applyUserUpdate(user *models.User, req models.UserUpdateRequest) error {
	if req.FullName != nil {
		trimmed := strings.TrimSpace(*req.FullName)
		if trimmed == "" {
			return errors.New("поле full_name не должно быть пустым")
		}
		user.FullName = trimmed
	}

	if req.Email != nil {
		email := strings.TrimSpace(*req.Email)
		if email == "" {
			return errors.New("поле email не должно быть пустым")
		}
		user.Email = email
	}

	if req.Phone != nil {
		user.Phone = *req.Phone
	}

	if req.DefaultAddress != nil {
		user.DefaultAddress = strings.TrimSpace(*req.DefaultAddress)
	}

	return nil
}
