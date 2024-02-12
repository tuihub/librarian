package bizangela

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/client"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelangela"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	"github.com/tuihub/librarian/app/sephirah/internal/supervisor"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewAngela,
	NewAngelaBase,
	NewPullAccountTopic,
	NewPullAccountAppRelationTopic,
	NewPullAppTopic,
	NewAppCache,
	NewPullFeedTopic,
	NewNotifyRouterTopic,
	NewNotifyPushTopic,
	NewFeedToNotifyFlowCache,
	NewNotifyFlowCache,
	NewNotifyTargetCache,
	NewParseFeedItemDigestTopic,
	NewUpdateAppIndexTopic,
)

type Angela struct {
	mq *libmq.MQ
}
type AngelaBase struct {
	repo AngelaRepo
	supv *supervisor.Supervisor
	g    bizgebura.GeburaRepo
	// mapper   mapper.LibrarianMapperServiceClient
	searcher *client.Searcher
	porter   porter.LibrarianPorterServiceClient
}

type AngelaRepo interface {
	UpdateAccount(context.Context, modeltiphereth.Account) error
	UpsertApp(context.Context, *modelgebura.App, *modelgebura.App) error
	UpsertApps(context.Context, []*modelgebura.App) error
	AccountPurchaseApps(context.Context, model.InternalID, []model.InternalID) error
	UpsertFeed(context.Context, *modelfeed.Feed) error
	UpsertFeedItems(context.Context, []*modelfeed.Item, model.InternalID) ([]string, error)
	GetFeedItem(context.Context, model.InternalID) (*modelfeed.Item, error)
	UpdateFeedItemDigest(context.Context, *modelfeed.Item) error
}

func NewAngelaBase(
	repo AngelaRepo,
	supv *supervisor.Supervisor,
	g bizgebura.GeburaRepo,
	// mClient mapper.LibrarianMapperServiceClient,
	pClient porter.LibrarianPorterServiceClient,
	sClient *client.Searcher,
) (*AngelaBase, error) {
	return &AngelaBase{
		repo: repo,
		supv: supv,
		g:    g,
		//mapper:   mClient,
		porter:   pClient,
		searcher: sClient,
	}, nil
}

func NewAngela(
	mq *libmq.MQ,
	pullAccount *libmq.Topic[modeltiphereth.PullAccountInfo],
	pullSteamAccountAppRelation *libmq.Topic[modelangela.PullAccountAppRelation],
	pullSteamApp *libmq.Topic[modelangela.PullApp],
	pullFeed *libmq.Topic[modelyesod.PullFeed],
	notifyRouter *libmq.Topic[modelangela.NotifyRouter],
	notifyPush *libmq.Topic[modelangela.NotifyPush],
	parseFeedItem *libmq.Topic[modelangela.ParseFeedItemDigest],
	updateAppIndex *libmq.Topic[modelangela.UpdateAppIndex],
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
	if err := mq.RegisterTopic(parseFeedItem); err != nil {
		return nil, err
	}
	if err := mq.RegisterTopic(updateAppIndex); err != nil {
		return nil, err
	}
	return &Angela{
		mq: mq,
	}, nil
}
