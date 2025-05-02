package internal

import (
	"context"
	"encoding/json"
	"sync"

	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

type Handler struct {
	porter.UnimplementedLibrarianPorterServiceServer
	clientMap sync.Map
}

func NewHandler() *Handler {
	return &Handler{
		porter.UnimplementedLibrarianPorterServiceServer{},
		sync.Map{},
	}
}

func (h *Handler) EnablePorter(ctx context.Context, req *porter.EnablePorterRequest) (
	*porter.EnablePorterResponse, error) {
	var contextIds []*librarian.InternalID
	h.clientMap.Range(func(key, value interface{}) bool {
		id, ok := key.(int64)
		if !ok {
			return true
		}
		contextIds = append(contextIds, &librarian.InternalID{Id: id})
		return true
	})
	return &porter.EnablePorterResponse{
		StatusMessage:    "",
		NeedRefreshToken: false,
		EnablesSummary: &porter.PorterEnablesSummary{
			ContextIds:    contextIds,
			FeedSetterIds: nil,
			FeedGetterIds: nil,
		},
	}, nil
}

func (h *Handler) EnableContext(ctx context.Context, req *porter.EnableContextRequest) (
	*porter.EnableContextResponse, error) {
	var config PorterContext
	err := json.Unmarshal([]byte(req.GetContextJson()), &config)
	if err != nil {
		return nil, errors.BadRequest("invalid context_json", err.Error())
	}
	h.clientMap.Store(req.GetContextId().GetId(), config)
	return &porter.EnableContextResponse{}, nil
}

func (h *Handler) DisableContext(ctx context.Context, req *porter.DisableContextRequest) (
	*porter.DisableContextResponse, error) {
	h.clientMap.Delete(req.GetContextId().GetId())
	return &porter.DisableContextResponse{}, nil
}

func (h *Handler) PushFeedItems(ctx context.Context, req *porter.PushFeedItemsRequest) (
	*porter.PushFeedItemsResponse, error) {
	var config PushFeedItems
	err := json.Unmarshal([]byte(req.GetDestination().GetConfigJson()), &config)
	if err != nil {
		return nil, errors.BadRequest("invalid config_json", err.Error())
	}
	clientAny, ok := h.clientMap.Load(req.GetDestination().GetContextId().GetId())
	if !ok {
		return nil, errors.BadRequest("context not found", "")
	}
	client, ok := clientAny.(PorterContext)
	if !ok {
		return nil, errors.BadRequest("invalid context", "")
	}
	messages := make(map[string]string)
	for _, item := range req.GetItems() {
		messages[item.GetTitle()] = item.GetLink()
	}
	err = SendBatch(ctx, client.Token, config.ChannelID, messages)
	if err != nil {
		return nil, err
	}
	return &porter.PushFeedItemsResponse{}, nil
}
