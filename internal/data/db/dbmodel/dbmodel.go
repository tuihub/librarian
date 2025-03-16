package dbmodel

import (
	"gorm.io/gorm"
	"time"
)

type InternalID int64

type Model struct {
	ID        InternalID `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	deletedAt gorm.DeletedAt `gorm:"index"`
}

type ID InternalID
