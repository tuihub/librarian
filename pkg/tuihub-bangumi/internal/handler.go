package internal

import (
	"context"
	"encoding/json"
	"strconv"
	"sync"

	"github.com/tuihub/librarian/pkg/tuihub-bangumi/internal/biz"
	"github.com/tuihub/librarian/pkg/tuihub-bangumi/internal/model"
	"github.com/tuihub/librarian/pkg/tuihub-go"
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

	// Validate required token
	if config.Token == "" {
		return nil, errors.BadRequest("token_required", "Bangumi API token is required")
	}

	h.clientMap.Store(req.GetContextId().GetId(), config)
	return &porter.EnableContextResponse{}, nil
}

func (h *Handler) DisableContext(ctx context.Context, req *porter.DisableContextRequest) (
	*porter.DisableContextResponse, error) {
	h.clientMap.Delete(req.GetContextId().GetId())
	return &porter.DisableContextResponse{}, nil
}

func (h *Handler) GetAppInfo(ctx context.Context, req *porter.GetAppInfoRequest) (
	*porter.GetAppInfoResponse, error) {
	var config model.GetAppInfoConfig
	err := json.Unmarshal([]byte(req.GetConfig().GetConfigJson()), &config)
	if err != nil {
		return nil, errors.BadRequest("invalid context_json", err.Error())
	}

	clientAny, ok := h.clientMap.Load(req.GetConfig().GetContextId().GetId())
	if !ok {
		return nil, errors.BadRequest("context not found", "")
	}
	client, ok := clientAny.(model.PorterContext)
	if !ok {
		return nil, errors.BadRequest("invalid context", "")
	}

	bangumi := biz.NewBangumiUseCase(client.Token)
	app, err := bangumi.GetSubject(ctx, config.AppID)
	if err != nil {
		return nil, err
	}

	return &porter.GetAppInfoResponse{
		AppInfo: &porter.AppInfo{
			Source:      tuihub.WellKnownToString(librarian.WellKnownAppInfoSource_WELL_KNOWN_APP_INFO_SOURCE_BANGUMI),
			SourceAppId: config.AppID,
			SourceUrl:   &app.StoreURL,
			RawDataJson: "",
			Details: &porter.AppInfoDetails{
				Description: app.Description,
				ReleaseDate: app.ReleaseDate,
				Developer:   app.Developer,
				Publisher:   app.Publisher,
				Version:     "",
				ImageUrls:   app.ImageURLs,
			},
			Name:               app.Name,
			Type:               ToPBAppType(app.Type),
			ShortDescription:   app.ShortDescription,
			IconImageUrl:       "",
			BackgroundImageUrl: app.BackgroundImageURL,
			CoverImageUrl:      app.CoverImageURL,
			Tags:               app.Tags,
			NameAlternatives:   []string{app.NameCN},
		},
	}, nil
}

func (h *Handler) SearchAppInfo(ctx context.Context, req *porter.SearchAppInfoRequest) (
	*porter.SearchAppInfoResponse, error) {
	var config model.SearchAppInfoConfig
	err := json.Unmarshal([]byte(req.GetConfig().GetConfigJson()), &config)
	if err != nil {
		return nil, errors.BadRequest("invalid context_json", err.Error())
	}

	clientAny, ok := h.clientMap.Load(req.GetConfig().GetContextId().GetId())
	if !ok {
		return nil, errors.BadRequest("context not found", "")
	}
	client, ok := clientAny.(model.PorterContext)
	if !ok {
		return nil, errors.BadRequest("invalid context", "")
	}

	bangumi := biz.NewBangumiUseCase(client.Token)
	apps, err := bangumi.SearchSubjects(ctx, config.NameLike)
	if err != nil {
		return nil, err
	}

	appList := make([]*porter.AppInfo, len(apps))
	for i, app := range apps {
		appList[i] = &porter.AppInfo{
			Source: tuihub.WellKnownToString(
				librarian.WellKnownAppInfoSource_WELL_KNOWN_APP_INFO_SOURCE_BANGUMI,
			),
			SourceAppId: strconv.Itoa(app.ID),
			SourceUrl:   &app.StoreURL,
			RawDataJson: "",
			Details: &porter.AppInfoDetails{
				Description: app.Description,
				ReleaseDate: app.ReleaseDate,
				Developer:   app.Developer,
				Publisher:   app.Publisher,
				Version:     "",
				ImageUrls:   app.ImageURLs,
			},
			Name:               app.Name,
			Type:               ToPBAppType(app.Type),
			ShortDescription:   app.ShortDescription,
			IconImageUrl:       "",
			BackgroundImageUrl: app.BackgroundImageURL,
			CoverImageUrl:      app.CoverImageURL,
			Tags:               app.Tags,
			NameAlternatives:   []string{app.NameCN},
		}
	}

	return &porter.SearchAppInfoResponse{AppInfos: appList}, nil
}

func ToPBAppType(t biz.AppType) porter.AppType {
	switch t {
	case biz.AppTypeGame:
		return porter.AppType_APP_TYPE_GAME
	case biz.AppTypeAnime, biz.AppTypeBook, biz.AppTypeMusic, biz.AppTypeReal:
		// Map other types to a generic media type or unspecified
		return porter.AppType_APP_TYPE_UNSPECIFIED
	default:
		return porter.AppType_APP_TYPE_UNSPECIFIED
	}
}
