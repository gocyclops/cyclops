// Package routes provides the routing configuration for the Cyclops application.
// It includes the necessary imports and the function to register all routes.
package routes

import (
	"cyclops/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

// RegisterRoutes sets up the routes for the application using the provided
// Fiber app instance, PostgreSQL connection pool, and Redis client.
//
// Parameters:
// - app: A pointer to the Fiber app instance.
// - db: A pointer to the PostgreSQL connection pool.
// - redisClient: A pointer to the Redis client.
func RegisterRoutes(app *fiber.App, redisClient *redis.Client) {
	controllers.RegisterRoutes(app)
}
