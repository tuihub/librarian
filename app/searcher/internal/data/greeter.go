package data

import (
	"context"

	"github.com/sony/sonyflake"
	"github.com/tuihub/librarian/app/searcher/internal/biz"
)

type greeterRepo struct {
	data *Data
	sf   *sonyflake.Sonyflake
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data) biz.GreeterRepo {
	sf := sonyflake.NewSonyflake(sonyflake.Settings{
		MachineID: func() (uint16, error) {
			return 0, nil
		},
	})
	return &greeterRepo{
		data,
		sf,
	}
}

func (r *greeterRepo) NewID(ctx context.Context) (int64, error) {
	id, err := r.sf.NextID()
	return int64(id), err
}
