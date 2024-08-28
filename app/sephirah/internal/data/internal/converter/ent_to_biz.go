package converter

import (
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/appinfo"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/deviceinfo"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feedconfig"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/image"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifyflow"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifytarget"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/portercontext"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/porterinstance"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/systemnotification"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelchesed"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelsupervisor"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model/modelfeed"
)

// goverter:converter
// goverter:output:format function
// goverter:output:file ./generated.go
// goverter:output:package github.com/tuihub/librarian/app/sephirah/internal/data/internal/converter
// goverter:matchIgnoreCase
// goverter:ignoreUnexported
// goverter:enum:exclude time:Duration
// goverter:extend TimeToTime
// goverter:extend TimeToTimePtr
type toBizConverter interface { //nolint:unused // used by generator
	// goverter:ignore PassWord
	ToBizUser(*ent.User) *modeltiphereth.User
	ToBizUserList([]*ent.User) []*modeltiphereth.User
	// goverter:enum:unknown UserTypeUnspecified
	// goverter:enum:map TypeAdmin UserTypeAdmin
	// goverter:enum:map TypeNormal UserTypeNormal
	// goverter:enum:map TypeSentinel UserTypeSentinel
	ToLibAuthUserType(user.Type) libauth.UserType
	// goverter:enum:unknown UserStatusUnspecified
	// goverter:enum:map StatusActive UserStatusActive
	// goverter:enum:map StatusBlocked UserStatusBlocked
	ToBizUserStatus(user.Status) modeltiphereth.UserStatus

	// goverter:ignore DeviceInfo
	// goverter:map CreatedAt CreateAt
	ToBizUserSession(*ent.UserSession) *modeltiphereth.UserSession
	ToBizUserSessionList([]*ent.UserSession) []*modeltiphereth.UserSession

	ToBizDeviceInfo(*ent.DeviceInfo) *modeltiphereth.DeviceInfo
	ToBizDeviceInfoList([]*ent.DeviceInfo) []*modeltiphereth.DeviceInfo
	// goverter:enum:unknown SystemTypeUnspecified
	// goverter:enum:map SystemTypeUnknown SystemTypeUnspecified
	// goverter:enum:map SystemTypeIos SystemTypeIOS
	// goverter:enum:map SystemTypeAndroid SystemTypeAndroid
	// goverter:enum:map SystemTypeWeb SystemTypeWeb
	// goverter:enum:map SystemTypeWindows SystemTypeWindows
	// goverter:enum:map SystemTypeMacos SystemTypeMacOS
	// goverter:enum:map SystemTypeLinux SystemTypeLinux
	ToBizSystemType(deviceinfo.SystemType) modeltiphereth.SystemType

	// goverter:map UpdatedAt LatestUpdateTime
	ToBizAccount(*ent.Account) *modeltiphereth.Account
	ToBizAccountList([]*ent.Account) []*modeltiphereth.Account

	// goverter:map . BinarySummary
	ToBizPorter(*ent.PorterInstance) *modelsupervisor.PorterInstance
	ToBizPorterList([]*ent.PorterInstance) []*modelsupervisor.PorterInstance
	// goverter:enum:unknown UserStatusUnspecified
	// goverter:enum:map StatusActive UserStatusActive
	// goverter:enum:map StatusBlocked UserStatusBlocked
	ToBizPorterStatus(porterinstance.Status) modeltiphereth.UserStatus

	ToBizPorterContext(*ent.PorterContext) *modelsupervisor.PorterContext
	ToBizPorterContextList([]*ent.PorterContext) []*modelsupervisor.PorterContext
	// goverter:enum:unknown PorterContextStatusUnspecified
	// goverter:enum:map StatusActive PorterContextStatusActive
	// goverter:enum:map StatusDisabled PorterContextStatusDisabled
	ToBizPorterContextStatus(portercontext.Status) modelsupervisor.PorterContextStatus

	// goverter:map . Details
	// goverter:map UpdatedAt LatestUpdateTime
	// goverter:ignore BoundInternal
	// goverter:ignore Tags
	ToBizAppInfo(*ent.AppInfo) *modelgebura.AppInfo
	ToBizAppInfoList([]*ent.AppInfo) []*modelgebura.AppInfo
	// goverter:enum:unknown AppTypeUnspecified
	// goverter:enum:map TypeUnknown AppTypeUnspecified
	// goverter:enum:map TypeGame AppTypeGame
	ToBizAppType(appinfo.Type) modelgebura.AppType

	// goverter:ignore AssignedAppInfoID
	ToBizApp(*ent.App) *modelgebura.App
	ToBizAppList([]*ent.App) []*modelgebura.App
	ToBizAppBinary(ent.AppBinary) modelgebura.AppBinary

	ToBizAppInst(*ent.AppInst) *modelgebura.AppInst
	ToBizAppInstList([]*ent.AppInst) []*modelgebura.AppInst

	// goverter:map LatestPullAt LatestPullTime
	// goverter:ignore ActionSets
	ToBizFeedConfig(*ent.FeedConfig) *modelyesod.FeedConfig
	ToBizFeedConfigList([]*ent.FeedConfig) []*modelyesod.FeedConfig
	// goverter:enum:unknown FeedConfigStatusUnspecified
	// goverter:enum:map StatusActive FeedConfigStatusActive
	// goverter:enum:map StatusSuspend FeedConfigStatusSuspend
	ToBizFeedConfigStatus(feedconfig.Status) modelyesod.FeedConfigStatus
	// goverter:enum:unknown FeedConfigPullStatusUnspecified
	// goverter:enum:map LatestPullStatusProcessing FeedConfigPullStatusProcessing
	// goverter:enum:map LatestPullStatusSuccess FeedConfigPullStatusSuccess
	// goverter:enum:map LatestPullStatusFailed FeedConfigPullStatusFailed
	ToBizFeedConfigPullStatus(feedconfig.LatestPullStatus) modelyesod.FeedConfigPullStatus

	ToBizFeedActionSet(*ent.FeedActionSet) *modelyesod.FeedActionSet
	ToBizFeedActionSetList([]*ent.FeedActionSet) []*modelyesod.FeedActionSet

	// goverter:ignore Items
	// goverter:ignore FeedType
	// goverter:ignore FeedVersion
	ToBizFeed(*ent.Feed) *modelfeed.Feed
	ToBizFeedItem(*ent.FeedItem) *modelfeed.Item
	ToBizFeedItemList([]*ent.FeedItem) []*modelfeed.Item

	ToBizFeedItemCollection(*ent.FeedItemCollection) *modelyesod.FeedItemCollection
	ToBizFeedItemCollectionList([]*ent.FeedItemCollection) []*modelyesod.FeedItemCollection

	ToBizNotifyTarget(*ent.NotifyTarget) *modelnetzach.NotifyTarget
	ToBizNotifyTargetList([]*ent.NotifyTarget) []*modelnetzach.NotifyTarget
	// goverter:enum:unknown NotifyTargetStatusUnspecified
	// goverter:enum:map StatusActive NotifyTargetStatusActive
	// goverter:enum:map StatusSuspend NotifyTargetStatusSuspend
	ToBizNotifyTargetStatus(notifytarget.Status) modelnetzach.NotifyTargetStatus
	// goverter:ignore Sources
	// goverter:ignore Targets
	ToBizNotifyFlow(*ent.NotifyFlow) *modelnetzach.NotifyFlow
	// goverter:enum:unknown NotifyFlowStatusUnspecified
	// goverter:enum:map StatusActive NotifyFlowStatusActive
	// goverter:enum:map StatusSuspend NotifyFlowStatusSuspend
	ToBizNotifyFlowStatus(notifyflow.Status) modelnetzach.NotifyFlowStatus

	// goverter:map CreatedAt CreateTime
	// goverter:map UpdatedAt UpdateTime
	ToBizSystemNotification(*ent.SystemNotification) *modelnetzach.SystemNotification
	ToBizSystemNotificationList([]*ent.SystemNotification) []*modelnetzach.SystemNotification
	// goverter:enum:unknown SystemNotificationTypeUnspecified
	// goverter:enum:map TypeSystem SystemNotificationTypeSystem
	// goverter:enum:map TypeUser SystemNotificationTypeUser
	ToBizSystemNotificationType(systemnotification.Type) modelnetzach.SystemNotificationType
	// goverter:enum:unknown SystemNotificationLevelUnspecified
	// goverter:enum:map LevelInfo SystemNotificationLevelInfo
	// goverter:enum:map LevelWarn SystemNotificationLevelWarning
	// goverter:enum:map LevelError SystemNotificationLevelError
	// goverter:enum:map LevelOngoing SystemNotificationLevelOngoing
	ToBizSystemNotificationLevel(systemnotification.Level) modelnetzach.SystemNotificationLevel
	// goverter:enum:unknown SystemNotificationStatusUnspecified
	// goverter:enum:map StatusUnread SystemNotificationStatusUnread
	// goverter:enum:map StatusRead SystemNotificationStatusRead
	// goverter:enum:map StatusDismissed SystemNotificationStatusDismissed
	ToBizSystemNotificationStatus(systemnotification.Status) modelnetzach.SystemNotificationStatus

	ToBizImage(*ent.Image) *modelchesed.Image
	ToBizImageList([]*ent.Image) []*modelchesed.Image
	// goverter:enum:unknown ImageStatusUnspecified
	// goverter:enum:map StatusUploaded ImageStatusUploaded
	// goverter:enum:map StatusScanned ImageStatusScanned
	ToBizImageStatus(image.Status) modelchesed.ImageStatus
}

func TimeToTime(t time.Time) time.Time {
	return t
}

func TimeToTimePtr(t *time.Time) *time.Time {
	return t
}
