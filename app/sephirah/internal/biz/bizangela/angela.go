package bizangela

import (
	"context"
	"strconv"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/internal/lib/libmq"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewAngela,
	NewAngelaBase,
	NewPullAccountTopic,
	NewPullSteamAccountAppRelationTopic,
	NewPullSteamAppTopic,
)

type Angela struct {
	mq *libmq.MQ
}
type AngelaBase struct {
	t        biztiphereth.TipherethRepo
	g        bizgebura.GeburaRepo
	mapper   mapper.LibrarianMapperServiceClient
	searcher searcher.LibrarianSearcherServiceClient
	porter   porter.LibrarianPorterServiceClient
}

func NewAngelaBase(
	t biztiphereth.TipherethRepo,
	g bizgebura.GeburaRepo,
	mClient mapper.LibrarianMapperServiceClient,
	pClient porter.LibrarianPorterServiceClient,
	sClient searcher.LibrarianSearcherServiceClient,
) (*AngelaBase, error) {
	return &AngelaBase{
		t:        t,
		g:        g,
		mapper:   mClient,
		porter:   pClient,
		searcher: sClient,
	}, nil
}

func NewAngela(
	mq *libmq.MQ,
	pullAccount *libmq.TopicImpl[biztiphereth.PullAccountInfo],
	pullSteamAccountAppRelation *libmq.TopicImpl[PullSteamAccountAppRelation],
	pullSteamApp *libmq.TopicImpl[PullSteamApp],
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
	return &Angela{
		mq: mq,
	}, nil
}

