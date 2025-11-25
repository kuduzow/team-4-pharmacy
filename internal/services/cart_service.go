package services

import (
	"errors"

	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"github.com/kuduzow/team-4-pharmacy/internal/repository"
	"gorm.io/gorm"
)

var (
	ErrCartNotFound     = errors.New("корзина не найдена")
	ErrCartItemNotFound = errors.New("позиция не найдена")
	ErrInvalidQuantity  = errors.New("количество должно быть положительным")
	ErrOutOfStock       = errors.New("недостаточно товара на складе")
	ErrMedicineMissing  = errors.New("лекарство не найдено")
)

type CartService interface {
	Create(userID uint) (*models.Cart, error)
	GetCart(userID uint) (*models.UpdateCart, error)
	ClearCart(userID uint) error
}

type cartService struct {
	cartRepo repository.CartRepository
}

func NewCartService(cartRepo repository.CartRepository) CartService {
	return &cartService{cartRepo: cartRepo}
}
func (s *cartService) Create(id uint) (*models.Cart, error) {
	_, err := s.cartRepo.GetByUserID(id)

	if err == nil {
		return nil, errors.New("Корзина у пользователя уже существует")
	}
	cart := &models.Cart{
		UserID:     id,
		TotalPrice: 0,
	}
	if err := s.cartRepo.Create(cart); err != nil {
		return nil, err
	}
	return cart, nil
}

func (s *cartService) GetCart(userID uint) (*models.UpdateCart, error) {

	cart, err := s.cartRepo.GetByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCartNotFound
		}
		return nil, err
	}

	var total int64
	for i := range cart.Items {
		total += cart.Items[i].LineTotal
	}

	updateCart := &models.UpdateCart{
		UserID: cart.UserID,
	}

	return updateCart, nil
}

func (s *cartService) ClearCart(userID uint) error {
	return s.cartRepo.ClearByUserID(userID)
}
