package porter

import (
	"context"

	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/service/sephirah/converter"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	pb "github.com/tuihub/protos/pkg/librarian/v1"
)

func (s *LibrarianSephirahPorterService) RefreshToken(ctx context.Context, req *porter.RefreshTokenRequest) (
	*porter.RefreshTokenResponse, error,
) {
	accessToken, refreshToken, err := s.t.RefreshToken(ctx, nil)
	if err != nil {
		logger.Infof("GetToken failed: %s", err.Error())
		return nil, err
	}
	return &porter.RefreshTokenResponse{
		AccessToken:  string(accessToken),
		RefreshToken: string(refreshToken),
	}, nil
}

func (s *LibrarianSephirahPorterService) AcquireUserToken(ctx context.Context, req *porter.AcquireUserTokenRequest) (
	*porter.AcquireUserTokenResponse, error,
) {
	if req.GetUserId() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	token, err := s.t.AcquireUserToken(ctx, converter.ToBizInternalID(req.GetUserId()))
	if err != nil {
		return nil, err
	}
	return &porter.AcquireUserTokenResponse{
		AccessToken: string(token),
	}, nil
}
