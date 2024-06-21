package bizgebura

import (
	"context"
	"fmt"
	"strconv"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizutils"
	"github.com/tuihub/librarian/app/sephirah/internal/model/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelangela"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	searcherpb "github.com/tuihub/protos/pkg/librarian/searcher/v1"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

func (g *Gebura) CreateAppInfo(
	ctx context.Context,
	appInfo *modelgebura.AppInfo,
) (*modelgebura.AppInfo, *errors.Error) {
	if libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin) == nil {
		return nil, bizutils.NoPermissionError()
	}
	id, err := g.searcher.NewID(ctx)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err)
	}
	appInfo.ID = id
	appInfo.Internal = true
	appInfo.SourceAppID = strconv.FormatInt(int64(appInfo.ID), 10)
	appInfo.SourceURL = ""
	appInfo.BoundInternal = appInfo.ID
	if err = g.repo.CreateAppInfo(ctx, appInfo); err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	_ = g.updateAppInfoIndex.Publish(ctx, modelangela.UpdateAppInfoIndex{IDs: []model.InternalID{appInfo.ID}})
	return appInfo, nil
}

func (g *Gebura) UpdateAppInfo(ctx context.Context, appInfo *modelgebura.AppInfo) *errors.Error {
	if libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin) == nil {
		return bizutils.NoPermissionError()
	}
	appInfo.Internal = true
	err := g.repo.UpdateAppInfo(ctx, appInfo)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	_ = g.updateAppInfoIndex.Publish(ctx, modelangela.UpdateAppInfoIndex{IDs: []model.InternalID{appInfo.ID}})
	return nil
}

func (g *Gebura) ListAppInfos(
	ctx context.Context,
	paging model.Paging,
	sources []string,
	types []modelgebura.AppType,
	ids []model.InternalID,
	containDetails bool,
) ([]*modelgebura.AppInfo, int64, *errors.Error) {
	if libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin) == nil {
		return nil, 0, bizutils.NoPermissionError()
	}
	infos, total, err := g.repo.ListAppInfos(ctx, paging, sources, types, ids, containDetails)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return infos, total, nil
}

func (g *Gebura) MergeAppInfos(ctx context.Context, base modelgebura.AppInfo, merged model.InternalID) *errors.Error {
	if libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin) == nil {
		return bizutils.NoPermissionError()
	}
	if !base.Internal {
		return pb.ErrorErrorReasonBadRequest("source must be INTERNAL")
	}
	if err := g.repo.MergeAppInfos(ctx, base, merged); err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err)
	}
	_ = g.updateAppInfoIndex.Publish(ctx, modelangela.UpdateAppInfoIndex{IDs: []model.InternalID{base.ID}})
	return nil
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
	ids, err := g.searcher.NewBatchIDs(ctx, len(infoIDs))
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err)
	}
	for i, infoID := range infoIDs {
		if infoID == nil {
			continue
		}
		if wait {
			err = g.pullAppInfo.LocalCall(ctx, modelangela.PullAppInfo{
				ID:              ids[i],
				AppInfoID:       *infoID,
				IgnoreRateLimit: false,
			})
			if err != nil {
				return nil, pb.ErrorErrorReasonUnspecified("%s", err)
			}
			var app *modelgebura.AppInfo
			app, err = g.appInfoCache.GetWithFallBack(ctx, *infoID, nil)
			if err != nil {
				continue
			}
			appInfos = append(appInfos, app)
		} else {
			_ = g.pullAppInfo.Publish(ctx, modelangela.PullAppInfo{
				ID:              ids[i],
				AppInfoID:       *infoID,
				IgnoreRateLimit: false,
			})
			appInfo := new(modelgebura.AppInfo)
			appInfo.ID = ids[i]
			appInfo.Internal = infoID.Internal
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

func (g *Gebura) SearchAppInfos(ctx context.Context, paging model.Paging, query string) (
	[]*modelgebura.AppInfoMixed, int, *errors.Error) {
	if libauth.FromContextAssertUserType(ctx) == nil {
		return nil, 0, bizutils.NoPermissionError()
	}
	ids, err := g.searcher.SearchID(ctx, paging, query, searcherpb.Index_INDEX_GEBURA_APP)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err)
	}
	infos, err := g.repo.GetBatchBoundAppInfos(ctx, ids)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err)
	}
	res := make([]*modelgebura.AppInfoMixed, 0, len(infos))
	for _, a := range infos {
		res = append(res, a.Flatten())
	}
	return res, 0, nil
}

