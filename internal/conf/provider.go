package conf

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	GetEnableServiceDiscovery,
	GetServer,
	GetDatabase,
	GetStorage,
	GetPorter,
	GetMinerData,
	GetAuth,
	GetMQ,
	GetCache,
	GetConsul,
	GetSearch,
)

func GetEnableServiceDiscovery(c *Config) *EnableServiceDiscovery {
	return c.EnableServiceDiscovery
}
func GetServer(c *Config) *Server {
	return c.Server
}
func GetDatabase(c *Config) *Database {
	return c.Database
}
func GetStorage(c *Config) *Storage {
	return c.Storage
}
func GetPorter(c *Config) *Porter {
	return c.Porter
}
func GetMinerData(c *Config) *Miner_Data {
	return c.Miner.GetData()
}
func GetAuth(c *Config) *Auth {
	return c.Auth
}
func GetMQ(c *Config) *MQ {
	return c.MQ
}
func GetCache(c *Config) *Cache {
	return c.Cache
}
func GetConsul(c *Config) *Consul {
	return c.Consul
}
func GetSearch(c *Config) *Search {
	return c.Search
}
