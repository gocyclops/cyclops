package main

import (
  "log"
  "net/http"
  "os"

  "test/database"
  "test/migrations"
  "test/myredis"
  "test/mys3"
  "test/routes"
  "github.com/gin-gonic/gin"

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
  // Initialize Gin router
  r := gin.Default()
  r.Use(gin.Logger())
  r.Use(gin.Recovery())

  // Register all routes
  routes.RegisterRoutes(r, myredis.RedisClient)

  // Set up cron job
  c := cron.New()
  c.AddFunc("@daily", func() {
    // Add cron jobs
  })
  c.Start()

  defer c.Stop()
  // Start the server
  if err := r.Run(":8080"); err != nil {
    log.Fatalf("Failed to run server: %v", err)
  }
}
