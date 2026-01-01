package model

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type NotifyTarget struct {
	ID          model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	OwnerID     model.InternalID `gorm:"column:user_notify_target"` // Inferred
	Name        string
	Description string
	Destination *model.FeatureRequest `gorm:"serializer:json"`
	Status      string
	UpdatedAt   time.Time
	CreatedAt   time.Time
	Owner       *User        `gorm:"foreignKey:OwnerID"`
	NotifyFlows []NotifyFlow `gorm:"many2many:notify_flow_targets;"`
}

func (NotifyTarget) TableName() string {
	return "notify_targets"
}
