package data

import (
	"context"
	"errors"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/account"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/deviceinfo"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/portercontext"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/porterinstance"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/userdevice"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/usersession"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"

	"entgo.io/ent/dialect/sql"
)

type tipherethRepo struct {
	data *Data
}

// NewTipherethRepo .
func NewTipherethRepo(data *Data) biztiphereth.TipherethRepo {
	return &tipherethRepo{
		data: data,
	}
}

func (t tipherethRepo) FetchUserByPassword(
	ctx context.Context,
	username, password string,
) (*modeltiphereth.User, error) {
	u, err := t.data.db.User.Query().Where(
		user.UsernameEQ(username),
		user.PasswordEQ(password),
	).First(ctx)
	if u == nil || err != nil {
		return nil, errors.New("invalid user")
	}
	return converter.ToBizUser(u), nil
}

func (t tipherethRepo) CreateDevice(
	ctx context.Context,
	info *modeltiphereth.DeviceInfo,
) error {
	return t.data.db.DeviceInfo.Create().
		SetID(info.ID).
		SetDeviceName(info.DeviceName).
		SetSystemType(converter.ToEntSystemType(info.SystemType)).
		SetSystemVersion(info.SystemVersion).
		SetClientName(info.ClientName).
		SetClientSourceCodeAddress(info.ClientSourceCodeAddress).
		SetClientVersion(info.ClientVersion).
		Exec(ctx)
}

func (t tipherethRepo) FetchDeviceInfo(
	ctx context.Context,
	deviceID model.InternalID,
) (*modeltiphereth.DeviceInfo, error) {
	res, err := t.data.db.DeviceInfo.Get(ctx, deviceID)
	if err != nil {
		return nil, err
	}
	return converter.ToBizDeviceInfo(res), nil
}

func (t tipherethRepo) ListDevices(ctx context.Context, id model.InternalID) ([]*modeltiphereth.DeviceInfo, error) {
	devices, err := t.data.db.DeviceInfo.Query().Where(
		deviceinfo.HasUserWith(user.IDEQ(id)),
	).All(ctx)
	if err != nil {
		return nil, err
	}
	return converter.ToBizDeviceInfoList(devices), nil
}

