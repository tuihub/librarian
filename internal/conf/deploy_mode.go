package conf

import (
	"errors"
	"fmt"

	"go.uber.org/zap"
)

func ApplyDeployMode(c *Config, l *zap.SugaredLogger) (*Config, error) {
	if c == nil {
		return nil, errors.New("config is nil")
	}
	deployMode := c.DeployMode
	switch deployMode {
	case DeployModeTemporary:
		return applyDeployModeTemporary(c, l)
	case DeployModeMinimize:
		return applyDeployModeMinimize(c, l)
	case DeployModeStandard:
		return applyDeployModeStandard(c, l)
	case DeployModeCluster:
		return applyDeployModeCluster(c, l)
	case "":
		return nil, errors.New("deploy mode is required")
	default:
		return nil, fmt.Errorf("unknown deploy mode %s", deployMode)
	}
}

func applyDeployModeTemporary(c *Config, l *zap.SugaredLogger) (*Config, error) {
	if c == nil {
		return nil, errors.New("config is nil")
	}
	if deployMode := c.DeployMode; deployMode != DeployModeTemporary {
		return nil, errors.New("deploy mode is not TEMPORARY")
	}

	// check Database
	if c.Database == nil {
		l.Warnf("[Database] not configured, using memory database")
		c.Database = new(Database)
		c.Database.Driver = DatabaseDriverMemory
	}
	if c.Database.Driver != DatabaseDriverMemory {
		l.Warnf("[Database] force to use memory database for temporary deploy mode")
		c.Database.Driver = DatabaseDriverMemory
	}

	// check Storage
	if c.Storage == nil {
		l.Warnf("[Storage] not configured, using memory storage")
		c.Storage = new(Storage)
		c.Storage.Driver = StorageDriverMemory
	}
	if c.Storage.Driver != StorageDriverMemory {
		l.Warnf("[Storage] force to use memory storage for temporary deploy mode")
		c.Storage.Driver = StorageDriverMemory
	}

	// check Auth
	if c.Auth == nil {
		l.Warnf("[Auth] not configured, using default auth")
		c.Auth = new(Auth)
	}

	// check MQ
	if c.MQ == nil {
		l.Warnf("[MQ] not configured, using memory mq")
		c.MQ = new(MQ)
		c.MQ.Driver = MQDriverMemory
	}

	// check Cache
	if c.Cache == nil {
		l.Warnf("[Cache] not configured, using memory cache")
		c.Cache = new(Cache)
		c.Cache.Driver = CacheDriverMemory
	}
	return c, nil
}

func applyDeployModeMinimize(c *Config, l *zap.SugaredLogger) (*Config, error) {
	if c == nil {
		return nil, errors.New("config is nil")
	}
	if deployMode := c.DeployMode; deployMode != DeployModeMinimize {
		return nil, errors.New("deploy mode is not MINIMIZE")
	}

	// check Database
	if c.Database == nil {
		l.Warnf("[Database] not configured, using file database")
		c.Database = new(Database)
		c.Database.Driver = DatabaseDriverSqlite
	}
	if c.Database.Driver == DatabaseDriverMemory {
		l.Errorf("[Database] memory driver is only allowed in temporary deploy mode")
		return nil, errors.New("[Database] memory driver is only allowed in temporary deploy mode")
	}

	// check Storage
	if c.Storage == nil {
		l.Errorf("[Storage] config required")
		return nil, errors.New("[Storage] config required")
	}
	if c.Storage.Driver == StorageDriverMemory {
		l.Errorf("[Storage] memory driver is only allowed in temporary deploy mode")
		return nil, errors.New("[Storage] memory driver is only allowed in temporary deploy mode")
	}

	// check Auth
	if c.Auth == nil {
		l.Errorf("[Auth] config required")
		return nil, errors.New("[Auth] config required")
	}
	if c.Auth.PasswordSalt == "" {
		l.Errorf("[Auth] password salt is required")
		return nil, errors.New("[Auth] password salt is required")
	}
	if c.Auth.TokenSecret == "" {
		l.Errorf("[Auth] token secret is required")
		return nil, errors.New("[Auth] token secret is required")
	}

	// check MQ
	if c.MQ == nil {
		l.Warnf("[MQ] not configured, using memory mq")
		c.MQ = new(MQ)
		c.MQ.Driver = MQDriverMemory
	}

	// check Cache
	if c.Cache == nil {
		l.Warnf("[Cache] not configured, using memory cache")
		c.Cache = new(Cache)
		c.Cache.Driver = CacheDriverMemory
	}
	return c, nil
}

