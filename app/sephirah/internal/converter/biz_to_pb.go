package converter

import (
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/internal/lib/libauth"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

// goverter:converter
type toPBConverter interface {
	// goverter:matchIgnoreCase
	// goverter:mapIdentity Id
	// goverter:map Type | ToPBUserType
	// goverter:map Status | ToPBUserStatus
	// goverter:ignore Password
	ToPBUser(biztiphereth.User) pb.User
	// goverter:map InternalID Id
	ToPBUserInternalID(biztiphereth.User) librarian.InternalID
	ToPBUserList([]*biztiphereth.User) []*pb.User

	// goverter:matchIgnoreCase
	// goverter:mapIdentity Id
	// goverter:map Source | ToPBAppSource
	// goverter:map Type | ToPBAppType
	ToPBApp(bizgebura.App) librarian.App
	// goverter:map InternalID Id
	ToPBAppInternalID(bizgebura.App) librarian.InternalID
	ToPBAppList([]*bizgebura.App) []*librarian.App

	// goverter:matchIgnoreCase
	// goverter:mapIdentity Id
	// goverter:mapIdentity SourceId
	// goverter:map Source | ToPBAppPackageSource
	// goverter:ignore SourceBindApp
	ToPBAppPackage(*bizgebura.AppPackage) *librarian.AppPackage
	// goverter:matchIgnoreCase
	ToPBAppPackageBinary(*bizgebura.AppPackageBinary) *librarian.AppPackageBinary
	// goverter:map InternalID Id
	ToPBAppPackageInternalID(bizgebura.AppPackage) librarian.InternalID
	ToPBAppPackageList([]*bizgebura.AppPackage) []*librarian.AppPackage
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

func ToPBUserStatus(s biztiphereth.UserStatus) pb.UserStatus {
	switch s {
	case biztiphereth.UserStatusUnspecified:
		return pb.UserStatus_USER_STATUS_UNSPECIFIED
	case biztiphereth.UserStatusActive:
		return pb.UserStatus_USER_STATUS_ACTIVE
	case biztiphereth.UserStatusBlocked:
		return pb.UserStatus_USER_STATUS_BLOCKED
	default:
		return pb.UserStatus_USER_STATUS_UNSPECIFIED
	}
}

func ToPBAppType(t bizgebura.AppType) librarian.AppType {
	switch t {
	case bizgebura.AppTypeUnspecified:
		return librarian.AppType_APP_TYPE_UNSPECIFIED
	case bizgebura.AppTypeGame:
		return librarian.AppType_APP_TYPE_GAME
	default:
		return librarian.AppType_APP_TYPE_UNSPECIFIED
	}
}

func ToPBAppSource(s bizgebura.AppSource) librarian.AppSource {
	switch s {
	case bizgebura.AppSourceUnspecified:
		return librarian.AppSource_APP_SOURCE_UNSPECIFIED
	case bizgebura.AppSourceInternal:
		return librarian.AppSource_APP_SOURCE_INTERNAL
	case bizgebura.AppSourceSteam:
		return librarian.AppSource_APP_SOURCE_STEAM
	default:
		return librarian.AppSource_APP_SOURCE_UNSPECIFIED
	}
}

func ToPBAppPackageSource(a bizgebura.AppPackageSource) librarian.AppPackageSource {
	switch a {
	case bizgebura.AppPackageSourceUnspecified:
		return librarian.AppPackageSource_APP_PACKAGE_SOURCE_UNSPECIFIED
	case bizgebura.AppPackageSourceManual:
		return librarian.AppPackageSource_APP_PACKAGE_SOURCE_MANUAL
	case bizgebura.AppPackageSourceSentinel:
		return librarian.AppPackageSource_APP_PACKAGE_SOURCE_SENTINEL
	default:
		return librarian.AppPackageSource_APP_PACKAGE_SOURCE_UNSPECIFIED
	}
}

func ToPBAccountPlatform(p biztiphereth.AccountPlatform) librarian.AccountPlatform {
	switch p {
	case biztiphereth.AccountPlatformUnspecified:
		return librarian.AccountPlatform_ACCOUNT_PLATFORM_UNSPECIFIED
	case biztiphereth.AccountPlatformSteam:
		return librarian.AccountPlatform_ACCOUNT_PLATFORM_STEAM
	default:
		return librarian.AccountPlatform_ACCOUNT_PLATFORM_UNSPECIFIED
	}
}
