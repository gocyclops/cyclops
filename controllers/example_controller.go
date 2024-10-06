// Package controllers provides the necessary functions and methods to handle
// various routes in the Cyclops application. This package is responsible for
// managing the routing logic and ensuring that the appropriate controllers are
// invoked for different endpoints.
package controllers

import (
	"cyclops/models"
	"cyclops/repository"

	"github.com/gofiber/fiber/v2"
)

func Welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to Cyclops!")
}

// CreateUser is an example function that handles the creation of a new user.
// It expects a JSON payload in the request body, which it parses into a User model.
// If the JSON parsing fails, it returns a 400 Bad Request status with an error message.
// If the user creation in the repository fails, it returns a 500 Internal Server Error status with an error message.
// On successful creation, it returns a 201 Created status with the created user in the response.
func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if err := repository.CreateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot create user"})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func RegisterRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return Welcome(c)
	})
}
