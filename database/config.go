// Package database provides functionality to initialize and manage the database connection.
// 
// This package uses GORM as the ORM and connects to a PostgreSQL database using environment variables
// for configuration.
//
// The following environment variables are expected:
// - DB_HOST: The hostname of the database server.
// - DB_USER: The username to connect to the database.
// - DB_PASSWORD: The password to connect to the database.
// - DB_NAME: The name of the database.
// - DB_PORT: The port on which the database server is running.
//
// The InitDB function initializes the database connection using the provided environment variables
// and assigns the connection to the global variable DB. If the connection fails, the function logs
// a fatal error and terminates the application.
package database

import (
	"fmt"
	"os"
  "log"

  "gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB

func InitDB() {
  dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
    os.Getenv("DB_HOST"),
    os.Getenv("DB_USER"),
    os.Getenv("DB_PASSWORD"),
    os.Getenv("DB_NAME"),
    os.Getenv("DB_PORT"),
  )

  database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    log.Fatalf("Error connecting to database: %v", err)
  }

  DB = database
}
