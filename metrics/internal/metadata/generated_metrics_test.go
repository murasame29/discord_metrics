// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
			mb := NewMetricsBuilder(loadMetricsBuilderConfig(t, test.name), settings, WithStartTime(start))

			expectedWarnings := 0

			assert.Equal(t, expectedWarnings, observedLogs.Len())

			defaultMetricsCount := 0
			allMetricsCount := 0

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordDiscordMetricsJoinCountDataPoint(ts, 1, "discord-metrics.environment-val", "discord-metrics.actor.id-val", "discord-metrics.actor.name-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordDiscordMetricsLeaveCountDataPoint(ts, 1, "discord-metrics.environment-val", "discord-metrics.actor.id-val", "discord-metrics.actor.name-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordDiscordMetricsMessagesCountDataPoint(ts, 1, "discord-metrics.environment-val", "discord-metrics.category.id-val", "discord-metrics.category.name-val", "discord-metrics.channel.id-val", "discord-metrics.channel.name-val", "discord-metrics.actor.id-val", "discord-metrics.actor.name-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordDiscordMetricsVcEventCountDataPoint(ts, 1, "discord-metrics.environment-val", "discord-metrics.category.id-val", "discord-metrics.category.name-val", "discord-metrics.channel.id-val", "discord-metrics.channel.name-val", "discord-metrics.actor.id-val", "discord-metrics.actor.name-val", "discord-metrics.vc.event-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordDiscordMetricsVcActiveMembersDataPoint(ts, 1, "discord-metrics.environment-val", "discord-metrics.category.id-val", "discord-metrics.category.name-val", "discord-metrics.channel.id-val", "discord-metrics.channel.name-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordDiscordMetricsVcDurationDataPoint(ts, 1, "discord-metrics.environment-val", "discord-metrics.category.id-val", "discord-metrics.category.name-val", "discord-metrics.channel.id-val", "discord-metrics.channel.name-val", "discord-metrics.actor.id-val", "discord-metrics.actor.name-val")

			res := pcommon.NewResource()
			metrics := mb.Emit(WithResource(res))

			if test.configSet == testSetNone {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			assert.Equal(t, res, rm.Resource())
			assert.Equal(t, 1, rm.ScopeMetrics().Len())
			ms := rm.ScopeMetrics().At(0).Metrics()
			if test.configSet == testSetDefault {
				assert.Equal(t, defaultMetricsCount, ms.Len())
			}
			if test.configSet == testSetAll {
				assert.Equal(t, allMetricsCount, ms.Len())
			}
			validatedMetrics := make(map[string]bool)
			for i := 0; i < ms.Len(); i++ {
				switch ms.At(i).Name() {
				case "discord-metrics.join.count":
					assert.False(t, validatedMetrics["discord-metrics.join.count"], "Found a duplicate in the metrics slice: discord-metrics.join.count")
					validatedMetrics["discord-metrics.join.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "サーバに入ったユーザの数", ms.At(i).Description())
					assert.Equal(t, "{join}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("discord-metrics.environment")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.environment-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("discord-metrics.actor.id")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.actor.id-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("discord-metrics.actor.name")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.actor.name-val", attrVal.Str())
				case "discord-metrics.leave.count":
					assert.False(t, validatedMetrics["discord-metrics.leave.count"], "Found a duplicate in the metrics slice: discord-metrics.leave.count")
					validatedMetrics["discord-metrics.leave.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "サーバから出たユーザの数", ms.At(i).Description())
					assert.Equal(t, "{leave}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("discord-metrics.environment")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.environment-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("discord-metrics.actor.id")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.actor.id-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("discord-metrics.actor.name")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.actor.name-val", attrVal.Str())
				case "discord-metrics.messages.count":
					assert.False(t, validatedMetrics["discord-metrics.messages.count"], "Found a duplicate in the metrics slice: discord-metrics.messages.count")
					validatedMetrics["discord-metrics.messages.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "サーバ内で入力されたメッセージの総数", ms.At(i).Description())
					assert.Equal(t, "{messages}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("discord-metrics.environment")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.environment-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("discord-metrics.category.id")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.category.id-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("discord-metrics.category.name")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.category.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("discord-metrics.channel.id")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.channel.id-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("discord-metrics.channel.name")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.channel.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("discord-metrics.actor.id")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.actor.id-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("discord-metrics.actor.name")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.actor.name-val", attrVal.Str())
				case "discord-metrics.vc-event.count":
					assert.False(t, validatedMetrics["discord-metrics.vc-event.count"], "Found a duplicate in the metrics slice: discord-metrics.vc-event.count")
					validatedMetrics["discord-metrics.vc-event.count"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "vcで行われたイベント", ms.At(i).Description())
					assert.Equal(t, "{vc-event}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("discord-metrics.environment")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.environment-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("discord-metrics.category.id")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.category.id-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("discord-metrics.category.name")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.category.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("discord-metrics.channel.id")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.channel.id-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("discord-metrics.channel.name")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.channel.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("discord-metrics.actor.id")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.actor.id-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("discord-metrics.actor.name")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.actor.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("discord-metrics.vc.event")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.vc.event-val", attrVal.Str())
				case "discord_metrics.vc_active_members":
					assert.False(t, validatedMetrics["discord_metrics.vc_active_members"], "Found a duplicate in the metrics slice: discord_metrics.vc_active_members")
					validatedMetrics["discord_metrics.vc_active_members"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, " VCのアクティブメンバー数", ms.At(i).Description())
					assert.Equal(t, "{members}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("discord-metrics.environment")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.environment-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("discord-metrics.category.id")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.category.id-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("discord-metrics.category.name")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.category.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("discord-metrics.channel.id")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.channel.id-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("discord-metrics.channel.name")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.channel.name-val", attrVal.Str())
				case "discord_metrics.vc_duration":
					assert.False(t, validatedMetrics["discord_metrics.vc_duration"], "Found a duplicate in the metrics slice: discord_metrics.vc_duration")
					validatedMetrics["discord_metrics.vc_duration"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "ユーザーごとのVC滞在時間", ms.At(i).Description())
					assert.Equal(t, "{seconds}", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("discord-metrics.environment")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.environment-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("discord-metrics.category.id")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.category.id-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("discord-metrics.category.name")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.category.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("discord-metrics.channel.id")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.channel.id-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("discord-metrics.channel.name")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.channel.name-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("discord-metrics.actor.id")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.actor.id-val", attrVal.Str())
					attrVal, ok = dp.Attributes().Get("discord-metrics.actor.name")
					assert.True(t, ok)
					assert.EqualValues(t, "discord-metrics.actor.name-val", attrVal.Str())
				}
			}
		})
	}
}
