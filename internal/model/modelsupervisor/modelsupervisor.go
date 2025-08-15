package modelsupervisor

import (
	"github.com/tuihub/librarian/internal/lib/libtype"
	"github.com/tuihub/librarian/internal/model"
)

type PorterInstanceController struct {
	PorterInstance
}

type PorterContextController struct {
	PorterContext
}

type PorterInstance struct {
	ID                      model.InternalID
	BinarySummary           *PorterBinarySummary
	GlobalName              string
	Address                 string
	Region                  string
	FeatureSummary          *PorterFeatureSummary
	Status                  model.UserStatus
	ContextJSONSchema       string
	ConnectionStatus        PorterConnectionStatus
	ConnectionStatusMessage string
}

type PorterBinarySummary struct {
	Name              string
	Version           string
	Description       string
	SourceCodeAddress string
	BuildVersion      string
	BuildDate         string
}

type PorterFeatureSummary struct {
	AccountPlatforms   []*model.FeatureFlag `json:"account_platforms"`
	AppInfoSources     []*model.FeatureFlag `json:"app_info_sources"`
	FeedSources        []*model.FeatureFlag `json:"feed_sources"`
	NotifyDestinations []*model.FeatureFlag `json:"notify_destinations"`
	FeedItemActions    []*model.FeatureFlag `json:"feed_item_actions"`
	FeedGetters        []*model.FeatureFlag `json:"feed_getters"`
	FeedSetters        []*model.FeatureFlag `json:"feed_setters"`
}

type ServerFeatureSummary struct {
	AccountPlatforms   []*model.FeatureFlag
	AppInfoSources     []*model.FeatureFlag
	FeedSources        []*model.FeatureFlag
	NotifyDestinations []*model.FeatureFlag
	FeedItemActions    []*model.FeatureFlag
	FeedSetters        []*model.FeatureFlag
	FeedGetters        []*model.FeatureFlag
}

type PorterConnectionStatus int

const (
	PorterConnectionStatusUnspecified PorterConnectionStatus = iota
	PorterConnectionStatusQueueing
	PorterConnectionStatusConnected
	PorterConnectionStatusDisconnected
	PorterConnectionStatusActive
	PorterConnectionStatusActivationFailed
	PorterConnectionStatusDowngraded
)

type PorterContext struct {
	ID                  model.InternalID
	GlobalName          string
	Region              string
	ContextJSON         string
	Name                string
	Description         string
	Status              PorterContextStatus
	HandleStatus        PorterContextHandleStatus
	HandleStatusMessage string
}

type PorterContextStatus int

const (
	PorterContextStatusUnspecified PorterContextStatus = iota
	PorterContextStatusActive
	PorterContextStatusDisabled
)

type PorterContextHandleStatus int

const (
	PorterContextHandleStatusUnspecified PorterContextHandleStatus = iota
	PorterContextHandleStatusActive
	PorterContextHandleStatusDowngraded
	PorterContextHandleStatusQueueing
	PorterContextHandleStatusBlocked
)

type PorterDigest struct {
	BinarySummary     *PorterBinarySummary
	GlobalName        string
	Regions           []string
	ContextJSONSchema string
	FeatureSummary    *PorterFeatureSummary
}

type ServerFeatureSummaryMap struct {
	AccountPlatforms   *libtype.SyncMap[string, []string]
	AppInfoSources     *libtype.SyncMap[string, []string]
	FeedSources        *libtype.SyncMap[string, []string]
	NotifyDestinations *libtype.SyncMap[string, []string]
	FeedItemActions    *libtype.SyncMap[string, []string]
}

func NewServerFeatureSummaryMap() *ServerFeatureSummaryMap {
	return &ServerFeatureSummaryMap{
		AccountPlatforms:   libtype.NewSyncMap[string, []string](),
		AppInfoSources:     libtype.NewSyncMap[string, []string](),
		FeedSources:        libtype.NewSyncMap[string, []string](),
		NotifyDestinations: libtype.NewSyncMap[string, []string](),
		FeedItemActions:    libtype.NewSyncMap[string, []string](),
	}
}
