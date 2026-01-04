package internal

import (
	v1 "github.com/tuihub/librarian/internal/data/internal/migration/internal/v1"

	"gorm.io/gorm"
)

type latestVersion v1.V1

func LatestVersion() int64 {
	v := latestVersion{}
	return v.Version()
}

func MigrationVersion(db *gorm.DB) error {
	v := &latestVersion{}
	return v.Migration(v, v.Version(), db)
}
