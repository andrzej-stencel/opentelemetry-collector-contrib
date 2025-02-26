// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package googlecloudmonitoringreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudmonitoringreceiver"

import (
	"errors"
	"fmt"
	"time"

	"go.opentelemetry.io/collector/scraper/scraperhelper"
)

const (
	defaultCollectionInterval = 300 * time.Second // Default value for collection interval
	minCollectionInterval     = 60 * time.Second  // Minimum value for collection interval
	defaultFetchDelay         = 60 * time.Second  // Default value for fetch delay
)

type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`

	ProjectID   string         `mapstructure:"project_id"`
	MetricsList []MetricConfig `mapstructure:"metrics_list"`
}

type MetricConfig struct {
	MetricName string `mapstructure:"metric_name"`
	// Filter for listing metric descriptors. Only support `project` and `metric.type` as filter objects.
	// See https://cloud.google.com/monitoring/api/v3/filters#metric-descriptor-filter for more details.
	MetricDescriptorFilter string `mapstructure:"metric_descriptor_filter"`
}

func (config *Config) Validate() error {
	if config.CollectionInterval < minCollectionInterval {
		return fmt.Errorf("\"collection_interval\" must be not lower than the collection interval: %v, current value is %v", minCollectionInterval, config.CollectionInterval)
	}

	if len(config.MetricsList) == 0 {
		return errors.New("missing required field \"metrics_list\" or its value is empty")
	}

	for _, metric := range config.MetricsList {
		if err := metric.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (metric MetricConfig) Validate() error {
	if metric.MetricName != "" && metric.MetricDescriptorFilter != "" {
		return errors.New("fields \"metric_name\" and \"metric_descriptor_filter\" cannot both have value")
	}

	if metric.MetricName == "" && metric.MetricDescriptorFilter == "" {
		return errors.New("fields \"metric_name\" and \"metric_descriptor_filter\" cannot both be empty")
	}

	return nil
}
