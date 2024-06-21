package modeltiphereth

import (
	"github.com/tuihub/librarian/internal/model"
)

type PorterInstance struct {
	ID               model.InternalID
	Name             string
	Version          string
	GlobalName       string
	Address          string
	FeatureSummary   *PorterFeatureSummary
	Status           PorterInstanceStatus
	ConnectionStatus PorterConnectionStatus
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
	Region           string `json:"region"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	ConfigJSONSchema string `json:"config_json_schema"`
}

type FeatureRequest struct {
	ID         string `json:"id"`
	Region     string `json:"region"`
	ConfigJSON string `json:"config_json"`
}

type PorterInstancePrivilege struct {
	All bool `json:"all"`
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
