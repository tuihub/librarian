package modelnetzach

import (
	"github.com/tuihub/librarian/model"
)

type NotifyFlow struct {
	ID          model.InternalID
	Name        string
	Description string
	Sources     []*NotifyFlowSource
	Targets     []*NotifyFlowTarget
	Status      NotifyFlowStatus
}

type NotifyFlowSource struct {
	SourceID model.InternalID
	Filter   *NotifyFilter
}

type NotifyFlowTarget struct {
	TargetID  model.InternalID
	Filter    *NotifyFilter
	ChannelID string
}

type NotifyFlowStatus int

const (
	NotifyFlowStatusUnspecified NotifyFlowStatus = iota
	NotifyFlowStatusActive
	NotifyFlowStatusSuspend
)

type NotifyTarget struct {
	ID          model.InternalID
	Name        string
	Description string
	Destination string
	Status      NotifyTargetStatus
	Token       string
}

type NotifyTargetStatus int

const (
	NotifyTargetStatusUnspecified NotifyTargetStatus = iota
	NotifyTargetStatusActive
	NotifyTargetStatusSuspend
)

type NotifyFilter struct {
	ExcludeKeywords []string
	IncludeKeywords []string
}
