package libobserve

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
)

func Server(observer *BuiltInObserver) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			reply, err := handler(ctx, req)
			if err != nil {
				observer.GetRequest().Failure()
			} else {
				observer.GetRequest().Success()
			}
			return reply, err
		}
	}
}
