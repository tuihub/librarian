package portersdk

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	capi "github.com/hashicorp/consul/api"
	pb "github.com/tuihub/protos/pkg/librarian/porter/v1"
	sephirah "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	"google.golang.org/grpc/metadata"
)

type controller struct {
	handler Handler
	config  *PorterConfig
	logger  log.Logger
	token   *tokenInfo
	client  sephirah.LibrarianSephirahServiceClient
}

type tokenInfo struct {
	enabler      int64
	accessToken  string
	refreshToken string
}

func (s *controller) GetPorterInformation(ctx context.Context, req *pb.GetPorterInformationRequest) (*pb.GetPorterInformationResponse, error) {
	return &pb.GetPorterInformationResponse{
		Name:           s.config.Name,
		Version:        s.config.Version,
		GlobalName:     s.config.GlobalName,
		FeatureSummary: s.config.FeatureSummary,
	}, nil
}
func (s *controller) EnablePorter(ctx context.Context, req *pb.EnablePorterRequest) (*pb.EnablePorterResponse, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+req.RefreshToken)
	resp, err := s.client.RefreshToken(ctx, &sephirah.RefreshTokenRequest{})
	if err != nil {
		return nil, err
	}
	s.token = &tokenInfo{
		enabler:      req.SephirahId,
		accessToken:  resp.AccessToken,
		refreshToken: resp.RefreshToken,
	}
	return &pb.EnablePorterResponse{}, nil
}
func (s *controller) Enabled() bool {
	return s.token != nil
}

func newSephirahClient() (sephirah.LibrarianSephirahServiceClient, error) {
	client, err := capi.NewClient(capi.DefaultConfig())
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///searcher"),
		grpc.WithDiscovery(consul.New(client)),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	cli := sephirah.NewLibrarianSephirahServiceClient(conn)
	return cli, err
}
