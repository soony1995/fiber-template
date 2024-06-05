package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	_ "login_module/docs"
	"login_module/internal/infrastructure/container"
	"login_module/internal/presentation/http"
	"time"
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

	// Configure CORS to handle requests from your React frontend
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           12 * time.Hour,
	}))

	http.SetupRoutes(r, ctn)

	err := r.Run(":3000")
	if err != nil {
		return
	}
}
