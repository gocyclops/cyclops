// Package models contains the data structures and related functions
// used in the Cyclops application. This package is responsible for
// defining the models that represent the core entities of the application.
package models

import "gorm.io/gorm"

// Example model
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"unique"`
	Age   int
}
