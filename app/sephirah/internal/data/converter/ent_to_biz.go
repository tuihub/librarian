package converter

import (
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizyesod"
	"github.com/tuihub/librarian/app/sephirah/internal/ent"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/account"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/app"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/apppackage"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/feedconfig"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/user"
	"github.com/tuihub/librarian/internal/lib/libauth"
)

// goverter:converter
type toBizConverter interface {
	// goverter:matchIgnoreCase
	// goverter:map ID ID
	// goverter:map Type | ToLibAuthUserType
	// goverter:map Status | ToBizUserStatus
	// goverter:ignore PassWord
	ToBizUser(*ent.User) *biztiphereth.User
	ToBizUserList([]*ent.User) []*biztiphereth.User

	// goverter:matchIgnoreCase
	// goverter:map ID ID
	// goverter:map Platform | ToBizAccountPlatform
	ToBizAccount(*ent.Account) *biztiphereth.Account
	ToBizAccountList([]*ent.Account) []*biztiphereth.Account

	// goverter:matchIgnoreCase
	// goverter:map ID ID
	// goverter:map Type | ToBizAppType
	// goverter:map Source | ToBizAppSource
	// goverter:map . Details
	ToBizApp(*ent.App) *bizgebura.App

	// goverter:matchIgnoreCase
	// goverter:map ID ID
	// goverter:map Source | ToBizAppPackageSource
	// goverter:mapIdentity Binary
	ToBizAppPackage(*ent.AppPackage) *bizgebura.AppPackage
	// goverter:map ID ID
	// goverter:map BinaryName Name
	// goverter:map BinarySize Size
	// goverter:map BinaryPublicURL PublicURL
	ToBizAppPacakgeBinary(ent.AppPackage) bizgebura.AppPackageBinary
	ToBizAppPackageList([]*ent.AppPackage) []*bizgebura.AppPackage

	// goverter:matchIgnoreCase
	// goverter:map ID ID
	// goverter:map Source | ToBizFeedConfigSource
	// goverter:map Status | ToBizFeedConfigStatus
	ToBizFeedConfig(*ent.FeedConfig) *bizyesod.FeedConfig
	ToBizFeedConfigList([]*ent.FeedConfig) []*bizyesod.FeedConfig
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

func ToBizUserStatus(s user.Status) biztiphereth.UserStatus {
	switch s {
	case user.StatusActive:
		return biztiphereth.UserStatusActive
	case user.StatusBlocked:
		return biztiphereth.UserStatusBlocked
	default:
		return biztiphereth.UserStatusUnspecified
	}
}

func ToBizAppType(t app.Type) bizgebura.AppType {
	switch t {
	case app.TypeGame:
		return bizgebura.AppTypeGame
	default:
		return bizgebura.AppTypeUnspecified
	}
}

func ToBizAppSource(s app.Source) bizgebura.AppSource {
	switch s {
	case app.SourceInternal:
		return bizgebura.AppSourceInternal
	case app.SourceSteam:
		return bizgebura.AppSourceSteam
	default:
		return bizgebura.AppSourceUnspecified
	}
}

func ToBizAppPackageSource(a apppackage.Source) bizgebura.AppPackageSource {
	switch a {
	case apppackage.SourceManual:
		return bizgebura.AppPackageSourceManual
	case apppackage.SourceSentinel:
		return bizgebura.AppPackageSourceSentinel
	default:
		return bizgebura.AppPackageSourceUnspecified
	}
}

func ToBizFeedConfigSource(s feedconfig.Source) bizyesod.FeedConfigSource {
	switch s {
	case feedconfig.SourceCommon:
		return bizyesod.FeedConfigSourceCommon
	default:
		return bizyesod.FeedConfigSourceUnspecified
	}
}

func ToBizFeedConfigStatus(s feedconfig.Status) bizyesod.FeedConfigStatus {
	switch s {
	case feedconfig.StatusActive:
		return bizyesod.FeedConfigStatusActive
	case feedconfig.StatusSuspend:
		return bizyesod.FeedConfigStatusSuspend
	default:
		return bizyesod.FeedConfigStatusUnspecified
	}
}

func ToBizAccountPlatform(p account.Platform) biztiphereth.AccountPlatform {
	switch p {
	case account.PlatformSteam:
		return biztiphereth.AccountPlatformSteam
	default:
		return biztiphereth.AccountPlatformUnspecified
	}
}
