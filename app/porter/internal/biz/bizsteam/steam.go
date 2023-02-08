package bizsteam

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/tuihub/librarian/app/porter/internal/client/steam"
	"github.com/tuihub/librarian/app/porter/internal/client/steam/model"
	"github.com/tuihub/librarian/app/porter/internal/data"
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
	ImageURL         string
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
	c *steam.Steam
}

func NewSteamUseCase(client *steam.Steam, _ *data.Data) *SteamUseCase {
	return &SteamUseCase{c: client}
}

func (s *SteamUseCase) GetUser(ctx context.Context, steamID string) (*User, error) {
	id, err := strconv.ParseUint(steamID, 10, 64)
	if err != nil {
		return nil, err
	}
	resp, err := s.c.WebAPI.GetPlayerSummary(ctx, model.GetPlayerSummariesRequest{
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
	id, err := strconv.ParseUint(steamID, 10, 64)
	if err != nil {
		return nil, err
	}
	resp, err := s.c.WebAPI.GetOwnedGames(ctx, model.GetOwnedGamesRequest{
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
	res := make([]*App, len(resp.Games))
	for i, game := range resp.Games {
		res[i] = &App{
			AppID:            game.AppID,
			StoreURL:         "",
			Name:             game.Name,
			Type:             AppTypeGame,
			ShortDescription: "",
			ImageURL:         "",
			Description:      "",
			ReleaseDate:      "",
			Developer:        "",
			Publisher:        "",
		}
	}
	return res, nil
}

func (s *SteamUseCase) GetAppDetails(ctx context.Context, appID int) (*App, error) {
	resp, err := s.c.StoreAPI.GetAppDetails(ctx, model.GetAppDetailsRequest{
		AppIDs:      []int{appID},
		CountryCode: "",
		Language:    "",
	})
	if err != nil {
		return nil, err
	}
	if len(resp) != 1 {
		return nil, errors.New("unexpected result")
	}
	var res *App
	for _, app := range resp {
		res = &App{
			AppID:            uint(appID),
			StoreURL:         fmt.Sprintf("https://store.steampowered.com/app/%d", appID),
			Name:             "",
			Type:             "",
			ShortDescription: "",
			ImageURL:         "",
			Description:      "",
			ReleaseDate:      "",
			Developer:        "",
			Publisher:        "",
		}
		if app.Success {
			res = &App{
				AppID:            uint(app.Data.AppID),
				StoreURL:         fmt.Sprintf("https://store.steampowered.com/app/%d", app.Data.AppID),
				Name:             app.Data.Name,
				Type:             AppType(app.Data.Type),
				ShortDescription: app.Data.ShortDescription,
				ImageURL:         app.Data.HeaderImage,
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
