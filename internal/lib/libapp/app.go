package libapp

import (
	"fmt"
	"strings"

	"github.com/tuihub/librarian/internal/lib/libzap"

	"github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
)

func InitLogger(id, name, version string) {
	logger := log.With(zap.NewLogger(libzap.NewDefaultLogger()),
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
		log.FilterKey("password"),
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

func LoadConfig(flagconf string, conf interface{}) {
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	if err := c.Scan(conf); err != nil {
		panic(err)
	}
}
