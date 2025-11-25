package services

import (
	"errors"

	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"github.com/kuduzow/team-4-pharmacy/internal/repository"
	"gorm.io/gorm"
)

var errOrderNotFound = errors.New("заказ не найден")

type OrderService interface {
	CreateOrder(req models.OrderCreate) (*models.Order, error)
	GetOrderByID(id uint) (*models.Order, error)
	UpdateOrder(id uint, req models.OrderUpdate) (*models.Order, error)
	DeleteOrder(id uint) error
}

type orderService struct {
	order   repository.OrderRepository
	payment repository.PaymentRepository
}

func NewOrderService(order repository.OrderRepository, payment repository.PaymentRepository) OrderService {
	return &orderService{
		order:   order,
		payment: payment,
	}
}
func (c *orderService) CreateOrder(req models.OrderCreate) (*models.Order, error) {
	if err := c.validateOrderCreate(req); err != nil {
		return nil, err
	}
	order := &models.Order{
		UserID:          req.UserID,
		OrderStatus:     req.OrderStatus,
		TotalPrice:      req.TotalPrice,
		DiscountTotal:   req.DiscountTotal,
		FinalPrice:      req.FinalPrice,
		DeliveryAddress: req.DeliveryAddress,
		Comment:         req.Comment,
	}
	if err := c.order.Create(order); err != nil {
		return nil, err
	}
	return order, nil
}
func (c *orderService) GetOrderByID(id uint) (*models.Order, error) {
	order, err := c.order.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errOrderNotFound
		}
		return nil, err
	}
	return order, nil
}
func (c *orderService) UpdateOrder(id uint, req models.OrderUpdate) (*models.Order, error) {
	order, err := c.order.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errOrderNotFound
		}
		return nil, err
	}
if req.OrderStatus != nil {
		order.OrderStatus = *req.OrderStatus
	}
	if req.DeliveryAddress != nil {
		order.DeliveryAddress = *req.DeliveryAddress
	}
	if req.Comment != nil {
		order.Comment = *req.Comment
	}

	if err := c.order.Update(order); err != nil {
		return nil, err
	}

	return order, nil

}
func (c *orderService) DeleteOrder(id uint) error {
	if _, err := c.order.GetByID(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errOrderNotFound
		}
		return err
	}
	return c.order.Delete(id)
}
func(c *orderService)validateOrderCreate(req models.OrderCreate)error{
	if req.UserID == 0 {
		return errors.New("UserID")
	}
	if req.FinalPrice < 0 {
		return errors.New("FinalPrice не может быть отрицательным")
	}
	if req.TotalPrice < 0 {
		return errors.New("TotalPrice не может быть отрицательным")
	}
	if req.DeliveryAddress == "" {
		return errors.New("адрес доставки обязателен")
	}
	return nil
}