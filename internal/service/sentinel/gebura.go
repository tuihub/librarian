package sentinel

import (
	"context"

	"github.com/tuihub/librarian/internal/service/sentinel/converter"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1/sentinel"
)

func (s *LibrarianSentinelService) ReportSentinelInformation(
	ctx context.Context, req *pb.ReportSentinelInformationRequest,
) (*pb.ReportSentinelInformationResponse, error) {
	err := s.g.UpsertSentinelInfo(ctx, converter.ToBizSentinel(req))
	if err != nil {
		return nil, err
	}
	return &pb.ReportSentinelInformationResponse{}, nil
}

func (s *LibrarianSentinelService) ReportAppBinaries(
	ctx context.Context, req *pb.ReportAppBinariesRequest,
) (*pb.ReportAppBinariesResponse, error) {
	err := s.g.UpsertAppBinaries(ctx, converter.ToBizSentinelAppBinaryList(req.GetAppBinaries()))
	if err != nil {
		return nil, err
	}
	return &pb.ReportAppBinariesResponse{}, nil
}
