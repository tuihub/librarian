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

func (t *Tiphereth) updatePorters(ctx context.Context) {
	newPorters, err := t.supv.RefreshAliveInstances(ctx)
	if err != nil {
		logger.Errorf("refresh alive instances failed: %s", err.Error())
		return
	}
	if len(newPorters) == 0 {
		return
	}
	ids, err := t.searcher.NewBatchIDs(ctx, len(newPorters))
	if err != nil {
		logger.Errorf("new batch ids failed: %s", err.Error())
		return
	}
	for i, porter := range newPorters {
		porter.ID = ids[i]
	}
	err = t.repo.UpsertPorters(ctx, newPorters)
	if err != nil {
		logger.Errorf("upsert porters failed: %s", err.Error())
		return
	}
}

func (t *Tiphereth) ListPorters(
	ctx context.Context,
	paging model.Paging,
) ([]*modeltiphereth.PorterInstance, int64, *errors.Error) {
	if libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin) == nil {
		return nil, 0, bizutils.NoPermissionError()
	}
	porters, total, err := t.repo.ListPorters(ctx, paging)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return porters, total, nil
}

func (t *Tiphereth) UpdatePorterStatus(
	ctx context.Context,
	id model.InternalID,
	status modeltiphereth.PorterInstanceStatus,
) *errors.Error {
	if libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin) == nil {
		return bizutils.NoPermissionError()
	}
	err := t.repo.UpdatePorterStatus(ctx, id, status)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (t *Tiphereth) UpdatePorterPrivilege(
	ctx context.Context,
	id model.InternalID,
	privilege *modeltiphereth.PorterInstancePrivilege,
) *errors.Error {
	claims := libauth.FromContextAssertUserType(ctx, libauth.UserTypeNormal)
	if claims == nil {
		return bizutils.NoPermissionError()
	}
	err := t.repo.UpdatePorterPrivilege(ctx, claims.UserID, id, privilege)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}
