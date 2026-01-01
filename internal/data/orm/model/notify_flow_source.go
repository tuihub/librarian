package model

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type NotifyFlowSource struct {
	NotifyFlowID          model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	NotifySourceID        model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	FilterIncludeKeywords []string         `gorm:"serializer:json"`
	FilterExcludeKeywords []string         `gorm:"serializer:json"`
	UpdatedAt             time.Time
	CreatedAt             time.Time
}

func (NotifyFlowSource) TableName() string {
	return "notify_flow_sources"
}
