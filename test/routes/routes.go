// Package routes provides the routing configuration for the test application.
// It includes the necessary imports and the function to register all routes.
package routes

import (
    "test/controllers"
    "github.com/gofiber/fiber/v2"
    "github.com/redis/go-redis/v9"
)

// RegisterRoutes sets up the routes for the application using the provided
// fiber app instance, and Redis client.
//
// Parameters:
// - app: A pointer to the fiber app instance.
// - redisClient: A pointer to the Redis client.
func RegisterRoutes(app *fiber.App, redisClient *redis.Client) {
  controllers.RegisterRoutes(app)
}
