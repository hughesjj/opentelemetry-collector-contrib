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

type testConfigCollection int

const (
	testSetDefault testConfigCollection = iota
	testSetAll
	testSetNone
)

func TestMetricsBuilder(t *testing.T) {
	tests := []struct {
		name      string
		configSet testConfigCollection
	}{
		{
			name:      "default",
			configSet: testSetDefault,
		},
		{
			name:      "all_set",
			configSet: testSetAll,
		},
		{
			name:      "none_set",
			configSet: testSetNone,
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
			mb.RecordPostgresqlBackendsDataPoint(ts, 1, "attr-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlBgwriterBuffersAllocatedDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlBgwriterBuffersWritesDataPoint(ts, 1, AttributeBgBufferSource(1))

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlBgwriterCheckpointCountDataPoint(ts, 1, AttributeBgCheckpointType(1))

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlBgwriterDurationDataPoint(ts, 1, AttributeBgDurationType(1))

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlBgwriterMaxwrittenDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlBlocksReadDataPoint(ts, 1, "attr-val", "attr-val", AttributeSource(1))

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlCommitsDataPoint(ts, 1, "attr-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlConnectionMaxDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlDatabaseCountDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlDbSizeDataPoint(ts, 1, "attr-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlIndexScansDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlIndexSizeDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlOperationsDataPoint(ts, 1, "attr-val", "attr-val", AttributeOperation(1))

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlReplicationDataDelayDataPoint(ts, 1, "attr-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlRollbacksDataPoint(ts, 1, "attr-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlRowsDataPoint(ts, 1, "attr-val", "attr-val", AttributeState(1))

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlTableCountDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlTableSizeDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlTableVacuumCountDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlWalAgeDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordPostgresqlWalLagDataPoint(ts, 1, AttributeWalOperationLag(1), "attr-val")

			metrics := mb.Emit(WithPostgresqlDatabaseName("attr-val"), WithPostgresqlIndexName("attr-val"), WithPostgresqlTableName("attr-val"))

			if test.configSet == testSetNone {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			attrCount := 0
			enabledAttrCount := 0
			attrVal, ok := rm.Resource().Attributes().Get("postgresql.database.name")
			attrCount++
			assert.Equal(t, mb.resourceAttributesSettings.PostgresqlDatabaseName.Enabled, ok)
			if mb.resourceAttributesSettings.PostgresqlDatabaseName.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "attr-val", attrVal.Str())
			}
			attrVal, ok = rm.Resource().Attributes().Get("postgresql.index.name")
			attrCount++
			assert.Equal(t, mb.resourceAttributesSettings.PostgresqlIndexName.Enabled, ok)
			if mb.resourceAttributesSettings.PostgresqlIndexName.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "attr-val", attrVal.Str())
			}
			attrVal, ok = rm.Resource().Attributes().Get("postgresql.table.name")
			attrCount++
			assert.Equal(t, mb.resourceAttributesSettings.PostgresqlTableName.Enabled, ok)
			if mb.resourceAttributesSettings.PostgresqlTableName.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "attr-val", attrVal.Str())
			}
			assert.Equal(t, enabledAttrCount, rm.Resource().Attributes().Len())
			assert.Equal(t, attrCount, 3)

			assert.Equal(t, 1, rm.ScopeMetrics().Len())
			ms := rm.ScopeMetrics().At(0).Metrics()
			if test.configSet == testSetDefault {
				assert.Equal(t, defaultMetricsCount, ms.Len())
			}
			if test.configSet == testSetNone {
				assert.Equal(t, allMetricsCount, ms.Len())
			}
			validatedMetrics := make(map[string]bool)
			for i := 0; i < ms.Len(); i++ {
				switch ms.At(i).Name() {
				case "postgresql.backends":
					assert.False(t, validatedMetrics["postgresql.backends"], "Found a duplicate in the metrics slice: postgresql.backends")
					validatedMetrics["postgresql.backends"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of backends.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("database")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
				case "postgresql.bgwriter.buffers.allocated":
					assert.False(t, validatedMetrics["postgresql.bgwriter.buffers.allocated"], "Found a duplicate in the metrics slice: postgresql.bgwriter.buffers.allocated")
					validatedMetrics["postgresql.bgwriter.buffers.allocated"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of buffers allocated.", ms.At(i).Description())
					assert.Equal(t, "{buffers}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "postgresql.bgwriter.buffers.writes":
					assert.False(t, validatedMetrics["postgresql.bgwriter.buffers.writes"], "Found a duplicate in the metrics slice: postgresql.bgwriter.buffers.writes")
					validatedMetrics["postgresql.bgwriter.buffers.writes"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of buffers written.", ms.At(i).Description())
					assert.Equal(t, "{buffers}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("source")
					assert.True(t, ok)
					assert.Equal(t, "backend", attrVal.Str())
				case "postgresql.bgwriter.checkpoint.count":
					assert.False(t, validatedMetrics["postgresql.bgwriter.checkpoint.count"], "Found a duplicate in the metrics slice: postgresql.bgwriter.checkpoint.count")
					validatedMetrics["postgresql.bgwriter.checkpoint.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of checkpoints performed.", ms.At(i).Description())
					assert.Equal(t, "{checkpoints}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("type")
					assert.True(t, ok)
					assert.Equal(t, "requested", attrVal.Str())
				case "postgresql.bgwriter.duration":
					assert.False(t, validatedMetrics["postgresql.bgwriter.duration"], "Found a duplicate in the metrics slice: postgresql.bgwriter.duration")
					validatedMetrics["postgresql.bgwriter.duration"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Total time spent writing and syncing files to disk by checkpoints.", ms.At(i).Description())
					assert.Equal(t, "ms", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("type")
					assert.True(t, ok)
					assert.Equal(t, "sync", attrVal.Str())
				case "postgresql.bgwriter.maxwritten":
					assert.False(t, validatedMetrics["postgresql.bgwriter.maxwritten"], "Found a duplicate in the metrics slice: postgresql.bgwriter.maxwritten")
					validatedMetrics["postgresql.bgwriter.maxwritten"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of times the background writer stopped a cleaning scan because it had written too many buffers.", ms.At(i).Description())
					assert.Equal(t, "", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "postgresql.blocks_read":
					assert.False(t, validatedMetrics["postgresql.blocks_read"], "Found a duplicate in the metrics slice: postgresql.blocks_read")
					validatedMetrics["postgresql.blocks_read"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of blocks read.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("database")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("table")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("source")
					assert.True(t, ok)
					assert.Equal(t, "heap_read", attrVal.Str())
				case "postgresql.commits":
					assert.False(t, validatedMetrics["postgresql.commits"], "Found a duplicate in the metrics slice: postgresql.commits")
					validatedMetrics["postgresql.commits"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of commits.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("database")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
				case "postgresql.connection.max":
					assert.False(t, validatedMetrics["postgresql.connection.max"], "Found a duplicate in the metrics slice: postgresql.connection.max")
					validatedMetrics["postgresql.connection.max"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Configured maximum number of client connections allowed", ms.At(i).Description())
					assert.Equal(t, "{connections}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "postgresql.database.count":
					assert.False(t, validatedMetrics["postgresql.database.count"], "Found a duplicate in the metrics slice: postgresql.database.count")
					validatedMetrics["postgresql.database.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of user databases.", ms.At(i).Description())
					assert.Equal(t, "{databases}", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "postgresql.db_size":
					assert.False(t, validatedMetrics["postgresql.db_size"], "Found a duplicate in the metrics slice: postgresql.db_size")
					validatedMetrics["postgresql.db_size"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The database disk usage.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("database")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
				case "postgresql.index.scans":
					assert.False(t, validatedMetrics["postgresql.index.scans"], "Found a duplicate in the metrics slice: postgresql.index.scans")
					validatedMetrics["postgresql.index.scans"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of index scans on a table.", ms.At(i).Description())
					assert.Equal(t, "{scans}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "postgresql.index.size":
					assert.False(t, validatedMetrics["postgresql.index.size"], "Found a duplicate in the metrics slice: postgresql.index.size")
					validatedMetrics["postgresql.index.size"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The size of the index on disk.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "postgresql.operations":
					assert.False(t, validatedMetrics["postgresql.operations"], "Found a duplicate in the metrics slice: postgresql.operations")
					validatedMetrics["postgresql.operations"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of db row operations.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("database")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("table")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("operation")
					assert.True(t, ok)
					assert.Equal(t, "ins", attrVal.Str())
				case "postgresql.replication.data_delay":
					assert.False(t, validatedMetrics["postgresql.replication.data_delay"], "Found a duplicate in the metrics slice: postgresql.replication.data_delay")
					validatedMetrics["postgresql.replication.data_delay"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The amount of data delayed in replication.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("replication_client")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
				case "postgresql.rollbacks":
					assert.False(t, validatedMetrics["postgresql.rollbacks"], "Found a duplicate in the metrics slice: postgresql.rollbacks")
					validatedMetrics["postgresql.rollbacks"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of rollbacks.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("database")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
				case "postgresql.rows":
					assert.False(t, validatedMetrics["postgresql.rows"], "Found a duplicate in the metrics slice: postgresql.rows")
					validatedMetrics["postgresql.rows"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The number of rows in the database.", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("database")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("table")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("state")
					assert.True(t, ok)
					assert.Equal(t, "dead", attrVal.Str())
				case "postgresql.table.count":
					assert.False(t, validatedMetrics["postgresql.table.count"], "Found a duplicate in the metrics slice: postgresql.table.count")
					validatedMetrics["postgresql.table.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of user tables in a database.", ms.At(i).Description())
					assert.Equal(t, "", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "postgresql.table.size":
					assert.False(t, validatedMetrics["postgresql.table.size"], "Found a duplicate in the metrics slice: postgresql.table.size")
					validatedMetrics["postgresql.table.size"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Disk space used by a table.", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "postgresql.table.vacuum.count":
					assert.False(t, validatedMetrics["postgresql.table.vacuum.count"], "Found a duplicate in the metrics slice: postgresql.table.vacuum.count")
					validatedMetrics["postgresql.table.vacuum.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of times a table has manually been vacuumed.", ms.At(i).Description())
					assert.Equal(t, "{vacuums}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "postgresql.wal.age":
					assert.False(t, validatedMetrics["postgresql.wal.age"], "Found a duplicate in the metrics slice: postgresql.wal.age")
					validatedMetrics["postgresql.wal.age"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Age of the oldest WAL file.", ms.At(i).Description())
					assert.Equal(t, "s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "postgresql.wal.lag":
					assert.False(t, validatedMetrics["postgresql.wal.lag"], "Found a duplicate in the metrics slice: postgresql.wal.lag")
					validatedMetrics["postgresql.wal.lag"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Time between flushing recent WAL locally and receiving notification that the standby server has completed an operation with it.", ms.At(i).Description())
					assert.Equal(t, "s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("operation")
					assert.True(t, ok)
					assert.Equal(t, "flush", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("replication_client")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
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
