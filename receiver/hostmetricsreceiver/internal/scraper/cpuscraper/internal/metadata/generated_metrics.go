// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/scraper"
	conventions "go.opentelemetry.io/collector/semconv/v1.9.0"
)

// AttributeState specifies the value state attribute.
type AttributeState int

const (
	_ AttributeState = iota
	AttributeStateIdle
	AttributeStateInterrupt
	AttributeStateNice
	AttributeStateSoftirq
	AttributeStateSteal
	AttributeStateSystem
	AttributeStateUser
	AttributeStateWait
)

// String returns the string representation of the AttributeState.
func (av AttributeState) String() string {
	switch av {
	case AttributeStateIdle:
		return "idle"
	case AttributeStateInterrupt:
		return "interrupt"
	case AttributeStateNice:
		return "nice"
	case AttributeStateSoftirq:
		return "softirq"
	case AttributeStateSteal:
		return "steal"
	case AttributeStateSystem:
		return "system"
	case AttributeStateUser:
		return "user"
	case AttributeStateWait:
		return "wait"
	}
	return ""
}

// MapAttributeState is a helper map of string to AttributeState attribute value.
var MapAttributeState = map[string]AttributeState{
	"idle":      AttributeStateIdle,
	"interrupt": AttributeStateInterrupt,
	"nice":      AttributeStateNice,
	"softirq":   AttributeStateSoftirq,
	"steal":     AttributeStateSteal,
	"system":    AttributeStateSystem,
	"user":      AttributeStateUser,
	"wait":      AttributeStateWait,
}

