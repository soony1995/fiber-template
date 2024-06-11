package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/oauth2"
	"login_module/internal/application/service"
	"login_module/internal/infrastructure/config"
	"login_module/internal/infrastructure/db"
	"login_module/internal/presentation/http/handler"
	"login_module/internal/presentation/http/middleware"
)

func InitRoutes(r *gin.Engine, redisClient *redis.Client) {
	authRepo := db.NewRedisAuthRepository(redisClient)

	// Initialize OAuth providers
	providers := map[string]*oauth2.Config{
		"google": config.GoogleOauthConfig,
		"kakao":  config.KakaoOauthConfig,
	}

	authService := service.NewAuthService(authRepo)
	oauthService := service.NewOAuthService(authRepo)

	authController := handler.NewAuthHandler(authService)
	oauthController := handler.NewOAuthHandler(oauthService, providers)

	// Set up routes
	api := r.Group("/api")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := api.Group("/auth")
	auth.POST("/login", authController.Login)
	auth.POST("/refresh_token", authController.RefreshToken)
	auth.POST("/logout", authController.Logout)

	oauth := api.Group("/oauth")
	oauth.GET("/:provider/login", oauthController.BeginOAuth)
	oauth.GET("/:provider/callback", oauthController.OAuthCallback)
	oauth.GET("/:provider/logout", middleware.ValidateIDToken(), oauthController.Logout)
}
