package container

import (
	"login_module/internal/handler"
	redisClient "login_module/internal/pkg/redis"
	"login_module/internal/service"

	"github.com/go-redis/redis/v8"
)

// Container는 모든 서비스와 핸들러 인스턴스를 포함합니다.
type Container struct {
	RedisClient *redis.Client

	AuthService *service.AuthService
	UserService *service.UserService

	AuthHandler *handler.AuthHandler
	UserHandler *handler.UserHandler
}

// BuildContainer는 모든 인스턴스를 생성하고 Container 구조체에 담아 반환합니다.
func BuildContainer() *Container {
	redisClient := redisClient.NewRedisClient()

	authService := service.NewAuthService(redisClient)
	userService := service.NewUserService(redisClient)

	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)

	return &Container{
		RedisClient: redisClient,

		AuthService: authService,
		UserService: userService,

		AuthHandler: authHandler,
		UserHandler: userHandler,
	}
}
