package data

import (
	"context"
	"errors"

	"github.com/tuihub/librarian/internal/data/internal/converter"
	"github.com/tuihub/librarian/internal/data/internal/ent"
	"github.com/tuihub/librarian/internal/data/internal/ent/account"
	"github.com/tuihub/librarian/internal/data/internal/ent/device"
	"github.com/tuihub/librarian/internal/data/internal/ent/portercontext"
	"github.com/tuihub/librarian/internal/data/internal/ent/porterinstance"
	"github.com/tuihub/librarian/internal/data/internal/ent/session"
	"github.com/tuihub/librarian/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelsupervisor"

	"entgo.io/ent/dialect/sql"
)

type TipherethRepo struct {
	data *Data
}

func NewTipherethRepo(data *Data) *TipherethRepo {
	return &TipherethRepo{
		data: data,
	}
}

func (t *TipherethRepo) FetchUserByPassword(
	ctx context.Context,
	username, password string,
) (*model.User, error) {
	u, err := t.data.db.User.Query().Where(
		user.UsernameEQ(username),
		user.PasswordEQ(password),
	).First(ctx)
	if u == nil || err != nil {
		return nil, errors.New("invalid user")
	}
	return converter.ToBizUser(u), nil
}

func (t *TipherethRepo) CreateDevice(
	ctx context.Context,
	info *model.Device,
	clientLocalID *string,
) (model.InternalID, error) {
	var res model.InternalID
	err := t.data.WithTx(ctx, func(tx *ent.Tx) error {
		if clientLocalID != nil {
			infos, err := tx.Device.Query().Where(
				device.ClientLocalIDEQ(*clientLocalID),
			).All(ctx)
			if err != nil {
				return err
			}
			if len(infos) > 0 {
				res = infos[0].ID
				return nil
			}
		}
		q := tx.Device.Create().
			SetID(info.ID).
			SetDeviceName(info.DeviceName).
			SetSystemType(converter.ToEntSystemType(info.SystemType)).
			SetSystemVersion(info.SystemVersion).
			SetClientName(info.ClientName).
			SetClientSourceCodeAddress(info.ClientSourceCodeAddress).
			SetClientVersion(info.ClientVersion)
		if clientLocalID != nil {
			q.SetClientLocalID(*clientLocalID)
		}
		res = info.ID
		return q.Exec(ctx)
	})
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (t *TipherethRepo) FetchDeviceInfo(
	ctx context.Context,
	deviceID model.InternalID,
) (*model.Device, error) {
	res, err := t.data.db.Device.Get(ctx, deviceID)
	if err != nil {
		return nil, err
	}
	return converter.ToBizDeviceInfo(res), nil
}

func (t *TipherethRepo) CreateUserSession(ctx context.Context, s *model.Session) error {
	return t.data.WithTx(ctx, func(tx *ent.Tx) error {
		q := tx.Session.Create().
			SetID(s.ID).
			SetUserID(s.UserID).
			SetRefreshToken(s.RefreshToken).
			SetCreatedAt(s.CreateAt).
			SetExpireAt(s.ExpireAt)
		if s.Device != nil {
			_, _ = tx.Session.Delete().Where(
				session.UserIDEQ(s.UserID),
				session.HasDeviceWith(
					device.IDEQ(s.Device.ID),
				)).Exec(ctx)
			q.SetDeviceID(s.Device.ID)
		}
		err := q.Exec(ctx)
		if err != nil {
			return err
		}
		return nil
	})
}

func (t *TipherethRepo) FetchUserSession(
	ctx context.Context,
	userID model.InternalID,
	token string,
) (*model.Session, error) {
	s, err := t.data.db.Session.Query().Where(
		session.UserIDEQ(userID),
		session.RefreshTokenEQ(token),
	).WithDevice().Only(ctx)
	if err != nil {
		return nil, err
	}
	res := converter.ToBizUserSession(s)
	if s.Edges.Device != nil {
		res.Device = converter.ToBizDeviceInfo(s.Edges.Device)
	}
	return res, nil
}

func (t *TipherethRepo) ListUserSessions(
	ctx context.Context,
	id model.InternalID,
) ([]*model.Session, error) {
	ss, err := t.data.db.Session.Query().Where(
		session.UserIDEQ(id),
	).WithDevice().All(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*model.Session, len(ss))
	for i, s := range ss {
		res[i] = converter.ToBizUserSession(s)
		if s.Edges.Device != nil {
			res[i].Device = converter.ToBizDeviceInfo(s.Edges.Device)
		}
	}
	return res, nil
}

func (t *TipherethRepo) UpdateUserSession(ctx context.Context, session *model.Session) error {
	return t.data.WithTx(ctx, func(tx *ent.Tx) error {
		q := tx.Session.UpdateOneID(session.ID).
			SetRefreshToken(session.RefreshToken).
			SetCreatedAt(session.CreateAt).
			SetExpireAt(session.ExpireAt)
		if session.Device != nil {
			q.SetDeviceID(session.Device.ID)
		}
		return q.Exec(ctx)
	})
}

func (t *TipherethRepo) DeleteUserSession(
	ctx context.Context,
	userID model.InternalID,
	sessionID model.InternalID,
) error {
	return t.data.db.Session.DeleteOneID(sessionID).Where(
		session.UserIDEQ(userID),
	).Exec(ctx)
}

func (t *TipherethRepo) CreateUser(ctx context.Context, u *model.User, c model.InternalID) error {
	q := t.data.db.User.Create().
		SetID(u.ID).
		SetUsername(u.Username).
		SetPassword(u.Password).
		SetStatus(converter.ToEntUserStatus(u.Status)).
		SetType(converter.ToEntUserType(u.Type)).
		SetCreatorID(c)
	return q.Exec(ctx)
}

func (t *TipherethRepo) UpdateUser(ctx context.Context, u *model.User, password string) error {
	q := t.data.db.User.Update().
		Where(user.IDEQ(u.ID))
	if u.Username != "" {
		q.SetUsername(u.Username)
	}
	if u.Password != "" {
		q.Where(user.PasswordEQ(password)).SetPassword(u.Password)
	}
	if u.Type != model.UserTypeUnspecified {
		q.SetType(converter.ToEntUserType(u.Type))
	}
	if u.Status != model.UserStatusUnspecified {
		q.SetStatus(converter.ToEntUserStatus(u.Status))
	}
	return q.Exec(ctx)
}

func (t *TipherethRepo) ListUsers(
	ctx context.Context,
	paging model.Paging,
	ids []model.InternalID,
	types []model.UserType,
	statuses []model.UserStatus,
	exclude []model.InternalID,
) ([]*model.User, int64, error) {
	q := t.data.db.User.Query()
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
	count, err := q.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	u, err := q.
		Limit(paging.ToLimit()).
		Offset(paging.ToOffset()).
		All(ctx)
	if err != nil {
		return nil, 0, err
	}
	return converter.ToBizUserList(u), int64(count), nil
}

func (t *TipherethRepo) GetUser(ctx context.Context, id model.InternalID) (*model.User, error) {
	u, err := t.data.db.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return converter.ToBizUser(u), nil
}

func (t *TipherethRepo) GetUserCount(ctx context.Context) (int, error) {
	return t.data.db.User.Query().Count(ctx)
}

func (t *TipherethRepo) LinkAccount(
	ctx context.Context,
	a model.Account,
	userID model.InternalID,
) (model.InternalID, error) {
	accountID := a.ID
	err := t.data.WithTx(ctx, func(tx *ent.Tx) error {
		u, err := tx.User.Get(ctx, userID)
		if err != nil {
			return err
		}
		exist, err := u.QueryAccount().Where(
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
				SetBoundUserID(userID).
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
		exist, err = acc.QueryBoundUser().Exist(ctx)
		if err != nil {
			return err
		}
		if exist {
			return errors.New("account already bound to an user")
		}
		accountID = acc.ID
		return tx.Account.UpdateOneID(acc.ID).
			SetBoundUserID(userID).
			Exec(ctx)
	})
	if err != nil {
		return 0, err
	}
	return accountID, nil
}

func (t *TipherethRepo) UnLinkAccount(ctx context.Context, a model.Account, u model.InternalID) error {
	return t.data.db.Account.Update().Where(
		account.PlatformEQ(a.Platform),
		account.PlatformAccountIDEQ(a.PlatformAccountID),
		account.HasBoundUserWith(user.IDEQ(u)),
	).
		ClearBoundUser().
		Exec(ctx)
}

func (t *TipherethRepo) ListLinkAccounts(
	ctx context.Context,
	userID model.InternalID,
) ([]*model.Account, error) {
	a, err := t.data.db.Account.Query().
		Where(
			account.HasBoundUserWith(user.IDEQ(userID)),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return converter.ToBizAccountList(a), nil
}

func (t *TipherethRepo) UpsertPorters(ctx context.Context, il []*modelsupervisor.PorterInstance) error {
	instances := make([]*ent.PorterInstanceCreate, len(il))
	for i, instance := range il {
		if instance.BinarySummary == nil {
			instance.BinarySummary = new(modelsupervisor.PorterBinarySummary)
		}
		instances[i] = t.data.db.PorterInstance.Create().
			SetID(instance.ID).
			SetName(instance.BinarySummary.Name).
			SetVersion(instance.BinarySummary.Version).
			SetDescription(instance.BinarySummary.Description).
			SetSourceCodeAddress(instance.BinarySummary.SourceCodeAddress).
			SetBuildVersion(instance.BinarySummary.BuildVersion).
			SetBuildDate(instance.BinarySummary.BuildDate).
			SetGlobalName(instance.GlobalName).
			SetRegion(instance.Region).
			SetAddress(instance.Address).
			SetStatus(converter.ToEntPorterInstanceStatus(instance.Status)).
			SetFeatureSummary(instance.FeatureSummary).
			SetContextJSONSchema(instance.ContextJSONSchema)
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

func (t *TipherethRepo) ListPorters(
	ctx context.Context,
	paging model.Paging,
) ([]*modelsupervisor.PorterInstance, int64, error) {
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

func (t *TipherethRepo) UpdatePorterStatus(
	ctx context.Context,
	id model.InternalID,
	status model.UserStatus,
) (*modelsupervisor.PorterInstance, error) {
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

func (t *TipherethRepo) FetchPorterByAddress(
	ctx context.Context,
	address string,
) (*modelsupervisor.PorterInstance, error) {
	p, err := t.data.db.PorterInstance.Query().Where(
		porterinstance.AddressEQ(address),
	).Only(ctx)
	if err != nil {
		return nil, err
	}
	return converter.ToBizPorter(p), nil
}

func (t *TipherethRepo) CreatePorterContext(
	ctx context.Context,
	userID model.InternalID,
	context *modelsupervisor.PorterContext,
) error {
	return t.data.db.PorterContext.Create().
		SetID(context.ID).
		SetOwnerID(userID).
		SetGlobalName(context.GlobalName).
		SetRegion(context.Region).
		SetContextJSON(context.ContextJSON).
		SetName(context.Name).
		SetDescription(context.Description).
		SetStatus(converter.ToEntPorterContextStatus(context.Status)).
		Exec(ctx)
}

func (t *TipherethRepo) ListPorterContexts(
	ctx context.Context,
	userID model.InternalID,
	paging model.Paging,
) ([]*modelsupervisor.PorterContext, int64, error) {
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

func (t *TipherethRepo) UpdatePorterContext(
	ctx context.Context,
	userID model.InternalID,
	context *modelsupervisor.PorterContext,
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

func (t *TipherethRepo) ListPorterDigests(
	ctx context.Context,
	status []model.UserStatus,
) ([]*modelsupervisor.PorterDigest, error) {
	var res []struct {
		ent.PorterInstance
		Min model.InternalID
	}
	q := t.data.db.PorterInstance.Query()
	if len(status) > 0 {
		q.Where(porterinstance.StatusIn(converter.ToEntPorterInstanceStatusList(status)...))
	}
	err := q.GroupBy(
		porterinstance.FieldGlobalName,
		porterinstance.FieldRegion,
	).
		Aggregate(ent.Min(porterinstance.FieldID)).
		Scan(ctx, &res)
	if err != nil {
		return nil, err
	}
	var ids []model.InternalID
	for _, p := range res {
		ids = append(ids, p.Min)
	}
	pi, err := t.data.db.PorterInstance.Query().Where(
		porterinstance.IDIn(ids...),
	).All(ctx)
	if err != nil {
		return nil, err
	}
	var pg []*modelsupervisor.PorterDigest
	pgm := make(map[string]*modelsupervisor.PorterDigest)
	for _, p := range pi {
		if len(p.ContextJSONSchema) == 0 {
			continue
		}
		if pgm[p.GlobalName] == nil {
			pgm[p.GlobalName] = &modelsupervisor.PorterDigest{
				BinarySummary:     converter.ToBizPorter(p).BinarySummary,
				GlobalName:        p.GlobalName,
				Regions:           []string{p.Region},
				ContextJSONSchema: p.ContextJSONSchema,
				FeatureSummary:    p.FeatureSummary,
			}
		} else {
			pgm[p.GlobalName].Regions = append(pgm[p.GlobalName].Regions, p.Region)
		}
	}
	for _, v := range pgm {
		pg = append(pg, v)
	}
	return pg, nil
}

func (t *TipherethRepo) FetchPorterContext(
	ctx context.Context,
	id model.InternalID,
) (*modelsupervisor.PorterContext, error) {
	res, err := t.data.db.PorterContext.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return converter.ToBizPorterContext(res), nil
}

func (t *TipherethRepo) GetEnabledPorterContexts(
	ctx context.Context,
) ([]*modelsupervisor.PorterContext, error) {
	pc, err := t.data.db.PorterContext.Query().Where(
		portercontext.StatusEQ(portercontext.StatusActive),
	).All(ctx)
	if err != nil {
		return nil, err
	}
	return converter.ToBizPorterContextList(pc), nil
}
