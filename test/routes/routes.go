// Package routes provides the routing configuration for the test application.
// It includes the necessary imports and the function to register all routes.
package routes

import (
    "test/controllers"
    "github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the routes for the application using the provided
// gin app instance.
//
// Parameters:
// - app: A pointer to the gin app instance.
func RegisterRoutes(app *gin.Engine) {
  controllers.RegisterRoutes(app)
}
