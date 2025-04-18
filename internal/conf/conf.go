package conf

import (
	"errors"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
)

func Load(path string) (*Config, error) {
	if path == "" {
		return nil, errors.New("path is empty")
	}
	c := config.New(
		config.WithSource(
			file.NewSource(path),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		return nil, err
	}

	ret := new(Config)
	if err := c.Scan(ret); err != nil {
		return nil, err
	}
	return ret, nil
}

type DeployMode string

const (
	DeployModeTemporary DeployMode = "temporary"
	DeployModeMinimize  DeployMode = "minimize"
	DeployModeStandard  DeployMode = "standard"
	DeployModeCluster   DeployMode = "cluster"
)

type Config struct {
	DeployMode             DeployMode
	Server                 *Server
	Database               *Database
	Storage                *Storage
	Auth                   *Auth
	MQ                     *MQ
	Cache                  *Cache
	Porter                 *Porter
	Consul                 *Consul
	EnableServiceDiscovery *EnableServiceDiscovery
	Search                 *Search
	Miner                  *Miner
	Otlp                   *OTLP
}

type DatabaseDriver string

const (
	DatabaseDriverMemory   DatabaseDriver = "memory"
	DatabaseDriverSqlite   DatabaseDriver = "sqlite3"
	DatabaseDriverPostgres DatabaseDriver = "postgres"
)

type Database struct {
	Driver   DatabaseDriver
	Host     string
	Port     uint32
	DBName   string
	Username string
	Password string
}

type StorageDriver string

const (
	StorageDriverMemory StorageDriver = "memory"
	StorageDriverFile   StorageDriver = "file"
	StorageDriverS3     StorageDriver = "s3"
)

type Storage struct {
	Driver    StorageDriver
	Host      string
	Port      uint32
	Region    string
	AccessKey string
	SecretKey string
	Token     string
}

type Auth struct {
	PasswordSalt string
	TokenIssuer  string
	TokenSecret  string
}

type MQDriver string

const (
	MQDriverMemory MQDriver = "memory"
	MQDriverSQL    MQDriver = "sql"
	MQDriverRedis  MQDriver = "redis"
)

type MQ struct {
	Driver MQDriver
}

type CacheDriver string

const (
	CacheDriverMemory CacheDriver = "memory"
	CacheDriverRedis  CacheDriver = "redis"
)

type Cache struct {
	Driver   CacheDriver
	Host     string
	Port     uint32
	DB       int32
	Username string
	Password string
}
