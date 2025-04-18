package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/Piotr-Skrobski/Alaska/user-service/internal/controllers"
	"github.com/Piotr-Skrobski/Alaska/user-service/internal/repositories"
	router "github.com/Piotr-Skrobski/Alaska/user-service/internal/routers"
	"github.com/Piotr-Skrobski/Alaska/user-service/internal/services"
	"github.com/Piotr-Skrobski/Alaska/user-service/internal/utils"
	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	cfg := utils.LoadConfig()

	mqConn, err := amqp.Dial(cfg.RabbitURI)
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v\n", err)
	}
	defer mqConn.Close()
	log.Println("✅ Connected to RabbitMQ")
	channel, err := mqConn.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()
	utils.SetUpQueue(channel)

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
	if err := repositories.MigrateUserTable(postgresDb); err != nil {
		log.Fatalf("database migration failed: %v\n", err)
	}
	log.Println("✅ User table migrated (or already exists)")

	opt, err := redis.ParseURL(cfg.RedisURI)
	if err != nil {
		log.Fatalf("failed to parse Redis URI: %v\n", err)
	}
	redisClient := redis.NewClient(opt)
	_, err = redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("failed to connect to Redis at ping stage: %v\n", err)
	}
	defer redisClient.Close()
	log.Println("✅ Connected to Redis")

	userRepo := repositories.NewUserRepository(postgresDb)
	sessionService := services.NewSessionService(redisClient)
	eventPublisher := services.NewEventPublisher(channel)
	userService := services.NewUserService(userRepo, sessionService, eventPublisher)
	userController := controllers.NewUserController(userService)

	r := router.NewRouter(userController)

	// TODO: Refactor later
	healthController := controllers.NewHealthController()
	healthController.RegisterRoutes(r)

	addr := ":" + cfg.Port
	log.Printf("🚀 Starting server on %s...", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("Server failed: %v\n", err)
	}

}
