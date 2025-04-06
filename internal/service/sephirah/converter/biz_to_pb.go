package converter

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	"github.com/tuihub/librarian/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/model/modelnetzach"
	"github.com/tuihub/librarian/internal/model/modelsupervisor"
	"github.com/tuihub/librarian/internal/model/modelyesod"
	sephirah "github.com/tuihub/protos/pkg/librarian/sephirah/v1/sephirah"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// goverter:converter
// goverter:output:format function
// goverter:output:file ./generated.go
// goverter:output:package github.com/tuihub/librarian/internal/service/sephirah/converter
// goverter:matchIgnoreCase
// goverter:ignoreUnexported
// goverter:extend ToPBInternalID
// goverter:extend ToPBInternalIDPtr
// goverter:extend ToPBTime
// goverter:extend ToPBTimePtr
// goverter:extend ToPBDuration
type toPBConverter interface { //nolint:unused // used by generator
	ToPBTimeRange(*model.TimeRange) *librarian.TimeRange
	ToPBInternalIDList([]model.InternalID) []*librarian.InternalID
	ToPBServerFeatureSummary(*modelsupervisor.ServerFeatureSummary) *librarian.FeatureSummary
	ToPBFeatureFlag(*modelsupervisor.FeatureFlag) *librarian.FeatureFlag
	ToPBFeatureFlagList([]*modelsupervisor.FeatureFlag) []*librarian.FeatureFlag
	ToPBFeatureRequest(*modelsupervisor.FeatureRequest) *librarian.FeatureRequest

	// goverter:map ID DeviceId
	ToPBDeviceInfo(*model.Device) *sephirah.Device
	ToPBDeviceInfoList([]*model.Device) []*sephirah.Device
	// goverter:enum:unknown SystemType_SYSTEM_TYPE_UNSPECIFIED
	// goverter:enum:map SystemTypeUnspecified SystemType_SYSTEM_TYPE_UNSPECIFIED
	// goverter:enum:map SystemTypeIOS SystemType_SYSTEM_TYPE_IOS
	// goverter:enum:map SystemTypeAndroid SystemType_SYSTEM_TYPE_ANDROID
	// goverter:enum:map SystemTypeWeb SystemType_SYSTEM_TYPE_WEB
	// goverter:enum:map SystemTypeWindows SystemType_SYSTEM_TYPE_WINDOWS
	// goverter:enum:map SystemTypeMacOS SystemType_SYSTEM_TYPE_MACOS
	// goverter:enum:map SystemTypeLinux SystemType_SYSTEM_TYPE_LINUX
	ToPBSystemType(model.SystemType) sephirah.SystemType

	// goverter:map CreateAt CreateTime
	// goverter:map ExpireAt ExpireTime
	// goverter:map Device DeviceInfo
	ToPBUserSession(*model.Session) *sephirah.UserSession
	ToPBUserSessionList([]*model.Session) []*sephirah.UserSession

	// goverter:ignore Password
	ToPBUser(*model.User) *sephirah.User
	ToPBUserList([]*model.User) []*sephirah.User
	// goverter:enum:unknown UserType_USER_TYPE_UNSPECIFIED
	// goverter:enum:map UserTypeUnspecified UserType_USER_TYPE_UNSPECIFIED
	// goverter:enum:map UserTypeAdmin UserType_USER_TYPE_ADMIN
	// goverter:enum:map UserTypeNormal UserType_USER_TYPE_NORMAL
	// goverter:enum:map UserTypeSentinel @ignore
	// goverter:enum:map UserTypePorter @ignore
	ToPBUserType(model.UserType) sephirah.UserType
	// goverter:enum:unknown UserStatus_USER_STATUS_UNSPECIFIED
	// goverter:enum:map UserStatusUnspecified UserStatus_USER_STATUS_UNSPECIFIED
	// goverter:enum:map UserStatusActive UserStatus_USER_STATUS_ACTIVE
	// goverter:enum:map UserStatusBlocked UserStatus_USER_STATUS_BLOCKED
	ToPBUserStatus(model.UserStatus) sephirah.UserStatus

	ToPBAccount(*model.Account) *sephirah.Account
	ToPBAccountList([]*model.Account) []*sephirah.Account

	// goverter:autoMap PorterInstance
	ToPBPorter(*modelsupervisor.PorterInstanceController) *sephirah.Porter
	ToPBPorterList([]*modelsupervisor.PorterInstanceController) []*sephirah.Porter
	// goverter:enum:unknown PorterConnectionStatus_PORTER_CONNECTION_STATUS_UNSPECIFIED
	// goverter:enum:map PorterConnectionStatusUnspecified PorterConnectionStatus_PORTER_CONNECTION_STATUS_UNSPECIFIED
	// goverter:enum:map PorterConnectionStatusConnected PorterConnectionStatus_PORTER_CONNECTION_STATUS_CONNECTED
	// goverter:enum:map PorterConnectionStatusDisconnected PorterConnectionStatus_PORTER_CONNECTION_STATUS_DISCONNECTED
	// goverter:enum:map PorterConnectionStatusActive PorterConnectionStatus_PORTER_CONNECTION_STATUS_ACTIVE
	// goverter:enum:map PorterConnectionStatusActivationFailed PorterConnectionStatus_PORTER_CONNECTION_STATUS_ACTIVATION_FAILED
	// goverter:enum:map PorterConnectionStatusDowngraded PorterConnectionStatus_PORTER_CONNECTION_STATUS_DOWNGRADED
	ToPBPorterConnectionStatus(modelsupervisor.PorterConnectionStatus) sephirah.PorterConnectionStatus

	// goverter:autoMap PorterContext
	ToPBPorterContext(*modelsupervisor.PorterContextController) *sephirah.PorterContext
	ToPBPorterContextList([]*modelsupervisor.PorterContextController) []*sephirah.PorterContext
	// goverter:enum:unknown PorterContextStatus_PORTER_CONTEXT_STATUS_UNSPECIFIED
	// goverter:enum:map PorterContextStatusUnspecified PorterContextStatus_PORTER_CONTEXT_STATUS_UNSPECIFIED
	// goverter:enum:map PorterContextStatusActive PorterContextStatus_PORTER_CONTEXT_STATUS_ACTIVE
	// goverter:enum:map PorterContextStatusDisabled PorterContextStatus_PORTER_CONTEXT_STATUS_DISABLED
	ToPBPorterContextStatus(modelsupervisor.PorterContextStatus) sephirah.PorterContextStatus
	// goverter:enum:unknown PorterContextHandleStatus_PORTER_CONTEXT_HANDLE_STATUS_UNSPECIFIED
	// goverter:enum:map PorterContextHandleStatusUnspecified PorterContextHandleStatus_PORTER_CONTEXT_HANDLE_STATUS_UNSPECIFIED
	// goverter:enum:map PorterContextHandleStatusActive PorterContextHandleStatus_PORTER_CONTEXT_HANDLE_STATUS_ACTIVE
	// goverter:enum:map PorterContextHandleStatusDowngraded PorterContextHandleStatus_PORTER_CONTEXT_HANDLE_STATUS_DOWNGRADED
	// goverter:enum:map PorterContextHandleStatusQueueing PorterContextHandleStatus_PORTER_CONTEXT_HANDLE_STATUS_QUEUEING
	// goverter:enum:map PorterContextHandleStatusBlocked PorterContextHandleStatus_PORTER_CONTEXT_HANDLE_STATUS_BLOCKED
	ToPBPorterContextHandleStatus(modelsupervisor.PorterContextHandleStatus) sephirah.PorterContextHandleStatus

	ToPBPorterDigest(*modelsupervisor.PorterDigest) *sephirah.PorterDigest
	ToPBPorterDigestList([]*modelsupervisor.PorterDigest) []*sephirah.PorterDigest

	// goverter:map AlternativeNames AltNames
	ToPBAppInfo(*modelgebura.AppInfo) *sephirah.AppInfo
	ToPBAppInfoList([]*modelgebura.AppInfo) []*sephirah.AppInfo
	// goverter:enum:unknown AppType_APP_TYPE_UNSPECIFIED
	// goverter:enum:map AppTypeUnspecified AppType_APP_TYPE_UNSPECIFIED
	// goverter:enum:map AppTypeGame AppType_APP_TYPE_GAME
	ToPBAppType(modelgebura.AppType) sephirah.AppType

	// goverter:map AlternativeNames AltNames
	ToPBApp(*modelgebura.App) *sephirah.App
	ToPBAppList([]*modelgebura.App) []*sephirah.App

	ToPBAppRunTime(*modelgebura.AppRunTime) *sephirah.AppRunTime
	ToPBAppRunTimeList([]*modelgebura.AppRunTime) []*sephirah.AppRunTime

	ToPBAppCategory(*modelgebura.AppCategory) *sephirah.AppCategory
	ToPBAppCategoryList([]*modelgebura.AppCategory) []*sephirah.AppCategory

	// ToPBAppInst(*modelgebura.AppInst) *sephirah.AppInst
	// ToPBAppInstList([]*modelgebura.AppInst) []*sephirah.AppInst

	ToPBFeed(*modelfeed.Feed) *librarian.Feed
	ToPBFeedItem(*modelfeed.Item) *librarian.FeedItem
	ToPBFeedItemList([]*modelfeed.Item) []*librarian.FeedItem
	ToPBFeedImage(*modelfeed.Image) *librarian.FeedImage
	ToPBEnclosure(*modelfeed.Enclosure) *librarian.FeedEnclosure
	// goverter:map LatestPullStatus | ToPBFeedConfigPullStatus
	ToPBFeedConfig(*modelyesod.FeedConfig) *sephirah.FeedConfig
	// goverter:enum:unknown FeedConfigStatus_FEED_CONFIG_STATUS_UNSPECIFIED
	// goverter:enum:map FeedConfigStatusUnspecified FeedConfigStatus_FEED_CONFIG_STATUS_UNSPECIFIED
	// goverter:enum:map FeedConfigStatusActive FeedConfigStatus_FEED_CONFIG_STATUS_ACTIVE
	// goverter:enum:map FeedConfigStatusSuspend FeedConfigStatus_FEED_CONFIG_STATUS_SUSPEND
	ToPBFeedConfigStatus(modelyesod.FeedConfigStatus) sephirah.FeedConfigStatus
	// goverter:map FeedConfig Config
	ToPBFeedWithConfig(*modelyesod.FeedWithConfig) *sephirah.ListFeedConfigsResponse_FeedWithConfig
	ToPBFeedWithConfigList([]*modelyesod.FeedWithConfig) []*sephirah.ListFeedConfigsResponse_FeedWithConfig
	ToPBFeedItemDigest(*modelyesod.FeedItemDigest) *sephirah.FeedItemDigest
	ToPBFeedItemDigestList([]*modelyesod.FeedItemDigest) []*sephirah.FeedItemDigest

	ToPBFeedActionSet(*modelyesod.FeedActionSet) *sephirah.FeedActionSet
	ToPBFeedActionSetList([]*modelyesod.FeedActionSet) []*sephirah.FeedActionSet

	ToPBFeedItemCollection(*modelyesod.FeedItemCollection) *sephirah.FeedItemCollection
	ToPBFeedItemCollectionList([]*modelyesod.FeedItemCollection) []*sephirah.FeedItemCollection

	ToPBNotifyTarget(*modelnetzach.NotifyTarget) *sephirah.NotifyTarget
	ToPBNotifyTargetList([]*modelnetzach.NotifyTarget) []*sephirah.NotifyTarget
	// goverter:enum:unknown NotifyTargetStatus_NOTIFY_TARGET_STATUS_UNSPECIFIED
	// goverter:enum:map NotifyTargetStatusUnspecified NotifyTargetStatus_NOTIFY_TARGET_STATUS_UNSPECIFIED
	// goverter:enum:map NotifyTargetStatusActive NotifyTargetStatus_NOTIFY_TARGET_STATUS_ACTIVE
	// goverter:enum:map NotifyTargetStatusSuspend NotifyTargetStatus_NOTIFY_TARGET_STATUS_SUSPEND
	ToPBNotifyTargetStatus(modelnetzach.NotifyTargetStatus) sephirah.NotifyTargetStatus

	ToPBNotifyFlow(*modelnetzach.NotifyFlow) *sephirah.NotifyFlow
	// goverter:enum:unknown NotifyFlowStatus_NOTIFY_FLOW_STATUS_UNSPECIFIED
	// goverter:enum:map NotifyFlowStatusUnspecified NotifyFlowStatus_NOTIFY_FLOW_STATUS_UNSPECIFIED
	// goverter:enum:map NotifyFlowStatusActive NotifyFlowStatus_NOTIFY_FLOW_STATUS_ACTIVE
	// goverter:enum:map NotifyFlowStatusSuspend NotifyFlowStatus_NOTIFY_FLOW_STATUS_SUSPEND
	ToPBNotifyFlowStatus(modelnetzach.NotifyFlowStatus) sephirah.NotifyFlowStatus
	ToPBNotifyFlowSource(*modelnetzach.NotifyFlowSource) *sephirah.NotifyFlowSource
	ToPBNotifyFlowTarget(*modelnetzach.NotifyFlowTarget) *sephirah.NotifyFlowTarget
	ToPBNotifyFlowList([]*modelnetzach.NotifyFlow) []*sephirah.NotifyFlow

	ToPBSystemNotification(*modelnetzach.SystemNotification) *sephirah.SystemNotification
	ToPBSystemNotificationList([]*modelnetzach.SystemNotification) []*sephirah.SystemNotification
	// goverter:enum:unknown SystemNotificationLevel_SYSTEM_NOTIFICATION_LEVEL_UNSPECIFIED
	// goverter:enum:map SystemNotificationLevelUnspecified SystemNotificationLevel_SYSTEM_NOTIFICATION_LEVEL_UNSPECIFIED
	// goverter:enum:map SystemNotificationLevelInfo SystemNotificationLevel_SYSTEM_NOTIFICATION_LEVEL_INFO
	// goverter:enum:map SystemNotificationLevelWarning SystemNotificationLevel_SYSTEM_NOTIFICATION_LEVEL_WARNING
	// goverter:enum:map SystemNotificationLevelError SystemNotificationLevel_SYSTEM_NOTIFICATION_LEVEL_ERROR
	// goverter:enum:map SystemNotificationLevelOngoing SystemNotificationLevel_SYSTEM_NOTIFICATION_LEVEL_ONGOING
	ToPBSystemNotificationLevel(modelnetzach.SystemNotificationLevel) sephirah.SystemNotificationLevel
	// goverter:enum:unknown SystemNotificationStatus_SYSTEM_NOTIFICATION_STATUS_UNSPECIFIED
	// goverter:enum:map SystemNotificationStatusUnspecified SystemNotificationStatus_SYSTEM_NOTIFICATION_STATUS_UNSPECIFIED
	// goverter:enum:map SystemNotificationStatusUnread SystemNotificationStatus_SYSTEM_NOTIFICATION_STATUS_UNREAD
	// goverter:enum:map SystemNotificationStatusRead SystemNotificationStatus_SYSTEM_NOTIFICATION_STATUS_READ
	// goverter:enum:map SystemNotificationStatusDismissed SystemNotificationStatus_SYSTEM_NOTIFICATION_STATUS_DISMISSED
	ToPBSystemNotificationStatus(modelnetzach.SystemNotificationStatus) sephirah.SystemNotificationStatus
}

