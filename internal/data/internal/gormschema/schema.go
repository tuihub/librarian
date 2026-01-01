// Package gormschema defines GORM models for database tables.
// These models derive their business fields from internal/model definitions.
package gormschema

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	"github.com/tuihub/librarian/internal/model/modelsupervisor"
)

// BaseTimestamps provides common timestamp fields for all entities.
type BaseTimestamps struct {
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// ============================================
// User/Account/Session related models
// ============================================

// User represents the users table.
type User struct {
	ID        model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	Username  string           `gorm:"uniqueIndex;not null"`
	Password  string           `gorm:"not null"`
	Status    string           `gorm:"not null"` // active, blocked
	Type      string           `gorm:"not null"` // admin, normal
	CreatorID model.InternalID `gorm:"not null"`
	BaseTimestamps
}

func (User) TableName() string { return "users" }

// Device represents the devices table.
type Device struct {
	ID                      model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	DeviceName              string           `gorm:"not null"`
	SystemType              string           `gorm:"not null"` // ios, android, web, windows, macos, linux, unknown
	SystemVersion           string           `gorm:"not null"`
	ClientName              string           `gorm:"not null"`
	ClientSourceCodeAddress string           `gorm:"not null"`
	ClientVersion           string           `gorm:"not null"`
	ClientLocalID           *string
	BaseTimestamps
}

func (Device) TableName() string { return "devices" }

// Session represents the sessions table.
type Session struct {
	ID           model.InternalID  `gorm:"primaryKey;autoIncrement:false"`
	UserID       model.InternalID  `gorm:"not null;uniqueIndex:idx_session_user_device"`
	DeviceID     *model.InternalID `gorm:"uniqueIndex:idx_session_user_device"`
	RefreshToken string            `gorm:"uniqueIndex;not null"`
	ExpireAt     time.Time         `gorm:"not null"`
	BaseTimestamps
}

func (Session) TableName() string { return "sessions" }

// Account represents the accounts table.
type Account struct {
	ID                model.InternalID  `gorm:"primaryKey;autoIncrement:false"`
	Platform          string            `gorm:"not null;uniqueIndex:idx_account_platform_id"`
	PlatformAccountID string            `gorm:"not null;uniqueIndex:idx_account_platform_id"`
	BoundUserID       *model.InternalID
	Name              string `gorm:"not null"`
	ProfileURL        string `gorm:"not null"`
	AvatarURL         string `gorm:"not null"`
	BaseTimestamps
}

func (Account) TableName() string { return "accounts" }

// ============================================
// Porter related models
// ============================================

// PorterInstance represents the porter_instances table.
type PorterInstance struct {
	ID                      model.InternalID         `gorm:"primaryKey;autoIncrement:false"`
	Name                    string                   `gorm:"not null"`
	Version                 string                   `gorm:"not null"`
	Description             string                   `gorm:"not null"`
	SourceCodeAddress       string                   `gorm:"not null"`
	BuildVersion            string                   `gorm:"not null"`
	BuildDate               string                   `gorm:"not null"`
	GlobalName              string                   `gorm:"not null;index:idx_porter_global_region"`
	Address                 string                   `gorm:"uniqueIndex;not null"`
	Region                  string                   `gorm:"not null;index:idx_porter_global_region"`
	FeatureSummary          *PorterFeatureSummaryVal `gorm:"type:text"`
	ContextJSONSchema       string
	Status                  string `gorm:"not null"` // active, blocked
	ConnectionStatus        string `gorm:"not null"` // unspecified, queueing, connected, disconnected, active, activation_failed, downgraded
	ConnectionStatusMessage string
	BaseTimestamps
}

func (PorterInstance) TableName() string { return "porter_instances" }

// PorterFeatureSummaryVal is a custom type for JSON serialization of PorterFeatureSummary.
type PorterFeatureSummaryVal modelsupervisor.PorterFeatureSummary

func (p PorterFeatureSummaryVal) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *PorterFeatureSummaryVal) Scan(value any) error {
	if value == nil {
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		s, ok2 := value.(string)
		if !ok2 {
			return errors.New("type assertion to []byte or string failed")
		}
		b = []byte(s)
	}
	return json.Unmarshal(b, p)
}

