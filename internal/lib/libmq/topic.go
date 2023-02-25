package libmq

import (
	"context"
	"encoding/json"
	"errors"
)

type Topic interface {
	Name() string
	Consume(context.Context, []byte) error
	SetMQ(*MQ)
}

func NewTopic[T any](topic string, payloadFunc func() T, consumerFunc func(context.Context, T) error) *TopicImpl[T] {
	return &TopicImpl[T]{
		mq:           nil,
		topicName:    topic,
		payloadFunc:  payloadFunc,
		consumerFunc: consumerFunc,
	}
}

type TopicImpl[T any] struct {
	mq           *MQ
	topicName    string
	payloadFunc  func() T
	consumerFunc func(context.Context, T) error
}

func (t *TopicImpl[T]) SetMQ(mq *MQ) {
	t.mq = mq
}

func (t *TopicImpl[T]) Name() string {
	return t.topicName
}

func (t *TopicImpl[T]) Payload() T {
	return t.payloadFunc()
}

func (t *TopicImpl[T]) Publish(ctx context.Context, i T) error {
	if t.mq == nil {
		return errors.New("topic not registered")
	}
	p, err := json.Marshal(i)
	if err != nil {
		return err
	}
	return t.mq.Publish(ctx, t.topicName, p)
}

func (t *TopicImpl[T]) Consume(ctx context.Context, i []byte) error {
	p := t.Payload()
	err := json.Unmarshal(i, &p)
	if err != nil {
		return err
	}
	return t.consumerFunc(ctx, p)
}
