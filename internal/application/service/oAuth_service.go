package service

import "github.com/go-redis/redis/v8"

type OAuthService struct {
	redisClient *redis.Client
}

func NewOAuthService(r *redis.Client) *OAuthService {
	return &OAuthService{
		redisClient: r,
	}
}

func (s *OAuthService) StoreTokens(accessToken string, refreshToken string) error {
	
	return nil
}
