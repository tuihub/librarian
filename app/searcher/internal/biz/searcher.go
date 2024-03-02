package biz

import (
	"context"

	"github.com/tuihub/librarian/internal/model"
)

type SearcherRepo interface {
	NewID(context.Context) (int64, error)
	DescribeID(context.Context, model.InternalID, Index, bool, string) error
	SearchID(context.Context, Index, model.Paging, string) ([]*SearchResult, error)
}

type Searcher struct {
	repo SearcherRepo
}

type SearchResult struct {
	ID   model.InternalID
	Rank int64
}

type Index int

const (
	IndexUnspecified Index = iota
	IndexGeneral
	IndexGeburaApp
	IndexChesedImage
)

func IndexNameMap() map[Index]string {
	return map[Index]string{
		IndexUnspecified: "",
		IndexGeneral:     "general",
		IndexGeburaApp:   "gebura",
		IndexChesedImage: "chesed",
	}
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

func (g *Searcher) DescribeID(
	ctx context.Context, id model.InternalID, index Index, append_ bool, description string,
) error {
	return g.repo.DescribeID(ctx, id, index, append_, description)
}

func (g *Searcher) SearchID(
	ctx context.Context, paging model.Paging, index Index, query string,
) ([]*SearchResult, error) {
	return g.repo.SearchID(ctx, index, paging, query)
}
