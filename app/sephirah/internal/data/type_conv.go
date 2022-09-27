package data

import (
	"github.com/tuihub/librarian/app/sephirah/internal/biz"
	"github.com/tuihub/librarian/app/sephirah/internal/ent"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/app"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/user"
	"github.com/tuihub/librarian/internal/lib/libauth"
)

func toEntUserType(t libauth.UserType) user.Type {
	switch t {
	case libauth.UserTypeAdmin:
		return user.TypeAdmin
	default:
		return ""
	}
}

func toLibAuthUserType(t user.Type) libauth.UserType {
	switch t {
	case user.TypeAdmin:
		return libauth.UserTypeAdmin
	default:
		return libauth.UserTypeUnspecified
	}
}

func toEntUserStatus(s biz.UserStatus) user.Status {
	switch s {
	case biz.UserStatusActive:
		return user.StatusActive
	case biz.UserStatusBlocked:
		return user.StatusBlocked
	default:
		return ""
	}
}

func toBizUser(u *ent.User) *biz.User {
	if u == nil {
		return nil
	}
	return &biz.User{
		InternalID: u.InternalID,
		UserName:   u.Username,
		UserType:   toLibAuthUserType(u.Type),
	}
}

func toEntAppType(t biz.AppType) app.Type {
	switch t {
	case biz.AppTypeGame:
		return app.TypeGame
	default:
		return app.TypeGeneral
	}
}

func toEntAppSource(s biz.AppSource) app.Source {
	switch s {
	case biz.AppSourceInternal:
		return app.SourceInternal
	case biz.AppSourceSteam:
		return app.SourceSteam
	default:
		return ""
	}
}
