package libobserve

import (
	"context"
	"os"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
	"github.com/shirou/gopsutil/v4/process"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

const defaultSystemMetricsInterval = 5 * time.Second

type SystemMetricsCollector struct {
	meter metric.Meter

	// System metrics
	cpuUsage metric.Float64Gauge

	memoryUsage     metric.Float64Gauge
	memoryAvailable metric.Float64Gauge

	networkBytesSent metric.Float64Counter
	networkBytesRecv metric.Float64Counter

	// Process metrics
	processCPUUsage metric.Float64Gauge

	processMemoryUsage metric.Float64Gauge
	processMemoryHeap  metric.Float64Gauge
	processMemoryStack metric.Float64Gauge
	processMemoryTotal metric.Float64Gauge

	processDiskWrite metric.Float64Gauge
	processDiskRead  metric.Float64Gauge

	// Go runtime metrics
	processGoroutines      metric.Int64Gauge
	processGCCount         metric.Int64Counter
	processGCPauseDuration metric.Float64Histogram

	// Process handle for CPU monitoring
	currentProcess *process.Process

	stopCh chan struct{}
}

func NewSystemMetricsCollector() *SystemMetricsCollector { //nolint:funlen // TODO
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

	processCPUUsage, _ := meter.Float64Gauge(
		"process.cpu.usage",
		metric.WithDescription("Process CPU usage percentage"),
		metric.WithUnit("percent"),
	)

	processMemoryUsage, _ := meter.Float64Gauge(
		"process.memory.usage",
		metric.WithDescription("Process memory usage percentage"),
		metric.WithUnit("percent"),
	)

	processMemoryHeap, _ := meter.Float64Gauge(
		"process.memory.heap",
		metric.WithDescription("Process heap memory usage in bytes"),
		metric.WithUnit("bytes"),
	)

	processMemoryStack, _ := meter.Float64Gauge(
		"process.memory.stack",
		metric.WithDescription("Process stack memory usage in bytes"),
		metric.WithUnit("bytes"),
	)

	processMemoryTotal, _ := meter.Float64Gauge(
		"process.memory.total",
		metric.WithDescription("Total process memory usage in bytes"),
		metric.WithUnit("bytes"),
	)

	processDiskWrite, _ := meter.Float64Gauge(
		"process.disk.write",
		metric.WithDescription("Process disk write in MB"),
		metric.WithUnit("MB"),
	)

	processDiskRead, _ := meter.Float64Gauge(
		"process.disk.read",
		metric.WithDescription("Process disk read in MB"),
		metric.WithUnit("MB"),
	)

	processGoroutines, _ := meter.Int64Gauge(
		"process.goroutines",
		metric.WithDescription("Number of goroutines"),
	)

	processGCCount, _ := meter.Int64Counter(
		"process.gc.count",
		metric.WithDescription("Number of garbage collection cycles"),
	)

	processGCPauseDuration, _ := meter.Float64Histogram(
		"process.gc.pause_duration",
		metric.WithDescription("Garbage collection pause duration"),
		metric.WithUnit("seconds"),
	)

	return &SystemMetricsCollector{
		meter:                  meter,
		cpuUsage:               cpuUsage,
		memoryUsage:            memoryUsage,
		memoryAvailable:        memoryAvailable,
		networkBytesSent:       networkBytesSent,
		networkBytesRecv:       networkBytesRecv,
		processCPUUsage:        processCPUUsage,
		processMemoryUsage:     processMemoryUsage,
		processMemoryHeap:      processMemoryHeap,
		processMemoryStack:     processMemoryStack,
		processMemoryTotal:     processMemoryTotal,
		processDiskWrite:       processDiskWrite,
		processDiskRead:        processDiskRead,
		processGoroutines:      processGoroutines,
		processGCCount:         processGCCount,
		processGCPauseDuration: processGCPauseDuration,
		currentProcess:         nil,
		stopCh:                 make(chan struct{}),
	}
}

func (c *SystemMetricsCollector) Collect(ctx context.Context) {
	// Get current process
	if c.currentProcess == nil {
		if p, err := process.NewProcess(int32(os.Getpid())); err == nil { //nolint:gosec // Get current process ID
			c.currentProcess = p
		}
	}
	// CPU
	if cpuPercent, err := cpu.Percent(time.Second, false); err == nil && len(cpuPercent) > 0 {
		c.cpuUsage.Record(ctx, cpuPercent[0],
			metric.WithAttributes(attribute.String("type", "total")))
	}
	if c.currentProcess != nil {
		if cpuPercent, err := c.currentProcess.CPUPercent(); err == nil {
			numCPU := float64(runtime.NumCPU())
			c.processCPUUsage.Record(ctx, cpuPercent/numCPU)
		}
	}

	// Memory
	if memInfo, err := mem.VirtualMemory(); err == nil {
		c.memoryUsage.Record(ctx, memInfo.UsedPercent,
			metric.WithAttributes(attribute.String("type", "virtual")))

		c.memoryAvailable.Record(ctx,
			float64(memInfo.Available)/1024/1024/1024, //nolint:mnd // GB
			metric.WithAttributes(attribute.String("type", "virtual")))
		if c.currentProcess != nil {
			if processMemInfo, err1 := c.currentProcess.MemoryInfo(); err1 == nil {
				memoryUsagePercent := float64(
					processMemInfo.RSS,
				) / float64(
					memInfo.Total,
				) * 100 //nolint:mnd // Percentage
				c.processMemoryUsage.Record(ctx, memoryUsagePercent,
					metric.WithAttributes(attribute.String("type", "rss")))
			}
		}
	}

	// Disk
	if c.currentProcess != nil {
		if ioCounters, err := c.currentProcess.IOCounters(); err == nil {
			c.processDiskWrite.Record(ctx,
				float64(ioCounters.WriteBytes)/1024/1024, //nolint:mnd // MB
				metric.WithAttributes(attribute.String("type", "write")))

			c.processDiskRead.Record(ctx,
				float64(ioCounters.ReadBytes)/1024/1024, //nolint:mnd // MB
				metric.WithAttributes(attribute.String("type", "read")))
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

	// Process metrics
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	c.processMemoryHeap.Record(ctx, float64(memStats.HeapAlloc),
		metric.WithAttributes(attribute.String("type", "heap")))

	c.processMemoryStack.Record(ctx, float64(memStats.StackSys),
		metric.WithAttributes(attribute.String("type", "stack")))

	c.processMemoryTotal.Record(ctx, float64(memStats.Alloc),
		metric.WithAttributes(attribute.String("type", "total")))

	c.processGoroutines.Record(ctx, int64(runtime.NumGoroutine()))

	// GC metrics from MemStats
	c.processGCCount.Add(ctx, int64(memStats.NumGC))

	// Record the pause duration using PauseTotalNs
	c.processGCPauseDuration.Record(
		ctx,
		float64(memStats.PauseTotalNs)/1e9, //nolint:mnd // Convert to seconds
	)
}

func (c *SystemMetricsCollector) Start(ctx context.Context) error {
	c.Collect(ctx)

	ticker := time.NewTicker(defaultSystemMetricsInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-c.stopCh:
			return nil
		case <-ticker.C:
			c.Collect(ctx)
		}
	}
}

func (c *SystemMetricsCollector) Stop(ctx context.Context) error {
	c.stopCh <- struct{}{}
	return nil
}
