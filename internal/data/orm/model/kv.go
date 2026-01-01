package model

import (
	"time"
)

type KV struct {
	Bucket    string `gorm:"primaryKey"`
	Key       string `gorm:"primaryKey"`
	Value     string
	UpdatedAt time.Time
	CreatedAt time.Time
}

func (KV) TableName() string {
	return "kvs"
}
