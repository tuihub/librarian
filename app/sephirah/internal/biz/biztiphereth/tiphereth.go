package biztiphereth

import (
	"context"
	"time"

	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/logger"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

// TipherethRepo is a User repo.
type TipherethRepo interface {
	UserActive(context.Context, *User) (bool, error)
	FetchUserByPassword(context.Context, *User) (*User, error)
	AddUser(context.Context, *User) error
	UpdateUser(context.Context, *User) error
	ListUser(context.Context, Paging, []libauth.UserType, []UserStatus) ([]*User, error)
	CreateAccount(context.Context, Account) error
}

// TipherethUseCase is a User use case.
type TipherethUseCase struct {
	auth        *libauth.Auth
	repo        TipherethRepo
	mapper      mapper.LibrarianMapperServiceClient
	searcher    searcher.LibrarianSearcherServiceClient
	porter      porter.LibrarianPorterServiceClient
	pullAccount *libmq.TopicImpl[PullAccountInfo]
}

// NewTipherethUseCase new a User use case.
func NewTipherethUseCase(
	repo TipherethRepo,
	auth *libauth.Auth,
	mClient mapper.LibrarianMapperServiceClient,
	pClient porter.LibrarianPorterServiceClient,
	sClient searcher.LibrarianSearcherServiceClient,
	pullAccount *libmq.TopicImpl[PullAccountInfo],
) (*TipherethUseCase, error) {
	return &TipherethUseCase{
		auth:        auth,
		repo:        repo,
		mapper:      mClient,
		porter:      pClient,
		searcher:    sClient,
		pullAccount: pullAccount,
	}, nil
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

	user, err = t.repo.FetchUserByPassword(ctx, user)
	if err != nil {
		logger.Errorf("FetchUserByPassword failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	var accessToken, refreshToken string
	accessToken, err = t.auth.GenerateToken(user.InternalID,
		libauth.ClaimsTypeAccessToken, user.Type, nil, time.Hour)
	if err != nil {
		logger.Infof("generate access token failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	refreshToken, err = t.auth.GenerateToken(user.InternalID,
		libauth.ClaimsTypeRefreshToken, user.Type, nil, time.Hour*24*7) //nolint:gomnd //TODO
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
	accessToken, err := t.auth.GenerateToken(claims.InternalID,
		libauth.ClaimsTypeAccessToken, claims.UserType, nil, time.Hour)
	if err != nil {
		logger.Infof("generate access token failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	refreshToken, err = t.auth.GenerateToken(claims.InternalID,
		libauth.ClaimsTypeRefreshToken, claims.UserType, nil, time.Hour*24*7) //nolint:gomnd //TODO
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
	user.PassWord = password
	resp, err := t.searcher.NewID(ctx, &searcher.NewIDRequest{})
	if err != nil {
		logger.Infof("NewID failed: %s", err.Error())
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	user.InternalID = resp.Id
	user.Status = UserStatusActive
	err = t.repo.AddUser(ctx, user)
	if err != nil {
		logger.Infof("repo AddUser failed: %s", err.Error())
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return &User{
		InternalID: resp.Id,
	}, nil
}

func (t *TipherethUseCase) UpdateUser(ctx context.Context, user *User) *errors.Error {
	if user.PassWord != "" {
		password, err := t.auth.GeneratePassword(user.PassWord)
		if err != nil {
			logger.Infof("generate password failed: %s", err.Error())
			return pb.ErrorErrorReasonBadRequest("invalid password")
		}
		user.PassWord = password
	}
	err := t.repo.UpdateUser(ctx, user)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (t *TipherethUseCase) ListUser(
	ctx context.Context,
	paging Paging,
	types []libauth.UserType,
	statuses []UserStatus,
) ([]*User, *errors.Error) {
	users, err := t.repo.ListUser(ctx, paging, types, statuses)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return users, nil
}

func (t *TipherethUseCase) LinkAccount(ctx context.Context, a Account) (*Account, *errors.Error) {
	claims, exist := libauth.FromContext(ctx)
	if !exist {
		return nil, pb.ErrorErrorReasonUnauthorized("invalid token")
	}
	if resp, err := t.searcher.NewID(ctx, &searcher.NewIDRequest{}); err != nil {
		logger.Infof("NewID failed: %s", err.Error())
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	} else {
		a.InternalID = resp.Id
	}
	if err := t.repo.CreateAccount(ctx, a); err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	if _, err := t.mapper.InsertEdge(ctx, &mapper.InsertEdgeRequest{EdgeList: []*mapper.Edge{
		{
			SrcVid: claims.InternalID,
			DstVid: a.InternalID,
			Type:   mapper.EdgeType_EDGE_TYPE_EQUAL,
			Prop:   nil,
		},
	}}); err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	if err := t.pullAccount.Publish(PullAccountInfo{
		Platform:          a.Platform,
		PlatformAccountID: a.PlatformAccountID,
	}); err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return &a, nil
}
