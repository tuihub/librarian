package libapp

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/tuihub/librarian/internal/lib/libzap"

	"github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/env"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewDiscovery, NewRegistrar, NewNodeFilter)

type Settings struct {
	Name string
	env  config.Config
	InherentSettings
	Flags
	Version           string
	ProtoVersion      string
	BuildDate         string
	SourceCodeAddress string
	DemoMode          bool
}

type Flags struct {
	ConfPath string
	DataPath string
}

type Env string

const (
	EnvCreateAdminUserName Env = "CREATE_ADMIN_USER"
	EnvCreateAdminPassword Env = "CREATE_ADMIN_PASS"
	EnvLogLevel            Env = "LOG_LEVEL"
	EnvDemoMode            Env = "DEMO_MODE"
	EnvAllowRegister       Env = "ALLOW_REGISTER"
	EnvUserCapacity        Env = "USER_CAPACITY"
)

func NewAppSettings(id, name, version, protoVersion, date string) (*Settings, error) {
	var as Settings
	flags := loadFlags()
	if err := checkDataPath(flags.DataPath); err != nil {
		return nil, err
	}
	initLogger(id, name, version, flags.DataPath, getInherentSettings().LogLevel, true)
	if e, err := loadEnv(); err != nil {
		return nil, err
	} else {
		as = Settings{
			name,
			e,
			getInherentSettings(),
			flags,
			version,
			protoVersion,
			date,
			"https://github.com/TuiHub/Librarian",
			false,
		}
	}
	if as.ConfPath == "" {
		as.ConfPath = as.DefaultConfPath
	}
	// override LogLevel by env
	if logLevel, err := as.Env(EnvLogLevel); err == nil {
		switch strings.ToLower(logLevel) {
		case "debug":
			as.LogLevel = libzap.DebugLevel
		case "info":
			as.LogLevel = libzap.InfoLevel
		case "warn":
			as.LogLevel = libzap.WarnLevel
		case "error":
			as.LogLevel = libzap.ErrorLevel
		}
	}
	initLogger(id, name, version, as.DataPath, as.LogLevel, as.ConfPath == "")
	return &as, nil
}

func (a *Settings) LoadConfig(conf interface{}) error {
	if a.ConfPath == "" {
		return nil
	}
	c := config.New(
		config.WithSource(
			file.NewSource(a.ConfPath),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		return err
	}

	if err := c.Scan(conf); err != nil {
		return err
	}
	return nil
}

func (a *Settings) Env(env Env) (string, error) {
	return a.env.Value(string(env)).String()
}

func (a *Settings) EnvExist(env Env) bool {
	v, err := a.env.Value(string(env)).String()
	return err == nil && v != ""
}

func checkDataPath(path string) error {
	if s, err := os.Stat(path); err != nil {
		return err
	} else if !s.IsDir() {
		return fmt.Errorf("%s: Is not a directory", path)
	}
	return nil
}

func loadFlags() Flags {
	var confPath string
	var dataPath string
	flag.StringVar(&confPath, "conf", "", "config path, eg: --conf config.yaml")
	flag.StringVar(&dataPath, "data", ".", "data path, eg: --data /opt/librarian/data")
	flag.Parse()
	return Flags{
		ConfPath: confPath,
		DataPath: dataPath,
	}
}

func loadEnv() (config.Config, error) {
	c := config.New(
		config.WithSource(
			env.NewSource(),
		),
	)

	if err := c.Load(); err != nil {
		return nil, err
	}
	return c, nil
}

func initLogger(id, name, version, dataPath string, logLevel libzap.Level, useStdout bool) {
	var l *zap.Logger
	if useStdout {
		l = zap.NewLogger(libzap.NewStdout(logLevel))
	} else {
		l = zap.NewLogger(libzap.New(dataPath, logLevel))
	}
	logger := log.With(l,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", name,
		"service.version", version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	log.SetLogger(logger)
}

func GetLogger() log.Logger {
	fuzzyStr := "***"
	return log.NewFilter(log.GetLogger(),
		log.FilterFunc(
			func(level log.Level, keyvals ...interface{}) bool {
				for i := 0; i < len(keyvals); i += 2 {
					if keyvals[i] == "reason" && keyvals[i+1] == "UNAUTHORISED" {
						return true
					}
				}
				return false
			},
		),
		log.FilterFunc(
			func(level log.Level, keyvals ...interface{}) bool {
				for i := 0; i < len(keyvals); i++ {
					if strings.Contains(fmt.Sprint(keyvals[i]), "password") {
						if i%2 == 0 {
							keyvals[i+1] = fuzzyStr
						} else {
							keyvals[i] = fuzzyStr
						}
					}
				}
				return false
			},
		),
	)
}
