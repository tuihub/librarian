package modelyesod

import (
	"database/sql/driver"
	"errors"
	"time"

	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	"github.com/tuihub/librarian/internal/model/modelnetzach"
)

type FeedItemDigest struct {
	FeedID              model.InternalID
	ItemID              model.InternalID
	AvatarURL           string
	Authors             string
	PublishedParsedTime time.Time
	Title               string
	ShortDescription    string
	ImageUrls           []string
	PublishPlatform     string
	FeedConfigName      string
	FeedAvatarURL       string
	ReadCount           int64
}

type FeedWithConfig struct {
	FeedConfig *FeedConfig
	Feed       *modelfeed.Feed
}

type FeedConfig struct {
	ID                model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	UserFeedConfig    model.InternalID `gorm:"column:user_feed_config;index"` // Mapped to UserID
	Name              string
	Description       string
	Source            *model.FeatureRequest `gorm:"serializer:json"`
	ActionSets        []model.InternalID    `gorm:"-"` // Ignored, use FeedActionSets relation
	Category          string                `gorm:"index"`
	Status            FeedConfigStatus
	PullInterval      time.Duration
	LatestPullTime    time.Time
	LatestPullStatus  FeedConfigPullStatus
	LatestPullMessage string
	HideItems         bool
	NextPullBeginAt   time.Time
	UpdatedAt         time.Time
	CreatedAt         time.Time
	Owner             *model.User     `gorm:"foreignKey:UserFeedConfig"`
	Feed              *modelfeed.Feed `gorm:"foreignKey:ID;references:ID"`
	FeedActionSets    []FeedActionSet `gorm:"many2many:feed_config_actions;"`
}

func (FeedConfig) TableName() string {
	return "feed_configs"
}

// FeedConfigAction is a helper for join table.
type FeedConfigAction struct {
	FeedConfigID    model.InternalID `gorm:"primaryKey"`
	FeedActionSetID model.InternalID `gorm:"primaryKey"`
}

func (FeedConfigAction) TableName() string {
	return "feed_config_actions"
}

type FeedConfigStatus int

const (
	FeedConfigStatusUnspecified FeedConfigStatus = iota
	FeedConfigStatusActive
	FeedConfigStatusSuspend
)

func (s FeedConfigStatus) Value() (driver.Value, error) {
	switch s {
	case FeedConfigStatusActive:
		return "active", nil
	case FeedConfigStatusSuspend:
		return "suspend", nil
	default:
		return "", nil
	}
}

func (s *FeedConfigStatus) Scan(value interface{}) error {
	v, ok := value.(string)
	if !ok {
		return errors.New("invalid type for FeedConfigStatus")
	}
	switch v {
	case "active":
		*s = FeedConfigStatusActive
	case "suspend":
		*s = FeedConfigStatusSuspend
	default:
		*s = FeedConfigStatusUnspecified
	}
	return nil
}

type FeedConfigPullStatus int

const (
	FeedConfigPullStatusUnspecified FeedConfigPullStatus = iota
	FeedConfigPullStatusProcessing
	FeedConfigPullStatusSuccess
	FeedConfigPullStatusFailed
)

func (s FeedConfigPullStatus) Value() (driver.Value, error) {
	switch s {
	case FeedConfigPullStatusProcessing:
		return "processing", nil
	case FeedConfigPullStatusSuccess:
		return "success", nil
	case FeedConfigPullStatusFailed:
		return "failed", nil
	default:
		return "", nil
	}
}

func (s *FeedConfigPullStatus) Scan(value interface{}) error {
	v, ok := value.(string)
	if !ok {
		return errors.New("invalid type for FeedConfigPullStatus")
	}
	switch v {
	case "processing":
		*s = FeedConfigPullStatusProcessing
	case "success":
		*s = FeedConfigPullStatusSuccess
	case "failed":
		*s = FeedConfigPullStatusFailed
	default:
		*s = FeedConfigPullStatusUnspecified
	}
	return nil
}

type ListFeedOrder int

const (
	ListFeedOrderUnspecified ListFeedOrder = iota
	ListFeedOrderNextPull
)

type PullFeed struct {
	InternalID   model.InternalID
	Source       *model.FeatureRequest
	SystemNotify *modelnetzach.SystemNotify
}

type GroupFeedItemsBy int

const (
	GroupFeedItemsByUnspecified GroupFeedItemsBy = iota
	GroupFeedItemsByYear
	GroupFeedItemsByMonth
	GroupFeedItemsByDay
	GroupFeedItemsByOverall
)

type FeedItemCollection struct {
	ID          model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	UserID      model.InternalID `gorm:"index"`
	Name        string
	Description string
	Category    string
	FeedItems   []modelfeed.Item `gorm:"many2many:feed_item_collection_feed_items;"`
	UpdatedAt   time.Time
	CreatedAt   time.Time
}

func (FeedItemCollection) TableName() string {
	return "feed_item_collections"
}

// FeedItemCollectionFeedItem is a join table for FeedItemCollection.
type FeedItemCollectionFeedItem struct {
	FeedItemCollectionID model.InternalID `gorm:"primaryKey"`
	FeedItemID           model.InternalID `gorm:"primaryKey"`
}

func (FeedItemCollectionFeedItem) TableName() string {
	return "feed_item_collection_feed_items"
}

type FeedActionSet struct {
	ID          model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	UserID      model.InternalID `gorm:"index"`
	Name        string
	Description string
	Actions     []*model.FeatureRequest `gorm:"serializer:json"`
	FeedConfigs []FeedConfig            `gorm:"many2many:feed_config_actions;"`
	UpdatedAt   time.Time
	CreatedAt   time.Time
}

func (FeedActionSet) TableName() string {
	return "feed_action_sets"
}
