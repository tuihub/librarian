package bizangela

import (
	"context"
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/model/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelangela"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/model"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

func NewPullAccountTopic(
	a *AngelaBase,
	sr *libmq.Topic[modelangela.PullSteamAccountAppRelation],
) *libmq.Topic[modeltiphereth.PullAccountInfo] {
	return libmq.NewTopic[modeltiphereth.PullAccountInfo](
		"PullAccountInfo",
		func(ctx context.Context, info *modeltiphereth.PullAccountInfo) error {
			switch info.Platform {
			case modeltiphereth.AccountPlatformUnspecified:
			case modeltiphereth.AccountPlatformSteam:
				ctx = libapp.NewContext(ctx, string(porter.FeatureFlag_FEATURE_FLAG_SOURCE_STEAM))
			default:
			}
			resp, err := a.porter.PullAccount(ctx, &porter.PullAccountRequest{AccountId: &librarian.AccountID{
				Platform:          converter.ToPBAccountPlatform(info.Platform),
				PlatformAccountId: info.PlatformAccountID,
			}})
			if err != nil {
				return err
			}
			switch info.Platform {
			case modeltiphereth.AccountPlatformUnspecified:
				return nil
			case modeltiphereth.AccountPlatformSteam:
				err = a.repo.UpdateAccount(ctx, modeltiphereth.Account{
					ID:                info.ID,
					Platform:          info.Platform,
					PlatformAccountID: info.PlatformAccountID,
					Name:              resp.GetAccount().GetName(),
					ProfileURL:        resp.GetAccount().GetProfileUrl(),
					AvatarURL:         resp.GetAccount().GetAvatarUrl(),
					LatestUpdateTime:  time.Time{},
				})
				if err != nil {
					return err
				}
				return sr.
					Publish(ctx, modelangela.PullSteamAccountAppRelation{
						ID:      info.ID,
						SteamID: info.PlatformAccountID,
					})
			default:
				return nil
			}
		},
	)
}

func NewPullSteamAccountAppRelationTopic(
	a *AngelaBase,
	sa *libmq.Topic[modelangela.PullSteamApp],
) *libmq.Topic[modelangela.PullSteamAccountAppRelation] {
	return libmq.NewTopic[modelangela.PullSteamAccountAppRelation](
		"PullSteamAccountAppRelation",
		func(ctx context.Context, r *modelangela.PullSteamAccountAppRelation) error {
			ctx = libapp.NewContext(ctx, string(porter.FeatureFlag_FEATURE_FLAG_SOURCE_STEAM))
			var appList []*librarian.App
			if resp, err := a.porter.PullAccountAppRelation(ctx, &porter.PullAccountAppRelationRequest{
				RelationType: porter.AccountAppRelationType_ACCOUNT_APP_RELATION_TYPE_OWN,
				AccountId: &librarian.AccountID{
					Platform:          librarian.AccountPlatform_ACCOUNT_PLATFORM_STEAM,
					PlatformAccountId: r.SteamID,
				},
			}); err != nil {
				return err
			} else {
				appList = resp.GetAppList()
			}
			appNum := len(appList)
			if appNum <= 0 {
				return nil
			}
			steamApps := make([]*modelgebura.App, 0, appNum)
			var steamAppIDs []model.InternalID
			if id, err := a.searcher.NewBatchIDs(ctx, appNum); err != nil {
				return err
			} else {
				steamAppIDs = id
			}
			for i, app := range appList {
				steamApps = append(steamApps, converter.ToBizApp(app))
				steamApps[i].ID = steamAppIDs[i]
				steamApps[i].Source = modelgebura.AppSourceSteam
			}
			if err := a.repo.UpsertApps(ctx, steamApps); err != nil {
				return err
			}
			if err := a.repo.AccountPurchaseApps(ctx, r.ID, steamAppIDs); err != nil {
				return err
			}
			for _, app := range steamApps {
				_ = sa.Publish(ctx, modelangela.PullSteamApp{
					ID:    app.ID,
					AppID: app.SourceAppID,
				})
			}
			return nil
		},
	)
}
