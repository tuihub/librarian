package bizkether

import (
	"github.com/tuihub/librarian/internal/data"
	"github.com/tuihub/librarian/internal/lib/libidgenerator"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/libsearch"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelkether"
	"github.com/tuihub/librarian/internal/model/modelyesod"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewKether,
	NewKetherBase,
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

type Kether struct {
	KetherBase
	mq *libmq.MQ
}
type KetherBase struct {
	repo   *data.KetherRepo
	supv   *data.SupervisorRepo
	g      *data.GeburaRepo
	porter porter.LibrarianPorterServiceClient
	search libsearch.Search
	id     *libidgenerator.IDGenerator
}

func NewKetherBase(
	repo *data.KetherRepo,
	supv *data.SupervisorRepo,
	g *data.GeburaRepo,
	pClient porter.LibrarianPorterServiceClient,
	search libsearch.Search,
	id *libidgenerator.IDGenerator,
) (*KetherBase, error) {
	return &KetherBase{
		repo:   repo,
		supv:   supv,
		g:      g,
		porter: pClient,
		search: search,
		id:     id,
	}, nil
}

func NewKether(
	base *KetherBase,
	mq *libmq.MQ,
	pullAccountInfo *libmq.Topic[model.PullAccountInfo],
	pullAccountAppInfoRelation *libmq.Topic[modelkether.PullAccountAppInfoRelation],
	pullAppInfo *libmq.Topic[modelkether.PullAppInfo],
	pullFeed *libmq.Topic[modelyesod.PullFeed],
	notifyRouter *libmq.Topic[modelkether.NotifyRouter],
	notifyPush *libmq.Topic[modelkether.NotifyPush],
	itemPostprocess *libmq.Topic[modelkether.FeedItemPostprocess],
	updateAppIndex *libmq.Topic[modelkether.UpdateAppInfoIndex],
) (*Kether, error) {
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
	return &Kether{
		KetherBase: *base,
		mq:         mq,
	}, nil
}
