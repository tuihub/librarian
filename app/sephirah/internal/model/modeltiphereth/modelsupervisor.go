package modeltiphereth

import "github.com/tuihub/librarian/model"

type PorterInstance struct {
	ID             model.InternalID
	Name           string
	Version        string
	GlobalName     string
	Address        string
	FeatureSummary *PorterFeatureSummary
	Status         PorterInstanceStatus
}

type PorterFeatureSummary struct {
	SupportedAccounts           []*SupportedAccount `json:"supported_accounts"`
	SupportedAppSources         []string            `json:"supported_app_sources"`
	SupportedFeedSources        []string            `json:"supported_feed_sources"`
	SupportedNotifyDestinations []string            `json:"supported_notify_destinations"`
}
type SupportedAccount struct {
	Platform         string                         `json:"platform"`
	AppRelationTypes []model.AccountAppRelationType `json:"app_relation_types"`
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
	SupportedAccountPlatforms   []string
	SupportedAppSources         []string
	SupportedFeedSources        []string
	SupportedNotifyDestinations []string
}
