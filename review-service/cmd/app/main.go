package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Piotr-Skrobski/Alaska/review-service/internal/controllers"
	"github.com/Piotr-Skrobski/Alaska/review-service/internal/repositories"
	router "github.com/Piotr-Skrobski/Alaska/review-service/internal/routers"
	"github.com/Piotr-Skrobski/Alaska/review-service/internal/services"
	"github.com/Piotr-Skrobski/Alaska/review-service/internal/utils"
	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	cfg := utils.LoadConfig()

	postgresDb, err := sql.Open("postgres", cfg.PostgresURI)
	if err != nil {
		log.Fatalf("failed to connect to PostgreSQL: %v\n", err)
	}
	defer postgresDb.Close()
	err = postgresDb.Ping()
	if err != nil {
		log.Fatalf("failed to connect to PostgreSQL at ping stage: %v\n", err)
	}
	log.Println("✅ Connected to PostgreSQL")

	if err := repositories.MigrateReviewTable(postgresDb); err != nil {
		log.Fatalf("database migration failed: %v\n", err)
	}
	log.Println("✅ Review table migrated (or already exists)")

	reviewRepo := repositories.NewReviewRepository(postgresDb)
	reviewService := services.NewReviewService(reviewRepo)
	reviewController := controllers.NewReviewController(reviewService)

	// Connect to RabbitMQ for event consumption
	mqConn, err := amqp.Dial(cfg.RabbitURI)
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v\n", err)
	}
	defer mqConn.Close()
	log.Println("✅ Connected to RabbitMQ")

	channel, err := mqConn.Channel()
	if err != nil {
		log.Fatalf("failed to open RabbitMQ channel: %v\n", err)
	}
	defer channel.Close()

	// Start consuming events
	eventConsumer := services.NewEventConsumer(channel, reviewService)
	eventConsumer.ConsumeUserDeletedEvents()
	eventConsumer.ConsumeMovieDeletedEvents()
	log.Println("✅ Started consuming events")

	r := router.NewRouter(reviewController)

	// TODO: Refactor later
	healthController := controllers.NewHealthController()
	healthController.RegisterRoutes(r)

	addr := ":" + cfg.Port
	log.Printf("🚀 Starting review service on %s...", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("Server failed: %v\n", err)
	}
}
