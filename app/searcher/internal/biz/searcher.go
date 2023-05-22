package biz

import (
	"context"

	"github.com/tuihub/librarian/internal/model"
)

type SearcherRepo interface {
	NewID(context.Context) (int64, error)
	DescribeID(context.Context, model.InternalID, string) error
	SearchID(context.Context, model.Paging, string) ([]*SearchResult, error)
}

type Searcher struct {
	repo SearcherRepo
}

type SearchResult struct {
	ID   model.InternalID
	Rank int64
}

func NewSearcher(repo SearcherRepo) *Searcher {
	return &Searcher{repo: repo}
}

func (g *Searcher) NewID(ctx context.Context) (int64, error) {
	return g.repo.NewID(ctx)
}

func (g *Searcher) NewBatchIDs(ctx context.Context, num int) ([]int64, error) {
	var res []int64
	for i := 0; i < num; i++ {
		id, err := g.repo.NewID(ctx)
		if err != nil {
			return nil, err
		}
		res = append(res, id)
	}
	return res, nil
}

func (g *Searcher) DescribeID(ctx context.Context, id model.InternalID, description string) error {
	return g.repo.DescribeID(ctx, id, description)
}

func (g *Searcher) SearchID(ctx context.Context, paging model.Paging, keyword string) ([]*SearchResult, error) {
	return g.repo.SearchID(ctx, paging, keyword)
}
