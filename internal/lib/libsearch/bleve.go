package libsearch

import (
	"context"
	"errors"
	"fmt"
	"path"
	"strconv"

	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libcodec"
	"github.com/tuihub/librarian/internal/model"

	"github.com/blevesearch/bleve/v2"
)

type bleveSearcherRepo struct {
	search map[SearchIndex]bleve.Index
}

func newBleve(c *conf.Search, app *libapp.Settings) (map[SearchIndex]bleve.Index, error) {
	if c.Driver != conf.SearchDriverBleve {
		return nil, nil //nolint:nilnil //TODO
	}
	res := make(map[SearchIndex]bleve.Index)
	for i, n := range SearchIndexNameMap() {
		if i == SearchIndexUnspecified {
			continue
		}
		mapping := bleve.NewIndexMapping()
		dbPath := path.Join(app.DataPath, fmt.Sprintf("bleve-%s.db", n))
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
		res[i] = index
	}
	return res, nil
}

func (r *bleveSearcherRepo) DescribeID(
	ctx context.Context, id model.InternalID, index SearchIndex, _ bool, description string,
) error {
	var jsonDesc interface{}
	err := libcodec.Unmarshal(libcodec.JSON, []byte(description), &jsonDesc)
	if err == nil {
		err = r.search[index].Index(strconv.FormatInt(int64(id), 10), jsonDesc)
		if err != nil {
			return err
		}
	} else {
		err = r.search[index].Index(strconv.FormatInt(int64(id), 10), description)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *bleveSearcherRepo) SearchID(ctx context.Context, index SearchIndex, paging model.Paging, query string) (
	[]*SearchResult, error) {
	search := bleve.NewSearchRequest(bleve.NewFuzzyQuery(query))
	search.From = paging.ToOffset()
	search.Size = paging.ToLimit()
	result, err := r.search[index].Search(search)
	if err != nil {
		return nil, err
	}
	res := make([]*SearchResult, 0, 20) //nolint:mnd // TODO
	for _, h := range result.Hits {
		var id int64
		id, err = strconv.ParseInt(h.ID, 10, 64)
		if err == nil {
			res = append(res, &SearchResult{
				ID:   model.InternalID(id),
				Rank: int64(h.Score),
			})
		}
	}
	return res, nil
}
