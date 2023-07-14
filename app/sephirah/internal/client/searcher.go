package client

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/model/converter"
	"github.com/tuihub/librarian/internal/lib/libcodec"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"
)

type Searcher struct {
	client searcher.LibrarianSearcherServiceClient
}

func NewSearcher(
	client searcher.LibrarianSearcherServiceClient,
) *Searcher {
	return &Searcher{client: client}
}

func (s *Searcher) NewID(ctx context.Context) (model.InternalID, error) {
	resp, err := s.client.NewID(ctx, &searcher.NewIDRequest{})
	if err != nil {
		logger.Infof("NewID failed: %s", err.Error())
		return 0, err
	}
	return converter.ToBizInternalID(resp.GetId()), nil
}

func (s *Searcher) NewBatchIDs(ctx context.Context, num int) ([]model.InternalID, error) {
	resp, err := s.client.NewBatchIDs(ctx, &searcher.NewBatchIDsRequest{
		Num: int32(num),
	})
	if err != nil {
		logger.Infof("NewBatchIDs failed: %s", err.Error())
		return nil, err
	}
	return converter.ToBizInternalIDList(resp.GetIds()), nil
}

func (s *Searcher) DescribeID(
	ctx context.Context,
	id model.InternalID,
	desc interface{},
	mode searcher.DescribeIDRequest_DescribeMode,
	index searcher.Index,
) error {
	descStr, err := libcodec.Marshal(libcodec.JSON, desc)
	if err != nil {
		return err
	}
	_, err = s.client.DescribeID(ctx, &searcher.DescribeIDRequest{
		Id:          converter.ToPBInternalID(id),
		Description: string(descStr),
		Mode:        mode,
		Index:       index,
	})
	return err
}

func (s *Searcher) SearchID(
	ctx context.Context, paging model.Paging, keyword string, index searcher.Index,
) ([]model.InternalID, error) {
	resp, err := s.client.SearchID(ctx, &searcher.SearchIDRequest{
		Paging:  model.ToPBPaging(paging),
		Keyword: keyword,
		Index:   index,
	})
	if err != nil {
		return nil, err
	}
	res := make([]model.InternalID, 0, len(resp.GetResult()))
	for _, r := range resp.GetResult() {
		res = append(res, converter.ToBizInternalID(r.Id))
	}
	return res, nil
}
