package http

import (
	"login_module/internal/infrastructure/container"
	"login_module/internal/presentation/http/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App, ctn *container.Container) {
	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/", ctn.AuthHandler.Login)
	auth.Post("/refresh_token", ctn.AuthHandler.RefreshToken)
	auth.Post("/logout", ctn.AuthHandler.Logout)

	users := api.Group("/users")
	users.Get("/:id", middleware.TokenValid, ctn.UserHandler.GetUser)
	users.Post("/", ctn.UserHandler.CreateUser)
}
