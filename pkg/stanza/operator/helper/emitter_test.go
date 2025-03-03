// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package helper

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
)

func TestLogEmitter(t *testing.T) {
	rwMtx := &sync.RWMutex{}
	var receivedEntries []*entry.Entry
	emitter := NewLogEmitter(
		componenttest.NewNopTelemetrySettings(),
		func(_ context.Context, entries []*entry.Entry) {
			rwMtx.Lock()
			defer rwMtx.Unlock()
			receivedEntries = entries
		},
	)

	require.NoError(t, emitter.Start(nil))

	defer func() {
		require.NoError(t, emitter.Stop())
	}()

	in := entry.New()

	assert.NoError(t, emitter.Process(context.Background(), in))

	require.Eventually(t, func() bool {
		rwMtx.RLock()
		defer rwMtx.RUnlock()
		return receivedEntries != nil
	}, time.Second, 10*time.Millisecond)
	require.Equal(t, in, receivedEntries[0])
}

func TestLogEmitterEmits(t *testing.T) {
	const (
		count   = 100
		timeout = time.Second
	)
	rwMtx := &sync.RWMutex{}
	var receivedEntries []*entry.Entry
	emitter := NewLogEmitter(
		componenttest.NewNopTelemetrySettings(),
		func(_ context.Context, entries []*entry.Entry) {
			rwMtx.Lock()
			defer rwMtx.Unlock()
			receivedEntries = append(receivedEntries, entries...)
		},
	)

	require.NoError(t, emitter.Start(nil))
	defer func() {
		require.NoError(t, emitter.Stop())
	}()

	entries := complexEntries(count)

	ctx := context.Background()
	for _, e := range entries {
		assert.NoError(t, emitter.Process(ctx, e))
	}

	require.Eventually(t, func() bool {
		rwMtx.RLock()
		defer rwMtx.RUnlock()
		return receivedEntries != nil
	}, timeout, 10*time.Millisecond)
	require.Len(t, receivedEntries, count)
}

func TestLogEmitterEmitsBatched(t *testing.T) {
	const (
		count   = 234
		timeout = time.Second
	)
	rwMtx := &sync.RWMutex{}
	var receivedEntries []*entry.Entry
	emitter := NewLogEmitter(
		componenttest.NewNopTelemetrySettings(),
		func(_ context.Context, entries []*entry.Entry) {
			rwMtx.Lock()
			defer rwMtx.Unlock()
			receivedEntries = append(receivedEntries, entries...)
		},
	)

	require.NoError(t, emitter.Start(nil))
	defer func() {
		require.NoError(t, emitter.Stop())
	}()

	entries := complexEntries(count)

	ctx := context.Background()
	assert.NoError(t, emitter.ProcessBatch(ctx, entries))

	require.Eventually(t, func() bool {
		rwMtx.RLock()
		defer rwMtx.RUnlock()
		return receivedEntries != nil
	}, timeout, 10*time.Millisecond)
	require.Len(t, receivedEntries, count)
}

func complexEntries(count int) []*entry.Entry {
	return complexEntriesForNDifferentHosts(count, 1)
}

func complexEntriesForNDifferentHosts(count int, n int) []*entry.Entry {
	ret := make([]*entry.Entry, count)
	for i := 0; i < count; i++ {
		e := entry.New()
		e.Severity = entry.Error
		e.Resource = map[string]any{
			"host":   fmt.Sprintf("host-%d", i%n),
			"bool":   true,
			"int":    123,
			"double": 12.34,
			"string": "hello",
			"object": map[string]any{
				"bool":   true,
				"int":    123,
				"double": 12.34,
				"string": "hello",
			},
		}
		e.Body = map[string]any{
			"bool":   true,
			"int":    123,
			"double": 12.34,
			"string": "hello",
			"bytes":  []byte("asdf"),
			"object": map[string]any{
				"bool":   true,
				"int":    123,
				"double": 12.34,
				"string": "hello",
				"bytes":  []byte("asdf"),
				"object": map[string]any{
					"bool":   true,
					"int":    123,
					"double": 12.34,
					"string": "hello",
					"bytes":  []byte("asdf"),
				},
			},
		}
		ret[i] = e
	}
	return ret
}
