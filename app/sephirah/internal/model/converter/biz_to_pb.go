package converter

import (
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelsupervisor"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// goverter:converter
// goverter:output:format function
// goverter:output:file ./generated.go
// goverter:output:package github.com/tuihub/librarian/app/sephirah/internal/model/converter
// goverter:matchIgnoreCase
// goverter:ignoreUnexported
// goverter:extend ToPBInternalID
// goverter:extend ToPBTime
// goverter:extend ToPBTimePtr
// goverter:extend ToPBDuration
type toPBConverter interface { //nolint:unused // used by generator
	ToPBTimeRange(*model.TimeRange) *librarian.TimeRange
	ToPBInternalIDList([]model.InternalID) []*librarian.InternalID
	ToPBServerFeatureSummary(*modelsupervisor.ServerFeatureSummary) *pb.ServerFeatureSummary
	ToPBFeatureFlag(*modelsupervisor.FeatureFlag) *librarian.FeatureFlag
	ToPBFeatureRequest(*modelsupervisor.FeatureRequest) *librarian.FeatureRequest

	// goverter:map ID DeviceId
	ToPBDeviceInfo(*modeltiphereth.DeviceInfo) *pb.DeviceInfo
	ToPBDeviceInfoList([]*modeltiphereth.DeviceInfo) []*pb.DeviceInfo
	// goverter:enum:unknown SystemType_SYSTEM_TYPE_UNSPECIFIED
	// goverter:enum:map SystemTypeUnspecified SystemType_SYSTEM_TYPE_UNSPECIFIED
	// goverter:enum:map SystemTypeIOS SystemType_SYSTEM_TYPE_IOS
	// goverter:enum:map SystemTypeAndroid SystemType_SYSTEM_TYPE_ANDROID
	// goverter:enum:map SystemTypeWeb SystemType_SYSTEM_TYPE_WEB
	// goverter:enum:map SystemTypeWindows SystemType_SYSTEM_TYPE_WINDOWS
	// goverter:enum:map SystemTypeMacOS SystemType_SYSTEM_TYPE_MACOS
	// goverter:enum:map SystemTypeLinux SystemType_SYSTEM_TYPE_LINUX
	ToPBSystemType(modeltiphereth.SystemType) pb.SystemType

	// goverter:map CreateAt CreateTime
	// goverter:map ExpireAt ExpireTime
	ToPBUserSession(*modeltiphereth.UserSession) *pb.UserSession
	ToPBUserSessionList([]*modeltiphereth.UserSession) []*pb.UserSession

	// goverter:ignore Password
	ToPBUser(*modeltiphereth.User) *pb.User
	ToPBUserList([]*modeltiphereth.User) []*pb.User
	// goverter:enum:unknown UserType_USER_TYPE_UNSPECIFIED
	// goverter:enum:map UserTypeUnspecified UserType_USER_TYPE_UNSPECIFIED
	// goverter:enum:map UserTypeAdmin UserType_USER_TYPE_ADMIN
	// goverter:enum:map UserTypeNormal UserType_USER_TYPE_NORMAL
	// goverter:enum:map UserTypeSentinel UserType_USER_TYPE_SENTINEL
	// goverter:enum:map UserTypePorter UserType_USER_TYPE_PORTER
	ToPBUserType(libauth.UserType) pb.UserType
	// goverter:enum:unknown UserStatus_USER_STATUS_UNSPECIFIED
	// goverter:enum:map UserStatusUnspecified UserStatus_USER_STATUS_UNSPECIFIED
	// goverter:enum:map UserStatusActive UserStatus_USER_STATUS_ACTIVE
	// goverter:enum:map UserStatusBlocked UserStatus_USER_STATUS_BLOCKED
	ToPBUserStatus(modeltiphereth.UserStatus) pb.UserStatus

	ToPBAccount(*modeltiphereth.Account) *librarian.Account
	ToPBAccountList([]*modeltiphereth.Account) []*librarian.Account

	// goverter:ignore FeatureSummary
	// goverter:autoMap PorterInstance
	ToPBPorter(*modelsupervisor.PorterInstanceController) *pb.Porter
	ToPBPorterList([]*modelsupervisor.PorterInstanceController) []*pb.Porter
	// goverter:enum:unknown PorterConnectionStatus_PORTER_CONNECTION_STATUS_UNSPECIFIED
	// goverter:enum:map PorterConnectionStatusUnspecified PorterConnectionStatus_PORTER_CONNECTION_STATUS_UNSPECIFIED
	// goverter:enum:map PorterConnectionStatusConnected PorterConnectionStatus_PORTER_CONNECTION_STATUS_CONNECTED
	// goverter:enum:map PorterConnectionStatusDisconnected PorterConnectionStatus_PORTER_CONNECTION_STATUS_DISCONNECTED
	// goverter:enum:map PorterConnectionStatusActive PorterConnectionStatus_PORTER_CONNECTION_STATUS_ACTIVE
	// goverter:enum:map PorterConnectionStatusActivationFailed PorterConnectionStatus_PORTER_CONNECTION_STATUS_ACTIVATION_FAILED
	// goverter:enum:map PorterConnectionStatusDowngraded PorterConnectionStatus_PORTER_CONNECTION_STATUS_DOWNGRADED
	ToPBPorterConnectionStatus(modelsupervisor.PorterConnectionStatus) pb.PorterConnectionStatus

	ToPBPorterContext(*modelsupervisor.PorterContext) *pb.PorterContext
	ToPBPorterContextList([]*modelsupervisor.PorterContext) []*pb.PorterContext
	// goverter:enum:unknown PorterContextStatus_PORTER_CONTEXT_STATUS_UNSPECIFIED
	// goverter:enum:map PorterContextStatusUnspecified PorterContextStatus_PORTER_CONTEXT_STATUS_UNSPECIFIED
	// goverter:enum:map PorterContextStatusActive PorterContextStatus_PORTER_CONTEXT_STATUS_ACTIVE
	// goverter:enum:map PorterContextStatusDisabled PorterContextStatus_PORTER_CONTEXT_STATUS_DISABLED
	ToPBPorterContextStatus(modelsupervisor.PorterContextStatus) pb.PorterContextStatus
	// goverter:enum:unknown PorterContextHandleStatus_PORTER_CONTEXT_HANDLE_STATUS_UNSPECIFIED
	// goverter:enum:map PorterContextHandleStatusUnspecified PorterContextHandleStatus_PORTER_CONTEXT_HANDLE_STATUS_UNSPECIFIED
	// goverter:enum:map PorterContextHandleStatusActive PorterContextHandleStatus_PORTER_CONTEXT_HANDLE_STATUS_ACTIVE
	// goverter:enum:map PorterContextHandleStatusDowngraded PorterContextHandleStatus_PORTER_CONTEXT_HANDLE_STATUS_DOWNGRADED
	// goverter:enum:map PorterContextHandleStatusQueueing PorterContextHandleStatus_PORTER_CONTEXT_HANDLE_STATUS_QUEUEING
	// goverter:enum:map PorterContextHandleStatusBlocked PorterContextHandleStatus_PORTER_CONTEXT_HANDLE_STATUS_BLOCKED
	ToPBPorterContextHandleStatus(modelsupervisor.PorterContextHandleStatus) pb.PorterContextHandleStatus

	ToPBPorterGroup(*modelsupervisor.PorterGroup) *pb.PorterGroup
	ToPBPorterGroupList([]*modelsupervisor.PorterGroup) []*pb.PorterGroup

	// goverter:ignore AltNames
	ToPBAppInfo(*modelgebura.AppInfo) *librarian.AppInfo
	// goverter:ignore ImageUrls
	ToPBAppInfoDetail(*modelgebura.AppInfoDetails) *librarian.AppInfoDetails
	ToPBAppInfoList([]*modelgebura.AppInfo) []*librarian.AppInfo
	// goverter:ignore AltNames
	ToPBAppInfoMixed(*modelgebura.AppInfoMixed) *librarian.AppInfoMixed
	ToPBAppInfoMixedList([]*modelgebura.AppInfoMixed) []*librarian.AppInfoMixed
	// goverter:enum:unknown AppType_APP_TYPE_UNSPECIFIED
	// goverter:enum:map AppTypeUnspecified AppType_APP_TYPE_UNSPECIFIED
	// goverter:enum:map AppTypeGame AppType_APP_TYPE_GAME
	ToPBAppType(modelgebura.AppType) librarian.AppType

	ToPBApp(*modelgebura.App) *pb.App
	ToPBAppList([]*modelgebura.App) []*pb.App
	// goverter:ignore Id
	// goverter:ignore TokenServerUrl
	// goverter:ignore Chunks
	ToPBAppBinary(*modelgebura.AppBinary) *pb.AppBinary

	ToPBAppInst(*modelgebura.AppInst) *pb.AppInst
	ToPBAppInstList([]*modelgebura.AppInst) []*pb.AppInst

	ToPBFeed(*modelfeed.Feed) *librarian.Feed
	ToPBFeedItem(*modelfeed.Item) *librarian.FeedItem
	ToPBFeedItemList([]*modelfeed.Item) []*librarian.FeedItem
	ToPBFeedImage(*modelfeed.Image) *librarian.FeedImage
	ToPBEnclosure(*modelfeed.Enclosure) *librarian.FeedEnclosure
	// goverter:map LatestPullStatus | ToPBFeedConfigPullStatus
	ToPBFeedConfig(*modelyesod.FeedConfig) *pb.FeedConfig
	// goverter:enum:unknown FeedConfigStatus_FEED_CONFIG_STATUS_UNSPECIFIED
	// goverter:enum:map FeedConfigStatusUnspecified FeedConfigStatus_FEED_CONFIG_STATUS_UNSPECIFIED
	// goverter:enum:map FeedConfigStatusActive FeedConfigStatus_FEED_CONFIG_STATUS_ACTIVE
	// goverter:enum:map FeedConfigStatusSuspend FeedConfigStatus_FEED_CONFIG_STATUS_SUSPEND
	ToPBFeedConfigStatus(modelyesod.FeedConfigStatus) pb.FeedConfigStatus
	// goverter:map FeedConfig Config
	ToPBFeedWithConfig(*modelyesod.FeedWithConfig) *pb.ListFeedConfigsResponse_FeedWithConfig
	ToPBFeedWithConfigList([]*modelyesod.FeedWithConfig) []*pb.ListFeedConfigsResponse_FeedWithConfig
	ToPBFeedItemDigest(*modelyesod.FeedItemDigest) *pb.FeedItemDigest
	ToPBFeedItemDigestList([]*modelyesod.FeedItemDigest) []*pb.FeedItemDigest

	ToPBFeedActionSet(*modelyesod.FeedActionSet) *pb.FeedActionSet
	ToPBFeedActionSetList([]*modelyesod.FeedActionSet) []*pb.FeedActionSet

	ToPBFeedItemCollection(*modelyesod.FeedItemCollection) *pb.FeedItemCollection
	ToPBFeedItemCollectionList([]*modelyesod.FeedItemCollection) []*pb.FeedItemCollection

	ToPBNotifyTarget(*modelnetzach.NotifyTarget) *pb.NotifyTarget
	ToPBNotifyTargetList([]*modelnetzach.NotifyTarget) []*pb.NotifyTarget
	// goverter:enum:unknown NotifyTargetStatus_NOTIFY_TARGET_STATUS_UNSPECIFIED
	// goverter:enum:map NotifyTargetStatusUnspecified NotifyTargetStatus_NOTIFY_TARGET_STATUS_UNSPECIFIED
	// goverter:enum:map NotifyTargetStatusActive NotifyTargetStatus_NOTIFY_TARGET_STATUS_ACTIVE
	// goverter:enum:map NotifyTargetStatusSuspend NotifyTargetStatus_NOTIFY_TARGET_STATUS_SUSPEND
	ToPBNotifyTargetStatus(modelnetzach.NotifyTargetStatus) pb.NotifyTargetStatus

	ToPBNotifyFlow(*modelnetzach.NotifyFlow) *pb.NotifyFlow
	// goverter:enum:unknown NotifyFlowStatus_NOTIFY_FLOW_STATUS_UNSPECIFIED
	// goverter:enum:map NotifyFlowStatusUnspecified NotifyFlowStatus_NOTIFY_FLOW_STATUS_UNSPECIFIED
	// goverter:enum:map NotifyFlowStatusActive NotifyFlowStatus_NOTIFY_FLOW_STATUS_ACTIVE
	// goverter:enum:map NotifyFlowStatusSuspend NotifyFlowStatus_NOTIFY_FLOW_STATUS_SUSPEND
	ToPBNotifyFlowStatus(modelnetzach.NotifyFlowStatus) pb.NotifyFlowStatus
	ToPBNotifyFlowSource(*modelnetzach.NotifyFlowSource) *pb.NotifyFlowSource
	ToPBNotifyFlowTarget(*modelnetzach.NotifyFlowTarget) *pb.NotifyFlowTarget
	ToPBNotifyFlowList([]*modelnetzach.NotifyFlow) []*pb.NotifyFlow

	ToPBSystemNotification(*modelnetzach.SystemNotification) *pb.SystemNotification
	ToPBSystemNotificationList([]*modelnetzach.SystemNotification) []*pb.SystemNotification
	// goverter:enum:unknown SystemNotificationType_SYSTEM_NOTIFICATION_TYPE_UNSPECIFIED
	// goverter:enum:map SystemNotificationTypeUnspecified SystemNotificationType_SYSTEM_NOTIFICATION_TYPE_UNSPECIFIED
	// goverter:enum:map SystemNotificationTypeSystem SystemNotificationType_SYSTEM_NOTIFICATION_TYPE_SYSTEM
	// goverter:enum:map SystemNotificationTypeUser SystemNotificationType_SYSTEM_NOTIFICATION_TYPE_USER
	ToPBSystemNotificationType(modelnetzach.SystemNotificationType) pb.SystemNotificationType
	// goverter:enum:unknown SystemNotificationLevel_SYSTEM_NOTIFICATION_LEVEL_UNSPECIFIED
	// goverter:enum:map SystemNotificationLevelUnspecified SystemNotificationLevel_SYSTEM_NOTIFICATION_LEVEL_UNSPECIFIED
	// goverter:enum:map SystemNotificationLevelInfo SystemNotificationLevel_SYSTEM_NOTIFICATION_LEVEL_INFO
	// goverter:enum:map SystemNotificationLevelWarning SystemNotificationLevel_SYSTEM_NOTIFICATION_LEVEL_WARNING
	// goverter:enum:map SystemNotificationLevelError SystemNotificationLevel_SYSTEM_NOTIFICATION_LEVEL_ERROR
	// goverter:enum:map SystemNotificationLevelOngoing SystemNotificationLevel_SYSTEM_NOTIFICATION_LEVEL_ONGOING
	ToPBSystemNotificationLevel(modelnetzach.SystemNotificationLevel) pb.SystemNotificationLevel
	// goverter:enum:unknown SystemNotificationStatus_SYSTEM_NOTIFICATION_STATUS_UNSPECIFIED
	// goverter:enum:map SystemNotificationStatusUnspecified SystemNotificationStatus_SYSTEM_NOTIFICATION_STATUS_UNSPECIFIED
	// goverter:enum:map SystemNotificationStatusUnread SystemNotificationStatus_SYSTEM_NOTIFICATION_STATUS_UNREAD
	// goverter:enum:map SystemNotificationStatusRead SystemNotificationStatus_SYSTEM_NOTIFICATION_STATUS_READ
	// goverter:enum:map SystemNotificationStatusDismissed SystemNotificationStatus_SYSTEM_NOTIFICATION_STATUS_DISMISSED
	ToPBSystemNotificationStatus(modelnetzach.SystemNotificationStatus) pb.SystemNotificationStatus
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

func ToPBFeedConfigPullStatus(s modelyesod.FeedConfigPullStatus) *pb.FeedConfigPullStatus {
	var status pb.FeedConfigPullStatus
	switch s {
	case modelyesod.FeedConfigPullStatusUnspecified:
		status = pb.FeedConfigPullStatus_FEED_CONFIG_PULL_STATUS_UNSPECIFIED
	case modelyesod.FeedConfigPullStatusProcessing:
		status = pb.FeedConfigPullStatus_FEED_CONFIG_PULL_STATUS_PROCESSING
	case modelyesod.FeedConfigPullStatusSuccess:
		status = pb.FeedConfigPullStatus_FEED_CONFIG_PULL_STATUS_SUCCESS
	case modelyesod.FeedConfigPullStatusFailed:
		status = pb.FeedConfigPullStatus_FEED_CONFIG_PULL_STATUS_FAILED
	default:
		status = pb.FeedConfigPullStatus_FEED_CONFIG_PULL_STATUS_UNSPECIFIED
	}
	return &status
}
