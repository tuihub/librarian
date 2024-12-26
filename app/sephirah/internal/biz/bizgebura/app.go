package bizgebura

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizutils"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

func (g *Gebura) CreateApp(
	ctx context.Context,
	a *modelgebura.App,
) (*modelgebura.App, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, bizutils.NoPermissionError()
	}
	id, err := g.id.New()
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err)
	}
	a.ID = id
	if err = g.repo.CreateApp(ctx, claims.UserID, a); err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return a, nil
}

func (g *Gebura) UpdateApp(ctx context.Context, a *modelgebura.App) *errors.Error {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return bizutils.NoPermissionError()
	}
	err := g.repo.UpdateApp(ctx, claims.UserID, a)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (g *Gebura) ListApps(
	ctx context.Context,
	paging model.Paging,
	ownerIDs []model.InternalID,
	appInfoIDs []model.InternalID,
	ids []model.InternalID,
) ([]*modelgebura.App, int, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, 0, bizutils.NoPermissionError()
	}
	publicOnly := true
	if len(ownerIDs) == 0 {
		ownerIDs = []model.InternalID{claims.UserID}
		publicOnly = false
	}
	res, total, err := g.repo.ListApps(ctx, paging, ownerIDs, appInfoIDs, ids, publicOnly)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return res, total, nil
}

func (g *Gebura) AssignApp(
	ctx context.Context,
	appID model.InternalID,
	appInfoID model.InternalID,
) *errors.Error {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return bizutils.NoPermissionError()
	}
	err := g.repo.AssignApp(ctx, claims.UserID, appID, appInfoID)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err)
	}
	return nil
}

func (g *Gebura) UnAssignApp(ctx context.Context, appID model.InternalID) *errors.Error {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return bizutils.NoPermissionError()
	} else {
		err := g.repo.UnAssignApp(ctx, claims.UserID, appID)
		if err != nil {
			return pb.ErrorErrorReasonUnspecified("%s", err)
		}
	}
	return nil
}

// func (g *Gebura) NewReportAppPackageHandler(ctx context.Context) (ReportAppPackageHandler, *errors.Error) {
//	claims := libauth.FromContext(ctx)
//	if claims == nil {
//		return nil, bizutils.NoPermissionError()
//	}
//	checksums, err := g.repo.ListAppPackageBinaryChecksumOfOneSource(ctx,
//		modelgebura.AppPackageSourceSentinel, claims.UserID)
//	if err != nil {
//		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
//	}
//	return &reportAppPackageHandler{
//		g:        g,
//		sourceID: claims.UserID,
//		sha256:   checksums,
//	}, nil
//}
//
// type reportAppPackageHandler struct {
//	g        *Gebura
//	sourceID model.InternalID
//	sha256   []string
//}
//
// func (r *reportAppPackageHandler) Handle(ctx context.Context, binaries []*modelgebura.AppBinary) *errors.Error {
//	var vl []*mapper.Vertex
//	packages := make([]*modelgebura.App, 0, len(binaries))
//	ids, err := r.g.searcher.NewBatchIDs(ctx, len(binaries))
//	if err != nil {
//		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
//	}
//	for i := range binaries {
//		packages = append(packages, new(modelgebura.App))
//		if !slices.Contains(r.sha256, string(binaries[i].Sha256)) {
//			packages[i].ID = ids[i]
//			vl = append(vl, &mapper.Vertex{
//				Vid:  int64(ids[i]),
//				Type: mapper.VertexType_VERTEX_TYPE_OBJECT,
//				Prop: nil,
//			})
//		}
//		packages[i].Source = modelgebura.AppPackageSourceSentinel
//		packages[i].SourceID = r.sourceID
//	}
//	if len(vl) > 0 {
//		// if _, err = r.g.mapper.InsertVertex(ctx, &mapper.InsertVertexRequest{VertexList: vl}); err != nil {
//		//	return pb.ErrorErrorReasonUnspecified("%s", err.Error())
//		//}
//		if err = r.g.repo.UpsertApps(ctx, r.sourceID, packages); err != nil {
//			return pb.ErrorErrorReasonUnspecified("%s", err.Error())
//		}
//	}
//	return nil
//}
