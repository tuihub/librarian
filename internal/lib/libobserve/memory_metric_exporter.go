package libobserve

import (
	"context"
	"sync"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
)

type MemoryMetricExporter struct {
	metrics   map[string][]MetricPoint
	maxSize   int
	maxMemory int64
	mu        sync.RWMutex
}

type MetricPoint struct {
	Name        string            `json:"name"`
	Value       float64           `json:"value"`
	Timestamp   time.Time         `json:"timestamp"`
	Labels      map[string]string `json:"labels"`
	Description string            `json:"description,omitempty"`
	Unit        string            `json:"unit,omitempty"`
}

func NewMemoryMetricExporter() *MemoryMetricExporter {
	return &MemoryMetricExporter{
		metrics:   make(map[string][]MetricPoint),
		maxSize:   1000,             //nolint:mnd // TODO
		maxMemory: 10 * 1024 * 1024, //nolint:mnd // TODO
		mu:        sync.RWMutex{},
	}
}

func (e *MemoryMetricExporter) Export(ctx context.Context, rm *metricdata.ResourceMetrics) error {
	for _, scopeMetrics := range rm.ScopeMetrics {
		for _, metrics := range scopeMetrics.Metrics {
			e.processMetric(metrics)
		}
	}
	return nil
}

func (e *MemoryMetricExporter) Aggregation(kind metric.InstrumentKind) metric.Aggregation {
	return metric.DefaultAggregationSelector(kind)
}

func (e *MemoryMetricExporter) Temporality(kind metric.InstrumentKind) metricdata.Temporality {
	return metric.DefaultTemporalitySelector(kind)
}

func (e *MemoryMetricExporter) processMetric(m metricdata.Metrics) {
	switch data := m.Data.(type) {
	case metricdata.Gauge[float64]:
		e.processGauge(m, data)
	case metricdata.Sum[float64]:
		e.processSum(m, data)
	case metricdata.Histogram[float64]:
		e.processHistogram(m, data)
	}
}

func (e *MemoryMetricExporter) processGauge(m metricdata.Metrics, data metricdata.Gauge[float64]) {
	for _, point := range data.DataPoints {
		e.addMetricPoint(MetricPoint{
			Name:        m.Name,
			Value:       point.Value,
			Timestamp:   point.Time,
			Labels:      attributesToMap(point.Attributes),
			Description: m.Description,
			Unit:        m.Unit,
		})
	}
}

func (e *MemoryMetricExporter) processSum(m metricdata.Metrics, data metricdata.Sum[float64]) {
	for _, point := range data.DataPoints {
		e.addMetricPoint(MetricPoint{
			Name:        m.Name,
			Value:       point.Value,
			Timestamp:   point.Time,
			Labels:      attributesToMap(point.Attributes),
			Description: m.Description,
			Unit:        m.Unit,
		})
	}
}

func (e *MemoryMetricExporter) processHistogram(m metricdata.Metrics, data metricdata.Histogram[float64]) {
	for _, point := range data.DataPoints {
		if point.Count > 0 {
			avg := point.Sum / float64(point.Count)
			e.addMetricPoint(MetricPoint{
				Name:        m.Name + "_avg",
				Value:       avg,
				Timestamp:   point.Time,
				Labels:      attributesToMap(point.Attributes),
				Description: m.Description + " (average)",
				Unit:        m.Unit,
			})
		}
	}
}

func (e *MemoryMetricExporter) addMetricPoint(point MetricPoint) {
	e.mu.Lock()
	defer e.mu.Unlock()

	// Ensure the timestamp is set
	if point.Timestamp.IsZero() {
		point.Timestamp = time.Now()
	}

	// Ensure the metric slice exists
	if e.metrics[point.Name] == nil {
		e.metrics[point.Name] = make([]MetricPoint, 0, e.maxSize)
	}

	// Append the new metric point
	e.metrics[point.Name] = append(e.metrics[point.Name], point)

	// If the size exceeds the maximum, remove the oldest
	if len(e.metrics[point.Name]) > e.maxSize {
		e.metrics[point.Name] = e.metrics[point.Name][1:]
	}

	// Check memory usage, if it exceeds the limit, clean up the oldest metrics
	e.cleanupIfNeeded()
}

func (e *MemoryMetricExporter) cleanupIfNeeded() {
	// Simple estimation: each metric point takes about 300 bytes
	estimatedSize := int64(len(e.metrics) * e.maxSize * 300) //nolint:mnd // TODO

	if estimatedSize > e.maxMemory {
		// Clear out the oldest half of the data for each metric
		for name, values := range e.metrics {
			if len(values) > 1 {
				keepCount := len(values) / 2 //nolint:mnd // TODO
				e.metrics[name] = values[len(values)-keepCount:]
			}
		}
	}
}

func (e *MemoryMetricExporter) GetMetrics(timeRange string) map[string][]MetricPoint {
	e.mu.RLock()
	defer e.mu.RUnlock()

	var cutoff time.Time
	now := time.Now()

	switch timeRange {
	case "5m":
		cutoff = now.Add(-5 * time.Minute)
	case "1h":
		cutoff = now.Add(-1 * time.Hour)
	case "24h":
		cutoff = now.Add(-24 * time.Hour)
	default:
		cutoff = now.Add(-5 * time.Minute) // 默认5分钟
	}

	result := make(map[string][]MetricPoint)
	for name, values := range e.metrics {
		var filtered []MetricPoint
		for _, v := range values {
			if v.Timestamp.After(cutoff) {
				filtered = append(filtered, v)
			}
		}
		if len(filtered) > 0 {
			result[name] = filtered
		}
	}

	return result
}

func (e *MemoryMetricExporter) GetLatestMetrics() map[string]MetricPoint {
	e.mu.RLock()
	defer e.mu.RUnlock()

	result := make(map[string]MetricPoint)
	for name, values := range e.metrics {
		if len(values) > 0 {
			result[name] = values[len(values)-1]
		}
	}

	return result
}

func attributesToMap(attrs attribute.Set) map[string]string {
	result := make(map[string]string)
	for _, kv := range attrs.ToSlice() {
		result[string(kv.Key)] = kv.Value.Emit()
	}
	return result
}

func (e *MemoryMetricExporter) Shutdown(ctx context.Context) error {
	return nil
}

func (e *MemoryMetricExporter) ForceFlush(ctx context.Context) error {
	return nil
}
