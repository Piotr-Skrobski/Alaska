package router

import (
	"net/http"

	"github.com/Piotr-Skrobski/Alaska/user-service/internal/controllers"
	"github.com/go-chi/chi/v5"
)

func NewRouter(userController *controllers.UserController, middleware ...func(http.Handler) http.Handler) chi.Router {
	r := chi.NewRouter()

	for _, mw := range middleware {
		r.Use(mw)
	}

	userController.RegisterRoutes(r)

	return r
}
