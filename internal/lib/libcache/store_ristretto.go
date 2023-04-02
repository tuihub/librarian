package libcache

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dgraph-io/ristretto"
)

// ristrettoStore is a store for Ristretto (memory) library.
type ristrettoStore struct {
	client  *ristretto.Cache
	options *Options
}

// newRistretto creates a new store to Ristretto (memory) library instance.
func newRistretto(client *ristretto.Cache, options ...Option) *ristrettoStore {
	return &ristrettoStore{
		client:  client,
		options: applyOptions(options...),
	}
}

// Get returns data stored from a given key.
func (s *ristrettoStore) Get(_ context.Context, key any) (any, error) {
	var err error

	value, exists := s.client.Get(key)
	if !exists {
		err = errors.New("value not found in Ristretto store")
	}

	return value, err
}

// GetWithTTL returns data stored from a given key and its corresponding TTL.
func (s *ristrettoStore) GetWithTTL(ctx context.Context, key any) (any, time.Duration, error) {
	value, err := s.Get(ctx, key)
	return value, 0, err
}

// Set defines data in Ristretto memoey cache for given key identifier.
func (s *ristrettoStore) Set(_ context.Context, key any, value any, options ...Option) error {
	opts := applyOptionsWithDefault(s.options, options...)

	var err error

	if set := s.client.SetWithTTL(key, value, opts.Cost, opts.Expiration); !set {
		err = fmt.Errorf("an error has occurred while setting value '%v' on key '%v'", value, key)
	}

	if err != nil {
		return err
	}

	return nil
}

// Delete removes data in Ristretto memoey cache for given key identifier.
func (s *ristrettoStore) Delete(_ context.Context, key any) error {
	s.client.Del(key)
	return nil
}

// Clear resets all data in the store.
func (s *ristrettoStore) Clear(_ context.Context) error {
	s.client.Clear()
	return nil
}
