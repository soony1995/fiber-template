package container

import (
	"login_module/internal/application/service"
	"login_module/internal/infrastructure/config"
	"login_module/internal/infrastructure/db"
	"login_module/internal/presentation/http/handler"

	"github.com/go-redis/redis/v8"
)

type Container struct {
	RedisClient *redis.Client

	AuthService *service.AuthService
	UserService *service.UserService

	AuthHandler  *handler.AuthHandler
	OAuthHandler *handler.OAuthHandler
	UserHandler  *handler.UserHandler
}

func BuildContainer() *Container {
	userRedisRepo := db.NewRedisAuthRepository(config.NewRedisClient())

	oAuthService := service.NewOAuthService(userRedisRepo)
	authService := service.NewAuthService(userRedisRepo)
	userService := service.NewUserService(userRedisRepo)

	authHandler := handler.NewAuthHandler(authService)
	oAuthHandler := handler.NewOAuthHandler(oAuthService)
	userHandler := handler.NewUserHandler(userService)

	return &Container{
		AuthHandler:  authHandler,
		OAuthHandler: oAuthHandler,
		UserHandler:  userHandler,
	}
}
