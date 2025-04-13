package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Piotr-Skrobski/Alaska/movie-service/internal/services"
	"github.com/go-chi/chi/v5"
)

type MovieController struct {
	movieService *services.MovieService
}

func NewMovieController(movieService *services.MovieService) *MovieController {
	return &MovieController{
		movieService: movieService,
	}
}

func (c *MovieController) RegisterRoutes(r chi.Router) {
	r.Get("/movies/title/{title}", c.GetMovieByTitle)
	r.Get("/movies/imdb/{id}", c.GetMovieByIMDbID)
}

func (c *MovieController) GetMovieByTitle(w http.ResponseWriter, r *http.Request) {
	title := chi.URLParam(r, "title")
	movie, err := c.movieService.GetMovieByTitle(title)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching movie: %v", err), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func (c *MovieController) GetMovieByIMDbID(w http.ResponseWriter, r *http.Request) {
	imdbID := chi.URLParam(r, "id")
	movie, err := c.movieService.GetMovieByIMDbID(imdbID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching movie: %v", err), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}
