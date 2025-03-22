package porter

import (
	"context"

	"github.com/tuihub/librarian/internal/service/sephirah/converter"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	porter "github.com/tuihub/protos/pkg/librarian/sephirah/v1/porter"
)

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
