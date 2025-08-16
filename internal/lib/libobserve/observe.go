package libobserve

import (
	"context"

	"github.com/tuihub/librarian/internal/conf"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewObserve,
	NewMemoryMetricExporter,
)

type Observe struct {
	systemCollector *SystemMetricsCollector
	memoryExporter  *MemoryMetricExporter
}

func NewObserve(
	c *conf.OpenTelemetry,
	memoryExporter *MemoryMetricExporter,
) (*Observe, error) {
	err := InitOTEL(c, memoryExporter)
	if err != nil {
		return nil, err
	}
	return &Observe{
		systemCollector: NewSystemMetricsCollector(),
		memoryExporter:  memoryExporter,
	}, nil
}

func (o *Observe) Start(ctx context.Context) error {
	return o.systemCollector.Start(ctx)
}

func (o *Observe) Stop(ctx context.Context) error {
	if o.systemCollector != nil {
		return o.systemCollector.Stop(ctx)
	}
	return nil
}

func (o *Observe) GetMetrics(timeRange string) map[string][]MetricPoint {
	return o.memoryExporter.GetMetrics(timeRange)
}

func (o *Observe) GetLatestMetrics() map[string]MetricPoint {
	return o.memoryExporter.GetLatestMetrics()
}
