package libsentry

import (
	"github.com/tuihub/librarian/internal/conf"

	"github.com/getsentry/sentry-go"
)

func InitSentry(c *conf.Sentry) error {
	if c == nil || c.GetDsn() == "" {
		return nil
	}
	// Initialize the Sentry SDK.
	err := sentry.Init(sentry.ClientOptions{ //nolint:exhaustruct // no need
		Dsn:              c.GetDsn(),
		AttachStacktrace: true,
		EnableTracing:    true,
		TracesSampleRate: 1.0,
		SampleRate:       1.0,
		Environment:      c.GetEnvironment(),
	})
	if err != nil {
		return err
	}
	return nil
}
