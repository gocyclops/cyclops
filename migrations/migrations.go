package migrations

import (
	"cyclops/database"
	"cyclops/models"
)

func Migrate() {
  database.DB.AutoMigrate(&models.User{})
}
