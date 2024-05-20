package main

import (
	"log"
	"login_module/internal/container"
	"login_module/internal/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	ctn := container.BuildContainer()
	app := fiber.New()

	route.SetupRoutes(app, ctn)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
