package modelangela

import (
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
)

type PullSteamAccountAppRelation struct {
	ID      model.InternalID
	SteamID string
}

type PullSteamApp struct {
	ID    model.InternalID
	AppID string
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
