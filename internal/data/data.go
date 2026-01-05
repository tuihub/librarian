package data

import (
	"context"
	stdsql "database/sql"
	"errors"
	"fmt"
	"path"
	"time"

	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/data/internal/migration"
	"github.com/tuihub/librarian/internal/data/internal/query"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/logger"

	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(
	NewData,
	GetDB,
	NewAngelaRepo,
	NewTipherethRepo,
	NewGeburaRepo,
	NewYesodRepo,
	NewNetzachRepo,
	NewChesedRepo,
	NewKetherRepo,
	NewBinahRepo,
	NewSupervisorRepo,
)

type Data struct {
	stdDB *stdsql.DB
	db    *gorm.DB
}

func NewData(c *conf.Database, app *libapp.Settings) (*Data, func(), error) {
	if c == nil {
		return nil, func() {}, errors.New("database config is nil")
	}
	var dialector gorm.Dialector
	switch c.Driver {
	case conf.DatabaseDriverMemory:
		dialector = sqlite.Open("file::memory:?cache=shared")
	case conf.DatabaseDriverSqlite:
		dialector = sqlite.Open(fmt.Sprintf(
			"file:%s?cache=shared&_journal=WAL&_busy_timeout=30000",
			path.Join(app.DataPath, "librarian.db"),
		))
	case conf.DatabaseDriverPostgres:
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
			c.Host, c.Username, c.Password, c.DBName, c.Port)
		dialector = postgres.Open(dsn)
	default:
		return nil, func() {}, fmt.Errorf("unsupported database driver %s", c.Driver)
	}

	db, err := gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		logger.Errorf("failed opening connection to database: %v", err)
		return nil, func() {}, fmt.Errorf("failed opening connection to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, func() {}, err
	}
	sqlDB.SetMaxIdleConns(10)  //nolint:mnd // default
	sqlDB.SetMaxOpenConns(100) //nolint:mnd // default
	sqlDB.SetConnMaxIdleTime(time.Hour)

	// Run database migrations
	if err = migration.Migrate(db); err != nil {
		logger.Errorf("failed running database migration: %v", err)
		return nil, func() {}, fmt.Errorf("failed running database migration: %w", err)
	}

	query.SetDefault(db)

	return &Data{
			stdDB: sqlDB,
			db:    db,
		}, func() {
			_ = sqlDB.Close()
		}, nil
}

func GetDB(d *Data) *stdsql.DB {
	return d.stdDB
}

func (d *Data) WithTx(ctx context.Context, fn func(tx *query.Query) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(query.Use(tx))
	})
}

func ErrorIsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
