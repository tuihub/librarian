package service

import (
	"context"
	"io"

	"github.com/tuihub/librarian/app/porter/internal/biz/bizfeed"
	"github.com/tuihub/librarian/app/porter/internal/biz/bizs3"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	pb "github.com/tuihub/protos/pkg/librarian/porter/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LibrarianPorterServiceService struct {
	pb.UnimplementedLibrarianPorterServiceServer

	feed *bizfeed.FeedUseCase
	s3   *bizs3.S3
}

func NewLibrarianPorterServiceService(
	feed *bizfeed.FeedUseCase,
	s3 *bizs3.S3,
) pb.LibrarianPorterServiceServer {
	return &LibrarianPorterServiceService{
		UnimplementedLibrarianPorterServiceServer: pb.UnimplementedLibrarianPorterServiceServer{},
		feed: feed,
		s3:   s3,
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
	return nil, status.Errorf(codes.InvalidArgument, "platform unexpected")
}
func (s *LibrarianPorterServiceService) PullApp(
	ctx context.Context,
	req *pb.PullAppRequest,
) (*pb.PullAppResponse, error) {
	return nil, status.Errorf(codes.InvalidArgument, "source unexpected")
}
func (s *LibrarianPorterServiceService) PullAccountAppRelation(
	ctx context.Context,
	req *pb.PullAccountAppRelationRequest,
) (*pb.PullAccountAppRelationResponse, error) {
	return nil, status.Errorf(codes.InvalidArgument, "platform unexpected")
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
