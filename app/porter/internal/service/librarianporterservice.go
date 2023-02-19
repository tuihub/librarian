package service

import (
	"context"
	"io"
	"strconv"

	"github.com/tuihub/librarian/app/porter/internal/biz/bizs3"
	"github.com/tuihub/librarian/app/porter/internal/biz/bizsteam"
	pb "github.com/tuihub/protos/pkg/librarian/porter/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LibrarianPorterServiceService struct {
	pb.UnimplementedLibrarianPorterServiceServer

	uc *bizsteam.SteamUseCase
	s3 *bizs3.S3
}

func NewLibrarianPorterServiceService(uc *bizsteam.SteamUseCase, s3 *bizs3.S3) pb.LibrarianPorterServiceServer {
	return &LibrarianPorterServiceService{
		UnimplementedLibrarianPorterServiceServer: pb.UnimplementedLibrarianPorterServiceServer{},
		uc: uc,
		s3: s3,
	}
}

func (s *LibrarianPorterServiceService) PullFeed(
	ctx context.Context,
	req *pb.PullFeedRequest,
) (*pb.PullFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullFeed not implemented")
}

func (s *LibrarianPorterServiceService) PullAccount(
	ctx context.Context,
	req *pb.PullAccountRequest,
) (*pb.PullAccountResponse, error) {
	switch req.GetAccountId().GetPlatform() {
	case librarian.AccountPlatform_ACCOUNT_PLATFORM_STEAM:
		u, err := s.uc.GetUser(ctx, req.GetAccountId().GetPlatformAccountId())
		if err != nil {
			return nil, err
		}
		return &pb.PullAccountResponse{Account: &librarian.Account{
			Id:                nil,
			Platform:          req.GetAccountId().GetPlatform(),
			PlatformAccountId: req.GetAccountId().GetPlatformAccountId(),
			Name:              u.Name,
			ProfileUrl:        u.ProfileURL,
			AvatarUrl:         u.AvatarURL,
		}}, nil
	default:
		return nil, status.Errorf(codes.InvalidArgument, "platform unexpected")
	}
}
func (s *LibrarianPorterServiceService) PullApp(
	ctx context.Context,
	req *pb.PullAppRequest,
) (*pb.PullAppResponse, error) {
	switch req.GetAppId().GetSource() {
	case librarian.AppSource_APP_SOURCE_STEAM:
		appID, err := strconv.ParseInt(req.GetAppId().GetSourceAppId(), 10, 64)
		if err != nil {
			return nil, err
		}
		a, err := s.uc.GetAppDetails(ctx, int(appID))
		if err != nil {
			return nil, err
		}
		return &pb.PullAppResponse{App: &librarian.App{
			Id:               nil,
			Source:           req.GetAppId().GetSource(),
			SourceAppId:      req.GetAppId().GetSourceAppId(),
			SourceUrl:        &a.StoreURL,
			Name:             a.Name,
			Type:             toPBAppType(a.Type),
			ShortDescription: a.ShortDescription,
			ImageUrl:         a.ImageURL,
			Details: &librarian.AppDetails{ // TODO
				Description: a.Description,
				ReleaseDate: a.ReleaseDate,
				Developer:   a.Developer,
				Publisher:   a.Publisher,
				Version:     "",
			},
		}}, nil
	default:
		return nil, status.Errorf(codes.InvalidArgument, "source unexpected")
	}
}
func (s *LibrarianPorterServiceService) PullAccountAppRelation(
	ctx context.Context,
	req *pb.PullAccountAppRelationRequest,
) (*pb.PullAccountAppRelationResponse, error) {
	switch req.GetAccountId().GetPlatform() {
	case librarian.AccountPlatform_ACCOUNT_PLATFORM_STEAM:
		al, err := s.uc.GetOwnedGames(ctx, req.GetAccountId().GetPlatformAccountId())
		if err != nil {
			return nil, err
		}
		appList := make([]*librarian.App, len(al))
		for i, a := range al {
			appList[i] = &librarian.App{ // TODO
				Id:               nil,
				Source:           librarian.AppSource_APP_SOURCE_STEAM,
				SourceAppId:      strconv.Itoa(int(a.AppID)),
				SourceUrl:        nil,
				Name:             a.Name,
				Type:             0,
				ShortDescription: "",
				ImageUrl:         "",
				Details:          nil,
			}
		}
		return &pb.PullAccountAppRelationResponse{AppList: appList}, nil
	default:
		return nil, status.Errorf(codes.InvalidArgument, "platform unexpected")
	}
}
func (s *LibrarianPorterServiceService) PushData(conn pb.LibrarianPorterService_PushDataServer) error {
	var file *bizs3.PutObject
	{
		req, err := conn.Recv()
		if err != nil {
			return err
		}
		if req.GetMetadata() == nil {
			return errors.BadRequest("missing metadata", "")
		}
		file, err = s.s3.NewPushData(
			conn.Context(),
			toBizBucket(req.GetMetadata().GetSource()),
			req.GetMetadata().GetContentId(),
		)
		if err != nil {
			return err
		}
	}

	for {
		if req, err := conn.Recv(); err != nil {
			if errors.Is(err, io.EOF) {
				return file.Close()
			}
			return err
		} else if len(req.GetData()) == 0 {
			return file.Close()
		} else if _, err = file.Write(req.GetData()); err != nil {
			return err
		}
	}
}