func (t tipherethRepo) CreateUserSession(ctx context.Context, session *modeltiphereth.UserSession) error {
	return t.data.WithTx(ctx, func(tx *ent.Tx) error {
		q := tx.UserSession.Create().
			SetID(session.ID).
			SetUserID(session.UserID).
			SetRefreshToken(session.RefreshToken).
			SetCreatedAt(session.CreateAt).
			SetExpireAt(session.ExpireAt)
		if session.DeviceInfo != nil {
			_, _ = tx.UserSession.Delete().Where(
				usersession.UserIDEQ(session.UserID),
				usersession.HasDeviceInfoWith(
					deviceinfo.IDEQ(session.DeviceInfo.ID),
				)).Exec(ctx)
			q.SetDeviceInfoID(session.DeviceInfo.ID)
		}
		err := q.Exec(ctx)
		if err != nil {
			return err
		}
		if session.DeviceInfo != nil {
			err = tx.UserDevice.Create().
				SetUserID(session.UserID).
				SetDeviceInfoID(session.DeviceInfo.ID).
				OnConflict(
					sql.ConflictColumns(userdevice.FieldUserID, userdevice.FieldDeviceID),
				).
				UpdateNewValues().
				Exec(ctx)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (t tipherethRepo) FetchUserSession(
	ctx context.Context,
	userID model.InternalID,
	token string,
) (*modeltiphereth.UserSession, error) {
	session, err := t.data.db.UserSession.Query().Where(
		usersession.UserIDEQ(userID),
		usersession.RefreshTokenEQ(token),
	).WithDeviceInfo().Only(ctx)
	if err != nil {
		return nil, err
	}
	res := converter.ToBizUserSession(session)
	if session.Edges.DeviceInfo != nil {
		res.DeviceInfo = converter.ToBizDeviceInfo(session.Edges.DeviceInfo)
	}
	return res, nil
}

func (t tipherethRepo) ListUserSessions(
	ctx context.Context,
	id model.InternalID,
) ([]*modeltiphereth.UserSession, error) {
	session, err := t.data.db.UserSession.Query().Where(
		usersession.UserIDEQ(id),
	).WithDeviceInfo().All(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*modeltiphereth.UserSession, len(session))
	for i, s := range session {
		res[i] = converter.ToBizUserSession(s)
		if s.Edges.DeviceInfo != nil {
			res[i].DeviceInfo = converter.ToBizDeviceInfo(s.Edges.DeviceInfo)
		}
	}
	return res, nil
}

func (t tipherethRepo) UpdateUserSession(ctx context.Context, session *modeltiphereth.UserSession) error {
	return t.data.WithTx(ctx, func(tx *ent.Tx) error {
		q := tx.UserSession.UpdateOneID(session.ID).
			SetRefreshToken(session.RefreshToken).
			SetCreatedAt(session.CreateAt).
			SetExpireAt(session.ExpireAt)
		if session.DeviceInfo != nil {
			q.SetDeviceInfoID(session.DeviceInfo.ID)

			err := tx.User.UpdateOneID(session.UserID).
				AddDeviceInfoIDs(session.DeviceInfo.ID).
				Exec(ctx)
			if err != nil {
				return err
			}
		}
		return q.Exec(ctx)
	})
}

func (t tipherethRepo) DeleteUserSession(
	ctx context.Context,
	userID model.InternalID,
	sessionID model.InternalID,
) error {
	return t.data.db.UserSession.DeleteOneID(sessionID).Where(
		usersession.UserIDEQ(userID),
	).Exec(ctx)
}

func (t tipherethRepo) CreateUser(ctx context.Context, u *modeltiphereth.User, c model.InternalID) error {
	q := t.data.db.User.Create().
		SetID(u.ID).
		SetUsername(u.UserName).
		SetPassword(u.PassWord).
		SetStatus(converter.ToEntUserStatus(u.Status)).
		SetType(converter.ToEntUserType(u.Type)).
		SetCreatorID(c)
	return q.Exec(ctx)
}

func (t tipherethRepo) UpdateUser(ctx context.Context, u *modeltiphereth.User, password string) error {
	q := t.data.db.User.Update().
		Where(user.IDEQ(u.ID))
	if u.UserName != "" {
		q.SetUsername(u.UserName)
	}
	if u.PassWord != "" {
		q.Where(user.PasswordEQ(password)).SetPassword(u.PassWord)
	}
	if u.Type != libauth.UserTypeUnspecified {
		q.SetType(converter.ToEntUserType(u.Type))
	}
	if u.Status != modeltiphereth.UserStatusUnspecified {
		q.SetStatus(converter.ToEntUserStatus(u.Status))
	}
	return q.Exec(ctx)
}

func (t tipherethRepo) ListUsers(
	ctx context.Context,
	paging model.Paging,
	ids []model.InternalID,
	types []libauth.UserType,
	statuses []modeltiphereth.UserStatus,
	exclude []model.InternalID,
	creator model.InternalID,
) ([]*modeltiphereth.User, int64, error) {
	q := t.data.db.User.Query().Where(user.IDEQ(creator)).QueryCreatedUser()
	if len(ids) > 0 {
		q.Where(user.IDIn(ids...))
	}
	if len(types) > 0 {
		q.Where(user.TypeIn(converter.ToEntUserTypeList(types)...))
	}
	if len(statuses) > 0 {
		q.Where(user.StatusIn(converter.ToEntUserStatusList(statuses)...))
	}
	if len(exclude) > 0 {
		q.Where(user.IDNotIn(exclude...))
	}
	u, err := q.
		Limit(paging.ToLimit()).
		Offset(paging.ToOffset()).
		All(ctx)
	if err != nil {
		return nil, 0, err
	}
	count, err := q.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	return converter.ToBizUserList(u), int64(count), nil
}

func (t tipherethRepo) GetUser(ctx context.Context, id model.InternalID) (*modeltiphereth.User, error) {
	u, err := t.data.db.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return converter.ToBizUser(u), nil
}

func (t tipherethRepo) GetUserCount(ctx context.Context) (int, error) {
	return t.data.db.User.Query().Count(ctx)
}

func (t tipherethRepo) LinkAccount(
	ctx context.Context,
	a modeltiphereth.Account,
	userID model.InternalID,
) (model.InternalID, error) {
	accountID := a.ID
	err := t.data.WithTx(ctx, func(tx *ent.Tx) error {
		u, err := tx.User.Get(ctx, userID)
		if err != nil {
			return err
		}
		exist, err := u.QueryBindAccount().Where(
			account.PlatformEQ(a.Platform),
		).Exist(ctx)
		if err != nil {
			return err
		}
		if exist {
			return errors.New("an account already bound to user")
		}
		acc, err := tx.Account.Query().Where(
			account.PlatformEQ(a.Platform),
			account.PlatformAccountIDEQ(a.PlatformAccountID),
		).Only(ctx)
		if ent.IsNotFound(err) {
			return tx.Account.Create().
				SetBindUserID(userID).
				SetID(a.ID).
				SetPlatform(a.Platform).
				SetPlatformAccountID(a.PlatformAccountID).
				SetName(a.Name).
				SetAvatarURL(a.AvatarURL).
				SetProfileURL(a.ProfileURL).
				Exec(ctx)
		}
		if err != nil {
			return err
		}
		exist, err = acc.QueryBindUser().Exist(ctx)
		if err != nil {
			return err
		}
		if exist {
			return errors.New("account already bound to an user")
		}
		accountID = acc.ID
		return tx.Account.UpdateOneID(acc.ID).
			SetBindUserID(userID).
			Exec(ctx)
	})
	if err != nil {
		return 0, err
	}
	return accountID, nil
}

func (t tipherethRepo) UnLinkAccount(ctx context.Context, a modeltiphereth.Account, u model.InternalID) error {
	return t.data.db.Account.Update().Where(
		account.PlatformEQ(a.Platform),
		account.PlatformAccountIDEQ(a.PlatformAccountID),
		account.HasBindUserWith(user.IDEQ(u)),
	).
		ClearBindUser().
		Exec(ctx)
}

func (t tipherethRepo) ListLinkAccounts(
	ctx context.Context,
	userID model.InternalID,
) ([]*modeltiphereth.Account, error) {
	a, err := t.data.db.Account.Query().
		Where(
			account.HasBindUserWith(user.IDEQ(userID)),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return converter.ToBizAccountList(a), nil
}

func (t tipherethRepo) UpsertPorters(ctx context.Context, il []*modeltiphereth.PorterInstance) error {
	instances := make([]*ent.PorterInstanceCreate, len(il))
	for i, instance := range il {
		instances[i] = t.data.db.PorterInstance.Create().
			SetID(instance.ID).
			SetName(instance.Name).
			SetVersion(instance.Version).
			SetGlobalName(instance.GlobalName).
			SetAddress(instance.Address).
			SetStatus(converter.ToEntPorterInstanceStatus(instance.Status)).
			SetFeatureSummary(instance.FeatureSummary)
	}
	return t.data.db.PorterInstance.
		CreateBulk(instances...).
		OnConflict(
			sql.ConflictColumns(porterinstance.FieldAddress),
			resolveWithIgnores([]string{
				porterinstance.FieldID,
				porterinstance.FieldStatus,
				porterinstance.FieldStatus,
			}),
		).
		Exec(ctx)
}

func (t tipherethRepo) ListPorters(
	ctx context.Context,
	paging model.Paging,
) ([]*modeltiphereth.PorterInstance, int64, error) {
	q := t.data.db.PorterInstance.Query()
	count, err := q.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	p, err := q.
		Limit(paging.ToLimit()).
		Offset(paging.ToOffset()).
		All(ctx)
	if err != nil {
		return nil, 0, err
	}
	return converter.ToBizPorterList(p), int64(count), nil
}

func (t tipherethRepo) UpdatePorterStatus(
	ctx context.Context,
	id model.InternalID,
	status modeltiphereth.PorterInstanceStatus,
) (*modeltiphereth.PorterInstance, error) {
	pi, err := t.data.db.PorterInstance.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	err = pi.Update().SetStatus(converter.ToEntPorterInstanceStatus(status)).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return converter.ToBizPorter(pi), nil
}

func (t tipherethRepo) FetchPorterByAddress(ctx context.Context, address string) (*modeltiphereth.PorterInstance, error) {
	p, err := t.data.db.PorterInstance.Query().Where(
		porterinstance.AddressEQ(address),
	).Only(ctx)
	if err != nil {
		return nil, err
	}
	return converter.ToBizPorter(p), nil
}

func (t tipherethRepo) CreatePorterContext(
	ctx context.Context,
	userID model.InternalID,
	context *modeltiphereth.PorterContext,
) error {
	return t.data.db.PorterContext.Create().
		SetOwnerID(userID).
		SetGlobalName(context.GlobalName).
		SetRegion(context.Region).
		SetContextJSON(context.ContextJSON).
		SetName(context.Name).
		SetDescription(context.Description).
		SetStatus(converter.ToEntPorterContextStatus(context.Status)).
		Exec(ctx)
}

func (t tipherethRepo) ListPorterContexts(
	ctx context.Context,
	userID model.InternalID,
	paging model.Paging,
) ([]*modeltiphereth.PorterContext, int64, error) {
	q := t.data.db.PorterContext.Query().Where(
		portercontext.HasOwnerWith(user.IDEQ(userID)),
	)
	count, err := q.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	p, err := q.
		Limit(paging.ToLimit()).
		Offset(paging.ToOffset()).
		All(ctx)
	if err != nil {
		return nil, 0, err
	}
	return converter.ToBizPorterContextList(p), int64(count), nil
}

func (t tipherethRepo) UpdatePorterContext(
	ctx context.Context,
	userID model.InternalID,
	context *modeltiphereth.PorterContext,
) error {
	return t.data.db.PorterContext.Update().Where(
		portercontext.IDEQ(context.ID),
		portercontext.HasOwnerWith(user.IDEQ(userID)),
	).
		SetContextJSON(context.ContextJSON).
		SetName(context.Name).
		SetDescription(context.Description).
		SetStatus(converter.ToEntPorterContextStatus(context.Status)).
		Exec(ctx)
}
