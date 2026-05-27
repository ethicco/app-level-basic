package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env         string
	StoragePath string
}

func New(envFile string) (*Config, error) {
	if err := godotenv.Load(envFile); err != nil {
		return nil, fmt.Errorf("load env file: %w", err)
	}

	return &Config{
		Env:         getEnv("APP_ENV", "development"),
		StoragePath: getEnv("STORAGE_PATH", "."),
	}, nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
