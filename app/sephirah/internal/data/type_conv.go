package data

import (
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/ent"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/account"
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
		Type:       toLibAuthUserType(u.Type),
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

func toBizAppType(t app.Type) bizgebura.AppType {
	switch t {
	case app.TypeGame:
		return bizgebura.AppTypeGame
	default:
		return bizgebura.AppTypeGeneral
	}
}

func toBizAppSource(s app.Source) bizgebura.AppSource {
	switch s {
	case app.SourceInternal:
		return bizgebura.AppSourceInternal
	case app.SourceSteam:
		return bizgebura.AppSourceSteam
	default:
		return bizgebura.AppSourceUnspecified
	}
}

func toBizAppDetails(a *ent.App) *bizgebura.AppDetails {
	return &bizgebura.AppDetails{
		Description: a.Description,
		ReleaseDate: a.ReleaseDate,
		Developer:   a.Developer,
		Publisher:   a.Publisher,
	}
}

func toBizApp(a *ent.App) *bizgebura.App {
	return &bizgebura.App{
		InternalID:      a.InternalID,
		Source:          toBizAppSource(a.Source),
		SourceAppID:     a.SourceAppID,
		SourceURL:       a.SourceURL,
		Name:            a.Name,
		Type:            toBizAppType(a.Type),
		ShorDescription: a.ShortDescription,
		ImageURL:        a.ImageURL,
	}
}

func toEntAccountPlatform(t biztiphereth.AccountPlatform) account.Platform {
	switch t {
	case biztiphereth.AccountPlatformSteam:
		return account.PlatformSteam
	default:
		return ""
	}
}
