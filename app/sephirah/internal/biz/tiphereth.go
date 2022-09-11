package biz

import (
	"context"
	"errors"
	"time"

	"github.com/tuihub/librarian/internal/lib/libauth"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"
)

// User is a User model.
type User struct {
	ID       int64
	UserName string
	PassWord string
}

type AccessToken string

// TipherethRepo is a Greater repo.
type TipherethRepo interface {
	UserActive(context.Context, *User) (bool, error)
	AddUser(context.Context, *User) (*User, error)
}

// TipherethUsecase is a User usecase.
type TipherethUsecase struct {
	auth     *libauth.Auth
	repo     TipherethRepo
	searcher searcher.LibrarianSearcherServiceClient
}

// NewTipherethUsecase new a User usecase.
func NewTipherethUsecase(repo TipherethRepo, auth *libauth.Auth, mClient mapper.LibrarianMapperServiceClient,
	sClient searcher.LibrarianSearcherServiceClient, pClient porter.LibrarianPorterServiceClient) *TipherethUsecase {
	return &TipherethUsecase{auth: auth, repo: repo, searcher: sClient}
}

func (t *TipherethUsecase) UserLogin(ctx context.Context, user *User) (AccessToken, error) {
	password, err := t.auth.GeneratePassword(user.PassWord)
	if err != nil {
		return "", errors.New("internal error")
	}
	ok, err := t.repo.UserActive(ctx, &User{
		UserName: user.UserName,
		PassWord: password,
	})
	if err != nil {
		return "", err
	}
	if ok {
		var token string
		token, err = t.auth.GenerateToken(1, 1, time.Hour)
		if err != nil {
			return "", err
		}
		return AccessToken(token), nil
	}
	return "", errors.New("invalid user")
}

func (t *TipherethUsecase) AddUser(ctx context.Context, user *User) (*User, error) {
	password, err := t.auth.GeneratePassword(user.PassWord)
	if err != nil {
		return nil, err
	}
	resp, err := t.searcher.NewID(ctx, &searcher.NewIDRequest{})
	if err != nil {
		return nil, err
	}
	_, err = t.repo.AddUser(ctx, &User{
		ID:       resp.Id,
		UserName: user.UserName,
		PassWord: password,
	})
	if err != nil {
		return nil, err
	}
	return &User{
		ID: resp.Id,
	}, nil
}
