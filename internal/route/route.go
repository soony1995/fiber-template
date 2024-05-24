package route

import (
	"login_module/internal/container"
	"login_module/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, ctn *container.Container) {
	// /api 그룹 생성
	api := app.Group("/api")

	// 로그인 라우트
	authGroup := api.Group("/auth")
	authGroup.Post("/login", ctn.AuthHandler.Login)
	authGroup.Post("/refresh_token", ctn.AuthHandler.RefreshToken)
	authGroup.Post("/logout", ctn.AuthHandler.Logout)

	// 사용자 라우트 그룹
	userGroup := api.Group("/user")
	userGroup.Get("/:id", middleware.TokenValid, ctn.UserHandler.GetUser)
}
