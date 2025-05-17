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

type Sentinel struct {
	ID                   model.InternalID
	Name                 string
	Description          string
	URL                  string
	AlternativeUrls      []string
	GetTokenPath         string
	DownloadFileBasePath string
	Libraries            []*SentinelLibrary
}

type SentinelSessionStatus int

const (
	SentinelSessionStatusUnspecified SentinelSessionStatus = iota
	SentinelSessionStatusActive
	SentinelSessionStatusSuspend
)

type SentinelSession struct {
	ID              model.InternalID
	SentinelID      model.InternalID
	RefreshToken    string
	Status          SentinelSessionStatus
	CreatorID       model.InternalID
	ExpireAt        time.Time
	LastUsedAt      *time.Time
	LastRefreshedAt *time.Time
	RefreshCount    int64
}

type SentinelLibrary struct {
	ID               model.InternalID
	ReportedID       int64
	DownloadBasePath string
	AppBinaries      []*SentinelAppBinary
}

type SentinelAppBinary struct {
	ID                model.InternalID
	UnionID           string
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
	ID             model.InternalID
	Name           string
	SizeBytes      int64
	Sha256         []byte
	ServerFilePath string
	ChunksInfo     string
}

type StoreApp struct {
	ID          model.InternalID
	Name        string
	Description string
}

type StoreAppBinary struct {
	ID        model.InternalID
	AppID     *model.InternalID
	UnionID   string
	SizeBytes int64
	NeedToken bool
	Name      string
	Version   string
	Developer string
	Publisher string
}
