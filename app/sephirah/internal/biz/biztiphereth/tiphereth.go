package biztiphereth

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizutils"
	"github.com/tuihub/librarian/app/sephirah/internal/client"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/supervisor"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/logger"
	"github.com/tuihub/librarian/model"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

type TipherethRepo interface {
	FetchUserByPassword(context.Context, *modeltiphereth.User) (*modeltiphereth.User, error)
	CreateUser(context.Context, *modeltiphereth.User, model.InternalID) error
	UpdateUser(context.Context, *modeltiphereth.User, string) error
	ListUsers(context.Context, model.Paging, []model.InternalID,
		[]libauth.UserType, []modeltiphereth.UserStatus, []model.InternalID,
		model.InternalID) ([]*modeltiphereth.User, int64, error)
	LinkAccount(context.Context, modeltiphereth.Account, model.InternalID) error
	UnLinkAccount(context.Context, modeltiphereth.Account, model.InternalID) error
	ListLinkAccounts(context.Context, model.InternalID) ([]*modeltiphereth.Account, error)
	GetUser(context.Context, model.InternalID) (*modeltiphereth.User, error)
}

type Tiphereth struct {
	auth        *libauth.Auth
	repo        TipherethRepo
	supv        *supervisor.Supervisor
	mapper      mapper.LibrarianMapperServiceClient
	searcher    *client.Searcher
	pullAccount *libmq.Topic[modeltiphereth.PullAccountInfo]
}

func NewTiphereth(
	repo TipherethRepo,
	auth *libauth.Auth,
	supv *supervisor.Supervisor,
	mClient mapper.LibrarianMapperServiceClient,
	sClient *client.Searcher,
	pullAccount *libmq.Topic[modeltiphereth.PullAccountInfo],
) (*Tiphereth, error) {
	return &Tiphereth{
		auth:        auth,
		repo:        repo,
		supv:        supv,
		mapper:      mClient,
		searcher:    sClient,
		pullAccount: pullAccount,
	}, nil
}

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
	accessToken, err = t.auth.GenerateToken(user.ID,
		libauth.ClaimsTypeAccessToken, user.Type, nil, libtime.Hour)
	if err != nil {
		logger.Infof("generate access token failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonUnspecified("generate access token failed: %s", err.Error())
	}
	refreshToken, err = t.auth.GenerateToken(user.ID,
		libauth.ClaimsTypeRefreshToken, user.Type, nil, libtime.SevenDays)
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
	accessToken, err := t.auth.GenerateToken(claims.InternalID,
		libauth.ClaimsTypeAccessToken, claims.UserType, nil, libtime.Hour)
	if err != nil {
		logger.Infof("generate access token failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	refreshToken, err = t.auth.GenerateToken(claims.InternalID,
		libauth.ClaimsTypeRefreshToken, claims.UserType, nil, libtime.SevenDays)
	if err != nil {
		logger.Infof("generate refresh token failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return modeltiphereth.AccessToken(accessToken), modeltiphereth.RefreshToken(refreshToken), nil
}

func (t *Tiphereth) CreateDefaultAdmin(ctx context.Context, user *modeltiphereth.User) {
	password, err := t.auth.GeneratePassword(user.PassWord)
	if err != nil {
		logger.Infof("generate password failed: %s", err.Error())
		return
	}
	user.PassWord = password
	id, err := t.searcher.NewID(ctx)
	if err != nil {
		return
	}
	user.ID = id
	user.Status = modeltiphereth.UserStatusActive
	user.Type = libauth.UserTypeAdmin
	if _, err = t.mapper.InsertVertex(ctx, &mapper.InsertVertexRequest{VertexList: []*mapper.Vertex{
		{
			Vid:  int64(user.ID),
			Type: mapper.VertexType_VERTEX_TYPE_ABSTRACT,
			Prop: nil,
		},
	}}); err != nil {
		return
	}
	if err = t.repo.CreateUser(ctx, user, user.ID); err != nil {
		logger.Infof("repo CreateUser failed: %s", err.Error())
		return
	}
}

func (t *Tiphereth) CreateUser(ctx context.Context, user *modeltiphereth.User) (*model.InternalID, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, bizutils.NoPermissionError()
	}
	if claims.UserType != libauth.UserTypeAdmin && user.Type != libauth.UserTypeSentinel {
		return nil, bizutils.NoPermissionError()
	}
	password, err := t.auth.GeneratePassword(user.PassWord)
	if err != nil {
		logger.Infof("generate password failed: %s", err.Error())
		return nil, pb.ErrorErrorReasonBadRequest("invalid password")
	}
	user.PassWord = password
	id, err := t.searcher.NewID(ctx)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err)
	}
	user.ID = id
	if _, err = t.mapper.InsertVertex(ctx, &mapper.InsertVertexRequest{VertexList: []*mapper.Vertex{
		{
			Vid:  int64(user.ID),
			Type: mapper.VertexType_VERTEX_TYPE_ABSTRACT,
			Prop: nil,
		},
	}}); err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	if err = t.repo.CreateUser(ctx, user, claims.InternalID); err != nil {
		logger.Infof("repo CreateUser failed: %s", err.Error())
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	res := user.ID
	return &res, nil
}

func (t *Tiphereth) UpdateUser(
	ctx context.Context, user *modeltiphereth.User, originPassword string,
) *errors.Error {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return bizutils.NoPermissionError()
	}
	if user.ID == 0 {
		return pb.ErrorErrorReasonBadRequest("internal id required")
	}
	if user.PassWord != "" && originPassword == "" {
		return pb.ErrorErrorReasonBadRequest("password required")
	}
	if claims.UserType != libauth.UserTypeAdmin &&
		claims.InternalID != user.ID {
		res, _, err := t.repo.ListUsers(ctx,
			model.Paging{
				PageSize: 1,
				PageNum:  1,
			},
			[]model.InternalID{user.ID},
			[]libauth.UserType{libauth.UserTypeSentinel},
			nil,
			nil,
			claims.InternalID,
		)
		if err != nil || len(res) != 1 || res[0].ID != user.ID {
			return bizutils.NoPermissionError()
		}
	}
	if user.PassWord != "" {
		password, err := t.auth.GeneratePassword(user.PassWord)
		if err != nil {
			logger.Infof("generate password failed: %s", err.Error())
			return pb.ErrorErrorReasonBadRequest("invalid password")
		}
		user.PassWord = password
		originPassword, err = t.auth.GeneratePassword(originPassword)
		if err != nil {
			logger.Infof("generate password failed: %s", err.Error())
			return pb.ErrorErrorReasonBadRequest("invalid password")
		}
	}
	err := t.repo.UpdateUser(ctx, user, originPassword)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (t *Tiphereth) ListUsers(
	ctx context.Context, paging model.Paging, types []libauth.UserType, statuses []modeltiphereth.UserStatus,
) ([]*modeltiphereth.User, int64, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, 0, bizutils.NoPermissionError()
	}
	var exclude []model.InternalID
	users, total, err := t.repo.ListUsers(ctx, paging, nil, types, statuses, exclude, claims.InternalID)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return users, total, nil
}

