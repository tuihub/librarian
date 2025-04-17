package conf

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	GetEnableServiceDiscovery,
	GetServer,
	GetDatabase,
	GetS3,
	GetPorter,
	GetMinerData,
	GetAuth,
	GetMQ,
	GetCache,
	GetConsul,
	GetSearch,
)

func GetEnableServiceDiscovery(c *Config) *EnableServiceDiscovery {
	return c.GetEnableServiceDiscovery()
}
func GetServer(c *Config) *SephirahServer {
	return c.GetServer()
}
func GetDatabase(c *Config) *Database {
	return c.GetDatabase()
}
func GetS3(c *Config) *S3 {
	return c.GetS3()
}
func GetPorter(c *Config) *Porter {
	return c.GetPorter()
}
func GetMinerData(c *Config) *Miner_Data {
	return c.GetMiner().GetData()
}
func GetAuth(c *Config) *Auth {
	return c.GetAuth()
}
func GetMQ(c *Config) *MQ {
	return c.GetMq()
}
func GetCache(c *Config) *Cache {
	return c.GetCache()
}
func GetConsul(c *Config) *Consul {
	return c.GetConsul()
}
func GetSearch(c *Config) *Search {
	return c.GetSearch()
}
