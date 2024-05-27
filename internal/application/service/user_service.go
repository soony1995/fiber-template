package service

import (
	"context"
	"encoding/json"
	"errors"
	"login_module/internal/application/dto"
	"login_module/internal/domain/model"

	"github.com/go-redis/redis/v8"
)

type UserService struct {
	redisClient *redis.Client
}

func NewUserService(redisClient *redis.Client) *UserService {
	return &UserService{
		redisClient: redisClient,
	}
}

func (s *UserService) GetUser(ctx context.Context, id string) (*dto.UserDTO, error) {
	result, err := s.redisClient.Get(ctx, id).Result()
	if err != nil {
		return nil, errors.New("user not found")
	}
	user := &model.User{}
	err = json.Unmarshal([]byte(result), user)
	if err != nil {
		return nil, err
	}
	userDTO := &dto.UserDTO{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
	return userDTO, nil
}

func (s *UserService) CreateUser(ctx context.Context, userDTO *dto.UserDTO) error {
	user := &model.User{
		ID:       userDTO.ID,
		Username: userDTO.Username,
		Email:    userDTO.Email,
	}
	userData, err := json.Marshal(user)
	if err != nil {
		return err
	}
	if err := s.redisClient.Set(ctx, user.ID, userData, 0).Err(); err != nil {
		return errors.New("failed to save user")
	}
	return nil
}
