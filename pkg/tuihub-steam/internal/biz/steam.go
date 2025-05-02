package biz

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/tuihub/librarian/pkg/tuihub-go/logger"
	"github.com/tuihub/librarian/pkg/tuihub-steam/internal/client"
	"github.com/tuihub/librarian/pkg/tuihub-steam/internal/data"
	"github.com/tuihub/librarian/pkg/tuihub-steam/internal/model"

	"github.com/go-kratos/kratos/v2/errors"
)

type User struct {
	SteamID    string
	Name       string
	ProfileURL string
	AvatarURL  string
}

type App struct {
	AppID              uint
	StoreURL           string
	Name               string
	Type               AppType
	ShortDescription   string
	BackgroundImageURL string
	CoverImageURL      string
	LogoImageURL       string
	IconImageURL       string
	Description        string
	ReleaseDate        string
	Developer          string
	Publisher          string
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

const (
	CDNCoverImageURL  = "https://steamcdn-a.akamaihd.net/steam/apps/%d/library_600x900_2x.jpg"
	CDNLogoImageURL   = "https://steamcdn-a.akamaihd.net/steam/apps/%d/logo.png"
	CDNHeaderImageURL = "https://steamcdn-a.akamaihd.net/steam/apps/%d/library_hero.jpg"
)

type SteamUseCase struct {
	c      *client.Steam
	locale Locale

	appNameIndex               *data.Data
	appIDNameMap               map[string]string
	appNameIndexUpdateAt       time.Time
	appNameIndexUpdateInterval time.Duration
	appNameIndexUpdateMu       sync.Mutex
	appNameIndexInitialized    bool
}

func NewSteamUseCase(apiKey string) *SteamUseCase {
	cli, err := client.NewSteam(apiKey)
	if err != nil {
		panic(err)
	}
	index, err := data.NewData()
	if err != nil {
		panic(err)
	}
	return &SteamUseCase{
		c:                          cli,
		locale:                     GetLocale(),
		appNameIndex:               index,
		appIDNameMap:               make(map[string]string),
		appNameIndexUpdateAt:       time.Time{},
		appNameIndexUpdateInterval: 24 * time.Hour, //nolint:mnd // no need
		appNameIndexUpdateMu:       sync.Mutex{},
		appNameIndexInitialized:    false,
	}
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
		AvatarURL:  resp.AvatarFull,
	}, nil
}

func (s *SteamUseCase) GetOwnedGames(ctx context.Context, steamID string) ([]*App, error) {
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
		res[i].BackgroundImageURL = fmt.Sprintf(CDNHeaderImageURL, game.AppID)
		res[i].CoverImageURL = fmt.Sprintf(CDNCoverImageURL, game.AppID)
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
		res.AppID = uint(appID) //nolint:gosec // TODO
		res.StoreURL = fmt.Sprintf("https://store.steampowered.com/app/%d", appID)
		if app.Success {
			res = &App{
				AppID:              uint(app.Data.AppID), //nolint:gosec // TODO
				StoreURL:           fmt.Sprintf("https://store.steampowered.com/app/%d", app.Data.AppID),
				Name:               app.Data.Name,
				Type:               AppType(app.Data.Type),
				ShortDescription:   app.Data.ShortDescription,
				BackgroundImageURL: fmt.Sprintf(CDNHeaderImageURL, app.Data.AppID),
				CoverImageURL:      fmt.Sprintf(CDNCoverImageURL, app.Data.AppID),
				LogoImageURL:       fmt.Sprintf(CDNLogoImageURL, app.Data.AppID),
				IconImageURL:       "",
				Description:        app.Data.DetailedDescription,
				ReleaseDate:        "Coming Soon",
				Developer:          strings.Join(app.Data.Developers, ","),
				Publisher:          strings.Join(app.Data.Publishers, ","),
			}
			if !app.Data.ReleaseDate.ComingSoon {
				res.ReleaseDate = app.Data.ReleaseDate.Date
			}
		}
		break
	}
	return res, nil
}

func (s *SteamUseCase) SearchAppByName(ctx context.Context, name string) ([]*App, error) {
	go s.updateAppNameIndex(context.Background())
	s.waitForAppNameIndexInitialized()
	appIDs, err := s.appNameIndex.Search(name)
	if err != nil {
		return nil, err
	}
	if len(appIDs) == 0 {
		return nil, errors.NotFound("app not found", "")
	}
	var res []*App
	for _, id := range appIDs {
		originName, ok := s.appIDNameMap[id]
		if !ok {
			continue
		}
		var appID int
		appID, err = strconv.Atoi(id)
		if err != nil {
			continue
		}
		res = append(res, &App{
			AppID:              uint(appID), //nolint:gosec // TODO
			StoreURL:           fmt.Sprintf("https://store.steampowered.com/app/%d", appID),
			Name:               originName,
			Type:               AppTypeGame,
			ShortDescription:   "",
			BackgroundImageURL: fmt.Sprintf(CDNHeaderImageURL, appID),
			CoverImageURL:      fmt.Sprintf(CDNCoverImageURL, appID),
			LogoImageURL:       fmt.Sprintf(CDNLogoImageURL, appID),
			IconImageURL:       "",
			Description:        "",
			ReleaseDate:        "",
			Developer:          "",
			Publisher:          "",
		})
	}
	return res, nil
}

func (s *SteamUseCase) appNameIndexOutdated() bool {
	return s.appNameIndexUpdateAt.IsZero() || time.Since(s.appNameIndexUpdateAt) > s.appNameIndexUpdateInterval
}

func (s *SteamUseCase) waitForAppNameIndexInitialized() {
	for {
		if s.appNameIndexInitialized {
			return
		}
		time.Sleep(time.Microsecond)
	}
}

func (s *SteamUseCase) updateAppNameIndex(ctx context.Context) {
	if !s.appNameIndexOutdated() {
		return
	}
	if !s.appNameIndexUpdateMu.TryLock() {
		return
	}
	defer s.appNameIndexUpdateMu.Unlock()
	logger.Info("updating app name index")
	list, err := s.c.GetAppList(ctx, model.GetAppListRequest{
		Language: s.language(),
	})
	if err != nil {
		logger.Errorf("update app name index failed: %s", err.Error())
		s.appNameIndexUpdateAt = time.Now()
		return
	}
	for _, app := range list.Apps {
		err = s.appNameIndex.Index(app.AppID, app.Name)
		if err != nil {
			continue
		}
		s.appIDNameMap[strconv.Itoa(app.AppID)] = app.Name
	}
	logger.Info("app name index updated")
	s.appNameIndexInitialized = true
	s.appNameIndexUpdateAt = time.Now()
}
