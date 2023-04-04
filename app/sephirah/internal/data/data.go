package data

import (
	"context"
	"errors"
	"fmt"

	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/migrate"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/logger"

	"entgo.io/ent/dialect/sql"
	"github.com/google/wire"
	"golang.org/x/exp/slices"

	_ "github.com/lib/pq"           // required by ent
	_ "github.com/mattn/go-sqlite3" // required by ent
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewSQLClient,
	NewTipherethRepo,
	NewGeburaRepo,
	NewYesodRepo,
	NewNetzachRepo,
)

// Data .
type Data struct {
	db *ent.Client
}

// NewData .
func NewData(db *ent.Client) *Data {
	return &Data{
		db: db,
	}
}

func NewSQLClient(c *conf.Sephirah_Data) (*ent.Client, func(), error) {
	var driverName, dataSourceName string
	driverName = c.Database.Driver
	switch driverName {
	case "sqlite3":
		dataSourceName = "file:ent?mode=memory&cache=shared&_fk=1"
	case "postgres":
		dataSourceName = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
			c.Database.Host,
			c.Database.Port,
			c.Database.User,
			c.Database.Dbname,
			c.Database.Password,
		)
		if c.Database.NoSsl {
			dataSourceName += " sslmode=disable"
		}
	default:
		return nil, func() {}, errors.New("unsupported sql database")
	}
	client, err := ent.Open(driverName, dataSourceName)
	if err != nil {
		logger.Errorf("failed opening connection to postgres: %v", err)
		return nil, func() {}, err
	}
	// Run the auto migration tool.
	if err = client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		logger.Errorf("failed creating schema resources: %v", err)
		return nil, func() {}, err
	}
	return client, func() {
		client.Close()
	}, err
}

func (data *Data) WithTx(ctx context.Context, fn func(tx *ent.Tx) error) error {
	tx, err := data.db.Tx(ctx)
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
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
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
