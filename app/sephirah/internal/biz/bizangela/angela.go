package bizangela

import (
	"context"
	"errors"
	"net/url"
	"strconv"
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizyesod"
	"github.com/tuihub/librarian/app/sephirah/internal/model/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelangela"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"github.com/google/wire"
	"golang.org/x/exp/slices"
)

var ProviderSet = wire.NewSet(
	NewAngela,
	NewAngelaBase,
	NewPullAccountTopic,
	NewPullSteamAccountAppRelationTopic,
	NewPullSteamAppTopic,
	NewPullFeedTopic,
	NewNotifyRouterTopic,
	NewNotifyPushTopic,
	NewFeedToNotifyFlowMap,
	NewNotifyFlowCache,
	NewNotifyTargetCache,
)

type Angela struct {
	mq *libmq.MQ
}
type AngelaBase struct {
	t        biztiphereth.TipherethRepo
	g        bizgebura.GeburaRepo
	y        bizyesod.YesodRepo
	mapper   mapper.LibrarianMapperServiceClient
	searcher searcher.LibrarianSearcherServiceClient
	porter   porter.LibrarianPorterServiceClient
}

func NewAngelaBase(
	t biztiphereth.TipherethRepo,
	g bizgebura.GeburaRepo,
	y bizyesod.YesodRepo,
	mClient mapper.LibrarianMapperServiceClient,
	pClient porter.LibrarianPorterServiceClient,
	sClient searcher.LibrarianSearcherServiceClient,
) (*AngelaBase, error) {
	return &AngelaBase{
		t:        t,
		g:        g,
		y:        y,
		mapper:   mClient,
		porter:   pClient,
		searcher: sClient,
	}, nil
}

