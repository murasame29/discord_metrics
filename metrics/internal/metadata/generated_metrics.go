// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
)

type metricDiscordMetricsJoinCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills discord_metrics.join.count metric with initial data.
func (m *metricDiscordMetricsJoinCount) init() {
	m.data.SetName("discord_metrics.join.count")
	m.data.SetDescription("サーバに入ったユーザの数")
	m.data.SetUnit("{join}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricDiscordMetricsJoinCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, discordMetricsGuildIDAttributeValue string, discordMetricsEnvironmentAttributeValue string, discordMetricsActorIDAttributeValue string, discordMetricsActorNameAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("discord_metrics.guild.id", discordMetricsGuildIDAttributeValue)
	dp.Attributes().PutStr("discord_metrics.environment", discordMetricsEnvironmentAttributeValue)
	dp.Attributes().PutStr("discord_metrics.actor.id", discordMetricsActorIDAttributeValue)
	dp.Attributes().PutStr("discord_metrics.actor.name", discordMetricsActorNameAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricDiscordMetricsJoinCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricDiscordMetricsJoinCount) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricDiscordMetricsJoinCount(cfg MetricConfig) metricDiscordMetricsJoinCount {
	m := metricDiscordMetricsJoinCount{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricDiscordMetricsLeaveCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills discord_metrics.leave.count metric with initial data.
func (m *metricDiscordMetricsLeaveCount) init() {
	m.data.SetName("discord_metrics.leave.count")
	m.data.SetDescription("サーバから出たユーザの数")
	m.data.SetUnit("{leave}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricDiscordMetricsLeaveCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, discordMetricsGuildIDAttributeValue string, discordMetricsEnvironmentAttributeValue string, discordMetricsActorIDAttributeValue string, discordMetricsActorNameAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("discord_metrics.guild.id", discordMetricsGuildIDAttributeValue)
	dp.Attributes().PutStr("discord_metrics.environment", discordMetricsEnvironmentAttributeValue)
	dp.Attributes().PutStr("discord_metrics.actor.id", discordMetricsActorIDAttributeValue)
	dp.Attributes().PutStr("discord_metrics.actor.name", discordMetricsActorNameAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricDiscordMetricsLeaveCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricDiscordMetricsLeaveCount) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricDiscordMetricsLeaveCount(cfg MetricConfig) metricDiscordMetricsLeaveCount {
	m := metricDiscordMetricsLeaveCount{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricDiscordMetricsMessagesCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills discord_metrics.messages.count metric with initial data.
func (m *metricDiscordMetricsMessagesCount) init() {
	m.data.SetName("discord_metrics.messages.count")
	m.data.SetDescription("サーバ内で入力されたメッセージの総数")
	m.data.SetUnit("{messages}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricDiscordMetricsMessagesCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, discordMetricsGuildIDAttributeValue string, discordMetricsEnvironmentAttributeValue string, discordMetricsCategoryIDAttributeValue string, discordMetricsCategoryNameAttributeValue string, discordMetricsChannelIDAttributeValue string, discordMetricsChannelNameAttributeValue string, discordMetricsActorIDAttributeValue string, discordMetricsActorNameAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("discord_metrics.guild.id", discordMetricsGuildIDAttributeValue)
	dp.Attributes().PutStr("discord_metrics.environment", discordMetricsEnvironmentAttributeValue)
	dp.Attributes().PutStr("discord_metrics.category.id", discordMetricsCategoryIDAttributeValue)
	dp.Attributes().PutStr("discord_metrics.category.name", discordMetricsCategoryNameAttributeValue)
	dp.Attributes().PutStr("discord_metrics.channel.id", discordMetricsChannelIDAttributeValue)
	dp.Attributes().PutStr("discord_metrics.channel.name", discordMetricsChannelNameAttributeValue)
	dp.Attributes().PutStr("discord_metrics.actor.id", discordMetricsActorIDAttributeValue)
	dp.Attributes().PutStr("discord_metrics.actor.name", discordMetricsActorNameAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricDiscordMetricsMessagesCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricDiscordMetricsMessagesCount) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricDiscordMetricsMessagesCount(cfg MetricConfig) metricDiscordMetricsMessagesCount {
	m := metricDiscordMetricsMessagesCount{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricDiscordMetricsVcEventCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills discord_metrics.vc-event.count metric with initial data.
func (m *metricDiscordMetricsVcEventCount) init() {
	m.data.SetName("discord_metrics.vc-event.count")
	m.data.SetDescription("vcで行われたイベント")
	m.data.SetUnit("{vc-event}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricDiscordMetricsVcEventCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, discordMetricsGuildIDAttributeValue string, discordMetricsEnvironmentAttributeValue string, discordMetricsCategoryIDAttributeValue string, discordMetricsCategoryNameAttributeValue string, discordMetricsChannelIDAttributeValue string, discordMetricsChannelNameAttributeValue string, discordMetricsActorIDAttributeValue string, discordMetricsActorNameAttributeValue string, discordMetricsVcEventAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("discord_metrics.guild.id", discordMetricsGuildIDAttributeValue)
	dp.Attributes().PutStr("discord_metrics.environment", discordMetricsEnvironmentAttributeValue)
	dp.Attributes().PutStr("discord_metrics.category.id", discordMetricsCategoryIDAttributeValue)
	dp.Attributes().PutStr("discord_metrics.category.name", discordMetricsCategoryNameAttributeValue)
	dp.Attributes().PutStr("discord_metrics.channel.id", discordMetricsChannelIDAttributeValue)
	dp.Attributes().PutStr("discord_metrics.channel.name", discordMetricsChannelNameAttributeValue)
	dp.Attributes().PutStr("discord_metrics.actor.id", discordMetricsActorIDAttributeValue)
	dp.Attributes().PutStr("discord_metrics.actor.name", discordMetricsActorNameAttributeValue)
	dp.Attributes().PutStr("discord_metrics.vc.event", discordMetricsVcEventAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricDiscordMetricsVcEventCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricDiscordMetricsVcEventCount) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricDiscordMetricsVcEventCount(cfg MetricConfig) metricDiscordMetricsVcEventCount {
	m := metricDiscordMetricsVcEventCount{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricDiscordMetricsVcActiveMembers struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills discord_metrics.vc_active_members metric with initial data.
func (m *metricDiscordMetricsVcActiveMembers) init() {
	m.data.SetName("discord_metrics.vc_active_members")
	m.data.SetDescription(" VCのアクティブメンバー数")
	m.data.SetUnit("{members}")
	m.data.SetEmptyGauge()
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricDiscordMetricsVcActiveMembers) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, discordMetricsGuildIDAttributeValue string, discordMetricsEnvironmentAttributeValue string, discordMetricsCategoryIDAttributeValue string, discordMetricsCategoryNameAttributeValue string, discordMetricsChannelIDAttributeValue string, discordMetricsChannelNameAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("discord_metrics.guild.id", discordMetricsGuildIDAttributeValue)
	dp.Attributes().PutStr("discord_metrics.environment", discordMetricsEnvironmentAttributeValue)
	dp.Attributes().PutStr("discord_metrics.category.id", discordMetricsCategoryIDAttributeValue)
	dp.Attributes().PutStr("discord_metrics.category.name", discordMetricsCategoryNameAttributeValue)
	dp.Attributes().PutStr("discord_metrics.channel.id", discordMetricsChannelIDAttributeValue)
	dp.Attributes().PutStr("discord_metrics.channel.name", discordMetricsChannelNameAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricDiscordMetricsVcActiveMembers) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricDiscordMetricsVcActiveMembers) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricDiscordMetricsVcActiveMembers(cfg MetricConfig) metricDiscordMetricsVcActiveMembers {
	m := metricDiscordMetricsVcActiveMembers{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricDiscordMetricsVcDuration struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills discord_metrics.vc_duration metric with initial data.
func (m *metricDiscordMetricsVcDuration) init() {
	m.data.SetName("discord_metrics.vc_duration")
	m.data.SetDescription("ユーザーごとのVC滞在時間")
	m.data.SetUnit("{seconds}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricDiscordMetricsVcDuration) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, discordMetricsGuildIDAttributeValue string, discordMetricsEnvironmentAttributeValue string, discordMetricsCategoryIDAttributeValue string, discordMetricsCategoryNameAttributeValue string, discordMetricsChannelIDAttributeValue string, discordMetricsChannelNameAttributeValue string, discordMetricsActorIDAttributeValue string, discordMetricsActorNameAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("discord_metrics.guild.id", discordMetricsGuildIDAttributeValue)
	dp.Attributes().PutStr("discord_metrics.environment", discordMetricsEnvironmentAttributeValue)
	dp.Attributes().PutStr("discord_metrics.category.id", discordMetricsCategoryIDAttributeValue)
	dp.Attributes().PutStr("discord_metrics.category.name", discordMetricsCategoryNameAttributeValue)
	dp.Attributes().PutStr("discord_metrics.channel.id", discordMetricsChannelIDAttributeValue)
	dp.Attributes().PutStr("discord_metrics.channel.name", discordMetricsChannelNameAttributeValue)
	dp.Attributes().PutStr("discord_metrics.actor.id", discordMetricsActorIDAttributeValue)
	dp.Attributes().PutStr("discord_metrics.actor.name", discordMetricsActorNameAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricDiscordMetricsVcDuration) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricDiscordMetricsVcDuration) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricDiscordMetricsVcDuration(cfg MetricConfig) metricDiscordMetricsVcDuration {
	m := metricDiscordMetricsVcDuration{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user config.
type MetricsBuilder struct {
	config                              MetricsBuilderConfig // config of the metrics builder.
	startTime                           pcommon.Timestamp    // start time that will be applied to all recorded data points.
	metricsCapacity                     int                  // maximum observed number of metrics per resource.
	metricsBuffer                       pmetric.Metrics      // accumulates metrics data before emitting.
	buildInfo                           component.BuildInfo  // contains version information.
	metricDiscordMetricsJoinCount       metricDiscordMetricsJoinCount
	metricDiscordMetricsLeaveCount      metricDiscordMetricsLeaveCount
	metricDiscordMetricsMessagesCount   metricDiscordMetricsMessagesCount
	metricDiscordMetricsVcEventCount    metricDiscordMetricsVcEventCount
	metricDiscordMetricsVcActiveMembers metricDiscordMetricsVcActiveMembers
	metricDiscordMetricsVcDuration      metricDiscordMetricsVcDuration
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(mbc MetricsBuilderConfig, settings receiver.CreateSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		config:                              mbc,
		startTime:                           pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                       pmetric.NewMetrics(),
		buildInfo:                           settings.BuildInfo,
		metricDiscordMetricsJoinCount:       newMetricDiscordMetricsJoinCount(mbc.Metrics.DiscordMetricsJoinCount),
		metricDiscordMetricsLeaveCount:      newMetricDiscordMetricsLeaveCount(mbc.Metrics.DiscordMetricsLeaveCount),
		metricDiscordMetricsMessagesCount:   newMetricDiscordMetricsMessagesCount(mbc.Metrics.DiscordMetricsMessagesCount),
		metricDiscordMetricsVcEventCount:    newMetricDiscordMetricsVcEventCount(mbc.Metrics.DiscordMetricsVcEventCount),
		metricDiscordMetricsVcActiveMembers: newMetricDiscordMetricsVcActiveMembers(mbc.Metrics.DiscordMetricsVcActiveMembers),
		metricDiscordMetricsVcDuration:      newMetricDiscordMetricsVcDuration(mbc.Metrics.DiscordMetricsVcDuration),
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
}

// ResourceMetricsOption applies changes to provided resource metrics.
type ResourceMetricsOption func(pmetric.ResourceMetrics)

// WithResource sets the provided resource on the emitted ResourceMetrics.
// It's recommended to use ResourceBuilder to create the resource.
func WithResource(res pcommon.Resource) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		res.CopyTo(rm.Resource())
	}
}

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
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
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricDiscordMetricsJoinCount.emit(ils.Metrics())
	mb.metricDiscordMetricsLeaveCount.emit(ils.Metrics())
	mb.metricDiscordMetricsMessagesCount.emit(ils.Metrics())
	mb.metricDiscordMetricsVcEventCount.emit(ils.Metrics())
	mb.metricDiscordMetricsVcActiveMembers.emit(ils.Metrics())
	mb.metricDiscordMetricsVcDuration.emit(ils.Metrics())

	for _, op := range rmo {
		op(rm)
	}
	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user config, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(rmo ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(rmo...)
	metrics := mb.metricsBuffer
	mb.metricsBuffer = pmetric.NewMetrics()
	return metrics
}

// RecordDiscordMetricsJoinCountDataPoint adds a data point to discord_metrics.join.count metric.
func (mb *MetricsBuilder) RecordDiscordMetricsJoinCountDataPoint(ts pcommon.Timestamp, val int64, discordMetricsGuildIDAttributeValue string, discordMetricsEnvironmentAttributeValue string, discordMetricsActorIDAttributeValue string, discordMetricsActorNameAttributeValue string) {
	mb.metricDiscordMetricsJoinCount.recordDataPoint(mb.startTime, ts, val, discordMetricsGuildIDAttributeValue, discordMetricsEnvironmentAttributeValue, discordMetricsActorIDAttributeValue, discordMetricsActorNameAttributeValue)
}

// RecordDiscordMetricsLeaveCountDataPoint adds a data point to discord_metrics.leave.count metric.
func (mb *MetricsBuilder) RecordDiscordMetricsLeaveCountDataPoint(ts pcommon.Timestamp, val int64, discordMetricsGuildIDAttributeValue string, discordMetricsEnvironmentAttributeValue string, discordMetricsActorIDAttributeValue string, discordMetricsActorNameAttributeValue string) {
	mb.metricDiscordMetricsLeaveCount.recordDataPoint(mb.startTime, ts, val, discordMetricsGuildIDAttributeValue, discordMetricsEnvironmentAttributeValue, discordMetricsActorIDAttributeValue, discordMetricsActorNameAttributeValue)
}

// RecordDiscordMetricsMessagesCountDataPoint adds a data point to discord_metrics.messages.count metric.
func (mb *MetricsBuilder) RecordDiscordMetricsMessagesCountDataPoint(ts pcommon.Timestamp, val int64, discordMetricsGuildIDAttributeValue string, discordMetricsEnvironmentAttributeValue string, discordMetricsCategoryIDAttributeValue string, discordMetricsCategoryNameAttributeValue string, discordMetricsChannelIDAttributeValue string, discordMetricsChannelNameAttributeValue string, discordMetricsActorIDAttributeValue string, discordMetricsActorNameAttributeValue string) {
	mb.metricDiscordMetricsMessagesCount.recordDataPoint(mb.startTime, ts, val, discordMetricsGuildIDAttributeValue, discordMetricsEnvironmentAttributeValue, discordMetricsCategoryIDAttributeValue, discordMetricsCategoryNameAttributeValue, discordMetricsChannelIDAttributeValue, discordMetricsChannelNameAttributeValue, discordMetricsActorIDAttributeValue, discordMetricsActorNameAttributeValue)
}

// RecordDiscordMetricsVcEventCountDataPoint adds a data point to discord_metrics.vc-event.count metric.
func (mb *MetricsBuilder) RecordDiscordMetricsVcEventCountDataPoint(ts pcommon.Timestamp, val int64, discordMetricsGuildIDAttributeValue string, discordMetricsEnvironmentAttributeValue string, discordMetricsCategoryIDAttributeValue string, discordMetricsCategoryNameAttributeValue string, discordMetricsChannelIDAttributeValue string, discordMetricsChannelNameAttributeValue string, discordMetricsActorIDAttributeValue string, discordMetricsActorNameAttributeValue string, discordMetricsVcEventAttributeValue string) {
	mb.metricDiscordMetricsVcEventCount.recordDataPoint(mb.startTime, ts, val, discordMetricsGuildIDAttributeValue, discordMetricsEnvironmentAttributeValue, discordMetricsCategoryIDAttributeValue, discordMetricsCategoryNameAttributeValue, discordMetricsChannelIDAttributeValue, discordMetricsChannelNameAttributeValue, discordMetricsActorIDAttributeValue, discordMetricsActorNameAttributeValue, discordMetricsVcEventAttributeValue)
}

// RecordDiscordMetricsVcActiveMembersDataPoint adds a data point to discord_metrics.vc_active_members metric.
func (mb *MetricsBuilder) RecordDiscordMetricsVcActiveMembersDataPoint(ts pcommon.Timestamp, val int64, discordMetricsGuildIDAttributeValue string, discordMetricsEnvironmentAttributeValue string, discordMetricsCategoryIDAttributeValue string, discordMetricsCategoryNameAttributeValue string, discordMetricsChannelIDAttributeValue string, discordMetricsChannelNameAttributeValue string) {
	mb.metricDiscordMetricsVcActiveMembers.recordDataPoint(mb.startTime, ts, val, discordMetricsGuildIDAttributeValue, discordMetricsEnvironmentAttributeValue, discordMetricsCategoryIDAttributeValue, discordMetricsCategoryNameAttributeValue, discordMetricsChannelIDAttributeValue, discordMetricsChannelNameAttributeValue)
}

// RecordDiscordMetricsVcDurationDataPoint adds a data point to discord_metrics.vc_duration metric.
func (mb *MetricsBuilder) RecordDiscordMetricsVcDurationDataPoint(ts pcommon.Timestamp, val int64, discordMetricsGuildIDAttributeValue string, discordMetricsEnvironmentAttributeValue string, discordMetricsCategoryIDAttributeValue string, discordMetricsCategoryNameAttributeValue string, discordMetricsChannelIDAttributeValue string, discordMetricsChannelNameAttributeValue string, discordMetricsActorIDAttributeValue string, discordMetricsActorNameAttributeValue string) {
	mb.metricDiscordMetricsVcDuration.recordDataPoint(mb.startTime, ts, val, discordMetricsGuildIDAttributeValue, discordMetricsEnvironmentAttributeValue, discordMetricsCategoryIDAttributeValue, discordMetricsCategoryNameAttributeValue, discordMetricsChannelIDAttributeValue, discordMetricsChannelNameAttributeValue, discordMetricsActorIDAttributeValue, discordMetricsActorNameAttributeValue)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
