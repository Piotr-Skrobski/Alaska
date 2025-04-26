package dtos

type CreateReviewRequest struct {
	UserID  int    `json:"user_id"`
	MovieID string `json:"movie_id"`
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}
