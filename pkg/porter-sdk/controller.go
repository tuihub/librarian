package portersdk

import (
	"context"
	"fmt"
	"time"

	pb "github.com/tuihub/protos/pkg/librarian/porter/v1"
	sephirah "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	capi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc/metadata"
)

type controller struct {
	handler Handler
	config  PorterConfig
	logger  log.Logger
	token   *tokenInfo
	client  sephirah.LibrarianSephirahServiceClient
}

type tokenInfo struct {
	enabler      int64
	accessToken  string
	refreshToken string
}

func (s *controller) GetPorterInformation(ctx context.Context, req *pb.GetPorterInformationRequest) (
	*pb.GetPorterInformationResponse, error) {
	return &pb.GetPorterInformationResponse{
		Name:           s.config.Name,
		Version:        s.config.Version,
		GlobalName:     s.config.GlobalName,
		FeatureSummary: s.config.FeatureSummary,
	}, nil
}
func (s *controller) EnablePorter(ctx context.Context, req *pb.EnablePorterRequest) (
	*pb.EnablePorterResponse, error) {
	if s.token != nil {
		if s.token.enabler == req.GetSephirahId() {
			return &pb.EnablePorterResponse{}, nil
		} else {
			return nil, fmt.Errorf("porter already enabled by %d", s.token.enabler)
		}
	}
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+req.GetRefreshToken())
	resp, err := s.client.RefreshToken(ctx, &sephirah.RefreshTokenRequest{})
	if err != nil {
		return nil, err
	}
	s.token = &tokenInfo{
		enabler:      req.GetSephirahId(),
		accessToken:  resp.GetAccessToken(),
		refreshToken: resp.GetRefreshToken(),
	}
	return &pb.EnablePorterResponse{}, nil
}
func (s *controller) Enabled() bool {
	return s.token != nil
}

func newRegistry() (*consul.Registry, error) {
	client, err := capi.NewClient(capi.DefaultConfig())
	if err != nil {
		return nil, err
	}
	return consul.New(client), nil
}

func newSephirahClient() (sephirah.LibrarianSephirahServiceClient, error) {
	r, err := newRegistry()
	if err != nil {
		return nil, err
	}
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///sephirah"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
		grpc.WithTimeout(time.Minute),
	)
	cli := sephirah.NewLibrarianSephirahServiceClient(conn)
	return cli, err
}