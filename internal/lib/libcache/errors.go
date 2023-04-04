package libcache

const NotFoundErr string = "value not found in store"

type NotFoundError struct {
	cause error
}

func newNotFound(e error) error {
	err := NotFoundError{
		cause: e,
	}
	return &err
}

func (e NotFoundError) Cause() error {
	return e.cause
}

func (e NotFoundError) Is(err error) bool {
	return err.Error() == NotFoundErr
}

func (e NotFoundError) Error() string {
	return NotFoundErr
}
func (e NotFoundError) Unwrap() error { return e.cause }
