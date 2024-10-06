package main

import (
	"cyclops/database"
	"cyclops/migrations"
	"cyclops/myredis"
	"cyclops/mys3"
	"cyclops/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize the database connection
	database.InitDB()
	migrations.Migrate()

	// Initialize the redis connection
	myredis.InitRedis()

	// Initialize S3 connection
	mys3.InitS3()
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")

	// Initialize Fiber router
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	// Register all routes
	routes.RegisterRoutes(app, myredis.RedisClient)

	// Set up cron job to delete expired folders
	c := cron.New()
	c.AddFunc("@daily", func() {
		// Add cron jobs
	})
	c.Start()

	defer c.Stop()

	// Start the server
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
