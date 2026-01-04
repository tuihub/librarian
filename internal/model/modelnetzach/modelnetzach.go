package modelnetzach

import (
	"database/sql/driver"
	"errors"
	"time"

	"github.com/tuihub/librarian/internal/model"

	"gorm.io/gorm"
)

type NotifyFlow struct {
	ID          model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	OwnerID     model.InternalID `gorm:"column:user_notify_flow;index"`
	Name        string
	Description string
	Sources     []*NotifyFlowSource `gorm:"foreignKey:NotifyFlowID"`
	Targets     []*NotifyFlowTarget `gorm:"foreignKey:NotifyFlowID"`
	Status      NotifyFlowStatus
	UpdatedAt   time.Time
	CreatedAt   time.Time
	Owner       *model.User `gorm:"foreignKey:OwnerID"`
}

func (NotifyFlow) TableName() string {
	return "notify_flows"
}

type NotifyFlowSource struct {
	NotifyFlowID          model.InternalID `gorm:"primaryKey"`
	NotifySourceID        model.InternalID `gorm:"primaryKey"`
	FilterIncludeKeywords []string         `gorm:"serializer:json"`
	FilterExcludeKeywords []string         `gorm:"serializer:json"`
	UpdatedAt             time.Time
	CreatedAt             time.Time

	// Biz fields
	SourceID model.InternalID `gorm:"-"` // Map to NotifySourceID
	Filter   *NotifyFilter    `gorm:"-"`
}

func (NotifyFlowSource) TableName() string {
	return "notify_flow_sources"
}

func (n *NotifyFlowSource) BeforeSave(tx *gorm.DB) error {
	if n.SourceID != 0 {
		n.NotifySourceID = n.SourceID
	}
	if n.Filter != nil {
		n.FilterIncludeKeywords = n.Filter.IncludeKeywords
		n.FilterExcludeKeywords = n.Filter.ExcludeKeywords
	}
	return nil
}

func (n *NotifyFlowSource) AfterFind(tx *gorm.DB) error {
	n.SourceID = n.NotifySourceID
	n.Filter = &NotifyFilter{
		IncludeKeywords: n.FilterIncludeKeywords,
		ExcludeKeywords: n.FilterExcludeKeywords,
	}
	return nil
}

type NotifyFlowTarget struct {
	NotifyFlowID          model.InternalID `gorm:"primaryKey"`
	NotifyTargetID        model.InternalID `gorm:"primaryKey"`
	FilterIncludeKeywords []string         `gorm:"serializer:json"`
	FilterExcludeKeywords []string         `gorm:"serializer:json"`
	UpdatedAt             time.Time
	CreatedAt             time.Time

	// Biz fields
	TargetID model.InternalID `gorm:"-"` // Map to NotifyTargetID
	Filter   *NotifyFilter    `gorm:"-"`
}

func (NotifyFlowTarget) TableName() string {
	return "notify_flow_targets"
}

func (n *NotifyFlowTarget) BeforeSave(tx *gorm.DB) error {
	if n.TargetID != 0 {
		n.NotifyTargetID = n.TargetID
	}
	if n.Filter != nil {
		n.FilterIncludeKeywords = n.Filter.IncludeKeywords
		n.FilterExcludeKeywords = n.Filter.ExcludeKeywords
	}
	return nil
}

func (n *NotifyFlowTarget) AfterFind(tx *gorm.DB) error {
	n.TargetID = n.NotifyTargetID
	n.Filter = &NotifyFilter{
		IncludeKeywords: n.FilterIncludeKeywords,
		ExcludeKeywords: n.FilterExcludeKeywords,
	}
	return nil
}

type NotifyFlowStatus int

const (
	NotifyFlowStatusUnspecified NotifyFlowStatus = iota
	NotifyFlowStatusActive
	NotifyFlowStatusSuspend
)

func (s NotifyFlowStatus) Value() (driver.Value, error) {
	switch s {
	case NotifyFlowStatusActive:
		return "active", nil
	case NotifyFlowStatusSuspend:
		return "suspend", nil
	default:
		return "", nil
	}
}

func (s *NotifyFlowStatus) Scan(value interface{}) error {
	v, ok := value.(string)
	if !ok {
		return errors.New("invalid type for NotifyFlowStatus")
	}
	switch v {
	case "active":
		*s = NotifyFlowStatusActive
	case "suspend":
		*s = NotifyFlowStatusSuspend
	default:
		*s = NotifyFlowStatusUnspecified
	}
	return nil
}

type NotifyTarget struct {
	ID          model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	OwnerID     model.InternalID `gorm:"column:user_notify_target;index"`
	Name        string
	Description string
	Destination *model.FeatureRequest `gorm:"serializer:json"`
	Status      NotifyTargetStatus
	UpdatedAt   time.Time
	CreatedAt   time.Time
	Owner       *model.User `gorm:"foreignKey:OwnerID"`
}

func (NotifyTarget) TableName() string {
	return "notify_targets"
}

