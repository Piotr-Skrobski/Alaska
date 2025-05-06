package main

import (
	"log"
	"net/http"

	"github.com/Piotr-Skrobski/Alaska/gateway-service/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:8081"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Gateway is healthy ✅"))
	})

	handlers.MountProxies(r)

	log.Println("🚀 Gateway running on :8080")
	http.ListenAndServe("0.0.0.0:8080", r)
}
