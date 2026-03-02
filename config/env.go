package config

import (
	"log"
	"log/slog"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Env struct {
	GO_ENV         string
	APP_LOG_LEVEL  slog.Level
	APP_LOG_FORMAT string // "text" = pretty readable, "json" = compact
	APP_PORT       string
	APP_NAME       string
	APP_VERSION    string
	APP_CORS       []string

	DB_HOST     string
	DB_PORT     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
}

func NewEnv() *Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	APP_CORS := getEnv("APP_CORS", "", false)
	cors := make([]string, 0)
	if APP_CORS != "" {
		cors = strings.Split(APP_CORS, ",")
	}

	APP_LOG_LEVEL := getEnv("APP_LOG_LEVEL", "DEBUG", false)
	logLevel := levelFromString(APP_LOG_LEVEL)

	return &Env{
		GO_ENV:        getEnv("GO_ENV", "development", false),
		APP_LOG_LEVEL: logLevel,
		APP_PORT:      getEnv("APP_PORT", "8000", false),
		APP_NAME:      getEnv("APP_NAME", "app-online-shop", false),
		APP_VERSION:   getEnv("APP_VERSION", "1.0.0", false),
		APP_CORS:      cors,

		DB_HOST:     getEnv("DB_HOST", "localhost", true),
		DB_PORT:     getEnv("DB_PORT", "1433", false),
		DB_USERNAME: getEnv("DB_USERNAME", "", true),
		DB_PASSWORD: getEnv("DB_PASSWORD", "", true),
		DB_NAME:     getEnv("DB_NAME", "db-online-shop", false),
	}
}

func getEnv(key string, defaultValue string, required bool) string {
	if value, exist := os.LookupEnv(key); exist && value != "" {
		return value
	}
	if required {
		log.Fatal("Error: Missing required environment variable", "key", key)
	}

	return defaultValue
}

func levelFromString(level string) slog.Level {
	switch level {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		log.Fatal("Invalid log level", "level", level)
		return slog.LevelDebug
	}
}
