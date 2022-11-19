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
	VertexTypeMetadata
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
	EdgeTypeDescribe
)

type MapperRepo interface {
	InsertVertex(context.Context, []*Vertex) error
	InsertEdge(context.Context, []*Edge) error
	FetchEqualVertex(context.Context, Vertex) ([]*Vertex, error)
}

type Mapper struct {
	repo MapperRepo
}

func NewMapper(repo MapperRepo) *Mapper {
	return &Mapper{repo: repo}
}

func (m *Mapper) InsertVertex(ctx context.Context, v []*Vertex) error {
	return m.repo.InsertVertex(ctx, v)
}

func (m *Mapper) InsertEdge(ctx context.Context, e []*Edge) error {
	return m.repo.InsertEdge(ctx, e)
}

func (m *Mapper) FetchEqualVertex(ctx context.Context, v Vertex) ([]*Vertex, error) {
	return m.repo.FetchEqualVertex(ctx, v)
}
