package jwt

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret_key")

type KakaoJwt struct {
	Keys []struct {
		Kid string `json:"kid"`
		Kty string `json:"kty"`
		Alg string `json:"alg"`
		Use string `json:"use"`
		N   string `json:"n"`
		E   string `json:"e"`
	} `json:"keys"`
}

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

func fetchPublicKeys() (KakaoJwt, error) {
	resp, err := http.Get(os.Getenv("KAKAO_ID_TOKEN_PUBLIC_KEY"))
	if err != nil {
		return KakaoJwt{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return KakaoJwt{}, err
	}
	var kakaoJwt KakaoJwt
	if err := json.Unmarshal(body, &kakaoJwt); err != nil {
		return kakaoJwt, err
	}
	return kakaoJwt, nil
}

func getPublicKey(kakaoJwt KakaoJwt, kid string) (interface{}, error) {
	for _, key := range kakaoJwt.Keys {
		if key.Kid == kid {
			nBytes, err := base64.RawURLEncoding.DecodeString(key.N)
			if err != nil {
				return nil, err
			}
			eBytes, err := base64.RawURLEncoding.DecodeString(key.E)
			if err != nil {
				return nil, err
			}
			n := new(big.Int).SetBytes(nBytes)
			e := int(new(big.Int).SetBytes(eBytes).Int64())

			return &rsa.PublicKey{N: n, E: e}, nil
		}
	}
	return nil, fmt.Errorf("public key not found")
}

func ValidateIDToken(idToken string) error {
	parts := strings.Split(idToken, ".")
	if len(parts) != 3 {
		return fmt.Errorf("invalid token format")
	}
	headerBytes, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return err
	}
	var header struct {
		Kid string `json:"kid"`
	}
	if err := json.Unmarshal(headerBytes, &header); err != nil {
		return err
	}
	KakaoJwt, err := fetchPublicKeys()
	if err != nil {
		return err
	}
	publicKey, err := getPublicKey(KakaoJwt, header.Kid)
	if err != nil {
		return err
	}
	token, err := jwt.Parse(idToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}
