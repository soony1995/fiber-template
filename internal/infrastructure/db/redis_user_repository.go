package db

import (
	"context"
	"encoding/json"

	"login_module/internal/domain/model"
	"login_module/internal/domain/repository"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

type redisUserRepository struct {
	client *redis.Client
}

func NewRedisUserRepository(client *redis.Client) repository.UserRepository {
	return &redisUserRepository{client: client}
}

func (r *redisUserRepository) SaveUser(ctx context.Context, user model.User) error {
	userData, err := json.Marshal(user)
	if err != nil {
		return errors.Wrap(err, "failed to marshal user data")
	}

	err = r.client.Set(ctx, user.Email, userData, 0).Err()
	if err != nil {
		return errors.Wrap(err, "failed to save user data to redis")
	}

	return nil
}

func (r *redisUserRepository) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	userData, err := r.client.Get(ctx, email).Result()
	if err == redis.Nil {
		return model.User{}, errors.New("user not found")
	} else if err != nil {
		return model.User{}, errors.Wrap(err, "failed to get user data from redis")
	}

	var user model.User
	err = json.Unmarshal([]byte(userData), &user)
	if err != nil {
		return model.User{}, errors.Wrap(err, "failed to unmarshal user data")
	}

	return user, nil
}
