package libtype

import "sync"

type SyncMap[K comparable, V any] struct {
	m sync.Map
}

func NewSyncMap[K comparable, V any]() *SyncMap[K, V] {
	return &SyncMap[K, V]{
		m: sync.Map{},
	}
}

func (m *SyncMap[K, V]) Load(key K) (V, bool) {
	var value V
	v, ok := m.m.Load(key)
	if !ok {
		return value, ok
	}
	return v.(V), ok
}

func (m *SyncMap[K, V]) LoadOrStore(key K, value V) (V, bool) {
	a, loaded := m.m.LoadOrStore(key, value)
	return a.(V), loaded
}

func (m *SyncMap[K, V]) LoadAndDelete(key K) (V, bool) {
	var value V
	v, loaded := m.m.LoadAndDelete(key)
	if !loaded {
		return value, loaded
	}
	return v.(V), loaded
}

func (m *SyncMap[K, V]) Store(key K, value V) {
	m.m.Store(key, value)
}

func (m *SyncMap[K, V]) Range(f func(key K, value V) bool) {
	m.m.Range(func(key, value any) bool { return f(key.(K), value.(V)) })
}

func (m *SyncMap[K, V]) Swap(key K, value V) (V, bool) {
	p, loaded := m.m.Swap(key, value)
	return p.(V), loaded
}

func (m *SyncMap[K, V]) CompareAndSwap(key K, oldValue, newValue V) bool {
	return m.m.CompareAndSwap(key, oldValue, newValue)
}

func (m *SyncMap[K, V]) Delete(key K) {
	m.m.Delete(key)
}

func (m *SyncMap[K, V]) CompareAndDelete(key string, value V) bool {
	return m.m.CompareAndDelete(key, value)
}