type NotifyTargetStatus int

const (
	NotifyTargetStatusUnspecified NotifyTargetStatus = iota
	NotifyTargetStatusActive
	NotifyTargetStatusSuspend
)

func (s NotifyTargetStatus) Value() (driver.Value, error) {
	switch s {
	case NotifyTargetStatusActive:
		return "active", nil
	case NotifyTargetStatusSuspend:
		return "suspend", nil
	default:
		return "", nil
	}
}

func (s *NotifyTargetStatus) Scan(value interface{}) error {
	v, ok := value.(string)
	if !ok {
		return errors.New("invalid type for NotifyTargetStatus")
	}
	switch v {
	case "active":
		*s = NotifyTargetStatusActive
	case "suspend":
		*s = NotifyTargetStatusSuspend
	default:
		*s = NotifyTargetStatusUnspecified
	}
	return nil
}

type NotifyFilter struct {
	ExcludeKeywords []string
	IncludeKeywords []string
}

type SystemNotification struct {
	ID         model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	UserID     model.InternalID `gorm:"index"`
	Type       SystemNotificationType
	Level      SystemNotificationLevel
	Status     SystemNotificationStatus
	Title      string
	Content    string
	CreateTime time.Time `gorm:"column:created_at"` // Map to created_at
	UpdateTime time.Time `gorm:"column:updated_at"` // Map to updated_at
}

func (SystemNotification) TableName() string {
	return "system_notifications"
}

type SystemNotificationType int

const (
	SystemNotificationTypeUnspecified SystemNotificationType = iota
	SystemNotificationTypeSystem
	SystemNotificationTypeUser
)

func (t SystemNotificationType) Value() (driver.Value, error) {
	switch t {
	case SystemNotificationTypeSystem:
		return "system", nil
	case SystemNotificationTypeUser:
		return "user", nil
	default:
		return "", nil
	}
}

func (t *SystemNotificationType) Scan(value interface{}) error {
	v, ok := value.(string)
	if !ok {
		return errors.New("invalid type for SystemNotificationType")
	}
	switch v {
	case "system":
		*t = SystemNotificationTypeSystem
	case "user":
		*t = SystemNotificationTypeUser
	default:
		*t = SystemNotificationTypeUnspecified
	}
	return nil
}

type SystemNotificationLevel int

const (
	SystemNotificationLevelUnspecified SystemNotificationLevel = iota
	SystemNotificationLevelOngoing
	SystemNotificationLevelError
	SystemNotificationLevelWarning
	SystemNotificationLevelInfo
)

func (l SystemNotificationLevel) Value() (driver.Value, error) {
	switch l {
	case SystemNotificationLevelInfo:
		return "info", nil
	case SystemNotificationLevelWarning:
		return "warn", nil
	case SystemNotificationLevelError:
		return "error", nil
	case SystemNotificationLevelOngoing:
		return "ongoing", nil
	default:
		return "", nil
	}
}

func (l *SystemNotificationLevel) Scan(value interface{}) error {
	v, ok := value.(string)
	if !ok {
		return errors.New("invalid type for SystemNotificationLevel")
	}
	switch v {
	case "info":
		*l = SystemNotificationLevelInfo
	case "warn":
		*l = SystemNotificationLevelWarning
	case "error":
		*l = SystemNotificationLevelError
	case "ongoing":
		*l = SystemNotificationLevelOngoing
	default:
		*l = SystemNotificationLevelUnspecified
	}
	return nil
}

type SystemNotificationStatus int

const (
	SystemNotificationStatusUnspecified SystemNotificationStatus = iota
	SystemNotificationStatusUnread
	SystemNotificationStatusRead
	SystemNotificationStatusDismissed
)

func (s SystemNotificationStatus) Value() (driver.Value, error) {
	switch s {
	case SystemNotificationStatusUnread:
		return "unread", nil
	case SystemNotificationStatusRead:
		return "read", nil
	case SystemNotificationStatusDismissed:
		return "dismissed", nil
	default:
		return "", nil
	}
}

func (s *SystemNotificationStatus) Scan(value interface{}) error {
	v, ok := value.(string)
	if !ok {
		return errors.New("invalid type for SystemNotificationStatus")
	}
	switch v {
	case "unread":
		*s = SystemNotificationStatusUnread
	case "read":
		*s = SystemNotificationStatusRead
	case "dismissed":
		*s = SystemNotificationStatusDismissed
	default:
		*s = SystemNotificationStatusUnspecified
	}
	return nil
}

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

type NotifySource struct {
	ID                   model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	OwnerID              model.InternalID `gorm:"column:user_notify_source;index"`
	FeedConfigID         *model.InternalID
	FeedItemCollectionID *model.InternalID
	UpdatedAt            time.Time
	CreatedAt            time.Time
	Owner                *model.User `gorm:"foreignKey:OwnerID"`
}

func (NotifySource) TableName() string {
	return "notify_sources"
}
