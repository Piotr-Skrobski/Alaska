package utils

import (
	"github.com/Piotr-Skrobski/Alaska/movie-service/internal/models"
)

func generateSampleMovie() models.Movie {
	return models.Movie{
		Title:      "The Shawshank Redemption",
		Year:       "1994",
		Genre:      "Drama",
		Director:   "Frank Darabont",
		Writer:     "Stephen King, Frank Darabont",
		Actors:     []string{"Tim Robbins", "Morgan Freeman", "Bob Gunton"},
		Plot:       "Two imprisoned men bond over a number of years, finding solace and eventual redemption through acts of common decency.",
		Language:   "English",
		Country:    "USA",
		PosterURL:  "https://m.media-amazon.com/images/M/MV5BMDFkYTc0MGEtZmNhMC00ZDIzLWFmNTEtODM1ZmRlYWMwMWFmXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_SX300.jpg",
		IMDbRating: 9.3,
		Runtime:    "142 min",
		BoxOffice: models.BoxOffice{
			Value:    28341469,
			Currency: "USD",
		},
		Awards:    "Nominated for 7 Oscars. Another 21 wins & 36 nominations.",
		Metascore: "80",
		IMDbID:    "tt0111161",
	}
}