// PorterContext represents the porter_contexts table.
type PorterContext struct {
	ID                  model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	OwnerID             model.InternalID `gorm:"not null"`
	GlobalName          string           `gorm:"not null"`
	Region              string           `gorm:"not null"`
	ContextJSON         string
	Name                string `gorm:"not null"`
	Description         string
	Status              string `gorm:"not null"` // active, disabled
	HandleStatus        string `gorm:"not null"` // unspecified, active, downgraded, queueing, blocked
	HandleStatusMessage string
	BaseTimestamps
}

func (PorterContext) TableName() string { return "porter_contexts" }

// ============================================
// App related models
// ============================================

// App represents the apps table.
type App struct {
	ID                 model.InternalID  `gorm:"primaryKey;autoIncrement:false"`
	UserID             model.InternalID  `gorm:"not null;index"`
	VersionNumber      uint64            `gorm:"not null"`
	VersionDate        time.Time         `gorm:"not null"`
	CreatorDeviceID    model.InternalID  `gorm:"not null"`
	AppSources         StringMapVal      `gorm:"type:text"`
	Public             bool              `gorm:"not null"`
	BoundStoreAppID    *model.InternalID
	StopStoreManage    *bool
	Name               string `gorm:"not null"`
	Type               string `gorm:"not null"` // unknown, game
	ShortDescription   string
	Description        string `gorm:"type:text"`
	IconImageURL       string
	IconImageID        model.InternalID
	BackgroundImageURL string
	BackgroundImageID  model.InternalID
	CoverImageURL      string
	CoverImageID       model.InternalID
	ReleaseDate        string
	Developer          string
	Publisher          string
	Tags               StringArrayVal `gorm:"type:text"`
	AlternativeNames   StringArrayVal `gorm:"type:text"`
	BaseTimestamps
}

func (App) TableName() string { return "apps" }

// AppInfo represents the app_infos table.
type AppInfo struct {
	ID                 model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	Source             string           `gorm:"not null;uniqueIndex:idx_appinfo_source_id"`
	SourceAppID        string           `gorm:"not null;uniqueIndex:idx_appinfo_source_id"`
	SourceURL          string
	Name               string `gorm:"not null"`
	Type               string `gorm:"not null"` // unknown, game
	ShortDescription   string
	Description        string `gorm:"type:text"`
	IconImageURL       string
	IconImageID        model.InternalID
	BackgroundImageURL string
	BackgroundImageID  model.InternalID
	CoverImageURL      string
	CoverImageID       model.InternalID
	ReleaseDate        string
	Developer          string
	Publisher          string
	Tags               StringArrayVal `gorm:"type:text"`
	AlternativeNames   StringArrayVal `gorm:"type:text"`
	RawData            string         `gorm:"type:text"`
	BaseTimestamps
}

func (AppInfo) TableName() string { return "app_infos" }

// AppRunTime represents the app_run_times table.
type AppRunTime struct {
	ID        model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	UserID    model.InternalID `gorm:"not null"`
	AppID     model.InternalID `gorm:"not null"`
	DeviceID  model.InternalID
	StartTime time.Time     `gorm:"not null"`
	Duration  time.Duration `gorm:"not null"`
	BaseTimestamps
}

func (AppRunTime) TableName() string { return "app_run_times" }

// AppCategory represents the app_categories table.
type AppCategory struct {
	ID            model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	UserID        model.InternalID `gorm:"not null"`
	VersionNumber uint64           `gorm:"not null"`
	VersionDate   time.Time        `gorm:"not null"`
	Name          string           `gorm:"not null"`
	BaseTimestamps
}

func (AppCategory) TableName() string { return "app_categories" }

// AppAppCategory represents the app_app_categories table (many-to-many join table).
type AppAppCategory struct {
	ID            model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	AppID         model.InternalID `gorm:"not null;uniqueIndex:idx_aac_app_category"`
	AppCategoryID model.InternalID `gorm:"not null;uniqueIndex:idx_aac_app_category"`
	BaseTimestamps
}

func (AppAppCategory) TableName() string { return "app_app_categories" }

// ============================================
// Sentinel related models
// ============================================

