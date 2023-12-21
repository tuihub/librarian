package bizyesod

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizutils"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

func (y *Yesod) CreateFeedConfig(ctx context.Context, config *modelyesod.FeedConfig) (model.InternalID, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return 0, bizutils.NoPermissionError()
	}
	id, err := y.searcher.NewID(ctx)
	if err != nil {
		return 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	config.ID = id
	if _, err = y.mapper.InsertVertex(ctx, &mapper.InsertVertexRequest{
		VertexList: []*mapper.Vertex{{
			Vid:  int64(config.ID),
			Type: mapper.VertexType_VERTEX_TYPE_METADATA,
		}},
	}); err != nil {
		return 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	err = y.repo.CreateFeedConfig(ctx, claims.InternalID, config)
	if err != nil {
		return 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return config.ID, nil
}

func (y *Yesod) UpdateFeedConfig(ctx context.Context, config *modelyesod.FeedConfig) *errors.Error {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return bizutils.NoPermissionError()
	}
	err := y.repo.UpdateFeedConfig(ctx, claims.InternalID, config)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (y *Yesod) ListFeedCategories(ctx context.Context) ([]string, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, bizutils.NoPermissionError()
	}
	res, err := y.repo.ListFeedCategories(ctx, claims.InternalID)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return res, nil
}
