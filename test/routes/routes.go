// Package routes provides the routing configuration for the test application.
// It includes the necessary imports and the function to register all routes.
package routes

import (
    "test/controllers"
    "github.com/gin-gonic/gin"
    "github.com/redis/go-redis/v9"
)

// RegisterRoutes sets up the routes for the application using the provided
// gin app instance, and Redis client.
//
// Parameters:
// - app: A pointer to the gin app instance.
// - redisClient: A pointer to the Redis client.
func RegisterRoutes(app *gin.Engine, redisClient *redis.Client) {
  controllers.RegisterRoutes(app)
}
