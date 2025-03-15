package bizangela

import (
	"context"

	"github.com/tuihub/librarian/internal/biz/bizgebura"
	"github.com/tuihub/librarian/internal/lib/libidgenerator"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/libsearch"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelangela"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	"github.com/tuihub/librarian/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/model/modelsupervisor"
	"github.com/tuihub/librarian/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/internal/model/modelyesod"
	"github.com/tuihub/librarian/internal/service/supervisor"
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
	NewFeedItemPostprocessTopic,
	NewUpdateAppInfoIndexTopic,
)

type Angela struct {
	AngelaBase
	mq *libmq.MQ
}
type AngelaBase struct {
	repo   AngelaRepo
	supv   *supervisor.Supervisor
	g      bizgebura.GeburaRepo
	porter porter.LibrarianPorterServiceClient
	search libsearch.Search
	id     *libidgenerator.IDGenerator
}

type AngelaRepo interface {
	UpsertAccount(context.Context, modeltiphereth.Account) error
	UpsertAppInfo(context.Context, *modelgebura.AppInfo, *modelgebura.AppInfo) error
	UpsertAppInfos(context.Context, []*modelgebura.AppInfo) error
	AccountPurchaseAppInfos(context.Context, model.InternalID, []model.InternalID) error
	UpsertFeed(context.Context, *modelfeed.Feed) error
	CheckNewFeedItems(context.Context, []*modelfeed.Item, model.InternalID) ([]string, error)
	UpsertFeedItems(context.Context, []*modelfeed.Item, model.InternalID) error
	UpdateFeedPullStatus(context.Context, *modelyesod.FeedConfig) error
	GetFeedItem(context.Context, model.InternalID) (*modelfeed.Item, error)
	GetFeedActions(context.Context, model.InternalID) ([]*modelyesod.FeedActionSet, error)
	GetNotifyTargetItems(context.Context, model.InternalID, model.Paging) (*modelsupervisor.FeatureRequest, []*modelfeed.Item, error)
	AddFeedItemsToCollection(context.Context, model.InternalID, []model.InternalID) error
}

func NewAngelaBase(
	repo AngelaRepo,
	supv *supervisor.Supervisor,
	g bizgebura.GeburaRepo,
	pClient porter.LibrarianPorterServiceClient,
	search libsearch.Search,
	id *libidgenerator.IDGenerator,
) (*AngelaBase, error) {
	return &AngelaBase{
		repo:   repo,
		supv:   supv,
		g:      g,
		porter: pClient,
		search: search,
		id:     id,
	}, nil
}

func NewAngela(
	base *AngelaBase,
	mq *libmq.MQ,
	pullAccountInfo *libmq.Topic[modeltiphereth.PullAccountInfo],
	pullAccountAppInfoRelation *libmq.Topic[modelangela.PullAccountAppInfoRelation],
	pullAppInfo *libmq.Topic[modelangela.PullAppInfo],
	pullFeed *libmq.Topic[modelyesod.PullFeed],
	notifyRouter *libmq.Topic[modelangela.NotifyRouter],
	notifyPush *libmq.Topic[modelangela.NotifyPush],
	itemPostprocess *libmq.Topic[modelangela.FeedItemPostprocess],
	updateAppIndex *libmq.Topic[modelangela.UpdateAppInfoIndex],
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
	if err := mq.RegisterTopic(itemPostprocess); err != nil {
		return nil, err
	}
	if err := mq.RegisterTopic(updateAppIndex); err != nil {
		return nil, err
	}
	return &Angela{
		AngelaBase: *base,
		mq:         mq,
	}, nil
}
