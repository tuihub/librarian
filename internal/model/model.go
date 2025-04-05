package model

import (
	"fmt"
	"time"

	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

type Paging struct {
	PageSize int64
	PageNum  int64
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
		PageSize: paging.GetPageSize(),
		PageNum:  paging.GetPageNum(),
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

type AccountAppRelationType int

const (
	AccountAppRelationTypeUnspecified AccountAppRelationType = iota
	AccountAppRelationTypeOwner
)

type ConfigDigest struct {
	Name    string
	Enabled *bool
	Driver  *string
	Listen  *string
}

func (d *ConfigDigest) Status() string {
	if d.Driver != nil {
		return fmt.Sprintf("Enable - Driver %s", *d.Driver)
	} else if d.Listen != nil {
		return fmt.Sprintf("Enable - Listen on %s", *d.Listen)
	} else if d.Enabled != nil {
		return "Enable"
	} else {
		return "Disable"
	}
}

func (d *ConfigDigest) String() string {
	return fmt.Sprintf("[%s\t] %s", d.Name, d.Status())
}
