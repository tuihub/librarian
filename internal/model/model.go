package model

import "time"

type Paging struct {
	PageSize int
	PageNum  int
}

func (p *Paging) ToLimit() int {
	return p.PageSize
}

func (p *Paging) ToOffset() int {
	return (p.PageNum - 1) * p.PageSize
}

type InternalID int64

type TimeRange struct {
	StartTime time.Time
	Duration  time.Duration
}
