package server

import (
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	pb "github.com/tuihub/protos/pkg/librarian/porter/v1"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Porter_Server, greeter pb.LibrarianPorterServiceServer, app *libapp.Settings) *grpc.Server {
	var middlewares = []middleware.Middleware{
		logging.Server(libapp.GetLogger()),
	}
	if app.EnablePanicRecovery {
		middlewares = append(middlewares, recovery.Recovery())
	}
	var opts = []grpc.ServerOption{
		grpc.Middleware(middlewares...),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	pb.RegisterLibrarianPorterServiceServer(srv, greeter)
	return srv
}
