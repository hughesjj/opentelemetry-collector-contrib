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

// MetricsSettings provides settings for memcachedreceiver metrics.
type MetricsSettings struct {
	MemcachedBytes              MetricSettings `mapstructure:"memcached.bytes"`
	MemcachedCommands           MetricSettings `mapstructure:"memcached.commands"`
	MemcachedConnectionsCurrent MetricSettings `mapstructure:"memcached.connections.current"`
	MemcachedConnectionsTotal   MetricSettings `mapstructure:"memcached.connections.total"`
	MemcachedCPUUsage           MetricSettings `mapstructure:"memcached.cpu.usage"`
	MemcachedCurrentItems       MetricSettings `mapstructure:"memcached.current_items"`
	MemcachedEvictions          MetricSettings `mapstructure:"memcached.evictions"`
	MemcachedNetwork            MetricSettings `mapstructure:"memcached.network"`
	MemcachedOperationHitRatio  MetricSettings `mapstructure:"memcached.operation_hit_ratio"`
	MemcachedOperations         MetricSettings `mapstructure:"memcached.operations"`
	MemcachedThreads            MetricSettings `mapstructure:"memcached.threads"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		MemcachedBytes: MetricSettings{
			Enabled: true,
		},
		MemcachedCommands: MetricSettings{
			Enabled: true,
		},
		MemcachedConnectionsCurrent: MetricSettings{
			Enabled: true,
		},
		MemcachedConnectionsTotal: MetricSettings{
			Enabled: true,
		},
		MemcachedCPUUsage: MetricSettings{
			Enabled: true,
		},
		MemcachedCurrentItems: MetricSettings{
			Enabled: true,
		},
		MemcachedEvictions: MetricSettings{
			Enabled: true,
		},
		MemcachedNetwork: MetricSettings{
			Enabled: true,
		},
		MemcachedOperationHitRatio: MetricSettings{
			Enabled: true,
		},
		MemcachedOperations: MetricSettings{
			Enabled: true,
		},
		MemcachedThreads: MetricSettings{
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

// ResourceAttributesSettings provides settings for memcachedreceiver metrics.
type ResourceAttributesSettings struct {
}

func DefaultResourceAttributesSettings() ResourceAttributesSettings {
	return ResourceAttributesSettings{}
}

// AttributeCommand specifies the a value command attribute.
type AttributeCommand int

const (
	_ AttributeCommand = iota
	AttributeCommandGet
	AttributeCommandSet
	AttributeCommandFlush
	AttributeCommandTouch
)

// String returns the string representation of the AttributeCommand.
func (av AttributeCommand) String() string {
	switch av {
	case AttributeCommandGet:
		return "get"
	case AttributeCommandSet:
		return "set"
	case AttributeCommandFlush:
		return "flush"
	case AttributeCommandTouch:
		return "touch"
	}
	return ""
}

// MapAttributeCommand is a helper map of string to AttributeCommand attribute value.
var MapAttributeCommand = map[string]AttributeCommand{
	"get":   AttributeCommandGet,
	"set":   AttributeCommandSet,
	"flush": AttributeCommandFlush,
	"touch": AttributeCommandTouch,
}

// AttributeDirection specifies the a value direction attribute.
type AttributeDirection int

const (
	_ AttributeDirection = iota
	AttributeDirectionSent
	AttributeDirectionReceived
)

// String returns the string representation of the AttributeDirection.
func (av AttributeDirection) String() string {
	switch av {
	case AttributeDirectionSent:
		return "sent"
	case AttributeDirectionReceived:
		return "received"
	}
	return ""
}

// MapAttributeDirection is a helper map of string to AttributeDirection attribute value.
var MapAttributeDirection = map[string]AttributeDirection{
	"sent":     AttributeDirectionSent,
	"received": AttributeDirectionReceived,
}

// AttributeOperation specifies the a value operation attribute.
type AttributeOperation int

const (
	_ AttributeOperation = iota
	AttributeOperationIncrement
	AttributeOperationDecrement
	AttributeOperationGet
)

// String returns the string representation of the AttributeOperation.
func (av AttributeOperation) String() string {
	switch av {
	case AttributeOperationIncrement:
		return "increment"
	case AttributeOperationDecrement:
		return "decrement"
	case AttributeOperationGet:
		return "get"
	}
	return ""
}

// MapAttributeOperation is a helper map of string to AttributeOperation attribute value.
var MapAttributeOperation = map[string]AttributeOperation{
	"increment": AttributeOperationIncrement,
	"decrement": AttributeOperationDecrement,
	"get":       AttributeOperationGet,
}

// AttributeState specifies the a value state attribute.
type AttributeState int

const (
	_ AttributeState = iota
	AttributeStateSystem
	AttributeStateUser
)

// String returns the string representation of the AttributeState.
func (av AttributeState) String() string {
	switch av {
	case AttributeStateSystem:
		return "system"
	case AttributeStateUser:
		return "user"
	}
	return ""
}

// MapAttributeState is a helper map of string to AttributeState attribute value.
var MapAttributeState = map[string]AttributeState{
	"system": AttributeStateSystem,
	"user":   AttributeStateUser,
}

// AttributeType specifies the a value type attribute.
type AttributeType int

const (
	_ AttributeType = iota
	AttributeTypeHit
	AttributeTypeMiss
)

// String returns the string representation of the AttributeType.
func (av AttributeType) String() string {
	switch av {
	case AttributeTypeHit:
		return "hit"
	case AttributeTypeMiss:
		return "miss"
	}
	return ""
}

// MapAttributeType is a helper map of string to AttributeType attribute value.
var MapAttributeType = map[string]AttributeType{
	"hit":  AttributeTypeHit,
	"miss": AttributeTypeMiss,
}

type metricMemcachedBytes struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills memcached.bytes metric with initial data.
func (m *metricMemcachedBytes) init() {
	m.data.SetName("memcached.bytes")
	m.data.SetDescription("Current number of bytes used by this server to store items.")
	m.data.SetUnit("By")
	m.data.SetEmptyGauge()
}

func (m *metricMemcachedBytes) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMemcachedBytes) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMemcachedBytes) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMemcachedBytes(settings MetricSettings) metricMemcachedBytes {
	m := metricMemcachedBytes{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricMemcachedCommands struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills memcached.commands metric with initial data.
func (m *metricMemcachedCommands) init() {
	m.data.SetName("memcached.commands")
	m.data.SetDescription("Commands executed.")
	m.data.SetUnit("{commands}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricMemcachedCommands) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, commandAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("command", commandAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMemcachedCommands) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMemcachedCommands) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMemcachedCommands(settings MetricSettings) metricMemcachedCommands {
	m := metricMemcachedCommands{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricMemcachedConnectionsCurrent struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills memcached.connections.current metric with initial data.
func (m *metricMemcachedConnectionsCurrent) init() {
	m.data.SetName("memcached.connections.current")
	m.data.SetDescription("The current number of open connections.")
	m.data.SetUnit("{connections}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricMemcachedConnectionsCurrent) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMemcachedConnectionsCurrent) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMemcachedConnectionsCurrent) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMemcachedConnectionsCurrent(settings MetricSettings) metricMemcachedConnectionsCurrent {
	m := metricMemcachedConnectionsCurrent{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricMemcachedConnectionsTotal struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills memcached.connections.total metric with initial data.
func (m *metricMemcachedConnectionsTotal) init() {
	m.data.SetName("memcached.connections.total")
	m.data.SetDescription("Total number of connections opened since the server started running.")
	m.data.SetUnit("{connections}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricMemcachedConnectionsTotal) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMemcachedConnectionsTotal) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMemcachedConnectionsTotal) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMemcachedConnectionsTotal(settings MetricSettings) metricMemcachedConnectionsTotal {
	m := metricMemcachedConnectionsTotal{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricMemcachedCPUUsage struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills memcached.cpu.usage metric with initial data.
func (m *metricMemcachedCPUUsage) init() {
	m.data.SetName("memcached.cpu.usage")
	m.data.SetDescription("Accumulated user and system time.")
	m.data.SetUnit("s")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricMemcachedCPUUsage) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, stateAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
	dp.Attributes().PutStr("state", stateAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMemcachedCPUUsage) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMemcachedCPUUsage) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMemcachedCPUUsage(settings MetricSettings) metricMemcachedCPUUsage {
	m := metricMemcachedCPUUsage{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricMemcachedCurrentItems struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills memcached.current_items metric with initial data.
func (m *metricMemcachedCurrentItems) init() {
	m.data.SetName("memcached.current_items")
	m.data.SetDescription("Number of items currently stored in the cache.")
	m.data.SetUnit("{items}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricMemcachedCurrentItems) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMemcachedCurrentItems) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMemcachedCurrentItems) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMemcachedCurrentItems(settings MetricSettings) metricMemcachedCurrentItems {
	m := metricMemcachedCurrentItems{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricMemcachedEvictions struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills memcached.evictions metric with initial data.
func (m *metricMemcachedEvictions) init() {
	m.data.SetName("memcached.evictions")
	m.data.SetDescription("Cache item evictions.")
	m.data.SetUnit("{evictions}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricMemcachedEvictions) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMemcachedEvictions) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMemcachedEvictions) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMemcachedEvictions(settings MetricSettings) metricMemcachedEvictions {
	m := metricMemcachedEvictions{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricMemcachedNetwork struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills memcached.network metric with initial data.
func (m *metricMemcachedNetwork) init() {
	m.data.SetName("memcached.network")
	m.data.SetDescription("Bytes transferred over the network.")
	m.data.SetUnit("by")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricMemcachedNetwork) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, directionAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("direction", directionAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMemcachedNetwork) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMemcachedNetwork) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMemcachedNetwork(settings MetricSettings) metricMemcachedNetwork {
	m := metricMemcachedNetwork{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricMemcachedOperationHitRatio struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills memcached.operation_hit_ratio metric with initial data.
func (m *metricMemcachedOperationHitRatio) init() {
	m.data.SetName("memcached.operation_hit_ratio")
	m.data.SetDescription("Hit ratio for operations, expressed as a percentage value between 0.0 and 100.0.")
	m.data.SetUnit("%")
	m.data.SetEmptyGauge()
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricMemcachedOperationHitRatio) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, operationAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
	dp.Attributes().PutStr("operation", operationAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMemcachedOperationHitRatio) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMemcachedOperationHitRatio) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMemcachedOperationHitRatio(settings MetricSettings) metricMemcachedOperationHitRatio {
	m := metricMemcachedOperationHitRatio{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricMemcachedOperations struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills memcached.operations metric with initial data.
func (m *metricMemcachedOperations) init() {
	m.data.SetName("memcached.operations")
	m.data.SetDescription("Operation counts.")
	m.data.SetUnit("{operations}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricMemcachedOperations) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, typeAttributeValue string, operationAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("type", typeAttributeValue)
	dp.Attributes().PutStr("operation", operationAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMemcachedOperations) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMemcachedOperations) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMemcachedOperations(settings MetricSettings) metricMemcachedOperations {
	m := metricMemcachedOperations{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricMemcachedThreads struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills memcached.threads metric with initial data.
func (m *metricMemcachedThreads) init() {
	m.data.SetName("memcached.threads")
	m.data.SetDescription("Number of threads used by the memcached instance.")
	m.data.SetUnit("{threads}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricMemcachedThreads) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricMemcachedThreads) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricMemcachedThreads) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricMemcachedThreads(settings MetricSettings) metricMemcachedThreads {
	m := metricMemcachedThreads{settings: settings}
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
	startTime                         pcommon.Timestamp   // start time that will be applied to all recorded data points.
	metricsCapacity                   int                 // maximum observed number of metrics per resource.
	resourceCapacity                  int                 // maximum observed number of resource attributes.
	metricsBuffer                     pmetric.Metrics     // accumulates metrics data before emitting.
	buildInfo                         component.BuildInfo // contains version information
	resourceAttributesSettings        ResourceAttributesSettings
	metricMemcachedBytes              metricMemcachedBytes
	metricMemcachedCommands           metricMemcachedCommands
	metricMemcachedConnectionsCurrent metricMemcachedConnectionsCurrent
	metricMemcachedConnectionsTotal   metricMemcachedConnectionsTotal
	metricMemcachedCPUUsage           metricMemcachedCPUUsage
	metricMemcachedCurrentItems       metricMemcachedCurrentItems
	metricMemcachedEvictions          metricMemcachedEvictions
	metricMemcachedNetwork            metricMemcachedNetwork
	metricMemcachedOperationHitRatio  metricMemcachedOperationHitRatio
	metricMemcachedOperations         metricMemcachedOperations
	metricMemcachedThreads            metricMemcachedThreads
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

func NewMetricsBuilderConfig(ms MetricsSettings, ras ResourceAttributesSettings) MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics:            ms,
		ResourceAttributes: ras,
	}
}

func NewMetricsBuilder(mbc MetricsBuilderConfig, settings receiver.CreateSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime:                         pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                     pmetric.NewMetrics(),
		buildInfo:                         settings.BuildInfo,
		resourceAttributesSettings:        mbc.ResourceAttributes,
		metricMemcachedBytes:              newMetricMemcachedBytes(mbc.Metrics.MemcachedBytes),
		metricMemcachedCommands:           newMetricMemcachedCommands(mbc.Metrics.MemcachedCommands),
		metricMemcachedConnectionsCurrent: newMetricMemcachedConnectionsCurrent(mbc.Metrics.MemcachedConnectionsCurrent),
		metricMemcachedConnectionsTotal:   newMetricMemcachedConnectionsTotal(mbc.Metrics.MemcachedConnectionsTotal),
		metricMemcachedCPUUsage:           newMetricMemcachedCPUUsage(mbc.Metrics.MemcachedCPUUsage),
		metricMemcachedCurrentItems:       newMetricMemcachedCurrentItems(mbc.Metrics.MemcachedCurrentItems),
		metricMemcachedEvictions:          newMetricMemcachedEvictions(mbc.Metrics.MemcachedEvictions),
		metricMemcachedNetwork:            newMetricMemcachedNetwork(mbc.Metrics.MemcachedNetwork),
		metricMemcachedOperationHitRatio:  newMetricMemcachedOperationHitRatio(mbc.Metrics.MemcachedOperationHitRatio),
		metricMemcachedOperations:         newMetricMemcachedOperations(mbc.Metrics.MemcachedOperations),
		metricMemcachedThreads:            newMetricMemcachedThreads(mbc.Metrics.MemcachedThreads),
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
	rm.Resource().Attributes().EnsureCapacity(mb.resourceCapacity)
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/memcachedreceiver")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricMemcachedBytes.emit(ils.Metrics())
	mb.metricMemcachedCommands.emit(ils.Metrics())
	mb.metricMemcachedConnectionsCurrent.emit(ils.Metrics())
	mb.metricMemcachedConnectionsTotal.emit(ils.Metrics())
	mb.metricMemcachedCPUUsage.emit(ils.Metrics())
	mb.metricMemcachedCurrentItems.emit(ils.Metrics())
	mb.metricMemcachedEvictions.emit(ils.Metrics())
	mb.metricMemcachedNetwork.emit(ils.Metrics())
	mb.metricMemcachedOperationHitRatio.emit(ils.Metrics())
	mb.metricMemcachedOperations.emit(ils.Metrics())
	mb.metricMemcachedThreads.emit(ils.Metrics())

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

// RecordMemcachedBytesDataPoint adds a data point to memcached.bytes metric.
func (mb *MetricsBuilder) RecordMemcachedBytesDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricMemcachedBytes.recordDataPoint(mb.startTime, ts, val)
}

// RecordMemcachedCommandsDataPoint adds a data point to memcached.commands metric.
func (mb *MetricsBuilder) RecordMemcachedCommandsDataPoint(ts pcommon.Timestamp, val int64, commandAttributeValue AttributeCommand) {
	mb.metricMemcachedCommands.recordDataPoint(mb.startTime, ts, val, commandAttributeValue.String())
}

// RecordMemcachedConnectionsCurrentDataPoint adds a data point to memcached.connections.current metric.
func (mb *MetricsBuilder) RecordMemcachedConnectionsCurrentDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricMemcachedConnectionsCurrent.recordDataPoint(mb.startTime, ts, val)
}

// RecordMemcachedConnectionsTotalDataPoint adds a data point to memcached.connections.total metric.
func (mb *MetricsBuilder) RecordMemcachedConnectionsTotalDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricMemcachedConnectionsTotal.recordDataPoint(mb.startTime, ts, val)
}

// RecordMemcachedCPUUsageDataPoint adds a data point to memcached.cpu.usage metric.
func (mb *MetricsBuilder) RecordMemcachedCPUUsageDataPoint(ts pcommon.Timestamp, val float64, stateAttributeValue AttributeState) {
	mb.metricMemcachedCPUUsage.recordDataPoint(mb.startTime, ts, val, stateAttributeValue.String())
}

// RecordMemcachedCurrentItemsDataPoint adds a data point to memcached.current_items metric.
func (mb *MetricsBuilder) RecordMemcachedCurrentItemsDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricMemcachedCurrentItems.recordDataPoint(mb.startTime, ts, val)
}

// RecordMemcachedEvictionsDataPoint adds a data point to memcached.evictions metric.
func (mb *MetricsBuilder) RecordMemcachedEvictionsDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricMemcachedEvictions.recordDataPoint(mb.startTime, ts, val)
}

// RecordMemcachedNetworkDataPoint adds a data point to memcached.network metric.
func (mb *MetricsBuilder) RecordMemcachedNetworkDataPoint(ts pcommon.Timestamp, val int64, directionAttributeValue AttributeDirection) {
	mb.metricMemcachedNetwork.recordDataPoint(mb.startTime, ts, val, directionAttributeValue.String())
}

// RecordMemcachedOperationHitRatioDataPoint adds a data point to memcached.operation_hit_ratio metric.
func (mb *MetricsBuilder) RecordMemcachedOperationHitRatioDataPoint(ts pcommon.Timestamp, val float64, operationAttributeValue AttributeOperation) {
	mb.metricMemcachedOperationHitRatio.recordDataPoint(mb.startTime, ts, val, operationAttributeValue.String())
}

// RecordMemcachedOperationsDataPoint adds a data point to memcached.operations metric.
func (mb *MetricsBuilder) RecordMemcachedOperationsDataPoint(ts pcommon.Timestamp, val int64, typeAttributeValue AttributeType, operationAttributeValue AttributeOperation) {
	mb.metricMemcachedOperations.recordDataPoint(mb.startTime, ts, val, typeAttributeValue.String(), operationAttributeValue.String())
}

// RecordMemcachedThreadsDataPoint adds a data point to memcached.threads metric.
func (mb *MetricsBuilder) RecordMemcachedThreadsDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricMemcachedThreads.recordDataPoint(mb.startTime, ts, val)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
