package model

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type FeedConfigAction struct {
	FeedConfigID    model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	FeedActionSetID model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	Index           int64
	UpdatedAt       time.Time
	CreatedAt       time.Time
}

func (FeedConfigAction) TableName() string {
	return "feed_config_actions"
}
