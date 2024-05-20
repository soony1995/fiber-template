package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	Login(creds Credentials) (string, error)
}

type authService struct{}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var jwtSecret = []byte("your_secret_key")

func NewAuthService() AuthService {
	return &authService{}
}

func (s *authService) Login(creds Credentials) (string, error) {
	// 사용자 인증 로직
	if creds.Username != "admin" || creds.Password != "password" {
		return "", fmt.Errorf("invalid credentials")
	}

	// JWT 토큰 생성
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": creds.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
