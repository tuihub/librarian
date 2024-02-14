package bizangela

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modelangela"
	"github.com/tuihub/librarian/internal/lib/libmq"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"
)

func NewUpdateAppIndexTopic(
	a *AngelaBase,
) *libmq.Topic[modelangela.UpdateAppIndex] {
	return libmq.NewTopic[modelangela.UpdateAppIndex](
		"UpdateAppIndex",
		func(ctx context.Context, r *modelangela.UpdateAppIndex) error {
			apps, err := a.g.GetBatchBoundApps(ctx, r.IDs)
			if err != nil {
				return err
			}
			for _, app := range apps {
				desc := app.Internal.Name
				for _, other := range app.Others {
					desc += other.Name
				}
				err = a.searcher.DescribeID(ctx,
					app.Internal.ID,
					desc,
					searcher.DescribeIDRequest_DESCRIBE_MODE_OVERRIDE,
					searcher.Index_INDEX_GEBURA_APP,
				)
				if err != nil {
					return err
				}
			}
			return nil
		},
	)
}
