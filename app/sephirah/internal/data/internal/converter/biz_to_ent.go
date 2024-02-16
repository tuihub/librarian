package converter

import (
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
)

// goverter:converter
// goverter:output:file ./generated.go
// goverter:output:package github.com/tuihub/librarian/app/sephirah/internal/data/internal/converter
// goverter:extend ToEntUserType
// goverter:extend ToEntUserStatus
// goverter:extend ToEntAppType
// goverter:extend ToEntFeedConfigStatus
// goverter:extend ToEntNotifyTargetStatus
// goverter:extend ToEntSystemType
type toEntConverter interface { //nolint:unused // used by generator
	ToEntUserTypeList([]libauth.UserType) []user.Type
	ToEntUserStatusList([]modeltiphereth.UserStatus) []user.Status

	// goverter:autoMap Details
	// goverter:useZeroValueOnPointerInconsistency
	// goverter:ignore Edges
	// goverter:ignore CreatedAt
	// goverter:ignore UpdatedAt
	ToEntAppInfo(modelgebura.AppInfo) ent.AppInfo

	ToEntFeedConfigStatusList([]modelyesod.FeedConfigStatus) []feedconfig.Status

	ToEntNotifyTargetStatusList([]modelnetzach.NotifyTargetStatus) []notifytarget.Status
}

func ToEntUserType(t libauth.UserType) user.Type {
	switch t {
	case libauth.UserTypeUnspecified:
		return ""
	case libauth.UserTypeAdmin:
		return user.TypeAdmin
	case libauth.UserTypeNormal:
		return user.TypeNormal
	case libauth.UserTypeSentinel:
		return user.TypeSentinel
	default:
		return ""
	}
}

func ToEntUserStatus(s modeltiphereth.UserStatus) user.Status {
	switch s {
	case modeltiphereth.UserStatusUnspecified:
		return ""
	case modeltiphereth.UserStatusActive:
		return user.StatusActive
	case modeltiphereth.UserStatusBlocked:
		return user.StatusBlocked
	default:
		return ""
	}
}

func ToEntPorterInstanceStatus(s modeltiphereth.PorterInstanceStatus) porterinstance.Status {
	switch s {
	case modeltiphereth.PorterInstanceStatusUnspecified:
		return ""
	case modeltiphereth.PorterInstanceStatusActive:
		return porterinstance.StatusActive
	case modeltiphereth.PorterInstanceStatusBlocked:
		return porterinstance.StatusBlocked
	default:
		return ""
	}
}

func ToEntAppType(t modelgebura.AppType) appinfo.Type {
	switch t {
	case modelgebura.AppTypeUnspecified:
		return appinfo.TypeUnknown
	case modelgebura.AppTypeGame:
		return appinfo.TypeGame
	default:
		return appinfo.TypeUnknown
	}
}

func ToEntFeedConfigStatus(s modelyesod.FeedConfigStatus) feedconfig.Status {
	switch s {
	case modelyesod.FeedConfigStatusUnspecified:
		return ""
	case modelyesod.FeedConfigStatusActive:
		return feedconfig.StatusActive
	case modelyesod.FeedConfigStatusSuspend:
		return feedconfig.StatusSuspend
	default:
		return ""
	}
}

func ToEntNotifyTargetStatus(s modelnetzach.NotifyTargetStatus) notifytarget.Status {
	switch s {
	case modelnetzach.NotifyTargetStatusUnspecified:
		return ""
	case modelnetzach.NotifyTargetStatusActive:
		return notifytarget.StatusActive
	case modelnetzach.NotifyTargetStatusSuspend:
		return notifytarget.StatusSuspend
	default:
		return ""
	}
}

func ToEntNotifySourceSource(s modelnetzach.NotifyFlowStatus) notifyflow.Status {
	switch s {
	case modelnetzach.NotifyFlowStatusUnspecified:
		return ""
	case modelnetzach.NotifyFlowStatusActive:
		return notifyflow.StatusActive
	case modelnetzach.NotifyFlowStatusSuspend:
		return notifyflow.StatusSuspend
	default:
		return ""
	}
}

func ToEntImageStatus(s modelchesed.ImageStatus) image.Status {
	switch s {
	case modelchesed.ImageStatusUnspecified:
		return ""
	case modelchesed.ImageStatusUploaded:
		return image.StatusUploaded
	case modelchesed.ImageStatusScanned:
		return image.StatusScanned
	default:
		return ""
	}
}

func ToEntSystemType(s modeltiphereth.SystemType) deviceinfo.SystemType {
	switch s {
	case modeltiphereth.SystemTypeUnspecified:
		return deviceinfo.SystemTypeUnknown
	case modeltiphereth.SystemTypeIOS:
		return deviceinfo.SystemTypeIos
	case modeltiphereth.SystemTypeAndroid:
		return deviceinfo.SystemTypeAndroid
	case modeltiphereth.SystemTypeWindows:
		return deviceinfo.SystemTypeWindows
	case modeltiphereth.SystemTypeMacOS:
		return deviceinfo.SystemTypeMacos
	case modeltiphereth.SystemTypeLinux:
		return deviceinfo.SystemTypeLinux
	case modeltiphereth.SystemTypeWeb:
		return deviceinfo.SystemTypeWeb
	default:
		return deviceinfo.SystemTypeUnknown
	}
}
