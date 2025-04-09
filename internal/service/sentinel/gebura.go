package sentinel

import (
	"context"

	"github.com/tuihub/librarian/internal/service/sentinel/converter"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1/sentinel"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LibrarianSentinelService) ReportSentinelInformation(
	ctx context.Context, req *pb.ReportSentinelInformationRequest,
) (*pb.ReportSentinelInformationResponse, error) {
	err := s.g.UpsertSentinelInfo(ctx, converter.ToBizSentinelInfo(req))
	if err != nil {
		return nil, err
	}
	return &pb.ReportSentinelInformationResponse{}, nil
}

func (s *LibrarianSentinelService) ReportAppBinaries(
	ctx context.Context, req *pb.ReportAppBinariesRequest,
) (*pb.ReportAppBinariesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportAppBinaries not implemented")
}
