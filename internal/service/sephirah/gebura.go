package sephirah

import (
	"context"

	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/service/sephirah/converter"
	sephirah "github.com/tuihub/protos/pkg/librarian/sephirah/v1/sephirah"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//
// func (s *LibrarianSephirahService) ListAppInfos(ctx context.Context, req *sephirah.ListAppInfosRequest) (
//
//	*sephirah.ListAppInfosResponse, error,
//
//	) {
//		a, total, err := s.g.ListAppInfos(ctx,
//			model.ToBizPaging(req.GetPaging()),
//			req.GetSourceFilter(),
//			converter.ToBizAppTypeList(req.GetTypeFilter()),
//			converter.ToBizInternalIDList(req.GetIdFilter()),
//			req.GetContainDetails(),
//		)
//		if err != nil {
//			return nil, err
//		}
//		return &sephirah.ListAppInfosResponse{
//			Paging:   &librarian.PagingResponse{TotalSize: total},
//			AppInfos: converter.ToPBAppInfoList(a),
//		}, nil
//	}
//
//
// func (s *LibrarianSephirahService) SyncAppInfos(ctx context.Context, req *sephirah.SyncAppInfosRequest) (
//
//	*sephirah.SyncAppInfosResponse, error,
//
//	) {
//		apps, err := s.g.SyncAppInfos(ctx, converter.ToBizAppInfoIDList(req.GetAppInfoIds()), req.GetWaitData())
//		if err != nil {
//			return nil, err
//		}
//		return &sephirah.SyncAppInfosResponse{AppInfos: converter.ToPBAppInfoList(apps)}, nil
//	}
// func (s *LibrarianSephirahService) SearchAppInfos(ctx context.Context, req *sephirah.SearchAppInfosRequest) (
//	*sephirah.SearchAppInfosResponse, error,
// ) {
//	infos, total, err := s.g.SearchAppInfos(ctx,
//		model.ToBizPaging(req.GetPaging()),
//		req.GetNameLike(),
//	)
//	if err != nil {
//		return nil, err
//	}
//	return &sephirah.SearchAppInfosResponse{
//		Paging:   &librarian.PagingResponse{TotalSize: int64(total)},
//		AppInfos: converter.ToPBAppInfoMixedList(infos),
//	}, nil
//}

// func (s *LibrarianSephirahService) SearchNewAppInfos(ctx context.Context, req *sephirah.SearchNewAppInfosRequest) (
//	*sephirah.SearchNewAppInfosResponse, error,
// ) {
//	infos, total, err := s.g.SearchNewAppInfos(ctx,
//		model.ToBizPaging(req.GetPaging()),
//		req.GetName(),
//		req.GetSourceFilter(),
//	)
//	if err != nil {
//		return nil, err
//	}
//	return &sephirah.SearchNewAppInfosResponse{
//		Paging:   &librarian.PagingResponse{TotalSize: int64(total)},
//		AppInfos: converter.ToPBAppInfoList(infos),
//	}, nil
//}
// func (s *LibrarianSephirahService) GetAppInfo(ctx context.Context, req *sephirah.GetAppInfoRequest) (
//	*sephirah.GetAppInfoResponse, error,
// ) {
//	res, err := s.g.GetAppInfo(ctx, converter.ToBizInternalID(req.GetAppInfoId()))
//	if err != nil {
//		return nil, err
//	}
//	return &sephirah.GetAppInfoResponse{AppInfo: converter.ToPBAppInfo(res)}, nil
//}
// func (s *LibrarianSephirahService) GetBoundAppInfos(ctx context.Context, req *sephirah.GetBoundAppInfosRequest) (
//	*sephirah.GetBoundAppInfosResponse, error,
// ) {
//	al, err := s.g.GetBoundAppInfos(ctx, converter.ToBizInternalID(req.GetAppInfoId()))
//	if err != nil {
//		return nil, err
//	}
//	return &sephirah.GetBoundAppInfosResponse{AppInfos: converter.ToPBAppInfoList(al)}, nil
//}
// func (s *LibrarianSephirahService) PurchaseAppInfo(ctx context.Context, req *sephirah.PurchaseAppInfoRequest) (
//	*sephirah.PurchaseAppInfoResponse, error,
// ) {
//	id, err := s.g.PurchaseAppInfo(ctx, converter.ToBizAppInfoID(req.GetAppInfoId()))
//	if err != nil {
//		return nil, err
//	}
//	return &sephirah.PurchaseAppInfoResponse{
//		Id: converter.ToPBInternalID(id),
//	}, nil
//}
// func (s *LibrarianSephirahService) GetPurchasedAppInfos(
//	ctx context.Context,
//	req *sephirah.GetPurchasedAppInfosRequest,
// ) (
//	*sephirah.GetPurchasedAppInfosResponse, error,
// ) {
//	infos, err := s.g.GetPurchasedAppInfos(ctx, req.GetSource())
//	if err != nil {
//		return nil, err
//	}
//	return &sephirah.GetPurchasedAppInfosResponse{
//		AppInfos: converter.ToPBAppInfoMixedList(infos),
//	}, nil
//}

func (s *LibrarianSephirahService) CreateApp(
	ctx context.Context,
	req *sephirah.CreateAppRequest,
) (*sephirah.CreateAppResponse, error) {
	ap, err := s.g.CreateApp(ctx, converter.ToBizApp(req.GetApp()))
	if err != nil {
		return nil, err
	}
	return &sephirah.CreateAppResponse{Id: converter.ToPBInternalID(ap.ID)}, nil
}
func (s *LibrarianSephirahService) UpdateApp(
	ctx context.Context,
	req *sephirah.UpdateAppRequest,
) (*sephirah.UpdateAppResponse, error) {
	err := s.g.UpdateApp(ctx, converter.ToBizApp(req.GetApp()))
	if err != nil {
		return nil, err
	}
	return &sephirah.UpdateAppResponse{}, nil
}
func (s *LibrarianSephirahService) ListApps(
	ctx context.Context,
	req *sephirah.ListAppsRequest,
) (*sephirah.ListAppsResponse, error) {
	ap, total, err := s.g.ListApps(ctx,
		model.ToBizPaging(req.GetPaging()),
		converter.ToBizInternalIDList(req.GetOwnerIdFilter()),
		converter.ToBizInternalIDList(req.GetIdFilter()),
	)
	if err != nil {
		return nil, err
	}
	return &sephirah.ListAppsResponse{
		Paging: &librarian.PagingResponse{TotalSize: int64(total)},
		Apps:   converter.ToPBAppList(ap),
	}, nil
}

// func (s *LibrarianSephirahService) AssignApp(
//	ctx context.Context,
//	req *sephirah.AssignAppRequest,
// ) (*sephirah.AssignAppResponse, error) {
//	err := s.g.AssignApp(ctx,
//		converter.ToBizInternalID(req.GetAppId()),
//		converter.ToBizInternalID(req.GetAppInfoId()),
//	)
//	if err != nil {
//		return nil, err
//	}
//	return &sephirah.AssignAppResponse{}, nil
//}
// func (s *LibrarianSephirahService) UnAssignApp(
//	ctx context.Context,
//	req *sephirah.UnAssignAppRequest,
// ) (*sephirah.UnAssignAppResponse, error) {
//	err := s.g.UnAssignApp(ctx, converter.ToBizInternalID(req.GetAppId()))
//	if err != nil {
//		return nil, err
//	}
//	return &sephirah.UnAssignAppResponse{}, nil
//}

// func (s *LibrarianSephirahService) ReportAppPackages(
//
//	conn sephirah.LibrarianSephirahService_ReportAppPackagesServer,
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
//			if err := conn.Send(&sephirah.ReportAppPackagesResponse{}); err != nil {
//				return err
//			}
//		}
//	}

func (s *LibrarianSephirahService) BatchCreateAppRunTime(
	ctx context.Context,
	req *sephirah.BatchCreateAppRunTimeRequest,
) (*sephirah.BatchCreateAppRunTimeResponse, error) {
	err := s.g.BatchCreateAppRunTime(ctx,
		converter.ToBizAppRunTimeList(req.GetAppRunTimes()),
	)
	if err != nil {
		return nil, err
	}
	return &sephirah.BatchCreateAppRunTimeResponse{}, nil
}
func (s *LibrarianSephirahService) SumAppRunTime(
	ctx context.Context,
	req *sephirah.SumAppRunTimeRequest,
) (*sephirah.SumAppRunTimeResponse, error) {
	res, err := s.g.SumAppRunTime(ctx,
		converter.ToBizInternalIDList(req.GetAppIdFilter()),
		converter.ToBizInternalIDList(req.GetDeviceIdFilter()),
		converter.ToBizTimeRange(req.GetTimeRangeCross()),
	)
	if err != nil {
		return nil, err
	}
	return &sephirah.SumAppRunTimeResponse{
		RunTimeSum: converter.ToPBDuration(lo.FromPtr(res)),
	}, nil
}
func (s *LibrarianSephirahService) ListAppRunTimes(
	ctx context.Context,
	req *sephirah.ListAppRunTimesRequest,
) (*sephirah.ListAppRunTimesResponse, error) {
	res, total, err := s.g.ListAppRunTimes(ctx,
		model.ToBizPaging(req.GetPaging()),
		converter.ToBizInternalIDList(req.GetAppIdFilter()),
		converter.ToBizInternalIDList(req.GetDeviceIdFilter()),
		converter.ToBizTimeRange(req.GetTimeRangeCross()),
	)
	if err != nil {
		return nil, err
	}
	return &sephirah.ListAppRunTimesResponse{
		Paging:      &librarian.PagingResponse{TotalSize: int64(total)},
		AppRunTimes: converter.ToPBAppRunTimeList(res),
	}, nil
}
func (s *LibrarianSephirahService) DeleteAppRunTime(
	ctx context.Context,
	req *sephirah.DeleteAppRunTimeRequest,
) (*sephirah.DeleteAppRunTimeResponse, error) {
	err := s.g.DeleteAppRunTime(ctx, converter.ToBizInternalID(req.GetId()))
	if err != nil {
		return nil, err
	}
	return &sephirah.DeleteAppRunTimeResponse{}, nil
}

func (s *LibrarianSephirahService) UploadAppSaveFile(
	ctx context.Context,
	req *sephirah.UploadAppSaveFileRequest,
) (*sephirah.UploadAppSaveFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadGameSaveFile not implemented")
}
func (s *LibrarianSephirahService) DownloadAppSaveFile(
	ctx context.Context,
	req *sephirah.DownloadAppSaveFileRequest,
) (*sephirah.DownloadAppSaveFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadGameSaveFile not implemented")
}
func (s *LibrarianSephirahService) ListAppSaveFiles(
	ctx context.Context,
	req *sephirah.ListAppSaveFilesRequest,
) (*sephirah.ListAppSaveFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGameSaveFile not implemented")
}

func (s *LibrarianSephirahService) ListAppCategories(
	ctx context.Context,
	req *sephirah.ListAppCategoriesRequest,
) (*sephirah.ListAppCategoriesResponse, error) {
	aps, err := s.g.ListAppCategories(ctx)
	if err != nil {
		return nil, err
	}
	return &sephirah.ListAppCategoriesResponse{
		AppCategories: converter.ToPBAppCategoryList(aps),
	}, nil
}
func (s *LibrarianSephirahService) CreateAppCategory(
	ctx context.Context,
	req *sephirah.CreateAppCategoryRequest,
) (*sephirah.CreateAppCategoryResponse, error) {
	ac, err := s.g.CreateAppCategory(ctx, converter.ToBizAppCategory(req.GetAppCategory()))
	if err != nil {
		return nil, err
	}
	return &sephirah.CreateAppCategoryResponse{Id: converter.ToPBInternalID(ac.ID)}, nil
}
func (s *LibrarianSephirahService) UpdateAppCategory(
	ctx context.Context,
	req *sephirah.UpdateAppCategoryRequest,
) (*sephirah.UpdateAppCategoryResponse, error) {
	err := s.g.UpdateAppCategory(ctx, converter.ToBizAppCategory(req.GetAppCategory()))
	if err != nil {
		return nil, err
	}
	return &sephirah.UpdateAppCategoryResponse{}, nil
}
func (s *LibrarianSephirahService) DeleteAppCategory(
	ctx context.Context,
	req *sephirah.DeleteAppCategoryRequest,
) (*sephirah.DeleteAppCategoryResponse, error) {
	err := s.g.DeleteAppCategory(ctx, converter.ToBizInternalID(req.GetId()))
	if err != nil {
		return nil, err
	}
	return &sephirah.DeleteAppCategoryResponse{}, nil
}
