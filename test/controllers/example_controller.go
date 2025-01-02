// Package controllers provides the necessary functions and methods to handle
// various routes in the test application. This package is responsible for
// managing the routing logic and ensuring that the appropriate controllers are
// invoked for different endpoints.
package controllers

import (
  "test/models"
  "test/repository"
  "encoding/json"
  "net/http"
  "github.com/gin-gonic/gin"
)
func Welcome(c *gin.Context) {
  c.String(200, "Welcome to test!")
}

// CreateUser is an example function that handles the creation of a new user.
// It expects a JSON payload in the request body, which it parses into a User model.
// If the JSON parsing fails, it returns a 400 Bad Request status with an error message.
// If the user creation in the repository fails, it returns a 500 Internal Server Error status with an error message.
// On successful creation, it returns a 201 Created status with the created user in the response.
func CreateUser(c *gin.Context) {
  var user models.User

  if err := c.ShouldBindJSON(&user); err != nil {
    c.JSON(400, gin.H{"error": "Cannot parse JSON"})
    return
  }

  if err := repository.CreateUser(&user); err != nil {
    c.JSON(500, gin.H{"error": "Cannot create user"})
    return
  }

  c.JSON(201, user)
}

func RegisterRoutes(router *gin.Engine) {
  router.GET("/", Welcome)
  router.POST("/users", CreateUser)
}
