package libcache

import (
	"context"
	"errors"
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
func (s *redisStore) Get(ctx context.Context, key string) (string, error) {
	object, err := s.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return "", newNotFound(err)
	}
	if err != nil {
		return "", err
	}
	return object, nil
}

// GetWithTTL returns data stored from a given key and its corresponding TTL.
func (s *redisStore) GetWithTTL(ctx context.Context, key string) (string, time.Duration, error) {
	object, err := s.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return "", 0, newNotFound(err)
	}
	if err != nil {
		return "", 0, err
	}

	ttl, err := s.client.TTL(ctx, key).Result()
	if err != nil {
		return "", 0, err
	}

	return object, ttl, err
}

// Set defines data in Redis for given key identifier.
func (s *redisStore) Set(ctx context.Context, key string, value string, options ...Option) error {
	opts := applyOptionsWithDefault(s.options, options...)

	err := s.client.Set(ctx, key, value, opts.Expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

// Delete removes data from Redis for given key identifier.
func (s *redisStore) Delete(ctx context.Context, key string) error {
	_, err := s.client.Del(ctx, key).Result()
	return err
}

// Lock acquires a distributed lock in Redis.
func (s *redisStore) Lock(ctx context.Context, key string, value string, options ...Option) error {
	opts := applyOptionsWithDefault(s.options, options...)
	success, err := s.client.SetNX(ctx, key, value, opts.Expiration).Result()
	if err != nil {
		return err
	}
	if !success {
		return errors.New("failed to acquire lock")
	}
	return nil
}

// Unlock releases a distributed lock in Redis.
func (s *redisStore) Unlock(ctx context.Context, key string, value string) (bool, error) {
	luaScript := `
    if redis.call("GET", KEYS[1]) == ARGV[1] then
        return redis.call("DEL", KEYS[1])
    else
        return 0
    end
	`
	result, err := s.client.Eval(ctx, luaScript, []string{key}, value).Result()
	if err != nil {
		return false, err
	}
	return result != 0, nil
}
