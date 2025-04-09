package bizgebura

import (
	"context"

	"github.com/tuihub/librarian/internal/biz/bizutils"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelgebura"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
)

func (g *Gebura) UpdateSentinelInfo(
	ctx context.Context, info *modelgebura.SentinelInfo,
) error {
	claims := libauth.FromContextAssertUserType(ctx, model.UserTypeSentinel)
	if claims == nil {
		return bizutils.NoPermissionError()
	}
	info.ID = claims.UserID
	var err error
	for _, lib := range info.Libraries {
		lib.ID, err = g.id.New()
		if err != nil {
			return pb.ErrorErrorReasonUnspecified("%s", err.Error())
		}
	}
	err = g.repo.UpsertSentinelInfo(ctx, info)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}
