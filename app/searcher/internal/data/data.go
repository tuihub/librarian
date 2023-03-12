package data

import (
	"context"
	"time"

	"github.com/tuihub/librarian/app/searcher/internal/biz"

	"github.com/blevesearch/bleve/v2"
	"github.com/google/wire"
	"github.com/sony/sonyflake"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewSearcherRepo, NewSnowFlake, NewBleve)

func NewSearcherRepo(b bleve.Index, sf *sonyflake.Sonyflake) biz.SearcherRepo {
	return &bleveSearcherRepo{
		sf:     sf,
		search: b,
	}
}

func NewSnowFlake() *sonyflake.Sonyflake {
	return sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: time.Time{},
		MachineID: func() (uint16, error) { // TODO
			return 0, nil
		},
		CheckMachineID: nil,
	})
}

func (r *bleveSearcherRepo) NewID(ctx context.Context) (int64, error) {
	id, err := r.sf.NextID()
	return int64(id), err
}
