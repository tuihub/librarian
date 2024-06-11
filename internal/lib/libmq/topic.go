package libmq

import (
	"context"
	"errors"

	"github.com/tuihub/librarian/internal/lib/libcodec"
	"github.com/tuihub/librarian/internal/lib/libobserve"
)

type TopicInterface interface {
	Name() string
	Consume(context.Context, []byte) error
	SetMQ(*MQ)
	GetOptions() *Options
}

func NewTopic[T any](topic string, consumerFunc func(context.Context, *T) error, opts ...Option) *Topic[T] {
	return &Topic[T]{
		mq:           nil,
		topicName:    topic,
		consumerFunc: consumerFunc,
		options:      applyOptions(opts...),
		pubObserver:  nil,
		subObserver:  nil,
	}
}

type Topic[T any] struct {
	mq           *MQ
	topicName    string
	consumerFunc func(context.Context, *T) error
	options      *Options
	pubObserver  *libobserve.ObserverCounter
	subObserver  *libobserve.ObserverCounter
}

func (t *Topic[T]) SetMQ(mq *MQ) {
	t.mq = mq
	if t.mq.observer != nil {
		t.pubObserver = t.mq.observer.NewMQ(t.Name() + "_pub")
		t.subObserver = t.mq.observer.NewMQ(t.Name() + "_sub")
	}
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
	err = t.mq.Publish(ctx, t.topicName, p)

	if t.pubObserver != nil && t.subObserver != nil {
		if err != nil {
			t.pubObserver.Failure()
		} else {
			t.pubObserver.Success()
		}
	}

	return err
}

func (t *Topic[T]) LocalCall(ctx context.Context, i T) error {
	err := t.consumerFunc(ctx, &i)

	if t.pubObserver != nil && t.subObserver != nil {
		t.pubObserver.Success()
		if err != nil {
			t.subObserver.Failure()
		} else {
			t.subObserver.Success()
		}
	}

	return err
}

func (t *Topic[T]) Consume(ctx context.Context, i []byte) error {
	p := new(T)
	err := libcodec.Unmarshal(libcodec.JSON, i, p)
	if err != nil {
		return err
	}
	err = t.consumerFunc(ctx, p)

	if t.pubObserver != nil && t.subObserver != nil {
		if err != nil {
			t.subObserver.Failure()
		} else {
			t.subObserver.Success()
		}
	}

	return err
}

func (t *Topic[T]) GetOptions() *Options {
	return t.options
}
