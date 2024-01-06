package portersdk

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	pb "github.com/tuihub/protos/pkg/librarian/porter/v1"
)

func newServer(c *ServerConfig, service pb.LibrarianPorterServiceServer, logger log.Logger) *grpc.Server {
	var middlewares = []middleware.Middleware{
		logging.Server(logger),
	}
	var opts = []grpc.ServerOption{
		grpc.Middleware(middlewares...),
	}
	if c.Network != "" {
		opts = append(opts, grpc.Network(c.Network))
	}
	if c.Addr != "" {
		opts = append(opts, grpc.Address(c.Addr))
	}
	if c.Timeout != nil {
		opts = append(opts, grpc.Timeout(*c.Timeout))
	}
	srv := grpc.NewServer(opts...)
	pb.RegisterLibrarianPorterServiceServer(srv, service)
	return srv
}

type service struct {
	pb.UnimplementedLibrarianPorterServiceServer
	p controller
}

func newService(p controller) pb.LibrarianPorterServiceServer {
	return &service{
		UnimplementedLibrarianPorterServiceServer: pb.UnimplementedLibrarianPorterServiceServer{},
		p: p,
	}
}

func (s *service) GetPorterInformation(ctx context.Context, req *pb.GetPorterInformationRequest) (*pb.GetPorterInformationResponse, error) {
	return s.p.GetPorterInformation(ctx, req)
}
func (s *service) EnablePorter(ctx context.Context, req *pb.EnablePorterRequest) (*pb.EnablePorterResponse, error) {
	return s.p.EnablePorter(ctx, req)
}
func (s *service) PullAccount(ctx context.Context, req *pb.PullAccountRequest) (*pb.PullAccountResponse, error) {
	if s.p.Enabled() {
		return s.p.handler.PullAccount(ctx, req)
	}
	return nil, errors.Forbidden("Unauthorized caller", "")
}
func (s *service) PullApp(ctx context.Context, req *pb.PullAppRequest) (*pb.PullAppResponse, error) {
	if s.p.Enabled() {
		return s.p.handler.PullApp(ctx, req)
	}
	return nil, errors.Forbidden("Unauthorized caller", "")
}
func (s *service) PullAccountAppRelation(ctx context.Context, req *pb.PullAccountAppRelationRequest) (*pb.PullAccountAppRelationResponse, error) {
	if s.p.Enabled() {
		return s.p.handler.PullAccountAppRelation(ctx, req)
	}
	return nil, errors.Forbidden("Unauthorized caller", "")
}
func (s *service) PullFeed(ctx context.Context, req *pb.PullFeedRequest) (*pb.PullFeedResponse, error) {
	if s.p.Enabled() {
		return s.p.handler.PullFeed(ctx, req)
	}
	return nil, errors.Forbidden("Unauthorized caller", "")
}
func (s *service) PushFeedItems(ctx context.Context, req *pb.PushFeedItemsRequest) (*pb.PushFeedItemsResponse, error) {
	if s.p.Enabled() {
		return s.p.handler.PushFeedItems(ctx, req)
	}
	return nil, errors.Forbidden("Unauthorized caller", "")
}
