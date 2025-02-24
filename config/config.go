package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName    string
	AppVersion string
	AppPort    string
	AppEnv     string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	return &Config{
		AppName:    getEnv("APP_NAME", "Go Hello World"),
		AppVersion: getEnv("APP_VERSION", "1.0.0"),
		AppPort:    getEnv("APP_PORT", "8080"),
		AppEnv:     getEnv("APP_ENV", "development"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
