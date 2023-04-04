package modelnetzach

import "github.com/tuihub/librarian/internal/model"

type NotifyFlow struct {
	ID          model.InternalID
	Name        string
	Description string
	Source      *NotifyFlowSource
	Targets     []*NotifyFlowTarget
	Status      NotifyFlowStatus
}

type NotifyFlowSource struct {
	FeedIDFilter []model.InternalID
}

type NotifyFlowTarget struct {
	TargetID  model.InternalID
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
	Type        NotifyTargetType
	Status      NotifyTargetStatus
	Token       string
}

type NotifyTargetType int

const (
	NotifyTargetTypeUnspecified NotifyTargetType = iota
	NotifyTargetTypeTelegram
)

type NotifyTargetStatus int

const (
	NotifyTargetStatusUnspecified NotifyTargetStatus = iota
	NotifyTargetStatusActive
	NotifyTargetStatusSuspend
)
