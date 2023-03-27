package converter

import (
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// goverter:converter
// goverter:extend ToPBInternalID
// goverter:extend ToPBTime
// goverter:extend ToPBTimePtr
// goverter:extend ToPBDuration
type toPBConverter interface {
	// goverter:matchIgnoreCase
	// goverter:map Type | ToPBUserType
	// goverter:map Status | ToPBUserStatus
	// goverter:ignore Password
	ToPBUser(*modeltiphereth.User) *pb.User
	ToPBUserList([]*modeltiphereth.User) []*pb.User

	// goverter:matchIgnoreCase
	// goverter:map Platform | ToPBAccountPlatform
	ToPBAccount(*modeltiphereth.Account) *librarian.Account
	ToPBAccountList([]*modeltiphereth.Account) []*librarian.Account

	// goverter:matchIgnoreCase
	// goverter:map Source | ToPBAppSource
	// goverter:map Type | ToPBAppType
	ToPBApp(*modelgebura.App) *librarian.App
	ToPBAppList([]*modelgebura.App) []*librarian.App

	// goverter:matchIgnoreCase
	// goverter:map Source | ToPBAppPackageSource
	// goverter:ignore SourceBindApp
	ToPBAppPackage(*modelgebura.AppPackage) *librarian.AppPackage
	// goverter:matchIgnoreCase
	ToPBAppPackageBinary(*modelgebura.AppPackageBinary) *librarian.AppPackageBinary
	ToPBAppPackageList([]*modelgebura.AppPackage) []*librarian.AppPackage

	// goverter:matchIgnoreCase
	ToPBFeed(*modelfeed.Feed) *librarian.Feed
	// goverter:matchIgnoreCase
	ToPBFeedItem(*modelfeed.Item) *librarian.FeedItem
	ToPBFeedItemList([]*modelfeed.Item) []*librarian.FeedItem
	// goverter:matchIgnoreCase
	ToPBFeedImage(*modelfeed.Image) *librarian.FeedImage
	// goverter:matchIgnoreCase
	ToPBEnclosure(*modelfeed.Enclosure) *librarian.FeedEnclosure
	// goverter:matchIgnoreCase
	// goverter:map Status | ToPBFeedConfigStatus
	// goverter:map Source | ToPBFeedConfigSource
	ToPBFeedConfig(*modelyesod.FeedConfig) *pb.FeedConfig
	// goverter:matchIgnoreCase
	// goverter:map FeedConfig Config
	ToPBFeedWithConfig(*modelyesod.FeedWithConfig) *pb.ListFeedConfigsResponse_FeedWithConfig
	ToPBFeedWithConfigList([]*modelyesod.FeedWithConfig) []*pb.ListFeedConfigsResponse_FeedWithConfig
	// goverter:matchIgnoreCase
	ToPBItemIDWithFeedID(*modelyesod.FeedItemIDWithFeedID) *pb.FeedItemIDWithFeedID
	ToPBItemIDWithFeedIDList([]*modelyesod.FeedItemIDWithFeedID) []*pb.FeedItemIDWithFeedID

	ToPBTimeRange(*model.TimeRange) *librarian.TimeRange
	ToPBInternalIDList([]model.InternalID) []*librarian.InternalID
}

func ToPBInternalID(id model.InternalID) *librarian.InternalID {
	return &librarian.InternalID{Id: int64(id)}
}

func ToPBTime(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}

func ToPBTimePtr(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}

func ToPBDuration(d time.Duration) *durationpb.Duration {
	return durationpb.New(d)
}

func ToPBUserType(u libauth.UserType) pb.UserType {
	switch u {
	case libauth.UserTypeUnspecified:
		return pb.UserType_USER_TYPE_UNSPECIFIED
	case libauth.UserTypeAdmin:
		return pb.UserType_USER_TYPE_ADMIN
	case libauth.UserTypeNormal:
		return pb.UserType_USER_TYPE_NORMAL
	case libauth.UserTypeSentinel:
		return pb.UserType_USER_TYPE_SENTINEL
	default:
		return pb.UserType_USER_TYPE_UNSPECIFIED
	}
}

func ToPBUserStatus(s modeltiphereth.UserStatus) pb.UserStatus {
	switch s {
	case modeltiphereth.UserStatusUnspecified:
		return pb.UserStatus_USER_STATUS_UNSPECIFIED
	case modeltiphereth.UserStatusActive:
		return pb.UserStatus_USER_STATUS_ACTIVE
	case modeltiphereth.UserStatusBlocked:
		return pb.UserStatus_USER_STATUS_BLOCKED
	default:
		return pb.UserStatus_USER_STATUS_UNSPECIFIED
	}
}

func ToPBAppType(t modelgebura.AppType) librarian.AppType {
	switch t {
	case modelgebura.AppTypeUnspecified:
		return librarian.AppType_APP_TYPE_UNSPECIFIED
	case modelgebura.AppTypeGame:
		return librarian.AppType_APP_TYPE_GAME
	default:
		return librarian.AppType_APP_TYPE_UNSPECIFIED
	}
}

func ToPBAppSource(s modelgebura.AppSource) librarian.AppSource {
	switch s {
	case modelgebura.AppSourceUnspecified:
		return librarian.AppSource_APP_SOURCE_UNSPECIFIED
	case modelgebura.AppSourceInternal:
		return librarian.AppSource_APP_SOURCE_INTERNAL
	case modelgebura.AppSourceSteam:
		return librarian.AppSource_APP_SOURCE_STEAM
	default:
		return librarian.AppSource_APP_SOURCE_UNSPECIFIED
	}
}

func ToPBAppPackageSource(a modelgebura.AppPackageSource) librarian.AppPackageSource {
	switch a {
	case modelgebura.AppPackageSourceUnspecified:
		return librarian.AppPackageSource_APP_PACKAGE_SOURCE_UNSPECIFIED
	case modelgebura.AppPackageSourceManual:
		return librarian.AppPackageSource_APP_PACKAGE_SOURCE_MANUAL
	case modelgebura.AppPackageSourceSentinel:
		return librarian.AppPackageSource_APP_PACKAGE_SOURCE_SENTINEL
	default:
		return librarian.AppPackageSource_APP_PACKAGE_SOURCE_UNSPECIFIED
	}
}

func ToPBAccountPlatform(p modeltiphereth.AccountPlatform) librarian.AccountPlatform {
	switch p {
	case modeltiphereth.AccountPlatformUnspecified:
		return librarian.AccountPlatform_ACCOUNT_PLATFORM_UNSPECIFIED
	case modeltiphereth.AccountPlatformSteam:
		return librarian.AccountPlatform_ACCOUNT_PLATFORM_STEAM
	default:
		return librarian.AccountPlatform_ACCOUNT_PLATFORM_UNSPECIFIED
	}
}

func ToPBFeedConfigStatus(s modelyesod.FeedConfigStatus) pb.FeedConfigStatus {
	switch s {
	case modelyesod.FeedConfigStatusUnspecified:
		return pb.FeedConfigStatus_FEED_CONFIG_STATUS_UNSPECIFIED
	case modelyesod.FeedConfigStatusActive:
		return pb.FeedConfigStatus_FEED_CONFIG_STATUS_ACTIVE
	case modelyesod.FeedConfigStatusSuspend:
		return pb.FeedConfigStatus_FEED_CONFIG_STATUS_SUSPEND
	default:
		return pb.FeedConfigStatus_FEED_CONFIG_STATUS_UNSPECIFIED
	}
}

func ToPBFeedConfigSource(s modelyesod.FeedConfigSource) pb.FeedConfigSource {
	switch s {
	case modelyesod.FeedConfigSourceUnspecified:
		return pb.FeedConfigSource_FEED_CONFIG_SOURCE_UNSPECIFIED
	case modelyesod.FeedConfigSourceCommon:
		return pb.FeedConfigSource_FEED_CONFIG_SOURCE_COMMON
	default:
		return pb.FeedConfigSource_FEED_CONFIG_SOURCE_UNSPECIFIED
	}
}
