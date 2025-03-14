package migration_test

import (
	"github.com/tuihub/librarian/app/sephirah/internal/data/migration"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestNewMigration(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	err = migration.NewMigration(db).Migrate()
	if err != nil {
		t.Fatalf("Failed to migrate: %v", err)
	}
}
