package bizangela

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/client"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelangela"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
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
	NewPullAccountAppInfoRelationTopic,
	NewPullAppInfoTopic,
	NewAppInfoCache,
	NewPullFeedTopic,
	NewNotifyRouterTopic,
	NewNotifyPushTopic,
	NewFeedToNotifyFlowCache,
	NewNotifyFlowCache,
	NewNotifyTargetCache,
	NewParseFeedItemDigestTopic,
	NewUpdateAppInfoIndexTopic,
	NewSystemNotificationTopic,
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
	UpsertAccount(context.Context, modeltiphereth.Account) error
	UpsertAppInfo(context.Context, *modelgebura.AppInfo, *modelgebura.AppInfo) error
	UpsertAppInfos(context.Context, []*modelgebura.AppInfo) error
	AccountPurchaseAppInfos(context.Context, model.InternalID, []model.InternalID) error
	UpsertFeed(context.Context, *modelfeed.Feed) error
	UpsertFeedItems(context.Context, []*modelfeed.Item, model.InternalID) ([]string, error)
	UpdateFeedPullStatus(context.Context, *modelyesod.FeedConfig) error
	GetFeedItem(context.Context, model.InternalID) (*modelfeed.Item, error)
	UpdateFeedItemDigest(context.Context, *modelfeed.Item) error
	UpsertSystemNotification(context.Context, model.InternalID, *modelnetzach.SystemNotification) error
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
	pullAccountInfo *libmq.Topic[modeltiphereth.PullAccountInfo],
	pullAccountAppInfoRelation *libmq.Topic[modelangela.PullAccountAppInfoRelation],
	pullAppInfo *libmq.Topic[modelangela.PullAppInfo],
	pullFeed *libmq.Topic[modelyesod.PullFeed],
	notifyRouter *libmq.Topic[modelangela.NotifyRouter],
	notifyPush *libmq.Topic[modelangela.NotifyPush],
	parseFeedItem *libmq.Topic[modelangela.ParseFeedItemDigest],
	updateAppIndex *libmq.Topic[modelangela.UpdateAppInfoIndex],
	systemNotification *libmq.Topic[modelangela.SystemNotify],
) (*Angela, error) {
	if err := mq.RegisterTopic(pullAccountInfo); err != nil {
		return nil, err
	}
	if err := mq.RegisterTopic(pullAccountAppInfoRelation); err != nil {
		return nil, err
	}
	if err := mq.RegisterTopic(pullAppInfo); err != nil {
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
	if err := mq.RegisterTopic(systemNotification); err != nil {
		return nil, err
	}
	return &Angela{
		mq: mq,
	}, nil
}
