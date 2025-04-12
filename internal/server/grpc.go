package server

import (
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libobserve"
	"github.com/tuihub/librarian/internal/lib/libsentry"
	sentinelpb "github.com/tuihub/protos/pkg/librarian/sephirah/v1/sentinel"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1/sephirah"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"go.opentelemetry.io/otel"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	c *conf.SephirahServer,
	auth *libauth.Auth,
	greeter pb.LibrarianSephirahServiceServer,
	sentinelserver sentinelpb.LibrarianSentinelServiceServer,
	app *libapp.Settings,
	observer *libobserve.BuiltInObserver,
) (*grpc.Server, error) {
	validator, err := libapp.NewValidator()
	if err != nil {
		return nil, err
	}
	secondsHistogram, err := metrics.DefaultSecondsHistogram(
		otel.Meter(app.Name),
		metrics.DefaultServerSecondsHistogramName,
	)
	if err != nil {
		return nil, err
	}
	requestsHistogram, err := metrics.DefaultRequestsCounter(
		otel.Meter(app.Name),
		metrics.DefaultServerRequestsCounterName,
	)
	if err != nil {
		return nil, err
	}
	var middlewares = []middleware.Middleware{
		logging.Server(libapp.GetLogger()),
		ratelimit.Server(),
		tracing.Server(),
		metrics.Server(
			metrics.WithSeconds(secondsHistogram),
			metrics.WithRequests(requestsHistogram),
		),
		validator,
	}
	if app.EnablePanicRecovery {
		middlewares = append(middlewares, recovery.Recovery())
	}
	middlewares = append(middlewares, libsentry.Server())
	middlewares = append(middlewares, libobserve.Server(observer))
	middlewares = append(middlewares, NewTokenMatcher(auth)...)
	var opts = []grpc.ServerOption{
		grpc.Middleware(middlewares...),
	}
	if c.GetGrpc().GetNetwork() != "" {
		opts = append(opts, grpc.Network(c.GetGrpc().GetNetwork()))
	}
	if c.GetGrpc().GetAddr() != "" {
		opts = append(opts, grpc.Address(c.GetGrpc().GetAddr()))
	}
	if c.GetGrpc().GetTimeout() != nil {
		opts = append(opts, grpc.Timeout(c.GetGrpc().GetTimeout().AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	pb.RegisterLibrarianSephirahServiceServer(srv, greeter)
	sentinelpb.RegisterLibrarianSentinelServiceServer(srv, sentinelserver)
	return srv, nil
}
