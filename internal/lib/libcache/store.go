package libcache

import (
	"context"
	"time"
)

// Store is the interface for all available stores
// Inspired by https://github.com/eko/gocache
type Store interface {
	Get(ctx context.Context, key any) (any, error)
	GetWithTTL(ctx context.Context, key any) (any, time.Duration, error)
	Set(ctx context.Context, key any, value any, options ...Option) error
	Delete(ctx context.Context, key any) error
	Clear(ctx context.Context) error
}

// Option represents a store option function.
type Option func(o *Options)

type Options struct {
	Cost       int64
	Expiration time.Duration
}

func (o *Options) IsEmpty() bool {
	return o.Cost == 0 && o.Expiration == 0
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

// WithCost allows setting the memory capacity used by the item when setting a value.
// Actually it seems to be used by Ristretto library only.
func WithCost(cost int64) Option {
	return func(o *Options) {
		o.Cost = cost
	}
}

// WithExpiration allows to specify an expiration time when setting a value.
func WithExpiration(expiration time.Duration) Option {
	return func(o *Options) {
		o.Expiration = expiration
	}
}
