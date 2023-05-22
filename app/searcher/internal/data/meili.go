package data

import (
	"context"

	"github.com/tuihub/librarian/app/searcher/internal/biz"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libcodec"
	"github.com/tuihub/librarian/internal/model"

	"github.com/meilisearch/meilisearch-go"
	"github.com/sony/sonyflake"
)

const IndexName = "default"

type meiliSearcherRepo struct {
	sf     *sonyflake.Sonyflake
	search *meilisearch.Client
}

func NewMeili(conf *conf.Searcher_Data, app *libapp.Settings) (*meilisearch.Client, error) {
	if conf.GetMeilisearch() == nil {
		return nil, nil //nolint:nilnil //TODO
	}
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:    conf.GetMeilisearch().GetAddr(),
		APIKey:  conf.GetMeilisearch().GetApiKey(),
		Timeout: 0,
	})
	return client, nil
}

type document struct {
	ID          model.InternalID
	Description interface{}
}

func (m *meiliSearcherRepo) DescribeID(ctx context.Context, id model.InternalID, description string) error {
	var jsonDesc interface{}
	err := libcodec.Unmarshal(libcodec.JSON, []byte(description), &jsonDesc)
	if err == nil {
		documents := map[string]interface{}{
			"id":          id,
			"description": jsonDesc,
		}
		_, err = m.search.Index(IndexName).AddDocuments(documents)
		if err != nil {
			return err
		}
	} else {
		documents := map[string]interface{}{
			"id":          id,
			"description": description,
		}
		_, err = m.search.Index(IndexName).AddDocuments(documents)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *meiliSearcherRepo) SearchID(ctx context.Context, paging model.Paging, keyword string) (
	[]*biz.SearchResult, error) {
	request := &meilisearch.SearchRequest{ //nolint:exhaustruct //TODO
		Limit:  int64(paging.ToLimit()),
		Offset: int64(paging.ToOffset()),
	}
	// https://github.com/meilisearch/meilisearch-go/issues/406
	resultRaw, err := m.search.Index(IndexName).SearchRaw(keyword, request)
	if err != nil {
		return nil, err
	}
	resultStr, err := resultRaw.MarshalJSON()
	if err != nil {
		return nil, err
	}
	result := struct {
		Hits []document
	}{}
	err = libcodec.Unmarshal(libcodec.JSON, resultStr, &result)
	if err != nil {
		return nil, err
	}
	res := make([]*biz.SearchResult, 0, 20) //nolint:gomnd // TODO
	for _, h := range result.Hits {
		var str []byte
		str, err = libcodec.Marshal(libcodec.JSON, h)
		if err != nil {
			continue
		}
		var d document
		err = libcodec.Unmarshal(libcodec.JSON, str, &d)
		if err != nil {
			continue
		}
		res = append(res, &biz.SearchResult{
			ID:   d.ID,
			Rank: 0,
		})
	}
	return res, nil
}