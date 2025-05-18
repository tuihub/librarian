package bizgebura

import (
	"context"
	"fmt"
	"time"

	"github.com/tuihub/librarian/internal/biz/bizutils"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelgebura"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
)

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
		sentinelID,
		0,
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
		ExpireAt:        time.Now().Add(sentinelRefreshTokenExpire),
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
	ctx context.Context, page *model.Paging, sentinelID model.InternalID,
) ([]*modelgebura.SentinelSession, int64, error) {
	if libauth.FromContextAssertUserType(ctx, model.UserTypeAdmin) == nil {
		return nil, 0, bizutils.NoPermissionError()
	}
	sessions, total, err := g.repo.ListSentinelSessions(ctx, page, sentinelID)
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
	if len(s.Libraries) > 0 {
		ids, err := g.id.BatchNew(len(s.Libraries))
		if err != nil {
			return pb.ErrorErrorReasonUnspecified("%s", err.Error())
		}
		for i := range s.Libraries {
			s.Libraries[i].ID = ids[i]
		}
	}
	err := g.repo.UpdateSentinelInfo(ctx, s)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (g *Gebura) UpsertAppBinaries(
	ctx context.Context, abs []*modelgebura.SentinelAppBinary, snapshot *time.Time, commit bool,
) (bool, error) {
	claims := libauth.FromContextAssertUserType(ctx, model.UserTypeSentinel)
	if claims == nil {
		return false, bizutils.NoPermissionError()
	}
	sentinelID := claims.UserID
	ids, err := g.id.BatchNew(len(abs))
	if err != nil {
		return false, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	for i := range abs {
		abs[i].ID = ids[i]
		abs[i].UnionID = fmt.Sprintf("%d-%d-%s", sentinelID, abs[i].SentinelLibraryID, abs[i].GeneratedID)
		if len(abs[i].Files) > 0 {
			fileIDs, err2 := g.id.BatchNew(len(abs[i].Files))
			if err2 != nil {
				return false, pb.ErrorErrorReasonUnspecified("%s", err2.Error())
			}
			for j := range abs[i].Files {
				abs[i].Files[j].ID = fileIDs[j]
			}
		}
	}
	err = g.repo.UpsertAppBinaries(ctx, sentinelID, abs, snapshot, commit)
	if err != nil {
		return false, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return commit, nil
}
