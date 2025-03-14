package migration

import "gorm.io/gorm"

type latestVersion struct {
	v2
}

type Migration struct {
	migrator latestVersion
}

func NewMigration(db *gorm.DB) *Migration {
	m := latestVersion{}
	m.init(db)
	return &Migration{
		migrator: m,
	}
}

func (m *Migration) Migrate() error {
	return m.migrator.migrate()
}
