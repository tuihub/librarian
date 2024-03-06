package service

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/model/converter"
	"github.com/tuihub/librarian/internal/model"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LibrarianSephirahServiceService) CreateAppInfo(ctx context.Context, req *pb.CreateAppInfoRequest) (
	*pb.CreateAppInfoResponse, error,
) {
	appInfo := req.GetAppInfo()
	if appInfo == nil {
		return nil, pb.ErrorErrorReasonBadRequest("appInfo info required")
	}
	a, err := s.g.CreateAppInfo(ctx, converter.ToBizAppInfo(req.GetAppInfo()))
	if err != nil {
		return nil, err
	}
	return &pb.CreateAppInfoResponse{
		Id: converter.ToPBInternalID(a.ID),
	}, nil
}
func (s *LibrarianSephirahServiceService) UpdateAppInfo(ctx context.Context, req *pb.UpdateAppInfoRequest) (
	*pb.UpdateAppInfoResponse, error,
) {
	appInfo := req.GetAppInfo()
	if appInfo == nil || appInfo.GetId() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("appInfo and internal_id required")
	}
	err := s.g.UpdateAppInfo(ctx, converter.ToBizAppInfo(req.GetAppInfo()))
	if err != nil {
		return nil, err
	}
	return &pb.UpdateAppInfoResponse{}, nil
}
func (s *LibrarianSephirahServiceService) ListAppInfos(ctx context.Context, req *pb.ListAppInfosRequest) (
	*pb.ListAppInfosResponse, error,
) {
	a, total, err := s.g.ListAppInfos(ctx,
		model.ToBizPaging(req.GetPaging()),
		req.GetSourceFilter(),
		converter.ToBizAppTypeList(req.GetTypeFilter()),
		converter.ToBizInternalIDList(req.GetIdFilter()),
		req.GetContainDetails(),
	)
	if err != nil {
		return nil, err
	}
	return &pb.ListAppInfosResponse{
		Paging:   &librarian.PagingResponse{TotalSize: total},
		AppInfos: converter.ToPBAppInfoList(a),
	}, nil
}
func (s *LibrarianSephirahServiceService) MergeAppInfos(ctx context.Context, req *pb.MergeAppInfosRequest) (
	*pb.MergeAppInfosResponse, error,
) {
	info := converter.ToBizAppInfo(req.GetBase())
	if info == nil {
		return nil, pb.ErrorErrorReasonBadRequest("base required")
	}
	err := s.g.MergeAppInfos(ctx,
		*info,
		converter.ToBizInternalID(req.GetMerged()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.MergeAppInfosResponse{}, nil
}
func (s *LibrarianSephirahServiceService) SyncAppInfos(ctx context.Context, req *pb.SyncAppInfosRequest) (
	*pb.SyncAppInfosResponse, error,
) {
	apps, err := s.g.SyncAppInfos(ctx, converter.ToBizAppInfoIDList(req.GetAppInfoIds()), req.GetWaitData())
	if err != nil {
		return nil, err
	}
	return &pb.SyncAppInfosResponse{AppInfos: converter.ToPBAppInfoList(apps)}, nil
}
func (s *LibrarianSephirahServiceService) SearchAppInfos(ctx context.Context, req *pb.SearchAppInfosRequest) (
	*pb.SearchAppInfosResponse, error,
) {
	infos, total, err := s.g.SearchAppInfos(ctx,
		model.ToBizPaging(req.GetPaging()),
		req.GetQuery(),
	)
	if err != nil {
		return nil, err
	}
	return &pb.SearchAppInfosResponse{
		Paging:   &librarian.PagingResponse{TotalSize: int64(total)},
		AppInfos: converter.ToPBAppInfoMixedList(infos),
	}, nil
}
func (s *LibrarianSephirahServiceService) SearchNewAppInfos(ctx context.Context, req *pb.SearchNewAppInfosRequest) (
	*pb.SearchNewAppInfosResponse, error,
) {
	infos, total, err := s.g.SearchNewAppInfos(ctx,
		model.ToBizPaging(req.GetPaging()),
		req.GetName(),
		req.GetSourceFilter(),
	)
	if err != nil {
		return nil, err
	}
	return &pb.SearchNewAppInfosResponse{
		Paging:   &librarian.PagingResponse{TotalSize: int64(total)},
		AppInfos: converter.ToPBAppInfoList(infos),
	}, nil
}
func (s *LibrarianSephirahServiceService) GetAppInfo(ctx context.Context, req *pb.GetAppInfoRequest) (
	*pb.GetAppInfoResponse, error,
) {
	res, err := s.g.GetAppInfo(ctx, converter.ToBizInternalID(req.GetAppInfoId()))
	if err != nil {
		return nil, err
	}
	return &pb.GetAppInfoResponse{AppInfo: converter.ToPBAppInfo(res)}, nil
}
func (s *LibrarianSephirahServiceService) GetBoundAppInfos(ctx context.Context, req *pb.GetBoundAppInfosRequest) (
	*pb.GetBoundAppInfosResponse, error,
) {
	al, err := s.g.GetBoundAppInfos(ctx, converter.ToBizInternalID(req.GetAppInfoId()))
	if err != nil {
		return nil, err
	}
	return &pb.GetBoundAppInfosResponse{AppInfos: converter.ToPBAppInfoList(al)}, nil
}
func (s *LibrarianSephirahServiceService) PurchaseAppInfo(ctx context.Context, req *pb.PurchaseAppInfoRequest) (
	*pb.PurchaseAppInfoResponse, error,
) {
	id, err := s.g.PurchaseAppInfo(ctx, converter.ToBizAppInfoID(req.GetAppInfoId()))
	if err != nil {
		return nil, err
	}
	return &pb.PurchaseAppInfoResponse{
		Id: converter.ToPBInternalID(id),
	}, nil
}
func (s *LibrarianSephirahServiceService) GetPurchasedAppInfos(
	ctx context.Context,
	req *pb.GetPurchasedAppInfosRequest,
) (
	*pb.GetPurchasedAppInfosResponse, error,
) {
	infos, err := s.g.GetPurchasedAppInfos(ctx, req.GetSource())
	if err != nil {
		return nil, err
	}
	return &pb.GetPurchasedAppInfosResponse{
		AppInfos: converter.ToPBAppInfoMixedList(infos),
	}, nil
}

func (s *LibrarianSephirahServiceService) CreateApp(
	ctx context.Context,
	req *pb.CreateAppRequest,
) (*pb.CreateAppResponse, error) {
	ap, err := s.g.CreateApp(ctx, converter.ToBizApp(req.GetApp()))
	if err != nil {
		return nil, err
	}
	return &pb.CreateAppResponse{Id: converter.ToPBInternalID(ap.ID)}, nil
}
func (s *LibrarianSephirahServiceService) UpdateApp(
	ctx context.Context,
	req *pb.UpdateAppRequest,
) (*pb.UpdateAppResponse, error) {
	err := s.g.UpdateApp(ctx, converter.ToBizApp(req.GetApp()))
	if err != nil {
		return nil, err
	}
	return &pb.UpdateAppResponse{}, nil
}
func (s *LibrarianSephirahServiceService) ListApps(
	ctx context.Context,
	req *pb.ListAppsRequest,
) (*pb.ListAppsResponse, error) {
	ap, total, err := s.g.ListApps(ctx,
		model.ToBizPaging(req.GetPaging()),
		converter.ToBizInternalIDList(req.GetOwnerIdFilter()),
		converter.ToBizInternalIDList(req.GetAssignedAppInfoIdFilter()),
		converter.ToBizInternalIDList(req.GetIdFilter()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.ListAppsResponse{
		Paging: &librarian.PagingResponse{TotalSize: int64(total)},
		Apps:   converter.ToPBAppList(ap),
	}, nil
}
func (s *LibrarianSephirahServiceService) AssignApp(
	ctx context.Context,
	req *pb.AssignAppRequest,
) (*pb.AssignAppResponse, error) {
	err := s.g.AssignApp(ctx,
		converter.ToBizInternalID(req.GetAppId()),
		converter.ToBizInternalID(req.GetAppInfoId()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.AssignAppResponse{}, nil
}
func (s *LibrarianSephirahServiceService) UnAssignApp(
	ctx context.Context,
	req *pb.UnAssignAppRequest,
) (*pb.UnAssignAppResponse, error) {
	err := s.g.UnAssignApp(ctx, converter.ToBizInternalID(req.GetAppId()))
	if err != nil {
		return nil, err
	}
	return &pb.UnAssignAppResponse{}, nil
}

// func (s *LibrarianSephirahServiceService) ReportAppPackages(
//
//	conn pb.LibrarianSephirahService_ReportAppPackagesServer,
//
//	) error {
//		ctx, err1 := s.authFunc(conn.Context())
//		if err1 != nil {
//			return err1
//		}
//		handler, err2 := s.g.NewReportAppPackageHandler(ctx)
//		if err2 != nil {
//			return err2
//		}
//		for {
//			var binaries []*modelgebura.AppPackageBinary
//			if req, err := conn.Recv(); err != nil {
//				if errors.Is(err, io.EOF) {
//					return nil
//				}
//				return err
//			} else {
//				binaries = converter.ToBizAppPackageBinaryList(req.GetSentinelAppPackageBinaries())
//			}
//			if err := handler.Handle(conn.Context(), binaries); err != nil {
//				return err
//			}
//			if err := conn.Send(&pb.ReportAppPackagesResponse{}); err != nil {
//				return err
//			}
//		}
//	}

func (s *LibrarianSephirahServiceService) CreateAppInst(
	ctx context.Context,
	req *pb.CreateAppInstRequest,
) (*pb.CreateAppInstResponse, error) {
	ap, err := s.g.CreateAppInst(ctx, converter.ToBizAppInst(req.GetAppInst()))
	if err != nil {
		return nil, err
	}
	return &pb.CreateAppInstResponse{Id: converter.ToPBInternalID(ap.ID)}, nil
}

func (s *LibrarianSephirahServiceService) UpdateAppInst(
	ctx context.Context,
	req *pb.UpdateAppInstRequest,
) (*pb.UpdateAppInstResponse, error) {
	err := s.g.UpdateAppInst(ctx, converter.ToBizAppInst(req.GetAppInst()))
	if err != nil {
		return nil, err
	}
	return &pb.UpdateAppInstResponse{}, nil
}

func (s *LibrarianSephirahServiceService) ListAppInsts(
	ctx context.Context,
	req *pb.ListAppInstsRequest,
) (*pb.ListAppInstsResponse, error) {
	ap, total, err := s.g.ListAppInsts(ctx,
		model.ToBizPaging(req.GetPaging()),
		converter.ToBizInternalIDList(req.GetIdFilter()),
		converter.ToBizInternalIDList(req.GetAppIdFilter()),
		converter.ToBizInternalIDList(req.GetDeviceIdFilter()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.ListAppInstsResponse{
		Paging:   &librarian.PagingResponse{TotalSize: int64(total)},
		AppInsts: converter.ToPBAppInstList(ap),
	}, nil
}

func (s *LibrarianSephirahServiceService) AddAppInstRunTime(
	ctx context.Context,
	req *pb.AddAppInstRunTimeRequest,
) (*pb.AddAppInstRunTimeResponse, error) {
	err := s.g.AddAppInstRunTime(ctx,
		converter.ToBizInternalID(req.GetAppInstId()),
		converter.ToBizTimeRange(req.GetTimeRange()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.AddAppInstRunTimeResponse{}, nil
}
func (s *LibrarianSephirahServiceService) SumAppInstRunTime(
	ctx context.Context,
	req *pb.SumAppInstRunTimeRequest,
) (*pb.SumAppInstRunTimeResponse, error) {
	if req.GetTimeAggregation().GetAggregationType() != librarian.TimeAggregation_AGGREGATION_TYPE_OVERALL {
		return nil, pb.ErrorErrorReasonBadRequest("unsupported aggregation type")
	}
	res, err := s.g.SumAppInstRunTime(ctx,
		converter.ToBizInternalID(req.GetAppInstId()),
		converter.ToBizTimeRange(req.GetTimeAggregation().GetTimeRange()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.SumAppInstRunTimeResponse{RunTimeGroups: []*pb.SumAppInstRunTimeResponse_Group{{
		TimeRange: req.GetTimeAggregation().GetTimeRange(),
		Duration:  converter.ToPBDuration(res),
	}}}, nil
}

func (s *LibrarianSephirahServiceService) UploadAppSaveFile(
	ctx context.Context,
	req *pb.UploadAppSaveFileRequest,
) (*pb.UploadAppSaveFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadGameSaveFile not implemented")
}
func (s *LibrarianSephirahServiceService) DownloadAppSaveFile(
	ctx context.Context,
	req *pb.DownloadAppSaveFileRequest,
) (*pb.DownloadAppSaveFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadGameSaveFile not implemented")
}
func (s *LibrarianSephirahServiceService) ListAppSaveFiles(
	ctx context.Context,
	req *pb.ListAppSaveFilesRequest,
) (*pb.ListAppSaveFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGameSaveFile not implemented")
}
