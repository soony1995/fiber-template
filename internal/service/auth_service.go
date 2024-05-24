package service

import (
	"context"
	"errors"
	jwt "login_module/internal/pkg/jwt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type AuthService struct {
	redisClient *redis.Client
}

func NewAuthService(redisClient *redis.Client) *AuthService {
	return &AuthService{redisClient: redisClient}
}

func (s *AuthService) Login(username, password string) (accessToken, refreshToken string, err error) {
	// 여기에 사용자 인증 로직 구현
	if username != "expected" || password != "password" {
		return "", "", errors.New("invalid credentials")
	}

	// 토큰 생성
	accessToken, err = jwt.GenerateToken(username)
	if err != nil {
		return "", "", err
	}
	refreshToken, err = jwt.GenerateRefreshToken(username)
	if err != nil {
		return "", "", err
	}

	// 리프레시 토큰을 Redis에 저장
	err = s.redisClient.Set(ctx, refreshToken, username, 24*time.Hour).Err()
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) RefreshToken(refreshToken string) (newAccessToken string, err error) {
	username, err := s.redisClient.Get(ctx, refreshToken).Result()
	if err != nil {
		return "", errors.New("invalid or expired refresh token")
	}

	// 새 액세스 토큰 생성
	newAccessToken, err = jwt.GenerateToken(username)
	if err != nil {
		return "", err
	}

	return newAccessToken, nil
}

func (s *AuthService) Logout(refreshToken string) error {
	return s.redisClient.Del(ctx, refreshToken).Err()
}
