package libcache

import (
	"context"
	"time"
)

// Store is the interface for all available stores
// Inspired by https://github.com/eko/gocache
type Store interface {
	Get(ctx context.Context, key string) (string, error)
	GetWithTTL(ctx context.Context, key string) (string, time.Duration, error)
	Set(ctx context.Context, key string, value string, options ...Option) error
	Delete(ctx context.Context, key string) error
	Lock(ctx context.Context, key string, value string, options ...Option) error
	Unlock(ctx context.Context, key string, value string) (bool, error)
}

// Option represents a store option function.
type Option func(o *Options)

type Options struct {
	Expiration time.Duration
}

func (o *Options) IsEmpty() bool {
	return o.Expiration == 0
}

func applyOptionsWithDefault(defaultOptions *Options, opts ...Option) *Options {
	returnedOptions := new(Options)
	*returnedOptions = *defaultOptions

	for _, opt := range opts {
		opt(returnedOptions)
	}

	return returnedOptions
}

func applyOptions(opts ...Option) *Options {
	o := new(Options)

	for _, opt := range opts {
		opt(o)
	}

	return o
}

// WithExpiration allows to specify an expiration time when setting a value.
func WithExpiration(expiration time.Duration) Option {
	return func(o *Options) {
		o.Expiration = expiration
	}
}
