package sentinel

import (
	"context"
	"time"

	"github.com/tuihub/librarian/internal/service/sentinel/converter"
	pb "github.com/tuihub/protos/pkg/librarian/sentinel/v1"

	"github.com/samber/lo"
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
	var snapshot *time.Time
	if req.GetSnapshotTime() != nil {
		snapshot = lo.ToPtr(req.GetSnapshotTime().AsTime())
	}
	success, err := s.g.UpsertAppBinaries(
		ctx,
		converter.ToBizSentinelAppBinaryList(req.GetAppBinaries()),
		snapshot,
		req.GetCommitSnapshot(),
	)
	if err != nil {
		return nil, err
	}
	return &pb.ReportAppBinariesResponse{
		CommitSnapshotSuccess: lo.ToPtr(success),
	}, nil
}
