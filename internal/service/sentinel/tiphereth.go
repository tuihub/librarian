package sentinel

import (
	"context"

	"github.com/tuihub/librarian/internal/lib/logger"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1/sentinel"
)

func (s *LibrarianSentinelService) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (
	*pb.RefreshTokenResponse, error,
) {
	accessToken, refreshToken, err := s.g.SentinelRefreshToken(ctx)
	if err != nil {
		logger.Infof("GetToken failed: %s", err.Error())
		return nil, err
	}
	return &pb.RefreshTokenResponse{
		AccessToken:  string(accessToken),
		RefreshToken: string(refreshToken),
	}, nil
}
