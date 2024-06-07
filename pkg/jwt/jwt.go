package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret_key")

type CustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &CustomClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}

func GenerateRefreshToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &CustomClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}

func VerifyJWTToken(tokenString string) (jwt.Claims, error) {
	// JWT 토큰 검증 로직 (간단한 예)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 키 가져오는 로직 필요
		return []byte("your-256-bit-secret"), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return token.Claims.(CustomClaims), nil
}

func ParseIDToken(idToken string) (string, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(idToken, jwt.MapClaims{})
	if err != nil {
		return "", fmt.Errorf("error parsing token: %s", err.Error())
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// "sub"는 일반적으로 사용자의 고유 ID를 나타냅니다.
		if userID, ok := claims["sub"].(string); ok {
			return userID, nil
		}
	}
	return "", fmt.Errorf("user ID not found in token")
}
