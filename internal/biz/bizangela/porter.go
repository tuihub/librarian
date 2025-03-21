package bizangela

import (
	"context"

	"github.com/tuihub/librarian/internal/biz/bizutils"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	"github.com/tuihub/librarian/internal/model/modelsupervisor"
)

func (a *Angela) PorterGetNotifyTargetItems(ctx context.Context, id model.InternalID, paging model.Paging) (*modelsupervisor.FeatureRequest, []*modelfeed.Item, error) {
	claims := libauth.FromContextAssertUserType(ctx, model.UserTypePorter)
	if claims == nil {
		return nil, nil, bizutils.NoPermissionError()
	}
	fr, items, err := a.repo.GetNotifyTargetItems(ctx, id, paging)
	if err != nil {
		return nil, nil, err
	}
	return fr, items, nil
}
