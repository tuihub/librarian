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

	"github.com/go-kratos/kratos/v2/errors"
)

func (y *Yesod) CreateFeedConfig(ctx context.Context, config *modelyesod.FeedConfig) (model.InternalID, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return 0, pb.ErrorErrorReasonForbidden("no permission")
	}
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
	err = y.repo.CreateFeedConfig(ctx, claims.InternalID, config)
	if err != nil {
		return 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return config.ID, nil
}

func (y *Yesod) UpdateFeedConfig(ctx context.Context, config *modelyesod.FeedConfig) *errors.Error {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return pb.ErrorErrorReasonForbidden("no permission")
	}
	claims, exist := libauth.FromContext(ctx)
	if !exist {
		return pb.ErrorErrorReasonUnauthorized("empty token")
	}
	err := y.repo.UpdateFeedConfig(ctx, claims.InternalID, config)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (y *Yesod) ListFeedConfigCategories(ctx context.Context) ([]string, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return nil, pb.ErrorErrorReasonForbidden("no permission")
	}
	claims, exist := libauth.FromContext(ctx)
	if !exist {
		return nil, pb.ErrorErrorReasonUnauthorized("empty token")
	}
	res, err := y.repo.ListFeedConfigCategories(ctx, claims.InternalID)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return res, nil
}
