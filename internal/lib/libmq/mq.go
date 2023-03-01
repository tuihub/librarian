package libmq

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/logger"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewMQ)

type MQ struct {
	router    *message.Router
	pubSub    *pubSub
	topicList map[string]bool
}

type pubSub struct {
	publisher  message.Publisher
	subscriber message.Subscriber
}

func NewMQ(c *conf.MQ) (*MQ, func(), error) {
	loggerAdapter := newMQLogger()
	var ps *pubSub
	switch c.Driver {
	case "memory":
		ps = newGoChannelAdapter(loggerAdapter)
	case "sql":
		var err error
		ps, err = newSQLAdapter(c.Database, loggerAdapter)
		if err != nil {
			return nil, func() {}, err
		}
	}
	router, err := message.NewRouter(
		message.RouterConfig{CloseTimeout: 0},
		loggerAdapter,
	)
	router.AddMiddleware(
		middleware.CorrelationID,
		middleware.Retry{
			MaxRetries:          5,                      //nolint:gomnd //TODO
			InitialInterval:     time.Millisecond * 100, //nolint:gomnd //TODO
			MaxInterval:         0,
			Multiplier:          0,
			MaxElapsedTime:      0,
			RandomizationFactor: 0,
			OnRetryHook:         nil,
			Logger:              loggerAdapter,
		}.Middleware,
		// middleware.Recoverer,
	)
	cleanup := func() {
		_ = router.Close()
	}
	return &MQ{
		router:    router,
		pubSub:    ps,
		topicList: make(map[string]bool),
	}, cleanup, err
}

func (a *MQ) Start(ctx context.Context) error {
	return a.router.Run(ctx)
}
func (a *MQ) Stop(ctx context.Context) error {
	return a.router.Close()
}

func (a *MQ) RegisterTopic(topic Topic) error {
	if _, exist := a.topicList[topic.Name()]; exist {
		return fmt.Errorf("topic %s already registered", topic)
	}
	a.topicList[topic.Name()] = true
	a.router.AddNoPublisherHandler(
		topic.Name(),
		topic.Name(),
		a.pubSub.subscriber,
		func(msg *message.Message) error {
			err := topic.Consume(msg.Context(), msg.Payload)
			if err != nil {
				return err
			}
			return nil
		},
	)
	topic.SetMQ(a)
	logger.Infof("topic %s registered", topic.Name())
	return nil
}

func (a *MQ) Publish(ctx context.Context, topic string, payload []byte) error {
	_, exist := a.topicList[topic]
	if !exist {
		return errors.New("unregistered topic")
	}
	msg := message.NewMessage(watermill.NewUUID(), payload)
	msg.SetContext(ctx)
	err := a.pubSub.publisher.Publish(topic, msg)
	for i := 0; err != nil && i < 16; i += 1 {
		err = a.pubSub.publisher.Publish(topic, msg)
	}
	return err
}
