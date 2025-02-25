// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package helper // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"

import (
	"context"

	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
)

// LogEmitter is a stanza operator that emits log entries to the consumer callback function `consumerFunc`
type LogEmitter struct {
	OutputOperator
	consumerFunc func(context.Context, []*entry.Entry)
}

// NewLogEmitter creates a new receiver output
func NewLogEmitter(set component.TelemetrySettings, consumerFunc func(context.Context, []*entry.Entry)) *LogEmitter {
	op, _ := NewOutputConfig("log_emitter", "log_emitter").Build(set)
	e := &LogEmitter{
		OutputOperator: op,
		consumerFunc:   consumerFunc,
	}
	return e
}

// Start starts the goroutine(s) required for this operator
func (e *LogEmitter) Start(_ operator.Persister) error {
	return nil
}

// Stop will close the log channel and stop running goroutines
func (e *LogEmitter) Stop() error {
	return nil
}

// ProcessBatch emits the entries to the consumerFunc
func (e *LogEmitter) ProcessBatch(ctx context.Context, entries []*entry.Entry) error {
	e.consumerFunc(ctx, entries)
	return nil
}

// Process will emit an entry to the consumerFunc
func (e *LogEmitter) Process(ctx context.Context, ent *entry.Entry) error {
	e.consumerFunc(ctx, []*entry.Entry{ent})
	return nil
}
