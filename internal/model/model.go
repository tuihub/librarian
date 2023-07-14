package model

import (
	"time"

	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

type Paging struct {
	PageSize int32
	PageNum  int32
}

func (p *Paging) ToLimit() int {
	return int(p.PageSize)
}

func (p *Paging) ToOffset() int {
	return int((p.PageNum - 1) * p.PageSize)
}

func ToBizPaging(paging *librarian.PagingRequest) Paging {
	const defaultPageSize = 10
	const defaultPageNum = 1
	if paging == nil {
		return Paging{
			PageSize: defaultPageSize,
			PageNum:  defaultPageNum,
		}
	}
	return Paging{
		PageSize: paging.PageSize,
		PageNum:  paging.PageNum,
	}
}

func ToPBPaging(paging Paging) *librarian.PagingRequest {
	return &librarian.PagingRequest{
		PageNum:  paging.PageNum,
		PageSize: paging.PageSize,
	}
}

type InternalID int64

type TimeRange struct {
	StartTime time.Time
	Duration  time.Duration
}
