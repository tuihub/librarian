package libobserve

import (
	"context"
	"errors"
	"strings"

	"github.com/tuihub/librarian/internal/conf"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
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

const (
	protocolGrpc = "grpc"
	protocolHTTP = "http"
)

func InitOTEL(c *conf.OTLP) error {
	if c == nil {
		return nil
	}
	headers, err := parseHeaders(c.GetHeaders())
	if err != nil {
		return err
	}
	ctx := context.Background()
	var conn *grpc.ClientConn
	switch c.GetProtocol() {
	case protocolGrpc:
		var grpcOpts []grpc.DialOption
		if c.GetGrpcInsecure() {
			grpcOpts = append(grpcOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
			grpcOpts = append(grpcOpts, grpc.WithPerRPCCredentials(globalHeaders{
				headers: headers,
				secure:  !c.GetGrpcInsecure(),
			}))
		}
		conn, err = grpc.NewClient(c.GetEndpoint(), grpcOpts...)
		if err != nil {
			return err
		}
	case protocolHTTP:
	default:
		return errors.New("invalid protocol")
	}

	_, err = newTraceProvider(ctx, c, conn)
	if err != nil {
		return err
	}
	_, err = newMeterProvider(ctx, c, conn)
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

func newTraceProvider(ctx context.Context, c *conf.OTLP, conn *grpc.ClientConn) (*trace.TracerProvider, error) {
	var exporter *otlptrace.Exporter
	headers, err := parseHeaders(c.GetHeaders())
	if err != nil {
		return nil, err
	}
	if c.GetProtocol() == protocolHTTP {
		exporter, err = otlptracehttp.New(ctx,
			otlptracehttp.WithEndpointURL(c.GetEndpoint()),
			otlptracehttp.WithHeaders(headers),
		)
	} else if c.GetProtocol() == protocolGrpc {
		exporter, err = otlptracegrpc.New(ctx,
			otlptracegrpc.WithGRPCConn(conn),
		)
	}
	if err != nil {
		return nil, err
	}

	tp := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithBatcher(exporter),
	)
	otel.SetTracerProvider(tp)
	return tp, nil
}

func newMeterProvider(ctx context.Context, c *conf.OTLP, conn *grpc.ClientConn) (*metric.MeterProvider, error) {
	var exporter metric.Exporter
	headers, err := parseHeaders(c.GetHeaders())
	if err != nil {
		return nil, err
	}
	if c.GetProtocol() == protocolHTTP {
		exporter, err = otlpmetrichttp.New(ctx,
			otlpmetrichttp.WithEndpointURL(c.GetEndpoint()),
			otlpmetrichttp.WithHeaders(headers),
		)
	} else if c.GetProtocol() == protocolGrpc {
		exporter, err = otlpmetricgrpc.New(ctx,
			otlpmetricgrpc.WithGRPCConn(conn),
		)
	}
	if err != nil {
		return nil, err
	}

	mp := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(exporter)),
	)
	otel.SetMeterProvider(mp)
	return mp, nil
}

func newLoggerProvider(ctx context.Context, c *conf.OTLP, conn *grpc.ClientConn) (*log.LoggerProvider, error) {
	var exporter log.Exporter
	headers, err := parseHeaders(c.GetHeaders())
	if err != nil {
		return nil, err
	}
	if c.GetProtocol() == protocolHTTP {
		exporter, err = otlploghttp.New(ctx,
			otlploghttp.WithEndpointURL(c.GetEndpoint()),
			otlploghttp.WithHeaders(headers),
		)
	} else if c.GetProtocol() == protocolGrpc {
		exporter, err = otlploggrpc.New(ctx,
			otlploggrpc.WithGRPCConn(conn),
		)
	}
	if err != nil {
		return nil, err
	}

	lp := log.NewLoggerProvider(
		log.WithProcessor(log.NewBatchProcessor(exporter)),
	)
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
		parts := strings.SplitN(header, "=", 2) //nolint:gomnd // no need
		if len(parts) != 2 {                    //nolint:gomnd // no need
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
