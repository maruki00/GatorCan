package config

import (
	"os"
	"strconv"
	"time"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type ServerConfig struct {
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

type CORSConfig struct {
	AllowedOrigins []string
}

type NotificationWebhook struct {
	URL    string
	RoleID string
}

type AppConfig struct {
	Environment         string
	Database            DatabaseConfig
	Server              ServerConfig
	CORS                CORSConfig
	NotificationWebhook NotificationWebhook
}

func LoadConfig() *AppConfig {
	return &AppConfig{
		Environment: getEnvOrDefault("APP_ENV", "development"),
		Database: DatabaseConfig{
			Host:     getEnvOrDefault("DB_HOST", "localhost"),
			Port:     getEnvAsIntOrDefault("DB_PORT", 5432),
			User:     getEnvOrDefault("DB_USER", "postgres"),
			Password: getEnvOrDefault("DB_PASSWORD", "postgres"),
			DBName:   getEnvOrDefault("DB_NAME", "gatorcan"),
			SSLMode:  getEnvOrDefault("DB_SSL_MODE", "disable"),
		},
		Server: ServerConfig{
			Port:         getEnvAsIntOrDefault("SERVER_PORT", 8080),
			ReadTimeout:  time.Duration(getEnvAsIntOrDefault("SERVER_READ_TIMEOUT", 10)) * time.Second,
			WriteTimeout: time.Duration(getEnvAsIntOrDefault("SERVER_WRITE_TIMEOUT", 10)) * time.Second,
			IdleTimeout:  time.Duration(getEnvAsIntOrDefault("SERVER_IDLE_TIMEOUT", 60)) * time.Second,
		},
		CORS: CORSConfig{
			AllowedOrigins: getEnvAsSliceOrDefault("CORS_ALLOWED_ORIGINS", []string{"*"}),
		},
		NotificationWebhook: NotificationWebhook{
			URL:    getEnvOrDefault("NOTIFICATION_WEBHOOK_URL", ""),
			RoleID: getEnvOrDefault("NOTIFICATION_ROLE_ID", ""),
		},
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsIntOrDefault(key string, defaultValue int) int {
	valueStr := getEnvOrDefault(key, "")
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}

func getEnvAsSliceOrDefault(key string, defaultValue []string) []string {
	valueStr := getEnvOrDefault(key, "")
	if valueStr == "" {
		return defaultValue
	}
	// You could split by comma or implement more complex parsing here
	// This is a simple version
	return []string{valueStr}
}