func DurationPBToDuration(t *durationpb.Duration) time.Duration {
	if t == nil {
		return time.Duration(0)
	}
	return t.AsDuration()
}

func ToPBInternalID(id model.InternalID) *librarian.InternalID {
	return &librarian.InternalID{Id: int64(id)}
}

func ToPBInternalIDPtr(id *model.InternalID) *librarian.InternalID {
	if id == nil {
		return nil
	}
	return &librarian.InternalID{Id: int64(*id)}
}

func ToPBTime(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}

func ToPBTimePtr(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}

func ToPBDuration(d time.Duration) *durationpb.Duration {
	return durationpb.New(d)
}

func ToPBFeedConfigPullStatus(s modelyesod.FeedConfigPullStatus) *sephirah.FeedConfigPullStatus {
	var status sephirah.FeedConfigPullStatus
	switch s {
	case modelyesod.FeedConfigPullStatusUnspecified:
		status = sephirah.FeedConfigPullStatus_FEED_CONFIG_PULL_STATUS_UNSPECIFIED
	case modelyesod.FeedConfigPullStatusProcessing:
		status = sephirah.FeedConfigPullStatus_FEED_CONFIG_PULL_STATUS_PROCESSING
	case modelyesod.FeedConfigPullStatusSuccess:
		status = sephirah.FeedConfigPullStatus_FEED_CONFIG_PULL_STATUS_SUCCESS
	case modelyesod.FeedConfigPullStatusFailed:
		status = sephirah.FeedConfigPullStatus_FEED_CONFIG_PULL_STATUS_FAILED
	default:
		status = sephirah.FeedConfigPullStatus_FEED_CONFIG_PULL_STATUS_UNSPECIFIED
	}
	return &status
}
