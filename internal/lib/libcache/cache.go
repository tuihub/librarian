package libcache

import (
	"errors"
	"fmt"
	"net"
	"strconv"

	"github.com/tuihub/librarian/internal/conf"

	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

var ProviderSet = wire.NewSet(NewStore)

func NewStore(c *conf.Cache) (Store, error) {
	var res Store
	var err error
	if c == nil {
		return nil, errors.New("cache config is nil")
	}
	switch c.Driver {
	case conf.CacheDriverMemory:
		res, err = newMemoryCache()
	case conf.CacheDriverRedis:
		res = newRedisCache(c)
	default:
		return nil, errors.New("unsupported cache driver")
	}
	if err != nil {
		return nil, fmt.Errorf("failed creating cache: %w", err)
	}
	initCaptchaStore(res)
	return res, nil
}

func newMemoryCache() (Store, error) {
	return newOtter(1024) //nolint:mnd // no need
}

func newRedisCache(c *conf.Cache) Store {
	return newRedis(redis.NewClient(&redis.Options{ //nolint:exhaustruct // no need
		Addr:     net.JoinHostPort(c.Host, strconv.Itoa(int(c.Port))),
		DB:       int(c.DB),
		Username: c.Username,
		Password: c.Password,
	}))
}
