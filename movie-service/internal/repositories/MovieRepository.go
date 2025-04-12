package repositories

import "github.com/Piotr-Skrobski/Alaska/movie-service/internal/models"

type MovieRepository interface {
	Create(movie models.Movie) error

	FindByTitle(title string) (models.Movie, error)

	FindByIMDbID(imdbID string) (models.Movie, error)

	Update(movie models.Movie) error

	Delete(imdbID string) error
}
