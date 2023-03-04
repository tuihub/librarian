package bizyesod

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/model/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
)

func (y *Yesod) CreateFeedConfig(ctx context.Context, config *modelyesod.FeedConfig) (model.InternalID, error) {
	claims, exist := libauth.FromContext(ctx)
	if !exist {
		return 0, pb.ErrorErrorReasonUnauthorized("empty token")
	}
	resp, err := y.searcher.NewID(ctx, &searcher.NewIDRequest{})
	if err != nil {
		logger.Infof("NewID failed: %s", err.Error())
		return 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	config.ID = converter.ToBizInternalID(resp.Id)
	if _, err = y.mapper.InsertVertex(ctx, &mapper.InsertVertexRequest{
		VertexList: []*mapper.Vertex{{
			Vid:  int64(config.ID),
			Type: mapper.VertexType_VERTEX_TYPE_METADATA,
		}},
	}); err != nil {
		return 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	err = y.repo.CreateFeedConfig(ctx, config, claims.InternalID)
	if err != nil {
		return 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return config.ID, nil
}

func (y *Yesod) UpdateFeedConfig(ctx context.Context, config *modelyesod.FeedConfig) error {
	err := y.repo.UpdateFeedConfig(ctx, config)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}
