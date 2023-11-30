package bizgebura

import (
	"context"
	"time"

	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

func (g *Gebura) AddAppPackageRunTime(
	ctx context.Context,
	packageID model.InternalID,
	timeRange *model.TimeRange,
) *errors.Error {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return pb.ErrorErrorReasonForbidden("no permission")
	}
	if timeRange == nil {
		return pb.ErrorErrorReasonBadRequest("empty time range")
	}
	claims, exist := libauth.FromContext(ctx)
	if !exist {
		return pb.ErrorErrorReasonUnauthorized("empty token")
	}
	err := g.repo.AddAppPackageRunTime(ctx, claims.InternalID, packageID, timeRange)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (g *Gebura) SumAppPackageRunTime(
	ctx context.Context,
	packageID model.InternalID,
	timeRange *model.TimeRange,
) (time.Duration, error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return time.Duration(0), pb.ErrorErrorReasonForbidden("no permission")
	}
	if timeRange == nil {
		return time.Duration(0), pb.ErrorErrorReasonBadRequest("empty time range")
	}
	claims, exist := libauth.FromContext(ctx)
	if !exist {
		return time.Duration(0), pb.ErrorErrorReasonUnauthorized("empty token")
	}
	res, err := g.repo.SumAppPackageRunTime(ctx, claims.InternalID, packageID, timeRange)
	if err != nil {
		return time.Duration(0), pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return res, nil
}
