package bizgebura

import (
	"context"

	"github.com/tuihub/librarian/internal/biz/bizutils"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelgebura"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

func (g *Gebura) CreateAppInst(
	ctx context.Context,
	inst *modelgebura.AppInst,
) (*modelgebura.AppInst, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, bizutils.NoPermissionError()
	}
	id, err := g.id.New()
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err)
	}
	inst.ID = id
	if err = g.repo.CreateAppInst(ctx, claims.UserID, inst); err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return inst, nil
}

func (g *Gebura) UpdateAppInst(ctx context.Context, inst *modelgebura.AppInst) *errors.Error {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return bizutils.NoPermissionError()
	}
	err := g.repo.UpdateAppInst(ctx, claims.UserID, inst)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (g *Gebura) ListAppInsts(
	ctx context.Context,
	paging model.Paging,
	ids []model.InternalID,
	appIDs []model.InternalID,
	deviceIDs []model.InternalID,
) ([]*modelgebura.AppInst, int, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, 0, bizutils.NoPermissionError()
	}
	res, total, err := g.repo.ListAppInsts(ctx, claims.UserID, paging, ids, appIDs, deviceIDs)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return res, total, nil
}
