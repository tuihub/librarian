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
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/porterinstance"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelchesed"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model/modelfeed"
)

// goverter:converter
// goverter:output:file ./generated.go
// goverter:output:package github.com/tuihub/librarian/app/sephirah/internal/data/internal/converter
// goverter:extend TimeToTime
// goverter:extend TimeToTimePtr
type toBizConverter interface { //nolint:unused // used by generator
	// goverter:matchIgnoreCase
	// goverter:map Type | ToLibAuthUserType
	// goverter:map Status | ToBizUserStatus
	// goverter:ignore PassWord
	ToBizUser(*ent.User) *modeltiphereth.User
	ToBizUserList([]*ent.User) []*modeltiphereth.User

	// goverter:matchIgnoreCase
	// goverter:ignore DeviceInfo
	// goverter:map CreatedAt CreateAt
	ToBizUserSession(*ent.UserSession) *modeltiphereth.UserSession
	ToBizUserSessionList([]*ent.UserSession) []*modeltiphereth.UserSession

	// goverter:matchIgnoreCase
	// goverter:map SystemType | ToBizSystemType
	ToBizDeviceInfo(*ent.DeviceInfo) *modeltiphereth.DeviceInfo
	ToBizDeviceInfoList([]*ent.DeviceInfo) []*modeltiphereth.DeviceInfo

	// goverter:matchIgnoreCase
	// goverter:map UpdatedAt LatestUpdateTime
	ToBizAccount(*ent.Account) *modeltiphereth.Account
	ToBizAccountList([]*ent.Account) []*modeltiphereth.Account

	// goverter:matchIgnoreCase
	// goverter:map Status | ToBizPorterStatus
	// goverter:ignore ConnectionStatus
	ToBizPorter(*ent.PorterInstance) *modeltiphereth.PorterInstance
	ToBizPorterList([]*ent.PorterInstance) []*modeltiphereth.PorterInstance

	// goverter:matchIgnoreCase
	// goverter:map Type | ToBizAppType
	// goverter:map . Details
	// goverter:map UpdatedAt LatestUpdateTime
	// goverter:ignore BoundInternal
	// goverter:ignore Tags
	ToBizAppInfo(*ent.AppInfo) *modelgebura.AppInfo
	ToBizAppInfoList([]*ent.AppInfo) []*modelgebura.AppInfo

	// goverter:matchIgnoreCase
	// goverter:ignore AssignedAppInfoID
	ToBizApp(*ent.App) *modelgebura.App
	ToBizAppList([]*ent.App) []*modelgebura.App
	ToBizAppBinary(ent.AppBinary) modelgebura.AppBinary

	// goverter:matchIgnoreCase
	ToBizAppInst(*ent.AppInst) *modelgebura.AppInst
	ToBizAppInstList([]*ent.AppInst) []*modelgebura.AppInst

	// goverter:matchIgnoreCase
	// goverter:map Status | ToBizFeedConfigStatus
	// goverter:map LatestPullAt LatestUpdateTime
	ToBizFeedConfig(*ent.FeedConfig) *modelyesod.FeedConfig
	ToBizFeedConfigList([]*ent.FeedConfig) []*modelyesod.FeedConfig

	// goverter:matchIgnoreCase
	// goverter:ignore Items
	// goverter:ignore FeedType
	// goverter:ignore FeedVersion
	ToBizFeed(*ent.Feed) *modelfeed.Feed
	// goverter:matchIgnoreCase
	ToBizFeedItem(*ent.FeedItem) *modelfeed.Item
	ToBizFeedItemList([]*ent.FeedItem) []*modelfeed.Item

	// goverter:matchIgnoreCase
	ToBizFeedItemCollection(*ent.FeedItemCollection) *modelyesod.FeedItemCollection
	ToBizFeedItemCollectionList([]*ent.FeedItemCollection) []*modelyesod.FeedItemCollection

	// goverter:map Status | ToBizNotifyTargetStatus
	ToBizNotifyTarget(*ent.NotifyTarget) *modelnetzach.NotifyTarget
	ToBizNotifyTargetList([]*ent.NotifyTarget) []*modelnetzach.NotifyTarget
	// goverter:ignore Sources
	// goverter:ignore Targets
	// goverter:map Status | ToBizNotifyFlowStatus
	ToBizNotifyFlow(*ent.NotifyFlow) *modelnetzach.NotifyFlow

	// goverter:map Status | ToBizImageStatus
	ToBizImage(*ent.Image) *modelchesed.Image
	ToBizImageList([]*ent.Image) []*modelchesed.Image
}

