package model

import "time"

type Paging struct {
	PageSize int
	PageNum  int
}

type InternalID int64

type TimeRange struct {
	StartTime time.Time
	Duration  time.Duration
}
