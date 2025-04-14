package repositories

import (
	"database/sql"

	"github.com/Piotr-Skrobski/Alaska/review-service/internal/models"
)

type ReviewRepository struct {
	DB *sql.DB
}

func NewReviewRepository(db *sql.DB) *ReviewRepository {
	return &ReviewRepository{DB: db}
}

func (r *ReviewRepository) CreateReview(review *models.Review) error {
	query := `INSERT INTO reviews (user_id, movie_id, rating, comment, created_at) VALUES ($1, $2, $3, $4, NOW()) RETURNING id`
	return r.DB.QueryRow(query, review.UserID, review.MovieID, review.Rating, review.Comment).Scan(&review.ID)
}

func (r *ReviewRepository) GetReviewsByMovieID(movieID int) ([]models.Review, error) {
	rows, err := r.DB.Query(`SELECT id, user_id, movie_id, rating, comment, created_at FROM reviews WHERE movie_id = $1 ORDER BY created_at DESC`, movieID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var review models.Review
		if err := rows.Scan(&review.ID, &review.UserID, &review.MovieID, &review.Rating, &review.Comment, &review.CreatedAt); err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}
