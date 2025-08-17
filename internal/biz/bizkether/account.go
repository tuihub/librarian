package bizkether

import (
	"context"
	"encoding/json"
	"time"

	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelkether"
	"github.com/tuihub/librarian/internal/service/sephirah/converter"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
)

func NewPullAccountTopic(
	a *KetherBase,
	sr *libmq.Topic[modelkether.PullAccountAppInfoRelation],
) *libmq.Topic[model.PullAccountInfo] {
	return libmq.NewTopic[model.PullAccountInfo](
		"PullAccountInfo",
		func(ctx context.Context, info *model.PullAccountInfo) error {
			var configJSONObj model.PullAccountInfoConfig
			if err := json.Unmarshal([]byte(info.Config.ConfigJSON), &configJSONObj); err != nil {
				return err
			}

			if !a.supv.HasAccountPlatform(info.Config) {
				return nil
			}
			resp, err := a.porter.GetAccount(
				a.supv.WithAccountPlatform(ctx, configJSONObj.Platform),
				&porter.GetAccountRequest{
					Config: converter.ToPBFeatureRequest(info.Config),
				},
			)
			if err != nil {
				return err
			}
			err = a.repo.UpsertAccount(ctx, model.Account{
				ID:                info.ID,
				Platform:          configJSONObj.Platform,
				PlatformAccountID: configJSONObj.PlatformAccountID,
				Name:              resp.GetAccount().GetName(),
				ProfileURL:        resp.GetAccount().GetProfileUrl(),
				AvatarURL:         resp.GetAccount().GetAvatarUrl(),
				LatestUpdateTime:  time.Time{},
			})
			if err != nil {
				return err
			}
			return sr.
				Publish(ctx, modelkether.PullAccountAppInfoRelation{
					ID:                info.ID,
					Platform:          configJSONObj.Platform,
					PlatformAccountID: configJSONObj.PlatformAccountID,
				})
		},
	)
}

func NewPullAccountAppInfoRelationTopic(
	a *KetherBase,
	sa *libmq.Topic[modelkether.PullAppInfo],
) *libmq.Topic[modelkether.PullAccountAppInfoRelation] {
	return libmq.NewTopic[modelkether.PullAccountAppInfoRelation](
		"PullAccountAppInfoRelation",
		func(ctx context.Context, r *modelkether.PullAccountAppInfoRelation) error {
			// if !a.supv.HasAccountPlatform(r.Platform) {
			//	return nil
			//}
			// var infoList []*librarian.AppInfo
			// if resp, err := a.porter.PullAccountAppInfoRelation(
			//	a.supv.WithAccountPlatform(ctx, r.Platform),
			//	&porter.PullAccountAppInfoRelationRequest{
			//		RelationType: librarian.AccountAppRelationType_ACCOUNT_APP_RELATION_TYPE_OWN,
			//		AccountId: &librarian.AccountID{
			//			Platform:          r.Platform,
			//			PlatformAccountId: r.PlatformAccountID,
			//		},
			//	},
			// ); err != nil {
			//	return err
			// } else {
			//	infoList = resp.GetAppInfos()
			//}
			// infoNum := len(infoList)
			// if infoNum <= 0 {
			//	return nil
			//}
			// infos := make([]*modelgebura.AppInfo, 0, infoNum)
			// var infoIDs []model.InternalID
			// if id, err := a.id.BatchNew(infoNum); err != nil {
			//	return err
			// } else {
			//	infoIDs = id
			//}
			// for i, info := range infoList {
			//	infos = append(infos, converter.ToBizAppInfo(info))
			//	infos[i].ID = infoIDs[i]
			//	infos[i].Source = r.Platform
			//}
			// if err := a.repo.UpsertAppInfos(ctx, infos); err != nil {
			//	return err
			//}
			// if err := a.repo.AccountPurchaseAppInfos(ctx, r.ID, infoIDs); err != nil {
			//	return err
			//}
			// for _, info := range infos {
			//	_ = sa.Publish(ctx, modelkether.PullAppInfo{
			//		ID: info.ID,
			//		AppInfoID: modelgebura.AppInfoID{
			//			Internal:    false,
			//			Source:      r.Platform,
			//			SourceAppID: info.SourceAppID,
			//		},
			//		IgnoreRateLimit: true,
			//	})
			//}
			return nil
		},
	)
}
