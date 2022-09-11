package data

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/ent/user"

	"github.com/tuihub/librarian/app/sephirah/internal/biz"
)

type tipherethRepo struct {
	data *Data
}

// NewTipherethRepo .
func NewTipherethRepo(data *Data) biz.TipherethRepo {
	return &tipherethRepo{
		data: data,
	}
}

func (t tipherethRepo) UserActive(ctx context.Context, userData *biz.User) (bool, error) {
	u, err := t.data.db.User.Query().Where(
		user.UsernameEQ(userData.UserName),
		user.PasswordEQ(userData.PassWord),
	).First(ctx)
	if err != nil {
		return false, err
	}
	if u != nil && u.State == user.StateActive {
		return true, err
	}
	return false, err
}

func (t tipherethRepo) AddUser(ctx context.Context, userData *biz.User) (*biz.User, error) {
	_, err := t.data.db.User.Create().
		SetInternalID(userData.Id).
		SetUsername(userData.UserName).
		SetPassword(userData.PassWord).
		SetState(user.StateActive).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.User{}, nil
}
