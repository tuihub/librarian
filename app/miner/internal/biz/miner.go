package biz

import (
	"bytes"
	"context"
	"io"
	"net/http"

	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

type MinerRepo interface {
	RecognizeImage(context.Context, []byte) ([]*OCRResults, error)
}

type Miner struct {
	repo MinerRepo
}

func NewMiner(repo MinerRepo) *Miner {
	return &Miner{repo: repo}
}

type OCRResults struct {
	Confidence float64 `json:"confidence"`
	Text       string  `json:"text"`
}

func (m *Miner) FeatureEnabled() bool {
	return m.repo != nil
}

func (m *Miner) RecognizeImageURL(ctx context.Context, url string) ([]*OCRResults, *errors.Error) {
	if !m.FeatureEnabled() {
		return nil, errors.BadRequest("request disabled feature", "")
	}
	hc := new(http.Client)
	getReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, bytes.NewBufferString(""))
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	getResp, err := hc.Do(getReq)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	imgBytes, err := io.ReadAll(getResp.Body)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	if len(imgBytes) == 0 {
		return []*OCRResults{}, nil
	}
	err = getResp.Body.Close()
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	res, err := m.repo.RecognizeImage(ctx, imgBytes)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return res, nil
}
