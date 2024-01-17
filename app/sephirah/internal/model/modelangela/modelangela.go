package modelangela

import (
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
)

type PullAccountAppRelation struct {
	ID                model.InternalID
	Platform          string
	PlatformAccountID string
}

type PullApp struct {
	ID              model.InternalID
	AppID           modelgebura.AppID
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

type ParseFeedItemDigest struct {
	ID model.InternalID
}

type UpdateAppIndex struct {
	IDs []model.InternalID
}
