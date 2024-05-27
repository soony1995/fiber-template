package container

import (
	"login_module/internal/application/service"
	"login_module/internal/infrastructure/db"
	"login_module/internal/presentation/http/handler"

	"github.com/go-redis/redis/v8"
)

type Container struct {
	RedisClient *redis.Client

	AuthService *service.AuthService
	UserService *service.UserService

	AuthHandler *handler.AuthHandler
	UserHandler *handler.UserHandler
}

func BuildContainer() *Container {
	redisClient := db.NewRedisClient()

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
