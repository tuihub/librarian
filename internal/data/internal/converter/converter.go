package converter

//go:generate go run github.com/jmattheis/goverter/cmd/goverter gen .

import (
	"time"

	"github.com/tuihub/librarian/internal/data/orm/model"
	libmodel "github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelchesed"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	"github.com/tuihub/librarian/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/model/modelnetzach"
	"github.com/tuihub/librarian/internal/model/modelsupervisor"
	"github.com/tuihub/librarian/internal/model/modelyesod"
)

// goverter:converter
// goverter:output:format function
// goverter:output:file ./generated.go
// goverter:output:package github.com/tuihub/librarian/internal/data/internal/converter
// goverter:matchIgnoreCase
// goverter:ignoreUnexported
// goverter:enum:exclude time:Duration
// goverter:extend TimeToTime
// goverter:extend TimeToTimePtr
// goverter:extend ToORMUserType
// goverter:extend ToORMUserStatus
// goverter:extend ToORMSystemType
// goverter:extend ToORMPorterInstanceStatus
// goverter:extend ToORMPorterConnectionStatus
// goverter:extend ToORMPorterContextStatus
// goverter:extend ToORMPorterContextHandleStatus
// goverter:extend ToORMSentinelSessionStatus
// goverter:extend ToORMAppInfoTypeManual
// goverter:extend ToORMFeedConfigStatus
// goverter:extend ToORMFeedConfigLatestPullStatus
// goverter:extend ToORMNotifySourceStatus
// goverter:extend ToORMNotifyTargetStatus
// goverter:extend ToORMSystemNotificationType
// goverter:extend ToORMSystemNotificationLevel
// goverter:extend ToORMSystemNotificationStatus
// goverter:extend ToORMImageStatus
// goverter:extend ToLibAuthUserType
// goverter:extend ToBizUserStatus
// goverter:extend ToBizSystemType
// goverter:extend ToBizPorterStatus
// goverter:extend ToBizPorterConnectionStatus
// goverter:extend ToBizPorterContextStatus
// goverter:extend ToBizPorterContextHandleStatus
// goverter:extend ToBizSentinelSessionStatus
// goverter:extend ToBizAppInfoType
// goverter:extend ToBizFeedConfigStatus
// goverter:extend ToBizFeedConfigPullStatus
// goverter:extend ToBizNotifyFlowStatus
// goverter:extend ToBizNotifyTargetStatus
// goverter:extend ToBizSystemNotificationType
// goverter:extend ToBizSystemNotificationLevel
// goverter:extend ToBizSystemNotificationStatus
// goverter:extend ToBizImageStatus
type toBizConverter interface { //nolint:unused // used by generator
	// goverter:ignore Password
	ToBizUser(*model.User) *libmodel.User
	ToBizUserList([]*model.User) []*libmodel.User

	// goverter:ignore Device
	// goverter:map CreatedAt CreateAt
	ToBizUserSession(*model.Session) *libmodel.Session
	ToBizUserSessionList([]*model.Session) []*libmodel.Session

	ToBizDeviceInfo(*model.Device) *libmodel.Device
	ToBizDeviceInfoList([]*model.Device) []*libmodel.Device

	// goverter:map UpdatedAt LatestUpdateTime
	ToBizAccount(*model.Account) *libmodel.Account
	ToBizAccountList([]*model.Account) []*libmodel.Account

	// goverter:map . BinarySummary
	ToBizPorter(*model.PorterInstance) *modelsupervisor.PorterInstance
	ToBizPorterList([]*model.PorterInstance) []*modelsupervisor.PorterInstance

	ToBizPorterContext(*model.PorterContext) *modelsupervisor.PorterContext
	ToBizPorterContextList([]*model.PorterContext) []*modelsupervisor.PorterContext

	// goverter:ignore Libraries
	ToBizSentinel(*model.Sentinel) *modelgebura.Sentinel
	ToBizSentinelList([]*model.Sentinel) []*modelgebura.Sentinel

	ToBizSentinelSession(*model.SentinelSession) *modelgebura.SentinelSession
	ToBizSentinelSessionList([]*model.SentinelSession) []*modelgebura.SentinelSession

	ToBizStoreApp(*model.StoreApp) *modelgebura.StoreApp
	ToBizStoreAppList([]*model.StoreApp) []*modelgebura.StoreApp

	// goverter:ignore AppID
	ToBizStoreAppBinary(*model.SentinelAppBinary) *modelgebura.StoreAppBinary
	ToBizStoreAppBinaryList([]*model.SentinelAppBinary) []*modelgebura.StoreAppBinary

	ToBizAppInfo(*model.AppInfo) *modelgebura.AppInfo
	ToBizAppInfoList([]*model.AppInfo) []*modelgebura.AppInfo

	ToBizApp(*model.App) *modelgebura.App
	ToBizAppList([]*model.App) []*modelgebura.App

	// goverter:map . RunTime
	ToBizAppRunTime(*model.AppRunTime) *modelgebura.AppRunTime
	ToBizAppRunTimeList([]*model.AppRunTime) []*modelgebura.AppRunTime
	// goverter:ignore AppIDs
	ToBizAppCategory(*model.AppCategory) *modelgebura.AppCategory

	// goverter:map LatestPullAt LatestPullTime
	// goverter:ignore ActionSets
	ToBizFeedConfig(*model.FeedConfig) *modelyesod.FeedConfig
	ToBizFeedConfigList([]*model.FeedConfig) []*modelyesod.FeedConfig

	ToBizFeedActionSet(*model.FeedActionSet) *modelyesod.FeedActionSet
	ToBizFeedActionSetList([]*model.FeedActionSet) []*modelyesod.FeedActionSet

	// goverter:ignore Items
	// goverter:ignore FeedType
	// goverter:ignore FeedVersion
	ToBizFeed(*model.Feed) *modelfeed.Feed
	ToBizFeedItem(*model.FeedItem) *modelfeed.Item
	ToBizFeedItemList([]*model.FeedItem) []*modelfeed.Item

	ToBizFeedItemCollection(*model.FeedItemCollection) *modelyesod.FeedItemCollection
	ToBizFeedItemCollectionList([]*model.FeedItemCollection) []*modelyesod.FeedItemCollection

	ToBizNotifyTarget(*model.NotifyTarget) *modelnetzach.NotifyTarget
	ToBizNotifyTargetList([]*model.NotifyTarget) []*modelnetzach.NotifyTarget

	// goverter:ignore Sources
	// goverter:ignore Targets
	ToBizNotifyFlow(*model.NotifyFlow) *modelnetzach.NotifyFlow

	// goverter:map CreatedAt CreateTime
	// goverter:map UpdatedAt UpdateTime
	ToBizSystemNotification(*model.SystemNotification) *modelnetzach.SystemNotification
	ToBizSystemNotificationList([]*model.SystemNotification) []*modelnetzach.SystemNotification

	ToBizImage(*model.Image) *modelchesed.Image
	ToBizImageList([]*model.Image) []*modelchesed.Image
}

