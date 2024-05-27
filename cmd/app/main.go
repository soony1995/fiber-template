package main

import (
	"log"
	"login_module/internal/infrastructure/container"
	"login_module/internal/presentation/http"

	_ "login_module/docs"

	"github.com/gofiber/fiber/v2"
)

//	@title			Order Api
//	@version		1.0
//	@description	This is an Order Api just for young people
//	@termsOfService	http://swagger.io/terms/

func main() {
	app := fiber.New()
	ctn := container.BuildContainer()

	// http.SetupRoutes(app, ctn)
	http.SetupRoutes(app, ctn)
	if err := app.Listen(":3000"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
