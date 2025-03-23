package bizbinah

import (
	"context"
	"io"
	"os"
	"strconv"

	"github.com/tuihub/librarian/internal/data"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelbinah"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/google/uuid"
)

type SimpleUploadFile struct {
	id       model.InternalID
	repo     *data.BinahRepo
	callback modelbinah.CallbackFunc
	file     *os.File
	Writer   io.Writer
}

func (b *Binah) NewSimpleUploadFile(ctx context.Context) (*SimpleUploadFile, *errors.Error) {
	callback, err := b.callback.GetUploadCallback(ctx)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	meta, err := b.callback.GetUploadFileMetadata(ctx)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	f, err := os.CreateTemp("", uuid.NewString())
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("create temp file failed")
	}
	return &SimpleUploadFile{
		id:       meta.ID,
		repo:     b.repo,
		callback: callback,
		file:     f,
		Writer:   f,
	}, nil
}

func (f *SimpleUploadFile) Finish(ctx context.Context) error {
	err := f.file.Sync()
	if err != nil {
		return err
	}
	_, err = f.file.Seek(0, 0)
	if err != nil {
		return err
	}
	return f.repo.PutObject(ctx, f.file, data.BucketDefault, strconv.FormatInt(int64(f.id), 10))
}
