package bizbinah

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
		Empty: emptyUploadCallback,
	}
	return CallbackControlBlock{uploadCallbackMap: uploadCallback}
}
func emptyUploadCallback(_ *UploadFile) error { return nil }

func (c *CallbackControlBlock) RegisterUploadCallback(callback UploadCallback) {
	c.uploadCallbackMap[callback.ID] = callback.Func
}
