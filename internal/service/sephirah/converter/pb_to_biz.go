package converter

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelbinah"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	"github.com/tuihub/librarian/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/model/modelnetzach"
	"github.com/tuihub/librarian/internal/model/modelsupervisor"
	"github.com/tuihub/librarian/internal/model/modelyesod"
	sephirah "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
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
// goverter:extend ToBizInternalID
// goverter:extend ToBizInternalIDPtr
// goverter:extend ToBizTime
// goverter:extend ToBizDuration
// goverter:extend PtrToString
// goverter:extend DurationPBToDuration
type toBizConverter interface { //nolint:unused // used by generator
	ToBizTimeRange(*librarian.TimeRange) *model.TimeRange
	ToBizTimeRangeList([]*librarian.TimeRange) []*model.TimeRange
	ToBizPorterFeatureSummary(*librarian.FeatureSummary) *modelsupervisor.PorterFeatureSummary
	ToBizFeatureFlag(*librarian.FeatureFlag) *model.FeatureFlag
	ToBizFeatureRequest(*librarian.FeatureRequest) *model.FeatureRequest

	ToBizInternalIDList([]*librarian.InternalID) []model.InternalID

	// goverter:map DeviceId ID
	ToBizDeviceInfo(*sephirah.Device) *model.Device
	// goverter:enum:unknown SystemTypeUnspecified
	// goverter:enum:map SystemType_SYSTEM_TYPE_UNSPECIFIED SystemTypeUnspecified
	// goverter:enum:map SystemType_SYSTEM_TYPE_IOS SystemTypeIOS
	// goverter:enum:map SystemType_SYSTEM_TYPE_ANDROID SystemTypeAndroid
	// goverter:enum:map SystemType_SYSTEM_TYPE_WEB SystemTypeWeb
	// goverter:enum:map SystemType_SYSTEM_TYPE_WINDOWS SystemTypeWindows
	// goverter:enum:map SystemType_SYSTEM_TYPE_MACOS SystemTypeMacOS
	// goverter:enum:map SystemType_SYSTEM_TYPE_LINUX SystemTypeLinux
	ToBizSystemType(sephirah.SystemType) model.SystemType
	ToBizUser(*sephirah.User) *model.User
	// goverter:enum:unknown UserTypeUnspecified
	// goverter:enum:map UserType_USER_TYPE_UNSPECIFIED UserTypeUnspecified
	// goverter:enum:map UserType_USER_TYPE_ADMIN UserTypeAdmin
	// goverter:enum:map UserType_USER_TYPE_NORMAL UserTypeNormal
	ToLibAuthUserType(sephirah.UserType) model.UserType
	ToLibAuthUserTypeList([]sephirah.UserType) []model.UserType
	ToBizUserStatusList([]sephirah.UserStatus) []model.UserStatus
	// goverter:enum:unknown UserStatusUnspecified
	// goverter:enum:map UserStatus_USER_STATUS_UNSPECIFIED UserStatusUnspecified
	// goverter:enum:map UserStatus_USER_STATUS_ACTIVE UserStatusActive
	// goverter:enum:map UserStatus_USER_STATUS_BLOCKED UserStatusBlocked
	ToBizUserStatus(sephirah.UserStatus) model.UserStatus

	ToBizPorterContext(*sephirah.PorterContext) *modelsupervisor.PorterContext
	// goverter:enum:unknown PorterContextStatusUnspecified
	// goverter:enum:map PorterContextStatus_PORTER_CONTEXT_STATUS_UNSPECIFIED PorterContextStatusUnspecified
	// goverter:enum:map PorterContextStatus_PORTER_CONTEXT_STATUS_ACTIVE PorterContextStatusActive
	// goverter:enum:map PorterContextStatus_PORTER_CONTEXT_STATUS_DISABLED PorterContextStatusDisabled
	ToBizPorterContextStatus(sephirah.PorterContextStatus) modelsupervisor.PorterContextStatus
	// goverter:enum:unknown PorterContextHandleStatusUnspecified
	// goverter:enum:map PorterContextHandleStatus_PORTER_CONTEXT_HANDLE_STATUS_UNSPECIFIED PorterContextHandleStatusUnspecified
	// goverter:enum:map PorterContextHandleStatus_PORTER_CONTEXT_HANDLE_STATUS_ACTIVE PorterContextHandleStatusActive
	// goverter:enum:map PorterContextHandleStatus_PORTER_CONTEXT_HANDLE_STATUS_DOWNGRADED PorterContextHandleStatusDowngraded
	// goverter:enum:map PorterContextHandleStatus_PORTER_CONTEXT_HANDLE_STATUS_QUEUEING PorterContextHandleStatusQueueing
	// goverter:enum:map PorterContextHandleStatus_PORTER_CONTEXT_HANDLE_STATUS_BLOCKED PorterContextHandleStatusBlocked
	ToBizPorterContextHandleStatus(sephirah.PorterContextHandleStatus) modelsupervisor.PorterContextHandleStatus

	ToBizPorterBinarySummary(*librarian.PorterBinarySummary) *modelsupervisor.PorterBinarySummary

	// goverter:ignore ID
	// goverter:ignore ShortDescription
	// goverter:ignore ReleaseDate
	// goverter:map NameAlternatives AlternativeNames
	// goverter:ignore RawData
	// goverter:ignore UpdatedAt
	ToBizAppInfo(*sephirah.AppInfo) *modelgebura.AppInfo
	ToBizAppInfoList([]*sephirah.AppInfo) []*modelgebura.AppInfo
	ToBizAppTypeList([]sephirah.AppType) []modelgebura.AppType

	// goverter:ignore ShortDescription
	// goverter:ignore ReleaseDate
	// goverter:map NameAlternatives AlternativeNames
	ToBizApp(*sephirah.App) *modelgebura.App
	// goverter:enum:unknown AppTypeUnspecified
	// goverter:enum:map AppType_APP_TYPE_UNSPECIFIED AppTypeUnspecified
	// goverter:enum:map AppType_APP_TYPE_GAME AppTypeGame
	ToBizAppType(sephirah.AppType) modelgebura.AppType

	ToBizAppRunTime(*sephirah.AppRunTime) *modelgebura.AppRunTime
	ToBizAppRunTimeList([]*sephirah.AppRunTime) []*modelgebura.AppRunTime

	ToBizAppCategory(*sephirah.AppCategory) *modelgebura.AppCategory

	// ToBizAppInst(*sephirah.AppInst) *modelgebura.AppInst

	// goverter:ignore DigestDescription
	// goverter:ignore DigestImages
	ToBizFeedItem(*librarian.FeedItem) *modelfeed.Item

	// goverter:ignore LatestPullTime
	// goverter:useZeroValueOnPointerInconsistency
	ToBizFeedConfig(*sephirah.FeedConfig) *modelyesod.FeedConfig
	// goverter:enum:unknown FeedConfigStatusUnspecified
	// goverter:enum:map FeedConfigStatus_FEED_CONFIG_STATUS_UNSPECIFIED FeedConfigStatusUnspecified
	// goverter:enum:map FeedConfigStatus_FEED_CONFIG_STATUS_ACTIVE FeedConfigStatusActive
	// goverter:enum:map FeedConfigStatus_FEED_CONFIG_STATUS_SUSPEND FeedConfigStatusSuspend
	ToBizFeedConfigStatus(sephirah.FeedConfigStatus) modelyesod.FeedConfigStatus
	// goverter:enum:unknown FeedConfigPullStatusUnspecified
	// goverter:enum:map FeedConfigPullStatus_FEED_CONFIG_PULL_STATUS_UNSPECIFIED FeedConfigPullStatusUnspecified
	// goverter:enum:map FeedConfigPullStatus_FEED_CONFIG_PULL_STATUS_PROCESSING FeedConfigPullStatusProcessing
	// goverter:enum:map FeedConfigPullStatus_FEED_CONFIG_PULL_STATUS_SUCCESS FeedConfigPullStatusSuccess
	// goverter:enum:map FeedConfigPullStatus_FEED_CONFIG_PULL_STATUS_FAILED FeedConfigPullStatusFailed
	ToBizFeedConfigPullStatus(sephirah.FeedConfigPullStatus) modelyesod.FeedConfigPullStatus
	ToBizFeedConfigStatusList([]sephirah.FeedConfigStatus) []modelyesod.FeedConfigStatus

	ToBizFeedActionSet(*sephirah.FeedActionSet) *modelyesod.FeedActionSet

	ToBizFeedItemCollection(*sephirah.FeedItemCollection) *modelyesod.FeedItemCollection

	ToBizNotifyTarget(*sephirah.NotifyTarget) *modelnetzach.NotifyTarget
	// goverter:enum:unknown NotifyTargetStatusUnspecified
	// goverter:enum:map NotifyTargetStatus_NOTIFY_TARGET_STATUS_UNSPECIFIED NotifyTargetStatusUnspecified
	// goverter:enum:map NotifyTargetStatus_NOTIFY_TARGET_STATUS_ACTIVE NotifyTargetStatusActive
	// goverter:enum:map NotifyTargetStatus_NOTIFY_TARGET_STATUS_SUSPEND NotifyTargetStatusSuspend
	ToBizNotifyTargetStatus(sephirah.NotifyTargetStatus) modelnetzach.NotifyTargetStatus
	ToBizNotifyTargetStatusList([]sephirah.NotifyTargetStatus) []modelnetzach.NotifyTargetStatus
	ToBizNotifyFlow(*sephirah.NotifyFlow) *modelnetzach.NotifyFlow
	// goverter:enum:unknown NotifyFlowStatusUnspecified
	// goverter:enum:map NotifyFlowStatus_NOTIFY_FLOW_STATUS_UNSPECIFIED NotifyFlowStatusUnspecified
	// goverter:enum:map NotifyFlowStatus_NOTIFY_FLOW_STATUS_ACTIVE NotifyFlowStatusActive
	// goverter:enum:map NotifyFlowStatus_NOTIFY_FLOW_STATUS_SUSPEND NotifyFlowStatusSuspend
	ToBizNotifyFlowStatus(sephirah.NotifyFlowStatus) modelnetzach.NotifyFlowStatus
	ToBizNotifyFlowSource(*sephirah.NotifyFlowSource) *modelnetzach.NotifyFlowSource
	ToBizNotifyFlowTarget(*sephirah.NotifyFlowTarget) *modelnetzach.NotifyFlowTarget
	ToBizNotifyFilter(*sephirah.NotifyFilter) *modelnetzach.NotifyFilter

	// goverter:enum:unknown SystemNotificationLevelUnspecified
	// goverter:enum:map SystemNotificationLevel_SYSTEM_NOTIFICATION_LEVEL_UNSPECIFIED SystemNotificationLevelUnspecified
	// goverter:enum:map SystemNotificationLevel_SYSTEM_NOTIFICATION_LEVEL_INFO SystemNotificationLevelInfo
	// goverter:enum:map SystemNotificationLevel_SYSTEM_NOTIFICATION_LEVEL_WARNING SystemNotificationLevelWarning
	// goverter:enum:map SystemNotificationLevel_SYSTEM_NOTIFICATION_LEVEL_ERROR SystemNotificationLevelError
	// goverter:enum:map SystemNotificationLevel_SYSTEM_NOTIFICATION_LEVEL_ONGOING SystemNotificationLevelOngoing
	ToBizSystemNotificationLevel(sephirah.SystemNotificationLevel) modelnetzach.SystemNotificationLevel
	ToBizSystemNotificationLevelList([]sephirah.SystemNotificationLevel) []modelnetzach.SystemNotificationLevel
	// goverter:enum:unknown SystemNotificationStatusUnspecified
	// goverter:enum:map SystemNotificationStatus_SYSTEM_NOTIFICATION_STATUS_UNSPECIFIED SystemNotificationStatusUnspecified
	// goverter:enum:map SystemNotificationStatus_SYSTEM_NOTIFICATION_STATUS_UNREAD SystemNotificationStatusUnread
	// goverter:enum:map SystemNotificationStatus_SYSTEM_NOTIFICATION_STATUS_READ SystemNotificationStatusRead
	// goverter:enum:map SystemNotificationStatus_SYSTEM_NOTIFICATION_STATUS_DISMISSED SystemNotificationStatusDismissed
	ToBizSystemNotificationStatus(sephirah.SystemNotificationStatus) modelnetzach.SystemNotificationStatus
	ToBizSystemNotificationStatusList([]sephirah.SystemNotificationStatus) []modelnetzach.SystemNotificationStatus

	ToBizFileMetadata(*librarian.FileMetadata) *modelbinah.FileMetadata
	// goverter:enum:unknown FileTypeUnspecified
	// goverter:enum:map FileType_FILE_TYPE_UNSPECIFIED FileTypeUnspecified
	// goverter:enum:map FileType_FILE_TYPE_GEBURA_SAVE FileTypeGeburaSave
	// goverter:enum:map FileType_FILE_TYPE_CHESED_IMAGE FileTypeChesedImage
	// goverter:enum:map FileType_FILE_TYPE_GEBURA_APP_INFO_IMAGE FileTypeGeburaAppInfoImage
	ToBizFileType(librarian.FileType) modelbinah.FileType
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
