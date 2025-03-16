package biztiphereth

import (
	"bytes"
	"context"
	"strconv"

	"github.com/tuihub/librarian/internal/biz/bizutils"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/dchest/captcha"
	"github.com/go-kratos/kratos/v2/errors"
)

func (t *Tiphereth) CreateUser(ctx context.Context, user *model.User) (*model.InternalID, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, bizutils.NoPermissionError()
	}
	if claims.UserType != model.UserTypeAdmin && user.Type != model.UserTypeSentinel {
		return nil, bizutils.NoPermissionError()
	}
	if t.app.EnvExist(libapp.EnvDemoMode) {
		if user.Type == model.UserTypeAdmin {
			return nil, pb.ErrorErrorReasonForbidden("server running in demo mode, create admin user is not allowed")
		}
	}
	password, err := t.auth.GeneratePassword(user.Password)
	if err != nil {
		logger.Infof("generate password failed: %s", err.Error())
		return nil, pb.ErrorErrorReasonBadRequest("invalid password")
	}
	user.Password = password
	id, err := t.id.New()
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err)
	}
	user.ID = id
	// if _, err = t.mapper.InsertVertex(ctx, &mapper.InsertVertexRequest{VertexList: []*mapper.Vertex{
	//	{
	//		Vid:  int64(user.ID),
	//		Type: mapper.VertexType_VERTEX_TYPE_ABSTRACT,
	//		Prop: nil,
	//	},
	// }}); err != nil {
	//	return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	//}
	if err = t.repo.CreateUser(ctx, user, claims.UserID); err != nil {
		logger.Infof("repo CreateUser failed: %s", err.Error())
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	_ = t.userCountCache.Delete(ctx)
	res := user.ID
	return &res, nil
}

func (t *Tiphereth) UpdateUser(
	ctx context.Context, user *model.User, originPassword string,
) *errors.Error {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return bizutils.NoPermissionError()
	}
	if user.ID == 0 {
		return pb.ErrorErrorReasonBadRequest("internal id required")
	}
	if user.Password != "" && originPassword == "" {
		return pb.ErrorErrorReasonBadRequest("password required")
	}
	if t.app.EnvExist(libapp.EnvDemoMode) {
		if user.Type == model.UserTypeAdmin {
			return pb.ErrorErrorReasonForbidden("server running in demo mode, modify admin user is not allowed")
		}
	}
	if claims.UserType != model.UserTypeAdmin &&
		claims.UserID != user.ID {
		res, _, err := t.repo.ListUsers(ctx,
			model.Paging{
				PageSize: 1,
				PageNum:  1,
			},
			[]model.InternalID{user.ID},
			[]model.UserType{model.UserTypeSentinel},
			nil,
			nil,
			claims.UserID,
		)
		if err != nil || len(res) != 1 || res[0].ID != user.ID {
			return bizutils.NoPermissionError()
		}
	}
	if user.Password != "" {
		password, err := t.auth.GeneratePassword(user.Password)
		if err != nil {
			logger.Infof("generate password failed: %s", err.Error())
			return pb.ErrorErrorReasonBadRequest("invalid password")
		}
		user.Password = password
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
	ctx context.Context, paging model.Paging, types []model.UserType, statuses []model.UserStatus,
) ([]*model.User, int64, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, 0, bizutils.NoPermissionError()
	}
	var exclude []model.InternalID
	users, total, err := t.repo.ListUsers(ctx, paging, nil, types, statuses, exclude, claims.UserID)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return users, total, nil
}

func (t *Tiphereth) GetUser(ctx context.Context, id *model.InternalID) (*model.User, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, bizutils.NoPermissionError()
	}
	var userID model.InternalID
	if id != nil {
		userID = *id
	} else {
		userID = claims.UserID
	}
	user, err := t.repo.GetUser(ctx, userID)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return user, nil
}

func (t *Tiphereth) RegisterUser(
	ctx context.Context,
	username string,
	password string,
	captchaReq *model.CaptchaAns,
) (*model.CaptchaQue, string, *errors.Error) {
	if t.app.EnvExist(libapp.EnvDemoMode) {
		return nil, "", pb.ErrorErrorReasonForbidden("server running in demo mode, register user is not allowed")
	}
	if !t.app.EnvExist(libapp.EnvAllowRegister) {
		return nil, "", pb.ErrorErrorReasonForbidden("server not allow register user")
	}
	if err := t.checkCapacity(ctx); err != nil {
		return nil, "", err
	}

	if captchaReq == nil {
		captchaID := captcha.New()
		captchaImg := bytes.NewBuffer(nil)
		err := captcha.WriteImage(captchaImg, captchaID, 200, 100) //nolint:mnd // hard code
		if err != nil {
			return nil, "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
		}
		return &model.CaptchaQue{
			ID:    captchaID,
			Image: captchaImg.Bytes(),
		}, "", nil
	}
	if !captcha.VerifyString(captchaReq.ID, captchaReq.Value) {
		return nil, "", pb.ErrorErrorReasonBadRequest("invalid captcha")
	}
	passwordParsed, err := t.auth.GeneratePassword(password)
	if err != nil {
		logger.Infof("generate password failed: %s", err.Error())
		return nil, "", pb.ErrorErrorReasonBadRequest("invalid password")
	}
	id, err := t.id.New()
	if err != nil {
		return nil, "", pb.ErrorErrorReasonUnspecified("%s", err)
	}
	user := &model.User{
		ID:       id,
		Username: username,
		Password: passwordParsed,
		Type:     model.UserTypeNormal,
		Status:   model.UserStatusActive,
	}
	if err = t.repo.CreateUser(ctx, user, user.ID); err != nil {
		logger.Infof("repo CreateUser failed: %s", err.Error())
		return nil, "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	_ = t.userCountCache.Delete(ctx)
	return nil, "TODO", nil // TODO: return refresh token
}

func (t *Tiphereth) checkCapacity(ctx context.Context) *errors.Error {
	if !t.app.EnvExist(libapp.EnvUserCapacity) {
		return nil
	}
	capacityStr, err := t.app.Env(libapp.EnvUserCapacity)
	if err != nil {
		return pb.ErrorErrorReasonForbidden("server have invalid setting")
	}
	capacity, err := strconv.Atoi(capacityStr)
	if err != nil {
		return pb.ErrorErrorReasonForbidden("server have invalid setting")
	}
	count, err := t.userCountCache.Get(ctx)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err)
	}
	if count.Count >= capacity {
		return pb.ErrorErrorReasonForbidden("server user capacity reached")
	}
	return nil
}
