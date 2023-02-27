package biztiphereth

import (
	"context"
	"time"

	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

type TipherethRepo interface {
	FetchUserByPassword(context.Context, *User) (*User, error)
	CreateUser(context.Context, *User, *model.InternalID) error
	UpdateUser(context.Context, *User) error
	ListUser(context.Context, model.Paging, []model.InternalID,
		[]libauth.UserType, []UserStatus, []model.InternalID, *model.InternalID) ([]*User, error)
	CreateAccount(context.Context, Account) error
	UpdateAccount(context.Context, Account) error
}

type Tiphereth struct {
	auth        *libauth.Auth
	repo        TipherethRepo
	mapper      mapper.LibrarianMapperServiceClient
	searcher    searcher.LibrarianSearcherServiceClient
	porter      porter.LibrarianPorterServiceClient
	pullAccount *libmq.TopicImpl[PullAccountInfo]
}

func NewTiphereth(
	repo TipherethRepo,
	auth *libauth.Auth,
	mClient mapper.LibrarianMapperServiceClient,
	pClient porter.LibrarianPorterServiceClient,
	sClient searcher.LibrarianSearcherServiceClient,
	pullAccount *libmq.TopicImpl[PullAccountInfo],
) (*Tiphereth, error) {
	return &Tiphereth{
		auth:        auth,
		repo:        repo,
		mapper:      mClient,
		porter:      pClient,
		searcher:    sClient,
		pullAccount: pullAccount,
	}, nil
}

func (t *Tiphereth) GetToken(ctx context.Context, user *User) (AccessToken, RefreshToken, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return "", "", pb.ErrorErrorReasonForbidden("no permission")
	}
	password, err := t.auth.GeneratePassword(user.PassWord)
	if err != nil {
		logger.Infof("generate password failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonBadRequest("invalid password")
	}
	user.PassWord = password

	user, err = t.repo.FetchUserByPassword(ctx, user)
	if err != nil {
		logger.Infof("FetchUserByPassword failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonBadRequest("invalid password")
	}
	if user.Status != UserStatusActive {
		return "", "", pb.ErrorErrorReasonBadRequest("user not active")
	}

	var accessToken, refreshToken string
	accessToken, err = t.auth.GenerateToken(user.InternalID,
		libauth.ClaimsTypeAccessToken, user.Type, nil, time.Hour)
	if err != nil {
		logger.Infof("generate access token failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonUnspecified("generate access token failed: %s", err.Error())
	}
	refreshToken, err = t.auth.GenerateToken(user.InternalID,
		libauth.ClaimsTypeRefreshToken, user.Type, nil, time.Hour*24*7) //nolint:gomnd //TODO
	if err != nil {
		logger.Infof("generate refresh token failed: %s", err.Error())
		return "", "", pb.ErrorErrorReasonUnspecified("generate access token failed: %s", err.Error())
	}
	// TODO save refreshToken to sql
	return AccessToken(accessToken), RefreshToken(refreshToken), nil
}

func (t *Tiphereth) RefreshToken(ctx context.Context) (AccessToken, RefreshToken, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal, libauth.UserTypeSentinel) {
		return "", "", pb.ErrorErrorReasonForbidden("no permission")
	}
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

func (t *Tiphereth) CreateUser(ctx context.Context, user *User) (*model.InternalID, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return nil, pb.ErrorErrorReasonForbidden("no permission")
	}
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin) {
		if user.Type != libauth.UserTypeSentinel {
			return nil, pb.ErrorErrorReasonForbidden("no permission")
		}
	}
	var creator model.InternalID
	if c, ok := libauth.FromContext(ctx); ok {
		creator = model.InternalID(c.InternalID)
	}
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
	if _, err = t.mapper.InsertVertex(ctx, &mapper.InsertVertexRequest{VertexList: []*mapper.Vertex{
		{
			Vid:  user.InternalID,
			Type: mapper.VertexType_VERTEX_TYPE_ABSTRACT,
			Prop: nil,
		},
	}}); err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	if err = t.repo.CreateUser(ctx, user, &creator); err != nil {
		logger.Infof("repo CreateUser failed: %s", err.Error())
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	res := model.InternalID(user.InternalID)
	return &res, nil
}

func (t *Tiphereth) UpdateUser(ctx context.Context, user *User) *errors.Error {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return pb.ErrorErrorReasonForbidden("no permission")
	}
	if user.InternalID == 0 {
		return pb.ErrorErrorReasonBadRequest("internal id required")
	}
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin) {
		c, ok := libauth.FromContext(ctx)
		if !ok {
			return pb.ErrorErrorReasonForbidden("no permission")
		}
		if c.InternalID != user.InternalID {
			res, err := t.repo.ListUser(ctx,
				model.Paging{
					PageSize: 1,
					PageNum:  1,
				},
				[]model.InternalID{model.InternalID(user.InternalID)},
				[]libauth.UserType{libauth.UserTypeSentinel},
				nil,
				nil,
				(*model.InternalID)(&c.InternalID),
			)
			if err != nil || len(res) != 1 || res[0].InternalID != user.InternalID {
				return pb.ErrorErrorReasonForbidden("no permission")
			}
		}
	}
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

func (t *Tiphereth) ListUser(
	ctx context.Context,
	paging model.Paging,
	types []libauth.UserType,
	statuses []UserStatus,
) ([]*User, *errors.Error) {
	var exclude []model.InternalID
	var creator *model.InternalID
	if c, ok := libauth.FromContext(ctx); !ok {
		return nil, pb.ErrorErrorReasonBadRequest("token required")
	} else {
		if c.UserType != libauth.UserTypeAdmin {
			creator = (*model.InternalID)(&c.InternalID)
		}
		exclude = append(exclude, model.InternalID(c.InternalID))
	}
	users, err := t.repo.ListUser(ctx, paging, nil, types, statuses, exclude, creator)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return users, nil
}

func (t *Tiphereth) LinkAccount(ctx context.Context, a Account) (*Account, *errors.Error) {
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
	if _, err := t.mapper.InsertVertex(ctx, &mapper.InsertVertexRequest{VertexList: []*mapper.Vertex{
		{
			Vid:  a.InternalID,
			Type: mapper.VertexType_VERTEX_TYPE_ENTITY,
			Prop: nil,
		},
	}}); err != nil {
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
	if err := t.repo.CreateAccount(ctx, a); err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	if err := t.pullAccount.Publish(ctx, PullAccountInfo{
		InternalID:        a.InternalID,
		Platform:          a.Platform,
		PlatformAccountID: a.PlatformAccountID,
	}); err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return &a, nil
}
