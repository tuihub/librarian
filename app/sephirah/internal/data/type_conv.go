package data

import (
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
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

func toEntUserStatus(s biztiphereth.UserStatus) user.Status {
	switch s {
	case biztiphereth.UserStatusActive:
		return user.StatusActive
	case biztiphereth.UserStatusBlocked:
		return user.StatusBlocked
	default:
		return ""
	}
}

func toBizUser(u *ent.User) *biztiphereth.User {
	if u == nil {
		return nil
	}
	return &biztiphereth.User{
		InternalID: u.InternalID,
		UserName:   u.Username,
		UserType:   toLibAuthUserType(u.Type),
	}
}

func toEntAppType(t bizgebura.AppType) app.Type {
	switch t {
	case bizgebura.AppTypeGame:
		return app.TypeGame
	default:
		return app.TypeGeneral
	}
}

func toEntAppSource(s bizgebura.AppSource) app.Source {
	switch s {
	case bizgebura.AppSourceInternal:
		return app.SourceInternal
	case bizgebura.AppSourceSteam:
		return app.SourceSteam
	default:
		return ""
	}
}
