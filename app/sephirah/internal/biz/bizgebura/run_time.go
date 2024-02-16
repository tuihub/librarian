package bizgebura

import (
	"context"
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizutils"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

func (g *Gebura) AddAppInstRunTime(
	ctx context.Context,
	packageID model.InternalID,
	timeRange *model.TimeRange,
) *errors.Error {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return bizutils.NoPermissionError()
	}
	if timeRange == nil {
		return pb.ErrorErrorReasonBadRequest("empty time range")
	}
	err := g.repo.AddAppInstRunTime(ctx, claims.UserID, packageID, timeRange)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (g *Gebura) SumAppInstRunTime(
	ctx context.Context,
	packageID model.InternalID,
	timeRange *model.TimeRange,
) (time.Duration, error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return time.Duration(0), bizutils.NoPermissionError()
	}
	if timeRange == nil {
		return time.Duration(0), pb.ErrorErrorReasonBadRequest("empty time range")
	}
	res, err := g.repo.SumAppInstRunTime(ctx, claims.UserID, packageID, timeRange)
	if err != nil {
		return time.Duration(0), pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return res, nil
}
