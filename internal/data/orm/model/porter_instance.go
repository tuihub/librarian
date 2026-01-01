package model

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelsupervisor"
)

type PorterInstance struct {
	ID                      model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	Name                    string
	Version                 string
	Description             string
	SourceCodeAddress       string
	BuildVersion            string
	BuildDate               string
	GlobalName              string                                `gorm:"index:idx_porter_instance_global_name_region,priority:1"`
	Address                 string                                `gorm:"uniqueIndex"`
	Region                  string                                `gorm:"index:idx_porter_instance_global_name_region,priority:2"`
	FeatureSummary          *modelsupervisor.PorterFeatureSummary `gorm:"serializer:json"`
	ContextJSONSchema       string
	Status                  string
	ConnectionStatus        string
	ConnectionStatusMessage string
	UpdatedAt               time.Time
	CreatedAt               time.Time
}

func (PorterInstance) TableName() string {
	return "porter_instances"
}
