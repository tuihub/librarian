package libmq

import (
	"github.com/tuihub/librarian/internal/conf"

	"github.com/ThreeDotsLabs/watermill-redisstream/pkg/redisstream"
	"github.com/redis/go-redis/v9"
)

func newRedisAdapter(c *conf.Cache, loggerAdapter *mqLogger) (*pubSub, error) {
	client := redis.NewClient(&redis.Options{ //nolint:exhaustruct // no need
		Addr:       c.GetAddr(),
		DB:         int(c.GetDb()),
		Username:   c.GetUser(),
		Password:   c.GetPassword(),
		MaxRetries: -1, // Use middleware to handle retry
	})
	subScriber, err := redisstream.NewSubscriber(
		redisstream.SubscriberConfig{ //nolint:exhaustruct // no need
			Client: client,
		},
		loggerAdapter,
	)
	if err != nil {
		return nil, err
	}
	publisher, err := redisstream.NewPublisher(
		redisstream.PublisherConfig{ //nolint:exhaustruct // no need
			Client: client,
		},
		loggerAdapter,
	)
	if err != nil {
		return nil, err
	}
	return &pubSub{publisher, subScriber}, nil
}
