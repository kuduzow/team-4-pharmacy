package services

import (
	"errors"

	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"github.com/kuduzow/team-4-pharmacy/internal/repository"
	"gorm.io/gorm"
)

var ErrPaymentNotFound = errors.New("оплата не найдена")

type PaymentService interface {
	CreatePayment(req models.PaymentCreate) (*models.Payment, error)

	GetPaymentByID(id uint) (*models.Payment, error)

	UpdatePayment(id uint, req models.PaymentUpdate) (*models.Payment, error)

	DeletePayment(id uint) error
}

type paymentService struct {
	payments repository.PaymentRepository
}

func NewPaymentService(payment repository.PaymentRepository) PaymentService {
	return &paymentService{
		payments: payment,
	}
}
func (c *paymentService) CreatePayment(req models.PaymentCreate) (*models.Payment, error) {

	payment := &models.Payment{
		OrderID: req.OrderID,
		Amount:  req.Amount,
		Status:  req.Status,
		Method:  req.Method,
		PaidAt:  req.PaidAt,
	}
	if err := c.payments.Create(payment); err != nil {
		return nil, err
	}
	return payment, nil
}

func (c *paymentService) GetPaymentByID(id uint) (*models.Payment, error) {
	payment, err := c.payments.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPaymentNotFound
		}
		return nil, err
	}
	return payment, nil
}
func (c *paymentService) UpdatePayment(id uint, req models.PaymentUpdate) (*models.Payment, error) {
	payment, err := c.payments.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrCheckConstraintViolated) {
			return nil, ErrPaymentNotFound
		}
		return nil, err
	}

	if err := c.payments.Update(payment); err != nil {
		return nil, err
	}
	return payment, nil
}
func (c *paymentService) DeletePayment(id uint) error {
	if _, err := c.payments.GetByID(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrPaymentNotFound
		}
		return err
	}
	return c.payments.Delete(id)
}
