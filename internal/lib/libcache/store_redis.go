package libcache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// redisStore is a store for Redis.
type redisStore struct {
	client  *redis.Client
	options *Options
}

// newRedis creates a new store to Redis instance(s).
func newRedis(client *redis.Client, options ...Option) *redisStore {
	return &redisStore{
		client:  client,
		options: applyOptions(options...),
	}
}

// Get returns data stored from a given key.
func (s *redisStore) Get(ctx context.Context, key any) (any, error) {
	object, err := s.client.Get(ctx, key.(string)).Result()
	if err != nil {
		return nil, err
	}
	return object, nil
}

// GetWithTTL returns data stored from a given key and its corresponding TTL.
func (s *redisStore) GetWithTTL(ctx context.Context, key any) (any, time.Duration, error) {
	object, err := s.client.Get(ctx, key.(string)).Result()
	if err != nil {
		return nil, 0, err
	}

	ttl, err := s.client.TTL(ctx, key.(string)).Result()
	if err != nil {
		return nil, 0, err
	}

	return object, ttl, err
}

// Set defines data in Redis for given key identifier.
func (s *redisStore) Set(ctx context.Context, key any, value any, options ...Option) error {
	opts := applyOptionsWithDefault(s.options, options...)

	err := s.client.Set(ctx, key.(string), value, opts.Expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

// Delete removes data from Redis for given key identifier.
func (s *redisStore) Delete(ctx context.Context, key any) error {
	_, err := s.client.Del(ctx, key.(string)).Result()
	return err
}

// Clear resets all data in the store.
func (s *redisStore) Clear(ctx context.Context) error {
	if err := s.client.FlushAll(ctx).Err(); err != nil {
		return err
	}

	return nil
}
