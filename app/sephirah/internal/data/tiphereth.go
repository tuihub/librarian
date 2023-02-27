package data

import (
	"context"
	"errors"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/data/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/account"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/user"
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
	userData *biztiphereth.User,
) (*biztiphereth.User, error) {
	u, err := t.data.db.User.Query().Where(
		user.UsernameEQ(userData.UserName),
		user.PasswordEQ(userData.PassWord),
	).First(ctx)
	if u == nil || err != nil {
		return nil, errors.New("invalid user")
	}
	return t.data.converter.ToBizUser(u), nil
}

func (t tipherethRepo) CreateUser(ctx context.Context, u *biztiphereth.User, c *model.InternalID) error {
	q := t.data.db.User.Create().
		SetID(u.InternalID).
		SetUsername(u.UserName).
		SetPassword(u.PassWord).
		SetStatus(converter.ToEntUserStatus(u.Status)).
		SetType(converter.ToEntUserType(u.Type))
	if c != nil {
		q.SetCreatorID(int64(*c))
	}
	return q.Exec(ctx)
}

func (t tipherethRepo) UpdateUser(ctx context.Context, u *biztiphereth.User) error {
	q := t.data.db.User.Update().
		Where(user.IDEQ(u.InternalID))
	if u.UserName != "" {
		q.SetUsername(u.UserName)
	}
	if u.PassWord != "" {
		q.SetPassword(u.PassWord)
	}
	if u.Type != libauth.UserTypeUnspecified {
		q.SetType(converter.ToEntUserType(u.Type))
	}
	if u.Status != biztiphereth.UserStatusUnspecified {
		q.SetStatus(converter.ToEntUserStatus(u.Status))
	}
	return q.Exec(ctx)
}

func (t tipherethRepo) ListUser(
	ctx context.Context,
	paging model.Paging,
	ids []model.InternalID,
	types []libauth.UserType,
	statuses []biztiphereth.UserStatus,
	exclude []model.InternalID,
	creator *model.InternalID,
) ([]*biztiphereth.User, error) {
	q := t.data.db.User.Query()
	if creator != nil {
		q = t.data.db.User.Query().Where(user.IDEQ(int64(*creator))).QueryCreate()
	}
	if len(ids) > 0 {
		q.Where(user.IDIn(t.data.converter.ToEntInternalIDList(ids)...))
	}
	if len(types) > 0 {
		q.Where(user.TypeIn(t.data.converter.ToEntUserTypeList(types)...))
	}
	if len(statuses) > 0 {
		q.Where(user.StatusIn(t.data.converter.ToEntUserStatusList(statuses)...))
	}
	if len(exclude) > 0 {
		q.Where(user.IDNotIn(t.data.converter.ToEntInternalIDList(exclude)...))
	}
	u, err := q.
		Limit(paging.PageSize).
		Offset((paging.PageNum - 1) * paging.PageSize).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return t.data.converter.ToBizUserList(u), nil
}

func (t tipherethRepo) CreateAccount(ctx context.Context, a biztiphereth.Account) error {
	return t.data.db.Account.Create().
		SetID(a.InternalID).
		SetPlatform(converter.ToEntAccountPlatform(a.Platform)).
		SetPlatformAccountID(a.PlatformAccountID).
		SetName(a.Name).
		SetAvatarURL(a.AvatarURL).
		SetProfileURL(a.ProfileURL).
		Exec(ctx)
}

func (t tipherethRepo) UpdateAccount(ctx context.Context, a biztiphereth.Account) error {
	return t.data.db.Account.Update().Where(
		account.IDEQ(a.InternalID),
		account.PlatformEQ(converter.ToEntAccountPlatform(a.Platform)),
		account.PlatformAccountIDEQ(a.PlatformAccountID),
	).
		SetName(a.Name).
		SetProfileURL(a.ProfileURL).
		SetAvatarURL(a.AvatarURL).
		Exec(ctx)
}
