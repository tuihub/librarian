package libmq

// Option represents a store option function.
type Option func(o *Options)

type Options struct {
	ConsumePoisoned bool
}

func applyOptions(opts ...Option) *Options {
	o := new(Options)

	for _, opt := range opts {
		opt(o)
	}

	return o
}

func WithConsumePoisoned() Option {
	return func(o *Options) {
		o.ConsumePoisoned = true
	}
}
