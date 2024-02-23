// Based on https://github.com/go-kratos/sentry

package libsentry

import (
	"context"
	"errors"
	"net"
	"os"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	http2 "github.com/go-kratos/kratos/v2/transport/http"
)

type valuesKey struct{}

type Option func(*Options)

type Options struct {
	// Repanic configures whether Sentry should repanic after recovery, in most cases it should be set to true.
	Repanic bool
	// WaitForDelivery configures whether you want to block the request before moving forward with the response.
	WaitForDelivery bool
	// Timeout for the event delivery requests.
	Timeout time.Duration
}

func WithRepanic(repanic bool) Option {
	return func(opts *Options) {
		opts.Repanic = repanic
	}
}

func WithWaitForDelivery(waitForDelivery bool) Option {
	return func(opts *Options) {
		opts.WaitForDelivery = waitForDelivery
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(opts *Options) {
		opts.Timeout = timeout
	}
}

// Server returns a new server middleware for Sentry.
func Server(opts ...Option) middleware.Middleware { //nolint:gocognit // ignore
	options := Options{Repanic: true} //nolint: exhaustruct // default options
	for _, o := range opts {
		o(&options)
	}
	if options.Timeout == 0 {
		options.Timeout = 2 * time.Second //nolint: gomnd // default timeout
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			hub := sentry.GetHubFromContext(ctx)
			if hub == nil {
				hub = sentry.CurrentHub().Clone()
			}
			scope := hub.Scope()
			if tr, ok := transport.FromServerContext(ctx); ok {
				switch tr.Kind() {
				case transport.KindGRPC:
					gtr, ok2 := tr.(*grpc.Transport)
					if !ok2 {
						break
					}
					scope.SetContext("gRPC", map[string]interface{}{
						"endpoint":  gtr.Endpoint(),
						"operation": gtr.Operation(),
					})
					headers := make(map[string]interface{})
					for _, k := range gtr.RequestHeader().Keys() {
						headers[k] = gtr.RequestHeader().Get(k)
					}
					scope.SetContext("Headers", headers)
				case transport.KindHTTP:
					htr, ok2 := tr.(*http2.Transport)
					if !ok2 {
						break
					}
					r := htr.Request()
					scope.SetRequest(r)
				}
			}

			ctx = context.WithValue(ctx, valuesKey{}, hub)
			defer recoverWithSentry(options, hub, ctx, req)
			return handler(ctx, req)
		}
	}
}

func recoverWithSentry(opts Options, hub *sentry.Hub, ctx context.Context, req interface{}) {
	if err := recover(); err != nil {
		if !isBrokenPipeError(err) {
			eventID := hub.RecoverWithContext(
				context.WithValue(ctx, sentry.RequestContextKey, req),
				err,
			)
			if eventID != nil && opts.WaitForDelivery {
				hub.Flush(opts.Timeout)
			}
		}
		if opts.Repanic {
			panic(err)
		}
	}
}

func isBrokenPipeError(err interface{}) bool {
	if err1, ok := err.(error); ok && err1 != nil { //nolint: nestif // ignore
		var netErr *net.OpError
		if errors.As(err1, &netErr) {
			var sysErr *os.SyscallError
			if errors.As(netErr.Err, &sysErr) {
				if strings.Contains(strings.ToLower(sysErr.Error()), "broken pipe") ||
					strings.Contains(strings.ToLower(sysErr.Error()), "connection reset by peer") {
					return true
				}
			}
		}
	}
	return false
}

// GetHubFromContext retrieves attached *sentry.Hub instance from context.
func GetHubFromContext(ctx context.Context) *sentry.Hub {
	if hub, ok := ctx.Value(valuesKey{}).(*sentry.Hub); ok {
		return hub
	}
	return nil
}
