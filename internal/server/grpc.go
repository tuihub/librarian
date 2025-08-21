package server

import (
	"net"
	"strconv"

	"github.com/tuihub/librarian/internal/client"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	porterpb "github.com/tuihub/protos/pkg/librarian/porter/v1"
	sentinelpb "github.com/tuihub/protos/pkg/librarian/sentinel/v1"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

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
	c *conf.Server,
	auth *libauth.Auth,
	sephirahserver pb.LibrarianSephirahServiceServer,
	sentinelserver sentinelpb.LibrarianSephirahSentinelServiceServer,
	porterserver porterpb.LibrarianSephirahPorterServiceServer,
	app *libapp.Settings,
	inprocPorter *client.InprocPorter,
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
	middlewares = append(middlewares, NewTokenMatcher(auth)...)
	var opts = []grpc.ServerOption{
		grpc.Middleware(middlewares...),
	}
	opts = append(opts, grpc.Address(net.JoinHostPort(c.Main.Host, strconv.Itoa(int(c.Main.Port)))))
	if c.Main.Timeout > 0 {
		opts = append(opts, grpc.Timeout(c.Main.Timeout))
	}
	srv := grpc.NewServer(opts...)
	pb.RegisterLibrarianSephirahServiceServer(srv, sephirahserver)
	sentinelpb.RegisterLibrarianSephirahSentinelServiceServer(srv, sentinelserver)
	porterpb.RegisterLibrarianSephirahPorterServiceServer(srv, porterserver)

	inprocPorter.SetSephirahServer(porterserver, middlewares)
	return srv, nil
}
