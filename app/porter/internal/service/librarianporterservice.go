package service

import (
	"context"
	"io"
	"strconv"

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
}

func NewLibrarianPorterServiceService(uc *bizsteam.SteamUseCase) pb.LibrarianPorterServiceServer {
	return &LibrarianPorterServiceService{uc: uc}
}

func (s *LibrarianPorterServiceService) PullFeed(ctx context.Context, req *pb.PullFeedRequest) (
	*pb.PullFeedResponse, error) {
	return &pb.PullFeedResponse{}, nil
}
func (s *LibrarianPorterServiceService) PullDB(ctx context.Context, req *pb.PullDBRequest) (
	*pb.PullDBResponse, error) {
	return &pb.PullDBResponse{}, nil
}
func (s *LibrarianPorterServiceService) PullWiki(ctx context.Context, req *pb.PullWikiRequest) (
	*pb.PullWikiResponse, error) {
	return &pb.PullWikiResponse{}, nil
}
func (s *LibrarianPorterServiceService) PullData(req *pb.PullDataRequest,
	conn pb.LibrarianPorterService_PullDataServer) error {
	for {
		err := conn.Send(&pb.PullDataResponse{})
		if err != nil {
			return err
		}
	}
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
			Source:           req.GetAppId().GetSource(),
			SourceAppId:      req.GetAppId().GetSourceAppId(),
			SourceUrl:        &a.StoreURL,
			Name:             a.Name,
			Type:             toPBAppType(a.Type),
			ShortDescription: a.ShortDescription,
			ImageUrl:         a.ImageURL,
			Details: &librarian.AppDetails{
				Description: a.Description,
				ReleaseDate: a.ReleaseDate,
				Developer:   a.Developer,
				Publisher:   a.Publisher,
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
			appList[i] = &librarian.App{
				Source:      librarian.AppSource_APP_SOURCE_STEAM,
				SourceAppId: strconv.Itoa(int(a.AppID)),
			}
		}
		return &pb.PullAccountAppRelationResponse{AppList: appList}, nil
	default:
		return nil, status.Errorf(codes.InvalidArgument, "platform unexpected")
	}
}
func (s *LibrarianPorterServiceService) PushData(conn pb.LibrarianPorterService_PushDataServer) error {
	for {
		_, err := conn.Recv()
		if errors.Is(err, io.EOF) {
			return conn.SendAndClose(&pb.PushDataResponse{})
		}
		if err != nil {
			return err
		}
	}
}
