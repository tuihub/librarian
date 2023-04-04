package libcache

import (
	"errors"

	"github.com/tuihub/librarian/internal/conf"

	"github.com/dgraph-io/ristretto"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

var ProviderSet = wire.NewSet(NewStore)

func NewStore(conf *conf.Cache) (Store, error) {
	var res Store
	var err error
	switch conf.Driver {
	case "memory":
		res, err = newRistrettoCache()
	case "redis":
		res = newRedisCache()
	default:
		return nil, errors.New("unsupported cache driver")
	}
	if err != nil {
		return nil, err
	}
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

func newRedisCache() Store {
	return newRedis(redis.NewClient(&redis.Options{ //nolint:exhaustruct // no need
		Addr: "127.0.0.1:6379",
	}))
}