func NewAngela(
	mq *libmq.MQ,
	pullAccount *libmq.Topic[modeltiphereth.PullAccountInfo],
	pullSteamAccountAppRelation *libmq.Topic[modelangela.PullSteamAccountAppRelation],
	pullSteamApp *libmq.Topic[modelangela.PullSteamApp],
	pullFeed *libmq.Topic[modelyesod.PullFeed],
	notifyRouter *libmq.Topic[modelangela.NotifyRouter],
	notifyPush *libmq.Topic[modelangela.NotifyPush],
) (*Angela, error) {
	if err := mq.RegisterTopic(pullAccount); err != nil {
		return nil, err
	}
	if err := mq.RegisterTopic(pullSteamAccountAppRelation); err != nil {
		return nil, err
	}
	if err := mq.RegisterTopic(pullSteamApp); err != nil {
		return nil, err
	}
	if err := mq.RegisterTopic(pullFeed); err != nil {
		return nil, err
	}
	if err := mq.RegisterTopic(notifyRouter); err != nil {
		return nil, err
	}
	if err := mq.RegisterTopic(notifyPush); err != nil {
		return nil, err
	}
	return &Angela{
		mq: mq,
	}, nil
}

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
				err = a.t.UpdateAccount(ctx, modeltiphereth.Account{
					ID:                info.ID,
					Platform:          info.Platform,
					PlatformAccountID: info.PlatformAccountID,
					Name:              resp.GetAccount().GetName(),
					ProfileURL:        resp.GetAccount().GetProfileUrl(),
					AvatarURL:         resp.GetAccount().GetAvatarUrl(),
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

func NewPullSteamAccountAppRelationTopic( //nolint:funlen,gocognit // TODO
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
				internalApps = append(internalApps, &modelgebura.App{
					ID:               internalAppIDs[i],
					Source:           modelgebura.AppSourceInternal,
					SourceAppID:      strconv.FormatInt(int64(internalAppIDs[i]), 10),
					SourceURL:        "",
					Name:             app.GetName(),
					Type:             modelgebura.AppTypeGame,
					ShortDescription: "",
					ImageURL:         "",
					Details:          nil,
					BoundInternal:    internalAppIDs[i],
				})
				steamApps = append(steamApps, &modelgebura.App{
					ID:               steamAppIDs[i],
					Source:           modelgebura.AppSourceSteam,
					SourceAppID:      app.GetSourceAppId(),
					SourceURL:        "",
					Name:             app.GetName(),
					Type:             modelgebura.AppTypeGame,
					ShortDescription: "",
					ImageURL:         "",
					Details:          nil,
					BoundInternal:    internalAppIDs[i],
				})
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
			if err := a.g.UpsertApps(ctx, internalApps); err != nil {
				return err
			}
			if err := a.g.UpsertApps(ctx, steamApps); err != nil {
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

func NewPullSteamAppTopic(
	a *AngelaBase,
) *libmq.Topic[modelangela.PullSteamApp] {
	return libmq.NewTopic[modelangela.PullSteamApp](
		"PullSteamApp",
		func(ctx context.Context, r *modelangela.PullSteamApp) error {
			ctx = libapp.NewContext(ctx, string(porter.FeatureFlag_FEATURE_FLAG_SOURCE_STEAM))
			resp, err := a.porter.PullApp(ctx, &porter.PullAppRequest{AppId: &librarian.AppID{
				Source:      librarian.AppSource_APP_SOURCE_STEAM,
				SourceAppId: r.AppID,
			}})
			if err != nil {
				return err
			}
			app := converter.ToBizApp(resp.GetApp())
			app.ID = r.ID
			app.Source = modelgebura.AppSourceSteam
			app.Type = modelgebura.AppTypeGame
			err = a.g.UpdateApp(ctx, app)
			if err != nil {
				return err
			}
			return nil
		},
	)
}

func NewPullFeedTopic( //nolint:gocognit // TODO
	a *AngelaBase,
	notify *libmq.Topic[modelangela.NotifyRouter],
) *libmq.Topic[modelyesod.PullFeed] {
	return libmq.NewTopic[modelyesod.PullFeed](
		"PullFeed",
		func(ctx context.Context, p *modelyesod.PullFeed) error {
			resp, err := a.porter.PullFeed(ctx, &porter.PullFeedRequest{
				Source:    porter.FeedSource_FEED_SOURCE_COMMON,
				ChannelId: p.URL,
			})
			if err != nil {
				return err
			}
			feed := modelfeed.NewConverter().FromPBFeed(resp.GetData())
			feed.ID = p.InternalID
			err = a.y.UpsertFeed(ctx, feed)
			if err != nil {
				return err
			}
			for _, item := range feed.Items {
				// generate internal_id
				var res *searcher.NewIDResponse
				res, err = a.searcher.NewID(ctx, &searcher.NewIDRequest{})
				if err != nil {
					return err
				}
				item.ID = converter.ToBizInternalID(res.GetId())
				// generate publish_platform
				if len(item.Link) > 0 {
					var linkParsed *url.URL
					linkParsed, err = url.Parse(item.Link)
					if err != nil {
						continue
					}
					item.PublishPlatform = linkParsed.Host
				}
				// generate published_parsed
				if item.PublishedParsed == nil {
					t := time.Now()
					item.PublishedParsed = &t
				}
			}
			newItemGUIDs, err := a.y.UpsertFeedItems(ctx, feed.Items, feed.ID)
			if err != nil {
				return err
			}
			newItems := make([]*modelfeed.Item, 0, len(newItemGUIDs))
			for _, item := range feed.Items {
				if slices.Contains(newItemGUIDs, item.GUID) {
					newItems = append(newItems, item)
				}
			}
			err = notify.Publish(ctx, modelangela.NotifyRouter{
				FeedID:   feed.ID,
				Messages: newItems,
			})
			if err != nil {
				return err
			}
			return nil
		},
	)
}

func NewNotifyRouterTopic(
	a *AngelaBase,
	flowMap *libcache.Map[model.InternalID, modelnetzach.NotifyFlow],
	feedToFlowMap *libcache.Map[model.InternalID, modelangela.FeedToNotifyFlowValue],
	push *libmq.Topic[modelangela.NotifyPush],
) *libmq.Topic[modelangela.NotifyRouter] {
	return libmq.NewTopic[modelangela.NotifyRouter](
		"NotifyRouter",
		func(ctx context.Context, r *modelangela.NotifyRouter) error {
			flowIDs, err := feedToFlowMap.GetWithFallBack(ctx, r.FeedID, nil)
			if err != nil {
				return err
			}
			if flowIDs == nil {
				return errors.New("nil result from feedToFlowMap")
			}
			for _, flowID := range *flowIDs {
				var flow *modelnetzach.NotifyFlow
				flow, err = flowMap.GetWithFallBack(ctx, flowID, nil)
				if err != nil {
					return err
				}
				for _, target := range flow.Targets {
					if target == nil {
						continue
					}
					err = push.Publish(ctx, modelangela.NotifyPush{
						Target:   *target,
						Messages: r.Messages,
					})
					if err != nil {
						return err
					}
				}
			}
			return nil
		},
	)
}

func NewNotifyPushTopic(
	a *AngelaBase,
	targetMap *libcache.Map[model.InternalID, modelnetzach.NotifyTarget],
) *libmq.Topic[modelangela.NotifyPush] {
	return libmq.NewTopic[modelangela.NotifyPush](
		"NotifyPush",
		func(ctx context.Context, p *modelangela.NotifyPush) error {
			target, err := targetMap.GetWithFallBack(ctx, p.Target.TargetID, nil)
			if err != nil {
				return err
			}
			_, err = a.porter.PushFeedItems(ctx, &porter.PushFeedItemsRequest{
				Destination: converter.ToPBFeedDestination(target.Type),
				ChannelId:   p.Target.ChannelID,
				Items:       converter.ToPBFeedItemList(p.Messages),
				Token:       target.Token,
			})
			if err != nil {
				return err
			}
			return nil
		},
	)
}
