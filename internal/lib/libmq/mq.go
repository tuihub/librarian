package libmq

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libobserve"
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
	observer  *libobserve.BuiltInObserver
}

type pubSub struct {
	publisher  message.Publisher
	subscriber message.Subscriber
}

func NewMQ(
	c *conf.MQ,
	db *sql.DB,
	cachec *conf.Cache,
	app *libapp.Settings,
	obs *libobserve.BuiltInObserver,
) (*MQ, func(), error) {
	loggerAdapter := newMQLogger()
	var ps *pubSub
	if c == nil {
		return nil, func() {}, errors.New("mq config is nil")
	}
	switch c.Driver {
	case conf.MQDriverMemory:
		ps = newGoChannelAdapter(loggerAdapter)
	case conf.MQDriverSQL:
		var err error
		ps, err = newSQLAdapter(db, loggerAdapter)
		if err != nil {
			return nil, func() {}, fmt.Errorf("failed creating sql adapter: %w", err)
		}
	case conf.MQDriverRedis:
		if cachec == nil || cachec.Driver != conf.CacheDriverRedis {
			return nil, func() {}, errors.New("invalid redis driver for mq")
		}
		var err error
		ps, err = newRedisAdapter(cachec, loggerAdapter)
		if err != nil {
			return nil, func() {}, fmt.Errorf("failed creating redis adapter: %w", err)
		}
	default:
		return nil, func() {}, fmt.Errorf("unsupported mq driver: %s", c.Driver)
	}
	router, err := message.NewRouter(
		message.RouterConfig{CloseTimeout: 0},
		loggerAdapter,
	)
	if err != nil {
		return nil, func() {}, fmt.Errorf("failed creating router: %w", err)
	}
	router.AddMiddleware(middleware.CorrelationID)
	if app.EnablePanicRecovery {
		router.AddMiddleware(middleware.Recoverer)
	}
	return &MQ{
		router:    router,
		pubSub:    ps,
		topicList: make(map[string]bool),
		observer:  obs,
	}, func() {}, nil
}

func (a *MQ) Start(ctx context.Context) error {
	return a.router.Run(ctx)
}
func (a *MQ) Stop(ctx context.Context) error {
	return a.router.Close()
}

// RegisterTopic register Topic to MQ
// If a message keep fail after retry, It will be sent to the PoisonQueue.
// PoisonQueue will retry messages in a very low rate.
func (a *MQ) RegisterTopic(topic TopicInterface) error {
	if _, exist := a.topicList[topic.Name()]; exist {
		return fmt.Errorf("topic %s already registered", topic)
	}
	a.topicList[topic.Name()] = true
	poisonQueue, err := middleware.PoisonQueue(a.pubSub.publisher, poisonedTopicName(topic.Name()))
	if err != nil {
		return err
	}
	// Normal queue
	a.router.AddNoPublisherHandler(
		topic.Name(),
		topic.Name(),
		a.pubSub.subscriber,
		func(msg *message.Message) error {
			return topic.Consume(msg.Context(), msg.Payload)
		},
	).AddMiddleware(poisonQueue, retryMiddleware())
	// Poison queue
	if topic.GetOptions() != nil && topic.GetOptions().ConsumePoisoned {
		a.router.AddNoPublisherHandler(
			poisonedTopicName(topic.Name()),
			poisonedTopicName(topic.Name()),
			a.pubSub.subscriber,
			func(msg *message.Message) error {
				return topic.Consume(msg.Context(), msg.Payload)
			},
		).AddMiddleware(poisonQueue, middleware.NewThrottle(1, time.Hour).Middleware)
	}
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
	for i := 0; err != nil && i < 5; i += 1 {
		if i > 0 {
			logger.Warnf("retry to publish message %s to topic %s: %s", msg.UUID, topic, err.Error())
			time.Sleep(time.Duration(i) * time.Second)
		}
		err = a.pubSub.publisher.Publish(topic, msg)
	}
	if err != nil {
		logger.Errorf("failed to publish message %s to topic %s: %s", msg.UUID, topic, err.Error())
	}
	return err
}

func poisonedTopicName(topic string) string {
	return topic + "_poisoned"
}

func retryMiddleware() func(h message.HandlerFunc) message.HandlerFunc {
	return middleware.Retry{
		MaxRetries:          3, //nolint:mnd //TODO
		InitialInterval:     time.Second,
		MaxInterval:         time.Minute,
		Multiplier:          4, //nolint:mnd //TODO
		MaxElapsedTime:      0,
		RandomizationFactor: 0.2, //nolint:mnd //TODO
		OnRetryHook:         nil,
		Logger:              newMQLogger(),
	}.Middleware
}
