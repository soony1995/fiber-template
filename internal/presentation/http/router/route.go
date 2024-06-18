package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"login_module/internal/application/service"
	"login_module/internal/infrastructure/db"
	"login_module/internal/presentation/http/handler"
)

func InitRoutes(r *gin.Engine, gorm *gorm.DB, redis *redis.Client) {
	userRepo := db.NewMySQLUserRepository(gorm)

	authService := service.NewAuthService(userRepo, redis)
	authController := handler.NewAuthHandler(authService)

	// Set up routes
	api := r.Group("/api")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := api.Group("/auth")
	auth.POST("/login", authController.Login)
	auth.POST("/logout", authController.Logout)

}
