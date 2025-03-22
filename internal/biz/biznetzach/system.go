package biznetzach

import (
	"context"

	"github.com/tuihub/librarian/internal/data"
	"github.com/tuihub/librarian/internal/lib/libidgenerator"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/model/modelnetzach"
)

func NewSystemNotificationTopic(
	repo *data.NetzachRepo,
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
