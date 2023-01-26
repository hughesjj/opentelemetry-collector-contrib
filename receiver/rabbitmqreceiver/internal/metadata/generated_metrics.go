// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
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

// MetricsSettings provides settings for rabbitmqreceiver metrics.
type MetricsSettings struct {
	RabbitmqConsumerCount       MetricSettings `mapstructure:"rabbitmq.consumer.count"`
	RabbitmqMessageAcknowledged MetricSettings `mapstructure:"rabbitmq.message.acknowledged"`
	RabbitmqMessageCurrent      MetricSettings `mapstructure:"rabbitmq.message.current"`
	RabbitmqMessageDelivered    MetricSettings `mapstructure:"rabbitmq.message.delivered"`
	RabbitmqMessageDropped      MetricSettings `mapstructure:"rabbitmq.message.dropped"`
	RabbitmqMessagePublished    MetricSettings `mapstructure:"rabbitmq.message.published"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		RabbitmqConsumerCount: MetricSettings{
			Enabled: true,
		},
		RabbitmqMessageAcknowledged: MetricSettings{
			Enabled: true,
		},
		RabbitmqMessageCurrent: MetricSettings{
			Enabled: true,
		},
		RabbitmqMessageDelivered: MetricSettings{
			Enabled: true,
		},
		RabbitmqMessageDropped: MetricSettings{
			Enabled: true,
		},
		RabbitmqMessagePublished: MetricSettings{
			Enabled: true,
		},
	}
}

// ResourceAttributeSettings provides common settings for a particular metric.
type ResourceAttributeSettings struct {
	Enabled bool `mapstructure:"enabled"`

	enabledProvidedByUser bool
}

func (ras *ResourceAttributeSettings) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ras, confmap.WithErrorUnused())
	if err != nil {
		return err
	}
	ras.enabledProvidedByUser = parser.IsSet("enabled")
	return nil
}

// ResourceAttributesSettings provides settings for rabbitmqreceiver metrics.
type ResourceAttributesSettings struct {
	RabbitmqNodeName  ResourceAttributeSettings `mapstructure:"rabbitmq.node.name"`
	RabbitmqQueueName ResourceAttributeSettings `mapstructure:"rabbitmq.queue.name"`
	RabbitmqVhostName ResourceAttributeSettings `mapstructure:"rabbitmq.vhost.name"`
}

func DefaultResourceAttributesSettings() ResourceAttributesSettings {
	return ResourceAttributesSettings{
		RabbitmqNodeName: ResourceAttributeSettings{
			Enabled: true,
		},
		RabbitmqQueueName: ResourceAttributeSettings{
			Enabled: true,
		},
		RabbitmqVhostName: ResourceAttributeSettings{
			Enabled: true,
		},
	}
}

// AttributeMessageState specifies the a value message.state attribute.
type AttributeMessageState int

const (
	_ AttributeMessageState = iota
	AttributeMessageStateReady
	AttributeMessageStateUnacknowledged
)

// String returns the string representation of the AttributeMessageState.
func (av AttributeMessageState) String() string {
	switch av {
	case AttributeMessageStateReady:
		return "ready"
	case AttributeMessageStateUnacknowledged:
		return "unacknowledged"
	}
	return ""
}

// MapAttributeMessageState is a helper map of string to AttributeMessageState attribute value.
var MapAttributeMessageState = map[string]AttributeMessageState{
	"ready":          AttributeMessageStateReady,
	"unacknowledged": AttributeMessageStateUnacknowledged,
}

type metricRabbitmqConsumerCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills rabbitmq.consumer.count metric with initial data.
func (m *metricRabbitmqConsumerCount) init() {
	m.data.SetName("rabbitmq.consumer.count")
	m.data.SetDescription("The number of consumers currently reading from the queue.")
	m.data.SetUnit("{consumers}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricRabbitmqConsumerCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricRabbitmqConsumerCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricRabbitmqConsumerCount) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricRabbitmqConsumerCount(settings MetricSettings) metricRabbitmqConsumerCount {
	m := metricRabbitmqConsumerCount{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricRabbitmqMessageAcknowledged struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills rabbitmq.message.acknowledged metric with initial data.
func (m *metricRabbitmqMessageAcknowledged) init() {
	m.data.SetName("rabbitmq.message.acknowledged")
	m.data.SetDescription("The number of messages acknowledged by consumers.")
	m.data.SetUnit("{messages}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricRabbitmqMessageAcknowledged) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricRabbitmqMessageAcknowledged) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricRabbitmqMessageAcknowledged) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricRabbitmqMessageAcknowledged(settings MetricSettings) metricRabbitmqMessageAcknowledged {
	m := metricRabbitmqMessageAcknowledged{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricRabbitmqMessageCurrent struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills rabbitmq.message.current metric with initial data.
func (m *metricRabbitmqMessageCurrent) init() {
	m.data.SetName("rabbitmq.message.current")
	m.data.SetDescription("The total number of messages currently in the queue.")
	m.data.SetUnit("{messages}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricRabbitmqMessageCurrent) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, messageStateAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("state", messageStateAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricRabbitmqMessageCurrent) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricRabbitmqMessageCurrent) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricRabbitmqMessageCurrent(settings MetricSettings) metricRabbitmqMessageCurrent {
	m := metricRabbitmqMessageCurrent{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricRabbitmqMessageDelivered struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills rabbitmq.message.delivered metric with initial data.
func (m *metricRabbitmqMessageDelivered) init() {
	m.data.SetName("rabbitmq.message.delivered")
	m.data.SetDescription("The number of messages delivered to consumers.")
	m.data.SetUnit("{messages}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricRabbitmqMessageDelivered) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricRabbitmqMessageDelivered) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricRabbitmqMessageDelivered) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricRabbitmqMessageDelivered(settings MetricSettings) metricRabbitmqMessageDelivered {
	m := metricRabbitmqMessageDelivered{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricRabbitmqMessageDropped struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills rabbitmq.message.dropped metric with initial data.
func (m *metricRabbitmqMessageDropped) init() {
	m.data.SetName("rabbitmq.message.dropped")
	m.data.SetDescription("The number of messages dropped as unroutable.")
	m.data.SetUnit("{messages}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricRabbitmqMessageDropped) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricRabbitmqMessageDropped) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricRabbitmqMessageDropped) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricRabbitmqMessageDropped(settings MetricSettings) metricRabbitmqMessageDropped {
	m := metricRabbitmqMessageDropped{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricRabbitmqMessagePublished struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills rabbitmq.message.published metric with initial data.
func (m *metricRabbitmqMessagePublished) init() {
	m.data.SetName("rabbitmq.message.published")
	m.data.SetDescription("The number of messages published to a queue.")
	m.data.SetUnit("{messages}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricRabbitmqMessagePublished) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricRabbitmqMessagePublished) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricRabbitmqMessagePublished) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricRabbitmqMessagePublished(settings MetricSettings) metricRabbitmqMessagePublished {
	m := metricRabbitmqMessagePublished{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilderConfig is a structural subset of an otherwise 1-1 copy of metadata.yaml
type MetricsBuilderConfig struct {
	Metrics            MetricsSettings            `mapstructure:",squash"`
	ResourceAttributes ResourceAttributesSettings `mapstructure:",squash"`
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                         pcommon.Timestamp   // start time that will be applied to all recorded data points.
	metricsCapacity                   int                 // maximum observed number of metrics per resource.
	resourceCapacity                  int                 // maximum observed number of resource attributes.
	metricsBuffer                     pmetric.Metrics     // accumulates metrics data before emitting.
	buildInfo                         component.BuildInfo // contains version information
	resourceAttributesSettings        ResourceAttributesSettings
	metricRabbitmqConsumerCount       metricRabbitmqConsumerCount
	metricRabbitmqMessageAcknowledged metricRabbitmqMessageAcknowledged
	metricRabbitmqMessageCurrent      metricRabbitmqMessageCurrent
	metricRabbitmqMessageDelivered    metricRabbitmqMessageDelivered
	metricRabbitmqMessageDropped      metricRabbitmqMessageDropped
	metricRabbitmqMessagePublished    metricRabbitmqMessagePublished
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
		startTime:                         pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                     pmetric.NewMetrics(),
		buildInfo:                         settings.BuildInfo,
		resourceAttributesSettings:        mbc.ResourceAttributes,
		metricRabbitmqConsumerCount:       newMetricRabbitmqConsumerCount(mbc.Metrics.RabbitmqConsumerCount),
		metricRabbitmqMessageAcknowledged: newMetricRabbitmqMessageAcknowledged(mbc.Metrics.RabbitmqMessageAcknowledged),
		metricRabbitmqMessageCurrent:      newMetricRabbitmqMessageCurrent(mbc.Metrics.RabbitmqMessageCurrent),
		metricRabbitmqMessageDelivered:    newMetricRabbitmqMessageDelivered(mbc.Metrics.RabbitmqMessageDelivered),
		metricRabbitmqMessageDropped:      newMetricRabbitmqMessageDropped(mbc.Metrics.RabbitmqMessageDropped),
		metricRabbitmqMessagePublished:    newMetricRabbitmqMessagePublished(mbc.Metrics.RabbitmqMessagePublished),
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

// WithRabbitmqNodeName sets provided value as "rabbitmq.node.name" attribute for current resource.
func WithRabbitmqNodeName(val string) ResourceMetricsOption {
	return func(ras ResourceAttributesSettings, rm pmetric.ResourceMetrics) {
		if ras.RabbitmqNodeName.Enabled {
			rm.Resource().Attributes().PutStr("rabbitmq.node.name", val)
		}
	}
}

// WithRabbitmqQueueName sets provided value as "rabbitmq.queue.name" attribute for current resource.
func WithRabbitmqQueueName(val string) ResourceMetricsOption {
	return func(ras ResourceAttributesSettings, rm pmetric.ResourceMetrics) {
		if ras.RabbitmqQueueName.Enabled {
			rm.Resource().Attributes().PutStr("rabbitmq.queue.name", val)
		}
	}
}

// WithRabbitmqVhostName sets provided value as "rabbitmq.vhost.name" attribute for current resource.
func WithRabbitmqVhostName(val string) ResourceMetricsOption {
	return func(ras ResourceAttributesSettings, rm pmetric.ResourceMetrics) {
		if ras.RabbitmqVhostName.Enabled {
			rm.Resource().Attributes().PutStr("rabbitmq.vhost.name", val)
		}
	}
}

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
	rm.Resource().Attributes().EnsureCapacity(mb.resourceCapacity)
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/rabbitmqreceiver")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricRabbitmqConsumerCount.emit(ils.Metrics())
	mb.metricRabbitmqMessageAcknowledged.emit(ils.Metrics())
	mb.metricRabbitmqMessageCurrent.emit(ils.Metrics())
	mb.metricRabbitmqMessageDelivered.emit(ils.Metrics())
	mb.metricRabbitmqMessageDropped.emit(ils.Metrics())
	mb.metricRabbitmqMessagePublished.emit(ils.Metrics())

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
	metrics := pmetric.NewMetrics()
	mb.metricsBuffer.MoveTo(metrics)
	return metrics
}

// RecordRabbitmqConsumerCountDataPoint adds a data point to rabbitmq.consumer.count metric.
func (mb *MetricsBuilder) RecordRabbitmqConsumerCountDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricRabbitmqConsumerCount.recordDataPoint(mb.startTime, ts, val)
}

// RecordRabbitmqMessageAcknowledgedDataPoint adds a data point to rabbitmq.message.acknowledged metric.
func (mb *MetricsBuilder) RecordRabbitmqMessageAcknowledgedDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricRabbitmqMessageAcknowledged.recordDataPoint(mb.startTime, ts, val)
}

// RecordRabbitmqMessageCurrentDataPoint adds a data point to rabbitmq.message.current metric.
func (mb *MetricsBuilder) RecordRabbitmqMessageCurrentDataPoint(ts pcommon.Timestamp, val int64, messageStateAttributeValue AttributeMessageState) {
	mb.metricRabbitmqMessageCurrent.recordDataPoint(mb.startTime, ts, val, messageStateAttributeValue.String())
}

// RecordRabbitmqMessageDeliveredDataPoint adds a data point to rabbitmq.message.delivered metric.
func (mb *MetricsBuilder) RecordRabbitmqMessageDeliveredDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricRabbitmqMessageDelivered.recordDataPoint(mb.startTime, ts, val)
}

// RecordRabbitmqMessageDroppedDataPoint adds a data point to rabbitmq.message.dropped metric.
func (mb *MetricsBuilder) RecordRabbitmqMessageDroppedDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricRabbitmqMessageDropped.recordDataPoint(mb.startTime, ts, val)
}

// RecordRabbitmqMessagePublishedDataPoint adds a data point to rabbitmq.message.published metric.
func (mb *MetricsBuilder) RecordRabbitmqMessagePublishedDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricRabbitmqMessagePublished.recordDataPoint(mb.startTime, ts, val)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
