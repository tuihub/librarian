package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Mapper is a Mapper model.
type Mapper struct {
	Hello string
}

type Vertex struct {
	id  int64
	tag string
}
type Edge struct {
	sid int64
	did int64
	tag string
}

// MapperRepo is a Greater repo.
type MapperRepo interface {
	Save(context.Context, *Mapper) (*Mapper, error)
}

// MapperUseCase is a Mapper use case.
type MapperUseCase struct {
	repo MapperRepo
}

// NewMapperUseCase new a Mapper use case.
func NewMapperUseCase(repo MapperRepo) *MapperUseCase {
	return &MapperUseCase{repo: repo}
}

// CreateMapper creates a Mapper, and returns the new Mapper.
func (uc *MapperUseCase) CreateMapper(ctx context.Context, g *Mapper) (*Mapper, error) {
	log.Context(ctx).Infof("CreateMapper: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
