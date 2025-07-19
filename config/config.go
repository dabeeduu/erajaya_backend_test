package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort  string
	DatabaseURL string
}

func LoadEnvConfig() (Config, error) {
	if err := godotenv.Load(); err != nil {
		return Config{}, err
	}

	config := Config{
		ServerPort:  os.Getenv("APP_PORT"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}

	return config, nil
}
