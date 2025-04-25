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
		PostgresURI: getEnv("POSTGRES_URI", "POSTGRES_URI_NOT_SET"),
		RabbitURI:   getEnv("RABBITMQ_URI", "RABBITMQ_URI_NOT_SET"),
		Port:        getEnv("PORT", "PORT_NOT_SET"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
