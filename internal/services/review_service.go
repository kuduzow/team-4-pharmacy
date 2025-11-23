package services

import (
	"errors"

	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"github.com/kuduzow/team-4-pharmacy/internal/repository"
	"gorm.io/gorm"
)

var ReviewNotFound = errors.New("Not Found")

type ModelService interface {
	CreateReview(req models.ReviewForPost) (*models.Review, error)
	UpdateReview(id uint, req models.ReviewForUpdate) (*models.Review, error)
	DeleteReview(id uint) error
	ListByMedicineID(medicineID uint) ([]models.Review, error)
	GetAvgRating(medicineID uint) (float64, error)
}

type ReviewService struct {
	repo repository.ReviewRepository
}

func NewReviewService(repo repository.ReviewRepository) ModelService {
	return &ReviewService{repo: repo}
}

func (s *ReviewService) CreateReview(req models.ReviewForPost) (*models.Review, error) {
	review := models.Review{
		UserID:     req.UserID,
		MedicineID: req.MedicineID,
		Rating:     req.Rating,
		Text:       req.Text,
	}

	err := s.repo.Create(&review)
	if err != nil {
		return nil, err
	}

	return &review, nil

}

func (s *ReviewService) UpdateReview(id uint, req models.ReviewForUpdate) (*models.Review, error) {

	review, err := s.repo.GetByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ReviewNotFound
		}
		return nil, err
	}

	if req.Rating != nil {
		review.Rating = *req.Rating
	}

	if req.Text != nil {
		review.Text = *req.Text
	}

	err = s.repo.Update(review)
	if err != nil {
		return nil, err
	}
	return review, nil
}

func (s *ReviewService) DeleteReview(id uint) error {
	review, err := s.repo.GetByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ReviewNotFound
		}
		return err
	}
	return s.repo.Delete(review.ID)
}

func (s *ReviewService) ListByMedicineID(medicineID uint) ([]models.Review, error) {
	reviews, err := s.repo.ListByMedicineID(medicineID)
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

func (s *ReviewService) GetAvgRating(medicineID uint) (float64, error) {
	reviews, err := s.repo.ListByMedicineID(medicineID)
	if err != nil {
		return 0, err
	}

	if len(reviews) == 0 {
		return 0, nil
	}

	accumulator := 0
	for _, v := range reviews {
		accumulator += int(v.Rating)
	}

	avg := float64(accumulator) / float64(len(reviews))

	return avg, nil

}
