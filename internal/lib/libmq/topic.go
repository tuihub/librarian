package libmq

import (
	"context"
	"errors"

	"github.com/tuihub/librarian/internal/lib/libcodec"
)

type TopicInterface interface {
	Name() string
	Consume(context.Context, []byte) error
	SetMQ(*MQ)
}

func NewTopic[T any](topic string, consumerFunc func(context.Context, *T) error) *Topic[T] {
	return &Topic[T]{
		mq:           nil,
		topicName:    topic,
		consumerFunc: consumerFunc,
	}
}

type Topic[T any] struct {
	mq           *MQ
	topicName    string
	consumerFunc func(context.Context, *T) error
}

func (t *Topic[T]) SetMQ(mq *MQ) {
	t.mq = mq
}

func (t *Topic[T]) Name() string {
	return t.topicName
}

func (t *Topic[T]) Publish(ctx context.Context, i T) error {
	if t.mq == nil {
		return errors.New("topic not registered")
	}
	p, err := libcodec.Marshal(libcodec.JSON, i)
	if err != nil {
		return err
	}
	return t.mq.Publish(ctx, t.topicName, p)
}

func (t *Topic[T]) Consume(ctx context.Context, i []byte) error {
	p := new(T)
	err := libcodec.Unmarshal(libcodec.JSON, i, p)
	if err != nil {
		return err
	}
	return t.consumerFunc(ctx, p)
}
