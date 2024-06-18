package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	_ "login_module/docs"
	"login_module/internal/infrastructure/config"
	"login_module/internal/presentation/http/router"
	"os"
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
	// Load environment variables
	config.InitEnv()

	// Initialize logger
	config.InitLog()

	// Initialize database connection
	db := config.ConnectToDB()

	// Initialize Redis client
	redisClient := config.NewRedisClient()

	r := gin.Default()

	store, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte(os.Getenv("SESSION_SECRET")))
	if err != nil {
		panic(err)
	}
	r.Use(sessions.Sessions("cookie_api", store))

	// Configure CORS to handle requests from your React frontend
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           12 * time.Hour,
	}))

	// Initialize routes
	router.InitRoutes(r, db, redisClient)

	// Run the server
	if err := r.Run(":3000"); err != nil {
		return
	}
}
