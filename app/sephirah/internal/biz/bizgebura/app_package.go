package bizgebura

import (
	"context"
	"strconv"

	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/logger"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"golang.org/x/exp/slices"
)

func (g *Gebura) CreateAppPackage(ctx context.Context, a *AppPackage) (*AppPackage, *errors.Error) {
	resp, err := g.searcher.NewID(ctx, &searcher.NewIDRequest{})
	if err != nil {
		logger.Infof("NewID failed: %s", err.Error())
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	a.InternalID = resp.Id
	a.Source = AppPackageSourceManual
	a.SourceID = 0
	a.SourcePackageID = strconv.FormatInt(resp.Id, 10)
	if _, err = g.mapper.InsertVertex(ctx, &mapper.InsertVertexRequest{
		VertexList: []*mapper.Vertex{{
			Vid:  a.InternalID,
			Type: mapper.VertexType_VERTEX_TYPE_ABSTRACT,
		}},
	}); err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	if err = g.repo.CreateAppPackage(ctx, a); err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return a, nil
}

func (g *Gebura) UpdateAppPackage(ctx context.Context, a *AppPackage) *errors.Error {
	a.Source = AppPackageSourceManual
	err := g.repo.UpdateAppPackage(ctx, a)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (g *Gebura) ListAppPackage(
	ctx context.Context,
	paging Paging,
	sources []AppPackageSource,
	ids []int64,
) ([]*AppPackage, *errors.Error) {
	res, err := g.repo.ListAppPackage(ctx, paging, sources, ids)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return res, nil
}

func (g *Gebura) AssignAppPackage(ctx context.Context, app App, appPackage AppPackage) *errors.Error {
	if err := g.repo.IsApp(ctx, app.InternalID); err != nil {
		return pb.ErrorErrorReasonBadRequest("%s", err.Error())
	}
	if err := g.repo.IsAppPackage(ctx, appPackage.InternalID); err != nil {
		return pb.ErrorErrorReasonBadRequest("%s", err.Error())
	}
	if _, err := g.mapper.InsertEdge(ctx, &mapper.InsertEdgeRequest{EdgeList: []*mapper.Edge{{
		SrcVid: app.InternalID,
		DstVid: appPackage.InternalID,
		Type:   mapper.EdgeType_EDGE_TYPE_DESCRIBE,
	}}}); err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (g *Gebura) NewReportAppPackageHandler(ctx context.Context) (ReportAppPackageHandler, *errors.Error) {
	claims, exist := libauth.FromContext(ctx)
	if !exist || claims == nil {
		return nil, pb.ErrorErrorReasonUnauthorized("token required")
	}
	ids, err := g.repo.ListAllAppPackageIDOfOneSource(ctx, AppPackageSourceSentinel, claims.InternalID)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return &reportAppPackageHandler{
		g:          g,
		sourceID:   claims.InternalID,
		packageIDs: ids,
	}, nil
}

type reportAppPackageHandler struct {
	g          *Gebura
	sourceID   int64
	packageIDs []string
}

func (r *reportAppPackageHandler) Handle(ctx context.Context, apl []*AppPackage) *errors.Error {
	var vl []*mapper.Vertex
	for i := range apl {
		if !slices.Contains(r.packageIDs, apl[i].SourcePackageID) {
			resp, err := r.g.searcher.NewID(ctx, &searcher.NewIDRequest{})
			if err != nil {
				logger.Infof("NewID failed: %s", err.Error())
				return pb.ErrorErrorReasonUnspecified("%s", err.Error())
			}
			apl[i].InternalID = resp.Id
			vl = append(vl, &mapper.Vertex{
				Vid:  resp.Id,
				Type: mapper.VertexType_VERTEX_TYPE_OBJECT,
				Prop: nil,
			})
		}
		apl[i].Source = AppPackageSourceSentinel
		apl[i].SourceID = r.sourceID
	}
	if len(vl) > 0 {
		if _, err := r.g.mapper.InsertVertex(ctx, &mapper.InsertVertexRequest{VertexList: vl}); err != nil {
			return pb.ErrorErrorReasonUnspecified("%s", err.Error())
		}
		if err := r.g.repo.UpsertAppPackage(ctx, apl); err != nil {
			return pb.ErrorErrorReasonUnspecified("%s", err.Error())
		}
	}
	return nil
}
