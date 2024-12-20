package config

import (
	"fmt"
	handler "realtime-score/internal/handlers"
	"realtime-score/internal/migration"
	"realtime-score/internal/repository"
	"realtime-score/internal/router"
	"realtime-score/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type AppConfig struct {
	app *gin.Engine
	log *logrus.Logger
}

func NewApp() AppConfig {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	app := gin.Default()
	log := logrus.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Replace with your front-end URL if needed
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	// Handle OPTIONS pre-flight requests for CORS
	app.OPTIONS("/*any", func(c *gin.Context) {
		c.AbortWithStatus(204)
	})
	db := DBNew()

	// Migration and seeder
	if err := migration.Migrate(db); err != nil {
		log.Printf("Error migrating database: %v", err)
	}
	if err = migration.Seeder(db); err != nil {
		log.Printf("Error seeding database: %v", err)
	}

	// Repositories
	userRepository := repository.NewUser(db)

	// Services
	userService := services.NewUser(userRepository)

	// Initialize handlers
	userHandler := handler.NewUser(log, userService)

	// Testing routes
	app.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Set up routes
	router.User(app, userHandler)
	router.TestingRouter(app)

	return AppConfig{
		app: app,
		log: log,
	}
}

func (ap *AppConfig) Run() {
	// Start the application
	if err := ap.app.Run(":8080"); err != nil {
		ap.log.Fatalf("Failed to start server: %v", err)
	}
}
