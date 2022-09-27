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
		userData.UserType = toLibAuthUserType(u.Type)
		return userData, nil
	}
	return nil, errors.New("invalid user")
}

func (t tipherethRepo) AddUser(ctx context.Context, userData *biztiphereth.User) error {
	userType := toEntUserType(userData.UserType)
	_, err := t.data.db.User.Create().
		SetInternalID(userData.InternalID).
		SetUsername(userData.UserName).
		SetPassword(userData.PassWord).
		SetStatus(user.StatusActive).
		SetType(userType).
		Save(ctx)
	return err
}

func (t tipherethRepo) ListUser(
	ctx context.Context,
	paging biztiphereth.Paging,
	types []libauth.UserType,
	statuses []biztiphereth.UserStatus,
) ([]*biztiphereth.User, error) {
	var typeFilter []user.Type
	for _, userType := range types {
		typeFilter = append(typeFilter, toEntUserType(userType))
	}
	var statusFilter []user.Status
	for _, userStatus := range statuses {
		statusFilter = append(statusFilter, toEntUserStatus(userStatus))
	}
	u, err := t.data.db.User.Query().
		Where(
			user.TypeIn(typeFilter...),
			user.StatusIn(statusFilter...),
		).
		Limit(paging.PageSize).
		Offset((paging.PageSize - 1) * paging.PageNum).
		All(ctx)
	if err != nil {
		return nil, err
	}
	var users []*biztiphereth.User
	for _, su := range u {
		users = append(users, toBizUser(su))
	}
	return users, nil
}
