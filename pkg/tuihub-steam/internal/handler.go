package internal

import (
	"context"
	"encoding/json"
	"strconv"
	"sync"

	"github.com/tuihub/librarian/pkg/tuihub-go"
	"github.com/tuihub/librarian/pkg/tuihub-steam/internal/biz"
	"github.com/tuihub/librarian/pkg/tuihub-steam/internal/model"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

type Handler struct {
	porter.UnimplementedLibrarianPorterServiceServer
	steam *biz.SteamUseCase
}

func NewHandler() *Handler {
	return &Handler{
		UnimplementedLibrarianPorterServiceServer: porter.UnimplementedLibrarianPorterServiceServer{},
		clientMap: sync.Map{},
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
	var config model.PorterContext
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

func (h *Handler) GetAccount(ctx context.Context, req *porter.GetAccountRequest) (
	*porter.GetAccountResponse, error) {
	clientAny, ok := h.clientMap.Load(req.GetConfig().GetContextId().GetId())
	if !ok {
		return nil, errors.BadRequest("context not found", "")
	}
	client, ok := clientAny.(model.PorterContext)
	if !ok {
		return nil, errors.BadRequest("invalid context", "")
	}
	steam := biz.NewSteamUseCase(client.APIKey)
	u, err := steam.GetUser(ctx, req.GetPlatformAccountId())
	if err != nil {
		return nil, err
	}
	return &porter.GetAccountResponse{Account: &porter.Account{
		Platform:          req.GetPlatform(),
		PlatformAccountId: req.GetPlatformAccountId(),
		Name:              u.Name,
		ProfileUrl:        u.ProfileURL,
		AvatarUrl:         u.AvatarURL,
	}}, nil
}

func (h *Handler) GetAppInfo(ctx context.Context, req *porter.GetAppInfoRequest) (
	*porter.GetAppInfoResponse, error) {
	clientAny, ok := h.clientMap.Load(req.GetConfig().GetContextId().GetId())
	if !ok {
		return nil, errors.BadRequest("context not found", "")
	}
	client, ok := clientAny.(model.PorterContext)
	if !ok {
		return nil, errors.BadRequest("invalid context", "")
	}
	appID, err := strconv.Atoi(req.GetSourceAppId())
	if err != nil {
		return nil, err
	}
	steam := biz.NewSteamUseCase(client.APIKey)
	a, err := steam.GetAppDetails(ctx, appID)
	if err != nil {
		return nil, err
	}
	return &porter.GetAppInfoResponse{
		AppInfo: &porter.AppInfo{
			Source:      req.GetSource(),
			SourceAppId: req.GetSourceAppId(),
			SourceUrl:   &a.StoreURL,
			RawDataJson: "",
			Details: &porter.AppInfoDetails{ // TODO
				Description: a.Description,
				ReleaseDate: a.ReleaseDate,
				Developer:   a.Developer,
				Publisher:   a.Publisher,
				Version:     "",
				ImageUrls:   nil,
			},
			Name:               a.Name,
			Type:               ToPBAppType(a.Type),
			ShortDescription:   a.ShortDescription,
			IconImageUrl:       "",
			BackgroundImageUrl: a.BackgroundImageURL,
			CoverImageUrl:      a.CoverImageURL,
			Tags:               nil,
			NameAlternatives:   nil,
		},
	}, nil
}

// func (h Handler) GetAccountAppInfoRelation(ctx context.Context, req *porter.GetAccountAppInfoRelationRequest) (
//	*porter.GetAccountAppInfoRelationResponse, error) {
//	al, err := h.steam.GetOwnedGames(ctx, req.GetAccountId().GetPlatformAccountId())
//	if err != nil {
//		return nil, err
//	}
//	appList := make([]*librarian.AppInfo, len(al))
//	for i, a := range al {
//		appList[i] = &librarian.AppInfo{ // TODO
//			Id:       nil,
//			Internal: false,
//			Source: tuihub.WellKnownToString(
//				librarian.WellKnownAppInfoSource_WELL_KNOWN_APP_INFO_SOURCE_STEAM,
//			),
//			SourceAppId:        strconv.Itoa(int(a.AppID)),
//			SourceUrl:          nil,
//			Details:            nil,
//			Name:               a.Name,
//			Type:               0,
//			ShortDescription:   "",
//			IconImageUrl:       a.IconImageURL,
//			BackgroundImageUrl: a.BackgroundImageURL,
//			CoverImageUrl:      a.CoverImageURL,
//			Tags:               nil,
//			AltNames:           nil,
//		}
//	}
//	return &porter.GetAccountAppInfoRelationResponse{AppInfos: appList}, nil
//}

func ToPBAppType(t biz.AppType) porter.AppType {
	switch t {
	case biz.AppTypeGame:
		return porter.AppType_APP_TYPE_GAME
	default:
		return porter.AppType_APP_TYPE_UNSPECIFIED
	}
}

func (h *Handler) SearchAppInfo(ctx context.Context, req *porter.SearchAppInfoRequest) (
	*porter.SearchAppInfoResponse, error) {
	clientAny, ok := h.clientMap.Load(req.GetConfig().GetContextId().GetId())
	if !ok {
		return nil, errors.BadRequest("context not found", "")
	}
	client, ok := clientAny.(model.PorterContext)
	if !ok {
		return nil, errors.BadRequest("invalid context", "")
	}
	steam := biz.NewSteamUseCase(client.APIKey)
	al, err := steam.SearchAppByName(ctx, req.GetNameLike())
	if err != nil {
		return nil, err
	}
	appList := make([]*porter.AppInfo, len(al))
	for i, a := range al {
		appList[i] = &porter.AppInfo{ // TODO
			Source: tuihub.WellKnownToString(
				librarian.WellKnownAppInfoSource_WELL_KNOWN_APP_INFO_SOURCE_STEAM,
			),
			SourceAppId:        strconv.Itoa(int(a.AppID)), //nolint:gosec // TODO
			SourceUrl:          nil,
			RawDataJson:        "",
			Details:            nil,
			Name:               a.Name,
			Type:               0,
			ShortDescription:   "",
			IconImageUrl:       a.IconImageURL,
			BackgroundImageUrl: a.BackgroundImageURL,
			CoverImageUrl:      a.CoverImageURL,
			Tags:               nil,
			NameAlternatives:   nil,
		}
	}
	return &porter.SearchAppInfoResponse{AppInfos: appList}, nil
}
