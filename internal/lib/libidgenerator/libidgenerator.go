package libidgenerator

import (
	"github.com/tuihub/librarian/internal/model"

	"github.com/google/wire"
	"github.com/sony/sonyflake/v2"
)

var ProviderSet = wire.NewSet(NewIDGenerator)

type IDGenerator struct {
	sf *sonyflake.Sonyflake
}

func NewIDGenerator() (*IDGenerator, error) {
	sf, err := sonyflake.New(sonyflake.Settings{ //nolint:exhaustruct // no need
		MachineID: func() (int, error) { // TODO
			return 0, nil
		},
	})
	if err != nil {
		return nil, err
	}
	return &IDGenerator{
		sf: sf,
	}, nil
}

func (i *IDGenerator) New() (model.InternalID, error) {
	id, err := i.sf.NextID()
	return model.InternalID(id), err
}

func (i *IDGenerator) BatchNew(n int) ([]model.InternalID, error) {
	ids := make([]model.InternalID, n)
	for idx := range n {
		id, err := i.New()
		if err != nil {
			return nil, err
		}
		ids[idx] = id
	}
	return ids, nil
}
