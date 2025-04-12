package services

import (
	"fmt"

	"github.com/Piotr-Skrobski/Alaska/movie-service/internal/errors"
	"github.com/Piotr-Skrobski/Alaska/movie-service/internal/models"
	"github.com/Piotr-Skrobski/Alaska/movie-service/internal/repositories"
)

type MovieService struct {
	repo        repositories.MovieRepository
	omdbService *OMDbService
}

func NewMovieService(repo repositories.MovieRepository, omdbService *OMDbService) *MovieService {
	return &MovieService{
		repo:        repo,
		omdbService: omdbService,
	}
}

func (s *MovieService) GetMovieByTitle(title string) (models.Movie, error) {
	if title == "" {
		return models.Movie{}, errors.ErrInvalidMovieTitle
	}

	movie, err := s.repo.FindByTitle(title)
	if err == nil && movie.Title != "" {
		return movie, nil
	}

	return s.fetchAndSaveFromOMDb(func() (models.Movie, error) {
		return s.omdbService.GetMovieByTitle(title)
	})
}

func (s *MovieService) GetMovieByIMDbID(imdbID string) (models.Movie, error) {
	if imdbID == "" {
		return models.Movie{}, errors.ErrInvalidMovieIMBdID
	}

	movie, err := s.repo.FindByIMDbID(imdbID)
	if err == nil && movie.IMDbID != "" {
		return movie, nil
	}

	return s.fetchAndSaveFromOMDb(func() (models.Movie, error) {
		return s.omdbService.GetMovieByIMDbID(imdbID)
	})
}

func (s *MovieService) fetchAndSaveFromOMDb(fetchFunc func() (models.Movie, error)) (models.Movie, error) {
	movie, err := fetchFunc()
	if err != nil {
		return models.Movie{}, fmt.Errorf("movie not found locally or in OMDb: %w", err)
	}

	if err := s.SaveMovie(movie); err != nil {
		fmt.Printf("Warning: Failed to save movie to database: %v\n", err)
	}

	return movie, nil
}

func (s *MovieService) SaveMovie(movie models.Movie) error {
	if movie.Title == "" {
		return errors.ErrInvalidMovieTitle
	}

	if movie.IMDbID == "" {
		return errors.ErrInvalidMovieIMBdID
	}

	existingMovie, err := s.repo.FindByIMDbID(movie.IMDbID)
	if err == nil && existingMovie.IMDbID != "" {
		return s.repo.Update(movie)
	}

	return s.repo.Create(movie)
}
