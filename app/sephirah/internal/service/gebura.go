package service

import (
	"context"
	"io"

	"github.com/tuihub/librarian/app/sephirah/internal/model/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/model"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"github.com/go-kratos/kratos/v2/errors"
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
	a, err := s.g.CreateApp(ctx, s.converter.ToBizApp(req.GetApp()))
	if err != nil {
		return nil, err
	}
	return &pb.CreateAppResponse{
		Id: converter.ToPBInternalID(a.ID),
	}, nil
}
func (s *LibrarianSephirahServiceService) UpdateApp(ctx context.Context, req *pb.UpdateAppRequest) (
	*pb.UpdateAppResponse, error,
) {
	app := req.GetApp()
	if app == nil || app.GetId() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("app and internal_id required")
	}
	err := s.g.UpdateApp(ctx, s.converter.ToBizApp(req.GetApp()))
	if err != nil {
		return nil, err
	}
	return &pb.UpdateAppResponse{}, nil
}
func (s *LibrarianSephirahServiceService) ListApps(ctx context.Context, req *pb.ListAppsRequest) (
	*pb.ListAppsResponse, error,
) {
	a, total, err := s.g.ListApp(ctx,
		model.Paging{
			PageSize: int(req.GetPaging().GetPageSize()),
			PageNum:  int(req.GetPaging().GetPageNum()),
		},
		s.converter.ToBizAppSourceList(req.GetSourceFilter()),
		s.converter.ToBizAppTypeList(req.GetTypeFilter()),
		s.converter.ToBizInternalIDList(req.GetIdFilter()),
		req.GetContainDetails())
	if err != nil {
		return nil, err
	}
	return &pb.ListAppsResponse{
		Paging: &librarian.PagingResponse{TotalSize: total},
		Apps:   s.converter.ToPBAppList(a),
	}, nil
}
func (s *LibrarianSephirahServiceService) RefreshApp(ctx context.Context, req *pb.RefreshAppRequest) (
	*pb.RefreshAppResponse, error,
) {
	return nil, pb.ErrorErrorReasonNotImplemented("impl in next version")
}
func (s *LibrarianSephirahServiceService) MergeApps(ctx context.Context, req *pb.MergeAppsRequest) (
	*pb.MergeAppsResponse, error,
) {
	return nil, pb.ErrorErrorReasonNotImplemented("impl in next version")
}
func (s *LibrarianSephirahServiceService) SearchApps(ctx context.Context, req *pb.SearchAppsRequest) (
	*pb.SearchAppsResponse, error,
) {
	return nil, pb.ErrorErrorReasonNotImplemented("impl in next version")
}
func (s *LibrarianSephirahServiceService) GetBindApps(ctx context.Context, req *pb.GetBindAppsRequest) (
	*pb.GetBindAppsResponse, error,
) {
	al, err := s.g.GetBindApps(ctx, converter.ToBizInternalID(req.GetAppId()))
	if err != nil {
		return nil, err
	}
	return &pb.GetBindAppsResponse{Apps: s.converter.ToPBAppList(al)}, nil
}
func (s *LibrarianSephirahServiceService) PurchaseApp(ctx context.Context, req *pb.PurchaseAppRequest) (
	*pb.PurchaseAppResponse, error,
) {
	return nil, pb.ErrorErrorReasonNotImplemented("impl in next version")
}
func (s *LibrarianSephirahServiceService) GetAppLibrary(ctx context.Context, req *pb.GetAppLibraryRequest) (
	*pb.GetAppLibraryResponse, error,
) {
	return nil, pb.ErrorErrorReasonNotImplemented("impl in next version")
}

func (s *LibrarianSephirahServiceService) CreateAppPackage(
	ctx context.Context,
	req *pb.CreateAppPackageRequest,
) (*pb.CreateAppPackageResponse, error) {
	ap, err := s.g.CreateAppPackage(ctx, s.converter.ToBizAppPackage(req.GetAppPackage()))
	if err != nil {
		return nil, err
	}
	return &pb.CreateAppPackageResponse{Id: converter.ToPBInternalID(ap.ID)}, nil
}
func (s *LibrarianSephirahServiceService) UpdateAppPackage(
	ctx context.Context,
	req *pb.UpdateAppPackageRequest,
) (*pb.UpdateAppPackageResponse, error) {
	err := s.g.UpdateAppPackage(ctx, s.converter.ToBizAppPackage(req.GetAppPackage()))
	if err == nil {
		return nil, err
	}
	return &pb.UpdateAppPackageResponse{}, nil
}
func (s *LibrarianSephirahServiceService) ListAppPackages(
	ctx context.Context,
	req *pb.ListAppPackagesRequest,
) (*pb.ListAppPackagesResponse, error) {
	ap, err := s.g.ListAppPackage(ctx,
		model.Paging{
			PageSize: int(req.GetPaging().GetPageSize()),
			PageNum:  int(req.GetPaging().GetPageNum()),
		},
		s.converter.ToBizAppPackageSourceList(req.GetSourceFilter()),
		s.converter.ToBizInternalIDList(req.GetIdFilter()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.ListAppPackagesResponse{
		Paging:      nil,
		AppPackages: s.converter.ToPBAppPackageList(ap),
	}, nil
}
func (s *LibrarianSephirahServiceService) AssignAppPackage(
	ctx context.Context,
	req *pb.AssignAppPackageRequest,
) (*pb.AssignAppPackageResponse, error) {
	err := s.g.AssignAppPackage(ctx, modelgebura.App{ // TODO
		ID:               converter.ToBizInternalID(req.GetAppId()),
		Source:           0,
		SourceAppID:      "",
		SourceURL:        "",
		Name:             "",
		Type:             0,
		ShortDescription: "",
		ImageURL:         "",
		Details:          nil,
	}, modelgebura.AppPackage{
		ID:              converter.ToBizInternalID(req.GetAppPackageId()),
		Source:          0,
		SourceID:        0,
		SourcePackageID: "",
		Name:            "",
		Description:     "",
		Binary: &modelgebura.AppPackageBinary{
			Name:      "",
			SizeByte:  0,
			PublicURL: "",
		},
		Public: false,
	})
	if err != nil {
		return nil, err
	}
	return &pb.AssignAppPackageResponse{}, nil
}
func (s *LibrarianSephirahServiceService) UnAssignAppPackage(
	ctx context.Context,
	req *pb.UnAssignAppPackageRequest,
) (*pb.UnAssignAppPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnBindAppPackage not implemented")
}
func (s *LibrarianSephirahServiceService) ReportAppPackages(
	conn pb.LibrarianSephirahService_ReportAppPackagesServer,
) error {
	handler, err0 := s.g.NewReportAppPackageHandler(conn.Context())
	if err0 != nil {
		return err0
	}
	for {
		var apl []*modelgebura.AppPackage
		if req, err := conn.Recv(); err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		} else {
			for id, a := range req.GetAppPackages() {
				apl = append(apl, &modelgebura.AppPackage{ // TODO
					ID:              0,
					Source:          0,
					SourceID:        0,
					SourcePackageID: id,
					Name:            "",
					Description:     "",
					Binary: &modelgebura.AppPackageBinary{
						Name:      a.GetName(),
						SizeByte:  a.GetSizeByte(),
						PublicURL: a.GetPublicUrl(),
					},
					Public: false,
				})
			}
		}
		if err := handler.Handle(conn.Context(), apl); err != nil {
			return err
		}
		if err := conn.Send(&pb.ReportAppPackagesResponse{}); err != nil {
			return err
		}
	}
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
