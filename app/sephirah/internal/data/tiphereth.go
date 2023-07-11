package data

import (
	"context"
	"errors"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/account"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
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
	userData *modeltiphereth.User,
) (*modeltiphereth.User, error) {
	u, err := t.data.db.User.Query().Where(
		user.UsernameEQ(userData.UserName),
		user.PasswordEQ(userData.PassWord),
	).First(ctx)
	if u == nil || err != nil {
		return nil, errors.New("invalid user")
	}
	return converter.ToBizUser(u), nil
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
	creator *model.InternalID,
) ([]*modeltiphereth.User, int64, error) {
	q := t.data.db.User.Query()
	if creator != nil {
		q = t.data.db.User.Query().Where(user.IDEQ(*creator)).QueryCreatedUser()
	}
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

func (t tipherethRepo) LinkAccount(ctx context.Context, a modeltiphereth.Account, u model.InternalID) error {
	return t.data.WithTx(ctx, func(tx *ent.Tx) error {
		acc, err := tx.Account.Query().Where(
			account.PlatformEQ(converter.ToEntAccountPlatform(a.Platform)),
			account.PlatformAccountIDEQ(a.PlatformAccountID),
		).Only(ctx)
		if ent.IsNotFound(err) {
			return t.data.db.Account.Create().
				SetBindUserID(u).
				SetID(a.ID).
				SetPlatform(converter.ToEntAccountPlatform(a.Platform)).
				SetPlatformAccountID(a.PlatformAccountID).
				SetName(a.Name).
				SetAvatarURL(a.AvatarURL).
				SetProfileURL(a.ProfileURL).
				Exec(ctx)
		}
		if err != nil {
			return err
		}
		exist, err := acc.QueryBindUser().Exist(ctx)
		if err != nil {
			return err
		}
		if exist {
			return errors.New("account already bound to an user")
		}
		return t.data.db.Account.UpdateOneID(acc.ID).
			SetBindUserID(u).
			Exec(ctx)
	})
}

func (t tipherethRepo) UnLinkAccount(ctx context.Context, a modeltiphereth.Account, u model.InternalID) error {
	return t.data.db.Account.Update().Where(
		account.PlatformEQ(converter.ToEntAccountPlatform(a.Platform)),
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
