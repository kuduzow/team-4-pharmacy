package services

import (
	"errors"

	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"github.com/kuduzow/team-4-pharmacy/internal/repository"
	"gorm.io/gorm"
)

type CartItemService interface {
	AddItem(userID uint, req models.CartCreateItemRequest) (*models.UpdateCart, error)
}

type cartItemService struct {
	cartRepo repository.CartRepository
	itemRepo repository.CartItemRepository
	medicine repository.MedicineRepository
}

func NewCartItemService(cartRepo repository.CartRepository, itemRepo repository.CartItemRepository,
	medicineRepo repository.MedicineRepository, db *gorm.DB) CartItemService {
	return &cartItemService{
		cartRepo: cartRepo,
		itemRepo: itemRepo,
		medicine: medicineRepo,
	}
}

func (s *cartItemService) AddItem(userID uint, req models.CartCreateItemRequest) (*models.UpdateCart, error) {
	if req.Quantity <= 0 {
		return nil, ErrInvalidQuantity
	}

	med, err := s.medicine.GetByID(req.MedicineID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrMedicineMissing
		}
		return nil, err
	}

	if int(req.Quantity) > med.StockQuantity {
		return nil, ErrOutOfStock
	}

	cart, err := s.cartRepo.GetByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c := &models.Cart{UserID: userID, TotalPrice: 0}
			if err := s.cartRepo.Create(c); err != nil {
				return nil, err
			}
			cart = c
		} else {
			return nil, err
		}
	}

	pricePerUnit := int64(med.Price * 100)
	var existing *models.CartItem
	for i := range cart.Items {
		if cart.Items[i].MedicineID == req.MedicineID {
			existing = &cart.Items[i]
			break
		}
	}

	if existing != nil {
		newQty := existing.Quantity + req.Quantity
		if int(newQty) > med.StockQuantity {
			return nil, ErrOutOfStock
		}
		existing.Quantity = newQty
		existing.PricePerUnit = pricePerUnit
		existing.LineTotal = existing.Quantity * existing.PricePerUnit

		if err := s.itemRepo.Update(existing); err != nil {
			return nil, err
		}
	} else {
		item := &models.CartItem{
			MedicineID:   req.MedicineID,
			Name:         med.Name,
			Quantity:     req.Quantity,
			PricePerUnit: pricePerUnit,
			LineTotal:    req.Quantity * pricePerUnit,
		}

		if err := s.itemRepo.Create(item); err != nil {
			return nil, err
		}
	}

	updatedCart, err := s.cartRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	var total int64
	for i := range updatedCart.Items {
		total += updatedCart.Items[i].LineTotal
	}

	return &models.UpdateCart{
		UserID: updatedCart.UserID,
	}, nil
}
