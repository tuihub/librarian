package model

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type FeedActionSet struct {
	ID          model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	UserID      model.InternalID `gorm:"column:user_feed_action_set"` // Inferred from edge.From("owner", User.Type).Ref("feed_action_set")
	Name        string
	Description string
	Actions     []*model.FeatureRequest `gorm:"serializer:json"`
	UpdatedAt   time.Time
	CreatedAt   time.Time
	Owner       *User        `gorm:"foreignKey:UserID"`
	FeedConfigs []FeedConfig `gorm:"many2many:feed_config_actions;"`
}

func (FeedActionSet) TableName() string {
	return "feed_action_sets"
}
