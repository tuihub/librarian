package bizgebura

import (
	"context"
	"time"

	"github.com/tuihub/librarian/internal/biz/bizutils"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelgebura"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
)

func (g *Gebura) CreateAppCategory(ctx context.Context, ac *modelgebura.AppCategory) (*modelgebura.AppCategory, error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, bizutils.NoPermissionError()
	}
	id, err := g.id.New()
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err)
	}
	ac.ID = id
	ac.VersionNumber = 1
	ac.VersionDate = time.Now()
	if err = g.repo.CreateAppCategory(ctx, claims.UserID, ac); err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return ac, nil
}

func (g *Gebura) ListAppCategories(ctx context.Context) ([]*modelgebura.AppCategory, error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, bizutils.NoPermissionError()
	}
	acList, err := g.repo.ListAppCategories(ctx, claims.UserID)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return acList, nil
}

func (g *Gebura) UpdateAppCategory(ctx context.Context, ac *modelgebura.AppCategory) error {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return bizutils.NoPermissionError()
	}
	err := g.repo.UpdateAppCategory(ctx, claims.UserID, ac)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (g *Gebura) DeleteAppCategory(ctx context.Context, id model.InternalID) error {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return bizutils.NoPermissionError()
	}
	err := g.repo.DeleteAppCategory(ctx, claims.UserID, id)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}