// Sentinel represents the sentinels table.
type Sentinel struct {
	ID                    model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	CreatorID             model.InternalID `gorm:"not null"`
	Name                  string           `gorm:"not null"`
	Description           string
	URL                   string
	AlternativeUrls       StringArrayVal `gorm:"type:text"`
	GetTokenPath          string
	DownloadFileBasePath  string
	LibraryReportSequence int64
	BaseTimestamps
}

func (Sentinel) TableName() string { return "sentinels" }

// SentinelSession represents the sentinel_sessions table.
type SentinelSession struct {
	ID              model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	SentinelID      model.InternalID `gorm:"not null"`
	RefreshToken    string           `gorm:"not null"`
	Status          string           `gorm:"not null"` // active, suspend
	CreatorID       model.InternalID `gorm:"not null"`
	ExpireAt        time.Time        `gorm:"not null"`
	LastUsedAt      *time.Time
	LastRefreshedAt *time.Time
	RefreshCount    int64
	BaseTimestamps
}

func (SentinelSession) TableName() string { return "sentinel_sessions" }

// SentinelLibrary represents the sentinel_libraries table.
type SentinelLibrary struct {
	ID                    model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	SentinelID            model.InternalID `gorm:"not null;uniqueIndex:idx_sentinel_lib"`
	ReportedID            int64            `gorm:"not null;uniqueIndex:idx_sentinel_lib"`
	DownloadBasePath      string
	LibraryReportSequence int64
	ActiveSnapshot        time.Time
	BaseTimestamps
}

func (SentinelLibrary) TableName() string { return "sentinel_libraries" }

// SentinelAppBinary represents the sentinel_app_binaries table.
type SentinelAppBinary struct {
	ID                        model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	UnionID                   string           `gorm:"not null"`
	SentinelID                model.InternalID `gorm:"not null;uniqueIndex:idx_sab_unique"`
	SentinelLibraryReportedID int64            `gorm:"not null;uniqueIndex:idx_sab_unique"`
	LibrarySnapshot           time.Time        `gorm:"not null;uniqueIndex:idx_sab_unique"`
	GeneratedID               string           `gorm:"not null;uniqueIndex:idx_sab_unique"`
	SizeBytes                 int64            `gorm:"not null"`
	NeedToken                 bool             `gorm:"not null"`
	Name                      string           `gorm:"not null"`
	Version                   string
	Developer                 string
	Publisher                 string
	StoreAppID                *model.InternalID
	BaseTimestamps
}

func (SentinelAppBinary) TableName() string { return "sentinel_app_binaries" }

// SentinelAppBinaryFile represents the sentinel_app_binary_files table.
type SentinelAppBinaryFile struct {
	ID                           model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	SentinelID                   model.InternalID `gorm:"not null;uniqueIndex:idx_sabf_unique"`
	SentinelLibraryReportedID    int64            `gorm:"not null;uniqueIndex:idx_sabf_unique"`
	LibrarySnapshot              time.Time        `gorm:"not null;uniqueIndex:idx_sabf_unique"`
	SentinelAppBinaryGeneratedID string           `gorm:"not null;uniqueIndex:idx_sabf_unique"`
	Name                         string           `gorm:"not null"`
	SizeBytes                    int64            `gorm:"not null"`
	Sha256                       []byte
	ServerFilePath               string `gorm:"not null;uniqueIndex:idx_sabf_unique"`
	ChunksInfo                   string `gorm:"type:text"`
	BaseTimestamps
}

func (SentinelAppBinaryFile) TableName() string { return "sentinel_app_binary_files" }

// ============================================
// Store related models
// ============================================

// StoreApp represents the store_apps table.
type StoreApp struct {
	ID          model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	Name        string           `gorm:"not null"`
	Description string
	BaseTimestamps
}

func (StoreApp) TableName() string { return "store_apps" }

// StoreAppBinary is not separate table - uses SentinelAppBinary with StoreAppID.

// ============================================
// Feed related models
// ============================================

// Feed represents the feeds table.
type Feed struct {
	ID          model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	Title       string
	Link        string
	Description string
	Language    string
	Authors     FeedPersonArrayVal `gorm:"type:text"`
	Image       *FeedImageVal      `gorm:"type:text"`
	BaseTimestamps
}

