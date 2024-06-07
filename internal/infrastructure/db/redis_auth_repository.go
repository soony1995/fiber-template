package db

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"login_module/internal/domain/model"
	"login_module/internal/domain/repository"
	"time"
)

type redisAuthRepository struct {
	client *redis.Client
}

func NewRedisAuthRepository(client *redis.Client) repository.AuthRepository {
	return &redisAuthRepository{client: client}
}

func (r *redisAuthRepository) SaveRefreshToken(ctx context.Context, m model.SaveRefreshToken) error {
	err := r.client.Set(ctx, m.UserUUID, m.RefreshToken, time.Duration(m.RefreshTokenExpiresIn)*time.Second).Err()
	if err != nil {
		return errors.Wrap(err, "failed to save refresh token in redis")
	}
	return nil
}

func (r *redisAuthRepository) GetUserByUserUUID(ctx context.Context, email string) (model.User, error) {
	userData, err := r.client.Get(ctx, email).Result()
	if errors.Is(err, redis.Nil) {
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

func (r *redisAuthRepository) DeleteIDToken(ctx context.Context, userUUID string) error {
	_, err := r.client.Del(ctx, userUUID).Result()
	return err
}
