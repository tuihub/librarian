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
// goverter:enum:exclude time:Duration
// goverter:extend TimeToTime
// goverter:extend TimeToTimePtr
type toBizConverter interface { //nolint:unused // used by generator
	// goverter:matchIgnoreCase
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

	// goverter:matchIgnoreCase
	// goverter:ignore DeviceInfo
	// goverter:map CreatedAt CreateAt
	ToBizUserSession(*ent.UserSession) *modeltiphereth.UserSession
	ToBizUserSessionList([]*ent.UserSession) []*modeltiphereth.UserSession

	// goverter:matchIgnoreCase
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

	// goverter:matchIgnoreCase
	// goverter:map UpdatedAt LatestUpdateTime
	ToBizAccount(*ent.Account) *modeltiphereth.Account
	ToBizAccountList([]*ent.Account) []*modeltiphereth.Account

	// goverter:matchIgnoreCase
	// goverter:ignore ConnectionStatus
	ToBizPorter(*ent.PorterInstance) *modeltiphereth.PorterInstance
	ToBizPorterList([]*ent.PorterInstance) []*modeltiphereth.PorterInstance
	// goverter:enum:unknown PorterInstanceStatusUnspecified
	// goverter:enum:map StatusActive PorterInstanceStatusActive
	// goverter:enum:map StatusBlocked PorterInstanceStatusBlocked
	ToBizPorterStatus(porterinstance.Status) modeltiphereth.PorterInstanceStatus

	// goverter:matchIgnoreCase
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

	// goverter:matchIgnoreCase
	// goverter:ignore AssignedAppInfoID
	ToBizApp(*ent.App) *modelgebura.App
	ToBizAppList([]*ent.App) []*modelgebura.App
	ToBizAppBinary(ent.AppBinary) modelgebura.AppBinary

	// goverter:matchIgnoreCase
	ToBizAppInst(*ent.AppInst) *modelgebura.AppInst
	ToBizAppInstList([]*ent.AppInst) []*modelgebura.AppInst

	// goverter:matchIgnoreCase
	// goverter:map LatestPullAt LatestPullTime
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
