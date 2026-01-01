package model

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type FeedConfig struct {
	ID                model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	UserFeedConfig    model.InternalID `gorm:"column:user_feed_config;index"` // This is UserID
	Name              string
	Description       string
	Source            *model.FeatureRequest `gorm:"serializer:json"`
	Status            string
	Category          string `gorm:"index"`
	PullInterval      time.Duration
	HideItems         bool
	LatestPullAt      time.Time
	LatestPullStatus  string
	LatestPullMessage string
	NextPullBeginAt   time.Time
	UpdatedAt         time.Time
	CreatedAt         time.Time
	Owner             *User           `gorm:"foreignKey:UserFeedConfig"`
	Feed              *Feed           `gorm:"foreignKey:ID;references:ID"` // HasOne Feed? ent: edge.To("feed", Feed.Type).Unique()
	NotifySource      []NotifySource  `gorm:"foreignKey:ID;references:ID"` // Polymorphic? Check ent schema.
	FeedActionSets    []FeedActionSet `gorm:"many2many:feed_config_actions;"`
}

func (FeedConfig) TableName() string {
	return "feed_configs"
}
