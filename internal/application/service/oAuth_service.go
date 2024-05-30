package service

import (
	"github.com/go-redis/redis/v8"
	"github.com/markbates/goth"
)

type OAuthService struct {
	redisClient *redis.Client
}

func NewOAuthService(r *redis.Client) *OAuthService {
	return &OAuthService{
		redisClient: r,
	}
}

func (s *OAuthService) StoreTokens(u goth.User) error {
	
	return nil
}
