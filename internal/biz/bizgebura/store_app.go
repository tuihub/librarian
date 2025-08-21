package bizgebura

import (
	"context"

	"github.com/tuihub/librarian/internal/biz/bizutils"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelgebura"
	pb "github.com/tuihub/protos/pkg/librarian/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

func (g *Gebura) GetStoreApp(
	ctx context.Context, appID model.InternalID,
) (*modelgebura.StoreApp, *errors.Error) {
	if libauth.FromContextAssertUserType(ctx) == nil {
		return nil, bizutils.NoPermissionError()
	}
	storeApp, err := g.repo.GetStoreApp(ctx, appID)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return storeApp, nil
}

func (g *Gebura) ListStoreApps(
	ctx context.Context, page *model.Paging,
) ([]*modelgebura.StoreApp, int64, *errors.Error) {
	if libauth.FromContextAssertUserType(ctx) == nil {
		return nil, 0, bizutils.NoPermissionError()
	}
	storeApps, total, err := g.repo.ListStoreApps(ctx, page)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return storeApps, int64(total), nil
}

func (g *Gebura) ListStoreAppBinaries(
	ctx context.Context, page *model.Paging, appIDs []model.InternalID,
) ([]*modelgebura.StoreAppBinary, int64, *errors.Error) {
	if libauth.FromContextAssertUserType(ctx, model.UserTypeAdmin) == nil {
		return nil, 0, bizutils.NoPermissionError()
	}
	storeAppBinaries, total, err := g.repo.ListStoreAppBinaries(ctx, page, appIDs)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return storeAppBinaries, int64(total), nil
}
