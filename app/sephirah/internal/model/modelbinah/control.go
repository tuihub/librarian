package modelbinah

import (
	"context"
	"errors"

	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/logger"
)

type ControlBlock struct {
	a                 *libauth.Auth
	uploadCallbackMap map[UploadCallbackID]UploadCallbackFunc
}

type UploadCallbackID int
type UploadCallbackFunc func() error

const (
	Empty UploadCallbackID = iota
	UploadArtifacts
	ChesedUploadImage
)

func NewControlBlock(a *libauth.Auth) *ControlBlock {
	return &ControlBlock{
		a: a,
		uploadCallbackMap: map[UploadCallbackID]UploadCallbackFunc{ //nolint:exhaustive //no need
			Empty: emptyUploadCallback,
		},
	}
}
func emptyUploadCallback() error { return nil }
func unregisteredCallback() error {
	err := errors.New("calling unregistered upload callback")
	logger.Error(err)
	return err
}

func (c *ControlBlock) RegisterUploadCallback(id UploadCallbackID, fn UploadCallbackFunc) *UploadCallBack {
	c.uploadCallbackMap[id] = fn
	return &UploadCallBack{
		id:           id,
		controlBlock: c,
	}
}

func (c *ControlBlock) GetUploadCallback(ctx context.Context) (UploadCallbackFunc, error) {
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

func (c *ControlBlock) GetUploadFileMetadata(ctx context.Context) (*FileMetadata, error) {
	payload, err := c.getUploadPayload(ctx)
	if err != nil {
		return nil, err
	}
	return &payload.FileMetadata, nil
}

type uploadTokenPayload struct {
	FileMetadata
	UploadCallbackID
}

func (c *ControlBlock) getUploadPayload(ctx context.Context) (*uploadTokenPayload, error) {
	claims, exist := libauth.FromContext(ctx)
	if !exist {
		return nil, errors.New("token required")
	}
	if claims.TransferMetadata == nil {
		return nil, errors.New("broken token")
	}
	payload, met := claims.TransferMetadata.(*uploadTokenPayload)
	if !met {
		return nil, errors.New("broken token")
	}
	return payload, nil
}
