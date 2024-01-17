package modelgebura

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type App struct {
	ID                 model.InternalID
	Internal           bool
	Source             string
	SourceAppID        string
	SourceURL          string
	Name               string
	Type               AppType
	ShortDescription   string
	IconImageURL       string
	BackgroundImageURL string
	CoverImageURL      string
	Tags               []string
	Details            *AppDetails
	// the bound Internal app id if self is external
	BoundInternal    model.InternalID
	LatestUpdateTime time.Time
}

type AppID struct {
	Internal    bool
	Source      string
	SourceAppID string
}

type AppMixed struct {
	ID                 model.InternalID
	Name               string
	Type               AppType
	ShortDescription   string
	IconImageURL       string
	BackgroundImageURL string
	CoverImageURL      string
	Tags               []string
	Details            *AppDetails
}

type AppDetails struct {
	Description string
	ReleaseDate string
	Developer   string
	Publisher   string
	Version     string
}

type AppSource int

const (
	AppSourceUnspecified AppSource = iota
	AppSourceInternal
	AppSourceSteam
)

type AppType int

const (
	AppTypeUnspecified AppType = iota
	AppTypeGame
)

type AppPackage struct {
	ID            model.InternalID
	Source        AppPackageSource
	SourceID      model.InternalID
	Name          string
	Description   string
	Binary        *AppPackageBinary
	Public        bool
	AssignedAppID model.InternalID
}

type AppPackageBinary struct {
	Name      string
	SizeBytes int64
	PublicURL string
	Sha256    []byte
}

type AppPackageSource int

const (
	AppPackageSourceUnspecified AppPackageSource = iota
	AppPackageSourceManual
	AppPackageSourceSentinel
)

type BoundApps struct {
	Internal *App
	Others   []*App
}

func (b *BoundApps) Flatten() *AppMixed {
	if b == nil {
		return nil
	}
	res := b.Internal
	for _, a := range b.Others {
		res = mergeApp(res, a)
	}
	return &AppMixed{
		ID:                 res.ID,
		Name:               res.Name,
		Type:               res.Type,
		ShortDescription:   res.ShortDescription,
		IconImageURL:       res.IconImageURL,
		BackgroundImageURL: res.BackgroundImageURL,
		CoverImageURL:      res.CoverImageURL,
		Tags:               res.Tags,
		Details:            res.Details,
	}
}

func mergeApp(base *App, merged *App) *App {
	if base == nil {
		base = merged
		return base
	}
	if merged == nil {
		return base
	}
	if len(base.Name) == 0 {
		base.Name = merged.Name
	}
	if base.Type == AppTypeUnspecified {
		base.Type = merged.Type
	}
	if len(base.ShortDescription) == 0 {
		base.ShortDescription = merged.ShortDescription
	}
	if len(base.IconImageURL) == 0 {
		base.IconImageURL = merged.IconImageURL
	}
	if len(base.BackgroundImageURL) == 0 {
		base.BackgroundImageURL = merged.BackgroundImageURL
	}
	if len(base.CoverImageURL) == 0 {
		base.CoverImageURL = merged.CoverImageURL
	}
	if base.Details == nil {
		base.Details = merged.Details
		return base
	}
	if merged.Details == nil {
		return base
	}
	if len(base.Details.Description) == 0 {
		base.Details.Description = merged.Details.Description
	}
	if len(base.Details.ReleaseDate) == 0 {
		base.Details.ReleaseDate = merged.Details.ReleaseDate
	}
	if len(base.Details.Developer) == 0 {
		base.Details.Developer = merged.Details.Developer
	}
	if len(base.Details.Publisher) == 0 {
		base.Details.Publisher = merged.Details.Publisher
	}
	if len(base.Details.Version) == 0 {
		base.Details.Version = merged.Details.Version
	}
	return base
}
