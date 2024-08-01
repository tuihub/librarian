package modeltiphereth

import (
	"github.com/tuihub/librarian/internal/model"
)

type PorterInstance struct {
	ID                model.InternalID
	Name              string
	Version           string
	GlobalName        string
	Address           string
	Region            string
	FeatureSummary    *PorterFeatureSummary
	Status            PorterInstanceStatus
	ConnectionStatus  PorterConnectionStatus
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
