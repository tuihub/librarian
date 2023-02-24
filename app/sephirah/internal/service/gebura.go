package service

import (
	"context"
	"io"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/service/converter"
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
	err := s.g.UpdateApp(ctx, s.converter.ToBizApp(req.GetApp()))
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
		s.converter.ToBizAppSourceList(req.GetSourceFilter()),
		s.converter.ToBizAppTypeList(req.GetTypeFilter()),
		s.converter.ToBizInternalIDList(req.GetIdFilter()),
		req.GetContainDetails())
	if err != nil {
		return nil, err
	}
	return &pb.ListAppResponse{ // TODO
		Paging:  nil,
		AppList: s.converter.ToPBAppList(a),
	}, nil
}
func (s *LibrarianSephirahServiceService) BindApp(ctx context.Context, req *pb.BindAppRequest) (
	*pb.BindAppResponse, error,
) {
	a, err := s.g.BindApp(ctx, // TODO
		bizgebura.App{
			InternalID:       req.GetInternalAppId().GetId(),
			Source:           0,
			SourceAppID:      "",
			SourceURL:        "",
			Name:             "",
			Type:             0,
			ShortDescription: "",
			ImageURL:         "",
			Details:          nil,
		},
		bizgebura.App{
			InternalID:       0,
			Source:           converter.ToBizAppSource(req.GetBindAppId().GetSource()),
			SourceAppID:      req.GetBindAppId().GetSourceAppId(),
			SourceURL:        "",
			Name:             "",
			Type:             0,
			ShortDescription: "",
			ImageURL:         "",
			Details:          nil,
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
func (s *LibrarianSephirahServiceService) ListBindApp(ctx context.Context, req *pb.ListBindAppRequest) (
	*pb.ListBindAppResponse, error,
) {
	al, err := s.g.ListBindApp(ctx, req.GetAppId().GetId())
	if err != nil {
		return nil, err
	}
	return &pb.ListBindAppResponse{AppList: s.converter.ToPBAppList(al)}, nil
}

func (s *LibrarianSephirahServiceService) CreateAppPackage(
	ctx context.Context,
	req *pb.CreateAppPackageRequest,
) (*pb.CreateAppPackageResponse, error) {
	ap, err := s.g.CreateAppPackage(ctx, s.converter.ToBizAppPackage(req.GetAppPackage()))
	if err != nil {
		return nil, err
	}
	return &pb.CreateAppPackageResponse{Id: &librarian.InternalID{Id: ap.InternalID}}, nil
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
func (s *LibrarianSephirahServiceService) ListAppPackage(
	ctx context.Context,
	req *pb.ListAppPackageRequest,
) (*pb.ListAppPackageResponse, error) {
	ap, err := s.g.ListAppPackage(ctx,
		bizgebura.Paging{
			PageSize: int(req.GetPaging().GetPageSize()),
			PageNum:  int(req.GetPaging().GetPageNum()),
		},
		s.converter.ToBizAppPackageSourceList(req.GetSourceFilter()),
		s.converter.ToBizInternalIDList(req.GetIdFilter()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.ListAppPackageResponse{
		Paging:         nil,
		AppPackageList: s.converter.ToPBAppPackageList(ap),
	}, nil
}
func (s *LibrarianSephirahServiceService) BindAppPackage(
	ctx context.Context,
	req *pb.AssignAppPackageRequest,
) (*pb.AssignAppPackageResponse, error) {
	err := s.g.AssignAppPackage(ctx, bizgebura.App{ // TODO
		InternalID:       req.GetAppId().GetId(),
		Source:           0,
		SourceAppID:      "",
		SourceURL:        "",
		Name:             "",
		Type:             0,
		ShortDescription: "",
		ImageURL:         "",
		Details:          nil,
	}, bizgebura.AppPackage{
		InternalID:      req.GetAppPackageId().GetId(),
		Source:          0,
		SourceID:        0,
		SourcePackageID: "",
		Name:            "",
		Description:     "",
		Binary: &bizgebura.AppPackageBinary{
			Name:      "",
			Size:      0,
			PublicURL: "",
		},
	})
	if err != nil {
		return nil, err
	}
	return &pb.AssignAppPackageResponse{}, nil
}
func (s *LibrarianSephirahServiceService) UnBindAppPackage(
	ctx context.Context,
	req *pb.UnAssignAppPackageRequest,
) (*pb.UnAssignAppPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnBindAppPackage not implemented")
}
func (s *LibrarianSephirahServiceService) ReportAppPackage(
	conn pb.LibrarianSephirahService_ReportAppPackageServer,
) error {
	handler, err0 := s.g.NewReportAppPackageHandler(conn.Context())
	if err0 != nil {
		return err0
	}
	for {
		var apl []*bizgebura.AppPackage
		if req, err := conn.Recv(); err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		} else {
			for id, a := range req.GetAppPackageList() {
				apl = append(apl, &bizgebura.AppPackage{ // TODO
					InternalID:      0,
					Source:          0,
					SourceID:        0,
					SourcePackageID: id,
					Name:            "",
					Description:     "",
					Binary: &bizgebura.AppPackageBinary{
						Name:      a.GetName(),
						Size:      a.GetSize(),
						PublicURL: a.GetPublicUrl(),
					},
				})
			}
		}
		if err := handler.Handle(conn.Context(), apl); err != nil {
			return err
		}
		if err := conn.Send(&pb.ReportAppPackageResponse{}); err != nil {
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
