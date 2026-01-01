package converter

import (
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelchesed"
	"github.com/tuihub/librarian/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/model/modelnetzach"
	"github.com/tuihub/librarian/internal/model/modelsupervisor"
	"github.com/tuihub/librarian/internal/model/modelyesod"
)

func ToORMUserType(t model.UserType) string {
	switch t {
	case model.UserTypeAdmin:
		return "admin"
	case model.UserTypeNormal:
		return "normal"
	default:
		return ""
	}
}

func ToORMUserStatus(s model.UserStatus) string {
	switch s {
	case model.UserStatusActive:
		return "active"
	case model.UserStatusBlocked:
		return "blocked"
	default:
		return ""
	}
}

func ToORMSystemType(t model.SystemType) string {
	switch t {
	case model.SystemTypeIOS:
		return "ios"
	case model.SystemTypeAndroid:
		return "android"
	case model.SystemTypeWeb:
		return "web"
	case model.SystemTypeWindows:
		return "windows"
	case model.SystemTypeMacOS:
		return "macos"
	case model.SystemTypeLinux:
		return "linux"
	default:
		return ""
	}
}

func ToORMPorterInstanceStatus(s model.UserStatus) string {
	return ToORMUserStatus(s)
}

func ToORMPorterConnectionStatus(s modelsupervisor.PorterConnectionStatus) string {
	switch s {
	case modelsupervisor.PorterConnectionStatusConnected:
		return "connected"
	case modelsupervisor.PorterConnectionStatusDisconnected:
		return "disconnected"
	case modelsupervisor.PorterConnectionStatusActive:
		return "active"
	case modelsupervisor.PorterConnectionStatusActivationFailed:
		return "activation_failed"
	case modelsupervisor.PorterConnectionStatusDowngraded:
		return "downgraded"
	case modelsupervisor.PorterConnectionStatusQueueing:
		return "queueing"
	default:
		return ""
	}
}

func ToORMPorterContextStatus(s modelsupervisor.PorterContextStatus) string {
	switch s {
	case modelsupervisor.PorterContextStatusActive:
		return "active"
	case modelsupervisor.PorterContextStatusDisabled:
		return "disabled"
	default:
		return ""
	}
}

func ToORMPorterContextHandleStatus(s modelsupervisor.PorterContextHandleStatus) string {
	switch s {
	case modelsupervisor.PorterContextHandleStatusActive:
		return "active"
	case modelsupervisor.PorterContextHandleStatusDowngraded:
		return "downgraded"
	case modelsupervisor.PorterContextHandleStatusQueueing:
		return "queueing"
	case modelsupervisor.PorterContextHandleStatusBlocked:
		return "blocked"
	default:
		return ""
	}
}

func ToORMSentinelSessionStatus(s modelgebura.SentinelSessionStatus) string {
	switch s {
	case modelgebura.SentinelSessionStatusActive:
		return "active"
	case modelgebura.SentinelSessionStatusSuspend:
		return "suspend"
	default:
		return ""
	}
}

func ToORMAppInfoTypeManual(t modelgebura.AppType) string {
	switch t {
	case modelgebura.AppTypeGame:
		return "game"
	default:
		return ""
	}
}

func ToORMFeedConfigStatus(s modelyesod.FeedConfigStatus) string {
	switch s {
	case modelyesod.FeedConfigStatusActive:
		return "active"
	case modelyesod.FeedConfigStatusSuspend:
		return "suspend"
	default:
		return ""
	}
}

func ToORMFeedConfigLatestPullStatus(s modelyesod.FeedConfigPullStatus) string {
	switch s {
	case modelyesod.FeedConfigPullStatusProcessing:
		return "processing"
	case modelyesod.FeedConfigPullStatusSuccess:
		return "success"
	case modelyesod.FeedConfigPullStatusFailed:
		return "failed"
	default:
		return ""
	}
}

