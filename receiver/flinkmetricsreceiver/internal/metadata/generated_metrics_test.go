// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap/confmaptest"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

type testMetricsSet int

const (
	testMetricsSetDefault testMetricsSet = iota
	testMetricsSetAll
	testMetricsSetNo
)

func TestMetricsBuilder(t *testing.T) {
	tests := []struct {
		name       string
		metricsSet testMetricsSet
	}{
		{
			name:       "default",
			metricsSet: testMetricsSetDefault,
		},
		{
			name:       "all_metrics",
			metricsSet: testMetricsSetAll,
		},
		{
			name:       "no_metrics",
			metricsSet: testMetricsSetNo,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			start := pcommon.Timestamp(1_000_000_000)
			ts := pcommon.Timestamp(1_000_001_000)
			observedZapCore, observedLogs := observer.New(zap.WarnLevel)
			settings := receivertest.NewNopCreateSettings()
			settings.Logger = zap.New(observedZapCore)
			mb := NewMetricsBuilder(loadConfig(t, test.name), settings, WithStartTime(start))

			expectedWarnings := 0
			assert.Equal(t, expectedWarnings, observedLogs.Len())

			defaultMetricsCount := 0
			allMetricsCount := 0

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordFlinkJobCheckpointCountDataPoint(ts, "1", AttributeCheckpoint(1))

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordFlinkJobCheckpointInProgressDataPoint(ts, "1")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordFlinkJobLastCheckpointSizeDataPoint(ts, "1")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordFlinkJobLastCheckpointTimeDataPoint(ts, "1")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordFlinkJobRestartCountDataPoint(ts, "1")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordFlinkJvmClassLoaderClassesLoadedDataPoint(ts, "1")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordFlinkJvmCPULoadDataPoint(ts, "1")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordFlinkJvmCPUTimeDataPoint(ts, "1")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordFlinkJvmGcCollectionsCountDataPoint(ts, "1", AttributeGarbageCollectorName(1))

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordFlinkJvmGcCollectionsTimeDataPoint(ts, "1", AttributeGarbageCollectorName(1))

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordFlinkJvmMemoryDirectTotalCapacityDataPoint(ts, "1")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordFlinkJvmMemoryDirectUsedDataPoint(ts, "1")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordFlinkJvmMemoryHeapCommittedDataPoint(ts, "1")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordFlinkJvmMemoryHeapMaxDataPoint(ts, "1")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordFlinkJvmMemoryHeapUsedDataPoint(ts, "1")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordFlinkJvmMemoryMappedTotalCapacityDataPoint(ts, "1")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordFlinkJvmMemoryMappedUsedDataPoint(ts, "1")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordFlinkJvmMemoryMetaspaceCommittedDataPoint(ts, "1")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordFlinkJvmMemoryMetaspaceMaxDataPoint(ts, "1")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordFlinkJvmMemoryMetaspaceUsedDataPoint(ts, "1")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordFlinkJvmMemoryNonheapCommittedDataPoint(ts, "1")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordFlinkJvmMemoryNonheapMaxDataPoint(ts, "1")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordFlinkJvmMemoryNonheapUsedDataPoint(ts, "1")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordFlinkJvmThreadsCountDataPoint(ts, "1")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordFlinkMemoryManagedTotalDataPoint(ts, "1")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordFlinkMemoryManagedUsedDataPoint(ts, "1")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordFlinkOperatorRecordCountDataPoint(ts, "1", "attr-val", AttributeRecord(1))

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordFlinkOperatorWatermarkOutputDataPoint(ts, "1", "attr-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordFlinkTaskRecordCountDataPoint(ts, "1", AttributeRecord(1))

			metrics := mb.Emit(WithFlinkJobName("attr-val"), WithFlinkResourceTypeJobmanager, WithFlinkSubtaskIndex("attr-val"), WithFlinkTaskName("attr-val"), WithFlinkTaskmanagerID("attr-val"), WithHostName("attr-val"))

			if test.metricsSet == testMetricsSetNo {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			attrCount := 0
			enabledAttrCount := 0
			attrVal, ok := rm.Resource().Attributes().Get("flink.job.name")
			attrCount++
			assert.Equal(t, mb.resourceAttributesSettings.FlinkJobName.Enabled, ok)
			if mb.resourceAttributesSettings.FlinkJobName.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "attr-val", attrVal.Str())
			}
			attrVal, ok = rm.Resource().Attributes().Get("flink.resource.type")
			attrCount++
			assert.Equal(t, mb.resourceAttributesSettings.FlinkResourceType.Enabled, ok)
			if mb.resourceAttributesSettings.FlinkResourceType.Enabled {
				enabledAttrCount++
				assert.Equal(t, "jobmanager", attrVal.Str())
			}
			attrVal, ok = rm.Resource().Attributes().Get("flink.subtask.index")
			attrCount++
			assert.Equal(t, mb.resourceAttributesSettings.FlinkSubtaskIndex.Enabled, ok)
			if mb.resourceAttributesSettings.FlinkSubtaskIndex.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "attr-val", attrVal.Str())
			}
			attrVal, ok = rm.Resource().Attributes().Get("flink.task.name")
			attrCount++
			assert.Equal(t, mb.resourceAttributesSettings.FlinkTaskName.Enabled, ok)
			if mb.resourceAttributesSettings.FlinkTaskName.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "attr-val", attrVal.Str())
			}
			attrVal, ok = rm.Resource().Attributes().Get("flink.taskmanager.id")
			attrCount++
			assert.Equal(t, mb.resourceAttributesSettings.FlinkTaskmanagerID.Enabled, ok)
			if mb.resourceAttributesSettings.FlinkTaskmanagerID.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "attr-val", attrVal.Str())
			}
			attrVal, ok = rm.Resource().Attributes().Get("host.name")
			attrCount++
			assert.Equal(t, mb.resourceAttributesSettings.HostName.Enabled, ok)
			if mb.resourceAttributesSettings.HostName.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "attr-val", attrVal.Str())
			}
			assert.Equal(t, enabledAttrCount, rm.Resource().Attributes().Len())
			assert.Equal(t, attrCount, 6)

			assert.Equal(t, 1, rm.ScopeMetrics().Len())
			ms := rm.ScopeMetrics().At(0).Metrics()
			if test.metricsSet == testMetricsSetDefault {
				assert.Equal(t, defaultMetricsCount, ms.Len())
			}
			if test.metricsSet == testMetricsSetAll {
				assert.Equal(t, allMetricsCount, ms.Len())
			}
			validatedMetrics := make(map[string]bool)
			for i := 0; i < ms.Len(); i++ {
				switch ms.At(i).Name() {
				case "flink.job.checkpoint.count":
					assert.False(t, validatedMetrics["flink.job.checkpoint.count"], "Found a duplicate in the metrics slice: flink.job.checkpoint.count")
					validatedMetrics["flink.job.checkpoint.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of checkpoints completed or failed.", ms.At(i).Description())
					assert.Equal(t, "{checkpoints}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("checkpoint")
					assert.True(t, ok)
					assert.Equal(t, "completed", attrVal.Str())
				case "flink.job.checkpoint.in_progress":
					assert.False(t, validatedMetrics["flink.job.checkpoint.in_progress"], "Found a duplicate in the metrics slice: flink.job.checkpoint.in_progress")
					validatedMetrics["flink.job.checkpoint.in_progress"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of checkpoints in progress.", ms.At(i).Description())
					assert.Equal(t, "{checkpoints}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "flink.job.last_checkpoint.size":
					assert.False(t, validatedMetrics["flink.job.last_checkpoint.size"], "Found a duplicate in the metrics slice: flink.job.last_checkpoint.size")
					validatedMetrics["flink.job.last_checkpoint.size"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The total size of the last checkpoint.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "flink.job.last_checkpoint.time":
					assert.False(t, validatedMetrics["flink.job.last_checkpoint.time"], "Found a duplicate in the metrics slice: flink.job.last_checkpoint.time")
					validatedMetrics["flink.job.last_checkpoint.time"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The end to end duration of the last checkpoint.", ms.At(i).Description())
					assert.Equal(t, "ms", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "flink.job.restart.count":
					assert.False(t, validatedMetrics["flink.job.restart.count"], "Found a duplicate in the metrics slice: flink.job.restart.count")
					validatedMetrics["flink.job.restart.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The total number of restarts since this job was submitted, including full restarts and fine-grained restarts.", ms.At(i).Description())
					assert.Equal(t, "{restarts}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "flink.jvm.class_loader.classes_loaded":
					assert.False(t, validatedMetrics["flink.jvm.class_loader.classes_loaded"], "Found a duplicate in the metrics slice: flink.jvm.class_loader.classes_loaded")
					validatedMetrics["flink.jvm.class_loader.classes_loaded"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The total number of classes loaded since the start of the JVM.", ms.At(i).Description())
					assert.Equal(t, "{classes}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "flink.jvm.cpu.load":
					assert.False(t, validatedMetrics["flink.jvm.cpu.load"], "Found a duplicate in the metrics slice: flink.jvm.cpu.load")
					validatedMetrics["flink.jvm.cpu.load"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The CPU usage of the JVM for a jobmanager or taskmanager.", ms.At(i).Description())
					assert.Equal(t, "%", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
				case "flink.jvm.cpu.time":
					assert.False(t, validatedMetrics["flink.jvm.cpu.time"], "Found a duplicate in the metrics slice: flink.jvm.cpu.time")
					validatedMetrics["flink.jvm.cpu.time"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The CPU time used by the JVM for a jobmanager or taskmanager.", ms.At(i).Description())
					assert.Equal(t, "ns", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "flink.jvm.gc.collections.count":
					assert.False(t, validatedMetrics["flink.jvm.gc.collections.count"], "Found a duplicate in the metrics slice: flink.jvm.gc.collections.count")
					validatedMetrics["flink.jvm.gc.collections.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The total number of collections that have occurred.", ms.At(i).Description())
					assert.Equal(t, "{collections}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("name")
					assert.True(t, ok)
					assert.Equal(t, "PS_MarkSweep", attrVal.Str())
				case "flink.jvm.gc.collections.time":
					assert.False(t, validatedMetrics["flink.jvm.gc.collections.time"], "Found a duplicate in the metrics slice: flink.jvm.gc.collections.time")
					validatedMetrics["flink.jvm.gc.collections.time"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The total time spent performing garbage collection.", ms.At(i).Description())
					assert.Equal(t, "ms", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("name")
					assert.True(t, ok)
					assert.Equal(t, "PS_MarkSweep", attrVal.Str())
				case "flink.jvm.memory.direct.total_capacity":
					assert.False(t, validatedMetrics["flink.jvm.memory.direct.total_capacity"], "Found a duplicate in the metrics slice: flink.jvm.memory.direct.total_capacity")
					validatedMetrics["flink.jvm.memory.direct.total_capacity"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The total capacity of all buffers in the direct buffer pool.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "flink.jvm.memory.direct.used":
					assert.False(t, validatedMetrics["flink.jvm.memory.direct.used"], "Found a duplicate in the metrics slice: flink.jvm.memory.direct.used")
					validatedMetrics["flink.jvm.memory.direct.used"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The amount of memory used by the JVM for the direct buffer pool.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "flink.jvm.memory.heap.committed":
					assert.False(t, validatedMetrics["flink.jvm.memory.heap.committed"], "Found a duplicate in the metrics slice: flink.jvm.memory.heap.committed")
					validatedMetrics["flink.jvm.memory.heap.committed"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The amount of heap memory guaranteed to be available to the JVM.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "flink.jvm.memory.heap.max":
					assert.False(t, validatedMetrics["flink.jvm.memory.heap.max"], "Found a duplicate in the metrics slice: flink.jvm.memory.heap.max")
					validatedMetrics["flink.jvm.memory.heap.max"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The maximum amount of heap memory that can be used for memory management.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "flink.jvm.memory.heap.used":
					assert.False(t, validatedMetrics["flink.jvm.memory.heap.used"], "Found a duplicate in the metrics slice: flink.jvm.memory.heap.used")
					validatedMetrics["flink.jvm.memory.heap.used"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The amount of heap memory currently used.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "flink.jvm.memory.mapped.total_capacity":
					assert.False(t, validatedMetrics["flink.jvm.memory.mapped.total_capacity"], "Found a duplicate in the metrics slice: flink.jvm.memory.mapped.total_capacity")
					validatedMetrics["flink.jvm.memory.mapped.total_capacity"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of buffers in the mapped buffer pool.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "flink.jvm.memory.mapped.used":
					assert.False(t, validatedMetrics["flink.jvm.memory.mapped.used"], "Found a duplicate in the metrics slice: flink.jvm.memory.mapped.used")
					validatedMetrics["flink.jvm.memory.mapped.used"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The amount of memory used by the JVM for the mapped buffer pool.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "flink.jvm.memory.metaspace.committed":
					assert.False(t, validatedMetrics["flink.jvm.memory.metaspace.committed"], "Found a duplicate in the metrics slice: flink.jvm.memory.metaspace.committed")
					validatedMetrics["flink.jvm.memory.metaspace.committed"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The amount of memory guaranteed to be available to the JVM in the Metaspace memory pool.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "flink.jvm.memory.metaspace.max":
					assert.False(t, validatedMetrics["flink.jvm.memory.metaspace.max"], "Found a duplicate in the metrics slice: flink.jvm.memory.metaspace.max")
					validatedMetrics["flink.jvm.memory.metaspace.max"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The maximum amount of memory that can be used in the Metaspace memory pool.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "flink.jvm.memory.metaspace.used":
					assert.False(t, validatedMetrics["flink.jvm.memory.metaspace.used"], "Found a duplicate in the metrics slice: flink.jvm.memory.metaspace.used")
					validatedMetrics["flink.jvm.memory.metaspace.used"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The amount of memory currently used in the Metaspace memory pool.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "flink.jvm.memory.nonheap.committed":
					assert.False(t, validatedMetrics["flink.jvm.memory.nonheap.committed"], "Found a duplicate in the metrics slice: flink.jvm.memory.nonheap.committed")
					validatedMetrics["flink.jvm.memory.nonheap.committed"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The amount of non-heap memory guaranteed to be available to the JVM.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "flink.jvm.memory.nonheap.max":
					assert.False(t, validatedMetrics["flink.jvm.memory.nonheap.max"], "Found a duplicate in the metrics slice: flink.jvm.memory.nonheap.max")
					validatedMetrics["flink.jvm.memory.nonheap.max"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The maximum amount of non-heap memory that can be used for memory management.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "flink.jvm.memory.nonheap.used":
					assert.False(t, validatedMetrics["flink.jvm.memory.nonheap.used"], "Found a duplicate in the metrics slice: flink.jvm.memory.nonheap.used")
					validatedMetrics["flink.jvm.memory.nonheap.used"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The amount of non-heap memory currently used.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "flink.jvm.threads.count":
					assert.False(t, validatedMetrics["flink.jvm.threads.count"], "Found a duplicate in the metrics slice: flink.jvm.threads.count")
					validatedMetrics["flink.jvm.threads.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The total number of live threads.", ms.At(i).Description())
					assert.Equal(t, "{threads}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "flink.memory.managed.total":
					assert.False(t, validatedMetrics["flink.memory.managed.total"], "Found a duplicate in the metrics slice: flink.memory.managed.total")
					validatedMetrics["flink.memory.managed.total"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The total amount of managed memory.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "flink.memory.managed.used":
					assert.False(t, validatedMetrics["flink.memory.managed.used"], "Found a duplicate in the metrics slice: flink.memory.managed.used")
					validatedMetrics["flink.memory.managed.used"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The amount of managed memory currently used.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "flink.operator.record.count":
					assert.False(t, validatedMetrics["flink.operator.record.count"], "Found a duplicate in the metrics slice: flink.operator.record.count")
					validatedMetrics["flink.operator.record.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of records an operator has.", ms.At(i).Description())
					assert.Equal(t, "{records}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("name")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("record")
					assert.True(t, ok)
					assert.Equal(t, "in", attrVal.Str())
				case "flink.operator.watermark.output":
					assert.False(t, validatedMetrics["flink.operator.watermark.output"], "Found a duplicate in the metrics slice: flink.operator.watermark.output")
					validatedMetrics["flink.operator.watermark.output"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The last watermark this operator has emitted.", ms.At(i).Description())
					assert.Equal(t, "ms", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("name")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
				case "flink.task.record.count":
					assert.False(t, validatedMetrics["flink.task.record.count"], "Found a duplicate in the metrics slice: flink.task.record.count")
					validatedMetrics["flink.task.record.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of records a task has.", ms.At(i).Description())
					assert.Equal(t, "{records}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("record")
					assert.True(t, ok)
					assert.Equal(t, "in", attrVal.Str())
				}
			}
		})
	}
}

func loadConfig(t *testing.T, name string) MetricsBuilderConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	cfg := DefaultMetricsBuilderConfig()
	require.NoError(t, component.UnmarshalConfig(sub, &cfg))
	return cfg
}
