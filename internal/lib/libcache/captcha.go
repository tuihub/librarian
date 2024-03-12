package libcache

import (
	"context"

	"github.com/dchest/captcha"
)

const captchaStoreKey = "captcha"

func initCaptchaStore(store Store) {
	captcha.SetCustomStore(&captchaStoreImpl{store})
}

type captchaStoreImpl struct {
	store Store
}

func (c *captchaStoreImpl) Set(id string, digits []byte) {
	_ = c.store.Set(context.Background(), captchaStoreKey+":"+id, string(digits), WithExpiration(captcha.Expiration))
}

func (c *captchaStoreImpl) Get(id string, clear bool) []byte {
	get, err := c.store.Get(context.Background(), captchaStoreKey+":"+id)
	if err != nil {
		return nil
	}
	digits, ok := get.(string)
	if !ok {
		return nil
	}
	if clear {
		_ = c.store.Delete(context.Background(), captchaStoreKey+id)
	}
	return []byte(digits)
}
