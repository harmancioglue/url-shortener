package config

import (
	"os"
	"strconv"
)

type Config struct {
	Server ServerConfig
	DB     DatabaseConfig
}

type ServerConfig struct {
	Host string
	Port int
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func Load() (*Config, error) {
	config := &Config{}

	// Server configuration
	config.Server.Host = getEnv("SERVER_HOST", "localhost")
	config.Server.Port = getEnvInt("SERVER_PORT")

	// Database configuration
	config.DB.Host = getEnv("DB_HOST", "localhost")
	config.DB.Port = getEnvInt("DB_PORT")
	config.DB.User = getEnv("DB_USER", "postgres")
	config.DB.Password = getEnv("DB_PASSWORD", "")
	config.DB.DBName = getEnv("DB_NAME", "url_shortener")
	config.DB.SSLMode = getEnv("DB_SSLMODE", "disable")

	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string) int {
	v, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		v = 0
	}

	return v
}

