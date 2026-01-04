package v1

import (
	"fmt"

	"github.com/tuihub/librarian/internal/data/internal/migration/internal/v0"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "embed"
)

//go:embed sqlite.sql
var v1SQLiteSQL string

//go:embed postgres.sql
var v1PostgresSQL string

type V1 struct {
	v0.V0
}

func (v *V1) Version() int64 {
	return 1
}

func (v *V1) GetPrev() v0.Version {
	return &v.V0
}

// Up upgrades from v0 to v1.
func (v *V1) Up(db *gorm.DB) error {
	var sqlContent string

	switch db.Dialector.(type) {
	case *sqlite.Dialector:
		sqlContent = v1SQLiteSQL
	case *postgres.Dialector:
		sqlContent = v1PostgresSQL
	default:
		return fmt.Errorf("unsupported database driver: %s", db.Dialector.Name())
	}

	return v.ExecuteSQL(db, sqlContent)
}
