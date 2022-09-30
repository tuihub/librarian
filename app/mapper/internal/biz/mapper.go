package biz

import (
	"context"
)

type Vertex struct {
	ID   int64
	Type VertexType
}

type VertexType int

const (
	VertexTypeUnspecified VertexType = iota
	VertexTypeAbstract
	VertexTypeEntity
	VertexTypeMessage
	VertexTypeObject
)

type Edge struct {
	SourceID      int64
	DestinationID int64
	Type          EdgeType
}

type EdgeType int

const (
	EdgeTypeUnspecified EdgeType = iota
	EdgeTypeGeneral
	EdgeTypeEqual
	EdgeTypeCreate
	EdgeTypeEnjoy
	EdgeTypeMention
	EdgeTypeDerive
	EdgeTypeControl
	EdgeTypeFollow
)

// MapperRepo is a Greater repo.
type MapperRepo interface {
	InsertVertex(context.Context, Vertex) error
	InsertEdge(context.Context, Edge) error
	FetchEqualVertex(context.Context, Vertex) ([]*Vertex, error)
}

// MapperUseCase is a Mapper use case.
type MapperUseCase struct {
	repo MapperRepo
}

// NewMapperUseCase new a Mapper use case.
func NewMapperUseCase(repo MapperRepo) *MapperUseCase {
	return &MapperUseCase{repo: repo}
}

func (m *MapperUseCase) InsertVertex(ctx context.Context, v Vertex) error {
	return m.repo.InsertVertex(ctx, v)
}

func (m *MapperUseCase) InsertEdge(ctx context.Context, e Edge) error {
	return m.repo.InsertEdge(ctx, e)
}

func (m *MapperUseCase) FetchEqualVertex(ctx context.Context, v Vertex) ([]*Vertex, error) {
	return m.repo.FetchEqualVertex(ctx, v)
}