func (Feed) TableName() string { return "feeds" }

// FeedPersonArrayVal is a custom type for JSON serialization.
type FeedPersonArrayVal []*modelfeed.Person

func (p FeedPersonArrayVal) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return json.Marshal(p)
}

func (p *FeedPersonArrayVal) Scan(value any) error {
	if value == nil {
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		s, ok2 := value.(string)
		if !ok2 {
			return errors.New("type assertion to []byte or string failed")
		}
		b = []byte(s)
	}
	return json.Unmarshal(b, p)
}

// FeedImageVal is a custom type for JSON serialization.
type FeedImageVal modelfeed.Image

func (p FeedImageVal) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *FeedImageVal) Scan(value any) error {
	if value == nil {
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		s, ok2 := value.(string)
		if !ok2 {
			return errors.New("type assertion to []byte or string failed")
		}
		b = []byte(s)
	}
	return json.Unmarshal(b, p)
}

// FeedConfig represents the feed_configs table.
type FeedConfig struct {
	ID                model.InternalID   `gorm:"primaryKey;autoIncrement:false"`
	OwnerID           model.InternalID   `gorm:"not null"`
	Name              string             `gorm:"not null"`
	Description       string
	Category          string
	Source            *FeatureRequestVal `gorm:"type:text"`
	Status            string             `gorm:"not null"` // active, suspend
	PullInterval      time.Duration      `gorm:"not null"`
	LatestPullAt      time.Time
	LatestPullStatus  string `gorm:"not null"` // processing, success, failed
	LatestPullMessage string
	NextPullBeginAt   time.Time
	HideItems         bool `gorm:"not null"`
	BaseTimestamps
}

func (FeedConfig) TableName() string { return "feed_configs" }

// FeatureRequestVal is a custom type for JSON serialization.
type FeatureRequestVal model.FeatureRequest

func (p FeatureRequestVal) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *FeatureRequestVal) Scan(value any) error {
	if value == nil {
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		s, ok2 := value.(string)
		if !ok2 {
			return errors.New("type assertion to []byte or string failed")
		}
		b = []byte(s)
	}
	return json.Unmarshal(b, p)
}

// FeedConfigAction represents the feed_config_actions table.
type FeedConfigAction struct {
	ID              model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	FeedConfigID    model.InternalID `gorm:"not null"`
	FeedActionSetID model.InternalID `gorm:"not null"`
	Index           int64            `gorm:"not null"`
	BaseTimestamps
}

func (FeedConfigAction) TableName() string { return "feed_config_actions" }

// FeedItem represents the feed_items table.
type FeedItem struct {
	ID                model.InternalID   `gorm:"primaryKey;autoIncrement:false"`
	FeedID            model.InternalID   `gorm:"not null;uniqueIndex:idx_feeditem_feed_guid"`
	Title             string             `gorm:"not null"`
	Description       string             `gorm:"type:text"`
	Content           string             `gorm:"type:text"`
	Link              string
	Updated           string
	UpdatedParsed     *time.Time
	Published         string
	PublishedParsed   time.Time `gorm:"not null;index"`
	Authors           FeedPersonArrayVal `gorm:"type:text"`
	GUID              string             `gorm:"not null;uniqueIndex:idx_feeditem_feed_guid"`
	Image             *FeedImageVal      `gorm:"type:text"`
	Enclosures        FeedEnclosureArrayVal `gorm:"type:text"`
	PublishPlatform   string
	ReadCount         int64
	DigestDescription string           `gorm:"type:text"`
	DigestImages      FeedImageArrayVal `gorm:"type:text"`
	BaseTimestamps
}

func (FeedItem) TableName() string { return "feed_items" }

// FeedEnclosureArrayVal is a custom type for JSON serialization.
type FeedEnclosureArrayVal []*modelfeed.Enclosure

func (p FeedEnclosureArrayVal) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return json.Marshal(p)
}

func (p *FeedEnclosureArrayVal) Scan(value any) error {
	if value == nil {
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		s, ok2 := value.(string)
		if !ok2 {
			return errors.New("type assertion to []byte or string failed")
		}
		b = []byte(s)
	}
	return json.Unmarshal(b, p)
}

