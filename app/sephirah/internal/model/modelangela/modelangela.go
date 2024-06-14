package modelangela

import (
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
)

type PullAccountAppInfoRelation struct {
	ID                model.InternalID
	Platform          string
	PlatformAccountID string
}

type PullAppInfo struct {
	ID              model.InternalID
	AppInfoID       modelgebura.AppInfoID
	IgnoreRateLimit bool
}

type NotifyRouter struct {
	FeedID   model.InternalID
	Messages []*modelfeed.Item
}

type FeedToNotifyFlowValue []model.InternalID

type NotifyPush struct {
	Target   modelnetzach.NotifyFlowTarget
	Messages []*modelfeed.Item
}

type ParseFeedItemDigest struct {
	ID model.InternalID
}

type UpdateAppInfoIndex struct {
	IDs []model.InternalID
}

type SystemNotify struct {
	UserID       model.InternalID
	Notification modelnetzach.SystemNotification
}

func NewSystemNotify(
	level modelnetzach.SystemNotificationLevel,
	status modelnetzach.SystemNotificationStatus,
	title string,
	content string,
) *SystemNotify {
	return &SystemNotify{
		UserID: 0,
		Notification: modelnetzach.SystemNotification{
			ID:         0,
			Type:       modelnetzach.SystemNotificationTypeSystem,
			Level:      level,
			Status:     status,
			Title:      title,
			Content:    content,
			CreateTime: time.Now(),
		},
	}
}

func NewUserNotify(
	userID model.InternalID,
	level modelnetzach.SystemNotificationLevel,
	status modelnetzach.SystemNotificationStatus,
	title string,
	content string,
) *SystemNotify {
	return &SystemNotify{
		UserID: userID,
		Notification: modelnetzach.SystemNotification{
			ID:         0,
			Type:       modelnetzach.SystemNotificationTypeUser,
			Level:      level,
			Status:     status,
			Title:      title,
			Content:    content,
			CreateTime: time.Now(),
		},
	}
}
