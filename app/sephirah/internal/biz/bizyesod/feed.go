package bizyesod

import (
	"context"

	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
)

func (y *Yesod) ListFeeds(
	ctx context.Context, paging model.Paging, ids []model.InternalID,
	authorIDs []model.InternalID, sources []FeedConfigSource, statuses []FeedConfigStatus,
) ([]*FeedWithConfig, error) {
	claims, exist := libauth.FromContext(ctx)
	if !exist {
		return nil, pb.ErrorErrorReasonUnauthorized("empty token")
	}
	return y.repo.ListFeeds(ctx, claims.InternalID, paging, ids, authorIDs, sources, statuses)
}