func (g *Gebura) SearchNewAppInfos(
	ctx context.Context,
	paging model.Paging,
	name string,
	sourceFilter []string,
) ([]*modelgebura.AppInfo, int, *errors.Error) {
	if libauth.FromContextAssertUserType(ctx) == nil {
		return nil, 0, bizutils.NoPermissionError()
	}
	//TODO: fix
	// if len(sourceFilter) == 0 {
	// sourceFilter = g.supv.GetFeatureSummary().AppInfoSources
	//}
	if len(sourceFilter) == 0 {
		return nil, 0, pb.ErrorErrorReasonBadRequest("no available info source")
	}
	var infos []*modelgebura.AppInfo
	for _, source := range sourceFilter {
		info, err := g.porter.SearchAppInfo(g.supv.CallAppInfoSource(ctx, source), &porter.SearchAppInfoRequest{
			Name: name,
		})
		if err != nil {
			continue
		}
		infos = append(infos, converter.ToBizAppInfoList(info.GetAppInfos())...)
	}
	return infos, 0, nil
}

func (g *Gebura) GetAppInfo(ctx context.Context, id model.InternalID) (*modelgebura.AppInfo, *errors.Error) {
	if libauth.FromContextAssertUserType(ctx) == nil {
		return nil, bizutils.NoPermissionError()
	}
	infos, err := g.repo.GetBoundAppInfos(ctx, id)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	for _, info := range infos {
		if info.ID == id {
			return info, nil
		}
	}
	return nil, pb.ErrorErrorReasonNotFound("info not found")
}

func (g *Gebura) GetBoundAppInfos(ctx context.Context, id model.InternalID) ([]*modelgebura.AppInfo, *errors.Error) {
	if libauth.FromContextAssertUserType(ctx) == nil {
		return nil, bizutils.NoPermissionError()
	}
	apps, err := g.repo.GetBoundAppInfos(ctx, id)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return apps, nil
}

func (g *Gebura) PurchaseAppInfo(ctx context.Context, infoID *modelgebura.AppInfoID) (model.InternalID, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return 0, bizutils.NoPermissionError()
	}
	id, err := g.repo.PurchaseAppInfo(ctx, claims.UserID, infoID, func(ctx2 context.Context) error {
		if infoID.Internal {
			return errors.BadRequest("app not found", "")
		}
		ids, err := g.searcher.NewBatchIDs(ctx2, 2) //nolint:gomnd // checked
		if err != nil {
			return err
		}
		internalInfo := new(modelgebura.AppInfo)
		internalInfo.ID = ids[0]
		internalInfo.Internal = true
		internalInfo.Source = ""
		internalInfo.SourceAppID = fmt.Sprintf("%d", ids[0])
		internalInfo.Type = modelgebura.AppTypeGame
		internalInfo.BoundInternal = ids[0]
		externalInfo := new(modelgebura.AppInfo)
		externalInfo.ID = ids[1]
		externalInfo.Internal = false
		externalInfo.Source = infoID.Source
		externalInfo.SourceAppID = infoID.SourceAppID
		externalInfo.Type = modelgebura.AppTypeGame
		externalInfo.BoundInternal = ids[0]
		err = g.repo.CreateAppInfo(ctx2, internalInfo)
		if err != nil {
			return err
		}
		err = g.repo.CreateAppInfo(ctx2, externalInfo)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, pb.ErrorErrorReasonUnspecified("%s", err)
	}
	return id, nil
}

func (g *Gebura) GetPurchasedAppInfos(ctx context.Context, source string) ([]*modelgebura.AppInfoMixed, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, bizutils.NoPermissionError()
	}
	infos, err := g.repo.GetPurchasedAppInfos(ctx, claims.UserID, source)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err)
	}
	res := make([]*modelgebura.AppInfoMixed, 0, len(infos))
	for _, a := range infos {
		res = append(res, a.Flatten())
	}
	return res, nil
}
