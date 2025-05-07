package data

import (
	"context"
	stdsql "database/sql"
	"errors"
	"fmt"
	"net"
	"path"
	"slices"
	"strconv"
	"time"

	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/data/internal/ent"
	"github.com/tuihub/librarian/internal/data/internal/ent/migrate"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/logger"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/google/wire"

	_ "github.com/jackc/pgx/v5/stdlib" // required by ent
	_ "github.com/mattn/go-sqlite3"    // required by ent
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
)

type Data struct {
	stdDB *stdsql.DB
	db    *ent.Client
}

func NewData(c *conf.Database, app *libapp.Settings) (*Data, func(), error) {
	var dataSourceName string
	if c == nil {
		return nil, func() {}, errors.New("database config is nil")
	}
	driverName := c.Driver
	var dialectName string
	switch driverName {
	case conf.DatabaseDriverMemory:
		dialectName = dialect.SQLite
		driverName = conf.DatabaseDriverSqlite
		dataSourceName = "file:ent?mode=memory&cache=shared&_fk=1"
	case conf.DatabaseDriverSqlite:
		dialectName = dialect.SQLite
		dataSourceName = fmt.Sprintf("file:%s?cache=shared&_fk=1&_journal=WAL", path.Join(app.DataPath, "librarian.db"))
	case conf.DatabaseDriverPostgres:
		dialectName = dialect.Postgres
		driverName = "pgx"
		dataSourceName = fmt.Sprintf("postgresql://%s:%s@%s/%s",
			c.Username,
			c.Password,
			net.JoinHostPort(c.Host, strconv.Itoa(int(c.Port))),
			c.DBName,
		)
	default:
		return nil, func() {}, fmt.Errorf("unsupported database driver %s", driverName)
	}

	db, err := stdsql.Open(string(driverName), dataSourceName)
	if err != nil {
		logger.Errorf("failed opening connection to database: %v", err)
		return nil, func() {}, fmt.Errorf("failed opening connection to database: %w", err)
	}
	drv := sql.OpenDB(dialectName, db)

	db.SetMaxIdleConns(10)  //nolint:mnd // no need
	db.SetMaxOpenConns(100) //nolint:mnd // no need
	db.SetConnMaxIdleTime(time.Hour)

	client := ent.NewClient(ent.Driver(drv))

	// Run the auto migration tool.
	if err = client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		logger.Errorf("failed creating schema resources: %v", err)
		return nil, func() {}, fmt.Errorf("failed creating schema resources: %w", err)
	}

	return &Data{
			stdDB: db,
			db:    client,
		}, func() {
			_ = client.Close()
			_ = db.Close()
		}, nil
}

func GetDB(d *Data) *stdsql.DB {
	return d.stdDB
}

func (d *Data) WithTx(ctx context.Context, fn func(tx *ent.Tx) error) error {
	tx, err := d.db.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			_ = tx.Rollback()
			panic(v)
		}
	}()
	if err = fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %s", err, rerr.Error())
		}
		return err
	}
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}

func resolveWithIgnores(ignores []string) sql.ConflictOption {
	return sql.ResolveWith(func(u *sql.UpdateSet) {
		for _, c := range u.Columns() {
			if slices.Contains(ignores, c) {
				continue
			}
			u.SetExcluded(c)
		}
	})
}