func (t *Tiphereth) GetUser(ctx context.Context, id *model.InternalID) (*modeltiphereth.User, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, bizutils.NoPermissionError()
	}
	var userID model.InternalID
	if id != nil {
		userID = *id
	} else {
		userID = claims.InternalID
	}
	user, err := t.repo.GetUser(ctx, userID)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return user, nil
}

func (t *Tiphereth) LinkAccount(
	ctx context.Context,
	a modeltiphereth.Account,
) (*modeltiphereth.Account, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, bizutils.NoPermissionError()
	}
	if !t.supv.CheckAccountPlatform(a.Platform) {
		return nil, bizutils.UnsupportedFeatureError()
	}
	id, err := t.searcher.NewID(ctx)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err)
	}
	a.ID = id
	if err = t.repo.LinkAccount(ctx, a, claims.InternalID); err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	if err = t.pullAccount.Publish(ctx, modeltiphereth.PullAccountInfo{
		ID:                a.ID,
		Platform:          a.Platform,
		PlatformAccountID: a.PlatformAccountID,
	}); err != nil {
		logger.Errorf("Publish PullAccountInfo failed %s", err.Error())
	}
	return &a, nil
}

func (t *Tiphereth) UnLinkAccount(ctx context.Context, a modeltiphereth.Account) *errors.Error {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return bizutils.NoPermissionError()
	}
	if !t.supv.CheckAccountPlatform(a.Platform) {
		return bizutils.UnsupportedFeatureError()
	}
	if err := t.repo.UnLinkAccount(ctx, a, claims.InternalID); err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (t *Tiphereth) ListLinkAccounts(
	ctx context.Context, id model.InternalID,
) ([]*modeltiphereth.Account, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, bizutils.NoPermissionError()
	}
	if id == 0 {
		id = claims.InternalID
	}
	a, err := t.repo.ListLinkAccounts(ctx, id)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return a, nil
}
