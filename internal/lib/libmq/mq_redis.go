package libmq

import (
	"net"
	"strconv"

	"github.com/tuihub/librarian/internal/conf"

	"github.com/ThreeDotsLabs/watermill-redisstream/pkg/redisstream"
	"github.com/redis/go-redis/v9"
)

func newRedisAdapter(c *conf.Cache, loggerAdapter *mqLogger) (*pubSub, error) {
	client := redis.NewClient(&redis.Options{ //nolint:exhaustruct // no need
		Addr:       net.JoinHostPort(c.Host, strconv.Itoa(int(c.Port))),
		DB:         int(c.DB),
		Username:   c.Username,
		Password:   c.Password,
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
