package services

import (
	"github.com/Piotr-Skrobski/Alaska/review-service/internal/models"
	"github.com/Piotr-Skrobski/Alaska/review-service/internal/repositories"
)

type ReviewService struct {
	ReviewRepo *repositories.ReviewRepository
}

func NewReviewService(reviewRepo *repositories.ReviewRepository) *ReviewService {
	return &ReviewService{ReviewRepo: reviewRepo}
}

func (s *ReviewService) CreateReview(review *models.Review) error {
	return s.ReviewRepo.CreateReview(review)
}

func (s *ReviewService) GetReviewsByMovieID(movieID string) ([]models.Review, error) {
	return s.ReviewRepo.GetReviewsByMovieID(movieID)
}

func (s *ReviewService) DeleteReviewsByUserID(userID int) error {
	return s.ReviewRepo.DeleteReviewsByUserID(userID)
}

func (s *ReviewService) DeleteReviewsByMovieID(movieID string) error {
	return s.ReviewRepo.DeleteReviewsByMovieID(movieID)
}
