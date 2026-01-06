package v0

import (
	"fmt"

	"gorm.io/gorm"
)

// Version defines the migration version interface.
type Version interface {
	// Version returns the current version number
	Version() int64
	// GetPrev returns the previous versioner
	GetPrev() Version
	// Up executes incremental migration from previous version to current version
	Up(*gorm.DB) error
	// Migration executes full migration to current version
	Migration(Version, int64, *gorm.DB) error
}

// V0 base version, does not contain actual SQL, only used for tooling.
type V0 struct{}

func (v *V0) Version() int64 {
	return 0
}

func (v *V0) GetPrev() Version {
	return nil
}

func (v *V0) Up(_ *gorm.DB) error {
	return nil
}

func (v *V0) Migration(current Version, target int64, db *gorm.DB) error {
	if current.Version() > target {
		return nil
	}
	if current.Version() < target {
		prev := current.GetPrev()
		if prev == nil {
			return fmt.Errorf("no previous version for v%d", current.Version())
		}
		if prev.Version()+1 != current.Version() {
			return fmt.Errorf("invalid version chain: v%d -> v%d", prev.Version(), current.Version())
		}
		if err := prev.Migration(prev, target, db); err != nil {
			return err
		}
	}
	if err := current.Up(db); err != nil {
		return fmt.Errorf("failed to migrate to v%d: %w", current.Version(), err)
	}
	return nil
}

func (v *V0) ExecuteSQL(db *gorm.DB, content string) error {
	if _, err := db.Statement.ConnPool.ExecContext(db.Statement.Context, content); err != nil {
		return fmt.Errorf("failed to execute SQL: %w", err)
	}

	return nil
}
