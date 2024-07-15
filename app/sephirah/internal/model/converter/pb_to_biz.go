package converter

import (
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modelbinah"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// goverter:converter
// goverter:output:file ./generated.go
// goverter:matchIgnoreCase
// goverter:output:package github.com/tuihub/librarian/app/sephirah/internal/model/converter
// goverter:extend ToBizInternalID
// goverter:extend ToBizTime
// goverter:extend ToBizDuration
// goverter:extend PtrToString
// goverter:extend DurationPBToDuration
type toBizConverter interface { //nolint:unused // used by generator
	ToBizTimeRange(*librarian.TimeRange) *model.TimeRange
	ToBizPorterFeatureSummary(*porter.PorterFeatureSummary) *modeltiphereth.PorterFeatureSummary
	ToBizFeatureFlag(*librarian.FeatureFlag) *modeltiphereth.FeatureFlag
	ToBizFeatureRequest(*librarian.FeatureRequest) *modeltiphereth.FeatureRequest
	// goverter:enum:unknown AccountAppRelationTypeUnspecified
	// goverter:enum:map AccountAppRelationType_ACCOUNT_APP_RELATION_TYPE_UNSPECIFIED AccountAppRelationTypeUnspecified
	// goverter:enum:map AccountAppRelationType_ACCOUNT_APP_RELATION_TYPE_OWN AccountAppRelationTypeOwner
	ToBizAccountAppRelationType(librarian.AccountAppRelationType) model.AccountAppRelationType

	ToBizInternalIDList([]*librarian.InternalID) []model.InternalID

	// goverter:map DeviceId ID
	ToBizDeviceInfo(*pb.DeviceInfo) *modeltiphereth.DeviceInfo
	// goverter:enum:unknown SystemTypeUnspecified
	// goverter:enum:map SystemType_SYSTEM_TYPE_UNSPECIFIED SystemTypeUnspecified
	// goverter:enum:map SystemType_SYSTEM_TYPE_IOS SystemTypeIOS
	// goverter:enum:map SystemType_SYSTEM_TYPE_ANDROID SystemTypeAndroid
	// goverter:enum:map SystemType_SYSTEM_TYPE_WEB SystemTypeWeb
	// goverter:enum:map SystemType_SYSTEM_TYPE_WINDOWS SystemTypeWindows
	// goverter:enum:map SystemType_SYSTEM_TYPE_MACOS SystemTypeMacOS
	// goverter:enum:map SystemType_SYSTEM_TYPE_LINUX SystemTypeLinux
	ToBizSystemType(pb.SystemType) modeltiphereth.SystemType
	ToBizUser(*pb.User) *modeltiphereth.User
	// goverter:enum:unknown UserTypeUnspecified
	// goverter:enum:map UserType_USER_TYPE_UNSPECIFIED UserTypeUnspecified
	// goverter:enum:map UserType_USER_TYPE_ADMIN UserTypeAdmin
	// goverter:enum:map UserType_USER_TYPE_NORMAL UserTypeNormal
	// goverter:enum:map UserType_USER_TYPE_SENTINEL UserTypeSentinel
	// goverter:enum:map UserType_USER_TYPE_PORTER UserTypePorter
	ToLibAuthUserType(pb.UserType) libauth.UserType
	ToLibAuthUserTypeList([]pb.UserType) []libauth.UserType
	ToBizUserStatusList([]pb.UserStatus) []modeltiphereth.UserStatus
	// goverter:enum:unknown UserStatusUnspecified
	// goverter:enum:map UserStatus_USER_STATUS_UNSPECIFIED UserStatusUnspecified
	// goverter:enum:map UserStatus_USER_STATUS_ACTIVE UserStatusActive
	// goverter:enum:map UserStatus_USER_STATUS_BLOCKED UserStatusBlocked
	ToBizUserStatus(pb.UserStatus) modeltiphereth.UserStatus

	ToBizPorterContext(*pb.PorterContext) *modeltiphereth.PorterInstanceContext

	// goverter:ignore BoundInternal
	// goverter:ignore LatestUpdateTime
	ToBizAppInfo(*librarian.AppInfo) *modelgebura.AppInfo
	ToBizAppInfoList([]*librarian.AppInfo) []*modelgebura.AppInfo
	ToBizAppInfoDetail(*librarian.AppInfoDetails) *modelgebura.AppInfoDetails
	ToBizAppTypeList([]librarian.AppType) []modelgebura.AppType
	ToBizAppInfoID(*librarian.AppInfoID) *modelgebura.AppInfoID
	ToBizAppInfoIDList([]*librarian.AppInfoID) []*modelgebura.AppInfoID

	ToBizApp(*pb.App) *modelgebura.App
	ToBizAppBinary(*pb.AppBinary) *modelgebura.AppBinary
	ToBizAppBinaryList([]*pb.AppBinary) []*modelgebura.AppBinary
	// goverter:enum:unknown AppTypeUnspecified
	// goverter:enum:map AppType_APP_TYPE_UNSPECIFIED AppTypeUnspecified
	// goverter:enum:map AppType_APP_TYPE_GAME AppTypeGame
	ToBizAppType(librarian.AppType) modelgebura.AppType

	ToBizAppInst(*pb.AppInst) *modelgebura.AppInst

	// goverter:ignore DigestDescription
	// goverter:ignore DigestImages
	ToBizFeedItem(*librarian.FeedItem) *modelfeed.Item

	// goverter:ignore LatestPullTime
	// goverter:useZeroValueOnPointerInconsistency
	ToBizFeedConfig(*pb.FeedConfig) *modelyesod.FeedConfig
	// goverter:enum:unknown FeedConfigStatusUnspecified
	// goverter:enum:map FeedConfigStatus_FEED_CONFIG_STATUS_UNSPECIFIED FeedConfigStatusUnspecified
	// goverter:enum:map FeedConfigStatus_FEED_CONFIG_STATUS_ACTIVE FeedConfigStatusActive
	// goverter:enum:map FeedConfigStatus_FEED_CONFIG_STATUS_SUSPEND FeedConfigStatusSuspend
	ToBizFeedConfigStatus(pb.FeedConfigStatus) modelyesod.FeedConfigStatus
	// goverter:enum:unknown FeedConfigPullStatusUnspecified
	// goverter:enum:map FeedConfigPullStatus_FEED_CONFIG_PULL_STATUS_UNSPECIFIED FeedConfigPullStatusUnspecified
	// goverter:enum:map FeedConfigPullStatus_FEED_CONFIG_PULL_STATUS_PROCESSING FeedConfigPullStatusProcessing
	// goverter:enum:map FeedConfigPullStatus_FEED_CONFIG_PULL_STATUS_SUCCESS FeedConfigPullStatusSuccess
	// goverter:enum:map FeedConfigPullStatus_FEED_CONFIG_PULL_STATUS_FAILED FeedConfigPullStatusFailed
	ToBizFeedConfigPullStatus(pb.FeedConfigPullStatus) modelyesod.FeedConfigPullStatus
	ToBizFeedConfigStatusList([]pb.FeedConfigStatus) []modelyesod.FeedConfigStatus

	ToBizFeedActionSet(*pb.FeedActionSet) *modelyesod.FeedActionSet

	ToBizFeedItemCollection(*pb.FeedItemCollection) *modelyesod.FeedItemCollection

	ToBizNotifyTarget(*pb.NotifyTarget) *modelnetzach.NotifyTarget
	// goverter:enum:unknown NotifyTargetStatusUnspecified
	// goverter:enum:map NotifyTargetStatus_NOTIFY_TARGET_STATUS_UNSPECIFIED NotifyTargetStatusUnspecified
	// goverter:enum:map NotifyTargetStatus_NOTIFY_TARGET_STATUS_ACTIVE NotifyTargetStatusActive
	// goverter:enum:map NotifyTargetStatus_NOTIFY_TARGET_STATUS_SUSPEND NotifyTargetStatusSuspend
	ToBizNotifyTargetStatus(pb.NotifyTargetStatus) modelnetzach.NotifyTargetStatus
	ToBizNotifyTargetStatusList([]pb.NotifyTargetStatus) []modelnetzach.NotifyTargetStatus
	ToBizNotifyFlow(*pb.NotifyFlow) *modelnetzach.NotifyFlow
	// goverter:enum:unknown NotifyFlowStatusUnspecified
	// goverter:enum:map NotifyFlowStatus_NOTIFY_FLOW_STATUS_UNSPECIFIED NotifyFlowStatusUnspecified
	// goverter:enum:map NotifyFlowStatus_NOTIFY_FLOW_STATUS_ACTIVE NotifyFlowStatusActive
	// goverter:enum:map NotifyFlowStatus_NOTIFY_FLOW_STATUS_SUSPEND NotifyFlowStatusSuspend
	ToBizNotifyFlowStatus(pb.NotifyFlowStatus) modelnetzach.NotifyFlowStatus
	ToBizNotifyFlowSource(*pb.NotifyFlowSource) *modelnetzach.NotifyFlowSource
	ToBizNotifyFlowTarget(*pb.NotifyFlowTarget) *modelnetzach.NotifyFlowTarget
	ToBizNotifyFilter(*pb.NotifyFilter) *modelnetzach.NotifyFilter

	// goverter:enum:unknown SystemNotificationTypeUnspecified
	// goverter:enum:map SystemNotificationType_SYSTEM_NOTIFICATION_TYPE_UNSPECIFIED SystemNotificationTypeUnspecified
	// goverter:enum:map SystemNotificationType_SYSTEM_NOTIFICATION_TYPE_SYSTEM SystemNotificationTypeSystem
	// goverter:enum:map SystemNotificationType_SYSTEM_NOTIFICATION_TYPE_USER SystemNotificationTypeUser
	ToBizSystemNotificationType(pb.SystemNotificationType) modelnetzach.SystemNotificationType
	ToBizSystemNotificationTypeList([]pb.SystemNotificationType) []modelnetzach.SystemNotificationType
	// goverter:enum:unknown SystemNotificationLevelUnspecified
	// goverter:enum:map SystemNotificationLevel_SYSTEM_NOTIFICATION_LEVEL_UNSPECIFIED SystemNotificationLevelUnspecified
	// goverter:enum:map SystemNotificationLevel_SYSTEM_NOTIFICATION_LEVEL_INFO SystemNotificationLevelInfo
	// goverter:enum:map SystemNotificationLevel_SYSTEM_NOTIFICATION_LEVEL_WARNING SystemNotificationLevelWarning
	// goverter:enum:map SystemNotificationLevel_SYSTEM_NOTIFICATION_LEVEL_ERROR SystemNotificationLevelError
	// goverter:enum:map SystemNotificationLevel_SYSTEM_NOTIFICATION_LEVEL_ONGOING SystemNotificationLevelOngoing
	ToBizSystemNotificationLevel(pb.SystemNotificationLevel) modelnetzach.SystemNotificationLevel
	ToBizSystemNotificationLevelList([]pb.SystemNotificationLevel) []modelnetzach.SystemNotificationLevel
	// goverter:enum:unknown SystemNotificationStatusUnspecified
	// goverter:enum:map SystemNotificationStatus_SYSTEM_NOTIFICATION_STATUS_UNSPECIFIED SystemNotificationStatusUnspecified
	// goverter:enum:map SystemNotificationStatus_SYSTEM_NOTIFICATION_STATUS_UNREAD SystemNotificationStatusUnread
	// goverter:enum:map SystemNotificationStatus_SYSTEM_NOTIFICATION_STATUS_READ SystemNotificationStatusRead
	// goverter:enum:map SystemNotificationStatus_SYSTEM_NOTIFICATION_STATUS_DISMISSED SystemNotificationStatusDismissed
	ToBizSystemNotificationStatus(pb.SystemNotificationStatus) modelnetzach.SystemNotificationStatus
	ToBizSystemNotificationStatusList([]pb.SystemNotificationStatus) []modelnetzach.SystemNotificationStatus

	ToBizFileMetadata(*pb.FileMetadata) *modelbinah.FileMetadata
	// goverter:enum:unknown FileTypeUnspecified
	// goverter:enum:map FileType_FILE_TYPE_UNSPECIFIED FileTypeUnspecified
	// goverter:enum:map FileType_FILE_TYPE_GEBURA_SAVE FileTypeGeburaSave
	// goverter:enum:map FileType_FILE_TYPE_CHESED_IMAGE FileTypeChesedImage
	ToBizFileType(pb.FileType) modelbinah.FileType
}

func PtrToString(u *string) string {
	if u == nil {
		return ""
	}
	return *u
}

func ToBizInternalID(id *librarian.InternalID) model.InternalID {
	if id == nil {
		return 0
	}
	return model.InternalID(id.GetId())
}

func ToBizInternalIDPtr(id *librarian.InternalID) *model.InternalID {
	if id == nil {
		return nil
	}
	i := model.InternalID(id.GetId())
	return &i
}

func ToBizTime(t *timestamppb.Timestamp) time.Time {
	return t.AsTime()
}

func ToBizDuration(d *durationpb.Duration) time.Duration {
	return d.AsDuration()
}

func ToBizPorterStatus(s pb.UserStatus) modeltiphereth.PorterInstanceStatus {
	switch s {
	case pb.UserStatus_USER_STATUS_UNSPECIFIED:
		return modeltiphereth.PorterInstanceStatusUnspecified
	case pb.UserStatus_USER_STATUS_ACTIVE:
		return modeltiphereth.PorterInstanceStatusActive
	case pb.UserStatus_USER_STATUS_BLOCKED:
		return modeltiphereth.PorterInstanceStatusBlocked
	default:
		return modeltiphereth.PorterInstanceStatusUnspecified
	}
}

func ToBizGroupFeedItemsBy(by librarian.TimeAggregation_AggregationType) modelyesod.GroupFeedItemsBy {
	switch by {
	case librarian.TimeAggregation_AGGREGATION_TYPE_UNSPECIFIED:
		return modelyesod.GroupFeedItemsByUnspecified
	case librarian.TimeAggregation_AGGREGATION_TYPE_YEAR:
		return modelyesod.GroupFeedItemsByYear
	case librarian.TimeAggregation_AGGREGATION_TYPE_MONTH:
		return modelyesod.GroupFeedItemsByMonth
	case librarian.TimeAggregation_AGGREGATION_TYPE_DAY:
		return modelyesod.GroupFeedItemsByDay
	case librarian.TimeAggregation_AGGREGATION_TYPE_OVERALL:
		return modelyesod.GroupFeedItemsByOverall
	default:
		return modelyesod.GroupFeedItemsByUnspecified
	}
}
