package migrations

import (
	"awesomeProject/src/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Category{})
}
