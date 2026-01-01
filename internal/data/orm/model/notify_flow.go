package model

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type NotifyFlow struct {
	ID                model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	OwnerID           model.InternalID `gorm:"column:user_notify_flow"` // Inferred
	Name              string
	Description       string
	Status            string
	UpdatedAt         time.Time
	CreatedAt         time.Time
	Owner             *User              `gorm:"foreignKey:OwnerID"`
	NotifyTargets     []NotifyTarget     `gorm:"many2many:notify_flow_targets;"`
	NotifySources     []NotifySource     `gorm:"many2many:notify_flow_sources;"`
	NotifyFlowTargets []NotifyFlowTarget `gorm:"foreignKey:NotifyFlowID"`
	NotifyFlowSources []NotifyFlowSource `gorm:"foreignKey:NotifyFlowID"`
}

func (NotifyFlow) TableName() string {
	return "notify_flows"
}
