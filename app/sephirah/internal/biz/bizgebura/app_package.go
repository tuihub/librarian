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
	"golang.org/x/exp/slices"
)

func (g *Gebura) CreateAppPackage(
	ctx context.Context,
	a *modelgebura.AppPackage,
) (*modelgebura.AppPackage, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return nil, pb.ErrorErrorReasonForbidden("no permission")
	}
	resp, err := g.searcher.NewID(ctx, &searcher.NewIDRequest{})
	if err != nil {
		logger.Infof("NewID failed: %s", err.Error())
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	a.ID = converter.ToBizInternalID(resp.Id)
	a.Source = modelgebura.AppPackageSourceManual
	a.SourceID = 0
	a.SourcePackageID = strconv.FormatInt(int64(converter.ToBizInternalID(resp.Id)), 10)
	if _, err = g.mapper.InsertVertex(ctx, &mapper.InsertVertexRequest{
		VertexList: []*mapper.Vertex{{
			Vid:  int64(a.ID),
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

func (g *Gebura) UpdateAppPackage(ctx context.Context, a *modelgebura.AppPackage) *errors.Error {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return pb.ErrorErrorReasonForbidden("no permission")
	}
	a.Source = modelgebura.AppPackageSourceManual
	err := g.repo.UpdateAppPackage(ctx, a)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (g *Gebura) ListAppPackages(
	ctx context.Context,
	paging model.Paging,
	sources []modelgebura.AppPackageSource,
	ids []model.InternalID,
) ([]*modelgebura.AppPackage, int, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return nil, 0, pb.ErrorErrorReasonForbidden("no permission")
	}
	res, total, err := g.repo.ListAppPackages(ctx, paging, sources, ids)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return res, total, nil
}

func (g *Gebura) AssignAppPackage(
	ctx context.Context,
	appID model.InternalID,
	appPackageID model.InternalID,
) *errors.Error {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return pb.ErrorErrorReasonForbidden("no permission")
	}
	if claims, ok := libauth.FromContext(ctx); !ok {
		return pb.ErrorErrorReasonForbidden("no permission")
	} else {
		err := g.repo.AssignAppPackage(ctx, claims.InternalID, appID, appPackageID)
		if err != nil {
			return pb.ErrorErrorReasonUnspecified("%s", err)
		}
	}
	return nil
}

func (g *Gebura) UnAssignAppPackage(ctx context.Context, appPackageID model.InternalID) *errors.Error {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return pb.ErrorErrorReasonForbidden("no permission")
	}
	if claims, ok := libauth.FromContext(ctx); !ok {
		return pb.ErrorErrorReasonForbidden("no permission")
	} else {
		err := g.repo.UnAssignAppPackage(ctx, claims.InternalID, appPackageID)
		if err != nil {
			return pb.ErrorErrorReasonUnspecified("%s", err)
		}
	}
	return nil
}

func (g *Gebura) NewReportAppPackageHandler(ctx context.Context) (ReportAppPackageHandler, *errors.Error) {
	claims, exist := libauth.FromContext(ctx)
	if !exist || claims == nil {
		return nil, pb.ErrorErrorReasonUnauthorized("token required")
	}
	ids, err := g.repo.ListAllAppPackageIDOfOneSource(ctx, modelgebura.AppPackageSourceSentinel, claims.InternalID)
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
	sourceID   model.InternalID
	packageIDs []string
}

func (r *reportAppPackageHandler) Handle(ctx context.Context, apl []*modelgebura.AppPackage) *errors.Error {
	var vl []*mapper.Vertex
	for i := range apl {
		if !slices.Contains(r.packageIDs, apl[i].SourcePackageID) {
			resp, err := r.g.searcher.NewID(ctx, &searcher.NewIDRequest{})
			if err != nil {
				logger.Infof("NewID failed: %s", err.Error())
				return pb.ErrorErrorReasonUnspecified("%s", err.Error())
			}
			apl[i].ID = converter.ToBizInternalID(resp.Id)
			vl = append(vl, &mapper.Vertex{
				Vid:  int64(converter.ToBizInternalID(resp.Id)),
				Type: mapper.VertexType_VERTEX_TYPE_OBJECT,
				Prop: nil,
			})
		}
		apl[i].Source = modelgebura.AppPackageSourceSentinel
		apl[i].SourceID = r.sourceID
	}
	if len(vl) > 0 {
		if _, err := r.g.mapper.InsertVertex(ctx, &mapper.InsertVertexRequest{VertexList: vl}); err != nil {
			return pb.ErrorErrorReasonUnspecified("%s", err.Error())
		}
		if err := r.g.repo.UpsertAppPackages(ctx, r.sourceID, apl); err != nil {
			return pb.ErrorErrorReasonUnspecified("%s", err.Error())
		}
	}
	return nil
}
