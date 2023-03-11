package bizgebura

import (
	"context"
	"strconv"

	"github.com/tuihub/librarian/app/sephirah/internal/model/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

func (g *Gebura) CreateApp(ctx context.Context, app *modelgebura.App) (*modelgebura.App, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin) {
		return nil, pb.ErrorErrorReasonForbidden("no permission")
	}
	resp, err := g.searcher.NewID(ctx, &searcher.NewIDRequest{})
	if err != nil {
		logger.Infof("NewID failed: %s", err.Error())
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	app.ID = converter.ToBizInternalID(resp.Id)
	app.Source = modelgebura.AppSourceInternal
	app.SourceAppID = strconv.FormatInt(int64(app.ID), 10)
	app.SourceURL = ""
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
	return nil
}

func (g *Gebura) ListApp(
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
	apps, total, err := g.repo.ListApp(ctx, paging, sources, types, ids, containDetails)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return apps, total, nil
}

func (g *Gebura) GetBindApps(ctx context.Context, id model.InternalID) ([]*modelgebura.App, *errors.Error) {
	app, _, err := g.repo.ListApp(ctx, model.Paging{
		PageSize: 1,
		PageNum:  1,
	}, nil, nil, []model.InternalID{id}, false)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	if len(app) != 1 {
		return nil, pb.ErrorErrorReasonBadRequest("No such app")
	}
	resp, err := g.mapper.FetchEqualVertex(ctx, &mapper.FetchEqualVertexRequest{SrcVid: int64(id)})
	if err != nil {
		logger.Infof("Fetch Equal Vertex failed: %s", err.Error())
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	appids := make([]model.InternalID, len(resp.GetVertexList()))
	for i, v := range resp.GetVertexList() {
		appids[i] = model.InternalID(v.GetVid())
	}
	apps, _, err := g.repo.ListApp(ctx, model.Paging{
		PageSize: 99, //nolint:gomnd //TODO
		PageNum:  1,
	}, nil, nil, appids, true)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return apps, nil
}
