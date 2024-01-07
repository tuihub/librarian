package internal

import (
	"context"
	"strconv"

	portersdk "github.com/tuihub/librarian/pkg/porter-sdk"
	"github.com/tuihub/librarian/pkg/porter-steam/internal/biz"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

type Handler struct {
	steam *biz.SteamUseCase
}

func NewHandler(apiKey string) *Handler {
	return &Handler{
		steam: biz.NewSteamUseCase(apiKey),
	}
}

func (h Handler) PullAccount(ctx context.Context, req *porter.PullAccountRequest) (
	*porter.PullAccountResponse, error) {
	u, err := h.steam.GetUser(ctx, req.GetAccountId().GetPlatformAccountId())
	if err != nil {
		return nil, err
	}
	return &porter.PullAccountResponse{Account: &librarian.Account{
		Id:                nil,
		Platform:          req.GetAccountId().GetPlatform(),
		PlatformAccountId: req.GetAccountId().GetPlatformAccountId(),
		Name:              u.Name,
		ProfileUrl:        u.ProfileURL,
		AvatarUrl:         u.AvatarURL,
		LatestUpdateTime:  nil,
	}}, nil
}

func (h Handler) PullApp(ctx context.Context, req *porter.PullAppRequest) (
	*porter.PullAppResponse, error) {
	appID, err := strconv.Atoi(req.GetAppId().GetSourceAppId())
	if err != nil {
		return nil, err
	}
	a, err := h.steam.GetAppDetails(ctx, appID)
	if err != nil {
		return nil, err
	}
	return &porter.PullAppResponse{App: &librarian.App{
		Id:          nil,
		Internal:    false,
		Source:      req.GetAppId().GetSource(),
		SourceAppId: req.GetAppId().GetSourceAppId(),
		SourceUrl:   &a.StoreURL,
		Details: &librarian.AppDetails{ // TODO
			Description: a.Description,
			ReleaseDate: a.ReleaseDate,
			Developer:   a.Developer,
			Publisher:   a.Publisher,
			Version:     "",
			ImageUrls:   nil,
		},
		Name:             a.Name,
		Type:             ToPBAppType(a.Type),
		ShortDescription: a.ShortDescription,
		IconImageUrl:     "",
		HeroImageUrl:     a.HeroImageURL,
		Tags:             nil,
		AltNames:         nil,
	}}, nil
}

func (h Handler) PullAccountAppRelation(ctx context.Context, req *porter.PullAccountAppRelationRequest) (
	*porter.PullAccountAppRelationResponse, error) {
	al, err := h.steam.GetOwnedGames(ctx, req.GetAccountId().GetPlatformAccountId())
	if err != nil {
		return nil, err
	}
	appList := make([]*librarian.App, len(al))
	for i, a := range al {
		appList[i] = &librarian.App{ // TODO
			Id:       nil,
			Internal: false,
			Source: portersdk.WellKnownToString(
				librarian.WellKnownAppSource_WELL_KNOWN_APP_SOURCE_STEAM,
			),
			SourceAppId:      strconv.Itoa(int(a.AppID)),
			SourceUrl:        nil,
			Details:          nil,
			Name:             a.Name,
			Type:             0,
			ShortDescription: "",
			IconImageUrl:     a.IconImageURL,
			HeroImageUrl:     a.HeroImageURL,
			Tags:             nil,
			AltNames:         nil,
		}
	}
	return &porter.PullAccountAppRelationResponse{AppList: appList}, nil
}

func (h Handler) PullFeed(ctx context.Context, request *porter.PullFeedRequest) (
	*porter.PullFeedResponse, error) {
	return nil, errors.BadRequest("not supported", "")
}

func (h Handler) PushFeedItems(ctx context.Context, request *porter.PushFeedItemsRequest) (
	*porter.PushFeedItemsResponse, error) {
	return nil, errors.BadRequest("not supported", "")
}

func ToPBAppType(t biz.AppType) librarian.AppType {
	switch t {
	case biz.AppTypeGame:
		return librarian.AppType_APP_TYPE_GAME
	default:
		return librarian.AppType_APP_TYPE_UNSPECIFIED
	}
}
