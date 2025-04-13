package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/Piotr-Skrobski/Alaska/movie-service/internal/controllers"
	"github.com/Piotr-Skrobski/Alaska/movie-service/internal/repositories"
	router "github.com/Piotr-Skrobski/Alaska/movie-service/internal/routers"
	"github.com/Piotr-Skrobski/Alaska/movie-service/internal/services"
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

	movieRepository := repositories.NewMovieRepository(mongoClient.Database("movies_db"))
	omdbService := services.NewOMDbService(cfg.OmdbAPIKey)
	movieService := services.NewMovieService(movieRepository, omdbService)

	movieController := controllers.NewMovieController(movieService)
	r := router.NewRouter(movieController, middleware.Logger)

	mqConn, err := amqp.Dial(cfg.RabbitURI)
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v\n", err)
	}
	defer mqConn.Close()
	log.Println("✅ Connected to RabbitMQ")

	addr := ":" + cfg.Port
	log.Printf("🚀 Starting server on %s...", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("Server failed: %v\n", err)
	}

}
