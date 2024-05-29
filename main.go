package main

import (
	_ "login_module/docs"
	"login_module/internal/infrastructure/container"
	"login_module/internal/presentation/http"
	"login_module/internal/presentation/http/handler"

	"github.com/gin-gonic/gin"
)

// @title Login Module API
// @version 1.0
// @description This is a sample server for a login module.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /api
func main() {
	r := gin.Default()
	ctn := container.BuildContainer()

	handler.InitOAuth()
	http.SetupRoutes(r, ctn)

	r.Run(":3000")
}
