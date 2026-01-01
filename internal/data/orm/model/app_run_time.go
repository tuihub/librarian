package model

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type AppRunTime struct {
	ID        model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	UserID    model.InternalID `gorm:"index:idx_app_run_time_user_id_app_id,priority:1"`
	AppID     model.InternalID `gorm:"index:idx_app_run_time_user_id_app_id,priority:2"`
	DeviceID  model.InternalID
	StartTime time.Time     `gorm:"uniqueIndex:idx_app_run_time_start_time_duration,priority:1"`
	Duration  time.Duration `gorm:"uniqueIndex:idx_app_run_time_start_time_duration,priority:2"`
	UpdatedAt time.Time
	CreatedAt time.Time
	App       *App `gorm:"foreignKey:AppID"`
}

func (AppRunTime) TableName() string {
	return "app_run_times"
}
