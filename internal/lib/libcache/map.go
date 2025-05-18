package libcache

import (
	"context"

	"github.com/tuihub/librarian/internal/lib/libcodec"
)

type mapKeyNameFunc[K any] func(K) string
type mapFallBackFunc[K any, V any] func(context.Context, K) (*V, error)

func NewMap[K any, V any](store Store, baseKeyName string, keyNameFunc mapKeyNameFunc[K],
	defaultFallBackFunc mapFallBackFunc[K, V], defaultOptions ...Option) *Map[K, V] {
	if defaultOptions == nil {
		defaultOptions = []Option{}
	}
	return &Map[K, V]{
		store,
		baseKeyName,
		defaultOptions,
		keyNameFunc,
		defaultFallBackFunc,
	}
}

type Map[K any, V any] struct {
	store               Store
	baseKeyName         string
	defaultOptions      []Option
	keyNameFunc         mapKeyNameFunc[K]
	defaultFallBackFunc mapFallBackFunc[K, V]
}

func (m *Map[K, V]) combineKeyName(key K) string {
	return m.baseKeyName + m.keyNameFunc(key)
}

func (m *Map[K, V]) get(ctx context.Context, key K) (*V, error) {
	res := new(V)
	value, err := m.store.Get(ctx, m.combineKeyName(key))
	if err != nil {
		return nil, err
	}
	err = libcodec.Unmarshal(libcodec.JSON, []byte(value), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *Map[K, V]) Get(ctx context.Context, key K) (*V, error) {
	res, err := m.get(ctx, key)
	if err == nil {
		return res, nil
	}
	if m.defaultFallBackFunc != nil {
		res, err = m.defaultFallBackFunc(ctx, key)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}
	_ = m.Set(ctx, key, res)
	return res, nil
}

// func (m *Map[K, V]) GetWithTTL(ctx context.Context, key K) (*V, time.Duration, error) {
//	res := new(V)
//	value, ttl, err := m.store.GetWithTTL(ctx, m.combineKeyName(key))
//	if err != nil {
//		return nil, 0, err
//	}
//	switch v := value.(type) {
//	case []byte:
//		err = libcodec.Unmarshal(libcodec.JSON, v, res)
//	case string:
//		err = libcodec.Unmarshal(libcodec.JSON, []byte(v), res)
//	default:
//		return nil, 0, errors.New("unexpected value type")
//	}
//	if err != nil {
//		return nil, 0, err
//	}
//	return res, ttl, nil
//}
//
// func (m *Map[K, V]) GetWithFallBack(ctx context.Context, key K,
//	fallBackFunc mapFallBackFunc[K, V], options ...Option) (*V, error) {
//	res, err := m.get(ctx, key)
//	if err == nil {
//		return res, nil
//	}
//	if fallBackFunc != nil {
//		res, err = fallBackFunc(ctx, key)
//		if err != nil {
//			return nil, err
//		}
//	} else if m.defaultFallBackFunc != nil {
//		res, err = m.defaultFallBackFunc(ctx, key)
//		if err != nil {
//			return nil, err
//		}
//	} else {
//		return nil, err
//	}
//	_ = m.Set(ctx, key, res, options...)
//	return res, nil
//}

func (m *Map[K, V]) Set(ctx context.Context, key K, value *V, options ...Option) error {
	b, err := libcodec.Marshal(libcodec.JSON, value)
	if err != nil {
		return err
	}
	return m.store.Set(ctx, m.combineKeyName(key), string(b), append(m.defaultOptions, options...)...)
}

func (m *Map[K, V]) Delete(ctx context.Context, key K) error {
	return m.store.Delete(ctx, m.combineKeyName(key))
}
