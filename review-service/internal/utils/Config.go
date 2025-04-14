package utils

import (
	"os"
)

type Config struct {
	PostgresURI string
	RabbitURI   string
	Port        string
}

func LoadConfig() Config {
	return Config{
		PostgresURI: getEnv("POSTGRES_URI", "postgres://user:pass@localhost:5432/dbname?sslmode=disable"),
		RabbitURI:   getEnv("RABBITMQ_URI", "amqp://admin:password@localhost:5672/"),
		Port:        getEnv("PORT", "10002"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
