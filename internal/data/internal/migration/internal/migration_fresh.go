package internal

import (
	"github.com/tuihub/librarian/internal/data/internal/models"

	"gorm.io/gorm"
)

func MigrationFresh(db *gorm.DB) error {
	return db.AutoMigrate(models.GetModels()...)
}
