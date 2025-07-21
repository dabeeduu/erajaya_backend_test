package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort   string
	DatabaseURL  string
	ShutdownTime int
	RedisAddr    string
}

func LoadEnvConfig() (Config, error) {
	if err := godotenv.Load(); err != nil {
		return Config{}, err
	}

	st, err := strconv.Atoi(os.Getenv("SHUTDOWN_TIME_SEC"))
	if err != nil {
		return Config{}, err
	}

	config := Config{
		ServerPort:   os.Getenv("APP_PORT"),
		DatabaseURL:  os.Getenv("DATABASE_URL"),
		RedisAddr:    os.Getenv("REDIS_ADDR"),
		ShutdownTime: st,
	}

	return config, nil
}
