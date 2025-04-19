package events

type MovieDeletedEvent struct {
	MovieID int    `json:"movie_id"`
	Title   string `json:"title"`
}
