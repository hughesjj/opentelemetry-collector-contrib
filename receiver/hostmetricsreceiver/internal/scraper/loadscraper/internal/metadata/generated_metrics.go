// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	conventions "go.opentelemetry.io/collector/semconv/v1.9.0"
)

// MetricSettings provides common settings for a particular metric.
type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (ms *MetricSettings) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms, confmap.WithErrorUnused())
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsSettings provides settings for hostmetricsreceiver/load metrics.
type MetricsSettings struct {
	SystemCPULoadAverage15m MetricSettings `mapstructure:"system.cpu.load_average.15m"`
	SystemCPULoadAverage1m  MetricSettings `mapstructure:"system.cpu.load_average.1m"`
	SystemCPULoadAverage5m  MetricSettings `mapstructure:"system.cpu.load_average.5m"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		SystemCPULoadAverage15m: MetricSettings{
			Enabled: true,
		},
		SystemCPULoadAverage1m: MetricSettings{
			Enabled: true,
		},
		SystemCPULoadAverage5m: MetricSettings{
			Enabled: true,
		},
	}
}

// ResourceAttributeSettings provides common settings for a particular metric.
type ResourceAttributeSettings struct {
	Enabled bool `mapstructure:"enabled"`
}

func (ras *ResourceAttributeSettings) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ras, confmap.WithErrorUnused())
	if err != nil {
		return err
	}
	return nil
}

// ResourceAttributesSettings provides settings for hostmetricsreceiver/load metrics.
type ResourceAttributesSettings struct {
}

func DefaultResourceAttributesSettings() ResourceAttributesSettings {
	return ResourceAttributesSettings{}
}

type metricSystemCPULoadAverage15m struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.cpu.load_average.15m metric with initial data.
func (m *metricSystemCPULoadAverage15m) init() {
	m.data.SetName("system.cpu.load_average.15m")
	m.data.SetDescription("Average CPU Load over 15 minutes.")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
}

func (m *metricSystemCPULoadAverage15m) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemCPULoadAverage15m) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemCPULoadAverage15m) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemCPULoadAverage15m(settings MetricSettings) metricSystemCPULoadAverage15m {
	m := metricSystemCPULoadAverage15m{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSystemCPULoadAverage1m struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.cpu.load_average.1m metric with initial data.
func (m *metricSystemCPULoadAverage1m) init() {
	m.data.SetName("system.cpu.load_average.1m")
	m.data.SetDescription("Average CPU Load over 1 minute.")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
}

func (m *metricSystemCPULoadAverage1m) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemCPULoadAverage1m) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemCPULoadAverage1m) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemCPULoadAverage1m(settings MetricSettings) metricSystemCPULoadAverage1m {
	m := metricSystemCPULoadAverage1m{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricSystemCPULoadAverage5m struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.cpu.load_average.5m metric with initial data.
func (m *metricSystemCPULoadAverage5m) init() {
	m.data.SetName("system.cpu.load_average.5m")
	m.data.SetDescription("Average CPU Load over 5 minutes.")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
}

func (m *metricSystemCPULoadAverage5m) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemCPULoadAverage5m) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemCPULoadAverage5m) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemCPULoadAverage5m(settings MetricSettings) metricSystemCPULoadAverage5m {
	m := metricSystemCPULoadAverage5m{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilderConfig is a structural subset of an otherwise 1-1 copy of metadata.yaml
type MetricsBuilderConfig struct {
	Metrics            MetricsSettings            `mapstructure:"metrics"`
	ResourceAttributes ResourceAttributesSettings `mapstructure:"resource_attributes"`
}

func (mbc *MetricsBuilderConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(mbc, confmap.WithErrorUnused())
	if err != nil {
		return err
	}
	return nil
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                     pcommon.Timestamp   // start time that will be applied to all recorded data points.
	metricsCapacity               int                 // maximum observed number of metrics per resource.
	resourceCapacity              int                 // maximum observed number of resource attributes.
	metricsBuffer                 pmetric.Metrics     // accumulates metrics data before emitting.
	buildInfo                     component.BuildInfo // contains version information
	resourceAttributesSettings    ResourceAttributesSettings
	metricSystemCPULoadAverage15m metricSystemCPULoadAverage15m
	metricSystemCPULoadAverage1m  metricSystemCPULoadAverage1m
	metricSystemCPULoadAverage5m  metricSystemCPULoadAverage5m
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics:            DefaultMetricsSettings(),
		ResourceAttributes: DefaultResourceAttributesSettings(),
	}
}

func (mbc MetricsBuilderConfig) WithMetrics(ms MetricsSettings) MetricsBuilderConfig {
	mbc.Metrics = ms
	return mbc
}

func (mbc MetricsBuilderConfig) WithResourceAttributes(ras ResourceAttributesSettings) MetricsBuilderConfig {
	mbc.ResourceAttributes = ras
	return mbc
}

func NewMetricsBuilder(mbc MetricsBuilderConfig, settings receiver.CreateSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime:                     pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                 pmetric.NewMetrics(),
		buildInfo:                     settings.BuildInfo,
		resourceAttributesSettings:    mbc.ResourceAttributes,
		metricSystemCPULoadAverage15m: newMetricSystemCPULoadAverage15m(mbc.Metrics.SystemCPULoadAverage15m),
		metricSystemCPULoadAverage1m:  newMetricSystemCPULoadAverage1m(mbc.Metrics.SystemCPULoadAverage1m),
		metricSystemCPULoadAverage5m:  newMetricSystemCPULoadAverage5m(mbc.Metrics.SystemCPULoadAverage5m),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
	if mb.resourceCapacity < rm.Resource().Attributes().Len() {
		mb.resourceCapacity = rm.Resource().Attributes().Len()
	}
}

// ResourceMetricsOption applies changes to provided resource metrics.
type ResourceMetricsOption func(ResourceAttributesSettings, pmetric.ResourceMetrics)

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return func(ras ResourceAttributesSettings, rm pmetric.ResourceMetrics) {
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
	}
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceMetricsOption arguments.
func (mb *MetricsBuilder) EmitForResource(rmo ...ResourceMetricsOption) {
	rm := pmetric.NewResourceMetrics()
	rm.SetSchemaUrl(conventions.SchemaURL)
	rm.Resource().Attributes().EnsureCapacity(mb.resourceCapacity)
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/hostmetricsreceiver/load")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricSystemCPULoadAverage15m.emit(ils.Metrics())
	mb.metricSystemCPULoadAverage1m.emit(ils.Metrics())
	mb.metricSystemCPULoadAverage5m.emit(ils.Metrics())

	for _, op := range rmo {
		op(mb.resourceAttributesSettings, rm)
	}
	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user settings, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(rmo ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(rmo...)
	metrics := mb.metricsBuffer
	mb.metricsBuffer = pmetric.NewMetrics()
	return metrics
}

// RecordSystemCPULoadAverage15mDataPoint adds a data point to system.cpu.load_average.15m metric.
func (mb *MetricsBuilder) RecordSystemCPULoadAverage15mDataPoint(ts pcommon.Timestamp, val float64) {
	mb.metricSystemCPULoadAverage15m.recordDataPoint(mb.startTime, ts, val)
}

// RecordSystemCPULoadAverage1mDataPoint adds a data point to system.cpu.load_average.1m metric.
func (mb *MetricsBuilder) RecordSystemCPULoadAverage1mDataPoint(ts pcommon.Timestamp, val float64) {
	mb.metricSystemCPULoadAverage1m.recordDataPoint(mb.startTime, ts, val)
}

// RecordSystemCPULoadAverage5mDataPoint adds a data point to system.cpu.load_average.5m metric.
func (mb *MetricsBuilder) RecordSystemCPULoadAverage5mDataPoint(ts pcommon.Timestamp, val float64) {
	mb.metricSystemCPULoadAverage5m.recordDataPoint(mb.startTime, ts, val)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
