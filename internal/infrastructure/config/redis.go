package config

import (
	"github.com/go-redis/redis/v8"
)

// NewRedisClient initializes and returns a new Redis client
func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return client
}
