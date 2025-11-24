package services

import (
	"errors"
	"strings"
	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"github.com/kuduzow/team-4-pharmacy/internal/repository"
	"gorm.io/gorm"
)

var ErrPromocodeNotFound = errors.New("промокод не найден")

type PromocodeService interface {
	CreatePromocode(req models.PromocodeCreateRequest) (*models.Promocode, error)

	GetAllPromocodes() ([]models.Promocode, error)

	UpdatePromocode(id uint, req models.PromocodeUpdateRequest) (*models.Promocode, error)

	DeletePromocode(id uint) error
}

type promocodeService struct {
	promocodes repository.PromocodeRepository
}

func NewPromocodeService(
	promocodes repository.PromocodeRepository,
) PromocodeService {
	return &promocodeService{promocodes: promocodes}
}

func (s *promocodeService) CreatePromocode(req models.PromocodeCreateRequest) (*models.Promocode, error) {
	if err := s.ValidateCreatePromocode(req); err != nil {
		return nil, err
	}

	promocode := &models.Promocode{
		Code:           req.Code,
		Description:    req.Description,
		DiscountType:   req.DiscountType,
		DiscountValue:  req.DiscountValue,
		ValidFrom:      req.ValidFrom,
		ValidTo:        req.ValidTo,
		MaxUses:        req.MaxUses,
		MaxUsesPerUser: req.MaxUsesPerUser,
	}

	if err := s.promocodes.Create(promocode); err != nil {
		return nil, err
	}

	return promocode, nil
}

func (s *promocodeService) GetAllPromocodes() ([]models.Promocode, error) {
	return s.promocodes.GetAll()
}

func (s *promocodeService) DeletePromocode(id uint) error {
	if _, err := s.promocodes.GetByID(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrPromocodeNotFound
		}
		return err
	}

	return s.promocodes.Delete(id)
}

func (s *promocodeService) UpdatePromocode(id uint, req models.PromocodeUpdateRequest) (*models.Promocode, error) {
	promocode, err := s.promocodes.GetByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPromocodeNotFound
		}
		return nil, err
	}

	if err := s.ApplyPromocodeUpdate(promocode, req); err != nil {
		return nil, err
	}

	if err := s.promocodes.Update(promocode); err != nil {
		return nil, err
	}

	return promocode, nil
}

func (s *promocodeService) ApplyPromocodeUpdate(promocode *models.Promocode, req models.PromocodeUpdateRequest) error {
	if req.Code != nil {
		trimmed := strings.TrimSpace(*req.Code)
		if trimmed == "" {
			return errors.New("код не может быть пустым")
		}
	}

	if req.DiscountType != nil {
		promocode.DiscountType = *req.DiscountType
	}

	if req.DiscountValue != nil {
		if *req.DiscountValue <= 0 {
			return errors.New("значение скидки не должно быть меньше 0")
		}

		if *req.DiscountType == models.DiscountTypePercent && *req.DiscountValue > 100 {
			return errors.New("процентная скидка не должна превышать 100")
		}

		promocode.DiscountValue = *req.DiscountValue
	}

	if req.ValidFrom != nil && req.ValidTo != nil {
		if req.ValidTo.Before(*req.ValidFrom) {
			return errors.New("дата окончания не должна быть раньше даты начала")
		}
	}

	return nil
}

func (s *promocodeService) ValidateCreatePromocode(req models.PromocodeCreateRequest) error {
	if req.Code == "" {
		return errors.New("код не может быть пустым")
	}

	if req.ValidTo.Before(req.ValidFrom) {
		return errors.New("дата окончания не должна быть раньше даты начала")
	}

	if !req.IsActive {
		return errors.New("промокод не активен")
	}

	return nil
}
