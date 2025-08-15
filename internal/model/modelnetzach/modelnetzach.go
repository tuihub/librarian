package modelnetzach

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
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
	TargetID model.InternalID
	Filter   *NotifyFilter
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
	Destination *model.FeatureRequest
	Status      NotifyTargetStatus
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

type SystemNotification struct {
	ID         model.InternalID
	Type       SystemNotificationType
	Level      SystemNotificationLevel
	Status     SystemNotificationStatus
	Title      string
	Content    string
	CreateTime time.Time
	UpdateTime time.Time
}

type SystemNotificationType int

const (
	SystemNotificationTypeUnspecified SystemNotificationType = iota
	SystemNotificationTypeSystem
	SystemNotificationTypeUser
)

type SystemNotificationLevel int

const (
	SystemNotificationLevelUnspecified SystemNotificationLevel = iota
	SystemNotificationLevelOngoing
	SystemNotificationLevelError
	SystemNotificationLevelWarning
	SystemNotificationLevelInfo
)

type SystemNotificationStatus int

const (
	SystemNotificationStatusUnspecified SystemNotificationStatus = iota
	SystemNotificationStatusUnread
	SystemNotificationStatusRead
	SystemNotificationStatusDismissed
)

type SystemNotify struct {
	UserID       model.InternalID
	Notification SystemNotification
}

const (
	SystemNotifyTitleCronJob = "Server Scheduled Task"
)

func NewSystemNotify(
	level SystemNotificationLevel,
	title string,
	content string,
) SystemNotify {
	return SystemNotify{
		UserID: 0,
		Notification: SystemNotification{
			ID:         0,
			Type:       SystemNotificationTypeSystem,
			Level:      level,
			Status:     SystemNotificationStatusUnread,
			Title:      title,
			Content:    content,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		},
	}
}

func NewUserNotify(
	userID model.InternalID,
	level SystemNotificationLevel,
	title string,
	content string,
) SystemNotify {
	return SystemNotify{
		UserID: userID,
		Notification: SystemNotification{
			ID:         0,
			Type:       SystemNotificationTypeUser,
			Level:      level,
			Status:     SystemNotificationStatusUnread,
			Title:      title,
			Content:    content,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		},
	}
}
