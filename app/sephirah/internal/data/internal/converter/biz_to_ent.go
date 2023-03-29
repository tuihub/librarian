package converter

import (
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/account"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/app"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/apppackage"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feedconfig"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	"github.com/tuihub/librarian/internal/lib/libauth"
)

// goverter:converter
// goverter:extend ToEntUserType
// goverter:extend ToEntUserStatus
// goverter:extend ToEntAppType
// goverter:extend ToEntAppSource
// goverter:extend ToEntAppPackageSource
// goverter:extend ToEntFeedConfigSource
// goverter:extend ToEntFeedConfigStatus
type toEntConverter interface { //nolint:unused // used by generator
	ToEntUserTypeList([]libauth.UserType) []user.Type
	ToEntUserStatusList([]modeltiphereth.UserStatus) []user.Status

	// goverter:autoMap Details
	// goverter:useZeroValueOnPointerInconsistency
	// goverter:ignore Edges
	// goverter:ignore CreatedAt
	// goverter:ignore UpdatedAt
	ToEntApp(modelgebura.App) ent.App
	ToEntAppPackageSourceList([]modelgebura.AppPackageSource) []apppackage.Source

	ToEntFeedConfigSourceList([]modelyesod.FeedConfigSource) []feedconfig.Source
	ToEntFeedConfigStatusList([]modelyesod.FeedConfigStatus) []feedconfig.Status
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

func ToEntAppType(t modelgebura.AppType) app.Type {
	switch t {
	case modelgebura.AppTypeUnspecified:
		return ""
	case modelgebura.AppTypeGame:
		return app.TypeGame
	default:
		return ""
	}
}

func ToEntAppSource(s modelgebura.AppSource) app.Source {
	switch s {
	case modelgebura.AppSourceUnspecified:
		return ""
	case modelgebura.AppSourceInternal:
		return app.SourceInternal
	case modelgebura.AppSourceSteam:
		return app.SourceSteam
	default:
		return ""
	}
}

func ToEntAccountPlatform(t modeltiphereth.AccountPlatform) account.Platform {
	switch t {
	case modeltiphereth.AccountPlatformUnspecified:
		return ""
	case modeltiphereth.AccountPlatformSteam:
		return account.PlatformSteam
	default:
		return ""
	}
}

func ToEntAppPackageSource(a modelgebura.AppPackageSource) apppackage.Source {
	switch a {
	case modelgebura.AppPackageSourceUnspecified:
		return ""
	case modelgebura.AppPackageSourceManual:
		return apppackage.SourceManual
	case modelgebura.AppPackageSourceSentinel:
		return apppackage.SourceSentinel
	default:
		return ""
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

func ToEntFeedConfigSource(s modelyesod.FeedConfigSource) feedconfig.Source {
	switch s {
	case modelyesod.FeedConfigSourceUnspecified:
		return ""
	case modelyesod.FeedConfigSourceCommon:
		return feedconfig.SourceCommon
	default:
		return ""
	}
}
