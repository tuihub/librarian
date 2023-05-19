package data

import (
	"context"
	"errors"
	"time"

	"github.com/tuihub/librarian/app/searcher/internal/biz"

	"github.com/blevesearch/bleve/v2"
	"github.com/google/wire"
	"github.com/meilisearch/meilisearch-go"
	"github.com/sony/sonyflake"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewSearcherRepo, NewSnowFlake, NewBleve, NewMeili)

func NewSearcherRepo(b bleve.Index, m *meilisearch.Client, sf *sonyflake.Sonyflake) (biz.SearcherRepo, error) {
	if m != nil {
		return &meiliSearcherRepo{
			sf:     sf,
			search: m,
		}, nil
	}
	if b != nil {
		return &bleveSearcherRepo{
			sf:     sf,
			search: b,
		}, nil
	}
	return nil, errors.New("no valid search backend")
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

func (m *meiliSearcherRepo) NewID(ctx context.Context) (int64, error) {
	id, err := m.sf.NextID()
	return int64(id), err
}
