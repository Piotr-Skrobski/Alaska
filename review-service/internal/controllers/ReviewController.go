package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Piotr-Skrobski/Alaska/review-service/internal/dtos"
	"github.com/Piotr-Skrobski/Alaska/review-service/internal/models"
	"github.com/Piotr-Skrobski/Alaska/review-service/internal/services"
	"github.com/go-chi/chi/v5"
)

type ReviewController struct {
	ReviewService *services.ReviewService
}

func NewReviewController(reviewService *services.ReviewService) *ReviewController {
	return &ReviewController{ReviewService: reviewService}
}

func (c *ReviewController) RegisterRoutes(r chi.Router) {
	r.Post("/reviews", c.CreateReview)
	r.Get("/reviews/movie/{movieID}", c.GetReviewsByMovieID)
}

func (c *ReviewController) CreateReview(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreateReviewRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	review := &models.Review{
		UserID:  req.UserID,
		MovieID: req.MovieID,
		Rating:  req.Rating,
		Comment: req.Comment,
	}

	if err := c.ReviewService.CreateReview(review); err != nil {
		http.Error(w, "failed to create review", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(review)
}

func (c *ReviewController) GetReviewsByMovieID(w http.ResponseWriter, r *http.Request) {
	movieID := chi.URLParam(r, "movieID")

	reviews, err := c.ReviewService.GetReviewsByMovieID(movieID)
	if err != nil {
		http.Error(w, "failed to fetch reviews", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reviews)
}
