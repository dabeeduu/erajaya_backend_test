package cache

import (
	"backend_golang/internal/entity"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type ProductCache interface {
	buildKey(version, sortBy, sortOrder string) string
	GetVersion(ctx context.Context) (string, error)
	BumpVersion(ctx context.Context) error
	GetAll(ctx context.Context, version, sortBy, sortOrder string) ([]entity.Product, error)
	SetAll(ctx context.Context, version, sortBy, sortOrder string, products []entity.Product) error
}

type productCache struct {
	redis *redis.Client
}

func NewProductCache(redis *redis.Client) *productCache {
	return &productCache{
		redis: redis,
	}
}

func (c *productCache) buildKey(version, sortBy, sortOrder string) string {
	key := fmt.Sprintf("products:all:v%s:sortBy=%s:sortOrder=%s", version, sortBy, sortOrder)
	return key
}

func (c *productCache) GetVersion(ctx context.Context) (string, error) {
	version, err := c.redis.Get(ctx, "products:cache:version").Result()
	if err == redis.Nil {
		return "0", nil
	}
	return version, err
}

func (c *productCache) BumpVersion(ctx context.Context) error {
	return c.redis.Incr(ctx, "products:cache:version").Err()
}

func (c *productCache) GetAll(ctx context.Context, version, sortBy, sortOrder string) ([]entity.Product, error) {
	key := c.buildKey(version, sortBy, sortOrder)

	data, err := c.redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var products []entity.Product
	if err := json.Unmarshal([]byte(data), &products); err != nil {
		return nil, nil
	}

	return products, nil
}

func (c *productCache) SetAll(ctx context.Context, version, sortBy, sortOrder string, products []entity.Product) error {
	key := c.buildKey(version, sortBy, sortOrder)

	bytes, err := json.Marshal(products)
	if err != nil {
		return err
	}

	return c.redis.Set(ctx, key, bytes, 10*time.Minute).Err()
}
