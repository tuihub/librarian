package converter

import (
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/ent"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/account"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/app"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/apppackage"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/feedconfig"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/user"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model/modelfeed"
)

// goverter:converter
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
	// goverter:map Platform | ToBizAccountPlatform
	ToBizAccount(*ent.Account) *modeltiphereth.Account
	ToBizAccountList([]*ent.Account) []*modeltiphereth.Account

	// goverter:matchIgnoreCase
	// goverter:map Type | ToBizAppType
	// goverter:map Source | ToBizAppSource
	// goverter:map . Details
	// goverter:ignore BoundInternal
	ToBizApp(*ent.App) *modelgebura.App
	ToBizAppList([]*ent.App) []*modelgebura.App

	// goverter:matchIgnoreCase
	// goverter:map Source | ToBizAppPackageSource
	// goverter:mapIdentity Binary
	ToBizAppPackage(*ent.AppPackage) *modelgebura.AppPackage
	// goverter:map BinaryName Name
	// goverter:map BinarySizeByte SizeByte
	// goverter:map BinaryPublicURL PublicURL
	ToBizAppPackageBinary(ent.AppPackage) modelgebura.AppPackageBinary
	ToBizAppPackageList([]*ent.AppPackage) []*modelgebura.AppPackage

	// goverter:matchIgnoreCase
	// goverter:map Source | ToBizFeedConfigSource
	// goverter:map Status | ToBizFeedConfigStatus
	// goverter:map LatestPullAt LatestPullTime
	// goverter:ignore Tags
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

func ToBizAppType(t app.Type) modelgebura.AppType {
	switch t {
	case app.TypeGame:
		return modelgebura.AppTypeGame
	default:
		return modelgebura.AppTypeUnspecified
	}
}

func ToBizAppSource(s app.Source) modelgebura.AppSource {
	switch s {
	case app.SourceInternal:
		return modelgebura.AppSourceInternal
	case app.SourceSteam:
		return modelgebura.AppSourceSteam
	default:
		return modelgebura.AppSourceUnspecified
	}
}

func ToBizAppPackageSource(a apppackage.Source) modelgebura.AppPackageSource {
	switch a {
	case apppackage.SourceManual:
		return modelgebura.AppPackageSourceManual
	case apppackage.SourceSentinel:
		return modelgebura.AppPackageSourceSentinel
	default:
		return modelgebura.AppPackageSourceUnspecified
	}
}

func ToBizFeedConfigSource(s feedconfig.Source) modelyesod.FeedConfigSource {
	switch s {
	case feedconfig.SourceCommon:
		return modelyesod.FeedConfigSourceCommon
	default:
		return modelyesod.FeedConfigSourceUnspecified
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

func ToBizAccountPlatform(p account.Platform) modeltiphereth.AccountPlatform {
	switch p {
	case account.PlatformSteam:
		return modeltiphereth.AccountPlatformSteam
	default:
		return modeltiphereth.AccountPlatformUnspecified
	}
}
