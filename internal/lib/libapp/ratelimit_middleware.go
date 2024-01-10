package libapp

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	"go.uber.org/ratelimit"
)

func NewLeakyBucketMiddleware(rate int) middleware.Middleware {
	rl := ratelimit.New(rate)
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			rl.Take()
			return handler(ctx, req)
		}
	}
}
