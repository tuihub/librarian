package modeltiphereth

import (
	"time"

	"github.com/tuihub/librarian/internal/lib/libtype"
	"github.com/tuihub/librarian/internal/model"
)

type PorterInstanceController struct {
	PorterInstance
	ConnectionStatus PorterConnectionStatus
	LastHeartbeat    time.Time
}

type PorterInstance struct {
	ID                model.InternalID
	Name              string
	Version           string
	GlobalName        string
	Address           string
	Region            string
	FeatureSummary    *PorterFeatureSummary
	Status            PorterInstanceStatus
	ContextJSONSchema string
}

type PorterFeatureSummary struct {
	AccountPlatforms   []*FeatureFlag `json:"account_platforms"`
	AppInfoSources     []*FeatureFlag `json:"app_info_sources"`
	FeedSources        []*FeatureFlag `json:"feed_sources"`
	NotifyDestinations []*FeatureFlag `json:"notify_destinations"`
	FeedItemActions    []*FeatureFlag `json:"feed_item_actions"`
}

type FeatureFlag struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	ConfigJSONSchema string `json:"config_json_schema"`
	RequireContext   bool   `json:"require_context"`
}

func (f *FeatureFlag) Match(request *FeatureRequest) bool {
	return f.ID == request.ID
}

type FeatureRequest struct {
	ID         string           `json:"id"`
	Region     string           `json:"region"`
	ConfigJSON string           `json:"config_json"`
	ContextID  model.InternalID `json:"context_id"`
}

type PorterInstanceContext struct {
	ID          model.InternalID
	PorterID    model.InternalID
	ContextJSON string
}

type PorterInstanceStatus int

const (
	PorterInstanceStatusUnspecified PorterInstanceStatus = iota
	PorterInstanceStatusActive
	PorterInstanceStatusBlocked
)

type ServerFeatureSummary struct {
	AccountPlatforms   []*FeatureFlag
	AppInfoSources     []*FeatureFlag
	FeedSources        []*FeatureFlag
	NotifyDestinations []*FeatureFlag
	FeedItemActions    []*FeatureFlag
}

type PorterConnectionStatus int

const (
	PorterConnectionStatusUnspecified PorterConnectionStatus = iota
	PorterConnectionStatusConnected
	PorterConnectionStatusDisconnected
	PorterConnectionStatusActive
	PorterConnectionStatusActivationFailed
)

type ServerFeatureSummaryMap struct {
	AccountPlatforms   *libtype.SyncMap[[]string]
	AppInfoSources     *libtype.SyncMap[[]string]
	FeedSources        *libtype.SyncMap[[]string]
	NotifyDestinations *libtype.SyncMap[[]string]
	FeedItemActions    *libtype.SyncMap[[]string]
}

func NewServerFeatureSummaryMap() *ServerFeatureSummaryMap {
	return &ServerFeatureSummaryMap{
		AccountPlatforms:   libtype.NewSyncMap[[]string](),
		AppInfoSources:     libtype.NewSyncMap[[]string](),
		FeedSources:        libtype.NewSyncMap[[]string](),
		NotifyDestinations: libtype.NewSyncMap[[]string](),
		FeedItemActions:    libtype.NewSyncMap[[]string](),
	}
}
