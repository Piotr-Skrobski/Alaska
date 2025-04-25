package utils

import (
	"os"
)

type Config struct {
	PostgresURI         string
	RabbitURI           string
	RedisURI            string
	JWTSecret           string
	DiscordClientID     string
	DiscordClientSecret string
	Port                string
}

func LoadConfig() Config {
	return Config{
		PostgresURI:         getEnv("POSTGRES_URI", "POSTGRES_URI_NOT_SET"),
		RabbitURI:           getEnv("RABBITMQ_URI", "RABBITMQ_URI_NOT_SET"),
		RedisURI:            getEnv("REDIS_URI", "REDIS_URI_NOT_SET"),
		JWTSecret:           getEnv("JWT_SECRET", "JWT_SECRET_NOT_SET"),
		DiscordClientID:     getEnv("DISCORD_CLIENT_ID", "DISCORD_CLIENT_ID_NOT_SET"),
		DiscordClientSecret: getEnv("DISCORD_CLIENT_SECRET", "DISCORD_CLIENT_SECRET_NOT_SET"),
		Port:                getEnv("PORT", "PORT_NOT_SET"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