// goverter:converter
// goverter:output:format function
// goverter:output:file ./generated.go
// goverter:output:package github.com/tuihub/librarian/internal/data/internal/converter
// goverter:matchIgnoreCase
// goverter:ignoreUnexported
// goverter:extend ToORMUserType
// goverter:extend ToORMUserStatus
// goverter:extend ToORMSystemType
// goverter:extend ToORMPorterInstanceStatus
// goverter:extend ToORMPorterConnectionStatus
// goverter:extend ToORMPorterContextStatus
// goverter:extend ToORMPorterContextHandleStatus
// goverter:extend ToORMSentinelSessionStatus
// goverter:extend ToORMAppInfoTypeManual
// goverter:extend ToORMFeedConfigStatus
// goverter:extend ToORMFeedConfigLatestPullStatus
// goverter:extend ToORMNotifySourceStatus
// goverter:extend ToORMNotifyTargetStatus
// goverter:extend ToORMSystemNotificationType
// goverter:extend ToORMSystemNotificationLevel
// goverter:extend ToORMSystemNotificationStatus
// goverter:extend ToORMImageStatus
type toORMConverter interface { //nolint:unused // used by generator
	// goverter:ignore CreatedAt
	// goverter:ignore UpdatedAt
	// goverter:ignore Sessions
	// goverter:ignore Account
	// goverter:ignore App
	// goverter:ignore FeedConfig
	// goverter:ignore NotifySource
	// goverter:ignore NotifyTarget
	// goverter:ignore NotifyFlow
	// goverter:ignore Image
	// goverter:ignore File
	// goverter:ignore Tag
	// goverter:ignore PorterContext
	// goverter:ignore CreatedUser
	// goverter:ignore Creator
	// goverter:ignore CreatorID
	ToORMUser(libmodel.User) model.User

	// goverter:ignore CreatedAt
	// goverter:ignore UpdatedAt
	ToORMAppInfo(modelgebura.AppInfo) model.AppInfo
}

func TimeToTime(t time.Time) time.Time {
	return t
}

func TimeToTimePtr(t *time.Time) *time.Time {
	return t
}
