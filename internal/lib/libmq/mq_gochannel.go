package libmq

import (
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

func newGoChannelAdapter(loggerAdapter *mqLogger) *pubSub {
	ps := gochannel.NewGoChannel( // TODO https://github.com/ThreeDotsLabs/watermill/issues/296
		gochannel.Config{
			OutputChannelBuffer:            0,
			Persistent:                     false,
			BlockPublishUntilSubscriberAck: false,
			PreserveContext:                false,
		},
		loggerAdapter,
	)
	return &pubSub{ps, ps}
}
