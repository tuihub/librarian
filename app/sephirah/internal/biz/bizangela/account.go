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
	sr *libmq.Topic[modelangela.PullAccountAppInfoRelation],
) *libmq.Topic[modeltiphereth.PullAccountInfo] {
	return libmq.NewTopic[modeltiphereth.PullAccountInfo](
		"PullAccountInfo",
		func(ctx context.Context, info *modeltiphereth.PullAccountInfo) error {
			if !a.supv.HasAccountPlatform(info.Platform) {
				return nil
			}
			resp, err := a.porter.PullAccount(
				a.supv.WithAccountPlatform(ctx, info.Platform),
				&porter.PullAccountRequest{AccountId: &librarian.AccountID{
					Platform:          info.Platform,
					PlatformAccountId: info.PlatformAccountID,
				}},
			)
			if err != nil {
				return err
			}
			err = a.repo.UpsertAccount(ctx, modeltiphereth.Account{
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
				Publish(ctx, modelangela.PullAccountAppInfoRelation{
					ID:                info.ID,
					Platform:          info.Platform,
					PlatformAccountID: info.PlatformAccountID,
				})
		},
	)
}

func NewPullAccountAppInfoRelationTopic(
	a *AngelaBase,
	sa *libmq.Topic[modelangela.PullAppInfo],
) *libmq.Topic[modelangela.PullAccountAppInfoRelation] {
	return libmq.NewTopic[modelangela.PullAccountAppInfoRelation](
		"PullAccountAppInfoRelation",
		func(ctx context.Context, r *modelangela.PullAccountAppInfoRelation) error {
			if !a.supv.HasAccountPlatform(r.Platform) {
				return nil
			}
			var infoList []*librarian.AppInfo
			if resp, err := a.porter.PullAccountAppInfoRelation(
				a.supv.WithAccountPlatform(ctx, r.Platform),
				&porter.PullAccountAppInfoRelationRequest{
					RelationType: librarian.AccountAppRelationType_ACCOUNT_APP_RELATION_TYPE_OWN,
					AccountId: &librarian.AccountID{
						Platform:          r.Platform,
						PlatformAccountId: r.PlatformAccountID,
					},
				},
			); err != nil {
				return err
			} else {
				infoList = resp.GetAppInfos()
			}
			infoNum := len(infoList)
			if infoNum <= 0 {
				return nil
			}
			infos := make([]*modelgebura.AppInfo, 0, infoNum)
			var infoIDs []model.InternalID
			if id, err := a.id.BatchNew(infoNum); err != nil {
				return err
			} else {
				infoIDs = id
			}
			for i, info := range infoList {
				infos = append(infos, converter.ToBizAppInfo(info))
				infos[i].ID = infoIDs[i]
				infos[i].Source = r.Platform
			}
			if err := a.repo.UpsertAppInfos(ctx, infos); err != nil {
				return err
			}
			if err := a.repo.AccountPurchaseAppInfos(ctx, r.ID, infoIDs); err != nil {
				return err
			}
			for _, info := range infos {
				_ = sa.Publish(ctx, modelangela.PullAppInfo{
					ID: info.ID,
					AppInfoID: modelgebura.AppInfoID{
						Internal:    false,
						Source:      r.Platform,
						SourceAppID: info.SourceAppID,
					},
					IgnoreRateLimit: true,
				})
			}
			return nil
		},
	)
}
