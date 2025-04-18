package libmq

import (
	stdsql "database/sql"

	"github.com/ThreeDotsLabs/watermill-sql/v3/pkg/sql"
)

func newSQLAdapter(db *stdsql.DB, loggerAdapter *mqLogger) (*pubSub, error) {
	sa := sql.DefaultPostgreSQLSchema{
		GenerateMessagesTableName: nil,
		GeneratePayloadType:       nil,
		SubscribeBatchSize:        0,
	}
	oa := sql.DefaultPostgreSQLOffsetsAdapter{
		GenerateMessagesOffsetsTableName: nil,
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