// FeedImageArrayVal is a custom type for JSON serialization.
type FeedImageArrayVal []*modelfeed.Image

func (p FeedImageArrayVal) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return json.Marshal(p)
}

func (p *FeedImageArrayVal) Scan(value any) error {
	if value == nil {
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		s, ok2 := value.(string)
		if !ok2 {
			return errors.New("type assertion to []byte or string failed")
		}
		b = []byte(s)
	}
	return json.Unmarshal(b, p)
}

// FeedActionSet represents the feed_action_sets table.
type FeedActionSet struct {
	ID          model.InternalID       `gorm:"primaryKey;autoIncrement:false"`
	OwnerID     model.InternalID       `gorm:"not null"`
	Name        string                 `gorm:"not null"`
	Description string
	Actions     FeatureRequestArrayVal `gorm:"type:text"`
	BaseTimestamps
}

func (FeedActionSet) TableName() string { return "feed_action_sets" }

// FeatureRequestArrayVal is a custom type for JSON serialization.
type FeatureRequestArrayVal []*model.FeatureRequest

func (p FeatureRequestArrayVal) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return json.Marshal(p)
}

func (p *FeatureRequestArrayVal) Scan(value any) error {
	if value == nil {
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		s, ok2 := value.(string)
		if !ok2 {
			return errors.New("type assertion to []byte or string failed")
		}
		b = []byte(s)
	}
	return json.Unmarshal(b, p)
}

// FeedItemCollection represents the feed_item_collections table.
type FeedItemCollection struct {
	ID          model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	OwnerID     model.InternalID `gorm:"not null"`
	Name        string           `gorm:"not null"`
	Description string
	Category    string
	BaseTimestamps
}

func (FeedItemCollection) TableName() string { return "feed_item_collections" }

// FeedItemCollectionItem represents the many-to-many relationship between FeedItemCollection and FeedItem.
type FeedItemCollectionItem struct {
	FeedItemCollectionID model.InternalID `gorm:"primaryKey"`
	FeedItemID           model.InternalID `gorm:"primaryKey"`
}

func (FeedItemCollectionItem) TableName() string { return "feed_item_collection_feed_items" }

// ============================================
// Notify related models
// ============================================

// NotifySource represents the notify_sources table.
type NotifySource struct {
	ID      model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	OwnerID model.InternalID `gorm:"not null"`
	BaseTimestamps
}

func (NotifySource) TableName() string { return "notify_sources" }

// NotifyTarget represents the notify_targets table.
type NotifyTarget struct {
	ID          model.InternalID   `gorm:"primaryKey;autoIncrement:false"`
	OwnerID     model.InternalID   `gorm:"not null"`
	Name        string             `gorm:"not null"`
	Description string
	Destination *FeatureRequestVal `gorm:"type:text"`
	Status      string             `gorm:"not null"` // active, suspend
	BaseTimestamps
}

func (NotifyTarget) TableName() string { return "notify_targets" }

// NotifyFlow represents the notify_flows table.
type NotifyFlow struct {
	ID          model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	OwnerID     model.InternalID `gorm:"not null"`
	Name        string           `gorm:"not null"`
	Description string
	Status      string `gorm:"not null"` // active, suspend
	BaseTimestamps
}

func (NotifyFlow) TableName() string { return "notify_flows" }

// NotifyFlowSource represents the notify_flow_sources table.
type NotifyFlowSource struct {
	ID                    model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	NotifyFlowID          model.InternalID `gorm:"not null"`
	NotifySourceID        model.InternalID `gorm:"not null"`
	FilterExcludeKeywords StringArrayVal   `gorm:"type:text"`
	FilterIncludeKeywords StringArrayVal   `gorm:"type:text"`
	BaseTimestamps
}

func (NotifyFlowSource) TableName() string { return "notify_flow_sources" }

// NotifyFlowTarget represents the notify_flow_targets table.
type NotifyFlowTarget struct {
	ID                    model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	NotifyFlowID          model.InternalID `gorm:"not null"`
	NotifyTargetID        model.InternalID `gorm:"not null"`
	FilterExcludeKeywords StringArrayVal   `gorm:"type:text"`
	FilterIncludeKeywords StringArrayVal   `gorm:"type:text"`
	BaseTimestamps
}

