package modelsupervisor

import (
	"database/sql/driver"
	"errors"
	"time"

	"github.com/tuihub/librarian/internal/lib/libtype"
	"github.com/tuihub/librarian/internal/model"

	"gorm.io/gorm"
)

type PorterInstanceController struct {
	PorterInstance
}

type PorterContextController struct {
	PorterContext
}

type PorterInstance struct {
	ID                      model.InternalID      `gorm:"primaryKey;autoIncrement:false"`
	BinarySummary           *PorterBinarySummary  `gorm:"-"` // Handled by hooks
	GlobalName              string                `gorm:"index:idx_porter_instance_global_name_region,priority:1"`
	Address                 string                `gorm:"uniqueIndex"`
	Region                  string                `gorm:"index:idx_porter_instance_global_name_region,priority:2"`
	FeatureSummary          *PorterFeatureSummary `gorm:"serializer:json"`
	Status                  model.UserStatus
	ContextJSONSchema       string
	ConnectionStatus        PorterConnectionStatus
	ConnectionStatusMessage string
	UpdatedAt               time.Time
	CreatedAt               time.Time

	// Flattened fields for DB
	Name              string `gorm:"column:name"`
	Version           string `gorm:"column:version"`
	Description       string `gorm:"column:description"`
	SourceCodeAddress string `gorm:"column:source_code_address"`
	BuildVersion      string `gorm:"column:build_version"`
	BuildDate         string `gorm:"column:build_date"`
}

func (PorterInstance) TableName() string {
	return "porter_instances"
}

func (p *PorterInstance) BeforeSave(tx *gorm.DB) error {
	if p.BinarySummary != nil {
		p.Name = p.BinarySummary.Name
		p.Version = p.BinarySummary.Version
		p.Description = p.BinarySummary.Description
		p.SourceCodeAddress = p.BinarySummary.SourceCodeAddress
		p.BuildVersion = p.BinarySummary.BuildVersion
		p.BuildDate = p.BinarySummary.BuildDate
	}
	return nil
}

func (p *PorterInstance) AfterFind(tx *gorm.DB) error {
	p.BinarySummary = &PorterBinarySummary{
		Name:              p.Name,
		Version:           p.Version,
		Description:       p.Description,
		SourceCodeAddress: p.SourceCodeAddress,
		BuildVersion:      p.BuildVersion,
		BuildDate:         p.BuildDate,
	}
	return nil
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

func (s PorterConnectionStatus) Value() (driver.Value, error) {
	switch s {
	case PorterConnectionStatusQueueing:
		return "queueing", nil
	case PorterConnectionStatusConnected:
		return "connected", nil
	case PorterConnectionStatusDisconnected:
		return "disconnected", nil
	case PorterConnectionStatusActive:
		return "active", nil
	case PorterConnectionStatusActivationFailed:
		return "activation_failed", nil
	case PorterConnectionStatusDowngraded:
		return "downgraded", nil
	default:
		return "", nil
	}
}

func (s *PorterConnectionStatus) Scan(value interface{}) error {
	v, ok := value.(string)
	if !ok {
		return errors.New("invalid type for PorterConnectionStatus")
	}
	switch v {
	case "queueing":
		*s = PorterConnectionStatusQueueing
	case "connected":
		*s = PorterConnectionStatusConnected
	case "disconnected":
		*s = PorterConnectionStatusDisconnected
	case "active":
		*s = PorterConnectionStatusActive
	case "activation_failed":
		*s = PorterConnectionStatusActivationFailed
	case "downgraded":
		*s = PorterConnectionStatusDowngraded
	default:
		*s = PorterConnectionStatusUnspecified
	}
	return nil
}

type PorterContext struct {
	ID                  model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	OwnerID             model.InternalID `gorm:"column:user_porter_context"`
	GlobalName          string           `gorm:"index:idx_porter_context_global_name_region,priority:1"`
	Region              string           `gorm:"index:idx_porter_context_global_name_region,priority:2"`
	ContextJSON         string
	Name                string
	Description         string
	Status              PorterContextStatus
	HandleStatus        PorterContextHandleStatus
	HandleStatusMessage string
	UpdatedAt           time.Time
	CreatedAt           time.Time
	Owner               *model.User `gorm:"foreignKey:OwnerID"`
}

func (PorterContext) TableName() string {
	return "porter_contexts"
}

type PorterContextStatus int

const (
	PorterContextStatusUnspecified PorterContextStatus = iota
	PorterContextStatusActive
	PorterContextStatusDisabled
)

func (s PorterContextStatus) Value() (driver.Value, error) {
	switch s {
	case PorterContextStatusActive:
		return "active", nil
	case PorterContextStatusDisabled:
		return "disabled", nil
	default:
		return "", nil
	}
}

func (s *PorterContextStatus) Scan(value interface{}) error {
	v, ok := value.(string)
	if !ok {
		return errors.New("invalid type for PorterContextStatus")
	}
	switch v {
	case "active":
		*s = PorterContextStatusActive
	case "disabled":
		*s = PorterContextStatusDisabled
	default:
		*s = PorterContextStatusUnspecified
	}
	return nil
}

type PorterContextHandleStatus int

const (
	PorterContextHandleStatusUnspecified PorterContextHandleStatus = iota
	PorterContextHandleStatusActive
	PorterContextHandleStatusDowngraded
	PorterContextHandleStatusQueueing
	PorterContextHandleStatusBlocked
)

func (s PorterContextHandleStatus) Value() (driver.Value, error) {
	switch s {
	case PorterContextHandleStatusActive:
		return "active", nil
	case PorterContextHandleStatusDowngraded:
		return "downgraded", nil
	case PorterContextHandleStatusQueueing:
		return "queueing", nil
	case PorterContextHandleStatusBlocked:
		return "blocked", nil
	default:
		return "", nil
	}
}

func (s *PorterContextHandleStatus) Scan(value interface{}) error {
	v, ok := value.(string)
	if !ok {
		return errors.New("invalid type for PorterContextHandleStatus")
	}
	switch v {
	case "active":
		*s = PorterContextHandleStatusActive
	case "downgraded":
		*s = PorterContextHandleStatusDowngraded
	case "queueing":
		*s = PorterContextHandleStatusQueueing
	case "blocked":
		*s = PorterContextHandleStatusBlocked
	default:
		*s = PorterContextHandleStatusUnspecified
	}
	return nil
}

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
