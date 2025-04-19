package events

type ReviewCreatedEvent struct {
	ReviewID int    `json:"review_id"`
	MovieID  int    `json:"movie_id"`
	UserID   int    `json:"user_id"`
	Rating   int    `json:"rating"`
	Comment  string `json:"comment"`
	Created  string `json:"created"`
}
