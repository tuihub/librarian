package bizgebura

import (
	"context"

	"github.com/tuihub/librarian/internal/biz/bizutils"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/model/modelkether"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

func (g *Gebura) ListAppInfos(
	ctx context.Context,
	paging model.Paging,
	sources []string,
	types []modelgebura.AppType,
	ids []model.InternalID,
	containDetails bool,
) ([]*modelgebura.AppInfo, int64, *errors.Error) {
	if libauth.FromContextAssertUserType(ctx, model.UserTypeAdmin) == nil {
		return nil, 0, bizutils.NoPermissionError()
	}
	infos, total, err := g.repo.ListAppInfos(ctx, paging, sources, types, ids)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return infos, total, nil
}

func (g *Gebura) SyncAppInfos(
	ctx context.Context,
	infoIDs []*modelgebura.AppInfoID,
	wait bool,
) ([]*modelgebura.AppInfo, *errors.Error) {
	if libauth.FromContextAssertUserType(ctx) == nil {
		return nil, bizutils.NoPermissionError()
	}
	appInfos := make([]*modelgebura.AppInfo, 0, len(infoIDs))
	ids, err := g.id.BatchNew(len(infoIDs))
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err)
	}
	for i, infoID := range infoIDs {
		if infoID == nil {
			continue
		}
		if wait {
			err = g.pullAppInfo.LocalCall(ctx, modelkether.PullAppInfo{
				ID:              ids[i],
				AppInfoID:       *infoID,
				IgnoreRateLimit: false,
			})
			if err != nil {
				return nil, pb.ErrorErrorReasonUnspecified("%s", err)
			}
			var app *modelgebura.AppInfo
			app, err = g.appInfoCache.Get(ctx, *infoID)
			if err != nil {
				continue
			}
			appInfos = append(appInfos, app)
		} else {
			_ = g.pullAppInfo.Publish(ctx, modelkether.PullAppInfo{
				ID:              ids[i],
				AppInfoID:       *infoID,
				IgnoreRateLimit: false,
			})
			appInfo := new(modelgebura.AppInfo)
			appInfo.ID = ids[i]
			appInfo.Source = infoID.Source
			appInfo.SourceAppID = infoID.SourceAppID
			appInfo, err = g.repo.CreateAppInfoOrGet(ctx, appInfo)
			if err != nil {
				continue
			}
			appInfos = append(appInfos, appInfo)
		}
	}
	return appInfos, nil
}

// func (g *Gebura) SearchAppInfos(ctx context.Context, paging model.Paging, query string) (
//	[]*modelgebura.AppInfoMixed, int, *errors.Error) {
//	if libauth.FromContextAssertUserType(ctx) == nil {
//		return nil, 0, bizutils.NoPermissionError()
//	}
//	results, err := g.search.SearchID(ctx, libsearch.SearchIndexGeburaApp, paging, query)
//	if err != nil {
//		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err)
//	}
//	ids := make([]model.InternalID, 0, len(results))
//	for _, r := range results {
//		ids = append(ids, r.ID)
//	}
//	infos, err := g.repo.GetBatchBoundAppInfos(ctx, ids)
//	if err != nil {
//		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err)
//	}
//	res := make([]*modelgebura.AppInfoMixed, 0, len(infos))
//	for _, a := range infos {
//		res = append(res, a.Flatten())
//	}
//	return res, 0, nil
//}

// func (g *Gebura) SearchNewAppInfos(
//	ctx context.Context,
//	paging model.Paging,
//	name string,
//	sourceFilter []string,
// ) ([]*modelgebura.AppInfo, int, *errors.Error) {
//	if libauth.FromContextAssertUserType(ctx) == nil {
//		return nil, 0, bizutils.NoPermissionError()
//	}
//	//TODO: fix
//	// if len(sourceFilter) == 0 {
//	// sourceFilter = g.supv.GetFeatureSummary().AppInfoSources
//	//}
//	if len(sourceFilter) == 0 {
//		return nil, 0, pb.ErrorErrorReasonBadRequest("no available info source")
//	}
//	var infos []*modelgebura.AppInfo
//	for _, source := range sourceFilter {
//		info, err := g.porter.SearchAppInfo(g.supv.WithAppInfoSource(ctx, source), &porter.SearchAppInfoRequest{
//			NameLike: name,
//		})
//		if err != nil {
//			continue
//		}
//		infos = append(infos, converter.ToBizAppInfoList(info.GetAppInfos())...)
//	}
//	return infos, 0, nil
//}

