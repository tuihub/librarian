package portersdk

import (
	"context"
	"time"

	pb "github.com/tuihub/protos/pkg/librarian/porter/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/transport/grpc"
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
	} else {
		opts = append(opts, grpc.Timeout(time.Minute))
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

func (s *service) GetPorterInformation(ctx context.Context, req *pb.GetPorterInformationRequest) (
	*pb.GetPorterInformationResponse, error) {
	return s.p.GetPorterInformation(ctx, req)
}
func (s *service) EnablePorter(ctx context.Context, req *pb.EnablePorterRequest) (
	*pb.EnablePorterResponse, error) {
	return s.p.EnablePorter(ctx, req)
}
func (s *service) PullAccount(ctx context.Context, req *pb.PullAccountRequest) (
	*pb.PullAccountResponse, error) {
	if !s.p.Enabled() {
		return nil, errors.Forbidden("Unauthorized caller", "")
	}
	if req.GetAccountId() == nil ||
		req.GetAccountId().GetPlatform() == "" ||
		req.GetAccountId().GetPlatformAccountId() == "" {
		return nil, errors.BadRequest("Invalid account id", "")
	}
	for _, account := range s.p.config.FeatureSummary.GetSupportedAccounts() {
		if account.GetPlatform() == req.GetAccountId().GetPlatform() {
			return s.p.handler.PullAccount(ctx, req)
		}
	}
	return nil, errors.BadRequest("Unsupported account platform", "")
}
func (s *service) PullApp(ctx context.Context, req *pb.PullAppRequest) (*pb.PullAppResponse, error) {
	if !s.p.Enabled() {
		return nil, errors.Forbidden("Unauthorized caller", "")
	}
	if req.GetAppId() == nil ||
		req.GetAppId().GetInternal() ||
		req.GetAppId().GetSource() == "" ||
		req.GetAppId().GetSourceAppId() == "" {
		return nil, errors.BadRequest("Invalid app id", "")
	}
	for _, source := range s.p.config.FeatureSummary.GetSupportedAppSources() {
		if source == req.GetAppId().GetSource() {
			return s.p.handler.PullApp(ctx, req)
		}
	}
	return nil, errors.BadRequest("Unsupported app source", "")
}
func (s *service) PullAccountAppRelation(ctx context.Context, req *pb.PullAccountAppRelationRequest) (
	*pb.PullAccountAppRelationResponse, error) {
	if !s.p.Enabled() {
		return nil, errors.Forbidden("Unauthorized caller", "")
	}
	if req.GetAccountId() == nil ||
		req.GetRelationType() == librarian.AccountAppRelationType_ACCOUNT_APP_RELATION_TYPE_UNSPECIFIED ||
		req.GetAccountId().GetPlatform() == "" || req.GetAccountId().GetPlatformAccountId() == "" {
		return nil, errors.BadRequest("Invalid account id", "")
	}
	for _, account := range s.p.config.FeatureSummary.GetSupportedAccounts() {
		if account.GetPlatform() == req.GetAccountId().GetPlatform() {
			for _, relationType := range account.GetAppRelationTypes() {
				if relationType == req.GetRelationType() {
					return s.p.handler.PullAccountAppRelation(ctx, req)
				}
			}
			return nil, errors.BadRequest("Unsupported relation type", "")
		}
	}
	return nil, errors.BadRequest("Unsupported account", "")
}
func (s *service) PullFeed(ctx context.Context, req *pb.PullFeedRequest) (*pb.PullFeedResponse, error) {
	if !s.p.Enabled() {
		return nil, errors.Forbidden("Unauthorized caller", "")
	}
	if req.GetSource() == "" ||
		req.GetChannelId() == "" {
		return nil, errors.BadRequest("Invalid feed id", "")
	}
	for _, source := range s.p.config.FeatureSummary.GetSupportedFeedSources() {
		if source == req.GetSource() {
			return s.p.handler.PullFeed(ctx, req)
		}
	}
	return nil, errors.BadRequest("Unsupported feed source", "")
}
func (s *service) PushFeedItems(ctx context.Context, req *pb.PushFeedItemsRequest) (
	*pb.PushFeedItemsResponse, error) {
	if !s.p.Enabled() {
		return nil, errors.Forbidden("Unauthorized caller", "")
	}
	if req.GetDestination() == "" || req.GetChannelId() == "" || len(req.GetItems()) == 0 {
		return nil, errors.BadRequest("Invalid feed id", "")
	}
	for _, destination := range s.p.config.FeatureSummary.GetSupportedNotifyDestinations() {
		if destination == req.GetDestination() {
			return s.p.handler.PushFeedItems(ctx, req)
		}
	}
	return nil, errors.BadRequest("Unsupported notify destination", "")
}