func applyDeployModeStandard(c *Config, l *zap.SugaredLogger) (*Config, error) {
	if c == nil {
		return nil, errors.New("config is nil")
	}
	if deployMode := c.DeployMode; deployMode != DeployModeStandard {
		return nil, errors.New("deploy mode is not STANDARD")
	}

	// check Database
	if c.Database == nil {
		l.Errorf("[Database] config required")
		return nil, errors.New("[Database] config required")
	}
	if c.Database.Driver == DatabaseDriverMemory {
		l.Errorf("[Database] memory driver is only allowed in temporary deploy mode")
		return nil, errors.New("[Database] memory driver is only allowed in temporary deploy mode")
	}

	// check Storage
	if c.Storage == nil {
		l.Errorf("[Storage] config required")
		return nil, errors.New("[Storage] config required")
	}
	if c.Storage.Driver == StorageDriverMemory {
		l.Errorf("[Storage] memory driver is only allowed in temporary deploy mode")
		return nil, errors.New("[Storage] memory driver is only allowed in temporary deploy mode")
	}

	// check Auth
	if c.Auth == nil {
		l.Errorf("[Auth] config required")
		return nil, errors.New("[Auth] config required")
	}
	if c.Auth.PasswordSalt == "" {
		l.Errorf("[Auth] password salt is required")
		return nil, errors.New("[Auth] password salt is required")
	}
	if c.Auth.TokenSecret == "" {
		l.Errorf("[Auth] token secret is required")
		return nil, errors.New("[Auth] token secret is required")
	}

	// check MQ
	if c.MQ == nil {
		l.Warnf("[MQ] not configured, using memory mq")
		c.MQ = new(MQ)
		c.MQ.Driver = MQDriverMemory
	}

	// check Cache
	if c.Cache == nil {
		l.Warnf("[Cache] not configured, using memory cache")
		c.Cache = new(Cache)
		c.Cache.Driver = CacheDriverMemory
	}
	return c, nil
}

func applyDeployModeCluster(c *Config, l *zap.SugaredLogger) (*Config, error) {
	if c == nil {
		return nil, errors.New("config is nil")
	}
	if deployMode := c.DeployMode; deployMode != DeployModeCluster {
		return nil, errors.New("deploy mode is not CLUSTER")
	}

	// check Database
	if c.Database == nil {
		l.Errorf("[Database] config required")
		return nil, errors.New("[Database] config required")
	}
	if c.Database.Driver == DatabaseDriverMemory {
		l.Errorf("[Database] memory driver is only allowed in temporary deploy mode")
		return nil, errors.New("[Database] memory driver is only allowed in temporary deploy mode")
	}
	if c.Database.Driver == DatabaseDriverSqlite {
		l.Errorf("[Database] sqlite driver is not supported in cluster mode")
		return nil, errors.New("[Database] sqlite driver is not supported in cluster mode")
	}

	// check Storage
	if c.Storage == nil {
		l.Errorf("[Storage] config required")
		return nil, errors.New("[Storage] config required")
	}
	if c.Storage.Driver == StorageDriverMemory {
		l.Errorf("[Storage] memory driver is only allowed in temporary deploy mode")
		return nil, errors.New("[Storage] memory driver is only allowed in temporary deploy mode")
	}
	if c.Storage.Driver == StorageDriverFile {
		l.Errorf("[Storage] file driver is not supported in cluster mode")
		return nil, errors.New("[Storage] file driver is not supported in cluster mode")
	}

	// check Auth
	if c.Auth == nil {
		l.Errorf("[Auth] config required")
		return nil, errors.New("[Auth] config required")
	}
	if c.Auth.PasswordSalt == "" {
		l.Errorf("[Auth] password salt is required")
		return nil, errors.New("[Auth] password salt is required")
	}
	if c.Auth.TokenSecret == "" {
		l.Errorf("[Auth] token secret is required")
		return nil, errors.New("[Auth] token secret is required")
	}

	// check MQ
	if c.MQ == nil {
		l.Errorf("[MQ] config required")
		return nil, errors.New("[MQ] config required")
	}
	if c.MQ.Driver == MQDriverMemory {
		l.Errorf("[MQ] memory driver is not allowed in cluster mode")
		return nil, errors.New("[MQ] memory driver is not allowed in cluster mode")
	}

	// check Cache
	if c.Cache == nil {
		l.Errorf("[Cache] config required")
		return nil, errors.New("[Cache] config required")
	}
	if c.Cache.Driver == CacheDriverMemory {
		l.Errorf("[Cache] memory driver is not allowed in cluster mode")
		return nil, errors.New("[Cache] memory driver is not allowed in cluster mode")
	}
	return c, nil
}
