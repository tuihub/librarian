package converter

import (
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizyesod"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/account"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/app"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/apppackage"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/feedconfig"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/user"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
)

// goverter:converter
// goverter:extend ToEntInternalID
// goverter:extend ToEntUserType
// goverter:extend ToEntUserStatus
// goverter:extend ToEntFeedConfigSource
// goverter:extend ToEntFeedConfigStatus
type toEntConverter interface {
	ToEntInternalIDList([]model.InternalID) []int64
	ToEntUserTypeList([]libauth.UserType) []user.Type
	ToEntUserStatusList([]biztiphereth.UserStatus) []user.Status

	ToEntFeedConfigSourceList([]bizyesod.FeedConfigSource) []feedconfig.Source
	ToEntFeedConfigStatusList([]bizyesod.FeedConfigStatus) []feedconfig.Status
}

func ToEntInternalID(id model.InternalID) int64 {
	return int64(id)
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

func ToEntUserStatus(s biztiphereth.UserStatus) user.Status {
	switch s {
	case biztiphereth.UserStatusUnspecified:
		return ""
	case biztiphereth.UserStatusActive:
		return user.StatusActive
	case biztiphereth.UserStatusBlocked:
		return user.StatusBlocked
	default:
		return ""
	}
}

func ToEntAppType(t bizgebura.AppType) app.Type {
	switch t {
	case bizgebura.AppTypeUnspecified:
		return ""
	case bizgebura.AppTypeGame:
		return app.TypeGame
	default:
		return ""
	}
}

func ToEntAppSource(s bizgebura.AppSource) app.Source {
	switch s {
	case bizgebura.AppSourceUnspecified:
		return ""
	case bizgebura.AppSourceInternal:
		return app.SourceInternal
	case bizgebura.AppSourceSteam:
		return app.SourceSteam
	default:
		return ""
	}
}

func ToEntAccountPlatform(t biztiphereth.AccountPlatform) account.Platform {
	switch t {
	case biztiphereth.AccountPlatformUnspecified:
		return ""
	case biztiphereth.AccountPlatformSteam:
		return account.PlatformSteam
	default:
		return ""
	}
}

func ToEntAppPackageSource(a bizgebura.AppPackageSource) apppackage.Source {
	switch a {
	case bizgebura.AppPackageSourceUnspecified:
		return ""
	case bizgebura.AppPackageSourceManual:
		return apppackage.SourceManual
	case bizgebura.AppPackageSourceSentinel:
		return apppackage.SourceSentinel
	default:
		return ""
	}
}

func ToEntFeedConfigStatus(s bizyesod.FeedConfigStatus) feedconfig.Status {
	switch s {
	case bizyesod.FeedConfigStatusUnspecified:
		return ""
	case bizyesod.FeedConfigStatusActive:
		return feedconfig.StatusActive
	case bizyesod.FeedConfigStatusSuspend:
		return feedconfig.StatusSuspend
	default:
		return ""
	}
}

func ToEntFeedConfigSource(s bizyesod.FeedConfigSource) feedconfig.Source {
	switch s {
	case bizyesod.FeedConfigSourceUnspecified:
		return ""
	case bizyesod.FeedConfigSourceCommon:
		return feedconfig.SourceCommon
	default:
		return ""
	}
}
