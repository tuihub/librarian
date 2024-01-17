package biztiphereth

import (
	"context"
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizutils"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

const accessTokenExpire = time.Hour
const refreshTokenExpire = libtime.SevenDays
const refreshTokenNeedRefresh = libtime.FiveDays

func (t *Tiphereth) GetToken(
	ctx context.Context,
	username, password string,
	deviceID *model.InternalID,
) (modeltiphereth.AccessToken, modeltiphereth.RefreshToken, *errors.Error) {
	password, err := t.auth.GeneratePassword(password)
	if err != nil {
		logger.Infof("generate password failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonUnauthorized("invalid user or password")
	}

	user, err := t.repo.FetchUserByPassword(ctx, username, password)
	if err != nil {
		logger.Infof("FetchUserByPassword failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonUnauthorized("invalid user or password")
	}
	if user.Status != modeltiphereth.UserStatusActive {
		return "", "", pb.ErrorErrorReasonUnauthorized("user not active")
	}

	sessionID, err := t.searcher.NewID(ctx)
	if err != nil {
		return "", "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	userSession := &modeltiphereth.UserSession{
		ID:           sessionID,
		UserID:       user.ID,
		RefreshToken: "",
		DeviceInfo:   nil,
		CreateAt:     time.Now(),
		ExpireAt:     time.Now().Add(refreshTokenExpire),
	}
	if deviceID != nil {
		userSession.DeviceInfo, err = t.repo.FetchDeviceInfo(ctx, *deviceID)
		if err != nil {
			logger.Infof("FetchDeviceInfo failed: %s", err.Error())
			return "", "", pb.ErrorErrorReasonUnauthorized("invalid device")
		}
	}

	var accessToken, refreshToken string
	accessToken, err = t.auth.GenerateToken(
		user.ID,
		0,
		libauth.ClaimsTypeAccessToken,
		user.Type,
		nil,
		accessTokenExpire,
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
		refreshTokenExpire,
	)
	if err != nil {
		logger.Infof("generate refresh token failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonUnspecified("generate access token failed: %s", err.Error())
	}
	userSession.RefreshToken = refreshToken
	err = t.repo.CreateUserSession(ctx, userSession)
	if err != nil {
		logger.Infof("create user session failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return modeltiphereth.AccessToken(accessToken), modeltiphereth.RefreshToken(refreshToken), nil
}

func (t *Tiphereth) RefreshToken( //nolint:gocognit // TODO
	ctx context.Context,
	deviceID *model.InternalID,
) (modeltiphereth.AccessToken, modeltiphereth.RefreshToken, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx,
		libauth.UserTypeAdmin, libauth.UserTypeNormal, libauth.UserTypeSentinel, libauth.UserTypePorter)
	if claims == nil {
		return "", "", bizutils.NoPermissionError()
	}
	oldRefreshToken := libauth.RawFromContext(ctx)
	if oldRefreshToken == "" {
		return "", "", bizutils.NoPermissionError()
	}
	needUpdate := false
	session := new(modeltiphereth.UserSession)
	var err error
	if claims.UserType != libauth.UserTypePorter { //nolint:nestif // TODO
		session, err = t.repo.FetchUserSession(ctx, claims.UserID, oldRefreshToken)
		if err != nil {
			return "", "", bizutils.NoPermissionError()
		}
		if session.RefreshToken != oldRefreshToken {
			return "", "", bizutils.NoPermissionError()
		}
		if session.DeviceInfo == nil && deviceID != nil {
			session.DeviceInfo, err = t.repo.FetchDeviceInfo(ctx, *deviceID)
			if err != nil {
				logger.Infof("FetchDeviceInfo failed: %s", err.Error())
				return "", "", pb.ErrorErrorReasonUnauthorized("invalid device")
			}
			needUpdate = true
		}
	}
	var accessToken, refreshToken string
	accessToken, err = t.auth.GenerateToken(
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
	if claims.ExpiresAt.After(time.Now().Add(refreshTokenNeedRefresh)) {
		refreshToken = oldRefreshToken
	} else {
		refreshToken, err = t.auth.GenerateToken(
			claims.UserID,
			claims.PorterID,
			libauth.ClaimsTypeRefreshToken,
			claims.UserType,
			nil,
			refreshTokenExpire,
		)
		if err != nil {
			logger.Infof("generate refresh token failed: %s", err.Error())
			return "", "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
		}
		session.RefreshToken = refreshToken
		session.CreateAt = time.Now()
		session.ExpireAt = time.Now().Add(refreshTokenExpire)
		needUpdate = true
	}
	if claims.UserType != libauth.UserTypePorter && needUpdate {
		err = t.repo.UpdateUserSession(ctx, session)
		if err != nil {
			logger.Infof("update user session failed: %s", err.Error())
			return "", "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
		}
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

func (t *Tiphereth) RegisterDevice(
	ctx context.Context,
	device *modeltiphereth.DeviceInfo,
) (model.InternalID, *errors.Error) {
	if libauth.FromContextAssertUserType(ctx) == nil {
		return 0, bizutils.NoPermissionError()
	}
	id, err := t.searcher.NewID(ctx)
	if err != nil {
		return 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	device.ID = id
	err = t.repo.CreateDevice(ctx, device)
	if err != nil {
		return 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return id, nil
}

func (t *Tiphereth) ListUserSessions(ctx context.Context) ([]*modeltiphereth.UserSession, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, bizutils.NoPermissionError()
	}
	sessions, err := t.repo.ListUserSessions(ctx, claims.UserID)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return sessions, nil
}

func (t *Tiphereth) DeleteUserSession(ctx context.Context, id model.InternalID) *errors.Error {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return bizutils.NoPermissionError()
	}
	err := t.repo.DeleteUserSession(ctx, claims.UserID, id)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}
