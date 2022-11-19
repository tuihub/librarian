package service

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LibrarianSephirahServiceService) CreateApp(ctx context.Context, req *pb.CreateAppRequest) (
	*pb.CreateAppResponse, error,
) {
	app := req.GetApp()
	if app == nil {
		return nil, pb.ErrorErrorReasonBadRequest("app required")
	}
	a, err := s.g.CreateApp(ctx, &bizgebura.App{
		Source:          bizgebura.AppSourceInternal,
		SourceAppID:     "",
		SourceURL:       "",
		Name:            app.GetName(),
		Type:            bizgebura.ToBizAppType(app.GetType()),
		ShorDescription: app.GetShortDescription(),
		ImageURL:        app.GetImageUrl(),
		Details:         bizgebura.ToBizAppDetail(app.GetDetails()),
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateAppResponse{
		Id: &librarian.InternalID{Id: a.InternalID},
	}, nil
}
func (s *LibrarianSephirahServiceService) UpdateApp(ctx context.Context, req *pb.UpdateAppRequest) (
	*pb.UpdateAppResponse, error,
) {
	app := req.GetApp()
	if app == nil || app.GetId() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("app and internal_id required")
	}
	err := s.g.UpdateApp(ctx, &bizgebura.App{
		InternalID:      app.GetId().GetId(),
		Source:          bizgebura.AppSourceInternal,
		SourceAppID:     "",
		SourceURL:       "",
		Name:            app.GetName(),
		Type:            bizgebura.ToBizAppType(app.GetType()),
		ShorDescription: app.GetShortDescription(),
		ImageURL:        app.GetImageUrl(),
		Details:         bizgebura.ToBizAppDetail(app.GetDetails()),
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateAppResponse{}, nil
}
func (s *LibrarianSephirahServiceService) ListApp(ctx context.Context, req *pb.ListAppRequest) (
	*pb.ListAppResponse, error,
) {
	a, err := s.g.ListApp(ctx,
		bizgebura.Paging{
			PageSize: int(req.GetPaging().GetPageSize()),
			PageNum:  int(req.GetPaging().GetPageNum()),
		},
		bizgebura.ToBizAppSourceList(req.GetSourceFilter()),
		bizgebura.ToBizAppTypeList(req.GetTypeFilter()),
		toBizInternalIDList(req.GetIdFilter()),
		req.GetContainDetails())
	if err != nil {
		return nil, err
	}
	return &pb.ListAppResponse{
		AppList: bizgebura.ToPBAppList(a, req.GetContainDetails()),
	}, nil
}
func (s *LibrarianSephirahServiceService) BindApp(ctx context.Context, req *pb.BindAppRequest) (
	*pb.BindAppResponse, error,
) {
	a, err := s.g.BindApp(ctx,
		bizgebura.App{
			InternalID: req.GetInternalAppId().GetId(),
		},
		bizgebura.App{
			Source:      bizgebura.ToBizAppSource(req.GetBindAppId().GetSource()),
			SourceAppID: req.GetBindAppId().GetSourceAppId(),
		})
	if err != nil {
		return nil, err
	}
	return &pb.BindAppResponse{BindAppId: &librarian.InternalID{
		Id: a.InternalID,
	}}, nil
}
func (s *LibrarianSephirahServiceService) UnBindApp(ctx context.Context, req *pb.UnBindAppRequest) (
	*pb.UnBindAppResponse, error,
) {
	return nil, pb.ErrorErrorReasonNotImplemented("impl in next version")
}
func (s *LibrarianSephirahServiceService) RefreshApp(ctx context.Context, req *pb.RefreshAppRequest) (
	*pb.RefreshAppResponse, error,
) {
	return nil, pb.ErrorErrorReasonNotImplemented("impl in next version")
}

func (s *LibrarianSephirahServiceService) CreateAppPackage(
	ctx context.Context,
	req *pb.CreateAppPackageRequest,
) (*pb.CreateAppPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppPackage not implemented")
}
func (s *LibrarianSephirahServiceService) UpdateAppPackage(
	ctx context.Context,
	req *pb.UpdateAppPackageRequest,
) (*pb.UpdateAppPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppPackage not implemented")
}
func (s *LibrarianSephirahServiceService) ListAppPackage(
	ctx context.Context,
	req *pb.ListAppPackageRequest,
) (*pb.ListAppPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAppPackage not implemented")
}
func (s *LibrarianSephirahServiceService) BindAppPackage(
	ctx context.Context,
	req *pb.BindAppPackageRequest,
) (*pb.BindAppPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindAppPackage not implemented")
}
func (s *LibrarianSephirahServiceService) UnBindAppPackage(
	ctx context.Context,
	req *pb.UnBindAppPackageRequest,
) (*pb.UnBindAppPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnBindAppPackage not implemented")
}
func (s *LibrarianSephirahServiceService) ReportAppPackage(
	ctx context.Context,
	req *pb.ReportAppPackageRequest,
) (*pb.ReportAppPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportAppPackage not implemented")
}
func (s *LibrarianSephirahServiceService) UploadGameSaveFile(
	ctx context.Context,
	req *pb.UploadGameSaveFileRequest,
) (*pb.UploadGameSaveFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadGameSaveFile not implemented")
}
func (s *LibrarianSephirahServiceService) DownloadGameSaveFile(
	ctx context.Context,
	req *pb.DownloadGameSaveFileRequest,
) (*pb.DownloadGameSaveFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadGameSaveFile not implemented")
}
func (s *LibrarianSephirahServiceService) ListGameSaveFile(
	ctx context.Context,
	req *pb.ListGameSaveFileRequest,
) (*pb.ListGameSaveFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGameSaveFile not implemented")
}
