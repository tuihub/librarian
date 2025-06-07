package libkratos

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	grpcstd "google.golang.org/grpc"
	grpcmd "google.golang.org/grpc/metadata"
)

func MiddlewareToGRPCUnaryInterceptor(ms ...middleware.Middleware) grpcstd.UnaryServerInterceptor {
	chain := middleware.Chain(ms...)

	return func(ctx context.Context, req interface{}, info *grpcstd.UnaryServerInfo, handler grpcstd.UnaryHandler) (interface{}, error) {
		md, _ := grpcmd.FromIncomingContext(ctx)
		ctx = transport.NewServerContext(ctx, &Transport{
			K:   transport.KindGRPC,
			E:   "inproc",
			O:   info.FullMethod,
			Req: GRPCHeaderCarrier(md),
			Res: GRPCHeaderCarrier(grpcmd.MD{}),
		})
		h := func(ctx context.Context, req interface{}) (interface{}, error) {
			return handler(ctx, req)
		}

		h = chain(h)

		return h(ctx, req)
	}
}
