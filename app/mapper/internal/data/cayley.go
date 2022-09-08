package data

import (
	"context"

	"github.com/cayleygraph/cayley"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tuihub/librarian/app/mapper/internal/biz"
	"github.com/tuihub/librarian/internal/conf"
)

type cayleyMapperRepo struct {
	data *Data
}

func (r *cayleyMapperRepo) Save(ctx context.Context, g *biz.Mapper) (*biz.Mapper, error) {
	return g, nil
}

// NewCayley .
func NewCayley(c *conf.Mapper_Data) (*cayley.Handle, func()) {
	if c == nil || c.GetCayley() == nil {
		return nil, func() {}
	}
	if c.GetCayley().GetStore() != "memory" {
		log.Errorf("Unsupported cayley store: %s, skip initialize", c.GetCayley().GetStore())
		return nil, func() {}
	}

	db, err := cayley.NewMemoryGraph()
	if err != nil {
		log.Errorf("Failed to initialize Cayley DB, %s", err.Error())
	}

	return db, func() {
		log.Info("closing the data resources")
		db.Close()
	}
}
