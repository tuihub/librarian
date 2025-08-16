package internal

import (
	"context"
	"strconv"

	"github.com/tuihub/librarian/pkg/tuihub-go"
	"github.com/tuihub/librarian/pkg/tuihub-steam/internal/biz"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

type Handler struct {
	porter.UnimplementedLibrarianPorterServiceServer

	steam *biz.SteamUseCase
}

func NewHandler(apiKey string) *Handler {
	return &Handler{
		UnimplementedLibrarianPorterServiceServer: porter.UnimplementedLibrarianPorterServiceServer{},
		steam: biz.NewSteamUseCase(apiKey),
	}
}

func (h Handler) GetAccount(ctx context.Context, req *porter.GetAccountRequest) (
	*porter.GetAccountResponse, error) {
	u, err := h.steam.GetUser(ctx, req.GetPlatformAccountId())
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

func (h Handler) GetAppInfo(ctx context.Context, req *porter.GetAppInfoRequest) (
	*porter.GetAppInfoResponse, error) {
	appID, err := strconv.Atoi(req.GetSourceAppId())
	if err != nil {
		return nil, err
	}
	a, err := h.steam.GetAppDetails(ctx, appID)
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

func (h Handler) SearchAppInfo(ctx context.Context, req *porter.SearchAppInfoRequest) (
	*porter.SearchAppInfoResponse, error) {
	al, err := h.steam.SearchAppByName(ctx, req.GetNameLike())
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
