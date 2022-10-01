package service

import (
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/internal/lib/libauth"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
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

func toPBUser(u biztiphereth.User) pb.User {
	return pb.User{
		Id:       &librarian.InternalID{Id: u.InternalID},
		Username: u.PassWord,
		Type:     toPBUserType(u.Type),
	}
}

func toPBUserList(ul []*biztiphereth.User) []*pb.User {
	res := make([]*pb.User, len(ul))
	for i, u := range ul {
		if u != nil {
			uu := toPBUser(*u)
			res[i] = &uu
		}
	}
	return res
}

func toPBAppType(t bizgebura.AppType) librarian.AppType {
	switch t {
	case bizgebura.AppTypeGame:
		return librarian.AppType_APP_TYPE_GAME
	default:
		return librarian.AppType_APP_TYPE_UNSPECIFIED
	}
}

func toBizAppType(t librarian.AppType) bizgebura.AppType {
	switch t {
	case librarian.AppType_APP_TYPE_GAME:
		return bizgebura.AppTypeGame
	default:
		return bizgebura.AppTypeGeneral
	}
}

func toBizAppTypeList(tl []librarian.AppType) []bizgebura.AppType {
	res := make([]bizgebura.AppType, len(tl))
	for i, s := range tl {
		res[i] = toBizAppType(s)
	}
	return res
}

func toPBAppSource(s bizgebura.AppSource) librarian.AppSource {
	switch s {
	case bizgebura.AppSourceInternal:
		return librarian.AppSource_APP_SOURCE_INTERNAL
	case bizgebura.AppSourceSteam:
		return librarian.AppSource_APP_SOURCE_STEAM
	default:
		return librarian.AppSource_APP_SOURCE_UNSPECIFIED
	}
}

func toBizAppSource(s librarian.AppSource) bizgebura.AppSource {
	switch s {
	case librarian.AppSource_APP_SOURCE_INTERNAL:
		return bizgebura.AppSourceInternal
	case librarian.AppSource_APP_SOURCE_STEAM:
		return bizgebura.AppSourceSteam
	default:
		return bizgebura.AppSourceUnspecified
	}
}

func toBizAppSourceList(sl []librarian.AppSource) []bizgebura.AppSource {
	res := make([]bizgebura.AppSource, len(sl))
	for i, s := range sl {
		res[i] = toBizAppSource(s)
	}
	return res
}

func toBizAppDetail(d *librarian.AppDetails) *bizgebura.AppDetails {
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

func toPBAppDetails(d *bizgebura.AppDetails) *librarian.AppDetails {
	if d == nil {
		return nil
	}
	return &librarian.AppDetails{
		Description: d.Description,
		ReleaseDate: d.ReleaseDate,
		Developer:   d.Developer,
		Publisher:   d.Publisher,
	}
}

func toPBApp(a *bizgebura.App, containDetails bool) *librarian.App {
	if a == nil {
		return nil
	}
	app := &librarian.App{
		Id:               &librarian.InternalID{Id: a.InternalID},
		Source:           toPBAppSource(a.Source),
		Name:             a.Name,
		Type:             toPBAppType(a.Type),
		ShortDescription: a.ShorDescription,
		ImageUrl:         a.ImageURL,
	}
	if containDetails {
		app.Details = toPBAppDetails(a.Details)
	}
	return app
}

func toPBAppList(al []*bizgebura.App, containDetails bool) []*librarian.App {
	res := make([]*librarian.App, len(al))
	for i, a := range al {
		res[i] = toPBApp(a, containDetails)
	}
	return res
}

func toBizAccountPlatform(p librarian.AccountPlatform) biztiphereth.AccountPlatform {
	switch p {
	case librarian.AccountPlatform_ACCOUNT_PLATFORM_STEAM:
		return biztiphereth.AccountPlatformSteam
	default:
		return biztiphereth.AccountPlatformUnspecified
	}
}
