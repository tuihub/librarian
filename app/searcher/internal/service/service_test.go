package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tuihub/librarian/app/searcher/internal/biz"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"
)

type mockedGreeterRepo struct {
	mock.Mock
}

func (m *mockedGreeterRepo) NewID(ctx context.Context) (int64, error) {
	args := m.Called(ctx)
	return args.Get(0).(int64), args.Error(1)
}

func TestLibrarianSearcherServiceService_NewID(t *testing.T) {
	m := mockedGreeterRepo{}
	s := LibrarianSearcherServiceService{
		UnimplementedLibrarianSearcherServiceServer: searcher.UnimplementedLibrarianSearcherServiceServer{},
		uc: biz.NewGreeterUseCase(&m),
	}

	mc := m.On("NewID", context.Background()).Return((int64)(123), nil)
	id, err := s.NewID(context.Background(), &searcher.NewIDRequest{})
	m.AssertExpectations(t)
	assert.Nil(t, err)
	if assert.NotNil(t, id) {
		assert.Equal(t, (int64)(123), id.Id)
	}
	mc.Unset()

	e := errors.New("test error")
	m.On("NewID", context.Background()).Return((int64)(0), e)
	id, err = s.NewID(context.Background(), &searcher.NewIDRequest{})
	m.AssertExpectations(t)
	assert.Equal(t, err, e)
	assert.Nil(t, id)
}
