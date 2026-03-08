package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	GO_ENV      string
	APP_PORT    string
	APP_HOST    string
	APP_NAME    string
	APP_VERSION string
	APP_CORS    string

	DB_HOST     string
	DB_PORT     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string

	GOOGLE_CLIENT_ID        string
	GOOGLE_CLIENT_SECRET    string
	GOOGLE_REDIRECT_URL     string
	GOOGLE_SCOPES           string
	GOOGLE_AUTH_URL         string
	GOOGLE_TOKEN_URL        string
	GOOGLE_DEVICE_AUTH_URL  string
	GOOGLE_USER_INFO_URL    string
	GOOGLE_REVOKE_TOKEN_URL string
}

func NewEnv() *Env {
	if os.Getenv("CI") != "true" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file", err)
		}
	}

	return &Env{
		GO_ENV:      getEnv("GO_ENV", "development", false),
		APP_PORT:    getEnv("APP_PORT", "8000", false),
		APP_HOST:    getEnv("APP_HOST", "", true),
		APP_NAME:    getEnv("APP_NAME", "app-online-shop", false),
		APP_VERSION: getEnv("APP_VERSION", "1.0.0", false),
		APP_CORS:    getEnv("APP_CORS", "", false),

		DB_HOST:     getEnv("DB_HOST", "localhost", true),
		DB_PORT:     getEnv("DB_PORT", "1433", false),
		DB_USERNAME: getEnv("DB_USERNAME", "", true),
		DB_PASSWORD: getEnv("DB_PASSWORD", "", true),
		DB_NAME:     getEnv("DB_NAME", "db-online-shop", false),

		GOOGLE_CLIENT_ID:        getEnv("GOOGLE_CLIENT_ID", "", true),
		GOOGLE_CLIENT_SECRET:    getEnv("GOOGLE_CLIENT_SECRET", "", true),
		GOOGLE_REDIRECT_URL:     getEnv("GOOGLE_REDIRECT_URL", "", true),
		GOOGLE_SCOPES:           getEnv("GOOGLE_SCOPES", "", true),
		GOOGLE_AUTH_URL:         getEnv("GOOGLE_AUTH_URL", "", true),
		GOOGLE_TOKEN_URL:        getEnv("GOOGLE_TOKEN_URL", "", true),
		GOOGLE_DEVICE_AUTH_URL:  getEnv("GOOGLE_DEVICE_AUTH_URL", "", true),
		GOOGLE_USER_INFO_URL:    getEnv("GOOGLE_USER_INFO_URL", "", true),
		GOOGLE_REVOKE_TOKEN_URL: getEnv("GOOGLE_REVOKE_TOKEN_URL", "", true),
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
