package database

import (
	"fmt"
	"os"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
	TimeZone string
}

func LoadConfig() (Config, error) {
	config := Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		SSLMode:  getenv("DB_SSLMODE", "disable"),
		TimeZone: getenv("DB_TIMEZONE", "Asia/Kolkata"),
	}

	switch {
	case config.Host == "":
		return Config{}, fmt.Errorf("DB_HOST is required")
	case config.Port == "":
		return Config{}, fmt.Errorf("DB_PORT is required")
	case config.User == "":
		return Config{}, fmt.Errorf("DB_USER is required")
	case config.Name == "":
		return Config{}, fmt.Errorf("DB_NAME is required")
	}

	return config, nil
}

func (config Config) DSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.Host,
		config.User,
		config.Password,
		config.Name,
		config.Port,
		config.SSLMode,
		config.TimeZone,
	)
}

func getenv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
