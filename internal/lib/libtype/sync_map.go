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
	value, ok = v.(V)
	return value, ok
}

func (m *SyncMap[K, V]) LoadOrStore(key K, value V) (V, bool) {
	a, loaded := m.m.LoadOrStore(key, value)
	existing, ok := a.(V)
	if !ok {
		return existing, false
	}
	return existing, loaded
}

func (m *SyncMap[K, V]) LoadAndDelete(key K) (V, bool) {
	var value V
	v, loaded := m.m.LoadAndDelete(key)
	if !loaded {
		return value, loaded
	}
	value, ok := v.(V)
	if !ok {
		return value, false
	}
	return value, loaded
}

func (m *SyncMap[K, V]) Store(key K, value V) {
	m.m.Store(key, value)
}

func (m *SyncMap[K, V]) Range(f func(key K, value V) bool) {
	m.m.Range(func(key, value any) bool {
		k, ok := key.(K)
		v, ok2 := value.(V)
		if !ok || !ok2 {
			return true
		}
		return f(k, v)
	})
}

func (m *SyncMap[K, V]) Swap(key K, value V) (V, bool) {
	p, loaded := m.m.Swap(key, value)
	value, ok := p.(V)
	if !ok {
		return value, false
	}
	return value, loaded
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
