package data

import (
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/ent"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/account"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/app"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/apppackage"
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

func toBizUserStatus(s user.Status) biztiphereth.UserStatus {
	switch s {
	case user.StatusActive:
		return biztiphereth.UserStatusActive
	case user.StatusBlocked:
		return biztiphereth.UserStatusBlocked
	default:
		return biztiphereth.UserStatusUnspecified
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
		Status:     toBizUserStatus(u.Status),
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
		return bizgebura.AppTypeUnspecified
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

func toEntAppPackageSource(a bizgebura.AppPackageSource) apppackage.Source {
	switch a {
	case bizgebura.AppPackageSourceManual:
		return apppackage.SourceManual
	case bizgebura.AppPackageSourceSentinel:
		return apppackage.SourceSentinel
	default:
		return ""
	}
}

func toBizAppPackageSource(a apppackage.Source) bizgebura.AppPackageSource {
	switch a {
	case apppackage.SourceManual:
		return bizgebura.AppPackageSourceManual
	case apppackage.SourceSentinel:
		return bizgebura.AppPackageSourceSentinel
	default:
		return bizgebura.AppPackageSourceUnspecified
	}
}

func toBizAppPackage(a *ent.AppPackage) *bizgebura.AppPackage {
	return &bizgebura.AppPackage{
		InternalID:      a.InternalID,
		Source:          toBizAppPackageSource(a.Source),
		SourceID:        a.SourceID,
		SourcePackageID: a.SourcePackageID,
		Name:            a.Name,
		Description:     a.Description,
		Binary: bizgebura.AppPackageBinary{
			Name: a.BinaryName,
			Size: a.BinarySize,
		},
	}
}

func toBizAppPackages(al []*ent.AppPackage) []*bizgebura.AppPackage {
	res := make([]*bizgebura.AppPackage, len(al))
	for i, a := range al {
		res[i] = toBizAppPackage(a)
	}
	return res
}
