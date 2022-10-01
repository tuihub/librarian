package data

import (
	"context"
	"errors"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/user"
	"github.com/tuihub/librarian/internal/lib/libauth"
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

func (t tipherethRepo) UserActive(ctx context.Context, userData *biztiphereth.User) (bool, error) {
	u, err := t.data.db.User.Query().Where(
		user.UsernameEQ(userData.UserName),
		user.PasswordEQ(userData.PassWord),
	).First(ctx)
	if err != nil {
		return false, err
	}
	if u != nil && u.Status == user.StatusActive {
		return true, err
	}
	return false, err
}

func (t tipherethRepo) FetchUserByPassword(
	ctx context.Context,
	userData *biztiphereth.User,
) (*biztiphereth.User, error) {
	u, err := t.data.db.User.Query().Where(
		user.UsernameEQ(userData.UserName),
		user.PasswordEQ(userData.PassWord),
	).First(ctx)
	if err != nil {
		return nil, err
	}
	if u != nil {
		userData.InternalID = u.InternalID
		userData.Type = toLibAuthUserType(u.Type)
		return userData, nil
	}
	return nil, errors.New("invalid user")
}

func (t tipherethRepo) AddUser(ctx context.Context, userData *biztiphereth.User) error {
	userType := toEntUserType(userData.Type)
	err := t.data.db.User.Create().
		SetInternalID(userData.InternalID).
		SetUsername(userData.UserName).
		SetPassword(userData.PassWord).
		SetStatus(toEntUserStatus(userData.Status)).
		SetType(userType).
		Exec(ctx)
	return err
}

func (t tipherethRepo) UpdateUser(ctx context.Context, u *biztiphereth.User) error {
	update := t.data.db.User.Update().
		Where(user.InternalIDEQ(u.InternalID))
	if u.UserName != "" {
		update = update.SetUsername(u.UserName)
	}
	if u.PassWord != "" {
		update = update.SetPassword(u.PassWord)
	}
	if u.Status != biztiphereth.UserStatusUnspecified {
		update = update.SetStatus(toEntUserStatus(u.Status))
	}
	return update.Exec(ctx)
}

func (t tipherethRepo) ListUser(
	ctx context.Context,
	paging biztiphereth.Paging,
	types []libauth.UserType,
	statuses []biztiphereth.UserStatus,
) ([]*biztiphereth.User, error) {
	q := t.data.db.User.Query()
	if len(types) > 0 {
		typeFilter := make([]user.Type, len(types))
		for i, userType := range types {
			typeFilter[i] = toEntUserType(userType)
		}
		q.Where(user.TypeIn(typeFilter...))
	}
	if len(statuses) > 0 {
		statusFilter := make([]user.Status, len(statuses))
		for i, userStatus := range statuses {
			statusFilter[i] = toEntUserStatus(userStatus)
		}
		q.Where(user.StatusIn(statusFilter...))
	}
	u, err := q.
		Limit(paging.PageSize).
		Offset((paging.PageNum - 1) * paging.PageSize).
		All(ctx)
	if err != nil {
		return nil, err
	}
	users := make([]*biztiphereth.User, len(u))
	for i, su := range u {
		users[i] = toBizUser(su)
	}
	return users, nil
}

func (t tipherethRepo) CreateAccount(ctx context.Context, a biztiphereth.Account) error {
	return t.data.db.Account.Create().
		SetInternalID(a.InternalID).
		SetPlatform(toEntAccountPlatform(a.Platform)).
		SetPlatformAccountID(a.PlatformAccountID).
		SetName(a.Name).
		SetAvatarURL(a.AvatarURL).
		SetProfileURL(a.ProfileURL).
		Exec(ctx)
}
