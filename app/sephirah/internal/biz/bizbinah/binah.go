package bizbinah

import (
	"context"
	"io"
	"os"

	"github.com/tuihub/librarian/internal/lib/libauth"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

type UploadFile struct {
	callback UploadCallbackFunc
	file     *os.File
	Writer   io.Writer
}

func (f *UploadFile) Finish() error {
	return f.callback(f)
}

type BinahRepo interface {
}

type Binah struct {
	callback CallbackControlBlock
	auth     *libauth.Auth
	mapper   mapper.LibrarianMapperServiceClient
	porter   porter.LibrarianPorterServiceClient
	searcher searcher.LibrarianSearcherServiceClient
}

func NewBinah(callback CallbackControlBlock, auth *libauth.Auth, mClient mapper.LibrarianMapperServiceClient,
	pClient porter.LibrarianPorterServiceClient, sClient searcher.LibrarianSearcherServiceClient) *Binah {
	return &Binah{
		callback: callback,
		auth:     auth,
		mapper:   mClient,
		porter:   pClient,
		searcher: sClient,
	}
}

func (b *Binah) getUploadCallback(id UploadCallbackID) UploadCallbackFunc {
	f, exist := b.callback.uploadCallbackMap[id]
	if exist {
		return f
	}
	return emptyUploadCallback
}

func (b *Binah) NewUploadFile(ctx context.Context) (*UploadFile, *errors.Error) {
	claims, exist := libauth.FromContext(ctx)
	if !exist || claims == nil || claims.TransferMetadata == nil {
		return nil, pb.ErrorErrorReasonUnauthorized("token required")
	}
	f, err := os.CreateTemp("", claims.TransferMetadata.Name)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("create temp file failed")
	}
	return &UploadFile{
		callback: b.getUploadCallback(UploadCallbackID(claims.TransferMetadata.CallBack)),
		file:     f,
		Writer:   f,
	}, nil
}
