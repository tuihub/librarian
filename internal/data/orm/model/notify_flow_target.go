package model

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type NotifyFlowTarget struct {
	NotifyFlowID          model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	NotifyTargetID        model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	FilterIncludeKeywords []string         `gorm:"serializer:json"`
	FilterExcludeKeywords []string         `gorm:"serializer:json"`
	UpdatedAt             time.Time
	CreatedAt             time.Time
}

func (NotifyFlowTarget) TableName() string {
	return "notify_flow_targets"
}
