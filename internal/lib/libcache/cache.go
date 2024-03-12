package libcache

import (
	"errors"

	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/logger"

	"github.com/dgraph-io/ristretto"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

var ProviderSet = wire.NewSet(NewStore)

func NewStore(c *conf.Cache) (Store, error) {
	var res Store
	var err error
	if c == nil {
		c = new(conf.Cache)
	}
	if c.GetDriver() == "" {
		logger.Warnf("cache driver is not set, using memory as default")
		c.Driver = "memory"
	}
	switch c.GetDriver() {
	case "memory":
		res, err = newRistrettoCache()
	case "redis":
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
		NumCounters: 1000, //nolint:gomnd //TODO
		MaxCost:     100,  //nolint:gomnd //TODO
		BufferItems: 64,   //nolint:gomnd //TODO
	})
	if err != nil {
		return nil, err
	}
	return newRistretto(ristrettoCache), nil
}

func newRedisCache(c *conf.Cache) Store {
	return newRedis(redis.NewClient(&redis.Options{ //nolint:exhaustruct // no need
		Addr:     c.GetAddr(),
		DB:       int(c.GetDb()),
		Username: c.GetUser(),
		Password: c.GetPassword(),
	}))
}
