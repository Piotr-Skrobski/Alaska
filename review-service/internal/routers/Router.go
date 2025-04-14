package router

import (
	"net/http"

	"github.com/Piotr-Skrobski/Alaska/review-service/internal/controllers"
	"github.com/go-chi/chi/v5"
)

func NewRouter(reviewController *controllers.ReviewController, middleware ...func(http.Handler) http.Handler) chi.Router {
	r := chi.NewRouter()

	for _, mw := range middleware {
		r.Use(mw)
	}

	reviewController.RegisterRoutes(r)

	return r
}
