package bizgebura

import (
	"context"
	"strconv"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modelangela"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	searcherpb "github.com/tuihub/protos/pkg/librarian/searcher/v1"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

func (g *Gebura) CreateApp(ctx context.Context, app *modelgebura.App) (*modelgebura.App, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin) {
		return nil, pb.ErrorErrorReasonForbidden("no permission")
	}
	id, err := g.searcher.NewID(ctx)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err)
	}
	app.ID = id
	app.Source = modelgebura.AppSourceInternal
	app.SourceAppID = strconv.FormatInt(int64(app.ID), 10)
	app.SourceURL = ""
	app.BoundInternal = app.ID
	if _, err = g.mapper.InsertVertex(ctx, &mapper.InsertVertexRequest{
		VertexList: []*mapper.Vertex{{
			Vid:  int64(app.ID),
			Type: mapper.VertexType_VERTEX_TYPE_ABSTRACT,
		}},
	}); err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	if err = g.repo.CreateApp(ctx, app); err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	_ = g.updateAppIndex.Publish(ctx, modelangela.UpdateAppIndex{IDs: []model.InternalID{app.ID}})
	return app, nil
}

func (g *Gebura) UpdateApp(ctx context.Context, app *modelgebura.App) *errors.Error {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin) {
		return pb.ErrorErrorReasonForbidden("no permission")
	}
	app.Source = modelgebura.AppSourceInternal
	err := g.repo.UpdateApp(ctx, app)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	_ = g.updateAppIndex.Publish(ctx, modelangela.UpdateAppIndex{IDs: []model.InternalID{app.ID}})
	return nil
}

func (g *Gebura) ListApps(
	ctx context.Context,
	paging model.Paging,
	sources []modelgebura.AppSource,
	types []modelgebura.AppType,
	ids []model.InternalID,
	containDetails bool,
) ([]*modelgebura.App, int64, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin) {
		return nil, 0, pb.ErrorErrorReasonForbidden("no permission")
	}
	apps, total, err := g.repo.ListApps(ctx, paging, sources, types, ids, containDetails)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return apps, total, nil
}

func (g *Gebura) MergeApps(ctx context.Context, base modelgebura.App, merged model.InternalID) *errors.Error {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin) {
		return pb.ErrorErrorReasonForbidden("no permission")
	}
	if base.Source != modelgebura.AppSourceInternal {
		return pb.ErrorErrorReasonBadRequest("source must be INTERNAL")
	}
	if err := g.repo.MergeApps(ctx, base, merged); err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err)
	}
	_ = g.updateAppIndex.Publish(ctx, modelangela.UpdateAppIndex{IDs: []model.InternalID{base.ID}})
	return nil
}

func (g *Gebura) SearchApps(ctx context.Context, paging model.Paging, keyword string) (
	[]*modelgebura.App, int, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return nil, 0, pb.ErrorErrorReasonForbidden("no permission")
	}
	ids, err := g.searcher.SearchID(ctx, paging, keyword, searcherpb.Index_INDEX_GEBURA_APP)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err)
	}
	apps, err := g.repo.GetBatchBoundApps(ctx, ids)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err)
	}
	res := make([]*modelgebura.App, 0, len(apps))
	for _, a := range apps {
		res = append(res, a.Flatten())
	}
	return res, 0, nil
}

func (g *Gebura) GetApp(ctx context.Context, id model.InternalID) (*modelgebura.App, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return nil, pb.ErrorErrorReasonForbidden("no permission")
	}
	apps, err := g.repo.GetBoundApps(ctx, id)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	var res modelgebura.BoundApps
	for _, app := range apps {
		if app.Source == modelgebura.AppSourceInternal {
			res.Internal = app
		}
		if app.Source == modelgebura.AppSourceSteam {
			res.Steam = app
		}
	}
	return res.Flatten(), nil
}

func (g *Gebura) GetBindApps(ctx context.Context, id model.InternalID) ([]*modelgebura.App, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return nil, pb.ErrorErrorReasonForbidden("no permission")
	}
	apps, err := g.repo.GetBoundApps(ctx, id)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return apps, nil
}

func (g *Gebura) PurchaseApp(ctx context.Context, id model.InternalID) *errors.Error {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return pb.ErrorErrorReasonForbidden("no permission")
	}
	if claims, ok := libauth.FromContext(ctx); !ok {
		return pb.ErrorErrorReasonForbidden("no permission")
	} else {
		err := g.repo.PurchaseApp(ctx, claims.InternalID, id)
		if err != nil {
			return pb.ErrorErrorReasonUnspecified("%s", err)
		}
	}
	return nil
}

func (g *Gebura) GetPurchasedApps(ctx context.Context) ([]*modelgebura.App, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return nil, pb.ErrorErrorReasonForbidden("no permission")
	}
	if claims, ok := libauth.FromContext(ctx); !ok {
		return nil, pb.ErrorErrorReasonForbidden("no permission")
	} else {
		apps, err := g.repo.GetPurchasedApps(ctx, claims.InternalID)
		if err != nil {
			return nil, pb.ErrorErrorReasonUnspecified("%s", err)
		}
		return apps, nil
	}
}
