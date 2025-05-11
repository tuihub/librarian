package bizgebura

import (
	"context"
	"time"

	"github.com/tuihub/librarian/internal/biz/bizutils"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelgebura"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
)

const sentinelAccessTokenExpire = libtime.ThreeDays
const sentinelRefreshTokenExpire = libtime.ThirtyDays

func (g *Gebura) CreateSentinel(
	ctx context.Context, s *modelgebura.Sentinel,
) error {
	claims := libauth.FromContextAssertUserType(ctx, model.UserTypeAdmin)
	if claims == nil {
		return bizutils.NoPermissionError()
	}
	id, err := g.id.New()
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	s.ID = id
	err = g.repo.CreateSentinel(ctx, claims.UserID, s)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (g *Gebura) GetSentinel(ctx context.Context, id model.InternalID) (*modelgebura.Sentinel, error) {
	if libauth.FromContextAssertUserType(ctx, model.UserTypeAdmin) == nil {
		return nil, bizutils.NoPermissionError()
	}
	sentinel, err := g.repo.GetSentinel(ctx, id)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return sentinel, nil
}

func (g *Gebura) ListSentinels(
	ctx context.Context, page *model.Paging,
) ([]*modelgebura.Sentinel, int64, error) {
	if libauth.FromContextAssertUserType(ctx, model.UserTypeAdmin) == nil {
		return nil, 0, bizutils.NoPermissionError()
	}
	sentinels, total, err := g.repo.ListSentinels(ctx, page)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return sentinels, int64(total), nil
}

func (g *Gebura) UpdateSentinel(
	ctx context.Context, s *modelgebura.Sentinel,
) error {
	if libauth.FromContextAssertUserType(ctx, model.UserTypeAdmin) == nil {
		return bizutils.NoPermissionError()
	}
	err := g.repo.UpdateSentinel(ctx, s)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (g *Gebura) CreateSentinelSession(
	ctx context.Context, sentinelID model.InternalID,
) error {
	claims := libauth.FromContextAssertUserType(ctx, model.UserTypeAdmin)
	if claims == nil {
		return bizutils.NoPermissionError()
	}
	id, err := g.id.New()
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	refreshToken, err := g.auth.GenerateToken(
		claims.UserID,
		claims.PorterID,
		libauth.ClaimsTypeRefreshToken,
		model.UserTypeSentinel,
		nil,
		sentinelRefreshTokenExpire,
	)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	session := &modelgebura.SentinelSession{
		ID:              id,
		SentinelID:      sentinelID,
		RefreshToken:    refreshToken,
		Status:          modelgebura.SentinelSessionStatusActive,
		CreatorID:       claims.UserID,
		ExpireAt:        time.Now().Add(sentinelAccessTokenExpire),
		LastUsedAt:      nil,
		LastRefreshedAt: nil,
		RefreshCount:    0,
	}
	err = g.repo.CreateSentinelSession(ctx, session)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (g *Gebura) ListSentinelSessions(
	ctx context.Context, page *model.Paging,
) ([]*modelgebura.SentinelSession, int64, error) {
	if libauth.FromContextAssertUserType(ctx, model.UserTypeAdmin) == nil {
		return nil, 0, bizutils.NoPermissionError()
	}
	sessions, total, err := g.repo.ListSentinelSessions(ctx, page)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return sessions, int64(total), nil
}

func (g *Gebura) UpdateSentinelSessionStatus(
	ctx context.Context, id model.InternalID, status modelgebura.SentinelSessionStatus,
) error {
	if libauth.FromContextAssertUserType(ctx, model.UserTypeAdmin) == nil {
		return bizutils.NoPermissionError()
	}
	err := g.repo.UpdateSentinelSession(ctx, id, status)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (g *Gebura) DeleteSentinelSession(
	ctx context.Context, id model.InternalID,
) error {
	if libauth.FromContextAssertUserType(ctx, model.UserTypeAdmin) == nil {
		return bizutils.NoPermissionError()
	}
	err := g.repo.DeleteSentinelSession(ctx, id)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (g *Gebura) UpsertSentinelInfo(
	ctx context.Context, s *modelgebura.Sentinel,
) error {
	claims := libauth.FromContextAssertUserType(ctx, model.UserTypeSentinel)
	if claims == nil {
		return bizutils.NoPermissionError()
	}
	s.ID = claims.UserID
	err := g.repo.UpdateSentinelInfo(ctx, s)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (g *Gebura) UpsertAppBinaries(
	ctx context.Context, abs []*modelgebura.SentinelAppBinary,
) error {
	claims := libauth.FromContextAssertUserType(ctx, model.UserTypeSentinel)
	if claims == nil {
		return bizutils.NoPermissionError()
	}
	sentinelID := claims.UserID
	err := g.repo.UpsertAppBinaries(ctx, sentinelID, abs)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}
