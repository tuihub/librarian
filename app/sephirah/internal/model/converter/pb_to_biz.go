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
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// goverter:converter
// goverter:extend ToBizInternalID
// goverter:extend ToBizTime
// goverter:extend ToBizDuration
// goverter:extend ToLibAuthUserType
// goverter:extend ToBizUserStatus
// goverter:extend ToBizAppType
// goverter:extend ToBizAppSource
// goverter:extend PtrToString
// goverter:extend ToBizAppPackageSource
// goverter:extend DurationPBToDuration
// goverter:extend ToBizFeedConfigSource
// goverter:extend ToBizFeedConfigStatus
// goverter:extend ToBizNotifyTargetStatus
// goverter:extend ToBizNotifyTargetType
// goverter:extend ToBizNotifyFlowStatus
// goverter:extend ToBizFileType
type toBizConverter interface { //nolint:unused // used by generator
	ToBizTimeRange(*librarian.TimeRange) *model.TimeRange

	ToBizInternalIDList(idl []*librarian.InternalID) []model.InternalID
	// goverter:matchIgnoreCase
	ToBizUser(*pb.User) *modeltiphereth.User
	ToLibAuthUserTypeList([]pb.UserType) []libauth.UserType
	ToBizUserStatusList([]pb.UserStatus) []modeltiphereth.UserStatus

	// goverter:matchIgnoreCase
	// goverter:ignore BoundInternal
	ToBizApp(*librarian.App) *modelgebura.App
	// goverter:matchIgnoreCase
	ToBizAppDetail(*librarian.AppDetails) *modelgebura.AppDetails
	ToBizAppTypeList([]librarian.AppType) []modelgebura.AppType
	ToBizAppSourceList([]librarian.AppSource) []modelgebura.AppSource

	// goverter:matchIgnoreCase
	ToBizAppPackage(*librarian.AppPackage) *modelgebura.AppPackage
	// goverter:matchIgnoreCase
	ToBizAppPackageBinary(*librarian.AppPackageBinary) *modelgebura.AppPackageBinary
	ToBizAppPackageBinaryList([]*librarian.AppPackageBinary) []*modelgebura.AppPackageBinary
	ToBizAppPackageSourceList([]librarian.AppPackageSource) []modelgebura.AppPackageSource

	// goverter:matchIgnoreCase
	// goverter:ignore LatestUpdateTime
	ToBizFeedConfig(*pb.FeedConfig) *modelyesod.FeedConfig
	ToBizFeedConfigSourceList([]pb.FeedConfigSource) []modelyesod.FeedConfigSource
	ToBizFeedConfigStatusList([]pb.FeedConfigStatus) []modelyesod.FeedConfigStatus

	// goverter:matchIgnoreCase
	ToBizNotifyTarget(*pb.NotifyTarget) *modelnetzach.NotifyTarget
	ToBizNotifyTargetTypeList([]pb.NotifyTargetType) []modelnetzach.NotifyTargetType
	ToBizNotifyTargetStatusList([]pb.NotifyTargetStatus) []modelnetzach.NotifyTargetStatus
	// goverter:matchIgnoreCase
	ToBizNotifyFlow(*pb.NotifyFlow) *modelnetzach.NotifyFlow
	// goverter:matchIgnoreCase
	ToBizNotifyFlowSource(*pb.NotifyFlowSource) *modelnetzach.NotifyFlowSource
	// goverter:matchIgnoreCase
	ToBizNotifyFlowTarget(*pb.NotifyFlowTarget) *modelnetzach.NotifyFlowTarget

	// goverter:matchIgnoreCase
	ToBizFileMetadata(*pb.FileMetadata) *modelbinah.FileMetadata
}

func ToBizInternalID(id *librarian.InternalID) model.InternalID {
	if id == nil {
		return 0
	}
	return model.InternalID(id.Id)
}

func ToBizInternalIDPtr(id *librarian.InternalID) *model.InternalID {
	if id == nil {
		return nil
	}
	i := model.InternalID(id.Id)
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

func ToBizAppSource(s librarian.AppSource) modelgebura.AppSource {
	switch s {
	case librarian.AppSource_APP_SOURCE_UNSPECIFIED:
		return modelgebura.AppSourceUnspecified
	case librarian.AppSource_APP_SOURCE_INTERNAL:
		return modelgebura.AppSourceInternal
	case librarian.AppSource_APP_SOURCE_STEAM:
		return modelgebura.AppSourceSteam
	default:
		return modelgebura.AppSourceUnspecified
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

func ToBizAccountPlatform(p librarian.AccountPlatform) modeltiphereth.AccountPlatform {
	switch p {
	case librarian.AccountPlatform_ACCOUNT_PLATFORM_UNSPECIFIED:
		return modeltiphereth.AccountPlatformUnspecified
	case librarian.AccountPlatform_ACCOUNT_PLATFORM_STEAM:
		return modeltiphereth.AccountPlatformSteam
	default:
		return modeltiphereth.AccountPlatformUnspecified
	}
}

func ToBizFeedConfigSource(s pb.FeedConfigSource) modelyesod.FeedConfigSource {
	switch s {
	case pb.FeedConfigSource_FEED_CONFIG_SOURCE_UNSPECIFIED:
		return modelyesod.FeedConfigSourceUnspecified
	case pb.FeedConfigSource_FEED_CONFIG_SOURCE_COMMON:
		return modelyesod.FeedConfigSourceCommon
	default:
		return modelyesod.FeedConfigSourceUnspecified
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

func ToBizNotifyTargetType(t pb.NotifyTargetType) modelnetzach.NotifyTargetType {
	switch t {
	case pb.NotifyTargetType_NOTIFY_TARGET_TYPE_UNSPECIFIED:
		return modelnetzach.NotifyTargetTypeUnspecified
	case pb.NotifyTargetType_NOTIFY_TARGET_TYPE_TELEGRAM:
		return modelnetzach.NotifyTargetTypeTelegram
	default:
		return modelnetzach.NotifyTargetTypeUnspecified
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
