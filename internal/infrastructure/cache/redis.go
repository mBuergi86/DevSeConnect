package cache

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/redis/go-redis/v9"
)

func InitRedis() (*redis.Client, error) {
	redisURL := os.Getenv("REDIS_URL")
	redisAddr := strings.TrimPrefix(redisURL, "redis://")
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "yb0YjB1qc",
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	fmt.Printf("Connected to Redis at %s\n", redisAddr)
	return client, nil
}
