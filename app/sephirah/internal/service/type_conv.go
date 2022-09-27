package service

import (
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
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

func toBizUserStatus(s pb.UserStatus) biztiphereth.UserStatus {
	switch s {
	case pb.UserStatus_USER_STATUS_ACTIVE:
		return biztiphereth.UserStatusActive
	case pb.UserStatus_USER_STATUS_BLOCKED:
		return biztiphereth.UserStatusBlocked
	default:
		return biztiphereth.UserStatusUnspecified
	}
}

func toBizUserStatusList(sl []pb.UserStatus) []biztiphereth.UserStatus {
	res := make([]biztiphereth.UserStatus, len(sl))
	for i, s := range sl {
		res[i] = toBizUserStatus(s)
	}
	return res
}

func toPBUser(u biztiphereth.User) pb.ListUserResponse_User {
	return pb.ListUserResponse_User{
		Id:       &pb.InternalID{Id: u.InternalID},
		Username: u.PassWord,
		Type:     toPBUserType(u.UserType),
	}
}

func toPBUserList(ul []*biztiphereth.User) []*pb.ListUserResponse_User {
	res := make([]*pb.ListUserResponse_User, len(ul))
	for i, u := range ul {
		if u != nil {
			uu := toPBUser(*u)
			res[i] = &uu
		}
	}
	return res
}

func toBizAppType(t pb.AppType) bizgebura.AppType {
	switch t {
	case pb.AppType_APP_TYPE_GAME:
		return bizgebura.AppTypeGame
	default:
		return bizgebura.AppTypeGeneral
	}
}

func toBizAppDetail(d *pb.AppDetails) *bizgebura.AppDetails {
	if d == nil {
		return nil
	}
	return &bizgebura.AppDetails{
		Description: d.GetDescription(),
		ReleaseDate: d.GetReleaseDate(),
		Developer:   d.GetDeveloper(),
		Publisher:   d.GetPublisher(),
	}
}