func (NotifyFlowTarget) TableName() string { return "notify_flow_targets" }

// SystemNotification represents the system_notifications table.
type SystemNotification struct {
	ID      model.InternalID  `gorm:"primaryKey;autoIncrement:false"`
	UserID  *model.InternalID
	Type    string `gorm:"not null"` // system, user
	Level   string `gorm:"not null"` // info, warn, error, ongoing
	Status  string `gorm:"not null"` // unread, read, dismissed
	Title   string `gorm:"not null"`
	Content string `gorm:"type:text"`
	BaseTimestamps
}

func (SystemNotification) TableName() string { return "system_notifications" }

// ============================================
// Image/File related models
// ============================================

// Image represents the images table.
type Image struct {
	ID          model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	OwnerID     model.InternalID `gorm:"not null"`
	FileID      model.InternalID `gorm:"not null"`
	Name        string           `gorm:"not null"`
	Description string
	Status      string `gorm:"not null"` // uploaded, scanned
	BaseTimestamps
}

func (Image) TableName() string { return "images" }

// File represents the files table.
type File struct {
	ID      model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	OwnerID model.InternalID `gorm:"not null"`
	Name    string           `gorm:"not null"`
	Size    int64
	Type    string
	Sha256  string
	BaseTimestamps
}

func (File) TableName() string { return "files" }

// Tag represents the tags table.
type Tag struct {
	ID          model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	OwnerID     model.InternalID `gorm:"not null"`
	Name        string           `gorm:"not null"`
	Description string
	BaseTimestamps
}

func (Tag) TableName() string { return "tags" }

// ============================================
// KV related models
// ============================================

// KV represents the kvs table for key-value storage.
type KV struct {
	ID     int64  `gorm:"primaryKey;autoIncrement"`
	Bucket string `gorm:"not null;uniqueIndex:idx_kv_bucket_key"`
	Key    string `gorm:"not null;uniqueIndex:idx_kv_bucket_key"`
	Value  string `gorm:"type:text;not null"`
	BaseTimestamps
}

func (KV) TableName() string { return "kvs" }

// ============================================
// Custom value types for JSON serialization
// ============================================

// StringArrayVal is a custom type for storing []string as JSON.
type StringArrayVal []string

func (s StringArrayVal) Value() (driver.Value, error) {
	if s == nil {
		return nil, nil
	}
	return json.Marshal(s)
}

func (s *StringArrayVal) Scan(value any) error {
	if value == nil {
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		str, ok2 := value.(string)
		if !ok2 {
			return errors.New("type assertion to []byte or string failed")
		}
		b = []byte(str)
	}
	return json.Unmarshal(b, s)
}

// StringMapVal is a custom type for storing map[string]string as JSON.
type StringMapVal map[string]string

func (s StringMapVal) Value() (driver.Value, error) {
	if s == nil {
		return nil, nil
	}
	return json.Marshal(s)
}

func (s *StringMapVal) Scan(value any) error {
	if value == nil {
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		str, ok2 := value.(string)
		if !ok2 {
			return errors.New("type assertion to []byte or string failed")
		}
		b = []byte(str)
	}
	return json.Unmarshal(b, s)
}

// AllModels returns all GORM models for auto-migration.
func AllModels() []any {
	return []any{
		&User{},
		&Device{},
		&Session{},
		&Account{},
		&PorterInstance{},
		&PorterContext{},
		&App{},
		&AppInfo{},
		&AppRunTime{},
		&AppCategory{},
		&AppAppCategory{},
		&Sentinel{},
		&SentinelSession{},
		&SentinelLibrary{},
		&SentinelAppBinary{},
		&SentinelAppBinaryFile{},
		&StoreApp{},
		&Feed{},
		&FeedConfig{},
		&FeedConfigAction{},
		&FeedItem{},
		&FeedActionSet{},
		&FeedItemCollection{},
		&FeedItemCollectionItem{},
		&NotifySource{},
		&NotifyTarget{},
		&NotifyFlow{},
		&NotifyFlowSource{},
		&NotifyFlowTarget{},
		&SystemNotification{},
		&Image{},
		&File{},
		&Tag{},
		&KV{},
	}
}
