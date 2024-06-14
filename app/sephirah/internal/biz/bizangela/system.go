package bizangela

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modelangela"
	"github.com/tuihub/librarian/internal/lib/libmq"
)

func NewSystemNotificationTopic(
	a *AngelaBase,
) *libmq.Topic[modelangela.SystemNotify] {
	return libmq.NewTopic[modelangela.SystemNotify](
		"SystemNotify",
		func(ctx context.Context, r *modelangela.SystemNotify) error {
			if r.Notification.ID == 0 {
				id, err := a.searcher.NewID(ctx)
				if err != nil {
					return err
				}
				r.Notification.ID = id
			}
			return a.repo.UpsertSystemNotification(ctx, r.UserID, &r.Notification)
		},
	)
}
