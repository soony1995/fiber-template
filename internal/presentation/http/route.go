package http

import (
	"login_module/internal/infrastructure/container"
	"login_module/internal/presentation/http/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRoutes(router *gin.Engine, ctn *container.Container) {
	api := router.Group("/api")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := api.Group("/auth")
	auth.POST("/login", ctn.AuthHandler.Login)
	auth.POST("/refresh_token", ctn.AuthHandler.RefreshToken)
	auth.POST("/logout", ctn.AuthHandler.Logout)

	oauth := api.Group("/oauth")
	oauth.GET("/:provider/login", ctn.OAuthHandler.BeginOAuth)
	oauth.GET("/:provider/callback", ctn.OAuthHandler.OAuthCallback)
	oauth.GET("/:provider/logout", middleware.ValidateIDToken(), ctn.OAuthHandler.Logout)

	users := api.Group("/users")
	users.Use(middleware.ValidateIDToken())
}
