package bizkether

import (
	"context"

	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/model/modelkether"
)

func NewUpdateAppInfoIndexTopic(
	a *KetherBase,
) *libmq.Topic[modelkether.UpdateAppInfoIndex] {
	return libmq.NewTopic[modelkether.UpdateAppInfoIndex](
		"UpdateAppInfoIndex",
		func(ctx context.Context, r *modelkether.UpdateAppInfoIndex) error {
			// infos, err := a.g.GetBatchBoundAppInfos(ctx, r.IDs)
			// if err != nil {
			//	return err
			//}
			// for _, info := range infos {
			//	desc := info.Internal.Name
			//	for _, other := range info.Others {
			//		desc += other.Name
			//	}
			//	err = a.search.DescribeID(ctx,
			//		info.Internal.ID,
			//		libsearch.SearchIndexGeburaApp,
			//		false,
			//		desc,
			//	)
			//	if err != nil {
			//		return err
			//	}
			//}
			return nil
		},
	)
}
