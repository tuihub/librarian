package bizbinah

import (
	"context"
	"io"
	"io/ioutil"
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

// BinahRepo is an App repo.
type BinahRepo interface {
}

// BinahUseCase is an App use case.
type BinahUseCase struct {
	callback CallbackControlBlock
	auth     *libauth.Auth
	mapper   mapper.LibrarianMapperServiceClient
	porter   porter.LibrarianPorterServiceClient
	searcher searcher.LibrarianSearcherServiceClient
}

// NewBinahUseCase new an App use case.
func NewBinahUseCase(callback CallbackControlBlock, auth *libauth.Auth, mClient mapper.LibrarianMapperServiceClient,
	pClient porter.LibrarianPorterServiceClient, sClient searcher.LibrarianSearcherServiceClient) *BinahUseCase {
	return &BinahUseCase{
		callback: callback,
		auth:     auth,
		mapper:   mClient,
		porter:   pClient,
		searcher: sClient,
	}
}

func (b *BinahUseCase) getUploadCallback(id UploadCallbackID) UploadCallbackFunc {
	f, exist := b.callback.uploadCallbackMap[id]
	if exist {
		return f
	}
	return emptyUploadCallback
}

func (b *BinahUseCase) NewUploadFile(ctx context.Context) (*UploadFile, *errors.Error) {
	claims, exist := libauth.FromContext(ctx)
	if !exist || claims == nil || claims.TransferMetadata == nil {
		return nil, pb.ErrorErrorReasonUnauthorized("token required")
	}
	f, err := ioutil.TempFile("", claims.TransferMetadata.Name)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("create temp file failed")
	}
	return &UploadFile{
		callback: b.getUploadCallback(UploadCallbackID(claims.TransferMetadata.CallBack)),
		file:     f,
		Writer:   f,
	}, nil
}
