package services

import (
	"errors"
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
	users repository.UserRepository
}

func NewUserService(users repository.UserRepository) UserService {
	return &userService{
		users: users,
	}
}

func (s *userService) CreateUser(req models.UserCreateRequest) (*models.User, error) {
	if err := s.validateUserCreate(req); err != nil {
		return nil, err
	}

	user := &models.User{
		FullName:       strings.TrimSpace(req.FullName),
		Email:          strings.TrimSpace(req.Email),
		Phone:          req.Phone,
		DefaultAddress: strings.TrimSpace(req.DefaultAddress),
	}

	if err := s.users.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	user, err := s.users.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	return user, nil
}

func (s *userService) UpdateUser(id uint, req models.UserUpdateRequest) (*models.User, error) {
	user, err := s.users.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	if err := s.applyUserUpdate(user, req); err != nil {
		return nil, err
	}

	if err := s.users.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) DeleteUser(id uint) error {
	if _, err := s.users.GetByID(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUserNotFound
		}

		return err
	}

	return s.users.Delete(id)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
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
