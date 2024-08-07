package modelsupervisor

import (
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/internal/lib/libtype"
	"github.com/tuihub/librarian/internal/model"
)

type PorterInstanceController struct {
	PorterInstance
	ConnectionStatus        PorterConnectionStatus
	ConnectionStatusMessage string
	LastHeartbeat           time.Time
}

type PorterInstance struct {
	ID                model.InternalID
	BinarySummary     *PorterBinarySummary
	GlobalName        string
	Address           string
	Region            string
	FeatureSummary    *PorterFeatureSummary
	Status            modeltiphereth.UserStatus
	ContextJSONSchema string
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

type PorterGroup struct {
	BinarySummary     *PorterBinarySummary
	GlobalName        string
	Regions           []string
	ContextJSONSchema string
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
