package converter

import (
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizyesod"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

// goverter:converter
// goverter:extend ToBizInternalID
// goverter:extend ToLibAuthUserType
// goverter:extend ToBizUserStatus
// goverter:extend ToBizAppType
// goverter:extend ToBizAppSource
// goverter:extend PtrToString
// goverter:extend ToBizAppPackageSource
// goverter:extend DurationPBToDuration
// goverter:extend ToBizFeedConfigSource
// goverter:extend ToBizFeedConfigStatus
type toBizConverter interface {
	ToBizInternalIDList(idl []*librarian.InternalID) []model.InternalID
	// goverter:matchIgnoreCase
	ToBizUser(*pb.User) *biztiphereth.User
	ToLibAuthUserTypeList([]pb.UserType) []libauth.UserType
	ToBizUserStatusList([]pb.UserStatus) []biztiphereth.UserStatus

	// goverter:matchIgnoreCase
	ToBizApp(*librarian.App) *bizgebura.App
	ToBizAppTypeList([]librarian.AppType) []bizgebura.AppType
	ToBizAppSourceList([]librarian.AppSource) []bizgebura.AppSource

	// goverter:matchIgnoreCase
	ToBizAppPackage(*librarian.AppPackage) *bizgebura.AppPackage
	// goverter:matchIgnoreCase
	ToBizAppPackageBinary(*librarian.AppPackageBinary) *bizgebura.AppPackageBinary
	ToBizAppPackageSourceList([]librarian.AppPackageSource) []bizgebura.AppPackageSource

	// goverter:matchIgnoreCase
	ToBizFeedConfig(*pb.FeedConfig) *bizyesod.FeedConfig
	ToBizFeedConfigSourceList([]pb.FeedConfigSource) []bizyesod.FeedConfigSource
	ToBizFeedConfigStatusList([]pb.FeedConfigStatus) []bizyesod.FeedConfigStatus
}

func ToBizInternalID(id *librarian.InternalID) model.InternalID {
	if id == nil {
		return 0
	}
	return model.InternalID(id.Id)
}

func ToLibAuthUserType(u pb.UserType) libauth.UserType {
	switch u {
	case pb.UserType_USER_TYPE_UNSPECIFIED:
		return libauth.UserTypeUnspecified
	case pb.UserType_USER_TYPE_ADMIN:
		return libauth.UserTypeAdmin
	case pb.UserType_USER_TYPE_NORMAL:
		return libauth.UserTypeNormal
	case pb.UserType_USER_TYPE_SENTINEL:
		return libauth.UserTypeSentinel
	default:

		return libauth.UserTypeUnspecified
	}
}

func ToBizUserStatus(s pb.UserStatus) biztiphereth.UserStatus {
	switch s {
	case pb.UserStatus_USER_STATUS_UNSPECIFIED:
		return biztiphereth.UserStatusUnspecified
	case pb.UserStatus_USER_STATUS_ACTIVE:
		return biztiphereth.UserStatusActive
	case pb.UserStatus_USER_STATUS_BLOCKED:
		return biztiphereth.UserStatusBlocked
	default:

		return biztiphereth.UserStatusUnspecified
	}
}

func ToBizAppType(t librarian.AppType) bizgebura.AppType {
	switch t {
	case librarian.AppType_APP_TYPE_UNSPECIFIED:
		return bizgebura.AppTypeUnspecified
	case librarian.AppType_APP_TYPE_GAME:
		return bizgebura.AppTypeGame
	default:

		return bizgebura.AppTypeUnspecified
	}
}

func ToBizAppSource(s librarian.AppSource) bizgebura.AppSource {
	switch s {
	case librarian.AppSource_APP_SOURCE_UNSPECIFIED:
		return bizgebura.AppSourceUnspecified
	case librarian.AppSource_APP_SOURCE_INTERNAL:
		return bizgebura.AppSourceInternal
	case librarian.AppSource_APP_SOURCE_STEAM:
		return bizgebura.AppSourceSteam
	default:

		return bizgebura.AppSourceUnspecified
	}
}

func ToBizAppPackageSource(a librarian.AppPackageSource) bizgebura.AppPackageSource {
	switch a {
	case librarian.AppPackageSource_APP_PACKAGE_SOURCE_UNSPECIFIED:
		return bizgebura.AppPackageSourceUnspecified
	case librarian.AppPackageSource_APP_PACKAGE_SOURCE_MANUAL:
		return bizgebura.AppPackageSourceManual
	case librarian.AppPackageSource_APP_PACKAGE_SOURCE_SENTINEL:
		return bizgebura.AppPackageSourceSentinel
	default:

		return bizgebura.AppPackageSourceUnspecified
	}
}

func ToBizAccountPlatform(p librarian.AccountPlatform) biztiphereth.AccountPlatform {
	switch p {
	case librarian.AccountPlatform_ACCOUNT_PLATFORM_UNSPECIFIED:
		return biztiphereth.AccountPlatformUnspecified
	case librarian.AccountPlatform_ACCOUNT_PLATFORM_STEAM:
		return biztiphereth.AccountPlatformSteam
	default:

		return biztiphereth.AccountPlatformUnspecified
	}
}

func ToBizFeedConfigSource(s pb.FeedConfigSource) bizyesod.FeedConfigSource {
	switch s {
	case pb.FeedConfigSource_FEED_CONFIG_SOURCE_UNSPECIFIED:
		return bizyesod.FeedConfigSourceUnspecified
	case pb.FeedConfigSource_FEED_CONFIG_SOURCE_COMMON:
		return bizyesod.FeedConfigSourceCommon
	default:

		return bizyesod.FeedConfigSourceUnspecified
	}
}

func ToBizFeedConfigStatus(s pb.FeedConfigStatus) bizyesod.FeedConfigStatus {
	switch s {
	case pb.FeedConfigStatus_FEED_CONFIG_STATUS_UNSPECIFIED:
		return bizyesod.FeedConfigStatusUnspecified
	case pb.FeedConfigStatus_FEED_CONFIG_STATUS_ACTIVE:
		return bizyesod.FeedConfigStatusActive
	case pb.FeedConfigStatus_FEED_CONFIG_STATUS_SUSPEND:
		return bizyesod.FeedConfigStatusSuspend
	default:

		return bizyesod.FeedConfigStatusUnspecified
	}
}
