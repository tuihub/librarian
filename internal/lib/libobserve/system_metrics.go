package libobserve

import (
	"context"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type SystemMetricsCollector struct {
	meter metric.Meter

	cpuUsage metric.Float64Gauge

	memoryUsage     metric.Float64Gauge
	memoryAvailable metric.Float64Gauge

	diskUsage metric.Float64Gauge

	networkBytesSent metric.Float64Counter
	networkBytesRecv metric.Float64Counter

	stopCh chan struct{}
}

func NewSystemMetricsCollector() *SystemMetricsCollector {
	meter := otel.GetMeterProvider().Meter("system")

	cpuUsage, _ := meter.Float64Gauge(
		"system.cpu.usage",
		metric.WithDescription("System CPU usage percentage"),
		metric.WithUnit("percent"),
	)

	memoryUsage, _ := meter.Float64Gauge(
		"system.memory.usage",
		metric.WithDescription("System memory usage percentage"),
		metric.WithUnit("percent"),
	)

	memoryAvailable, _ := meter.Float64Gauge(
		"system.memory.available",
		metric.WithDescription("Available memory in GB"),
		metric.WithUnit("GB"),
	)

	diskUsage, _ := meter.Float64Gauge(
		"system.disk.usage",
		metric.WithDescription("Disk usage percentage"),
		metric.WithUnit("percent"),
	)

	networkBytesSent, _ := meter.Float64Counter(
		"system.network.bytes_sent",
		metric.WithDescription("Network bytes sent in MB"),
		metric.WithUnit("MB"),
	)

	networkBytesRecv, _ := meter.Float64Counter(
		"system.network.bytes_recv",
		metric.WithDescription("Network bytes received in MB"),
		metric.WithUnit("MB"),
	)

	return &SystemMetricsCollector{
		meter:            meter,
		cpuUsage:         cpuUsage,
		memoryUsage:      memoryUsage,
		memoryAvailable:  memoryAvailable,
		diskUsage:        diskUsage,
		networkBytesSent: networkBytesSent,
		networkBytesRecv: networkBytesRecv,
		stopCh:           make(chan struct{}),
	}
}

func (c *SystemMetricsCollector) Collect() {
	// CPU
	if cpuPercent, err := cpu.Percent(time.Second, false); err == nil && len(cpuPercent) > 0 {
		c.cpuUsage.Record(context.Background(), cpuPercent[0],
			metric.WithAttributes(attribute.String("type", "total")))
	}

	// Memory
	if memInfo, err := mem.VirtualMemory(); err == nil {
		c.memoryUsage.Record(context.Background(), memInfo.UsedPercent,
			metric.WithAttributes(attribute.String("type", "virtual")))

		c.memoryAvailable.Record(context.Background(),
			float64(memInfo.Available)/1024/1024/1024, //nolint:mnd // GB
			metric.WithAttributes(attribute.String("type", "virtual")))
	}

	// Disk
	if partitions, err := disk.Partitions(false); err == nil {
		for _, partition := range partitions {
			if usage, err1 := disk.Usage(partition.Mountpoint); err1 == nil {
				c.diskUsage.Record(context.Background(), usage.UsedPercent,
					metric.WithAttributes(attribute.String("mountpoint", partition.Mountpoint)))
			}
		}
	}

	// Network I/O
	if netIO, err := net.IOCounters(false); err == nil && len(netIO) > 0 {
		io := netIO[0]
		c.networkBytesSent.Add(context.Background(),
			float64(io.BytesSent)/1024/1024, //nolint:mnd // MB
			metric.WithAttributes(attribute.String("interface", io.Name)))

		c.networkBytesRecv.Add(context.Background(),
			float64(io.BytesRecv)/1024/1024, //nolint:mnd // MB
			metric.WithAttributes(attribute.String("interface", io.Name)))
	}
}

func (c *SystemMetricsCollector) Start(ctx context.Context) error {
	c.Collect()

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-c.stopCh:
			return nil
		case <-ticker.C:
			c.Collect()
		}
	}
}

func (c *SystemMetricsCollector) Stop(ctx context.Context) error {
	c.stopCh <- struct{}{}
	return nil
}
