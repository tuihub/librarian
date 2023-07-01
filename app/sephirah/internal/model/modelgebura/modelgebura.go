package modelgebura

import "github.com/tuihub/librarian/internal/model"

type App struct {
	ID               model.InternalID
	Source           AppSource
	SourceAppID      string
	SourceURL        string
	Name             string
	Type             AppType
	ShortDescription string
	ImageURL         string
	Details          *AppDetails
	// the bound internal app id if self is external
	BoundInternal model.InternalID
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
	ID          model.InternalID
	Source      AppPackageSource
	SourceID    model.InternalID
	Name        string
	Description string
	Binary      *AppPackageBinary
	Public      bool
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
