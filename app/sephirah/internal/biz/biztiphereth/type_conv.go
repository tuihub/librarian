package biztiphereth

import (
	"github.com/tuihub/librarian/internal/lib/libauth"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

func ToLibAuthUserType(u pb.UserType) libauth.UserType {
	switch u {
	case pb.UserType_USER_TYPE_ADMIN:
		return libauth.UserTypeAdmin
	case pb.UserType_USER_TYPE_NORMAL:
		return libauth.UserTypeNormal
	default:
		return libauth.UserTypeUnspecified
	}
}

func ToPBUserType(u libauth.UserType) pb.UserType {
	switch u {
	case libauth.UserTypeAdmin:
		return pb.UserType_USER_TYPE_ADMIN
	case libauth.UserTypeNormal:
		return pb.UserType_USER_TYPE_NORMAL
	default:
		return pb.UserType_USER_TYPE_UNSPECIFIED
	}
}

func ToLibAuthUserTypeList(tl []pb.UserType) []libauth.UserType {
	res := make([]libauth.UserType, len(tl))
	for i, t := range tl {
		res[i] = ToLibAuthUserType(t)
	}
	return res
}

func ToBizUserStatus(s pb.UserStatus) UserStatus {
	switch s {
	case pb.UserStatus_USER_STATUS_ACTIVE:
		return UserStatusActive
	case pb.UserStatus_USER_STATUS_BLOCKED:
		return UserStatusBlocked
	default:
		return UserStatusUnspecified
	}
}

func ToBizUserStatusList(sl []pb.UserStatus) []UserStatus {
	res := make([]UserStatus, len(sl))
	for i, s := range sl {
		res[i] = ToBizUserStatus(s)
	}
	return res
}

func ToPBUser(u User) pb.User {
	return pb.User{
		Id:       &librarian.InternalID{Id: u.InternalID},
		Username: u.PassWord,
		Type:     ToPBUserType(u.Type),
	}
}

func ToPBUserList(ul []*User) []*pb.User {
	res := make([]*pb.User, len(ul))
	for i, u := range ul {
		if u != nil {
			uu := ToPBUser(*u)
			res[i] = &uu
		}
	}
	return res
}

func ToBizAccountPlatform(p librarian.AccountPlatform) AccountPlatform {
	switch p {
	case librarian.AccountPlatform_ACCOUNT_PLATFORM_STEAM:
		return AccountPlatformSteam
	default:
		return AccountPlatformUnspecified
	}
}

func ToPBAccountPlatform(p AccountPlatform) librarian.AccountPlatform {
	switch p {
	case AccountPlatformSteam:
		return librarian.AccountPlatform_ACCOUNT_PLATFORM_STEAM
	default:
		return librarian.AccountPlatform_ACCOUNT_PLATFORM_UNSPECIFIED
	}
}
