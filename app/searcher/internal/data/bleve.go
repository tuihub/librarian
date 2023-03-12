package data

import (
	"context"
	"errors"
	"path"
	"strconv"

	"github.com/tuihub/librarian/app/searcher/internal/biz"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libcodec"
	"github.com/tuihub/librarian/internal/model"

	"github.com/blevesearch/bleve/v2"
	"github.com/sony/sonyflake"
)

type bleveSearcherRepo struct {
	sf     *sonyflake.Sonyflake
	search bleve.Index
}

func NewBleve() (bleve.Index, error) {
	mapping := bleve.NewIndexMapping()
	dbPath := path.Join(libapp.GetDataPath(), "bleve.db")
	index, err := bleve.Open(dbPath)
	if err != nil {
		if !errors.Is(err, bleve.ErrorIndexPathDoesNotExist) {
			return nil, err
		} else {
			index, err = bleve.New(dbPath, mapping)
			if err != nil {
				return nil, err
			}
		}
	}
	return index, nil
}

func (r *bleveSearcherRepo) DescribeID(ctx context.Context, id model.InternalID, description string) error {
	var jsonDesc interface{}
	err := libcodec.Unmarshal(libcodec.JSON, []byte(description), &jsonDesc)
	if err == nil {
		err = r.search.Index(strconv.FormatInt(int64(id), 10), jsonDesc)
		if err != nil {
			return err
		}
	} else {
		err = r.search.Index(strconv.FormatInt(int64(id), 10), description)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *bleveSearcherRepo) SearchID(ctx context.Context, keyword string) ([]*biz.SearchResult, error) {
	query := bleve.NewFuzzyQuery(keyword)
	search := bleve.NewSearchRequest(query)
	result, err := r.search.Search(search)
	if err != nil {
		return nil, err
	}
	res := make([]*biz.SearchResult, 0, 20) //nolint:gomnd // TODO
	for _, h := range result.Hits {
		var id int64
		id, err = strconv.ParseInt(h.ID, 10, 64)
		if err == nil {
			res = append(res, &biz.SearchResult{
				ID:   model.InternalID(id),
				Rank: int64(h.Score),
			})
		}
	}
	return res, nil
}