func TimeToTime(t time.Time) time.Time {
	return t
}

func TimeToTimePtr(t *time.Time) *time.Time {
	return t
}

func ToLibAuthUserType(t user.Type) libauth.UserType {
	switch t {
	case user.TypeAdmin:
		return libauth.UserTypeAdmin
	case user.TypeNormal:
		return libauth.UserTypeNormal
	case user.TypeSentinel:
		return libauth.UserTypeSentinel
	default:
		return libauth.UserTypeUnspecified
	}
}

func ToBizUserStatus(s user.Status) modeltiphereth.UserStatus {
	switch s {
	case user.StatusActive:
		return modeltiphereth.UserStatusActive
	case user.StatusBlocked:
		return modeltiphereth.UserStatusBlocked
	default:
		return modeltiphereth.UserStatusUnspecified
	}
}

func ToBizPorterStatus(s porterinstance.Status) modeltiphereth.PorterInstanceStatus {
	switch s {
	case porterinstance.StatusActive:
		return modeltiphereth.PorterInstanceStatusActive
	case porterinstance.StatusBlocked:
		return modeltiphereth.PorterInstanceStatusBlocked
	default:
		return modeltiphereth.PorterInstanceStatusUnspecified
	}
}

func ToBizAppType(t appinfo.Type) modelgebura.AppType {
	switch t {
	case appinfo.TypeUnknown:
		return modelgebura.AppTypeUnspecified
	case appinfo.TypeGame:
		return modelgebura.AppTypeGame
	default:
		return modelgebura.AppTypeUnspecified
	}
}

func ToBizFeedConfigStatus(s feedconfig.Status) modelyesod.FeedConfigStatus {
	switch s {
	case feedconfig.StatusActive:
		return modelyesod.FeedConfigStatusActive
	case feedconfig.StatusSuspend:
		return modelyesod.FeedConfigStatusSuspend
	default:
		return modelyesod.FeedConfigStatusUnspecified
	}
}

func ToBizNotifyTargetStatus(s notifytarget.Status) modelnetzach.NotifyTargetStatus {
	switch s {
	case notifytarget.StatusActive:
		return modelnetzach.NotifyTargetStatusActive
	case notifytarget.StatusSuspend:
		return modelnetzach.NotifyTargetStatusSuspend
	default:
		return modelnetzach.NotifyTargetStatusUnspecified
	}
}

func ToBizNotifyFlowStatus(s notifyflow.Status) modelnetzach.NotifyFlowStatus {
	switch s {
	case notifyflow.StatusActive:
		return modelnetzach.NotifyFlowStatusActive
	case notifyflow.StatusSuspend:
		return modelnetzach.NotifyFlowStatusSuspend
	default:
		return modelnetzach.NotifyFlowStatusUnspecified
	}
}

func ToBizImageStatus(s image.Status) modelchesed.ImageStatus {
	switch s {
	case image.StatusUploaded:
		return modelchesed.ImageStatusUploaded
	case image.StatusScanned:
		return modelchesed.ImageStatusScanned
	default:
		return modelchesed.ImageStatusUnspecified
	}
}

func ToBizSystemType(s deviceinfo.SystemType) modeltiphereth.SystemType {
	switch s {
	case deviceinfo.SystemTypeUnknown:
		return modeltiphereth.SystemTypeUnspecified
	case deviceinfo.SystemTypeIos:
		return modeltiphereth.SystemTypeIOS
	case deviceinfo.SystemTypeAndroid:
		return modeltiphereth.SystemTypeAndroid
	case deviceinfo.SystemTypeWeb:
		return modeltiphereth.SystemTypeWeb
	case deviceinfo.SystemTypeWindows:
		return modeltiphereth.SystemTypeWindows
	case deviceinfo.SystemTypeMacos:
		return modeltiphereth.SystemTypeMacOS
	case deviceinfo.SystemTypeLinux:
		return modeltiphereth.SystemTypeLinux
	default:
		return modeltiphereth.SystemTypeUnspecified
	}
}
