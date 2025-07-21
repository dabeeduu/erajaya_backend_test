package database

import (
	"backend_golang/config"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis(cfg config.Config) (*redis.Client, error) {
	addr := cfg.RedisAddr
	if addr == "" {
		addr = "localhost:6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("redis ping error: %w", err)
	}

	return client, nil
}
