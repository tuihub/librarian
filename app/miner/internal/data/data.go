package data

import (
	"bytes"
	"context"
	"encoding/base64"
	"io"
	"net/http"

	"github.com/tuihub/librarian/app/miner/internal/biz"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libcodec"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewMinerRepo)

// NewMinerRepo .
func NewMinerRepo(conf *conf.Miner_Data) biz.MinerRepo {
	if conf.Ocr == nil || len(conf.Ocr.GetAddress()) == 0 {
		return nil
	}
	return &minerRepo{
		c:    conf,
		http: new(http.Client),
	}
}

type minerRepo struct {
	c    *conf.Miner_Data
	http *http.Client
}

type OCRResponse struct {
	Msg     string         `json:"msg"`
	Results [][]OCRResults `json:"results"`
	Status  string         `json:"status"`
}
type OCRResults struct {
	Confidence float64 `json:"confidence"`
	Text       string  `json:"text"`
	TextRegion [][]int `json:"text_region"`
}

func (m *minerRepo) RecognizeImage(ctx context.Context, imgBytes []byte) ([]*biz.OCRResults, error) {
	ocrReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		m.c.Ocr.GetAddress(),
		bytes.NewBuffer([]byte("{\"images\":[\""+base64.StdEncoding.EncodeToString(imgBytes)+"\"]}")),
	)
	if err != nil {
		return nil, err
	}
	ocrReq.Header.Set("Content-Type", "application/json")
	ocrResp, err := m.http.Do(ocrReq)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(ocrResp.Body)
	if err != nil {
		return nil, err
	}
	err = ocrResp.Body.Close()
	if err != nil {
		return nil, err
	}
	ocrResponse := OCRResponse{} //nolint:exhaustruct //TODO
	err = libcodec.Unmarshal(libcodec.JSON, body, &ocrResponse)
	if err != nil {
		return nil, err
	}
	if len(ocrResponse.Results) != 1 {
		return nil, err
	}
	res := make([]*biz.OCRResults, 0, len(ocrResponse.Results))
	for _, resp := range ocrResponse.Results[0] {
		res = append(res, &biz.OCRResults{
			Confidence: resp.Confidence,
			Text:       resp.Text,
		})
	}
	return res, nil
}
