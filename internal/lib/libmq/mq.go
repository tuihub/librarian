package libmq

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
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

func NewMQ(c *conf.MQ, app *libapp.Settings) (*MQ, func(), error) {
	loggerAdapter := newMQLogger()
	var ps *pubSub
	if c == nil {
		c = new(conf.MQ)
	}
	if c.GetDriver() == "" {
		logger.Warnf("mq driver is not set, using memory as default")
		c.Driver = "memory"
	}
	switch c.GetDriver() {
	case "memory":
		ps = newGoChannelAdapter(loggerAdapter)
	case "sql":
		var err error
		ps, err = newSQLAdapter(c.GetDatabase(), loggerAdapter)
		if err != nil {
			return nil, func() {}, err
		}
	}
	router, err := message.NewRouter(
		message.RouterConfig{CloseTimeout: 0},
		loggerAdapter,
	)
	router.AddMiddleware(middleware.CorrelationID)
	if app.EnablePanicRecovery {
		router.AddMiddleware(middleware.Recoverer)
	}
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
	for i := 0; err != nil && i < 16; i += 1 {
		err = a.pubSub.publisher.Publish(topic, msg)
	}
	return err
}

func poisonedTopicName(topic string) string {
	return topic + "_poisoned"
}

func retryMiddleware() func(h message.HandlerFunc) message.HandlerFunc {
	return middleware.Retry{
		MaxRetries:          10, //nolint:gomnd //TODO
		InitialInterval:     time.Second,
		MaxInterval:         time.Minute,
		Multiplier:          2, //nolint:gomnd //TODO
		MaxElapsedTime:      0,
		RandomizationFactor: 0.1, //nolint:gomnd //TODO
		OnRetryHook:         nil,
		Logger:              newMQLogger(),
	}.Middleware
}
