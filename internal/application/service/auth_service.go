package service

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"login_module/internal/application/dto"
	"login_module/internal/domain/model"
	"login_module/internal/domain/repository"
)

type AuthService struct {
	userRepository repository.UserRepository
	redis          *redis.Client
}

func NewAuthService(userRepo repository.UserRepository, redis *redis.Client) *AuthService {
	return &AuthService{
		userRepository: userRepo,
		redis:          redis,
	}
}

func (s *AuthService) Login(ctx context.Context, req dto.LoginRequest) (*model.User, error) {
	user, err := s.userRepository.FindByUsername(ctx, req.Username)
	if err != nil {
		return nil, errors.New("user not found")
	}
	if !user.Authenticate(req.Password) {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}
