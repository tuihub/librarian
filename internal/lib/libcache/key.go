package libcache

import (
	"context"
	"errors"
	"time"

	"github.com/tuihub/librarian/internal/lib/libcodec"
)

func NewKey[T any](store StoreInterface, key string) *Key[T] {
	return &Key[T]{
		store:   store,
		keyName: key,
	}
}

type Key[T any] struct {
	store   StoreInterface
	keyName string
}

func (k *Key[T]) Get(ctx context.Context) (*T, error) {
	res := new(T)
	value, err := k.store.Get(ctx, k.keyName)
	if err != nil {
		return nil, err
	}
	switch v := value.(type) {
	case []byte:
		err = libcodec.Unmarshal(libcodec.JSON, v, res)
	case string:
		err = libcodec.Unmarshal(libcodec.JSON, []byte(v), res)
	default:
		return nil, errors.New("unexpected value type")
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (k *Key[T]) GetWithTTL(ctx context.Context) (*T, time.Duration, error) {
	res := new(T)
	value, ttl, err := k.store.GetWithTTL(ctx, k.keyName)
	if err != nil {
		return nil, 0, err
	}
	switch v := value.(type) {
	case []byte:
		err = libcodec.Unmarshal(libcodec.JSON, v, res)
	case string:
		err = libcodec.Unmarshal(libcodec.JSON, []byte(v), res)
	default:
		return nil, 0, errors.New("unexpected value type")
	}
	if err != nil {
		return nil, 0, err
	}
	return res, ttl, nil
}

func (k *Key[T]) Set(ctx context.Context, value *T, options ...Option) error {
	b, err := libcodec.Marshal(libcodec.JSON, value)
	if err != nil {
		return err
	}
	return k.store.Set(ctx, k.keyName, b, options...)
}

func (k *Key[T]) Delete(ctx context.Context) error {
	return k.store.Delete(ctx, k.keyName)
}
