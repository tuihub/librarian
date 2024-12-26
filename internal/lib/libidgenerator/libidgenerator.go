package libidgenerator

import (
	"time"

	"github.com/tuihub/librarian/internal/model"

	"github.com/google/wire"
	"github.com/sony/sonyflake"
)

var ProviderSet = wire.NewSet(NewIDGenerator)

type IDGenerator struct {
	sf *sonyflake.Sonyflake
}

func NewIDGenerator() *IDGenerator {
	return &IDGenerator{
		sf: sonyflake.NewSonyflake(sonyflake.Settings{
			StartTime: time.Time{},
			MachineID: func() (uint16, error) { // TODO
				return 0, nil
			},
			CheckMachineID: nil,
		}),
	}
}

func (i *IDGenerator) New() (model.InternalID, error) {
	id, err := i.sf.NextID()
	return model.InternalID(id), err //nolint:gosec // safe
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
