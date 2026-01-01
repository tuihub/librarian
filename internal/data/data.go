package data

import (
	"context"
	stdsql "database/sql"
	"errors"
	"fmt"
	"net"
	"path"
	"strconv"
	"time"

	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/data/internal/gormschema"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/logger"

	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	_ "github.com/jackc/pgx/v5/stdlib" // required by postgres
	_ "github.com/mattn/go-sqlite3"    // required by sqlite
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
	var dataSourceName string
	if c == nil {
		return nil, func() {}, errors.New("database config is nil")
	}
	driverName := c.Driver

	gormConfig := &gorm.Config{
		Logger:                                   gormlogger.Default.LogMode(gormlogger.Silent),
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	var dialector gorm.Dialector
	switch driverName {
	case conf.DatabaseDriverMemory:
		driverName = conf.DatabaseDriverSqlite
		dataSourceName = "file:librarian?mode=memory&cache=shared&_fk=1&_busy_timeout=30000&_timeout=30000"
		dialector = sqlite.Open(dataSourceName)
	case conf.DatabaseDriverSqlite:
		dataSourceName = fmt.Sprintf(
			"file:%s?cache=shared&_fk=1&_journal=WAL&_busy_timeout=30000",
			path.Join(app.DataPath, "librarian.db"),
		)
		dialector = sqlite.Open(dataSourceName)
	case conf.DatabaseDriverPostgres:
		dataSourceName = fmt.Sprintf("postgresql://%s:%s@%s/%s",
			c.Username,
			c.Password,
			net.JoinHostPort(c.Host, strconv.Itoa(int(c.Port))),
			c.DBName,
		)
		dialector = postgres.Open(dataSourceName)
	default:
		return nil, func() {}, fmt.Errorf("unsupported database driver %s", driverName)
	}

	db, err := gorm.Open(dialector, gormConfig)
	if err != nil {
		logger.Errorf("failed opening connection to database: %v", err)
		return nil, func() {}, fmt.Errorf("failed opening connection to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Errorf("failed getting underlying sql.DB: %v", err)
		return nil, func() {}, fmt.Errorf("failed getting underlying sql.DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)  //nolint:mnd // no need
	sqlDB.SetMaxOpenConns(100) //nolint:mnd // no need
	sqlDB.SetConnMaxIdleTime(time.Hour)

	// Run auto migration (without foreign keys)
	if err = db.AutoMigrate(gormschema.AllModels()...); err != nil {
		logger.Errorf("failed creating schema resources: %v", err)
		return nil, func() {}, fmt.Errorf("failed creating schema resources: %w", err)
	}

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

// WithTx executes a function within a database transaction.
// If the function returns an error or panics, the transaction is rolled back.
// On success, the transaction is committed.
func (d *Data) WithTx(ctx context.Context, fn func(tx *gorm.DB) error) error {
	tx := d.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()

	if err := fn(tx); err != nil {
		if rerr := tx.Rollback().Error; rerr != nil {
			return fmt.Errorf("%w: rolling back transaction: %s", err, rerr.Error())
		}
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}

// ErrorIsNotFound returns true if the error indicates a record was not found.
func ErrorIsNotFound(err error) bool {
	if err == nil {
		return false
	}
	return errors.Is(err, gorm.ErrRecordNotFound)
}
