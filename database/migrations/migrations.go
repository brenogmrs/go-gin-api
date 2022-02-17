package migrations

import (
	"github.com/brenogmrs/go-gin-api/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
