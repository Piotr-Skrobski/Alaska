package router

import (
	"net/http"

	"github.com/Piotr-Skrobski/Alaska/movie-service/internal/controllers"
	"github.com/go-chi/chi/v5"
)

func NewRouter(movieController *controllers.MovieController, middleware ...func(http.Handler) http.Handler) chi.Router {
	r := chi.NewRouter()

	for _, mw := range middleware {
		r.Use(mw)
	}

	movieController.RegisterRoutes(r)

	return r
}
