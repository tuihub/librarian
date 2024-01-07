package bizbinah

import (
	"context"
	"io"
	"os"
	"strconv"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modelbinah"
	"github.com/tuihub/librarian/model"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/google/uuid"
)

type SimpleUploadFile struct {
	id       model.InternalID
	porter   porter.LibrarianPorterServiceClient
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
		porter:   b.porter,
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
	cli, err := f.porter.PushData(ctx)
	if err != nil {
		return err
	}
	err = cli.Send(&porter.PushDataRequest{
		Content: &porter.PushDataRequest_Metadata{
			Metadata: &porter.PushDataRequest_DataMeta{
				Source:    porter.DataSource_DATA_SOURCE_INTERNAL_DEFAULT,
				ContentId: strconv.FormatInt(int64(f.id), 10),
			},
		},
	})
	if err != nil {
		return err
	}
	buf := make([]byte, 32<<10) //nolint:gomnd //TODO
	for {
		var n int
		n, err = f.file.Read(buf)
		if n == 0 && errors.Is(err, io.EOF) {
			_, err = cli.CloseAndRecv()
			if !errors.Is(err, io.EOF) {
				return err
			}
			break
		}
		if err != nil {
			return err
		}
		err = cli.Send(&porter.PushDataRequest{
			Content: &porter.PushDataRequest_Data{
				Data: buf[:n],
			},
		})
		if err != nil {
			return err
		}
	}
	return f.callback(ctx, f.id)
}