func ToORMNotifySourceStatus(s modelnetzach.NotifyFlowStatus) string {
	switch s {
	case modelnetzach.NotifyFlowStatusActive:
		return "active"
	case modelnetzach.NotifyFlowStatusSuspend:
		return "suspend"
	default:
		return ""
	}
}

func ToORMNotifyTargetStatus(s modelnetzach.NotifyTargetStatus) string {
	switch s {
	case modelnetzach.NotifyTargetStatusActive:
		return "active"
	case modelnetzach.NotifyTargetStatusSuspend:
		return "suspend"
	default:
		return ""
	}
}

func ToORMSystemNotificationType(t modelnetzach.SystemNotificationType) string {
	switch t {
	case modelnetzach.SystemNotificationTypeSystem:
		return "system"
	case modelnetzach.SystemNotificationTypeUser:
		return "user"
	default:
		return ""
	}
}

func ToORMSystemNotificationLevel(l modelnetzach.SystemNotificationLevel) string {
	switch l {
	case modelnetzach.SystemNotificationLevelInfo:
		return "info"
	case modelnetzach.SystemNotificationLevelWarning:
		return "warn"
	case modelnetzach.SystemNotificationLevelError:
		return "error"
	case modelnetzach.SystemNotificationLevelOngoing:
		return "ongoing"
	default:
		return ""
	}
}

func ToORMSystemNotificationStatus(s modelnetzach.SystemNotificationStatus) string {
	switch s {
	case modelnetzach.SystemNotificationStatusUnread:
		return "unread"
	case modelnetzach.SystemNotificationStatusRead:
		return "read"
	case modelnetzach.SystemNotificationStatusDismissed:
		return "dismissed"
	default:
		return ""
	}
}

func ToORMImageStatus(s modelchesed.ImageStatus) string {
	switch s {
	case modelchesed.ImageStatusUploaded:
		return "uploaded"
	case modelchesed.ImageStatusScanned:
		return "scanned"
	default:
		return ""
	}
}

// Reverse converters

func ToLibAuthUserType(s string) model.UserType {
	switch s {
	case "admin":
		return model.UserTypeAdmin
	case "normal":
		return model.UserTypeNormal
	default:
		return model.UserTypeUnspecified
	}
}

func ToBizUserStatus(s string) model.UserStatus {
	switch s {
	case "active":
		return model.UserStatusActive
	case "blocked":
		return model.UserStatusBlocked
	default:
		return model.UserStatusUnspecified
	}
}

func ToBizSystemType(s string) model.SystemType {
	switch s {
	case "ios":
		return model.SystemTypeIOS
	case "android":
		return model.SystemTypeAndroid
	case "web":
		return model.SystemTypeWeb
	case "windows":
		return model.SystemTypeWindows
	case "macos":
		return model.SystemTypeMacOS
	case "linux":
		return model.SystemTypeLinux
	default:
		return model.SystemTypeUnspecified
	}
}

func ToBizPorterStatus(s string) model.UserStatus {
	return ToBizUserStatus(s)
}

func ToBizPorterConnectionStatus(s string) modelsupervisor.PorterConnectionStatus {
	switch s {
	case "connected":
		return modelsupervisor.PorterConnectionStatusConnected
	case "disconnected":
		return modelsupervisor.PorterConnectionStatusDisconnected
	case "active":
		return modelsupervisor.PorterConnectionStatusActive
	case "activation_failed":
		return modelsupervisor.PorterConnectionStatusActivationFailed
	case "downgraded":
		return modelsupervisor.PorterConnectionStatusDowngraded
	case "queueing":
		return modelsupervisor.PorterConnectionStatusQueueing
	default:
		return modelsupervisor.PorterConnectionStatusUnspecified
	}
}

func ToBizPorterContextStatus(s string) modelsupervisor.PorterContextStatus {
	switch s {
	case "active":
		return modelsupervisor.PorterContextStatusActive
	case "disabled":
		return modelsupervisor.PorterContextStatusDisabled
	default:
		return modelsupervisor.PorterContextStatusUnspecified
	}
}

