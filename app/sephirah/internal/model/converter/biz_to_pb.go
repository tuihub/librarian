package converter

import (
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
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
// goverter:output:file ./generated.go
// goverter:output:package github.com/tuihub/librarian/app/sephirah/internal/model/converter
// goverter:extend ToPBInternalID
// goverter:extend ToPBTime
// goverter:extend ToPBTimePtr
// goverter:extend ToPBDuration
type toPBConverter interface { //nolint:unused // used by generator
	ToPBTimeRange(*model.TimeRange) *librarian.TimeRange
	ToPBInternalIDList([]model.InternalID) []*librarian.InternalID
	ToPBServerFeatureSummary(*modeltiphereth.ServerFeatureSummary) *pb.ServerFeatureSummary

	// goverter:matchIgnoreCase
	// goverter:map ID DeviceId
	// goverter:map SystemType | ToPBSystemType
	ToPBDeviceInfo(*modeltiphereth.DeviceInfo) *pb.DeviceInfo
	ToPBDeviceInfoList([]*modeltiphereth.DeviceInfo) []*pb.DeviceInfo

	// goverter:matchIgnoreCase
	// goverter:map CreateAt CreateTime
	// goverter:map ExpireAt ExpireTime
	ToPBUserSession(*modeltiphereth.UserSession) *pb.UserSession
	ToPBUserSessionList([]*modeltiphereth.UserSession) []*pb.UserSession

	// goverter:matchIgnoreCase
	// goverter:map Type | ToPBUserType
	// goverter:map Status | ToPBUserStatus
	// goverter:ignore Password
	ToPBUser(*modeltiphereth.User) *pb.User
	ToPBUserList([]*modeltiphereth.User) []*pb.User

	// goverter:matchIgnoreCase
	ToPBAccount(*modeltiphereth.Account) *librarian.Account
	ToPBAccountList([]*modeltiphereth.Account) []*librarian.Account

	// goverter:matchIgnoreCase
	// goverter:map Status | ToPBPorterStatus
	// goverter:map ConnectionStatus | ToPBPorterConnectionStatus
	// goverter:ignore FeatureSummary
	ToPBPorter(*modeltiphereth.PorterInstance) *pb.Porter
	ToPBPorterList([]*modeltiphereth.PorterInstance) []*pb.Porter

	// goverter:matchIgnoreCase
	// goverter:map Type | ToPBAppType
	// goverter:ignore AltNames
	ToPBAppInfo(*modelgebura.AppInfo) *librarian.AppInfo
	// goverter:matchIgnoreCase
	// goverter:ignore ImageUrls
	ToPBAppInfoDetail(*modelgebura.AppInfoDetails) *librarian.AppInfoDetails
	ToPBAppInfoList([]*modelgebura.AppInfo) []*librarian.AppInfo
	// goverter:matchIgnoreCase
	// goverter:map Type | ToPBAppType
	// goverter:ignore AltNames
	ToPBAppInfoMixed(*modelgebura.AppInfoMixed) *librarian.AppInfoMixed
	ToPBAppInfoMixedList([]*modelgebura.AppInfoMixed) []*librarian.AppInfoMixed

	// goverter:matchIgnoreCase
	ToPBApp(*modelgebura.App) *pb.App
	ToPBAppList([]*modelgebura.App) []*pb.App
	// goverter:matchIgnoreCase
	// goverter:ignore TokenServerUrl
	// goverter:ignore Chunks
	ToPBAppBinary(*modelgebura.AppBinary) *pb.AppBinary

	// goverter:matchIgnoreCase
	ToPBAppInst(*modelgebura.AppInst) *pb.AppInst
	ToPBAppInstList([]*modelgebura.AppInst) []*pb.AppInst

	// goverter:matchIgnoreCase
	ToPBFeed(*modelfeed.Feed) *librarian.Feed
	// goverter:matchIgnoreCase
	ToPBFeedItem(*modelfeed.Item) *librarian.FeedItem
	ToPBFeedItemList([]*modelfeed.Item) []*librarian.FeedItem
	// goverter:matchIgnoreCase
	ToPBFeedImage(*modelfeed.Image) *librarian.FeedImage
	// goverter:matchIgnoreCase
	ToPBEnclosure(*modelfeed.Enclosure) *librarian.FeedEnclosure
	// goverter:matchIgnoreCase
	// goverter:map Status | ToPBFeedConfigStatus
	ToPBFeedConfig(*modelyesod.FeedConfig) *pb.FeedConfig
	// goverter:matchIgnoreCase
	// goverter:map FeedConfig Config
	ToPBFeedWithConfig(*modelyesod.FeedWithConfig) *pb.ListFeedConfigsResponse_FeedWithConfig
	ToPBFeedWithConfigList([]*modelyesod.FeedWithConfig) []*pb.ListFeedConfigsResponse_FeedWithConfig
	// goverter:matchIgnoreCase
	ToPBFeedItemDigest(*modelyesod.FeedItemDigest) *pb.FeedItemDigest
	ToPBFeedItemDigestList([]*modelyesod.FeedItemDigest) []*pb.FeedItemDigest

	// goverter:matchIgnoreCase
	// goverter:map Status | ToPBNotifyTargetStatus
	ToPBNotifyTarget(*modelnetzach.NotifyTarget) *pb.NotifyTarget
	ToPBNotifyTargetList([]*modelnetzach.NotifyTarget) []*pb.NotifyTarget

	// goverter:matchIgnoreCase
	// goverter:map Status | ToPBNotifyFlowStatus
	ToPBNotifyFlow(*modelnetzach.NotifyFlow) *pb.NotifyFlow
	// goverter:matchIgnoreCase
	ToPBNotifyFlowSource(*modelnetzach.NotifyFlowSource) *pb.NotifyFlowSource
	// goverter:matchIgnoreCase
	ToPBNotifyFlowTarget(*modelnetzach.NotifyFlowTarget) *pb.NotifyFlowTarget
	ToPBNotifyFlowList([]*modelnetzach.NotifyFlow) []*pb.NotifyFlow
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

func ToPBUserType(u libauth.UserType) pb.UserType {
	switch u {
	case libauth.UserTypeUnspecified:
		return pb.UserType_USER_TYPE_UNSPECIFIED
	case libauth.UserTypeAdmin:
		return pb.UserType_USER_TYPE_ADMIN
	case libauth.UserTypeNormal:
		return pb.UserType_USER_TYPE_NORMAL
	case libauth.UserTypeSentinel:
		return pb.UserType_USER_TYPE_SENTINEL
	default:
		return pb.UserType_USER_TYPE_UNSPECIFIED
	}
}

func ToPBUserStatus(s modeltiphereth.UserStatus) pb.UserStatus {
	switch s {
	case modeltiphereth.UserStatusUnspecified:
		return pb.UserStatus_USER_STATUS_UNSPECIFIED
	case modeltiphereth.UserStatusActive:
		return pb.UserStatus_USER_STATUS_ACTIVE
	case modeltiphereth.UserStatusBlocked:
		return pb.UserStatus_USER_STATUS_BLOCKED
	default:
		return pb.UserStatus_USER_STATUS_UNSPECIFIED
	}
}

func ToPBPorterStatus(s modeltiphereth.PorterInstanceStatus) pb.UserStatus {
	switch s {
	case modeltiphereth.PorterInstanceStatusUnspecified:
		return pb.UserStatus_USER_STATUS_UNSPECIFIED
	case modeltiphereth.PorterInstanceStatusActive:
		return pb.UserStatus_USER_STATUS_ACTIVE
	case modeltiphereth.PorterInstanceStatusBlocked:
		return pb.UserStatus_USER_STATUS_BLOCKED
	default:
		return pb.UserStatus_USER_STATUS_UNSPECIFIED
	}
}

func ToPBAppType(t modelgebura.AppType) librarian.AppType {
	switch t {
	case modelgebura.AppTypeUnspecified:
		return librarian.AppType_APP_TYPE_UNSPECIFIED
	case modelgebura.AppTypeGame:
		return librarian.AppType_APP_TYPE_GAME
	default:
		return librarian.AppType_APP_TYPE_UNSPECIFIED
	}
}

func ToPBFeedConfigStatus(s modelyesod.FeedConfigStatus) pb.FeedConfigStatus {
	switch s {
	case modelyesod.FeedConfigStatusUnspecified:
		return pb.FeedConfigStatus_FEED_CONFIG_STATUS_UNSPECIFIED
	case modelyesod.FeedConfigStatusActive:
		return pb.FeedConfigStatus_FEED_CONFIG_STATUS_ACTIVE
	case modelyesod.FeedConfigStatusSuspend:
		return pb.FeedConfigStatus_FEED_CONFIG_STATUS_SUSPEND
	default:
		return pb.FeedConfigStatus_FEED_CONFIG_STATUS_UNSPECIFIED
	}
}

func ToPBNotifyTargetStatus(s modelnetzach.NotifyTargetStatus) pb.NotifyTargetStatus {
	switch s {
	case modelnetzach.NotifyTargetStatusUnspecified:
		return pb.NotifyTargetStatus_NOTIFY_TARGET_STATUS_UNSPECIFIED
	case modelnetzach.NotifyTargetStatusActive:
		return pb.NotifyTargetStatus_NOTIFY_TARGET_STATUS_ACTIVE
	case modelnetzach.NotifyTargetStatusSuspend:
		return pb.NotifyTargetStatus_NOTIFY_TARGET_STATUS_SUSPEND
	default:
		return pb.NotifyTargetStatus_NOTIFY_TARGET_STATUS_UNSPECIFIED
	}
}

func ToPBNotifyFlowStatus(s modelnetzach.NotifyFlowStatus) pb.NotifyFlowStatus {
	switch s {
	case modelnetzach.NotifyFlowStatusUnspecified:
		return pb.NotifyFlowStatus_NOTIFY_FLOW_STATUS_UNSPECIFIED
	case modelnetzach.NotifyFlowStatusActive:
		return pb.NotifyFlowStatus_NOTIFY_FLOW_STATUS_ACTIVE
	case modelnetzach.NotifyFlowStatusSuspend:
		return pb.NotifyFlowStatus_NOTIFY_FLOW_STATUS_SUSPEND
	default:
		return pb.NotifyFlowStatus_NOTIFY_FLOW_STATUS_UNSPECIFIED
	}
}

func ToPBSystemType(s modeltiphereth.SystemType) pb.SystemType {
	switch s {
	case modeltiphereth.SystemTypeUnspecified:
		return pb.SystemType_SYSTEM_TYPE_UNSPECIFIED
	case modeltiphereth.SystemTypeIOS:
		return pb.SystemType_SYSTEM_TYPE_IOS
	case modeltiphereth.SystemTypeAndroid:
		return pb.SystemType_SYSTEM_TYPE_ANDROID
	case modeltiphereth.SystemTypeWeb:
		return pb.SystemType_SYSTEM_TYPE_WEB
	case modeltiphereth.SystemTypeWindows:
		return pb.SystemType_SYSTEM_TYPE_WINDOWS
	case modeltiphereth.SystemTypeMacOS:
		return pb.SystemType_SYSTEM_TYPE_MACOS
	case modeltiphereth.SystemTypeLinux:
		return pb.SystemType_SYSTEM_TYPE_LINUX
	default:
		return pb.SystemType_SYSTEM_TYPE_UNSPECIFIED
	}
}

func ToPBPorterConnectionStatus(s modeltiphereth.PorterConnectionStatus) pb.PorterConnectionStatus {
	switch s {
	case modeltiphereth.PorterConnectionStatusUnspecified:
		return pb.PorterConnectionStatus_PORTER_CONNECTION_STATUS_UNSPECIFIED
	case modeltiphereth.PorterConnectionStatusConnected:
		return pb.PorterConnectionStatus_PORTER_CONNECTION_STATUS_CONNECTED
	case modeltiphereth.PorterConnectionStatusDisconnected:
		return pb.PorterConnectionStatus_PORTER_CONNECTION_STATUS_DISCONNECTED
	case modeltiphereth.PorterConnectionStatusActive:
		return pb.PorterConnectionStatus_PORTER_CONNECTION_STATUS_ACTIVE
	case modeltiphereth.PorterConnectionStatusActivationFailed:
		return pb.PorterConnectionStatus_PORTER_CONNECTION_STATUS_ACTIVATION_FAILED
	default:
		return pb.PorterConnectionStatus_PORTER_CONNECTION_STATUS_UNSPECIFIED
	}
}
