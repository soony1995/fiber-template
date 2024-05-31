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

/*
	type User struct {
		RawData           map[string]interface{}
		Provider          string
		Email             string
		Name              string
		FirstName         string
		LastName          string
		NickName          string
		Description       string
		UserID            string
		AvatarURL         string
		Location          string
		AccessToken       string
		AccessTokenSecret string
		RefreshToken      string
		ExpiresAt         time.Time
		IDToken           string
	}
*/

func (s *OAuthService) StoreTokens(u goth.User) error {
	/*** store user info to Mysql data
	Email pk
	Pwd
	NickName
	Provider (google, naver, local)
	RegisteredAt
	LastLoginAt
	***/

	return nil
}
