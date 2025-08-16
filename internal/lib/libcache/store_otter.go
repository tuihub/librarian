package libcache

import (
	"context"
	"errors"
	"math"
	"time"

	"github.com/tuihub/librarian/internal/lib/libtype"

	"github.com/maypok86/otter/v2"
)

// otterStore is a store for Otter.
type otterStore struct {
	client  *otter.Cache[string, string]
	locks   *libtype.SyncMap[string, *lockInfo]
	options *Options
}

// newOtter creates a new store for Otter.
func newOtter(capacity int, options ...Option) (*otterStore, error) {
	cache, err := otter.New(&otter.Options[string, string]{ //nolint:exhaustruct // no need
		MaximumSize: capacity,
	})
	if err != nil {
		return nil, err
	}
	return &otterStore{
		client:  cache,
		locks:   libtype.NewSyncMap[string, *lockInfo](),
		options: applyOptions(options...),
	}, nil
}

type lockInfo struct {
	value  string
	expire time.Time
}

func (l *lockInfo) Equal(other *lockInfo) bool {
	if l == nil || other == nil {
		return l == other
	}
	return l.value == other.value && l.expire.Equal(other.expire)
}

// Get retrieves data from the cache using the given key.
func (s *otterStore) Get(ctx context.Context, key string) (string, error) {
	value, ok := s.client.GetEntry(key)
	if !ok {
		return "", newNotFound(errors.New("key not found"))
	}
	return value.Value, nil
}

// GetWithTTL retrieves data and its TTL from the cache using the given key.
func (s *otterStore) GetWithTTL(ctx context.Context, key string) (string, time.Duration, error) {
	entry, ok := s.client.GetEntry(key)
	if !ok {
		return "", 0, newNotFound(errors.New("key not found"))
	}
	ttl := entry.ExpiresAfter()
	return entry.Value, ttl, nil
}

// Set stores data in the cache with the given key and options.
func (s *otterStore) Set(ctx context.Context, key string, value string, options ...Option) error {
	opts := applyOptionsWithDefault(s.options, options...)
	ttl := opts.Expiration
	if ttl <= 0 {
		ttl = time.Duration(math.MaxInt64)
	}
	s.client.Set(key, value)
	s.client.SetExpiresAfter(key, ttl)
	return nil
}

// Delete removes data from the cache using the given key.
func (s *otterStore) Delete(ctx context.Context, key string) error {
	s.client.Invalidate(key)
	return nil
}

// Lock acquires a lock for a given key.
func (s *otterStore) Lock(ctx context.Context, key string, value string, options ...Option) error {
	opts := applyOptionsWithDefault(s.options, options...)
	ttl := opts.Expiration
	if ttl <= 0 {
		ttl = time.Duration(math.MaxInt64)
	}
	t := time.Now()
	info, loaded := s.locks.LoadOrStore(key, &lockInfo{
		value:  value,
		expire: t.Add(ttl),
	})
	if loaded {
		if info.expire.Before(t) {
			swapped := s.locks.CompareAndSwap(key, info, &lockInfo{
				value:  value,
				expire: t.Add(ttl),
			})
			if swapped {
				return nil
			}
		}
		return errors.New("lock already acquired")
	}
	return nil
}

// Unlock releases a lock for a given key.
func (s *otterStore) Unlock(ctx context.Context, key string, value string) (bool, error) {
	info, _ := s.locks.Load(key)
	if info == nil || info.value != value {
		return false, nil
	}
	deleted := s.locks.CompareAndDelete(key, info)
	return deleted, nil
}
