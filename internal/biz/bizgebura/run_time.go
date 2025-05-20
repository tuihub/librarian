package bizgebura

import (
	"context"
	"time"

	"github.com/tuihub/librarian/internal/biz/bizutils"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelgebura"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

func (g *Gebura) BatchCreateAppRunTime(
	ctx context.Context,
	runTimes []*modelgebura.AppRunTime,
) *errors.Error {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return bizutils.NoPermissionError()
	}
	for _, runTime := range runTimes {
		if runTime == nil || runTime.RunTime == nil {
			return pb.ErrorErrorReasonBadRequest("empty time range")
		}
		if runTime.RunTime.Duration <= 0 {
			return pb.ErrorErrorReasonBadRequest("invalid time range")
		}
	}
	ids, err := g.id.BatchNew(len(runTimes))
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err)
	}
	for i, runTime := range runTimes {
		runTime.ID = ids[i]
	}
	err = g.repo.BatchCreateAppRunTime(ctx, claims.UserID, runTimes)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (g *Gebura) SumAppRunTime(
	ctx context.Context,
	appIDs []model.InternalID,
	deviceIDs []model.InternalID,
	timeRange *model.TimeRange,
) (*time.Duration, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, bizutils.NoPermissionError()
	}
	if timeRange == nil {
		return nil, pb.ErrorErrorReasonBadRequest("empty time range")
	}
	if timeRange.Duration <= 0 {
		return nil, pb.ErrorErrorReasonBadRequest("invalid time range")
	}
	res, err := g.repo.SumAppRunTime(ctx, claims.UserID, appIDs, deviceIDs, timeRange)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return &res, nil
}

func (g *Gebura) ListAppRunTimes(
	ctx context.Context,
	paging model.Paging,
	appIDs []model.InternalID,
	deviceIDs []model.InternalID,
	timeRange *model.TimeRange,
) ([]*modelgebura.AppRunTime, int, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, 0, bizutils.NoPermissionError()
	}
	if timeRange == nil {
		return nil, 0, pb.ErrorErrorReasonBadRequest("empty time range")
	}
	if timeRange.Duration <= 0 {
		return nil, 0, pb.ErrorErrorReasonBadRequest("invalid time range")
	}
	res, total, err := g.repo.ListAppRunTimes(ctx, claims.UserID, paging, appIDs, deviceIDs, timeRange)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return res, total, nil
}

func (g *Gebura) DeleteAppRunTime(ctx context.Context, id model.InternalID) *errors.Error {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return bizutils.NoPermissionError()
	}
	err := g.repo.DeleteAppRunTime(ctx, claims.UserID, id)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}
