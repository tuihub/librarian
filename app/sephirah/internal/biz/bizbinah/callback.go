package bizbinah

import (
	"errors"

	"github.com/tuihub/librarian/internal/lib/logger"
)

type CallbackControlBlock struct {
	uploadCallbackMap map[UploadCallbackID]UploadCallbackFunc
}
type UploadCallback struct {
	ID   UploadCallbackID
	Func UploadCallbackFunc
}
type UploadCallbackID int
type UploadCallbackFunc func(*UploadFile) error

const (
	Empty UploadCallbackID = iota
	UploadArtifacts
)

func NewCallbackControl() CallbackControlBlock {
	uploadCallback := map[UploadCallbackID]UploadCallbackFunc{
		Empty:           emptyUploadCallback,
		UploadArtifacts: unregisteredCallback,
	}
	return CallbackControlBlock{uploadCallbackMap: uploadCallback}
}
func emptyUploadCallback(_ *UploadFile) error { return nil }
func unregisteredCallback(_ *UploadFile) error {
	err := errors.New("calling unregistered upload callback")
	logger.Error(err)
	return err
}

func (c *CallbackControlBlock) RegisterUploadCallback(callback UploadCallback) {
	c.uploadCallbackMap[callback.ID] = callback.Func
}
