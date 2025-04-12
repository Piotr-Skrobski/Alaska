package models

type Movie struct {
	Title      string    `bson:"title" json:"title"`
	Year       string    `bson:"year" json:"year"`
	Genre      string    `bson:"genre" json:"genre"`
	Director   string    `bson:"director" json:"director"`
	Writer     string    `bson:"writer" json:"writer"`
	Actors     []string  `bson:"actors" json:"actors"`
	Plot       string    `bson:"plot" json:"plot"`
	Language   string    `bson:"language" json:"language"`
	Country    string    `bson:"country" json:"country"`
	PosterURL  string    `bson:"poster_url" json:"poster_url"`
	IMDbRating float64   `bson:"imdb_rating" json:"imdb_rating"`
	Runtime    string    `bson:"runtime" json:"runtime"`
	BoxOffice  BoxOffice `bson:"box_office" json:"box_office"`
	Awards     string    `bson:"awards" json:"awards"`
	Metascore  string    `bson:"metascore" json:"metascore"`
	IMDbID     string    `bson:"imdb_id" json:"imdb_id"`
}
