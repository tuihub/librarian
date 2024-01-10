package bizangela

import (
	"context"
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/model/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelangela"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/model"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

func NewPullAccountTopic(
	a *AngelaBase,
	sr *libmq.Topic[modelangela.PullAccountAppRelation],
) *libmq.Topic[modeltiphereth.PullAccountInfo] {
	return libmq.NewTopic[modeltiphereth.PullAccountInfo](
		"PullAccountInfo",
		func(ctx context.Context, info *modeltiphereth.PullAccountInfo) error {
			if !a.supv.CheckAccountPlatform(info.Platform) {
				return nil
			}
			resp, err := a.porter.PullAccount(
				a.supv.CallAccountPlatform(ctx, info.Platform),
				&porter.PullAccountRequest{AccountId: &librarian.AccountID{
					Platform:          info.Platform,
					PlatformAccountId: info.PlatformAccountID,
				}},
			)
			if err != nil {
				return err
			}
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
				Publish(ctx, modelangela.PullAccountAppRelation{
					ID:                info.ID,
					Platform:          info.Platform,
					PlatformAccountID: info.PlatformAccountID,
				})
		},
	)
}

func NewPullAccountAppRelationTopic(
	a *AngelaBase,
	sa *libmq.Topic[modelangela.PullApp],
) *libmq.Topic[modelangela.PullAccountAppRelation] {
	return libmq.NewTopic[modelangela.PullAccountAppRelation](
		"PullAccountAppRelation",
		func(ctx context.Context, r *modelangela.PullAccountAppRelation) error {
			if !a.supv.CheckAccountPlatform(r.Platform) {
				return nil
			}
			var appList []*librarian.App
			if resp, err := a.porter.PullAccountAppRelation(
				a.supv.CallAccountPlatform(ctx, r.Platform),
				&porter.PullAccountAppRelationRequest{
					RelationType: librarian.AccountAppRelationType_ACCOUNT_APP_RELATION_TYPE_OWN,
					AccountId: &librarian.AccountID{
						Platform:          r.Platform,
						PlatformAccountId: r.PlatformAccountID,
					},
				},
			); err != nil {
				return err
			} else {
				appList = resp.GetAppList()
			}
			appNum := len(appList)
			if appNum <= 0 {
				return nil
			}
			apps := make([]*modelgebura.App, 0, appNum)
			var appIDs []model.InternalID
			if id, err := a.searcher.NewBatchIDs(ctx, appNum); err != nil {
				return err
			} else {
				appIDs = id
			}
			for i, app := range appList {
				apps = append(apps, converter.ToBizApp(app))
				apps[i].ID = appIDs[i]
				apps[i].Source = r.Platform
			}
			if err := a.repo.UpsertApps(ctx, apps); err != nil {
				return err
			}
			if err := a.repo.AccountPurchaseApps(ctx, r.ID, appIDs); err != nil {
				return err
			}
			for _, app := range apps {
				_ = sa.Publish(ctx, modelangela.PullApp{
					ID:     app.ID,
					Source: r.Platform,
					AppID:  app.SourceAppID,
				})
			}
			return nil
		},
	)
}
