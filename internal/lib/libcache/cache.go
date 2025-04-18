package libcache

import (
	"errors"
	"net"
	"strconv"

	"github.com/tuihub/librarian/internal/conf"

	"github.com/dgraph-io/ristretto"
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
		res, err = newRistrettoCache()
	case conf.CacheDriverRedis:
		res = newRedisCache(c)
	default:
		return nil, errors.New("unsupported cache driver")
	}
	if err != nil {
		return nil, err
	}
	initCaptchaStore(res)
	return res, nil
}

func newRistrettoCache() (Store, error) {
	ristrettoCache, err := ristretto.NewCache(&ristretto.Config{ //nolint:exhaustruct // no need
		NumCounters: 1000, //nolint:mnd //TODO
		MaxCost:     100,  //nolint:mnd //TODO
		BufferItems: 64,   //nolint:mnd //TODO
	})
	if err != nil {
		return nil, err
	}
	return newRistretto(ristrettoCache), nil
}

func newRedisCache(c *conf.Cache) Store {
	return newRedis(redis.NewClient(&redis.Options{ //nolint:exhaustruct // no need
		Addr:     net.JoinHostPort(c.Host, strconv.Itoa(int(c.Port))),
		DB:       int(c.DB),
		Username: c.Username,
		Password: c.Password,
	}))
}
