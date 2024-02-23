package libsentry

import (
	"github.com/tuihub/librarian/internal/conf"

	"github.com/getsentry/sentry-go"
	sentryotel "github.com/getsentry/sentry-go/otel"
	"go.opentelemetry.io/otel"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

func InitSentry(c *conf.Sentry) error {
	if c == nil || c.GetDsn() == "" {
		return nil
	}
	// Initialize the Sentry SDK.
	err := sentry.Init(sentry.ClientOptions{ //nolint:exhaustruct // no need
		Dsn:                c.GetDsn(),
		AttachStacktrace:   true,
		EnableTracing:      true,
		TracesSampleRate:   1.0,
		ProfilesSampleRate: 1.0,
		Environment:        c.GetEnvironment(),
	})
	if err != nil {
		return err
	}
	// Initialize the OpenTelemetry tracing integration.
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithSpanProcessor(sentryotel.NewSentrySpanProcessor()),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(sentryotel.NewSentryPropagator())
	return nil
}
