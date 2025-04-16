package bizyesod

import (
	"context"

	"github.com/tuihub/librarian/internal/biz/bizutils"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelyesod"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

func (y *Yesod) CreateFeedActionSet(
	ctx context.Context,
	set *modelyesod.FeedActionSet,
) (model.InternalID, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return 0, bizutils.NoPermissionError()
	}
	if len(set.Actions) == 0 {
		return 0, pb.ErrorErrorReasonBadRequest("actions is empty")
	}
	id, err := y.id.New()
	if err != nil {
		return 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	set.ID = id
	err = y.repo.CreateFeedActionSet(ctx, claims.UserID, set)
	if err != nil {
		return 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return set.ID, nil
}

func (y *Yesod) UpdateFeedActionSet(ctx context.Context, set *modelyesod.FeedActionSet) *errors.Error {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return bizutils.NoPermissionError()
	}
	if len(set.Actions) == 0 {
		return pb.ErrorErrorReasonBadRequest("actions is empty")
	}
	err := y.repo.UpdateFeedActionSet(ctx, claims.UserID, set)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (y *Yesod) ListFeedActionSets(
	ctx context.Context,
	paging model.Paging,
) ([]*modelyesod.FeedActionSet, int, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, 0, bizutils.NoPermissionError()
	}
	sets, total, err := y.repo.ListFeedActionSets(ctx, claims.UserID, paging)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return sets, total, nil
}
