package modelbinah

import (
	"context"
	"errors"

	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcodec"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/logger"
)

type ControlBlock struct {
	a                   *libauth.Auth
	uploadCallbackMap   map[UploadCallbackID]CallbackFunc
	downloadCallbackMap map[DownloadCallbackID]CallbackFunc
}

type UploadCallbackID int
type DownloadCallbackID int
type CallbackFunc func(context.Context, model.InternalID) error

const (
	UploadEmpty UploadCallbackID = iota
	UploadArtifacts
	UploadChesedImage
)

const (
	DownloadEmpty DownloadCallbackID = iota
)

func NewControlBlock(a *libauth.Auth) *ControlBlock {
	return &ControlBlock{
		a: a,
		uploadCallbackMap: map[UploadCallbackID]CallbackFunc{ //nolint:exhaustive //no need
			UploadEmpty: emptyCallback,
		},
		downloadCallbackMap: map[DownloadCallbackID]CallbackFunc{
			DownloadEmpty: emptyCallback,
		},
	}
}
func emptyCallback(_ context.Context, _ model.InternalID) error { return nil }
func unregisteredCallback(_ context.Context, _ model.InternalID) error {
	err := errors.New("calling unregistered upload callback")
	logger.Error(err)
	return err
}

func (c *ControlBlock) RegisterUploadCallback(id UploadCallbackID, fn CallbackFunc) *UploadCallBack {
	if id != UploadEmpty {
		c.uploadCallbackMap[id] = fn
	}
	return &UploadCallBack{
		id:           id,
		controlBlock: c,
	}
}

func (c *ControlBlock) RegisterDownloadCallback(id DownloadCallbackID, fn CallbackFunc) *DownloadCallBack {
	if id != DownloadEmpty {
		c.downloadCallbackMap[id] = fn
	}
	return &DownloadCallBack{
		id:           id,
		controlBlock: c,
	}
}

func (c *ControlBlock) GetUploadCallback(ctx context.Context) (CallbackFunc, error) {
	payload, err := c.getUploadPayload(ctx)
	if err != nil {
		return nil, err
	}
	f, exist := c.uploadCallbackMap[payload.UploadCallbackID]
	if exist {
		return f, nil
	}
	return unregisteredCallback, nil
}

func (c *ControlBlock) GetDownloadCallback(ctx context.Context) (CallbackFunc, error) {
	payload, err := c.getDownloadPayload(ctx)
	if err != nil {
		return nil, err
	}
	f, exist := c.downloadCallbackMap[payload.DownloadCallbackID]
	if exist {
		return f, nil
	}
	return unregisteredCallback, nil
}

func (c *ControlBlock) GetUploadFileMetadata(ctx context.Context) (*FileMetadata, error) {
	payload, err := c.getUploadPayload(ctx)
	if err != nil {
		return nil, err
	}
	return &payload.FileMetadata, nil
}

func (c *ControlBlock) GetDownloadFileMetadata(ctx context.Context) (*FileMetadata, error) {
	payload, err := c.getDownloadPayload(ctx)
	if err != nil {
		return nil, err
	}
	return &payload.FileMetadata, nil
}

type uploadTokenPayload struct {
	FileMetadata
	UploadCallbackID
}

type downloadTokenPayload struct {
	FileMetadata
	DownloadCallbackID
}

func (c *ControlBlock) getUploadPayload(ctx context.Context) (*uploadTokenPayload, error) {
	claims := libauth.FromContext(ctx)
	if claims == nil {
		return nil, errors.New("token required")
	}
	if claims.TransferMetadata == nil {
		return nil, errors.New("broken token")
	}
	bytes, err := libcodec.Marshal(libcodec.JSON, claims.TransferMetadata)
	if err != nil {
		return nil, err
	}
	payload := new(uploadTokenPayload)
	err = libcodec.Unmarshal(libcodec.JSON, bytes, &payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func (c *ControlBlock) getDownloadPayload(ctx context.Context) (*downloadTokenPayload, error) {
	claims := libauth.FromContext(ctx)
	if claims == nil {
		return nil, errors.New("token required")
	}
	if claims.TransferMetadata == nil {
		return nil, errors.New("broken token")
	}
	bytes, err := libcodec.Marshal(libcodec.JSON, claims.TransferMetadata)
	if err != nil {
		return nil, err
	}
	payload := new(downloadTokenPayload)
	err = libcodec.Unmarshal(libcodec.JSON, bytes, &payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}
