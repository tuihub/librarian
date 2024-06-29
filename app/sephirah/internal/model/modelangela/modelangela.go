package modelangela

import (
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
)

type PullAccountAppInfoRelation struct {
	ID                model.InternalID
	Platform          string
	PlatformAccountID string
}

type PullAppInfo struct {
	ID              model.InternalID
	AppInfoID       modelgebura.AppInfoID
	IgnoreRateLimit bool
}

type NotifyRouter struct {
	FeedID   model.InternalID
	Messages []*modelfeed.Item
}

type FeedToNotifyFlowValue []model.InternalID

type NotifyPush struct {
	Target   modelnetzach.NotifyFlowTarget
	Messages []*modelfeed.Item
}

type FeedItemPostprocess struct {
	FeedID       model.InternalID
	ItemID       model.InternalID
	SystemNotify *modelnetzach.SystemNotify
}

type UpdateAppInfoIndex struct {
	IDs []model.InternalID
}
