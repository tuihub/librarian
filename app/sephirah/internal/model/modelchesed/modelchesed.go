package modelchesed

import (
	"github.com/tuihub/librarian/model"
)

type Image struct {
	ID          model.InternalID
	Name        string
	Description string
	Status      ImageStatus
}

type ImageStatus int

const (
	ImageStatusUnspecified ImageStatus = iota
	ImageStatusUploaded
	ImageStatusScanned
)

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
