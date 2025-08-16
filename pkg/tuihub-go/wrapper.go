package tuihub

import (
	"context"
	"fmt"
	"sync"
	"time"

	pb "github.com/tuihub/protos/pkg/librarian/porter/v1"
	porter "github.com/tuihub/protos/pkg/librarian/sephirah/v1/porter"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

const (
	defaultHeartbeatInterval  = time.Second * 10
	defaultHeartbeatDowngrade = time.Second * 30
	defaultHeartbeatTimeout   = time.Second * 60
	defaultRefreshToken       = time.Hour / 2
)

type serviceWrapper struct {
	pb.LibrarianPorterServiceServer

	Info         *pb.GetPorterInformationResponse
	Logger       log.Logger
	Client       porter.LibrarianSephirahPorterServiceClient
	RequireToken bool
	Token        *tokenInfo
	tokenMu      sync.Mutex

	lastHeartbeat    time.Time
	lastRefreshToken time.Time
}

type tokenInfo struct {
	enabler      int64
	AccessToken  string
	refreshToken string
}

func (s *serviceWrapper) GetPorterInformation(ctx context.Context, req *pb.GetPorterInformationRequest) (
	*pb.GetPorterInformationResponse, error) {
	return s.Info, nil
}
func (s *serviceWrapper) EnablePorter(ctx context.Context, req *pb.EnablePorterRequest) ( //nolint:gocognit //TODO
	*pb.EnablePorterResponse, error) {
	needRefreshToken := false
	f := func() error {
		s.tokenMu.Lock()
		defer s.tokenMu.Unlock()
		if s.Token != nil { //nolint:nestif //TODO
			if s.Token.enabler == req.GetSephirahId() {
				if req.GetRefreshToken() != "" {
					resp, err := s.Client.RefreshToken(
						WithToken(ctx, req.GetRefreshToken()),
						&porter.RefreshTokenRequest{},
					)
					if err != nil {
						return err
					}
					s.Token.AccessToken = resp.GetAccessToken()
					s.Token.refreshToken = resp.GetRefreshToken()
					s.lastRefreshToken = time.Now()
				}
				if s.RequireToken &&
					(s.lastRefreshToken.Add(defaultRefreshToken).Before(time.Now()) || s.Token.refreshToken == "") {
					needRefreshToken = true
				}
				return nil
			} else if s.lastHeartbeat.Add(defaultHeartbeatTimeout).After(time.Now()) {
				return fmt.Errorf("porter already enabled by %d", s.Token.enabler)
			}
		}
		s.Token = new(tokenInfo)
		s.Token.enabler = req.GetSephirahId()
		s.lastHeartbeat = time.Now()
		if s.RequireToken {
			if req.GetRefreshToken() == "" {
				needRefreshToken = true
				return nil
			}
			resp, err := s.Client.RefreshToken(
				WithToken(ctx, req.GetRefreshToken()),
				&porter.RefreshTokenRequest{},
			)
			if err != nil {
				return err
			}
			s.Token = &tokenInfo{
				enabler:      req.GetSephirahId(),
				AccessToken:  resp.GetAccessToken(),
				refreshToken: resp.GetRefreshToken(),
			}
			s.lastRefreshToken = time.Now()
		}
		return nil
	}
	if err := f(); err != nil {
		return nil, err
	}
	if resp, err := s.LibrarianPorterServiceServer.EnablePorter(ctx, req); isUnimplementedError(err) {
		return &pb.EnablePorterResponse{
			StatusMessage:    "",
			NeedRefreshToken: needRefreshToken,
			EnablesSummary:   nil,
		}, nil
	} else {
		resp.NeedRefreshToken = needRefreshToken || resp.GetNeedRefreshToken()
		return resp, err
	}
}
func (s *serviceWrapper) Enabled() bool {
	return s.Token != nil
}

func NewServer(c *ServerConfig, service pb.LibrarianPorterServiceServer, logger log.Logger) *grpc.Server {
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

type serviceServer struct {
	*serviceWrapper
}

func NewService(p *serviceWrapper) pb.LibrarianPorterServiceServer {
	return &serviceServer{
		p,
	}
}

func (s *serviceServer) GetPorterInformation(ctx context.Context, req *pb.GetPorterInformationRequest) (
	*pb.GetPorterInformationResponse, error) {
	return s.serviceWrapper.GetPorterInformation(ctx, req)
}
func (s *serviceServer) EnablePorter(ctx context.Context, req *pb.EnablePorterRequest) (
	*pb.EnablePorterResponse, error) {
	return s.serviceWrapper.EnablePorter(ctx, req)
}
func (s *serviceServer) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (
	*pb.GetAccountResponse, error) {
	if !s.serviceWrapper.Enabled() {
		return nil, errors.Forbidden("Unauthorized caller", "")
	}
	if req == nil ||
		req.GetPlatform() == "" ||
		req.GetPlatformAccountId() == "" {
		return nil, errors.BadRequest("Invalid account id", "")
	}
	for _, account := range s.serviceWrapper.Info.GetFeatureSummary().GetAccountPlatforms() {
		if account.GetId() == req.GetPlatform() {
			return s.serviceWrapper.LibrarianPorterServiceServer.GetAccount(ctx, req)
		}
	}
	return nil, errors.BadRequest("Unsupported account platform", "")
}
func (s *serviceServer) GetAppInfo(ctx context.Context, req *pb.GetAppInfoRequest) (*pb.GetAppInfoResponse, error) {
	if !s.serviceWrapper.Enabled() {
		return nil, errors.Forbidden("Unauthorized caller", "")
	}
	if req == nil ||
		req.GetSource() == "" ||
		req.GetSourceAppId() == "" {
		return nil, errors.BadRequest("Invalid app id", "")
	}
	for _, source := range s.serviceWrapper.Info.GetFeatureSummary().GetAppInfoSources() {
		if source.GetId() == req.GetSource() {
			return s.serviceWrapper.LibrarianPorterServiceServer.GetAppInfo(ctx, req)
		}
	}
	return nil, errors.BadRequest("Unsupported app source", "")
}

// func (s *serviceServer) GetAccountAppInfoRelation(ctx context.Context, req *pb.GetAccountAppInfoRelationRequest) (
//
//		*pb.GetAccountAppInfoRelationResponse, error) {
//		if !s.serviceWrapper.Enabled() {
//			return nil, errors.Forbidden("Unauthorized caller", "")
//		}
//		if req.GetAccountId() == nil ||
//			req.GetRelationType() == librarian.AccountAppRelationType_ACCOUNT_APP_RELATION_TYPE_UNSPECIFIED ||
//			req.GetAccountId().GetPlatform() == "" || req.GetAccountId().GetPlatformAccountId() == "" {
//			return nil, errors.BadRequest("Invalid account id", "")
//		}
//		for _, account := range s.serviceWrapper.Info.GetFeatureSummary().GetAccountPlatforms() {
//			if account.GetId() == req.GetAccountId().GetPlatform() {
//				return s.serviceWrapper.LibrarianPorterServiceServer.GetAccountAppInfoRelation(ctx, req)
//			}
//		}
//		return nil, errors.BadRequest("Unsupported account", "")
//	}

func (s *serviceServer) SearchAppInfo(
	ctx context.Context,
	req *pb.SearchAppInfoRequest,
) (*pb.SearchAppInfoResponse, error) {
	if !s.serviceWrapper.Enabled() {
		return nil, errors.Forbidden("Unauthorized caller", "")
	}
	if req.GetNameLike() == "" {
		return nil, errors.BadRequest("Invalid app name", "")
	}
	if len(s.serviceWrapper.Info.GetFeatureSummary().GetAppInfoSources()) > 0 {
		return s.serviceWrapper.LibrarianPorterServiceServer.SearchAppInfo(ctx, req)
	}
	return nil, errors.BadRequest("Unsupported app source", "")
}
func (s *serviceServer) PullFeed(ctx context.Context, req *pb.PullFeedRequest) (*pb.PullFeedResponse, error) {
	if !s.serviceWrapper.Enabled() {
		return nil, errors.Forbidden("Unauthorized caller", "")
	}
	for _, source := range s.serviceWrapper.Info.GetFeatureSummary().GetFeedSources() {
		if source.GetId() == req.GetSource().GetId() {
			return s.serviceWrapper.LibrarianPorterServiceServer.PullFeed(ctx, req)
		}
	}
	return nil, errors.BadRequest("Unsupported feed source", "")
}
func (s *serviceServer) PushFeedItems(ctx context.Context, req *pb.PushFeedItemsRequest) (
	*pb.PushFeedItemsResponse, error) {
	if !s.serviceWrapper.Enabled() {
		return nil, errors.Forbidden("Unauthorized caller", "")
	}
	for _, destination := range s.serviceWrapper.Info.GetFeatureSummary().GetNotifyDestinations() {
		if destination.GetId() == req.GetDestination().GetId() {
			return s.serviceWrapper.LibrarianPorterServiceServer.PushFeedItems(ctx, req)
		}
	}
	return nil, errors.BadRequest("Unsupported notify destination", "")
}
