package internal

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/tuihub/librarian/model/modelfeed"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	"strconv"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h Handler) PullAccount(ctx context.Context, req *porter.PullAccountRequest) (
	*porter.PullAccountResponse, error) {
	return nil, errors.BadRequest("not supported", "")
}

func (h Handler) PullApp(ctx context.Context, req *porter.PullAppRequest) (
	*porter.PullAppResponse, error) {
	return nil, errors.BadRequest("not supported", "")
}

func (h Handler) PullAccountAppRelation(ctx context.Context, req *porter.PullAccountAppRelationRequest) (
	*porter.PullAccountAppRelationResponse, error) {
	return nil, errors.BadRequest("not supported", "")
}

func (h Handler) PullFeed(ctx context.Context, req *porter.PullFeedRequest) (
	*porter.PullFeedResponse, error) {
	return nil, errors.BadRequest("not supported", "")
}

func (h Handler) PushFeedItems(ctx context.Context, req *porter.PushFeedItemsRequest) (
	*porter.PushFeedItemsResponse, error) {
	messages := make(map[string]string)
	for _, item := range modelfeed.NewConverter().FromPBFeedItemList(req.GetItems()) {
		messages[item.Title] = item.Link
	}
	channelIDInt64, err := strconv.ParseInt(req.GetChannelId(), 10, 64)
	if err != nil {
		return nil, errors.BadRequest("invalid channel_id", "")
	}
	err = SendBatch(ctx, req.GetToken(), channelIDInt64, messages)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &porter.PushFeedItemsResponse{}, nil
}
