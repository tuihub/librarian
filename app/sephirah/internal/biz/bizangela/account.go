package bizangela

import (
	"context"
	"strconv"
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/model/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelangela"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/model"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"
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

func NewPullSteamAccountAppRelationTopic( //nolint:gocognit // TODO
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
			internalApps := make([]*modelgebura.App, 0, appNum)
			var steamAppIDs, internalAppIDs []model.InternalID
			if resp, err := a.searcher.NewBatchIDs(ctx, &searcher.NewBatchIDsRequest{
				Num: int32(appNum),
			}); err != nil {
				return err
			} else {
				steamAppIDs = converter.ToBizInternalIDList(resp.GetIds())
			}
			if resp, err := a.searcher.NewBatchIDs(ctx, &searcher.NewBatchIDsRequest{
				Num: int32(appNum),
			}); err != nil {
				return err
			} else {
				internalAppIDs = converter.ToBizInternalIDList(resp.GetIds())
			}
			for i, app := range appList {
				internalApps = append(internalApps, converter.ToBizApp(app))
				internalApps[i].ID = internalAppIDs[i]
				internalApps[i].Source = modelgebura.AppSourceInternal
				internalApps[i].SourceAppID = strconv.FormatInt(int64(internalAppIDs[i]), 10)
				internalApps[i].BoundInternal = internalAppIDs[i]

				steamApps = append(steamApps, converter.ToBizApp(app))
				steamApps[i].ID = steamAppIDs[i]
				steamApps[i].Source = modelgebura.AppSourceSteam
				steamApps[i].BoundInternal = internalAppIDs[i]
			}
			vl := make([]*mapper.Vertex, len(steamApps)*2) //nolint:gomnd // double
			el := make([]*mapper.Edge, len(steamApps)*2)   //nolint:gomnd // double
			for i := range steamApps {
				vl[i*2] = &mapper.Vertex{
					Vid:  int64(internalApps[i].ID),
					Type: mapper.VertexType_VERTEX_TYPE_ABSTRACT,
					Prop: nil,
				}
				vl[i*2+1] = &mapper.Vertex{
					Vid:  int64(steamApps[i].ID),
					Type: mapper.VertexType_VERTEX_TYPE_METADATA,
					Prop: nil,
				}
				el[i*2] = &mapper.Edge{
					SrcVid: int64(internalApps[i].ID),
					DstVid: int64(steamApps[i].ID),
					Type:   mapper.EdgeType_EDGE_TYPE_EQUAL,
					Prop:   nil,
				}
				el[i*2+1] = &mapper.Edge{
					SrcVid: int64(r.ID),
					DstVid: int64(steamApps[i].ID),
					Type:   mapper.EdgeType_EDGE_TYPE_ENJOY,
					Prop:   nil,
				}
			}
			if _, err := a.mapper.InsertVertex(ctx, &mapper.InsertVertexRequest{VertexList: vl}); err != nil {
				return err
			}
			if _, err := a.mapper.InsertEdge(ctx, &mapper.InsertEdgeRequest{EdgeList: el}); err != nil {
				return err
			}
			if err := a.repo.UpsertApps(ctx, internalApps); err != nil {
				return err
			}
			if err := a.repo.UpsertApps(ctx, steamApps); err != nil {
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
