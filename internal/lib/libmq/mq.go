package libmq

import (
	"errors"
	"fmt"

	"github.com/tuihub/librarian/internal/lib/logger"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewMQ)

type MQ struct {
	router    *message.Router
	pubSub    *gochannel.GoChannel
	topicList map[string]bool
}

func NewMQ() (*MQ, func(), error) {
	loggerAdapter := watermill.NewStdLoggerWithOut(logger.NewWriter(), false, false)
	router, err := message.NewRouter(message.RouterConfig{}, loggerAdapter)
	pubSub := gochannel.NewGoChannel(
		gochannel.Config{},
		loggerAdapter,
	)
	cleanup := func() {
		_ = pubSub.Close()
		_ = router.Close()
	}
	return &MQ{
		router: router,
		pubSub: pubSub,
	}, cleanup, err
}

func (a *MQ) RegisterTopic(topic Topic) error {
	if _, exist := a.topicList[topic.Name()]; exist {
		return fmt.Errorf("topic %s already registered", topic)
	}
	a.topicList[topic.Name()] = true
	a.router.AddNoPublisherHandler(
		topic.Name(),
		topic.Name(),
		a.pubSub,
		func(msg *message.Message) error {
			err := topic.Consume(msg.Payload)
			if err != nil {
				return err
			}
			return nil
		},
	)
	topic.SetMQ(a)
	return nil
}

func (a *MQ) Publish(topic string, payload []byte) error {
	_, exist := a.topicList[topic]
	if !exist {
		return errors.New("unregistered topic")
	}
	msg := message.NewMessage(watermill.NewUUID(), payload)
	return a.pubSub.Publish(topic, msg)
}
