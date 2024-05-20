package container

import (
	"login_module/internal/handler"
	"login_module/internal/service"
	redisClient "login_module/pkg/redis"

	"github.com/go-redis/redis/v8"
)

// Container는 모든 서비스와 핸들러 인스턴스를 포함합니다.
type Container struct {
	RedisClient *redis.Client

	AuthService    service.AuthService
	UserService    service.UserService
	ProductService service.ProductService

	AuthHandler    *handler.AuthHandler
	UserHandler    *handler.UserHandler
	ProductHandler *handler.ProductHandler
}

// BuildContainer는 모든 인스턴스를 생성하고 Container 구조체에 담아 반환합니다.
func BuildContainer() *Container {
	redisClient := redisClient.NewRedisClient()

	authService := service.NewAuthService()
	userService := service.NewUserService()
	productService := service.NewProductService()

	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService, redisClient)
	productHandler := handler.NewProductHandler(productService)

	return &Container{
		RedisClient: redisClient,

		AuthService:    authService,
		UserService:    userService,
		ProductService: productService,

		AuthHandler:    authHandler,
		UserHandler:    userHandler,
		ProductHandler: productHandler,
	}
}
