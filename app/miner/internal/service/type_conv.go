package service

import (
	"github.com/tuihub/librarian/app/miner/internal/biz"
	pb "github.com/tuihub/protos/pkg/librarian/miner/v1"
)

func ToPBOCRResults(rs []*biz.OCRResults) []*pb.RecognizeImageResult {
	res := make([]*pb.RecognizeImageResult, 0, len(rs))
	for _, r := range rs {
		res = append(res, &pb.RecognizeImageResult{
			Confidence: r.Confidence,
			Text:       r.Text,
		})
	}
	return res
}
