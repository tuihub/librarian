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
		Type:     toPBUserType(u.Type),
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

func toPBAppType(t bizgebura.AppType) pb.AppType {
	switch t {
	case bizgebura.AppTypeGame:
		return pb.AppType_APP_TYPE_GAME
	default:
		return pb.AppType_APP_TYPE_UNSPECIFIED
	}
}

func toBizAppType(t pb.AppType) bizgebura.AppType {
	switch t {
	case pb.AppType_APP_TYPE_GAME:
		return bizgebura.AppTypeGame
	default:
		return bizgebura.AppTypeGeneral
	}
}

func toBizAppTypeList(tl []pb.AppType) []bizgebura.AppType {
	res := make([]bizgebura.AppType, len(tl))
	for i, s := range tl {
		res[i] = toBizAppType(s)
	}
	return res
}

func toPBAppSource(s bizgebura.AppSource) pb.AppSource {
	switch s {
	case bizgebura.AppSourceInternal:
		return pb.AppSource_APP_SOURCE_INTERNAL
	case bizgebura.AppSourceSteam:
		return pb.AppSource_APP_SOURCE_STEAM
	default:
		return pb.AppSource_APP_SOURCE_UNSPECIFIED
	}
}

func toBizAppSource(s pb.AppSource) bizgebura.AppSource {
	switch s {
	case pb.AppSource_APP_SOURCE_INTERNAL:
		return bizgebura.AppSourceInternal
	case pb.AppSource_APP_SOURCE_STEAM:
		return bizgebura.AppSourceSteam
	default:
		return bizgebura.AppSourceUnspecified
	}
}

func toBizAppSourceList(sl []pb.AppSource) []bizgebura.AppSource {
	res := make([]bizgebura.AppSource, len(sl))
	for i, s := range sl {
		res[i] = toBizAppSource(s)
	}
	return res
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

func toPBAppDetails(d *bizgebura.AppDetails) *pb.AppDetails {
	if d == nil {
		return nil
	}
	return &pb.AppDetails{
		Description: d.Description,
		ReleaseDate: d.ReleaseDate,
		Developer:   d.Developer,
		Publisher:   d.Publisher,
	}
}

func toPBApp(a *bizgebura.App, containDetails bool) *pb.App {
	if a == nil {
		return nil
	}
	app := &pb.App{
		Id:               &pb.InternalID{Id: a.InternalID},
		Source:           toPBAppSource(a.Source),
		Name:             a.Name,
		Type:             toPBAppType(a.Type),
		ShortDescription: a.ShorDescription,
		ImageUrl:         a.ImageURL,
	}
	if containDetails {
		appDetails := toPBAppDetails(a.Details)
		app.XDetails = &pb.App_Details{
			Details: appDetails,
		}
	}
	return app
}

func toPBAppList(al []*bizgebura.App, containDetails bool) []*pb.App {
	res := make([]*pb.App, len(al))
	for i, a := range al {
		res[i] = toPBApp(a, containDetails)
	}
	return res
}
