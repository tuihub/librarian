package service

import (
	"context"
	"io"
	"strconv"

	"github.com/tuihub/librarian/app/porter/internal/biz/bizfeed"
	"github.com/tuihub/librarian/app/porter/internal/biz/bizs3"
	"github.com/tuihub/librarian/app/porter/internal/biz/bizsteam"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	pb "github.com/tuihub/protos/pkg/librarian/porter/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LibrarianPorterServiceService struct {
	pb.UnimplementedLibrarianPorterServiceServer

	feed  *bizfeed.FeedUseCase
	steam *bizsteam.SteamUseCase
	s3    *bizs3.S3
}

func NewLibrarianPorterServiceService(
	feed *bizfeed.FeedUseCase,
	steam *bizsteam.SteamUseCase,
	s3 *bizs3.S3,
) pb.LibrarianPorterServiceServer {
	return &LibrarianPorterServiceService{
		UnimplementedLibrarianPorterServiceServer: pb.UnimplementedLibrarianPorterServiceServer{},
		feed:  feed,
		steam: steam,
		s3:    s3,
	}
}

func (s *LibrarianPorterServiceService) PullFeed(
	ctx context.Context,
	req *pb.PullFeedRequest,
) (*pb.PullFeedResponse, error) {
	switch req.GetSource() {
	case pb.FeedSource_FEED_SOURCE_UNSPECIFIED:
		return nil, status.Errorf(codes.InvalidArgument, "source unexpected")
	case pb.FeedSource_FEED_SOURCE_COMMON:
		{
			feed, err := s.feed.GetFeed(ctx, req.GetChannelId())
			if err != nil {
				return nil, err
			}
			res := modelfeed.NewConverter().ToPBFeed(feed)
			return &pb.PullFeedResponse{Data: res}, nil
		}
	default:
		return nil, status.Errorf(codes.InvalidArgument, "source unexpected")
	}
}

func (s *LibrarianPorterServiceService) PushFeedItems(ctx context.Context, req *pb.PushFeedItemsRequest) (
	*pb.PushFeedItemsResponse, error) {
	err := s.feed.PushFeedItems(ctx,
		ToBizFeedDestination(req.GetDestination()),
		modelfeed.NewConverter().FromPBFeedItemList(req.GetItems()),
		req.GetChannelId(),
		req.GetToken(),
	)
	if err != nil {
		return nil, err
	}
	return &pb.PushFeedItemsResponse{}, nil
}

func (s *LibrarianPorterServiceService) PullAccount(
	ctx context.Context,
	req *pb.PullAccountRequest,
) (*pb.PullAccountResponse, error) {
	switch req.GetAccountId().GetPlatform() {
	case librarian.AccountPlatform_ACCOUNT_PLATFORM_UNSPECIFIED:
		return nil, status.Errorf(codes.InvalidArgument, "platform unexpected")
	case librarian.AccountPlatform_ACCOUNT_PLATFORM_STEAM:
		u, err := s.steam.GetUser(ctx, req.GetAccountId().GetPlatformAccountId())
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
			LatestUpdateTime:  nil,
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
	case librarian.AppSource_APP_SOURCE_UNSPECIFIED:
		return nil, status.Errorf(codes.InvalidArgument, "source unexpected")
	case librarian.AppSource_APP_SOURCE_INTERNAL:
		return nil, status.Errorf(codes.InvalidArgument, "source unexpected")
	case librarian.AppSource_APP_SOURCE_STEAM:
		appID, err := strconv.Atoi(req.GetAppId().GetSourceAppId())
		if err != nil {
			return nil, err
		}
		a, err := s.steam.GetAppDetails(ctx, appID)
		if err != nil {
			return nil, err
		}
		return &pb.PullAppResponse{App: &librarian.App{
			Id:               nil,
			Source:           req.GetAppId().GetSource(),
			SourceAppId:      req.GetAppId().GetSourceAppId(),
			SourceUrl:        &a.StoreURL,
			Name:             a.Name,
			Type:             ToPBAppType(a.Type),
			ShortDescription: a.ShortDescription,
			IconImageUrl:     "",
			Tags:             nil,
			Details: &librarian.AppDetails{ // TODO
				Description:  a.Description,
				ReleaseDate:  a.ReleaseDate,
				Developer:    a.Developer,
				Publisher:    a.Publisher,
				Version:      "",
				HeroImageUrl: a.HeroImageURL,
				LogoImageUrl: "",
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
	case librarian.AccountPlatform_ACCOUNT_PLATFORM_UNSPECIFIED:
		return nil, status.Errorf(codes.InvalidArgument, "platform unexpected")
	case librarian.AccountPlatform_ACCOUNT_PLATFORM_STEAM:
		al, err := s.steam.GetOwnedGames(ctx, req.GetAccountId().GetPlatformAccountId())
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
				IconImageUrl:     a.IconImageURL,
				Tags:             nil,
				Details: &librarian.AppDetails{
					Description:  "",
					ReleaseDate:  "",
					Developer:    "",
					Publisher:    "",
					Version:      "",
					HeroImageUrl: "",
					LogoImageUrl: a.LogoImageURL,
				},
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
			ToBizBucket(req.GetMetadata().GetSource()),
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

func (s *LibrarianPorterServiceService) PresignedPushData(ctx context.Context, req *pb.PresignedPushDataRequest) (
	*pb.PresignedPushDataResponse, error) {
	res, err := s.s3.PresignedPutData(ctx,
		ToBizBucket(req.GetSource()),
		req.GetContentId(),
		req.GetExpireTime().AsDuration(),
	)
	if err != nil {
		return nil, err
	}
	return &pb.PresignedPushDataResponse{PushUrl: res}, nil
}

func (s *LibrarianPorterServiceService) PresignedPullData(ctx context.Context, req *pb.PresignedPullDataRequest) (
	*pb.PresignedPullDataResponse, error) {
	res, err := s.s3.PresignedGetData(ctx,
		ToBizBucket(req.GetSource()),
		req.GetContentId(),
		req.GetExpireTime().AsDuration(),
	)
	if err != nil {
		return nil, err
	}
	return &pb.PresignedPullDataResponse{PullUrl: res}, nil
}
