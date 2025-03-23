package bizbinah

import (
	"context"
	"strconv"

	"github.com/tuihub/librarian/internal/data"
	"github.com/tuihub/librarian/internal/lib/libtime"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

func (b *Binah) PresignedUploadFile(ctx context.Context) (string, *errors.Error) {
	metadata, err := b.callback.GetDownloadFileMetadata(ctx)
	if err != nil {
		return "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	res, err := b.repo.PresignedPutObject(ctx, data.BucketDefault, strconv.FormatInt(int64(metadata.ID), 10), libtime.Day)
	if err != nil {
		return "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return res, nil
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
	res, err := b.repo.PresignedGetObject(ctx, data.BucketDefault, strconv.FormatInt(int64(metadata.ID), 10), libtime.Day)
	if err != nil {
		return "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return res, nil
}
