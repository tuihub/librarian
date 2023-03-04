package bizyesod

import (
	"context"

	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
)

func (y *Yesod) CreateFeedConfig(ctx context.Context, config *FeedConfig) (int64, error) {
	claims, exist := libauth.FromContext(ctx)
	if !exist {
		return 0, pb.ErrorErrorReasonUnauthorized("empty token")
	}
	resp, err := y.searcher.NewID(ctx, &searcher.NewIDRequest{})
	if err != nil {
		logger.Infof("NewID failed: %s", err.Error())
		return 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	config.ID = model.InternalID(resp.Id)
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
	return resp.Id, nil
}

func (y *Yesod) UpdateFeedConfig(ctx context.Context, config *FeedConfig) error {
	err := y.repo.UpdateFeedConfig(ctx, config)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}
