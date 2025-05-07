package libsearch

import (
	"context"
	"fmt"

	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewSearch)

type Search interface {
	DescribeID(context.Context, model.InternalID, SearchIndex, bool, string) error
	SearchID(context.Context, SearchIndex, model.Paging, string) ([]*SearchResult, error)
}

type SearchResult struct {
	ID   model.InternalID
	Rank int64
}

type SearchIndex int

const (
	SearchIndexUnspecified SearchIndex = iota
	SearchIndexGeneral
	SearchIndexGeburaApp
	SearchIndexChesedImage
)

func SearchIndexNameMap() map[SearchIndex]string {
	return map[SearchIndex]string{
		SearchIndexUnspecified: "",
		SearchIndexGeneral:     "general",
		SearchIndexGeburaApp:   "gebura",
		SearchIndexChesedImage: "chesed",
	}
}

func NewSearch(
	c *conf.Search, app *libapp.Settings,
) (Search, error) {
	switch c.Driver {
	case conf.SearchDriverMeili:
		m := newMeili(c)
		return &meiliSearcherRepo{
			search: m,
		}, nil
	case conf.SearchDriverBleve:
		b, err := newBleve(c, app)
		if err != nil {
			return nil, fmt.Errorf("failed creating bleve searcher: %w", err)
		}
		return &bleveSearcherRepo{
			search: b,
		}, nil
	default:
		logger.Warnf("no valid search backend, search function will not work")
		return &defaultSearcherRepo{}, nil
	}
}

type defaultSearcherRepo struct {
}

func (d defaultSearcherRepo) DescribeID(
	context.Context, model.InternalID, SearchIndex, bool, string) error {
	return nil // search disabled
}

func (d defaultSearcherRepo) SearchID(
	context.Context, SearchIndex, model.Paging, string) ([]*SearchResult, error) {
	return nil, nil // search disabled
}
