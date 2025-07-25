// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package googlecloudpubsubreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudpubsubreceiver"

import (
	"context"
	"strings"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudpubsubreceiver/internal/metadata"
)

const (
	reportTransport      = "pubsub"
	reportFormatProtobuf = "protobuf"
)

func NewFactory() receiver.Factory {
	f := &pubsubReceiverFactory{
		receivers: make(map[*Config]*pubsubReceiver),
	}
	return receiver.NewFactory(
		metadata.Type,
		f.CreateDefaultConfig,
		receiver.WithTraces(f.CreateTraces, metadata.TracesStability),
		receiver.WithMetrics(f.CreateMetrics, metadata.MetricsStability),
		receiver.WithLogs(f.CreateLogs, metadata.LogsStability),
	)
}

type pubsubReceiverFactory struct {
	receivers map[*Config]*pubsubReceiver
}

func (*pubsubReceiverFactory) CreateDefaultConfig() component.Config {
	return &Config{}
}

func (factory *pubsubReceiverFactory) ensureReceiver(settings receiver.Settings, config component.Config) (*pubsubReceiver, error) {
	receiver := factory.receivers[config.(*Config)]
	if receiver != nil {
		return receiver, nil
	}
	rconfig := config.(*Config)
	obsrecv, err := receiverhelper.NewObsReport(receiverhelper.ObsReportSettings{
		ReceiverID:             settings.ID,
		Transport:              reportTransport,
		ReceiverCreateSettings: settings,
	})
	if err != nil {
		return nil, err
	}
	receiver = &pubsubReceiver{
		settings:  settings,
		obsrecv:   obsrecv,
		userAgent: strings.ReplaceAll(rconfig.UserAgent, "{{version}}", settings.BuildInfo.Version),
		config:    rconfig,
	}
	factory.receivers[config.(*Config)] = receiver
	return receiver, nil
}

func (factory *pubsubReceiverFactory) CreateTraces(
	_ context.Context,
	params receiver.Settings,
	cfg component.Config,
	consumer consumer.Traces,
) (receiver.Traces, error) {
	err := cfg.(*Config).validate()
	if err != nil {
		return nil, err
	}
	receiver, err := factory.ensureReceiver(params, cfg)
	if err != nil {
		return nil, err
	}
	receiver.tracesConsumer = consumer
	return receiver, nil
}

func (factory *pubsubReceiverFactory) CreateMetrics(
	_ context.Context,
	params receiver.Settings,
	cfg component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	err := cfg.(*Config).validate()
	if err != nil {
		return nil, err
	}
	receiver, err := factory.ensureReceiver(params, cfg)
	if err != nil {
		return nil, err
	}
	receiver.metricsConsumer = consumer
	return receiver, nil
}

func (factory *pubsubReceiverFactory) CreateLogs(
	_ context.Context,
	params receiver.Settings,
	cfg component.Config,
	consumer consumer.Logs,
) (receiver.Logs, error) {
	err := cfg.(*Config).validate()
	if err != nil {
		return nil, err
	}
	receiver, err := factory.ensureReceiver(params, cfg)
	if err != nil {
		return nil, err
	}
	receiver.logsConsumer = consumer
	return receiver, nil
}
