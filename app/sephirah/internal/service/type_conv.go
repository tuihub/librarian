package service

import (
	"github.com/tuihub/librarian/app/sephirah/internal/biz"
	"github.com/tuihub/librarian/internal/lib/libauth"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
)

func toLibAuthUserType(u pb.UserType) libauth.UserType {
	switch u {
	case pb.UserType_USER_TYPE_ADMIN:
		return libauth.UserTypeAdmin
	case pb.UserType_USER_TYPE_NORMAL:
		return libauth.UserTypeNormal
	default:
		return libauth.UserTypeUnspecified
	}
}

func toPBUserType(u libauth.UserType) pb.UserType {
	switch u {
	case libauth.UserTypeAdmin:
		return pb.UserType_USER_TYPE_ADMIN
	case libauth.UserTypeNormal:
		return pb.UserType_USER_TYPE_NORMAL
	default:
		return pb.UserType_USER_TYPE_UNSPECIFIED
	}
}

func toLibAuthUserTypeList(tl []pb.UserType) []libauth.UserType {
	res := make([]libauth.UserType, len(tl))
	for i, t := range tl {
		res[i] = toLibAuthUserType(t)
	}
	return res
}

func toBizUserStatus(s pb.UserStatus) biz.UserStatus {
	switch s {
	case pb.UserStatus_USER_STATUS_ACTIVE:
		return biz.UserStatusActive
	case pb.UserStatus_USER_STATUS_BLOCKED:
		return biz.UserStatusBlocked
	default:
		return biz.UserStatusUnspecified
	}
}

func toBizUserStatusList(sl []pb.UserStatus) []biz.UserStatus {
	res := make([]biz.UserStatus, len(sl))
	for i, s := range sl {
		res[i] = toBizUserStatus(s)
	}
	return res
}

func toPBUser(u biz.User) pb.ListUserResponse_User {
	return pb.ListUserResponse_User{
		Id:       &pb.InternalID{Id: u.InternalID},
		Username: u.PassWord,
		Type:     toPBUserType(u.UserType),
	}
}

func toPBUserList(ul []*biz.User) []*pb.ListUserResponse_User {
	res := make([]*pb.ListUserResponse_User, len(ul))
	for i, u := range ul {
		if u != nil {
			uu := toPBUser(*u)
			res[i] = &uu
		}
	}
	return res
}

func toBizAppType(t pb.AppType) biz.AppType {
	switch t {
	case pb.AppType_APP_TYPE_GAME:
		return biz.AppTypeGame
	default:
		return biz.AppTypeGeneral
	}
}

func toBizAppDetail(d *pb.AppDetails) *biz.AppDetails {
	if d == nil {
		return nil
	}
	return &biz.AppDetails{
		Description: d.GetDescription(),
		ReleaseDate: d.GetReleaseDate(),
		Developer:   d.GetDeveloper(),
		Publisher:   d.GetPublisher(),
	}
}
