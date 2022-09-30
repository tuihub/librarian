package data

import (
	"errors"

	"github.com/tuihub/librarian/app/mapper/internal/biz"

	"github.com/cayleygraph/cayley"
	"github.com/google/wire"
	"github.com/zhihu/norm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewNebula, NewCayley, NewMapperRepo)

// NewMapperRepo .
func NewMapperRepo(n *norm.DB, c *cayley.Handle) (biz.MapperRepo, error) {
	if n != nil {
		return &nebulaMapperRepo{
			db: n,
		}, nil
	}
	if c != nil {
		return &cayleyMapperRepo{
			db: c,
		}, nil
	}
	return nil, errors.New("no valid db backend")
}
