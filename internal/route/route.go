package route

import (
	"login_module/internal/container"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, ctn *container.Container) {
	// /api 그룹 생성
	api := app.Group("/api")

	// 로그인 라우트
	api.Post("/login", ctn.AuthHandler.Login)

	// 사용자 라우트 그룹
	userGroup := api.Group("/user")
	userGroup.Get("/:id", ctn.UserHandler.GetUser)
	userGroup.Post("/", ctn.UserHandler.CreateUser)

	// 제품 라우트 그룹
	productGroup := api.Group("/product")
	productGroup.Get("/:id", ctn.ProductHandler.GetProduct)
	productGroup.Post("/", ctn.ProductHandler.CreateProduct)
}
