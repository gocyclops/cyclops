// Package repository provides functions to interact with the user data in the database.
// This is an example of how to create and retrieve user records using GORM.

/// CreateUser inserts a new user record into the database.
/// 
/// Parameters:
/// - user: A pointer to the User model that needs to be created.
///
/// Returns:
/// - error: An error object if the operation fails, otherwise nil.
///
/// Example:
/// ```go
/// newUser := &models.User{Name: "John Doe", Email: "john.doe@example.com"}
/// err := repository.CreateUser(newUser)
/// if err != nil {
///     log.Fatalf("Failed to create user: %v", err)
/// }
/// ```

/// GetUserByID retrieves a user record from the database by its ID.
/// 
/// Parameters:
/// - id: The unique identifier of the user to be retrieved.
///
/// Returns:
/// - models.User: The user model corresponding to the given ID.
/// - error: An error object if the operation fails, otherwise nil.
///
/// Example:
/// ```go
/// userID := uint(1)
/// user, err := repository.GetUserByID(userID)
/// if err != nil {
///     log.Fatalf("Failed to get user: %v", err)
/// }
/// fmt.Printf("User: %+v\n", user)
/// ```
package repository

import (
    "cyclops/models"
    "cyclops/database"
)

func CreateUser(user *models.User) error {
  return database.DB.Create(user).Error
}

func GetUserByID(id uint) (models.User, error) {
  var user models.User
  err := database.DB.First(&user, id).Error
  return user, err
}
