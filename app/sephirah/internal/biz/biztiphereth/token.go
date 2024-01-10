package biztiphereth

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizutils"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

func (t *Tiphereth) GetToken(
	ctx context.Context,
	user *modeltiphereth.User,
) (modeltiphereth.AccessToken, modeltiphereth.RefreshToken, *errors.Error) {
	password, err := t.auth.GeneratePassword(user.PassWord)
	if err != nil {
		logger.Infof("generate password failed: %s", err.Error())

		return "", "", pb.ErrorErrorReasonUnauthorized("invalid user or password")
	}
	user.PassWord = password

	user, err = t.repo.FetchUserByPassword(ctx, user)
	if err != nil {
		logger.Infof("FetchUserByPassword failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonUnauthorized("invalid user or password")
	}
	if user.Status != modeltiphereth.UserStatusActive {
		return "", "", pb.ErrorErrorReasonUnauthorized("user not active")
	}

	var accessToken, refreshToken string
	accessToken, err = t.auth.GenerateToken(
		user.ID,
		0,
		libauth.ClaimsTypeAccessToken,
		user.Type,
		nil,
		libtime.Hour,
	)
	if err != nil {
		logger.Infof("generate access token failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonUnspecified("generate access token failed: %s", err.Error())
	}
	refreshToken, err = t.auth.GenerateToken(
		user.ID,
		0,
		libauth.ClaimsTypeRefreshToken,
		user.Type,
		nil,
		libtime.SevenDays,
	)
	if err != nil {
		logger.Infof("generate refresh token failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonUnspecified("generate access token failed: %s", err.Error())
	}
	// TODO save refreshToken to sql
	return modeltiphereth.AccessToken(accessToken), modeltiphereth.RefreshToken(refreshToken), nil
}

func (t *Tiphereth) RefreshToken(
	ctx context.Context,
) (modeltiphereth.AccessToken, modeltiphereth.RefreshToken, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx,
		libauth.UserTypeAdmin, libauth.UserTypeNormal, libauth.UserTypeSentinel, libauth.UserTypePorter)
	if claims == nil {
		return "", "", bizutils.NoPermissionError()
	}
	var accessToken, refreshToken string
	accessToken, err := t.auth.GenerateToken(
		claims.UserID,
		claims.PorterID,
		libauth.ClaimsTypeAccessToken,
		claims.UserType,
		nil,
		libtime.Hour,
	)
	if err != nil {
		logger.Infof("generate access token failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	refreshToken, err = t.auth.GenerateToken(
		claims.UserID,
		claims.PorterID,
		libauth.ClaimsTypeRefreshToken,
		claims.UserType,
		nil,
		libtime.SevenDays,
	)
	if err != nil {
		logger.Infof("generate refresh token failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return modeltiphereth.AccessToken(accessToken), modeltiphereth.RefreshToken(refreshToken), nil
}

func (t *Tiphereth) GainUserPrivilege(
	ctx context.Context,
	userID model.InternalID,
) (modeltiphereth.AccessToken, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx, libauth.UserTypePorter)
	if claims == nil || claims.PorterID != 0 {
		return "", bizutils.NoPermissionError()
	}
	privilege, err := t.repo.FetchPorterPrivilege(ctx, claims.UserID, userID)
	if err != nil {
		return "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	if !privilege.All {
		return "", bizutils.NoPermissionError()
	}
	accessToken, err := t.auth.GenerateToken(
		userID,
		claims.UserID,
		libauth.ClaimsTypeAccessToken,
		libauth.UserTypeNormal,
		nil,
		libtime.Day,
	)
	if err != nil {
		logger.Infof("generate access token failed: %s", err.Error())
		return "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return modeltiphereth.AccessToken(accessToken), nil
}
