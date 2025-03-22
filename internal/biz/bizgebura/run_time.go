package bizgebura

import (
	"context"
	"time"

	"github.com/tuihub/librarian/internal/biz/bizutils"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

func (g *Gebura) AddAppRunTime(
	ctx context.Context,
	instID model.InternalID,
	timeRange *model.TimeRange,
) *errors.Error {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return bizutils.NoPermissionError()
	}
	if timeRange == nil {
		return pb.ErrorErrorReasonBadRequest("empty time range")
	}
	if timeRange.Duration <= 0 {
		return pb.ErrorErrorReasonBadRequest("invalid time range")
	}
	err := g.repo.AddAppRunTime(ctx, claims.UserID, instID, timeRange)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (g *Gebura) SumAppInstRunTime(
	ctx context.Context,
	instID model.InternalID,
	timeRange *model.TimeRange,
) (time.Duration, error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return time.Duration(0), bizutils.NoPermissionError()
	}
	if timeRange == nil {
		return time.Duration(0), pb.ErrorErrorReasonBadRequest("empty time range")
	}
	if timeRange.Duration <= 0 {
		return time.Duration(0), pb.ErrorErrorReasonBadRequest("invalid time range")
	}
	res, err := g.repo.SumAppRunTime(ctx, claims.UserID, instID, timeRange)
	if err != nil {
		return time.Duration(0), pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return res, nil
}
