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
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// goverter:converter
// goverter:output:file ./generated.go
// goverter:output:package github.com/tuihub/librarian/app/sephirah/internal/model/converter
// goverter:extend ToBizInternalID
// goverter:extend ToBizTime
// goverter:extend ToBizDuration
// goverter:extend ToLibAuthUserType
// goverter:extend ToBizUserStatus
// goverter:extend ToBizAppType
// goverter:extend PtrToString
// goverter:extend ToBizAppPackageSource
// goverter:extend DurationPBToDuration
// goverter:extend ToBizFeedConfigStatus
// goverter:extend ToBizNotifyTargetStatus
// goverter:extend ToBizNotifyFlowStatus
// goverter:extend ToBizFileType
// goverter:extend ToBizAccountAppRelationType
// goverter:extend ToBizSystemType
// goverter:extend ToBizPorterConnectionStatus
type toBizConverter interface { //nolint:unused // used by generator
	ToBizTimeRange(*librarian.TimeRange) *model.TimeRange
	// goverter:matchIgnoreCase
	ToBizPorterFeatureSummary(*porter.PorterFeatureSummary) *modeltiphereth.PorterFeatureSummary

	ToBizInternalIDList(idl []*librarian.InternalID) []model.InternalID

	// goverter:matchIgnoreCase
	// goverter:map DeviceId ID
	ToBizDeviceInfo(*pb.DeviceInfo) *modeltiphereth.DeviceInfo
	// goverter:matchIgnoreCase
	ToBizUser(*pb.User) *modeltiphereth.User
	ToLibAuthUserTypeList([]pb.UserType) []libauth.UserType
	ToBizUserStatusList([]pb.UserStatus) []modeltiphereth.UserStatus

	// goverter:matchIgnoreCase
	ToBizPorterPrivilege(*pb.PorterPrivilege) *modeltiphereth.PorterInstancePrivilege

	// goverter:matchIgnoreCase
	// goverter:ignore BoundInternal
	// goverter:ignore LatestUpdateTime
	ToBizApp(*librarian.App) *modelgebura.App
	// goverter:matchIgnoreCase
	ToBizAppDetail(*librarian.AppDetails) *modelgebura.AppDetails
	ToBizAppTypeList([]librarian.AppType) []modelgebura.AppType
	// goverter:matchIgnoreCase
	ToBizAppID(*librarian.AppID) *modelgebura.AppID
	ToBizAppIDList([]*librarian.AppID) []*modelgebura.AppID

	// goverter:matchIgnoreCase
	ToBizAppPackage(*librarian.AppPackage) *modelgebura.AppPackage
	// goverter:matchIgnoreCase
	ToBizAppPackageBinary(*librarian.AppPackageBinary) *modelgebura.AppPackageBinary
	ToBizAppPackageBinaryList([]*librarian.AppPackageBinary) []*modelgebura.AppPackageBinary
	ToBizAppPackageSourceList([]librarian.AppPackageSource) []modelgebura.AppPackageSource

	// goverter:matchIgnoreCase
	// goverter:ignore LatestUpdateTime
	ToBizFeedConfig(*pb.FeedConfig) *modelyesod.FeedConfig
	ToBizFeedConfigStatusList([]pb.FeedConfigStatus) []modelyesod.FeedConfigStatus

	// goverter:matchIgnoreCase
	ToBizNotifyTarget(*pb.NotifyTarget) *modelnetzach.NotifyTarget
	ToBizNotifyTargetStatusList([]pb.NotifyTargetStatus) []modelnetzach.NotifyTargetStatus
	// goverter:matchIgnoreCase
	ToBizNotifyFlow(*pb.NotifyFlow) *modelnetzach.NotifyFlow
	// goverter:matchIgnoreCase
	ToBizNotifyFlowSource(*pb.NotifyFlowSource) *modelnetzach.NotifyFlowSource
	// goverter:matchIgnoreCase
	ToBizNotifyFlowTarget(*pb.NotifyFlowTarget) *modelnetzach.NotifyFlowTarget
	// goverter:matchIgnoreCase
	ToBizNotifyFilter(*pb.NotifyFilter) *modelnetzach.NotifyFilter

	// goverter:matchIgnoreCase
	ToBizFileMetadata(*pb.FileMetadata) *modelbinah.FileMetadata
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

func ToLibAuthUserType(u pb.UserType) libauth.UserType {
	switch u {
	case pb.UserType_USER_TYPE_UNSPECIFIED:
		return libauth.UserTypeUnspecified
	case pb.UserType_USER_TYPE_ADMIN:
		return libauth.UserTypeAdmin
	case pb.UserType_USER_TYPE_NORMAL:
		return libauth.UserTypeNormal
	case pb.UserType_USER_TYPE_SENTINEL:
		return libauth.UserTypeSentinel
	default:
		return libauth.UserTypeUnspecified
	}
}

func ToBizUserStatus(s pb.UserStatus) modeltiphereth.UserStatus {
	switch s {
	case pb.UserStatus_USER_STATUS_UNSPECIFIED:
		return modeltiphereth.UserStatusUnspecified
	case pb.UserStatus_USER_STATUS_ACTIVE:
		return modeltiphereth.UserStatusActive
	case pb.UserStatus_USER_STATUS_BLOCKED:
		return modeltiphereth.UserStatusBlocked
	default:
		return modeltiphereth.UserStatusUnspecified
	}
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

func ToBizAppType(t librarian.AppType) modelgebura.AppType {
	switch t {
	case librarian.AppType_APP_TYPE_UNSPECIFIED:
		return modelgebura.AppTypeUnspecified
	case librarian.AppType_APP_TYPE_GAME:
		return modelgebura.AppTypeGame
	default:
		return modelgebura.AppTypeUnspecified
	}
}

func ToBizAppPackageSource(a librarian.AppPackageSource) modelgebura.AppPackageSource {
	switch a {
	case librarian.AppPackageSource_APP_PACKAGE_SOURCE_UNSPECIFIED:
		return modelgebura.AppPackageSourceUnspecified
	case librarian.AppPackageSource_APP_PACKAGE_SOURCE_MANUAL:
		return modelgebura.AppPackageSourceManual
	case librarian.AppPackageSource_APP_PACKAGE_SOURCE_SENTINEL:
		return modelgebura.AppPackageSourceSentinel
	default:
		return modelgebura.AppPackageSourceUnspecified
	}
}

func ToBizFeedConfigStatus(s pb.FeedConfigStatus) modelyesod.FeedConfigStatus {
	switch s {
	case pb.FeedConfigStatus_FEED_CONFIG_STATUS_UNSPECIFIED:
		return modelyesod.FeedConfigStatusUnspecified
	case pb.FeedConfigStatus_FEED_CONFIG_STATUS_ACTIVE:
		return modelyesod.FeedConfigStatusActive
	case pb.FeedConfigStatus_FEED_CONFIG_STATUS_SUSPEND:
		return modelyesod.FeedConfigStatusSuspend
	default:
		return modelyesod.FeedConfigStatusUnspecified
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

func ToBizNotifyTargetStatus(s pb.NotifyTargetStatus) modelnetzach.NotifyTargetStatus {
	switch s {
	case pb.NotifyTargetStatus_NOTIFY_TARGET_STATUS_UNSPECIFIED:
		return modelnetzach.NotifyTargetStatusUnspecified
	case pb.NotifyTargetStatus_NOTIFY_TARGET_STATUS_ACTIVE:
		return modelnetzach.NotifyTargetStatusActive
	case pb.NotifyTargetStatus_NOTIFY_TARGET_STATUS_SUSPEND:
		return modelnetzach.NotifyTargetStatusSuspend
	default:
		return modelnetzach.NotifyTargetStatusUnspecified
	}
}

func ToBizNotifyFlowStatus(s pb.NotifyFlowStatus) modelnetzach.NotifyFlowStatus {
	switch s {
	case pb.NotifyFlowStatus_NOTIFY_FLOW_STATUS_UNSPECIFIED:
		return modelnetzach.NotifyFlowStatusUnspecified
	case pb.NotifyFlowStatus_NOTIFY_FLOW_STATUS_ACTIVE:
		return modelnetzach.NotifyFlowStatusActive
	case pb.NotifyFlowStatus_NOTIFY_FLOW_STATUS_SUSPEND:
		return modelnetzach.NotifyFlowStatusSuspend
	default:
		return modelnetzach.NotifyFlowStatusUnspecified
	}
}

func ToBizFileType(t pb.FileType) modelbinah.FileType {
	switch t {
	case pb.FileType_FILE_TYPE_UNSPECIFIED:
		return modelbinah.FileTypeUnspecified
	case pb.FileType_FILE_TYPE_GEBURA_SAVE:
		return modelbinah.FileTypeGeburaSave
	case pb.FileType_FILE_TYPE_CHESED_IMAGE:
		return modelbinah.FileTypeChesedImage
	default:
		return modelbinah.FileTypeUnspecified
	}
}

func ToBizAccountAppRelationType(t librarian.AccountAppRelationType) model.AccountAppRelationType {
	switch t {
	case librarian.AccountAppRelationType_ACCOUNT_APP_RELATION_TYPE_UNSPECIFIED:
		return model.AccountAppRelationTypeUnspecified
	case librarian.AccountAppRelationType_ACCOUNT_APP_RELATION_TYPE_OWN:
		return model.AccountAppRelationTypeOwner
	default:
		return model.AccountAppRelationTypeUnspecified
	}
}

func ToBizSystemType(t pb.SystemType) modeltiphereth.SystemType {
	switch t {
	case pb.SystemType_SYSTEM_TYPE_UNSPECIFIED:
		return modeltiphereth.SystemTypeUnspecified
	case pb.SystemType_SYSTEM_TYPE_IOS:
		return modeltiphereth.SystemTypeIOS
	case pb.SystemType_SYSTEM_TYPE_ANDROID:
		return modeltiphereth.SystemTypeAndroid
	case pb.SystemType_SYSTEM_TYPE_WEB:
		return modeltiphereth.SystemTypeWeb
	case pb.SystemType_SYSTEM_TYPE_WINDOWS:
		return modeltiphereth.SystemTypeWindows
	case pb.SystemType_SYSTEM_TYPE_MACOS:
		return modeltiphereth.SystemTypeMacOS
	case pb.SystemType_SYSTEM_TYPE_LINUX:
		return modeltiphereth.SystemTypeLinux
	default:
		return modeltiphereth.SystemTypeUnspecified
	}
}

func ToBizPorterConnectionStatus(s pb.PorterConnectionStatus) modeltiphereth.PorterConnectionStatus {
	switch s {
	case pb.PorterConnectionStatus_PORTER_CONNECTION_STATUS_UNSPECIFIED:
		return modeltiphereth.PorterConnectionStatusUnspecified
	case pb.PorterConnectionStatus_PORTER_CONNECTION_STATUS_CONNECTED:
		return modeltiphereth.PorterConnectionStatusConnected
	case pb.PorterConnectionStatus_PORTER_CONNECTION_STATUS_DISCONNECTED:
		return modeltiphereth.PorterConnectionStatusDisconnected
	case pb.PorterConnectionStatus_PORTER_CONNECTION_STATUS_ACTIVE:
		return modeltiphereth.PorterConnectionStatusActive
	case pb.PorterConnectionStatus_PORTER_CONNECTION_STATUS_ACTIVATION_FAILED:
		return modeltiphereth.PorterConnectionStatusActivationFailed
	default:
		return modeltiphereth.PorterConnectionStatusUnspecified
	}
}