func NewPullAccountTopic(
	a *AngelaBase,
	sr *libmq.TopicImpl[PullSteamAccountAppRelation],
) *libmq.TopicImpl[biztiphereth.PullAccountInfo] {
	return libmq.NewTopic[biztiphereth.PullAccountInfo](
		"PullAccountInfo",
		func() biztiphereth.PullAccountInfo {
			return biztiphereth.PullAccountInfo{
				InternalID:        0,
				Platform:          0,
				PlatformAccountID: "",
			}
		},
		func(ctx context.Context, info biztiphereth.PullAccountInfo) error {
			resp, err := a.porter.PullAccount(ctx, &porter.PullAccountRequest{AccountId: &librarian.AccountID{
				Platform:          biztiphereth.ToPBAccountPlatform(info.Platform),
				PlatformAccountId: info.PlatformAccountID,
			}})
			if err != nil {
				return err
			}
			switch info.Platform {
			case biztiphereth.AccountPlatformSteam:
				err = a.t.UpdateAccount(ctx, biztiphereth.Account{
					InternalID:        info.InternalID,
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
					Publish(ctx, PullSteamAccountAppRelation{
						InternalID: info.InternalID,
						SteamID:    info.PlatformAccountID,
					})
			default:
				return nil
			}
		},
	)
}

func NewPullSteamAccountAppRelationTopic(
	a *AngelaBase,
	sa *libmq.TopicImpl[PullSteamApp],
) *libmq.TopicImpl[PullSteamAccountAppRelation] {
	return libmq.NewTopic[PullSteamAccountAppRelation](
		"PullSteamAccountAppRelation",
		func() PullSteamAccountAppRelation {
			return PullSteamAccountAppRelation{
				InternalID: 0,
				SteamID:    "",
			}
		},
		func(ctx context.Context, r PullSteamAccountAppRelation) error {
			resp, err := a.porter.PullAccountAppRelation(ctx, &porter.PullAccountAppRelationRequest{
				RelationType: porter.AccountAppRelationType_ACCOUNT_APP_RELATION_TYPE_OWN,
				AccountId: &librarian.AccountID{
					Platform:          librarian.AccountPlatform_ACCOUNT_PLATFORM_STEAM,
					PlatformAccountId: r.SteamID,
				},
			})
			if err != nil {
				return err
			}
			steamApps := make([]*bizgebura.App, 0, len(resp.GetAppList()))
			internalApps := make([]*bizgebura.App, len(resp.GetAppList()))
			for i, app := range resp.GetAppList() {
				resp2, err2 := a.searcher.NewID(ctx, &searcher.NewIDRequest{})
				if err2 != nil {
					return err2
				}
				internalApps[i] = &bizgebura.App{ // TODO
					InternalID:      resp2.Id,
					Source:          bizgebura.AppSourceInternal,
					SourceAppID:     strconv.FormatInt(resp2.Id, 10),
					SourceURL:       "",
					Name:            app.GetName(),
					Type:            0,
					ShorDescription: "",
					ImageURL:        "",
					Details:         nil,
				}
				resp2, err2 = a.searcher.NewID(ctx, &searcher.NewIDRequest{})
				if err2 != nil {
					return err2
				}
				steamApps[i] = &bizgebura.App{ // TODO
					InternalID:      resp2.Id,
					Source:          bizgebura.AppSourceSteam,
					SourceAppID:     app.GetSourceAppId(),
					SourceURL:       "",
					Name:            app.GetName(),
					Type:            0,
					ShorDescription: "",
					ImageURL:        "",
					Details:         nil,
				}
			}
			vl := make([]*mapper.Vertex, len(steamApps)*2) //nolint:gomnd // double
			el := make([]*mapper.Edge, len(steamApps)*2)   //nolint:gomnd // double
			for i := range steamApps {
				vl[i*2] = &mapper.Vertex{
					Vid:  internalApps[i].InternalID,
					Type: mapper.VertexType_VERTEX_TYPE_ABSTRACT,
					Prop: nil,
				}
				vl[i*2+1] = &mapper.Vertex{
					Vid:  steamApps[i].InternalID,
					Type: mapper.VertexType_VERTEX_TYPE_METADATA,
					Prop: nil,
				}
				el[i*2] = &mapper.Edge{
					SrcVid: internalApps[i].InternalID,
					DstVid: steamApps[i].InternalID,
					Type:   mapper.EdgeType_EDGE_TYPE_EQUAL,
					Prop:   nil,
				}
				el[i*2+1] = &mapper.Edge{
					SrcVid: r.InternalID,
					DstVid: steamApps[i].InternalID,
					Type:   mapper.EdgeType_EDGE_TYPE_ENJOY,
					Prop:   nil,
				}
			}
			if _, err = a.mapper.InsertVertex(ctx, &mapper.InsertVertexRequest{VertexList: vl}); err != nil {
				return err
			}
			if _, err = a.mapper.InsertEdge(ctx, &mapper.InsertEdgeRequest{EdgeList: el}); err != nil {
				return err
			}
			if err = a.g.UpsertApp(ctx, append(steamApps, internalApps...)); err != nil {
				return err
			}
			for _, app := range steamApps {
				_ = sa.Publish(ctx, PullSteamApp{
					InternalID: app.InternalID,
					AppID:      app.SourceAppID,
				})
			}
			return nil
		},
	)
}

func NewPullSteamAppTopic(
	a *AngelaBase,
) *libmq.TopicImpl[PullSteamApp] {
	return libmq.NewTopic[PullSteamApp](
		"PullSteamApp",
		func() PullSteamApp {
			return PullSteamApp{
				InternalID: 0,
				AppID:      "",
			}
		},
		func(ctx context.Context, r PullSteamApp) error {
			resp, err := a.porter.PullApp(ctx, &porter.PullAppRequest{AppId: &librarian.AppID{
				Source:      librarian.AppSource_APP_SOURCE_STEAM,
				SourceAppId: r.AppID,
			}})
			if err != nil {
				return err
			}
			app := resp.GetApp()
			err = a.g.UpdateApp(ctx, &bizgebura.App{
				InternalID:      r.InternalID,
				Source:          bizgebura.AppSourceSteam,
				SourceAppID:     app.GetSourceAppId(),
				SourceURL:       app.GetSourceUrl(),
				Name:            app.GetName(),
				Type:            bizgebura.ToBizAppType(app.GetType()),
				ShorDescription: app.GetShortDescription(),
				ImageURL:        app.GetImageUrl(),
				Details: &bizgebura.AppDetails{
					Description: app.GetDetails().GetDescription(),
					ReleaseDate: app.GetDetails().GetReleaseDate(),
					Developer:   app.GetDetails().GetDeveloper(),
					Publisher:   app.GetDetails().GetPublisher(),
				},
			})
			if err != nil {
				return err
			}
			return nil
		},
	)
}
