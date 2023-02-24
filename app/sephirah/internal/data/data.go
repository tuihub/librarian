package data

import (
	"context"
	"errors"
	"fmt"

	"github.com/tuihub/librarian/app/sephirah/internal/data/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/ent"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/logger"

	"github.com/google/wire"

	_ "github.com/lib/pq"           // required by ent
	_ "github.com/mattn/go-sqlite3" // required by ent
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewSQLClient, NewTipherethRepo, NewGeburaRepo, NewYesodRepo)

// Data .
type Data struct {
	converter converter.Converter
	db        *ent.Client
}

// NewData .
func NewData(db *ent.Client) *Data {
	return &Data{
		converter: converter.NewConverter(),
		db:        db,
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
	if err = client.Schema.Create(context.Background()); err != nil {
		logger.Errorf("failed creating schema resources: %v", err)
		return nil, func() {}, err
	}
	return client, func() {
		client.Close()
	}, err
}