// func (g *Gebura) GetAppInfo(ctx context.Context, id model.InternalID) (*modelgebura.AppInfo, *errors.Error) {
//	if libauth.FromContextAssertUserType(ctx) == nil {
//		return nil, bizutils.NoPermissionError()
//	}
//	infos, err := g.repo.GetBoundAppInfos(ctx, id)
//	if err != nil {
//		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
//	}
//	for _, info := range infos {
//		if info.ID == id {
//			return info, nil
//		}
//	}
//	return nil, pb.ErrorErrorReasonNotFound("info not found")
//}
//
// func (g *Gebura) GetBoundAppInfos(ctx context.Context, id model.InternalID) ([]*modelgebura.AppInfo, *errors.Error) {
//	if libauth.FromContextAssertUserType(ctx) == nil {
//		return nil, bizutils.NoPermissionError()
//	}
//	apps, err := g.repo.GetBoundAppInfos(ctx, id)
//	if err != nil {
//		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
//	}
//	return apps, nil
//}
//
// func (g *Gebura) PurchaseAppInfo(ctx context.Context, infoID *modelgebura.AppInfoID) (model.InternalID, *errors.Error) {
//	claims := libauth.FromContextAssertUserType(ctx)
//	if claims == nil {
//		return 0, bizutils.NoPermissionError()
//	}
//	id, err := g.repo.PurchaseAppInfo(ctx, claims.UserID, infoID, func(ctx2 context.Context) error {
//		if infoID.Internal {
//			return errors.BadRequest("app not found", "")
//		}
//		ids, err := g.id.BatchNew(2) //nolint:mnd // checked
//		if err != nil {
//			return err
//		}
//		internalInfo := new(modelgebura.AppInfo)
//		internalInfo.ID = ids[0]
//		internalInfo.Internal = true
//		internalInfo.Source = ""
//		internalInfo.SourceAppID = fmt.Sprintf("%d", ids[0])
//		internalInfo.Type = modelgebura.AppTypeGame
//		internalInfo.BoundInternal = ids[0]
//		externalInfo := new(modelgebura.AppInfo)
//		externalInfo.ID = ids[1]
//		externalInfo.Internal = false
//		externalInfo.Source = infoID.Source
//		externalInfo.SourceAppID = infoID.SourceAppID
//		externalInfo.Type = modelgebura.AppTypeGame
//		externalInfo.BoundInternal = ids[0]
//		err = g.repo.CreateAppInfo(ctx2, internalInfo)
//		if err != nil {
//			return err
//		}
//		err = g.repo.CreateAppInfo(ctx2, externalInfo)
//		if err != nil {
//			return err
//		}
//		return nil
//	})
//	if err != nil {
//		return 0, pb.ErrorErrorReasonUnspecified("%s", err)
//	}
//	return id, nil
//}
//
// func (g *Gebura) GetPurchasedAppInfos(ctx context.Context, source string) ([]*modelgebura.AppInfoMixed, *errors.Error) {
//	claims := libauth.FromContextAssertUserType(ctx)
//	if claims == nil {
//		return nil, bizutils.NoPermissionError()
//	}
//	infos, err := g.repo.GetPurchasedAppInfos(ctx, claims.UserID, source)
//	if err != nil {
//		return nil, pb.ErrorErrorReasonUnspecified("%s", err)
//	}
//	res := make([]*modelgebura.AppInfoMixed, 0, len(infos))
//	for _, a := range infos {
//		res = append(res, a.Flatten())
//	}
//	return res, nil
//}
