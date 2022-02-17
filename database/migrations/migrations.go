package migrations

import (
	"github.com/brenogmrs/go-gin-api/models"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
