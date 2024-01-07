package biz

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/tuihub/librarian/pkg/porter-steam/internal/client"
	"github.com/tuihub/librarian/pkg/porter-steam/internal/model"

	"github.com/go-kratos/kratos/v2/errors"
)

type User struct {
	SteamID    string
	Name       string
	ProfileURL string
	AvatarURL  string
}

type App struct {
	AppID            uint
	StoreURL         string
	Name             string
	Type             AppType
	ShortDescription string
	HeroImageURL     string
	LogoImageURL     string
	IconImageURL     string
	Description      string
	ReleaseDate      string
	Developer        string
	Publisher        string
}

type AppType string

const (
	AppTypeGame        AppType = "game"
	AppTypeApplication AppType = "application"
	AppTypeTool        AppType = "tool"
	AppTypeDemo        AppType = "demo"
	AppTypeDeprected   AppType = "deprected"
	AppTypeDLC         AppType = "dlc"
	AppTypeGuide       AppType = "guide"
	AppTypeDriver      AppType = "driver"
	AppTypeConfig      AppType = "config"
	AppTypeHardware    AppType = "hardware"
	AppTypeFranchise   AppType = "franchise"
	AppTypeVideo       AppType = "video"
	AppTypePlugin      AppType = "plugin"
	AppTypeMusic       AppType = "music"
	AppTypeSeries      AppType = "series"
	AppTypeComic       AppType = "comic"
	AppTypeBeta        AppType = "beta"
)

type SteamUseCase struct {
	c      *client.Steam
	locale Locale
}

func NewSteamUseCase(apiKey string) *SteamUseCase {
	cli, err := client.NewSteam(apiKey)
	if err != nil {
		panic(err)
	}
	return &SteamUseCase{
		c:      cli,
		locale: GetLocale(),
	}
}

func (s *SteamUseCase) FeatureEnabled() bool {
	return s.c != nil
}

func (s *SteamUseCase) language() model.LanguageCode {
	switch s.locale {
	case LocaleEn:
		return model.LanguageEnglish
	case LocaleChs:
		return model.LanguageChineseSimplified
	case LocaleCht:
		return model.LanguageChineseTraditional
	default:
		return model.LanguageEnglish
	}
}

func (s *SteamUseCase) GetUser(ctx context.Context, steamID string) (*User, error) {
	if !s.FeatureEnabled() {
		return nil, errors.BadRequest("request disabled feature", "")
	}
	id, err := strconv.ParseUint(steamID, 10, 64)
	if err != nil {
		return nil, err
	}
	resp, err := s.c.GetPlayerSummary(ctx, model.GetPlayerSummariesRequest{
		SteamID: id,
	})
	if err != nil {
		return nil, err
	}
	return &User{
		SteamID:    steamID,
		Name:       resp.Nickname,
		ProfileURL: resp.ProfileURL,
		AvatarURL:  resp.Avatar,
	}, nil
}

func (s *SteamUseCase) GetOwnedGames(ctx context.Context, steamID string) ([]*App, error) {
	if !s.FeatureEnabled() {
		return nil, errors.BadRequest("request disabled feature", "")
	}
	id, err := strconv.ParseUint(steamID, 10, 64)
	if err != nil {
		return nil, err
	}
	resp, err := s.c.GetOwnedGames(ctx, model.GetOwnedGamesRequest{
		SteamID:                id,
		IncludeAppInfo:         true,
		IncludePlayedFreeGames: true,
		IncludeFreeSub:         true,
		SkipUnvettedApps:       false,
		Language:               "",
		IncludeExtendedAppInfo: false,
	})
	if err != nil {
		return nil, err
	}
	const imgBaseURL = "http://media.steampowered.com/steamcommunity/public/images/apps"
	res := make([]*App, len(resp.Games))
	for i, game := range resp.Games {
		res[i] = new(App)
		res[i].AppID = game.AppID
		res[i].Name = game.Name
		if len(game.ImgLogoURL) > 0 {
			res[i].LogoImageURL = fmt.Sprintf("%s/%d/%s.jpg", imgBaseURL, game.AppID, game.ImgLogoURL)
		}
		if len(game.ImgIconURL) > 0 {
			res[i].IconImageURL = fmt.Sprintf("%s/%d/%s.jpg", imgBaseURL, game.AppID, game.ImgIconURL)
		}
	}
	return res, nil
}

func (s *SteamUseCase) GetAppDetails(ctx context.Context, appID int) (*App, error) {
	if !s.FeatureEnabled() {
		return nil, errors.BadRequest("request disabled feature", "")
	}
	resp, err := s.c.GetAppDetails(ctx, model.GetAppDetailsRequest{
		AppIDs:      []int{appID},
		CountryCode: "",
		Language:    s.language(),
	})
	if err != nil {
		return nil, err
	}
	if len(resp) != 1 {
		return nil, errors.InternalServer("unexpected result", "")
	}
	var res *App
	for _, app := range resp {
		res = new(App)
		res.AppID = uint(appID)
		res.StoreURL = fmt.Sprintf("https://store.steampowered.com/app/%d", appID)
		if app.Success {
			res = &App{
				AppID:            uint(app.Data.AppID),
				StoreURL:         fmt.Sprintf("https://store.steampowered.com/app/%d", app.Data.AppID),
				Name:             app.Data.Name,
				Type:             AppType(app.Data.Type),
				ShortDescription: app.Data.ShortDescription,
				HeroImageURL:     app.Data.HeaderImage,
				LogoImageURL:     "",
				IconImageURL:     "",
				Description:      app.Data.DetailedDescription,
				ReleaseDate:      "Coming Soon",
				Developer:        strings.Join(app.Data.Developers, ","),
				Publisher:        strings.Join(app.Data.Publishers, ","),
			}
			if !app.Data.ReleaseDate.ComingSoon {
				res.ReleaseDate = app.Data.ReleaseDate.Date
			}
		}
		break
	}
	return res, nil
}
