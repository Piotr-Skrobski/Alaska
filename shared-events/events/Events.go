package events

type MovieDeletedEvent struct {
	MovieID int    `json:"movie_id"`
	Title   string `json:"title"`
}

type ReviewCreatedEvent struct {
	ReviewID int    `json:"review_id"`
	MovieID  int    `json:"movie_id"`
	UserID   int    `json:"user_id"`
	Rating   int    `json:"rating"`
	Comment  string `json:"comment"`
	Created  string `json:"created"`
}

type UserRegisteredEvent struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
