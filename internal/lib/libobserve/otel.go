package libobserve

import (
	"context"
	"errors"
	"strings"

	"github.com/tuihub/librarian/internal/conf"

	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitOTEL(
	c *conf.OpenTelemetry,
	memoryMetricExporter *MemoryMetricExporter,
) error {
	if c == nil {
		return nil
	}
	ctx := context.Background()
	var conn *grpc.ClientConn
	switch c.Protocol {
	case conf.OpenTelemetryProtocolDisable:
		break
	case conf.OpenTelemetryProtocolGRPC:
		headers, err := parseHeaders(c.Headers)
		if err != nil {
			return err
		}
		var grpcOpts []grpc.DialOption
		if c.GRPCInsecure {
			grpcOpts = append(grpcOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
			grpcOpts = append(grpcOpts, grpc.WithPerRPCCredentials(globalHeaders{
				headers: headers,
				secure:  !c.GRPCInsecure,
			}))
		}
		conn, err = grpc.NewClient(c.Endpoint, grpcOpts...)
		if err != nil {
			return err
		}
	case conf.OpenTelemetryProtocolHTTP:
		break
	default:
		return errors.New("invalid protocol")
	}

	_, err := newTraceProvider(ctx, c, conn)
	if err != nil {
		return err
	}
	_, err = newMeterProvider(ctx, c, conn, memoryMetricExporter)
	if err != nil {
		return err
	}
	_, err = newLoggerProvider(ctx, c, conn)
	if err != nil {
		return err
	}
	newPropagator()
	return nil
}

func newPropagator() propagation.TextMapPropagator {
	prop := propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
	otel.SetTextMapPropagator(prop)
	return prop
}

func newTraceProvider(
	ctx context.Context,
	c *conf.OpenTelemetry,
	conn *grpc.ClientConn,
) (*trace.TracerProvider, error) {
	options := []trace.TracerProviderOption{
		trace.WithSampler(trace.AlwaysSample()),
	}
	switch c.Protocol {
	case conf.OpenTelemetryProtocolDisable:
		break
	case conf.OpenTelemetryProtocolHTTP:
		headers, err := parseHeaders(c.Headers)
		if err != nil {
			return nil, err
		}
		exporter, err := otlptracehttp.New(ctx,
			otlptracehttp.WithEndpointURL(c.Endpoint),
			otlptracehttp.WithHeaders(headers),
		)
		if err != nil {
			return nil, err
		}
		options = append(options, trace.WithBatcher(exporter))
	case conf.OpenTelemetryProtocolGRPC:
		exporter, err := otlptracegrpc.New(ctx,
			otlptracegrpc.WithGRPCConn(conn),
		)
		if err != nil {
			return nil, err
		}
		options = append(options, trace.WithBatcher(exporter))
	}

	tp := trace.NewTracerProvider(options...)

	otel.SetTracerProvider(tp)
	return tp, nil
}

func newMeterProvider(
	ctx context.Context,
	c *conf.OpenTelemetry,
	conn *grpc.ClientConn,
	memExporter *MemoryMetricExporter,
) (*metric.MeterProvider, error) {
	view := metrics.DefaultSecondsHistogramView(metrics.DefaultServerSecondsHistogramName)
	options := []metric.Option{metric.WithView(view)}

	if c.EnableMemoryMetrics {
		options = append(options, metric.WithReader(metric.NewPeriodicReader(memExporter)))
	}

	switch c.Protocol {
	case conf.OpenTelemetryProtocolDisable:
		break
	case conf.OpenTelemetryProtocolHTTP:
		headers, err := parseHeaders(c.Headers)
		if err != nil {
			return nil, err
		}
		externalExporter, err := otlpmetrichttp.New(ctx,
			otlpmetrichttp.WithEndpointURL(c.Endpoint),
			otlpmetrichttp.WithHeaders(headers),
		)
		if err != nil {
			return nil, err
		}
		options = append(options, metric.WithReader(metric.NewPeriodicReader(externalExporter)))
	case conf.OpenTelemetryProtocolGRPC:
		externalExporter, err := otlpmetricgrpc.New(ctx,
			otlpmetricgrpc.WithGRPCConn(conn),
		)
		if err != nil {
			return nil, err
		}
		options = append(options, metric.WithReader(metric.NewPeriodicReader(externalExporter)))
	default:
		return nil, errors.New("invalid protocol for metrics")
	}

	mp := metric.NewMeterProvider(options...)
	otel.SetMeterProvider(mp)
	return mp, nil
}

func newLoggerProvider(ctx context.Context, c *conf.OpenTelemetry, conn *grpc.ClientConn) (*log.LoggerProvider, error) {
	var options []log.LoggerProviderOption

	switch c.Protocol {
	case conf.OpenTelemetryProtocolDisable:
		break
	case conf.OpenTelemetryProtocolHTTP:
		headers, err := parseHeaders(c.Headers)
		if err != nil {
			return nil, err
		}
		exporter, err := otlploghttp.New(ctx,
			otlploghttp.WithEndpointURL(c.Endpoint),
			otlploghttp.WithHeaders(headers),
		)
		if err != nil {
			return nil, err
		}
		options = append(options, log.WithProcessor(log.NewBatchProcessor(exporter)))
	case conf.OpenTelemetryProtocolGRPC:
		exporter, err := otlploggrpc.New(ctx,
			otlploggrpc.WithGRPCConn(conn),
		)
		if err != nil {
			return nil, err
		}
		options = append(options, log.WithProcessor(log.NewBatchProcessor(exporter)))
	default:
		return nil, errors.New("invalid protocol for logging")
	}

	lp := log.NewLoggerProvider(options...)
	global.SetLoggerProvider(lp)
	return lp, nil
}

func parseHeaders(headersRaw string) (map[string]string, error) {
	headersMap := make(map[string]string)
	if headersRaw == "" {
		return headersMap, nil
	}
	headers := strings.Split(headersRaw, ",")
	for _, header := range headers {
		parts := strings.SplitN(header, "=", 2) //nolint:mnd // no need
		if len(parts) != 2 {                    //nolint:mnd // no need
			return nil, errors.New("invalid header format: " + header)
		}
		headersMap[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}
	return headersMap, nil
}

type globalHeaders struct {
	headers map[string]string
	secure  bool
}

func (g globalHeaders) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	return g.headers, nil
}

func (g globalHeaders) RequireTransportSecurity() bool {
	return g.secure
}
