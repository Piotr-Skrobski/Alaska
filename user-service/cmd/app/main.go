package main

import (
	"context"
	"database/sql"
	"log"

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

}
