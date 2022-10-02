package biz

import (
	"context"
)

type SearcherRepo interface {
	NewID(ctx context.Context) (int64, error)
}

type Searcher struct {
	repo SearcherRepo
}

func NewSearcher(repo SearcherRepo) *Searcher {
	return &Searcher{repo: repo}
}

func (g *Searcher) NewID(ctx context.Context) (int64, error) {
	return g.repo.NewID(ctx)
}
