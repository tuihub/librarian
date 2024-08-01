package libtype

import "sync"

type SyncMap[T any] struct {
	m sync.Map
}

func NewSyncMap[T any]() *SyncMap[T] {
	return &SyncMap[T]{
		m: sync.Map{},
	}
}

func (m *SyncMap[T]) Load(key string) *T {
	v, ok := m.m.Load(key)
	if !ok {
		return nil
	}
	value, ok := v.(T)
	if !ok {
		return nil
	}
	return &value
}

func (m *SyncMap[T]) LoadOrStore(key string, value T) *T {
	v, loaded := m.m.LoadOrStore(key, value)
	if !loaded {
		return nil
	}
	actual, ok := v.(T)
	if !ok {
		return nil
	}
	return &actual
}

func (m *SyncMap[T]) LoadAndDelete(key string) *T {
	v, loaded := m.m.LoadAndDelete(key)
	if !loaded {
		return nil
	}
	value, ok := v.(T)
	if !ok {
		return nil
	}
	return &value
}

func (m *SyncMap[T]) Store(key string, value T) {
	m.m.Store(key, value)
}

func (m *SyncMap[T]) Range(f func(key string, value T) bool) {
	m.m.Range(func(key, value interface{}) bool {
		return f(key.(string), value.(T))
	})
}

func (m *SyncMap[T]) Swap(key string, value T) *T {
	v, loaded := m.m.Swap(key, value)
	if !loaded {
		return nil
	}
	previous, ok := v.(T)
	if !ok {
		return nil
	}
	return &previous
}

func (m *SyncMap[T]) CompareAndSwap(key string, oldValue, newValue T) bool {
	return m.m.CompareAndSwap(key, oldValue, newValue)
}

func (m *SyncMap[T]) Delete(key string) {
	m.m.Delete(key)
}

func (m *SyncMap[T]) CompareAndDelete(key string, value T) bool {
	return m.m.CompareAndDelete(key, value)
}
