package main

import (
	"database/sql"
	"log"

	"github.com/Piotr-Skrobski/Alaska/user-service/internal/utils"
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

}
