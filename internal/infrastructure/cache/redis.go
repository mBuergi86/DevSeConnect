package cache

import (
	"context"
	"os"
	"strings"

	"github.com/redis/go-redis/v9"
)

func InitRedis() (*redis.Client, error) {
	redisURL := os.Getenv("REDIS_URL")
	redisAddr := strings.TrimPrefix(redisURL, "redis://")
	client := redis.NewClient(&redis.Options{
		Addr: redisAddr,
		// Password: os.Getenv("REDIS_PASSWORD"),
		DB: 0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
