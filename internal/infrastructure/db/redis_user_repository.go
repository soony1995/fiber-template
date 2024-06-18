package db

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"login_module/internal/domain/model"
)

type RedisUserRepository struct {
	client *redis.Client
}

func NewRedisUserRepository(redisClient *redis.Client) *RedisUserRepository {
	return &RedisUserRepository{
		client: redisClient,
	}
}

func (repo *RedisUserRepository) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	val, err := repo.client.Get(ctx, username).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var user model.User
	if err := json.Unmarshal([]byte(val), &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *RedisUserRepository) Save(ctx context.Context, user *model.User) error {
	val, err := json.Marshal(user)
	if err != nil {
		return err
	}

	err = repo.client.Set(ctx, user.Username, val, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
