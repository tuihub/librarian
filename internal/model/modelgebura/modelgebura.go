package modelgebura

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type AppInfo struct {
	ID                 model.InternalID
	Source             string
	SourceAppID        string
	SourceURL          string
	Name               string
	Type               AppType
	ShortDescription   string
	Description        string
	IconImageURL       string
	IconImageID        model.InternalID
	BackgroundImageURL string
	BackgroundImageID  model.InternalID
	CoverImageURL      string
	CoverImageID       model.InternalID
	ReleaseDate        string
	Developer          string
	Publisher          string
	Tags               []string
	AlternativeNames   []string
	RawData            string
	UpdatedAt          time.Time
}

type AppInfoID struct {
	Source      string
	SourceAppID string
}

type AppInfoMixed struct {
	ID                 model.InternalID
	Name               string
	Type               AppType
	ShortDescription   string
	IconImageURL       string
	BackgroundImageURL string
	CoverImageURL      string
	Tags               []string
}

type AppType int

const (
	AppTypeUnspecified AppType = iota
	AppTypeGame
)

type App struct {
	ID                 model.InternalID
	VersionNumber      uint64
	VersionDate        time.Time
	CreatorDeviceID    model.InternalID
	AppSources         map[string]string
	Public             bool
	BoundStoreAppID    *model.InternalID
	StopStoreManage    *bool
	Name               string
	Type               AppType
	ShortDescription   string
	Description        string
	IconImageURL       string
	IconImageID        model.InternalID
	BackgroundImageURL string
	BackgroundImageID  model.InternalID
	CoverImageURL      string
	CoverImageID       model.InternalID
	ReleaseDate        string
	Developer          string
	Publisher          string
	Tags               []string
	AlternativeNames   []string
}

type AppRunTime struct {
	ID       model.InternalID
	AppID    model.InternalID
	DeviceID model.InternalID
	RunTime  *model.TimeRange
}

type AppBinary struct {
	Name      string
	SizeBytes int64
	PublicURL string
	Sha256    []byte
}

type AppBinaryChunk struct {
	Sequence  int64
	SizeBytes int64
	PublicURL string
	Sha256    []byte
}

type AppCategory struct {
	ID            model.InternalID
	VersionNumber uint64
	VersionDate   time.Time
	Name          string
	AppIDs        []model.InternalID
}

type SentinelInfo struct {
	ID                   model.InternalID
	URL                  string
	AlternativeUrls      []string
	GetTokenPath         string
	DownloadFileBasePath string
	Libraries            []*SentinelLibrary
}

type SentinelLibrary struct {
	ReportedID       int64
	DownloadBasePath string
	AppBinaries      []*SentinelAppBinary
}

type SentinelAppBinary struct {
	SentinelLibraryID int64
	GeneratedID       string
	SizeBytes         int64
	NeedToken         bool
	Files             []*SentinelAppBinaryFile
	Name              string
	Version           string
	Developer         string
	Publisher         string
}

type SentinelAppBinaryFile struct {
	Name           string
	SizeBytes      int64
	Sha256         []byte
	ServerFilePath string
	ChunksInfo     string
}
