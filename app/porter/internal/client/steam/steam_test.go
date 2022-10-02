package steam_test

import (
	"context"
	"os"
	"strconv"
	"testing"

	"github.com/tuihub/librarian/app/porter/internal/client/steam"
	"github.com/tuihub/librarian/app/porter/internal/client/steam/model"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/logger"
)

func getAPIKey() string {
	return os.Getenv("STEAM_API_KEY")
}
func getSteamID() uint64 {
	idStr, exist := os.LookupEnv("STEAM_ID")
	if !exist {
		return 0
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		panic(err)
	}
	return uint64(id)
}
func getAppID() int {
	return 10
}

func TestWebAPI_GetPlayerSummary(t *testing.T) {
	r := steam.NewWebAPI(&conf.Porter_Data{SteamApiKey: getAPIKey()})
	res, err := r.GetPlayerSummary(context.Background(), model.GetPlayerSummariesRequest{SteamID: getSteamID()})
	logger.Infof("res %+v, err: %+v", res, err)
}

func TestWebAPI_GetOwnedGames(t *testing.T) {
	r := steam.NewWebAPI(&conf.Porter_Data{SteamApiKey: getAPIKey()})
	res, err := r.GetOwnedGames(context.Background(), model.GetOwnedGamesRequest{
		SteamID: getSteamID(),
	})
	logger.Infof("res %+v, err: %+v", res, err)
}

func TestStoreAPI_GetAppDetails(t *testing.T) {
	r := steam.NewStoreAPI()
	res, err := r.GetAppDetails(context.Background(), model.GetAppDetailsRequest{
		AppIDs:      []int{getAppID()},
		CountryCode: model.ProductCCUS,
		Language:    model.LanguageWebEN,
	})
	logger.Infof("res %+v, err: %+v", res, err)
}
