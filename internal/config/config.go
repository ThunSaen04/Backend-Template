package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Config holds all configuration values loaded from .env
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
	AppPort    string

	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
}

// AppConfig is the global configuration instance
var AppConfig *Config

// LoadConfig reads configuration from .env file and populates AppConfig
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	AppConfig = &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "backend_template"),
		JWTSecret:  getEnv("JWT_SECRET", "default-secret"),
		AppPort:    getEnv("APP_PORT", "8080"),

		AccessTokenDuration:  parseDuration("ACCESS_TOKEN_DURATION", "6h"),
		RefreshTokenDuration: parseDuration("REFRESH_TOKEN_DURATION", "24h"),
	}

	log.Println("Configuration loaded successfully")
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// parseDuration reads a duration string from env (e.g. "6h", "24h", "30m") and falls back to a default
func parseDuration(key, defaultValue string) time.Duration {
	raw := getEnv(key, defaultValue)
	d, err := time.ParseDuration(raw)
	if err != nil {
		log.Printf("Warning: invalid duration for %s=%q, using default %s", key, raw, defaultValue)
		d, _ = time.ParseDuration(defaultValue)
	}
	return d
}
