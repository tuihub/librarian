package data

import (
	"strconv"

	"github.com/blevesearch/bleve/v2"
)

type Data struct {
	index bleve.Index
}

func NewData() (*Data, error) {
	mapping := bleve.NewIndexMapping()
	index, err := bleve.NewMemOnly(mapping)
	if err != nil {
		return nil, err
	}
	return &Data{index: index}, nil
}

func (d *Data) Index(id int, data string) error {
	return d.index.Index(strconv.Itoa(id), data)
}

func (d *Data) Search(query string) ([]string, error) {
	q := bleve.NewFuzzyQuery(query)
	search := bleve.NewSearchRequest(q)
	searchResults, err := d.index.Search(search)
	if err != nil {
		return nil, err
	}
	if len(searchResults.Hits) == 0 {
		q2 := bleve.NewMatchPhraseQuery(query)
		search = bleve.NewSearchRequest(q2)
		searchResults, err = d.index.Search(search)
		if err != nil {
			return nil, err
		}
	}
	var res []string
	for _, hit := range searchResults.Hits {
		res = append(res, hit.ID)
	}
	return res, nil
}
