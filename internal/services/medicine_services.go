package services

import (
	"errors"
	"strings"

	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"github.com/kuduzow/team-4-pharmacy/internal/repository"
	"gorm.io/gorm"
)

var ErrMedicineNotFound = errors.New("Лекарство не найдено")

type MedicineService interface {
	CreateMedicine(req models.MedicineCreateRequest) (*models.Medicine, error)

	GetMedicineByID(id uint) (*models.Medicine, error)

	UpdateMedicine(id uint, req models.MedicineUpdateRequest) (*models.Medicine, error)

	DeleteMedicine(id uint) error

	GetAllMedicines() ([]models.Medicine, error)

	GetInStockMedicines() ([]models.Medicine, error)
}

type medicineService struct {
	medicines repository.MedicineRepository
}

func NewMedicineService(
	medicines repository.MedicineRepository,
) MedicineService {
	return &medicineService{
		medicines: medicines,
	}
}

func (s *medicineService) CreateMedicine(req models.MedicineCreateRequest) (*models.Medicine, error) {
	if err := s.ValidateCreateMedicine(req); err != nil {
		return nil, err
	}

	inStock := req.StockQuantity > 0
	medicine := &models.Medicine{
		Name:                 req.Name,
		Description:          req.Description,
		Price:                req.Price,
		InStock:              inStock,
		StockQuantity:        req.StockQuantity,
		CategoryID:           req.CategoryID,
		SubcategoryID:        req.SubcategoryID,
		Manufacturer:         req.Manufacturer,
		PrescriptionRequired: req.PrescriptionRequired,
		AvgRating:            req.AvgRating,
	}

	if err := s.medicines.Create(medicine); err != nil {
		return nil, err
	}

	return medicine, nil
}

func (s *medicineService) GetMedicineByID(id uint) (*models.Medicine, error) {
	medicine, err := s.medicines.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrMedicineNotFound
		}
		return nil, err
	}

	return medicine, nil
}

func (s *medicineService) UpdateMedicine(id uint, req models.MedicineUpdateRequest) (*models.Medicine, error) {
	medicine, err := s.medicines.GetByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrMedicineNotFound
		}
		return nil, err
	}

	if err := s.ApplyMedicineUpdate(medicine, req); err != nil {
		return nil, err
	}

	if err := s.medicines.Update(medicine); err != nil {
		return nil, err
	}

	return medicine, nil
}

func (s *medicineService) DeleteMedicine(id uint) error {
	_, err := s.medicines.GetByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrMedicineNotFound
		}
		return err
	}

	return s.medicines.Delete(id)
}

func (s *medicineService) GetAllMedicines() ([]models.Medicine, error) {
	return s.medicines.GetAll()
}

func (s *medicineService) GetInStockMedicines() ([]models.Medicine, error) {
	return s.medicines.GetInStock()
}

func (s *medicineService) ApplyMedicineUpdate(medicine *models.Medicine, req models.MedicineUpdateRequest) error {
	if req.Name != nil {
		trimmed := strings.TrimSpace(*req.Name)

		if trimmed == "" {
			return errors.New("Название лекарства обязательна")
		}
		medicine.Name = trimmed
	}

	if req.Price != nil {
		if *req.Price <= 0 {
			return errors.New("цена лекарства должна быть больше 0")
		}
		medicine.Price = *req.Price
	}

	inStock := *req.StockQuantity > 0

	if req.InStock != nil {
		medicine.InStock = inStock
	}

	return nil
}

func (s *medicineService) ValidateCreateMedicine(req models.MedicineCreateRequest) error {
	if req.CategoryID == 0 {
		return errors.New("поле category_id должно быть больше 0")
	}

	if req.SubcategoryID == 0 {
		return errors.New("поле subcategory_id должно быть больше 0")
	}

	if req.Name == "" {
		return errors.New("Название лекарства обязательна")
	}

	if req.Price <= 0 {
		return errors.New("цена лекарства должна быть больше 0")
	}

	if req.StockQuantity < 0 {
		return errors.New("количество лекарств на складе не должно быть отрицательным")
	}

	return nil
}
