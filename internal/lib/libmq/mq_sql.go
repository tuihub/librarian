package libmq

import (
	stdSql "database/sql"
	"errors"
	"fmt"

	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/logger"

	"github.com/ThreeDotsLabs/watermill-sql/v3/pkg/sql"

	_ "github.com/mattn/go-sqlite3" // required
)

func newSQLAdapter(c *conf.Database, loggerAdapter *mqLogger) (*pubSub, error) {
	var driverName, dataSourceName string
	driverName = c.GetDriver()
	var sa sql.SchemaAdapter
	var oa sql.OffsetsAdapter
	switch driverName {
	case "postgres":
		dataSourceName = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
			c.GetHost(),
			c.GetPort(),
			c.GetUser(),
			c.GetDbname(),
			c.GetPassword(),
		)
		if c.GetNoSsl() {
			dataSourceName += " sslmode=disable"
		}
		sa = sql.DefaultPostgreSQLSchema{
			GenerateMessagesTableName: nil,
			GeneratePayloadType:       nil,
			SubscribeBatchSize:        0,
		}
		oa = sql.DefaultPostgreSQLOffsetsAdapter{
			GenerateMessagesOffsetsTableName: nil,
		}
	default:
		return nil, errors.New("unsupported sql database")
	}
	db, err := stdSql.Open(driverName, dataSourceName)
	if err != nil {
		logger.Errorf("failed opening connection to postgres: %v", err)
		return nil, err
	}

	subscriber, err := sql.NewSubscriber(db, sql.SubscriberConfig{ //nolint:exhaustruct // no need
		SchemaAdapter:    sa,
		OffsetsAdapter:   oa,
		InitializeSchema: true,
	}, loggerAdapter)
	if err != nil {
		return nil, err
	}
	publisher, err := sql.NewPublisher(db, sql.PublisherConfig{
		SchemaAdapter:        sa,
		AutoInitializeSchema: true,
	}, loggerAdapter)
	if err != nil {
		return nil, err
	}
	return &pubSub{publisher, subscriber}, nil
}
