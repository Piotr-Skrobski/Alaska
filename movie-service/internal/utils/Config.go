package utils

import (
	"os"
)

type Config struct {
	MongoURI   string
	RabbitURI  string
	OmdbAPIKey string
	Port       string
}

func LoadConfig() Config {
	return Config{
		MongoURI:   getEnv("MONGO_URI", "MONGO_URI_NOT_SET"),
		RabbitURI:  getEnv("RABBITMQ_URI", "RABBIT_URI_NOT_SET"),
		OmdbAPIKey: getEnv("OMDB_API_KEY", "OMDB_API_NOT_SET"),
		Port:       getEnv("PORT", "PORT_NOT_SET"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
