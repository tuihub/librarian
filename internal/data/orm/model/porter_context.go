package model

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type PorterContext struct {
	ID                  model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	OwnerID             model.InternalID `gorm:"column:user_porter_context"` // Inferred
	GlobalName          string           `gorm:"index:idx_porter_context_global_name_region,priority:1"`
	Region              string           `gorm:"index:idx_porter_context_global_name_region,priority:2"`
	ContextJSON         string
	Name                string
	Description         string
	Status              string
	HandleStatus        string
	HandleStatusMessage string
	UpdatedAt           time.Time
	CreatedAt           time.Time
	Owner               *User `gorm:"foreignKey:OwnerID"`
}

func (PorterContext) TableName() string {
	return "porter_contexts"
}
