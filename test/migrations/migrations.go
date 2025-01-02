package migrations

import (
	"test/database"
	"test/models"
)

func Migrate() {
  database.DB.AutoMigrate(&models.User{})
}