func ToBizPorterContextHandleStatus(s string) modelsupervisor.PorterContextHandleStatus {
	switch s {
	case "active":
		return modelsupervisor.PorterContextHandleStatusActive
	case "downgraded":
		return modelsupervisor.PorterContextHandleStatusDowngraded
	case "queueing":
		return modelsupervisor.PorterContextHandleStatusQueueing
	case "blocked":
		return modelsupervisor.PorterContextHandleStatusBlocked
	default:
		return modelsupervisor.PorterContextHandleStatusUnspecified
	}
}

func ToBizSentinelSessionStatus(s string) modelgebura.SentinelSessionStatus {
	switch s {
	case "active":
		return modelgebura.SentinelSessionStatusActive
	case "suspend":
		return modelgebura.SentinelSessionStatusSuspend
	default:
		return modelgebura.SentinelSessionStatusUnspecified
	}
}

func ToBizAppInfoType(s string) modelgebura.AppType {
	switch s {
	case "game":
		return modelgebura.AppTypeGame
	default:
		return modelgebura.AppTypeUnspecified
	}
}

func ToBizFeedConfigStatus(s string) modelyesod.FeedConfigStatus {
	switch s {
	case "active":
		return modelyesod.FeedConfigStatusActive
	case "suspend":
		return modelyesod.FeedConfigStatusSuspend
	default:
		return modelyesod.FeedConfigStatusUnspecified
	}
}

func ToBizFeedConfigPullStatus(s string) modelyesod.FeedConfigPullStatus {
	switch s {
	case "processing":
		return modelyesod.FeedConfigPullStatusProcessing
	case "success":
		return modelyesod.FeedConfigPullStatusSuccess
	case "failed":
		return modelyesod.FeedConfigPullStatusFailed
	default:
		return modelyesod.FeedConfigPullStatusUnspecified
	}
}

func ToBizNotifyFlowStatus(s string) modelnetzach.NotifyFlowStatus {
	switch s {
	case "active":
		return modelnetzach.NotifyFlowStatusActive
	case "suspend":
		return modelnetzach.NotifyFlowStatusSuspend
	default:
		return modelnetzach.NotifyFlowStatusUnspecified
	}
}

func ToBizNotifyTargetStatus(s string) modelnetzach.NotifyTargetStatus {
	switch s {
	case "active":
		return modelnetzach.NotifyTargetStatusActive
	case "suspend":
		return modelnetzach.NotifyTargetStatusSuspend
	default:
		return modelnetzach.NotifyTargetStatusUnspecified
	}
}

func ToBizSystemNotificationType(s string) modelnetzach.SystemNotificationType {
	switch s {
	case "system":
		return modelnetzach.SystemNotificationTypeSystem
	case "user":
		return modelnetzach.SystemNotificationTypeUser
	default:
		return modelnetzach.SystemNotificationTypeUnspecified
	}
}

func ToBizSystemNotificationLevel(s string) modelnetzach.SystemNotificationLevel {
	switch s {
	case "info":
		return modelnetzach.SystemNotificationLevelInfo
	case "warn":
		return modelnetzach.SystemNotificationLevelWarning
	case "error":
		return modelnetzach.SystemNotificationLevelError
	case "ongoing":
		return modelnetzach.SystemNotificationLevelOngoing
	default:
		return modelnetzach.SystemNotificationLevelUnspecified
	}
}

func ToBizSystemNotificationStatus(s string) modelnetzach.SystemNotificationStatus {
	switch s {
	case "unread":
		return modelnetzach.SystemNotificationStatusUnread
	case "read":
		return modelnetzach.SystemNotificationStatusRead
	case "dismissed":
		return modelnetzach.SystemNotificationStatusDismissed
	default:
		return modelnetzach.SystemNotificationStatusUnspecified
	}
}

func ToBizImageStatus(s string) modelchesed.ImageStatus {
	switch s {
	case "uploaded":
		return modelchesed.ImageStatusUploaded
	case "scanned":
		return modelchesed.ImageStatusScanned
	default:
		return modelchesed.ImageStatusUnspecified
	}
}
