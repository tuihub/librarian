package data

import (
	"errors"

	"github.com/tuihub/librarian/app/mapper/internal/biz"

	"github.com/cayleygraph/cayley"
	"github.com/google/wire"
	"github.com/zhihu/norm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewNebula, NewCayley, NewMapperRepo)

// Data .
type Data struct {
	ndb *norm.DB
	cdb *cayley.Handle
}

// NewData .
func NewData(ndb *norm.DB, cdb *cayley.Handle) *Data {
	return &Data{
		ndb,
		cdb,
	}
}

// NewMapperRepo .
func NewMapperRepo(data *Data) (biz.MapperRepo, error) {
	if data.ndb != nil {
		return &nebulaMapperRepo{
			data: data,
		}, nil
	}
	if data.cdb != nil {
		return &cayleyMapperRepo{
			data: data,
		}, nil
	}
	return nil, errors.New("no valid db backend")
}
