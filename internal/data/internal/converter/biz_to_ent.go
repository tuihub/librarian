package converter

import (
	"github.com/tuihub/librarian/internal/data/internal/ent"
	"github.com/tuihub/librarian/internal/data/internal/ent/app"
	"github.com/tuihub/librarian/internal/data/internal/ent/appinfo"
	"github.com/tuihub/librarian/internal/data/internal/ent/device"
	"github.com/tuihub/librarian/internal/data/internal/ent/feedconfig"
	"github.com/tuihub/librarian/internal/data/internal/ent/image"
	"github.com/tuihub/librarian/internal/data/internal/ent/notifyflow"
	"github.com/tuihub/librarian/internal/data/internal/ent/notifytarget"
	"github.com/tuihub/librarian/internal/data/internal/ent/portercontext"
	"github.com/tuihub/librarian/internal/data/internal/ent/porterinstance"
	"github.com/tuihub/librarian/internal/data/internal/ent/systemnotification"
	"github.com/tuihub/librarian/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelchesed"
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
type toEntConverter interface { //nolint:unused // used by generator
	// goverter:enum:unknown @ignore
	// goverter:enum:map UserTypeUnspecified @ignore
	// goverter:enum:map UserTypeAdmin TypeAdmin
	// goverter:enum:map UserTypeNormal TypeNormal
	// goverter:enum:map UserTypeSentinel @ignore
	// goverter:enum:map UserTypePorter @ignore
	ToEntUserType(model.UserType) user.Type
	ToEntUserTypeList([]model.UserType) []user.Type
	// goverter:enum:unknown @ignore
	// goverter:enum:map UserStatusUnspecified @ignore
	// goverter:enum:map UserStatusActive StatusActive
	// goverter:enum:map UserStatusBlocked StatusBlocked
	ToEntUserStatus(model.UserStatus) user.Status
	ToEntUserStatusList([]model.UserStatus) []user.Status

	// goverter:enum:unknown @ignore
	// goverter:enum:map SystemTypeUnspecified @ignore
	// goverter:enum:map SystemTypeIOS SystemTypeIos
	// goverter:enum:map SystemTypeAndroid SystemTypeAndroid
	// goverter:enum:map SystemTypeWindows SystemTypeWindows
	// goverter:enum:map SystemTypeMacOS SystemTypeMacos
	// goverter:enum:map SystemTypeLinux SystemTypeLinux
	// goverter:enum:map SystemTypeWeb SystemTypeWeb
	ToEntSystemType(model.SystemType) device.SystemType

	// goverter:enum:unknown @ignore
	// goverter:enum:map UserStatusUnspecified @ignore
	// goverter:enum:map UserStatusActive StatusActive
	// goverter:enum:map UserStatusBlocked StatusBlocked
	ToEntPorterInstanceStatus(model.UserStatus) porterinstance.Status
	ToEntPorterInstanceStatusList([]model.UserStatus) []porterinstance.Status

	// goverter:enum:unknown @ignore
	// goverter:enum:map PorterContextStatusUnspecified @ignore
	// goverter:enum:map PorterContextStatusActive StatusActive
	// goverter:enum:map PorterContextStatusDisabled StatusDisabled
	ToEntPorterContextStatus(modelsupervisor.PorterContextStatus) portercontext.Status

	// goverter:ignore CreatedAt
	// goverter:ignore UpdatedAt
	ToEntAppInfo(modelgebura.AppInfo) ent.AppInfo
	// goverter:enum:unknown @ignore
	// goverter:enum:map AppTypeUnspecified @ignore
	// goverter:enum:map AppTypeGame TypeGame
	ToEntAppInfoType(modelgebura.AppType) appinfo.Type
	// goverter:enum:unknown @ignore
	// goverter:enum:map AppTypeUnspecified @ignore
	// goverter:enum:map AppTypeGame TypeGame
	ToEntAppType(modelgebura.AppType) app.Type

	// goverter:enum:unknown @ignore
	// goverter:enum:map FeedConfigStatusUnspecified @ignore
	// goverter:enum:map FeedConfigStatusActive StatusActive
	// goverter:enum:map FeedConfigStatusSuspend StatusSuspend
	ToEntFeedConfigStatus(modelyesod.FeedConfigStatus) feedconfig.Status
	ToEntFeedConfigStatusList([]modelyesod.FeedConfigStatus) []feedconfig.Status

	// goverter:enum:unknown @ignore
	// goverter:enum:map FeedConfigPullStatusUnspecified @ignore
	// goverter:enum:map FeedConfigPullStatusProcessing LatestPullStatusProcessing
	// goverter:enum:map FeedConfigPullStatusSuccess LatestPullStatusSuccess
	// goverter:enum:map FeedConfigPullStatusFailed LatestPullStatusFailed
	ToEntFeedConfigLatestPullStatus(modelyesod.FeedConfigPullStatus) feedconfig.LatestPullStatus

	// goverter:enum:unknown @ignore
	// goverter:enum:map NotifyFlowStatusUnspecified @ignore
	// goverter:enum:map NotifyFlowStatusActive StatusActive
	// goverter:enum:map NotifyFlowStatusSuspend StatusSuspend
	ToEntNotifySourceStatus(modelnetzach.NotifyFlowStatus) notifyflow.Status

	// goverter:enum:unknown @ignore
	// goverter:enum:map NotifyTargetStatusUnspecified @ignore
	// goverter:enum:map NotifyTargetStatusActive StatusActive
	// goverter:enum:map NotifyTargetStatusSuspend StatusSuspend
	ToEntNotifyTargetStatus(modelnetzach.NotifyTargetStatus) notifytarget.Status
	ToEntNotifyTargetStatusList([]modelnetzach.NotifyTargetStatus) []notifytarget.Status

	// goverter:enum:unknown @ignore
	// goverter:enum:map SystemNotificationTypeUnspecified @ignore
	// goverter:enum:map SystemNotificationTypeSystem TypeSystem
	// goverter:enum:map SystemNotificationTypeUser TypeUser
	ToEntSystemNotificationType(modelnetzach.SystemNotificationType) systemnotification.Type
	ToEntSystemNotificationTypeList([]modelnetzach.SystemNotificationType) []systemnotification.Type

	// goverter:enum:unknown @ignore
	// goverter:enum:map SystemNotificationLevelUnspecified @ignore
	// goverter:enum:map SystemNotificationLevelInfo LevelInfo
	// goverter:enum:map SystemNotificationLevelWarning LevelWarn
	// goverter:enum:map SystemNotificationLevelError LevelError
	// goverter:enum:map SystemNotificationLevelOngoing LevelOngoing
	ToEntSystemNotificationLevel(modelnetzach.SystemNotificationLevel) systemnotification.Level
	ToEntSystemNotificationLevelList([]modelnetzach.SystemNotificationLevel) []systemnotification.Level

	// goverter:enum:unknown @ignore
	// goverter:enum:map SystemNotificationStatusUnspecified @ignore
	// goverter:enum:map SystemNotificationStatusUnread StatusUnread
	// goverter:enum:map SystemNotificationStatusRead StatusRead
	// goverter:enum:map SystemNotificationStatusDismissed StatusDismissed
	ToEntSystemNotificationStatus(modelnetzach.SystemNotificationStatus) systemnotification.Status
	ToEntSystemNotificationStatusList([]modelnetzach.SystemNotificationStatus) []systemnotification.Status

	// goverter:enum:unknown @ignore
	// goverter:enum:map ImageStatusUnspecified @ignore
	// goverter:enum:map ImageStatusUploaded StatusUploaded
	// goverter:enum:map ImageStatusScanned StatusScanned
	ToEntImageStatus(modelchesed.ImageStatus) image.Status
}
