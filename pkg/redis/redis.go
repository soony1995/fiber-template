package redis

import (
	"github.com/go-redis/redis/v8"
)

// NewRedisClient는 Redis 클라이언트를 생성하는 팩토리 함수입니다.
func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // Redis 비밀번호가 없는 경우 빈 문자열로 설정
		DB:       0,  // 기본 DB 사용
	})
}
