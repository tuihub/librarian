package bizgebura

import (
	"context"
	"fmt"
	"time"

	"github.com/tuihub/librarian/internal/biz/bizutils"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelgebura"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

const sentinelAccessTokenExpire = libtime.ThreeDays
const sentinelRefreshTokenExpire = libtime.ThirtyDays
const sentinelRefreshTokenNeedRefresh = libtime.TwentyTwoDays

func (g *Gebura) SentinelRefreshToken(
	ctx context.Context,
) (model.AccessToken, model.RefreshToken, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx, model.UserTypeSentinel)
	if claims == nil {
		return "", "", bizutils.NoPermissionError()
	}
	oldRefreshToken := libauth.RawFromContext(ctx)
	if oldRefreshToken == "" {
		return "", "", bizutils.NoPermissionError()
	}
	session, err := g.repo.GetSentinelSession(ctx, claims.UserID, oldRefreshToken)
	if err != nil || session.RefreshToken != oldRefreshToken {
		return "", "", bizutils.NoPermissionError()
	}
	var accessToken, refreshToken string
	accessToken, err = g.auth.GenerateToken(
		claims.UserID,
		claims.PorterID,
		libauth.ClaimsTypeAccessToken,
		claims.UserType,
		sentinelAccessTokenExpire,
		libauth.WithClaimsSentinelExtra(&libauth.ClaimsSentinelExtra{
			SentinelSessionID: session.ID,
		}),
	)
	if err != nil {
		logger.Infof("generate access token failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	expireAt := claims.ExpiresAt.Local()
	if claims.ExpiresAt.After(time.Now().Add(sentinelRefreshTokenNeedRefresh)) {
		refreshToken = oldRefreshToken
	} else {
		expireAt = time.Now()
		refreshToken, err = g.auth.GenerateToken(
			claims.UserID,
			claims.PorterID,
			libauth.ClaimsTypeRefreshToken,
			claims.UserType,
			sentinelRefreshTokenExpire,
			libauth.WithClaimsSentinelExtra(&libauth.ClaimsSentinelExtra{
				SentinelSessionID: session.ID,
			}),
		)
		if err != nil {
			logger.Infof("generate refresh token failed: %s", err.Error())
			return "", "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
		}
	}
	err = g.repo.UpdateSentinelSessionToken(ctx, session.ID, refreshToken, expireAt, time.Now())
	if err != nil {
		return "", "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return model.AccessToken(accessToken), model.RefreshToken(refreshToken), nil
}

func (g *Gebura) UpsertSentinelInfo(
	ctx context.Context, s *modelgebura.Sentinel,
) *errors.Error {
	claims := libauth.FromContextAssertUserType(ctx, model.UserTypeSentinel)
	if claims == nil {
		return bizutils.NoPermissionError()
	}
	_ = g.repo.UpdateSentinelSessionLastUsed(ctx, claims.SentinelSessionID, time.Now())
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
) (bool, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx, model.UserTypeSentinel)
	if claims == nil {
		return false, bizutils.NoPermissionError()
	}
	_ = g.repo.UpdateSentinelSessionLastUsed(ctx, claims.SentinelSessionID, time.Now())
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
