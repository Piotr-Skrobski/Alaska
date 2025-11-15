package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Piotr-Skrobski/Alaska/movie-service/internal/controllers"
	"github.com/Piotr-Skrobski/Alaska/movie-service/internal/repositories"
	router "github.com/Piotr-Skrobski/Alaska/movie-service/internal/routers"
	"github.com/Piotr-Skrobski/Alaska/movie-service/internal/services"
	"github.com/Piotr-Skrobski/Alaska/movie-service/internal/utils"
	"github.com/go-chi/chi/v5/middleware"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg := utils.LoadConfig()

	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v\n", err)
	}
	defer func() {
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			log.Printf("failed to disconnect MongoDB: %v\n", err)
		}
	}()

	if err := mongoClient.Ping(context.Background(), nil); err != nil {
		log.Fatalf("MongoDB ping failed: %v\n", err)
	}
	// Emojis should not be in code but since I *can*, I *will*.
	log.Println("✅ Connected to MongoDB")

	movieRepository := repositories.NewMovieRepository(mongoClient.Database("movies_db"))
	omdbService := services.NewOMDbService(cfg.OmdbAPIKey)
	movieService := services.NewMovieService(movieRepository, omdbService)

	movieController := controllers.NewMovieController(movieService)
	// TODO: Refactor later
	r := router.NewRouter(movieController, middleware.Logger)

	healthController := controllers.NewHealthController()
	healthController.RegisterRoutes(r)

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
