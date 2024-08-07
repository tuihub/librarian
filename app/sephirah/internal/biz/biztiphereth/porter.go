package biztiphereth

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizutils"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelsupervisor"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

func (t *Tiphereth) updatePorters(ctx context.Context) error {
	newPorters, err := t.supv.RefreshAliveInstances(ctx)
	if err != nil {
		logger.Errorf("refresh alive instances failed: %s", err.Error())
		return err
	}
	if len(newPorters) == 0 {
		return nil
	}
	ids, err := t.searcher.NewBatchIDs(ctx, len(newPorters))
	if err != nil {
		logger.Errorf("new batch ids failed: %s", err.Error())
		return err
	}
	for i, porter := range newPorters {
		porter.ID = ids[i]
		porter.Status = modeltiphereth.UserStatusBlocked
	}
	err = t.repo.UpsertPorters(ctx, newPorters)
	if err != nil {
		logger.Errorf("upsert porters failed: %s", err.Error())
		return err
	}
	return nil
}

func (t *Tiphereth) ListPorters(
	ctx context.Context,
	paging model.Paging,
) ([]*modelsupervisor.PorterInstance, int64, *errors.Error) {
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
	status modeltiphereth.UserStatus,
) *errors.Error {
	if libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin) == nil {
		return bizutils.NoPermissionError()
	}
	pi, err := t.repo.UpdatePorterStatus(ctx, id, status)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	_ = t.porterInstanceCache.Delete(ctx, pi.Address)
	return nil
}

func (t *Tiphereth) CreatePorterContext(ctx context.Context, context *modelsupervisor.PorterContext) (model.InternalID, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return 0, bizutils.NoPermissionError()
	}
	id, err := t.searcher.NewID(ctx)
	if err != nil {
		return 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	context.ID = id
	err = t.repo.CreatePorterContext(ctx, claims.UserID, context)
	if err != nil {
		return 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return id, nil
}

func (t *Tiphereth) ListPorterContexts(ctx context.Context, paging model.Paging) ([]*modelsupervisor.PorterContext, int64, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, 0, bizutils.NoPermissionError()
	}
	contexts, total, err := t.repo.ListPorterContexts(ctx, claims.UserID, paging)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return contexts, total, nil
}

func (t *Tiphereth) UpdatePorterContext(
	ctx context.Context,
	context *modelsupervisor.PorterContext,
) *errors.Error {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return bizutils.NoPermissionError()
	}
	err := t.repo.UpdatePorterContext(ctx, claims.UserID, context)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (t *Tiphereth) ListPorterGroups(
	ctx context.Context,
	paging model.Paging,
	status []modeltiphereth.UserStatus,
) ([]*modelsupervisor.PorterGroup, int64, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, 0, bizutils.NoPermissionError()
	}
	if claims.UserType != libauth.UserTypeAdmin {
		status = []modeltiphereth.UserStatus{modeltiphereth.UserStatusActive}
	}
	groups, err := t.repo.ListPorterGroups(ctx, status)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return groups, int64(len(groups)), nil
}
