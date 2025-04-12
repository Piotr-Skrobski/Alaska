package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Piotr-Skrobski/Alaska/movie-service/internal/models"
	"github.com/Piotr-Skrobski/Alaska/movie-service/internal/repositories"
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	MongoURI   string
	RabbitURI  string
	OmdbAPIKey string
	Port       string
}

func LoadConfig() Config {
	return Config{
		MongoURI:   getEnv("MONGO_URI", "TODO"),
		RabbitURI:  getEnv("RABBITMQ_URI", "TODO"),
		OmdbAPIKey: getEnv("OMDB_API_KEY", "TODO"),
		Port:       getEnv("PORT", "10002"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func main() {
	cfg := LoadConfig()

	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v\n", err)
	}

	if err := mongoClient.Ping(context.Background(), nil); err != nil {
		log.Fatalf("MongoDB ping failed: %v\n", err)
	}
	// Emojis should not be in code but since I *can*, I *will*.
	log.Println("✅ Connected to MongoDB")

	repository := repositories.NewMovieRepository(mongoClient.Database("movies_db"))
	repository.Create(generateSampleMovie())

	mqConn, err := amqp.Dial(cfg.RabbitURI)
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v\n", err)
	}
	defer mqConn.Close()
	log.Println("✅ Connected to RabbitMQ")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/movies/{title}", func(w http.ResponseWriter, r *http.Request) {
		title := chi.URLParam(r, "title")
		fmt.Fprintf(w, "Movie requested: %s\n", title)
	})

	addr := ":" + cfg.Port
	log.Printf("🚀 Starting server on %s...", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("Server failed: %v\n", err)
	}

}

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
