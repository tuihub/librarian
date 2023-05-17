package bizbinah

import (
	"context"
	"io"
	"os"
	"strconv"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modelbinah"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/google/uuid"
)

type UploadFile struct {
	id       model.InternalID
	porter   porter.LibrarianPorterServiceClient
	callback modelbinah.UploadCallbackFunc
	file     *os.File
	Writer   io.Writer
}

func (f *UploadFile) Finish(ctx context.Context) error {
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

type BinahRepo interface {
}

type Binah struct {
	callback *modelbinah.ControlBlock
	auth     *libauth.Auth
	mapper   mapper.LibrarianMapperServiceClient
	porter   porter.LibrarianPorterServiceClient
	searcher searcher.LibrarianSearcherServiceClient
}

func NewBinah(
	callback *modelbinah.ControlBlock,
	auth *libauth.Auth,
	mClient mapper.LibrarianMapperServiceClient,
	pClient porter.LibrarianPorterServiceClient,
	sClient searcher.LibrarianSearcherServiceClient,
) *Binah {
	return &Binah{
		callback: callback,
		auth:     auth,
		mapper:   mClient,
		porter:   pClient,
		searcher: sClient,
	}
}

func NewControlBlock(a *libauth.Auth) *modelbinah.ControlBlock {
	return modelbinah.NewControlBlock(a)
}

func (b *Binah) NewUploadFile(ctx context.Context) (*UploadFile, *errors.Error) {
	claims, exist := libauth.FromContext(ctx)
	if !exist || claims == nil || claims.TransferMetadata == nil {
		return nil, pb.ErrorErrorReasonUnauthorized("token required")
	}
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
	return &UploadFile{
		id:       meta.ID,
		porter:   b.porter,
		callback: callback,
		file:     f,
		Writer:   f,
	}, nil
}
