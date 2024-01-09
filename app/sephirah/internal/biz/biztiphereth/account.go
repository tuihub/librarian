package biztiphereth

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizutils"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/logger"
	"github.com/tuihub/librarian/model"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

func (t *Tiphereth) LinkAccount(
	ctx context.Context,
	a modeltiphereth.Account,
) (*modeltiphereth.Account, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, bizutils.NoPermissionError()
	}
	if !t.supv.CheckAccountPlatform(a.Platform) {
		return nil, bizutils.UnsupportedFeatureError()
	}
	id, err := t.searcher.NewID(ctx)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err)
	}
	a.ID = id
	if err = t.repo.LinkAccount(ctx, a, claims.UserID); err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	if err = t.pullAccount.Publish(ctx, modeltiphereth.PullAccountInfo{
		ID:                a.ID,
		Platform:          a.Platform,
		PlatformAccountID: a.PlatformAccountID,
	}); err != nil {
		logger.Errorf("Publish PullAccountInfo failed %s", err.Error())
	}
	return &a, nil
}

func (t *Tiphereth) UnLinkAccount(ctx context.Context, a modeltiphereth.Account) *errors.Error {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return bizutils.NoPermissionError()
	}
	if !t.supv.CheckAccountPlatform(a.Platform) {
		return bizutils.UnsupportedFeatureError()
	}
	if err := t.repo.UnLinkAccount(ctx, a, claims.UserID); err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (t *Tiphereth) ListLinkAccounts(
	ctx context.Context, id model.InternalID,
) ([]*modeltiphereth.Account, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, bizutils.NoPermissionError()
	}
	if id == 0 {
		id = claims.UserID
	}
	a, err := t.repo.ListLinkAccounts(ctx, id)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return a, nil
}
