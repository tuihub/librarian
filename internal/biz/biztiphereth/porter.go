package biztiphereth

import (
	"context"

	"github.com/tuihub/librarian/internal/biz/bizutils"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelsupervisor"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
)

func (t *Tiphereth) updatePorters(ctx context.Context) error {
	if contexts, err := t.repo.GetEnabledPorterContexts(ctx); err != nil {
		logger.Errorf("get enabled porter contexts failed: %s", err.Error())
	} else {
		for _, c := range contexts {
			err = t.supv.QueuePorterContext(ctx, *c)
			if err != nil {
				logger.Errorf("queue porter context failed: %s", err.Error())
			}
		}
	}
	newPorters, err := t.supv.RefreshAliveInstances(ctx)
	if err != nil {
		logger.Errorf("refresh alive instances failed: %s", err.Error())
		return err
	}
	if len(newPorters) == 0 {
		return nil
	}
	ids, err := t.id.BatchNew(len(newPorters))
	if err != nil {
		logger.Errorf("new batch ids failed: %s", err.Error())
		return err
	}
	for i, porter := range newPorters {
		porter.ID = ids[i]
		porter.Status = model.UserStatusBlocked
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
) ([]*modelsupervisor.PorterInstance, int64, error) {
	if libauth.FromContextAssertUserType(ctx, model.UserTypeAdmin) == nil {
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
	status model.UserStatus,
) error {
	if libauth.FromContextAssertUserType(ctx, model.UserTypeAdmin) == nil {
		return bizutils.NoPermissionError()
	}
	pi, err := t.repo.UpdatePorterStatus(ctx, id, status)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	_ = t.porterInstanceCache.Delete(ctx, pi.Address)
	return nil
}

func (t *Tiphereth) CreatePorterContext(ctx context.Context, context *modelsupervisor.PorterContext) (model.InternalID, error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return 0, bizutils.NoPermissionError()
	}
	id, err := t.id.New()
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

func (t *Tiphereth) ListPorterContexts(ctx context.Context, paging model.Paging) ([]*modelsupervisor.PorterContext, int64, error) {
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
) error {
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

func (t *Tiphereth) ListPorterDigests(
	ctx context.Context,
	paging model.Paging,
	status []model.UserStatus,
) ([]*modelsupervisor.PorterDigest, int64, error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, 0, bizutils.NoPermissionError()
	}
	if claims.UserType != model.UserTypeAdmin {
		status = []model.UserStatus{model.UserStatusActive}
	}
	groups, err := t.repo.ListPorterDigests(ctx, status)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return groups, int64(len(groups)), nil
}
