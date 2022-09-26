package biz

import (
	"context"
	"time"

	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/logger"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

// User is a User model.
type User struct {
	UniqueID int64
	UserName string
	PassWord string
}

type AccessToken string
type RefreshToken string

// TipherethRepo is a Greater repo.
type TipherethRepo interface {
	UserActive(context.Context, *User) (bool, error)
	GetUserID(context.Context, *User) (*User, error)
	AddUser(context.Context, *User) (*User, error)
}

// TipherethUseCase is a User use case.
type TipherethUseCase struct {
	auth     *libauth.Auth
	repo     TipherethRepo
	searcher searcher.LibrarianSearcherServiceClient
}

// NewTipherethUseCase new a User use case.
func NewTipherethUseCase(repo TipherethRepo, auth *libauth.Auth, mClient mapper.LibrarianMapperServiceClient,
	sClient searcher.LibrarianSearcherServiceClient, pClient porter.LibrarianPorterServiceClient) *TipherethUseCase {
	return &TipherethUseCase{auth: auth, repo: repo, searcher: sClient}
}

func (t *TipherethUseCase) GetToken(ctx context.Context, user *User) (AccessToken, RefreshToken, *errors.Error) {
	password, err := t.auth.GeneratePassword(user.PassWord)
	if err != nil {
		logger.Infof("generate password failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonBadRequest("invalid password")
	}
	user.PassWord = password
	ok, err := t.repo.UserActive(ctx, user)
	if err != nil {
		logger.Infof("UserActive failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	if !ok {
		return "", "", pb.ErrorErrorReasonBadRequest("invalid password")
	}

	user, err = t.repo.GetUserID(ctx, user)
	if err != nil {
		logger.Errorf("GetUserID failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	var accessToken, refreshToken string
	accessToken, err = t.auth.GenerateToken(user.UniqueID,
		libauth.ClaimsTypeAccessToken, time.Hour)
	if err != nil {
		logger.Infof("generate access token failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	refreshToken, err = t.auth.GenerateToken(user.UniqueID,
		libauth.ClaimsTypeRefreshToken, time.Hour*24*7) //nolint:gomnd //TODO
	if err != nil {
		logger.Infof("generate refresh token failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return AccessToken(accessToken), RefreshToken(refreshToken), nil
}

func (t *TipherethUseCase) RefreshToken(ctx context.Context) (AccessToken, RefreshToken, *errors.Error) {
	claims, exist := libauth.FromContext(ctx)
	if !exist {
		return "", "", pb.ErrorErrorReasonUnauthorized("empty token")
	}
	var accessToken, refreshToken string
	accessToken, err := t.auth.GenerateToken(claims.ID,
		libauth.ClaimsTypeAccessToken, time.Hour)
	if err != nil {
		logger.Infof("generate access token failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	refreshToken, err = t.auth.GenerateToken(claims.ID,
		libauth.ClaimsTypeRefreshToken, time.Hour*24*7) //nolint:gomnd //TODO
	if err != nil {
		logger.Infof("generate refresh token failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return AccessToken(accessToken), RefreshToken(refreshToken), nil
}

func (t *TipherethUseCase) AddUser(ctx context.Context, user *User) (*User, *errors.Error) {
	password, err := t.auth.GeneratePassword(user.PassWord)
	if err != nil {
		logger.Infof("generate password failed: %s", err.Error())
		return nil, pb.ErrorErrorReasonBadRequest("invalid password")
	}
	resp, err := t.searcher.NewID(ctx, &searcher.NewIDRequest{})
	if err != nil {
		logger.Infof("NewID failed: %s", err.Error())
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	_, err = t.repo.AddUser(ctx, &User{
		UniqueID: resp.Id,
		UserName: user.UserName,
		PassWord: password,
	})
	if err != nil {
		logger.Infof("repo AddUser failed: %s", err.Error())
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return &User{
		UniqueID: resp.Id,
	}, nil
}
