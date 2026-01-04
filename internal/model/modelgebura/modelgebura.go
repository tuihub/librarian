package modelgebura

import (
	"database/sql/driver"
	"errors"
	"time"

	"github.com/tuihub/librarian/internal/model"

	"gorm.io/gorm"
)

type AppInfo struct {
	ID                 model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	Source             string           `gorm:"index:idx_app_info_source_source_app_id,priority:1"`
	SourceAppID        string           `gorm:"index:idx_app_info_source_source_app_id,priority:2"`
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
	Tags               []string `gorm:"serializer:json"`
	AlternativeNames   []string `gorm:"serializer:json"`
	RawData            string
	UpdatedAt          time.Time
	CreatedAt          time.Time
}

func (AppInfo) TableName() string {
	return "app_infos"
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

func (t AppType) Value() (driver.Value, error) {
	switch t {
	case AppTypeGame:
		return "game", nil
	default:
		return "", nil
	}
}

func (t *AppType) Scan(value interface{}) error {
	v, ok := value.(string)
	if !ok {
		return errors.New("invalid type for AppType")
	}
	switch v {
	case "game":
		*t = AppTypeGame
	default:
		*t = AppTypeUnspecified
	}
	return nil
}

type App struct {
	ID                 model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	UserID             model.InternalID `gorm:"index"`
	VersionNumber      uint64
	VersionDate        time.Time
	CreatorDeviceID    model.InternalID
	AppSources         map[string]string `gorm:"serializer:json"`
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
	Tags               []string `gorm:"serializer:json"`
	AlternativeNames   []string `gorm:"serializer:json"`
	UpdatedAt          time.Time
	CreatedAt          time.Time
	User               *model.User   `gorm:"foreignKey:UserID"`
	AppRunTime         []AppRunTime  `gorm:"foreignKey:AppID"`
	AppCategories      []AppCategory `gorm:"many2many:app_app_categories;"`
}

func (App) TableName() string {
	return "apps"
}

type AppRunTime struct {
	ID        model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	UserID    model.InternalID `gorm:"index"` // Added UserID
	AppID     model.InternalID
	DeviceID  model.InternalID
	RunTime   *model.TimeRange `gorm:"-"` // Handled by hooks
	StartTime time.Time        `gorm:"column:start_time"`
	Duration  time.Duration    `gorm:"column:duration"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (AppRunTime) TableName() string {
	return "app_run_times"
}

func (a *AppRunTime) BeforeSave(tx *gorm.DB) error {
	if a.RunTime != nil {
		a.StartTime = a.RunTime.StartTime
		a.Duration = a.RunTime.Duration
	}
	return nil
}

func (a *AppRunTime) AfterFind(tx *gorm.DB) error {
	a.RunTime = &model.TimeRange{
		StartTime: a.StartTime,
		Duration:  a.Duration,
	}
	return nil
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
	ID               model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	UserID           model.InternalID `gorm:"index"`
	VersionNumber    uint64
	VersionDate      time.Time
	Name             string
	AppIDs           []model.InternalID `gorm:"-"` // Ignore this, use Apps relation
	Apps             []App              `gorm:"many2many:app_app_categories;"`
	AppAppCategories []AppAppCategory   `gorm:"foreignKey:AppCategoryID"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (AppCategory) TableName() string {
	return "app_categories"
}

// AppAppCategory is a helper struct for many2many join table if needed explicitly, otherwise GORM handles it.
// ORM had AppAppCategory.
type AppAppCategory struct {
	AppID         model.InternalID `gorm:"primaryKey"`
	AppCategoryID model.InternalID `gorm:"primaryKey"`
}

func (AppAppCategory) TableName() string {
	return "app_app_categories"
}

type Sentinel struct {
	ID                    model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	Name                  string
	Description           string
	URL                   string
	AlternativeUrls       []string `gorm:"serializer:json"`
	GetTokenPath          string
	DownloadFileBasePath  string
	Libraries             []*SentinelLibrary `gorm:"foreignKey:SentinelID"`
	CreatorID             model.InternalID
	LibraryReportSequence int64
	UpdatedAt             time.Time
	CreatedAt             time.Time
	SentinelSessions      []SentinelSession `gorm:"foreignKey:SentinelID"`
}

func (Sentinel) TableName() string {
	return "sentinels"
}

type SentinelSessionStatus int

const (
	SentinelSessionStatusUnspecified SentinelSessionStatus = iota
	SentinelSessionStatusActive
	SentinelSessionStatusSuspend
)

func (s SentinelSessionStatus) Value() (driver.Value, error) {
	switch s {
	case SentinelSessionStatusActive:
		return "active", nil
	case SentinelSessionStatusSuspend:
		return "suspend", nil
	default:
		return "", nil
	}
}

func (s *SentinelSessionStatus) Scan(value interface{}) error {
	v, ok := value.(string)
	if !ok {
		return errors.New("invalid type for SentinelSessionStatus")
	}
	switch v {
	case "active":
		*s = SentinelSessionStatusActive
	case "suspend":
		*s = SentinelSessionStatusSuspend
	default:
		*s = SentinelSessionStatusUnspecified
	}
	return nil
}

type SentinelSession struct {
	ID              model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	SentinelID      model.InternalID
	RefreshToken    string
	Status          SentinelSessionStatus
	CreatorID       model.InternalID
	ExpireAt        time.Time
	LastUsedAt      *time.Time
	LastRefreshedAt *time.Time
	RefreshCount    int64
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (SentinelSession) TableName() string {
	return "sentinel_sessions"
}

type SentinelLibrary struct {
	ID               model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	SentinelID       model.InternalID
	ReportedID       int64
	DownloadBasePath string
	AppBinaries      []*SentinelAppBinary `gorm:"foreignKey:SentinelLibraryID"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (SentinelLibrary) TableName() string {
	return "sentinel_libraries"
}

type SentinelAppBinary struct {
	ID                model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	UnionID           string
	SentinelLibraryID int64
	GeneratedID       string
	SizeBytes         int64
	NeedToken         bool
	Files             []*SentinelAppBinaryFile `gorm:"foreignKey:SentinelAppBinaryID"`
	Name              string
	Version           string
	Developer         string
	Publisher         string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (SentinelAppBinary) TableName() string {
	return "sentinel_app_binaries"
}

type SentinelAppBinaryFile struct {
	ID                  model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	SentinelAppBinaryID model.InternalID
	Name                string
	SizeBytes           int64
	Sha256              []byte
	ServerFilePath      string
	ChunksInfo          string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func (SentinelAppBinaryFile) TableName() string {
	return "sentinel_app_binary_files"
}

type StoreApp struct {
	ID          model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	Source      string
	SourceAppID string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (StoreApp) TableName() string {
	return "store_apps"
}

type StoreAppBinary struct {
	ID        model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	AppID     *model.InternalID
	UnionID   string
	SizeBytes int64
	NeedToken bool
	Name      string
	Version   string
	Developer string
	Publisher string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (StoreAppBinary) TableName() string {
	return "store_app_binaries"
}
