package data

import (
	"context"
	"time"

	"github.com/tuihub/librarian/app/searcher/internal/biz"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"

	"github.com/blevesearch/bleve/v2"
	"github.com/google/wire"
	"github.com/meilisearch/meilisearch-go"
	"github.com/sony/sonyflake"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewSearcherRepo, NewSnowFlake, NewBleve, NewMeili)

func NewSearcherRepo(
	b map[biz.Index]bleve.Index, m meilisearch.ServiceManager, sf *sonyflake.Sonyflake,
) (biz.SearcherRepo, error) {
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
	logger.Warnf("no valid search backend, search function will not work")
	return &defaultSearcherRepo{sf: sf}, nil
}

type defaultSearcherRepo struct {
	sf *sonyflake.Sonyflake
}

func (d defaultSearcherRepo) DescribeID(
	context.Context, model.InternalID, biz.Index, bool, string) error {
	return nil // search disabled
}

func (d defaultSearcherRepo) SearchID(
	context.Context, biz.Index, model.Paging, string) ([]*biz.SearchResult, error) {
	return nil, nil // search disabled
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

func (d defaultSearcherRepo) NewID(ctx context.Context) (int64, error) {
	id, err := d.sf.NextID()
	return int64(id), err //nolint:gosec // safe
}

func (r *bleveSearcherRepo) NewID(ctx context.Context) (int64, error) {
	id, err := r.sf.NextID()
	return int64(id), err //nolint:gosec // safe
}

func (m *meiliSearcherRepo) NewID(ctx context.Context) (int64, error) {
	id, err := m.sf.NextID()
	return int64(id), err //nolint:gosec // safe
}
