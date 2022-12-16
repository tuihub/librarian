package bizgebura

import (
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

func ToPBAppType(t AppType) librarian.AppType {
	switch t {
	case AppTypeGame:
		return librarian.AppType_APP_TYPE_GAME
	default:
		return librarian.AppType_APP_TYPE_UNSPECIFIED
	}
}

func ToBizAppType(t librarian.AppType) AppType {
	switch t {
	case librarian.AppType_APP_TYPE_GAME:
		return AppTypeGame
	default:
		return AppTypeUnspecified
	}
}

func ToBizAppTypeList(tl []librarian.AppType) []AppType {
	res := make([]AppType, len(tl))
	for i, s := range tl {
		res[i] = ToBizAppType(s)
	}
	return res
}

func ToPBAppSource(s AppSource) librarian.AppSource {
	switch s {
	case AppSourceInternal:
		return librarian.AppSource_APP_SOURCE_INTERNAL
	case AppSourceSteam:
		return librarian.AppSource_APP_SOURCE_STEAM
	default:
		return librarian.AppSource_APP_SOURCE_UNSPECIFIED
	}
}

func ToBizAppSource(s librarian.AppSource) AppSource {
	switch s {
	case librarian.AppSource_APP_SOURCE_INTERNAL:
		return AppSourceInternal
	case librarian.AppSource_APP_SOURCE_STEAM:
		return AppSourceSteam
	default:
		return AppSourceUnspecified
	}
}

func ToBizAppSourceList(sl []librarian.AppSource) []AppSource {
	res := make([]AppSource, len(sl))
	for i, s := range sl {
		res[i] = ToBizAppSource(s)
	}
	return res
}

func ToBizAppDetail(d *librarian.AppDetails) *AppDetails {
	if d == nil {
		return nil
	}
	return &AppDetails{
		Description: d.GetDescription(),
		ReleaseDate: d.GetReleaseDate(),
		Developer:   d.GetDeveloper(),
		Publisher:   d.GetPublisher(),
	}
}

func ToPBAppDetails(d *AppDetails) *librarian.AppDetails {
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

func ToPBApp(a *App, containDetails bool) *librarian.App {
	if a == nil {
		return nil
	}
	app := &librarian.App{
		Id:               &librarian.InternalID{Id: a.InternalID},
		Source:           ToPBAppSource(a.Source),
		SourceAppId:      a.SourceAppID,
		SourceUrl:        &a.SourceURL,
		Name:             a.Name,
		Type:             ToPBAppType(a.Type),
		ShortDescription: a.ShorDescription,
		ImageUrl:         a.ImageURL,
	}
	if containDetails {
		app.Details = ToPBAppDetails(a.Details)
	}
	return app
}

func ToPBAppList(al []*App, containDetails bool) []*librarian.App {
	res := make([]*librarian.App, len(al))
	for i, a := range al {
		res[i] = ToPBApp(a, containDetails)
	}
	return res
}

func ToBizAppPackageSource(a librarian.AppPackageSource) AppPackageSource {
	switch a {
	case librarian.AppPackageSource_APP_PACKAGE_SOURCE_MANUAL:
		return AppPackageSourceManual
	case librarian.AppPackageSource_APP_PACKAGE_SOURCE_SENTINEL:
		return AppPackageSourceSentinel
	default:
		return AppPackageSourceUnspecified
	}
}

func ToBizAppPackageSourceList(al []librarian.AppPackageSource) []AppPackageSource {
	res := make([]AppPackageSource, len(al))
	for i, a := range al {
		res[i] = ToBizAppPackageSource(a)
	}
	return res
}

func ToPBAppPackageSource(a AppPackageSource) librarian.AppPackageSource {
	switch a {
	case AppPackageSourceManual:
		return librarian.AppPackageSource_APP_PACKAGE_SOURCE_MANUAL
	case AppPackageSourceSentinel:
		return librarian.AppPackageSource_APP_PACKAGE_SOURCE_SENTINEL
	default:
		return librarian.AppPackageSource_APP_PACKAGE_SOURCE_UNSPECIFIED
	}
}

func ToPBAppPackage(a *AppPackage) *librarian.AppPackage {
	return &librarian.AppPackage{
		Id:              &librarian.InternalID{Id: a.InternalID},
		Source:          ToPBAppPackageSource(a.Source),
		SourceId:        &librarian.InternalID{Id: a.SourceID},
		SourcePackageId: a.SourcePackageID,
		Name:            a.Name,
		Description:     a.Description,
		Binary: &librarian.AppPackageBinary{
			Name:      a.Binary.Name,
			Size:      a.Binary.Size,
			PublicUrl: "",
		},
		SourceBindApp: nil,
	}
}

func ToPBAppPackageList(al []*AppPackage) []*librarian.AppPackage {
	res := make([]*librarian.AppPackage, len(al))
	for i, a := range al {
		res[i] = ToPBAppPackage(a)
	}
	return res
}
