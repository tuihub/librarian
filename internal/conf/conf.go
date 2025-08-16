package conf

import (
	"errors"
	"time"

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
	DeployMode             DeployMode              `json:"deploy_mode"              toml:"deploy_mode"              yaml:"deploy_mode"`
	Server                 *Server                 `json:"server"                   toml:"server"                   yaml:"server"`
	Database               *Database               `json:"database"                 toml:"database"                 yaml:"database"`
	Storage                *Storage                `json:"storage"                  toml:"storage"                  yaml:"storage"`
	Auth                   *Auth                   `json:"auth"                     toml:"auth"                     yaml:"auth"`
	MQ                     *MQ                     `json:"mq"                       toml:"mq"                       yaml:"mq"`
	Cache                  *Cache                  `json:"cache"                    toml:"cache"                    yaml:"cache"`
	Porter                 *Porter                 `json:"porter"                   toml:"porter"                   yaml:"porter"`
	Consul                 *Consul                 `json:"consul"                   toml:"consul"                   yaml:"consul"`
	EnableServiceDiscovery *EnableServiceDiscovery `json:"enable_service_discovery" toml:"enable_service_discovery" yaml:"enable_service_discovery"`
	Search                 *Search                 `json:"search"                   toml:"search"                   yaml:"search"`
	OTLP                   *OpenTelemetry          `json:"open_telemetry"           toml:"open_telemetry"           yaml:"open_telemetry"`
}

type Server struct {
	Admin   *HTTP `json:"admin"    toml:"admin"    yaml:"admin"`
	Main    *GRPC `json:"main"     toml:"main"     yaml:"main"`
	MainWeb *GRPC `json:"main_web" toml:"main_web" yaml:"main_web"`
}

type HTTP struct {
	Host    string        `json:"host"    toml:"host"    yaml:"host"`
	Port    uint32        `json:"port"    toml:"port"    yaml:"port"`
	Timeout time.Duration `json:"timeout" toml:"timeout" yaml:"timeout"`
}

type GRPC struct {
	Host    string        `json:"host"    toml:"host"    yaml:"host"`
	Port    uint32        `json:"port"    toml:"port"    yaml:"port"`
	Timeout time.Duration `json:"timeout" toml:"timeout" yaml:"timeout"`
}

type DatabaseDriver string

const (
	DatabaseDriverMemory   DatabaseDriver = "memory"
	DatabaseDriverSqlite   DatabaseDriver = "sqlite3"
	DatabaseDriverPostgres DatabaseDriver = "postgres"
)

type Database struct {
	Driver   DatabaseDriver `json:"driver"   toml:"driver"   yaml:"driver"`
	Host     string         `json:"host"     toml:"host"     yaml:"host"`
	Port     uint32         `json:"port"     toml:"port"     yaml:"port"`
	DBName   string         `json:"db_name"  toml:"db_name"  yaml:"db_name"`
	Username string         `json:"username" toml:"username" yaml:"username"`
	Password string         `json:"password" toml:"password" yaml:"password"`
}

type StorageDriver string

const (
	StorageDriverMemory StorageDriver = "memory"
	StorageDriverFile   StorageDriver = "file"
	StorageDriverS3     StorageDriver = "s3"
)

type Storage struct {
	Driver    StorageDriver `json:"driver"     toml:"driver"     yaml:"driver"`
	Host      string        `json:"host"       toml:"host"       yaml:"host"`
	Port      uint32        `json:"port"       toml:"port"       yaml:"port"`
	Region    string        `json:"region"     toml:"region"     yaml:"region"`
	AccessKey string        `json:"access_key" toml:"access_key" yaml:"access_key"`
	SecretKey string        `json:"secret_key" toml:"secret_key" yaml:"secret_key"`
	Token     string        `json:"token"      toml:"token"      yaml:"token"`
}

type Auth struct {
	PasswordSalt string `json:"password_salt" toml:"password_salt" yaml:"password_salt"`
	TokenIssuer  string `json:"token_issuer"  toml:"token_issuer"  yaml:"token_issuer"`
	TokenSecret  string `json:"token_secret"  toml:"token_secret"  yaml:"token_secret"`
}

type MQDriver string

const (
	MQDriverMemory MQDriver = "memory"
	MQDriverSQL    MQDriver = "sql"
	MQDriverRedis  MQDriver = "redis"
)

type MQ struct {
	Driver MQDriver `json:"driver" toml:"driver" yaml:"driver"`
}

type CacheDriver string

const (
	CacheDriverMemory CacheDriver = "memory"
	CacheDriverRedis  CacheDriver = "redis"
)

type Cache struct {
	Driver   CacheDriver `json:"driver"   toml:"driver"   yaml:"driver"`
	Host     string      `json:"host"     toml:"host"     yaml:"host"`
	Port     uint32      `json:"port"     toml:"port"     yaml:"port"`
	DB       int32       `json:"db"       toml:"db"       yaml:"db"`
	Username string      `json:"username" toml:"username" yaml:"username"`
	Password string      `json:"password" toml:"password" yaml:"password"`
}

type SearchDriver string

const (
	SearchDriverDisable SearchDriver = ""
	SearchDriverBleve   SearchDriver = "bleve"
	SearchDriverMeili   SearchDriver = "meili"
)

type Search struct {
	Driver    SearchDriver `json:"driver"     toml:"driver"     yaml:"driver"`
	MeiliHost string       `json:"meili_host" toml:"meili_host" yaml:"meili_host"`
	MeiliPort uint32       `json:"meili_port" toml:"meili_port" yaml:"meili_port"`
	MeiliKey  string       `json:"meili_key"  toml:"meili_key"  yaml:"meili_key"`
}

type OpenTelemetryProtocol string

const (
	OpenTelemetryProtocolDisable OpenTelemetryProtocol = ""
	OpenTelemetryProtocolGRPC    OpenTelemetryProtocol = "grpc"
	OpenTelemetryProtocolHTTP    OpenTelemetryProtocol = "http"
)

type OpenTelemetry struct {
	EnableMemoryMetrics bool                  `json:"enable_memory_metrics" toml:"enable_memory_metrics" yaml:"enable_memory_metrics"`
	Protocol            OpenTelemetryProtocol `json:"protocol"              toml:"protocol"              yaml:"protocol"`
	Endpoint            string                `json:"endpoint"              toml:"endpoint"              yaml:"endpoint"`
	Headers             string                `json:"headers"               toml:"headers"               yaml:"headers"`
	GRPCInsecure        bool                  `json:"grpc_insecure"         toml:"grpc_insecure"         yaml:"grpc_insecure"`
}

type Consul struct {
	Host  string `json:"host"  toml:"host"  yaml:"host"`
	Port  uint32 `json:"port"  toml:"port"  yaml:"port"`
	Token string `json:"token" toml:"token" yaml:"token"`
}

type EnableServiceDiscovery struct {
	Porter bool `json:"porter" toml:"porter" yaml:"porter"`
}

type Porter struct {
	Addresses []string `json:"addresses" toml:"addresses" yaml:"addresses"`
}
