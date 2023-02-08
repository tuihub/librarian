package data

import (
	"context"
	"time"

	"github.com/tuihub/librarian/app/searcher/internal/biz"

	"github.com/sony/sonyflake"
)

type searcherRepo struct {
	data *Data
	sf   *sonyflake.Sonyflake
}

func NewSearcherRepo(data *Data) biz.SearcherRepo {
	sf := sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: time.Time{},
		MachineID: func() (uint16, error) { // TODO
			return 0, nil
		},
		CheckMachineID: nil,
	})
	return &searcherRepo{
		data,
		sf,
	}
}

func (r *searcherRepo) NewID(ctx context.Context) (int64, error) {
	id, err := r.sf.NextID()
	return int64(id), err
}