type metricSystemCPUFrequency struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.cpu.frequency metric with initial data.
func (m *metricSystemCPUFrequency) init() {
	m.data.SetName("system.cpu.frequency")
	m.data.SetDescription("Current frequency of the CPU core in Hz.")
	m.data.SetUnit("Hz")
	m.data.SetEmptyGauge()
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemCPUFrequency) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, cpuAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
	dp.Attributes().PutStr("cpu", cpuAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemCPUFrequency) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemCPUFrequency) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemCPUFrequency(cfg MetricConfig) metricSystemCPUFrequency {
	m := metricSystemCPUFrequency{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSystemCPULogicalCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.cpu.logical.count metric with initial data.
func (m *metricSystemCPULogicalCount) init() {
	m.data.SetName("system.cpu.logical.count")
	m.data.SetDescription("Number of available logical CPUs.")
	m.data.SetUnit("{cpu}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricSystemCPULogicalCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemCPULogicalCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemCPULogicalCount) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemCPULogicalCount(cfg MetricConfig) metricSystemCPULogicalCount {
	m := metricSystemCPULogicalCount{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSystemCPUPhysicalCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.cpu.physical.count metric with initial data.
func (m *metricSystemCPUPhysicalCount) init() {
	m.data.SetName("system.cpu.physical.count")
	m.data.SetDescription("Number of available physical CPUs.")
	m.data.SetUnit("{cpu}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricSystemCPUPhysicalCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemCPUPhysicalCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemCPUPhysicalCount) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemCPUPhysicalCount(cfg MetricConfig) metricSystemCPUPhysicalCount {
	m := metricSystemCPUPhysicalCount{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSystemCPUTime struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.cpu.time metric with initial data.
func (m *metricSystemCPUTime) init() {
	m.data.SetName("system.cpu.time")
	m.data.SetDescription("Total seconds each logical CPU spent on each mode.")
	m.data.SetUnit("s")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemCPUTime) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, cpuAttributeValue string, stateAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
	dp.Attributes().PutStr("cpu", cpuAttributeValue)
	dp.Attributes().PutStr("state", stateAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemCPUTime) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemCPUTime) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemCPUTime(cfg MetricConfig) metricSystemCPUTime {
	m := metricSystemCPUTime{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSystemCPUUtilization struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.cpu.utilization metric with initial data.
func (m *metricSystemCPUUtilization) init() {
	m.data.SetName("system.cpu.utilization")
	m.data.SetDescription("Difference in system.cpu.time since the last measurement per logical CPU, divided by the elapsed time (value in interval [0,1]).")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemCPUUtilization) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, cpuAttributeValue string, stateAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
	dp.Attributes().PutStr("cpu", cpuAttributeValue)
	dp.Attributes().PutStr("state", stateAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemCPUUtilization) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemCPUUtilization) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemCPUUtilization(cfg MetricConfig) metricSystemCPUUtilization {
	m := metricSystemCPUUtilization{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user config.
type MetricsBuilder struct {
	config                       MetricsBuilderConfig // config of the metrics builder.
	startTime                    pcommon.Timestamp    // start time that will be applied to all recorded data points.
	metricsCapacity              int                  // maximum observed number of metrics per resource.
	metricsBuffer                pmetric.Metrics      // accumulates metrics data before emitting.
	buildInfo                    component.BuildInfo  // contains version information.
	metricSystemCPUFrequency     metricSystemCPUFrequency
	metricSystemCPULogicalCount  metricSystemCPULogicalCount
	metricSystemCPUPhysicalCount metricSystemCPUPhysicalCount
	metricSystemCPUTime          metricSystemCPUTime
	metricSystemCPUUtilization   metricSystemCPUUtilization
}

// MetricBuilderOption applies changes to default metrics builder.
type MetricBuilderOption interface {
	apply(*MetricsBuilder)
}

type metricBuilderOptionFunc func(mb *MetricsBuilder)

func (mbof metricBuilderOptionFunc) apply(mb *MetricsBuilder) {
	mbof(mb)
}

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) MetricBuilderOption {
	return metricBuilderOptionFunc(func(mb *MetricsBuilder) {
		mb.startTime = startTime
	})
}
func NewMetricsBuilder(mbc MetricsBuilderConfig, settings scraper.Settings, options ...MetricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		config:                       mbc,
		startTime:                    pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                pmetric.NewMetrics(),
		buildInfo:                    settings.BuildInfo,
		metricSystemCPUFrequency:     newMetricSystemCPUFrequency(mbc.Metrics.SystemCPUFrequency),
		metricSystemCPULogicalCount:  newMetricSystemCPULogicalCount(mbc.Metrics.SystemCPULogicalCount),
		metricSystemCPUPhysicalCount: newMetricSystemCPUPhysicalCount(mbc.Metrics.SystemCPUPhysicalCount),
		metricSystemCPUTime:          newMetricSystemCPUTime(mbc.Metrics.SystemCPUTime),
		metricSystemCPUUtilization:   newMetricSystemCPUUtilization(mbc.Metrics.SystemCPUUtilization),
	}

	for _, op := range options {
		op.apply(mb)
	}
	return mb
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
}

// ResourceMetricsOption applies changes to provided resource metrics.
type ResourceMetricsOption interface {
	apply(pmetric.ResourceMetrics)
}

type resourceMetricsOptionFunc func(pmetric.ResourceMetrics)

func (rmof resourceMetricsOptionFunc) apply(rm pmetric.ResourceMetrics) {
	rmof(rm)
}

// WithResource sets the provided resource on the emitted ResourceMetrics.
// It's recommended to use ResourceBuilder to create the resource.
func WithResource(res pcommon.Resource) ResourceMetricsOption {
	return resourceMetricsOptionFunc(func(rm pmetric.ResourceMetrics) {
		res.CopyTo(rm.Resource())
	})
}

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return resourceMetricsOptionFunc(func(rm pmetric.ResourceMetrics) {
		var dps pmetric.NumberDataPointSlice
		metrics := rm.ScopeMetrics().At(0).Metrics()
		for i := 0; i < metrics.Len(); i++ {
			switch metrics.At(i).Type() {
			case pmetric.MetricTypeGauge:
				dps = metrics.At(i).Gauge().DataPoints()
			case pmetric.MetricTypeSum:
				dps = metrics.At(i).Sum().DataPoints()
			}
			for j := 0; j < dps.Len(); j++ {
				dps.At(j).SetStartTimestamp(start)
			}
		}
	})
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceMetricsOption arguments.
func (mb *MetricsBuilder) EmitForResource(options ...ResourceMetricsOption) {
	rm := pmetric.NewResourceMetrics()
	rm.SetSchemaUrl(conventions.SchemaURL)
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/cpuscraper")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricSystemCPUFrequency.emit(ils.Metrics())
	mb.metricSystemCPULogicalCount.emit(ils.Metrics())
	mb.metricSystemCPUPhysicalCount.emit(ils.Metrics())
	mb.metricSystemCPUTime.emit(ils.Metrics())
	mb.metricSystemCPUUtilization.emit(ils.Metrics())

	for _, op := range options {
		op.apply(rm)
	}

	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user config, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(options ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(options...)
	metrics := mb.metricsBuffer
	mb.metricsBuffer = pmetric.NewMetrics()
	return metrics
}

// RecordSystemCPUFrequencyDataPoint adds a data point to system.cpu.frequency metric.
func (mb *MetricsBuilder) RecordSystemCPUFrequencyDataPoint(ts pcommon.Timestamp, val float64, cpuAttributeValue string) {
	mb.metricSystemCPUFrequency.recordDataPoint(mb.startTime, ts, val, cpuAttributeValue)
}

// RecordSystemCPULogicalCountDataPoint adds a data point to system.cpu.logical.count metric.
func (mb *MetricsBuilder) RecordSystemCPULogicalCountDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricSystemCPULogicalCount.recordDataPoint(mb.startTime, ts, val)
}

// RecordSystemCPUPhysicalCountDataPoint adds a data point to system.cpu.physical.count metric.
func (mb *MetricsBuilder) RecordSystemCPUPhysicalCountDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricSystemCPUPhysicalCount.recordDataPoint(mb.startTime, ts, val)
}

// RecordSystemCPUTimeDataPoint adds a data point to system.cpu.time metric.
func (mb *MetricsBuilder) RecordSystemCPUTimeDataPoint(ts pcommon.Timestamp, val float64, cpuAttributeValue string, stateAttributeValue AttributeState) {
	mb.metricSystemCPUTime.recordDataPoint(mb.startTime, ts, val, cpuAttributeValue, stateAttributeValue.String())
}

// RecordSystemCPUUtilizationDataPoint adds a data point to system.cpu.utilization metric.
func (mb *MetricsBuilder) RecordSystemCPUUtilizationDataPoint(ts pcommon.Timestamp, val float64, cpuAttributeValue string, stateAttributeValue AttributeState) {
	mb.metricSystemCPUUtilization.recordDataPoint(mb.startTime, ts, val, cpuAttributeValue, stateAttributeValue.String())
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...MetricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op.apply(mb)
	}
}
