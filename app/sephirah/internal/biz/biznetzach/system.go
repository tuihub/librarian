package biznetzach

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
	"github.com/tuihub/librarian/internal/lib/libidgenerator"
	"github.com/tuihub/librarian/internal/lib/libmq"
)

func NewSystemNotificationTopic(
	repo NetzachRepo,
	idGenerator *libidgenerator.IDGenerator,
) *libmq.Topic[modelnetzach.SystemNotify] {
	return libmq.NewTopic[modelnetzach.SystemNotify](
		"SystemNotify",
		func(ctx context.Context, r *modelnetzach.SystemNotify) error {
			if r.Notification.ID == 0 {
				id, err := idGenerator.New()
				if err != nil {
					return err
				}
				r.Notification.ID = id
			}
			return repo.UpsertSystemNotification(ctx, r.UserID, &r.Notification)
		},
	)
}
