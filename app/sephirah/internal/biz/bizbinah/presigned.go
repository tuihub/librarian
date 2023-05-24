package bizbinah

import (
	"context"
	"strconv"

	"github.com/tuihub/librarian/internal/lib/libtime"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/protobuf/types/known/durationpb"
)

func (b *Binah) PresignedUploadFile(ctx context.Context) (string, *errors.Error) {
	metadata, err := b.callback.GetDownloadFileMetadata(ctx)
	if err != nil {
		return "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	res, err := b.porter.PresignedPushData(ctx, &porter.PresignedPushDataRequest{
		Source:     porter.DataSource_DATA_SOURCE_INTERNAL_DEFAULT,
		ContentId:  strconv.FormatInt(int64(metadata.ID), 10),
		ExpireTime: durationpb.New(libtime.Day),
	})
	if err != nil {
		return "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return res.GetPushUrl(), nil
}

func (b *Binah) PresignedUploadFileComplete(ctx context.Context) *errors.Error {
	callback, err := b.callback.GetUploadCallback(ctx)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	metadata, err := b.callback.GetDownloadFileMetadata(ctx)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	err = callback(ctx, metadata.ID)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (b *Binah) PresignedDownloadFile(ctx context.Context) (string, *errors.Error) {
	metadata, err := b.callback.GetDownloadFileMetadata(ctx)
	if err != nil {
		return "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	res, err := b.porter.PresignedPullData(ctx, &porter.PresignedPullDataRequest{
		Source:     porter.DataSource_DATA_SOURCE_INTERNAL_DEFAULT,
		ContentId:  strconv.FormatInt(int64(metadata.ID), 10),
		ExpireTime: durationpb.New(libtime.Day),
	})
	if err != nil {
		return "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return res.GetPullUrl(), nil
}
