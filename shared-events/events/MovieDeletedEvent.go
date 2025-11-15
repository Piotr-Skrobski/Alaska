package events

type MovieDeletedEvent struct {
	MovieID string `json:"movie_id"`
	Title   string `json:"title"`
}
