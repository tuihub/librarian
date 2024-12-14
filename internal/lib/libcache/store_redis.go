package libcache

import (
	"context"
	"errors"
	"fmt"
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

func strKey(key any) string {
	str, ok := key.(string)
	if ok {
		return str
	}
	return fmt.Sprintf("%v", key)
}

// Get returns data stored from a given key.
func (s *redisStore) Get(ctx context.Context, key any) (any, error) {
	object, err := s.client.Get(ctx, strKey(key)).Result()
	if errors.Is(err, redis.Nil) {
		return nil, newNotFound(err)
	}
	if err != nil {
		return nil, err
	}
	return object, nil
}

// GetWithTTL returns data stored from a given key and its corresponding TTL.
func (s *redisStore) GetWithTTL(ctx context.Context, key any) (any, time.Duration, error) {
	object, err := s.client.Get(ctx, strKey(key)).Result()
	if errors.Is(err, redis.Nil) {
		return nil, 0, newNotFound(err)
	}
	if err != nil {
		return nil, 0, err
	}

	ttl, err := s.client.TTL(ctx, strKey(key)).Result()
	if err != nil {
		return nil, 0, err
	}

	return object, ttl, err
}

// Set defines data in Redis for given key identifier.
func (s *redisStore) Set(ctx context.Context, key any, value any, options ...Option) error {
	opts := applyOptionsWithDefault(s.options, options...)

	err := s.client.Set(ctx, strKey(key), value, opts.Expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

// Delete removes data from Redis for given key identifier.
func (s *redisStore) Delete(ctx context.Context, key any) error {
	_, err := s.client.Del(ctx, strKey(key)).Result()
	return err
}

// Clear resets all data in the store.
func (s *redisStore) Clear(ctx context.Context) error {
	if err := s.client.FlushAll(ctx).Err(); err != nil {
		return err
	}

	return nil
}
