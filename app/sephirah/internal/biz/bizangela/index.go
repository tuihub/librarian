package bizangela

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modelangela"
	"github.com/tuihub/librarian/internal/lib/libmq"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"
)

func NewUpdateAppInfoIndexTopic(
	a *AngelaBase,
) *libmq.Topic[modelangela.UpdateAppInfoIndex] {
	return libmq.NewTopic[modelangela.UpdateAppInfoIndex](
		"UpdateAppInfoIndex",
		func(ctx context.Context, r *modelangela.UpdateAppInfoIndex) error {
			infos, err := a.g.GetBatchBoundAppInfos(ctx, r.IDs)
			if err != nil {
				return err
			}
			for _, info := range infos {
				desc := info.Internal.Name
				for _, other := range info.Others {
					desc += other.Name
				}
				err = a.searcher.DescribeID(ctx,
					info.Internal.ID,
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
