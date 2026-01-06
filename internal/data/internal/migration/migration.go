package migration

import (
	"errors"
	"fmt"
	"time"

	"github.com/tuihub/librarian/internal/data/internal/migration/internal"
	"github.com/tuihub/librarian/internal/lib/logger"

	"gorm.io/gorm"
)

// SchemaVersion represents the schema version table.
type SchemaVersion struct {
	Version   int64     `gorm:"primaryKey"`
	AppliedAt time.Time `gorm:"autoCreateTime"`
}

// TableName specifies the table name for SchemaVersion.
func (SchemaVersion) TableName() string {
	return "schema_versions"
}

func Migrate(db *gorm.DB) error {
	if db == nil {
		return errors.New("database connection is nil")
	}

	if err := db.AutoMigrate(&SchemaVersion{}); err != nil {
		return fmt.Errorf("failed to ensure version table: %w", err)
	}

	current := getCurrentVersion(db)
	latest := internal.LatestVersion()

	logger.Infof("Current database version: %d, Latest version: %d", current, latest)

	if current > latest {
		return fmt.Errorf("database version %d is newer than application supports (%d)", current, latest)
	}

	if current == latest {
		logger.Infof("Database is up to date")
		return nil
	}

	// Fresh install: use GORM AutoMigrate
	if current == 0 {
		logger.Infof("Performing fresh installation using GORM AutoMigrate")
		return db.Transaction(func(tx *gorm.DB) error {
			if err := internal.MigrationFresh(db); err != nil {
				return fmt.Errorf("auto migrate failed: %w", err)
			}
			return setVersion(tx, latest)
		})
	}

	// Incremental upgrade: replay migrations step by step
	logger.Infof("Performing incremental migration from version %d to %d", current, latest)
	return db.Transaction(func(tx *gorm.DB) error {
		if err := internal.MigrationVersion(tx); err != nil {
			return fmt.Errorf("incremental migration failed: %w", err)
		}
		return setVersion(tx, latest)
	})
}

// getCurrentVersion retrieves the current version number from the database.
func getCurrentVersion(db *gorm.DB) int64 {
	var result SchemaVersion
	err := db.Order("version desc").Limit(1).First(&result).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// No records, indicating fresh install
			return 0
		}
		logger.Warnf("Failed to get current version, assuming fresh install: %v", err)
		return 0
	}

	return result.Version
}

// setVersion updates the database version number.
func setVersion(db *gorm.DB, version int64) error {
	return db.Create(&SchemaVersion{
		Version: version,
	}).Error
}
