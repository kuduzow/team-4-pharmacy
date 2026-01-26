package services

import (
	"errors"

	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"github.com/kuduzow/team-4-pharmacy/internal/repository"
	"gorm.io/gorm"
)

type CartItemService interface {
	AddItem(userID uint, req models.CartCreateItemRequest) (*models.Cart, error)
}

type cartItemService struct {
	cartRepo repository.CartRepository
	itemRepo repository.CartItemRepository
	medicine repository.MedicineRepository
	db       *gorm.DB
}

func NewCartItemService(cartRepo repository.CartRepository, itemRepo repository.CartItemRepository,
	medicineRepo repository.MedicineRepository, db *gorm.DB) CartItemService {
	return &cartItemService{
		cartRepo: cartRepo,
		itemRepo: itemRepo,
		medicine: medicineRepo,
		db:       db,
	}
}

func (s *cartItemService) AddItem(userID uint, req models.CartCreateItemRequest) (*models.Cart, error) {
	if req.Quantity <= 0 {
		return nil, ErrInvalidQuantity
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		medRepo := repository.NewMedicineRepository(tx)
		cartRepo := repository.NewCartRepository(tx)
		itemRepo := repository.NewCartItemRepository(tx)

		med, err := medRepo.GetByID(req.MedicineID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrMedicineMissing
			}
			return err
		}

		if int(req.Quantity) > med.StockQuantity {
			return ErrOutOfStock
		}

		cart, err := cartRepo.GetByUserID(userID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c := &models.Cart{UserID: userID, TotalPrice: 0}
				if err := cartRepo.Create(c); err != nil {
					return err
				}
				cart = c
			} else {
				return err
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
				return ErrOutOfStock
			}
			existing.Quantity = newQty
			existing.PricePerUnit = pricePerUnit
			existing.LineTotal = existing.Quantity * existing.PricePerUnit

			if err := itemRepo.Update(existing); err != nil {
				return err
			}
		} else {
			item := &models.CartItem{
				CartID:       cart.ID,
				MedicineID:   req.MedicineID,
				Name:         med.Name,
				Quantity:     req.Quantity,
				PricePerUnit: pricePerUnit,
				LineTotal:    req.Quantity * pricePerUnit,
			}

			if err := itemRepo.Create(item); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	updatedCart, err := s.cartRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	var total int64
	for i := range updatedCart.Items {
		total += updatedCart.Items[i].LineTotal
	}
	updatedCart.TotalPrice = total

	return updatedCart, nil
}
