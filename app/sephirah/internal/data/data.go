package data

import (
	"context"
	"errors"

	"github.com/tuihub/librarian/app/sephirah/internal/ent"
	"github.com/tuihub/librarian/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewSqlClient, NewGreeterRepo)

// Data .
type Data struct {
	db *ent.Client
}

// NewData .
func NewData(db *ent.Client) *Data {
	return &Data{
		db,
	}
}

func NewSqlClient(c *conf.Sephirah_Data) (*ent.Client, func(), error) {
	var driverName, dataSourceName string
	driverName = c.Database.Driver
	switch driverName {
	case "sqlite3":
		dataSourceName = "file:ent?mode=memory&cache=shared&_fk=1"
	case "postgres":
		dataSourceName = "host=<host> port=<port> user=<user> dbname=<database> password=<pass>"
	default:
		return nil, func() {}, errors.New("unsupported sql database")
	}
	client, err := ent.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client, func() {
		client.Close()
	}, err
}
