// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottl

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

type mockGetter struct {
	values []string
}

func (s mockGetter) GetStatements() []string {
	return s.values
}

func (s mockGetter) GetConditions() []string {
	return s.values
}

func (s mockGetter) GetValueExpressions() []string {
	return s.values
}

type mockFailingContextInferrer struct {
	err error
}

func (r *mockFailingContextInferrer) infer(_, _, _ []string) (string, error) {
	return "", r.err
}

func (r *mockFailingContextInferrer) inferFromStatements(_ []string) (string, error) {
	return "", r.err
}

func (r *mockFailingContextInferrer) inferFromConditions(_ []string) (string, error) {
	return "", r.err
}

func (r *mockFailingContextInferrer) inferFromValueExpressions(_ []string) (string, error) {
	return "", r.err
}

type mockStaticContextInferrer struct {
	value string
}

func (r *mockStaticContextInferrer) infer(_, _, _ []string) (string, error) {
	return r.value, nil
}

func (r *mockStaticContextInferrer) inferFromStatements(_ []string) (string, error) {
	return r.value, nil
}

func (r *mockStaticContextInferrer) inferFromConditions(_ []string) (string, error) {
	return r.value, nil
}

func (r *mockStaticContextInferrer) inferFromValueExpressions(_ []string) (string, error) {
	return r.value, nil
}

type mockSetArguments[K any] struct {
	Target Setter[K]
	Value  Getter[K]
}

func Test_NewParserCollection(t *testing.T) {
	settings := componenttest.NewNopTelemetrySettings()
	pc, err := NewParserCollection[any](settings)
	require.NoError(t, err)

	assert.NotNil(t, pc)
	assert.NotNil(t, pc.contextParsers)
	assert.NotNil(t, pc.contextInferrer)
}

func Test_NewParserCollection_OptionError(t *testing.T) {
	_, err := NewParserCollection[any](
		componenttest.NewNopTelemetrySettings(),
		func(_ *ParserCollection[any]) error {
			return errors.New("option error")
		},
	)

	require.Error(t, err, "option error")
}

func Test_WithParserCollectionContext(t *testing.T) {
	ps := mockParser(t, WithPathContextNames[any]([]string{"testContext"}))
	conv := newNopParsedStatementsConverter[any]()
	option := WithParserCollectionContext("testContext", ps, WithStatementConverter(conv))

	pc, err := NewParserCollection(componenttest.NewNopTelemetrySettings(), option)
	require.NoError(t, err)

	pw, exists := pc.contextParsers["testContext"]
	assert.True(t, exists)
	assert.NotNil(t, pw)
}

func Test_ParseStatements_NoStatementConverter(t *testing.T) {
	ps := mockParser(t, WithPathContextNames[any]([]string{"testContext"}))
	option := WithParserCollectionContext[any, any]("testContext", ps)

	pc, err := NewParserCollection[any](componenttest.NewNopTelemetrySettings(), option)
	require.NoError(t, err)

	pw, exists := pc.contextParsers["testContext"]
	assert.True(t, exists)
	assert.NotNil(t, pw)
	_, parseErr := pc.ParseStatementsWithContext("testContext", mockGetter{[]string{`set(testContext.attributes["foo"], "foo")`}}, true)
	assert.Error(t, parseErr)
	assert.Contains(t, parseErr.Error(), "no configured converter for statements")
}

func Test_ParseStatements_NoConditionConverter(t *testing.T) {
	ps := mockParser(t, WithPathContextNames[any]([]string{"testContext"}))
	option := WithParserCollectionContext[any, any]("testContext", ps)

	pc, err := NewParserCollection[any](componenttest.NewNopTelemetrySettings(), option)
	require.NoError(t, err)

	pw, exists := pc.contextParsers["testContext"]
	assert.True(t, exists)
	assert.NotNil(t, pw)
	_, parseErr := pc.ParseConditionsWithContext("testContext", mockGetter{[]string{`foo.attributes["bar"] == "foo"`}}, true)
	assert.Error(t, parseErr)
	assert.Contains(t, parseErr.Error(), "no configured converter for conditions")
}

func Test_WithParserCollectionContext_UnsupportedContext(t *testing.T) {
	ps := mockParser(t, WithPathContextNames[any]([]string{"foo"}))
	conv := newNopParsedStatementsConverter[any]()
	option := WithParserCollectionContext("bar", ps, WithStatementConverter(conv))

	_, err := NewParserCollection[any](componenttest.NewNopTelemetrySettings(), option)

	require.ErrorContains(t, err, `context "bar" must be a valid "*ottl.Parser[interface {}]" path context name`)
}

func Test_WithParserCollectionContext_contextInferrerCandidates(t *testing.T) {
	pc, err := NewParserCollection[any](componenttest.NewNopTelemetrySettings(),
		WithParserCollectionContext("foo", mockParser(t, WithPathContextNames[any]([]string{"foo", "bar"})), WithStatementConverter(newNopParsedStatementsConverter[any]())),
		WithParserCollectionContext("bar", mockParser(t, WithPathContextNames[any]([]string{"bar"})), WithStatementConverter(newNopParsedStatementsConverter[any]())),
	)
	require.NoError(t, err)
	require.NotNil(t, pc.contextInferrer)
	require.Contains(t, pc.contextInferrerCandidates, "foo")

	validEnumSymbol := EnumSymbol("TEST_ENUM")
	invalidEnumSymbol := EnumSymbol("DUMMY")

	fooCandidate := pc.contextInferrerCandidates["foo"]
	assert.NotNil(t, fooCandidate)
	assert.True(t, fooCandidate.hasFunctionName("set"))
	assert.False(t, fooCandidate.hasFunctionName("dummy"))
	assert.True(t, fooCandidate.hasEnumSymbol(&validEnumSymbol))
	assert.False(t, fooCandidate.hasEnumSymbol(&invalidEnumSymbol))
	assert.Nil(t, fooCandidate.getLowerContexts("foo"))

	barCandidate := pc.contextInferrerCandidates["bar"]
	assert.NotNil(t, barCandidate)
	assert.True(t, barCandidate.hasFunctionName("set"))
	assert.False(t, barCandidate.hasFunctionName("dummy"))
	assert.True(t, barCandidate.hasEnumSymbol(&validEnumSymbol))
	assert.False(t, barCandidate.hasEnumSymbol(&invalidEnumSymbol))
	assert.Equal(t, []string{"foo"}, barCandidate.getLowerContexts("bar"))
}

func Test_WithParserCollectionErrorMode(t *testing.T) {
	pc, err := NewParserCollection[any](
		componenttest.NewNopTelemetrySettings(),
		WithParserCollectionErrorMode[any](PropagateError),
	)

	require.NoError(t, err)
	require.NotNil(t, pc)
	require.Equal(t, PropagateError, pc.ErrorMode)
}

func Test_EnableParserCollectionModifiedPathsLogging_True(t *testing.T) {
	ps := mockParser(t, WithPathContextNames[any]([]string{"dummy"}))
	core, observedLogs := observer.New(zap.InfoLevel)
	telemetrySettings := componenttest.NewNopTelemetrySettings()
	telemetrySettings.Logger = zap.New(core)

	pc, err := NewParserCollection(
		telemetrySettings,
		WithParserCollectionContext("dummy", ps, WithStatementConverter(newNopParsedStatementsConverter[any]())),
		EnableParserCollectionModifiedPathsLogging[any](true),
	)
	require.NoError(t, err)

	originalStatements := []string{
		`set(attributes["foo"], "foo")`,
		`set(attributes["bar"], "bar")`,
	}

	_, err = pc.ParseStatementsWithContext("dummy", mockGetter{originalStatements}, true)
	require.NoError(t, err)

	logEntries := observedLogs.TakeAll()
	require.Len(t, logEntries, 1)
	logEntry := logEntries[0]
	require.Equal(t, zap.InfoLevel, logEntry.Level)
	require.Contains(t, logEntry.Message, "one or more paths were modified")
	logEntryStatements := logEntry.ContextMap()["values"]
	require.NotNil(t, logEntryStatements)

	for i, originalStatement := range originalStatements {
		k := fmt.Sprintf("[%d]", i)
		logEntryStatementContext := logEntryStatements.(map[string]any)[k]
		require.Equal(t, logEntryStatementContext.(map[string]any)["original"], originalStatement)
		modifiedStatement, err := ps.prependContextToStatementPaths("dummy", originalStatement)
		require.NoError(t, err)
		require.Equal(t, logEntryStatementContext.(map[string]any)["modified"], modifiedStatement)
	}
}

func Test_EnableParserCollectionModifiedPathsLogging_False(t *testing.T) {
	ps := mockParser(t, WithPathContextNames[any]([]string{"dummy"}))
	core, observedLogs := observer.New(zap.InfoLevel)
	telemetrySettings := componenttest.NewNopTelemetrySettings()
	telemetrySettings.Logger = zap.New(core)

	pc, err := NewParserCollection(
		telemetrySettings,
		WithParserCollectionContext("dummy", ps, WithStatementConverter(newNopParsedStatementsConverter[any]())),
		EnableParserCollectionModifiedPathsLogging[any](false),
	)
	require.NoError(t, err)

	_, err = pc.ParseStatementsWithContext("dummy", mockGetter{[]string{`set(attributes["foo"], "foo")`}}, true)
	require.NoError(t, err)
	require.Empty(t, observedLogs.TakeAll())
}

func Test_NopParsedStatementsConverter(t *testing.T) {
	type dummyContext struct{}

	noop := newNopParsedStatementsConverter[dummyContext]()
	parsedStatements := []*Statement[dummyContext]{{}}
	convertedStatements, err := noop(nil, mockGetter{values: []string{}}, parsedStatements)

	require.NoError(t, err)
	require.NotNil(t, convertedStatements)
	assert.Equal(t, parsedStatements, convertedStatements)
}

func Test_NewParserCollection_DefaultContextInferrer(t *testing.T) {
	pc, err := NewParserCollection[any](componenttest.NewNopTelemetrySettings())
	require.NoError(t, err)
	require.NotNil(t, pc)
	require.NotNil(t, pc.contextInferrer)
}

func Test_ParseStatements_Success(t *testing.T) {
	ps := mockParser(t, WithPathContextNames[any]([]string{"foo"}))

	pc, err := NewParserCollection(
		componenttest.NewNopTelemetrySettings(),
		WithParserCollectionContext("foo", ps, WithStatementConverter(newNopParsedStatementsConverter[any]())),
	)
	require.NoError(t, err)

	statements := mockGetter{values: []string{`set(foo.attributes["bar"], "foo")`, `set(foo.attributes["bar"], "bar")`}}
	result, err := pc.ParseStatements(statements)
	require.NoError(t, err)

	assert.IsType(t, []*Statement[any]{}, result)
	assert.Len(t, result.([]*Statement[any]), 2)
	assert.NotNil(t, result)
}

func Test_ParseStatements_MultipleContexts_Success(t *testing.T) {
	fooParser := mockParser(t, WithPathContextNames[any]([]string{"foo"}))
	barParser := mockParser(t, WithPathContextNames[any]([]string{"bar"}))
	failingConverter := func(
		_ *ParserCollection[any],
		_ StatementsGetter,
		_ []*Statement[any],
	) (any, error) {
		return nil, errors.New("failing converter")
	}

	pc, err := NewParserCollection(
		componenttest.NewNopTelemetrySettings(),
		WithParserCollectionContext("foo", fooParser, WithStatementConverter(failingConverter)),
		WithParserCollectionContext("bar", barParser, WithStatementConverter(newNopParsedStatementsConverter[any]())),
	)
	require.NoError(t, err)

	// The `foo` context is never used, so these statements will successfully parse.
	statements := mockGetter{values: []string{`set(bar.attributes["bar"], "foo")`, `set(bar.attributes["bar"], "bar")`}}
	result, err := pc.ParseStatements(statements)
	require.NoError(t, err)

	assert.IsType(t, []*Statement[any]{}, result)
	assert.Len(t, result.([]*Statement[any]), 2)
	assert.NotNil(t, result)
}

func Test_ParseStatements_NoContextInferredError(t *testing.T) {
	pc, err := NewParserCollection[any](componenttest.NewNopTelemetrySettings())
	require.NoError(t, err)

	statements := mockGetter{values: []string{`set(bar.attributes["bar"], "foo")`}}
	_, err = pc.ParseStatements(statements)

	assert.ErrorContains(t, err, "unable to infer a valid context")
}

func Test_ParseStatements_ContextInferenceError(t *testing.T) {
	pc, err := NewParserCollection[any](componenttest.NewNopTelemetrySettings(), WithParserCollectionContext[any, any]("foo", mockParser(t, WithPathContextNames[any]([]string{"foo"})), WithStatementConverter(newNopParsedStatementsConverter[any]())))
	require.NoError(t, err)
	pc.contextInferrer = &mockFailingContextInferrer{err: errors.New("inference error")}

	statements := mockGetter{values: []string{`set(bar.attributes["bar"], "foo")`}}
	_, err = pc.ParseStatements(statements)

	assert.ErrorContains(t, err, "inference error")
}

func Test_ParseStatements_UnknownContextError(t *testing.T) {
	pc, err := NewParserCollection[any](componenttest.NewNopTelemetrySettings(),
		WithParserCollectionContext("bar", mockParser(t, WithPathContextNames[any]([]string{"bar"})), WithStatementConverter(newNopParsedStatementsConverter[any]())),
		WithParserCollectionContext("te", mockParser(t, WithPathContextNames[any]([]string{"te"})), WithStatementConverter(newNopParsedStatementsConverter[any]())),
	)
	require.NoError(t, err)
	pc.contextInferrer = &mockStaticContextInferrer{"foo"}

	statements := mockGetter{values: []string{`set(foo.attributes["bar"], "foo")`}}
	_, err = pc.ParseStatements(statements)

	assert.ErrorContains(t, err, `context "foo" inferred from the statements`)
	assert.ErrorContains(t, err, "is not a supported context")
}

func Test_ParseStatements_ParseStatementsError(t *testing.T) {
	ps := mockParser(t, WithPathContextNames[any]([]string{"foo"}))
	ps.pathParser = func(_ Path[any]) (GetSetter[any], error) {
		return nil, errors.New("parse statements error")
	}

	pc, err := NewParserCollection(
		componenttest.NewNopTelemetrySettings(),
		WithParserCollectionContext("foo", ps, WithStatementConverter(newNopParsedStatementsConverter[any]())),
	)
	require.NoError(t, err)

	statements := mockGetter{values: []string{`set(foo.attributes["bar"], "foo")`}}
	_, err = pc.ParseStatements(statements)
	assert.ErrorContains(t, err, "parse statements error")
}

func Test_ParseStatements_ConverterError(t *testing.T) {
	ps := mockParser(t, WithPathContextNames[any]([]string{"dummy"}))
	conv := func(_ *ParserCollection[any], _ StatementsGetter, _ []*Statement[any]) (any, error) {
		return nil, errors.New("converter error")
	}

	pc, err := NewParserCollection(
		componenttest.NewNopTelemetrySettings(),
		WithParserCollectionContext("dummy", ps, WithStatementConverter(conv)),
	)
	require.NoError(t, err)

	statements := mockGetter{values: []string{`set(dummy.attributes["bar"], "foo")`}}
	_, err = pc.ParseStatements(statements)

	assert.EqualError(t, err, "converter error")
}

func Test_ParseStatements_ConverterNilReturn(t *testing.T) {
	ps := mockParser(t, WithPathContextNames[any]([]string{"dummy"}))
	conv := func(_ *ParserCollection[any], _ StatementsGetter, _ []*Statement[any]) (any, error) {
		return nil, nil
	}

	pc, err := NewParserCollection(
		componenttest.NewNopTelemetrySettings(),
		WithParserCollectionContext("dummy", ps, WithStatementConverter(conv)),
	)
	require.NoError(t, err)

	statements := mockGetter{values: []string{`set(dummy.attributes["bar"], "foo")`}}
	result, err := pc.ParseStatements(statements)
	assert.NoError(t, err)
	assert.Nil(t, result)
}

func Test_ParseStatements_StatementsConverterGetterType(t *testing.T) {
	ps := mockParser(t, WithPathContextNames[any]([]string{"dummy"}))
	statements := mockGetter{values: []string{`set(dummy.attributes["bar"], "foo")`}}
	conv := func(_ *ParserCollection[any], statementsGetter StatementsGetter, _ []*Statement[any]) (any, error) {
		switch statementsGetter.(type) {
		case mockGetter:
			return statements, nil
		default:
			return nil, fmt.Errorf("invalid StatementsGetter type, expected: mockGetter, got: %T", statementsGetter)
		}
	}

	pc, err := NewParserCollection(componenttest.NewNopTelemetrySettings(), WithParserCollectionContext("dummy", ps, WithStatementConverter(conv)))
	require.NoError(t, err)

	_, err = pc.ParseStatements(statements)
	require.NoError(t, err)
}

func Test_ParseStatements_WithContextInferenceConditions(t *testing.T) {
	metricParser := mockParser(t, WithPathContextNames[any]([]string{"metric", "resource"}))
	resourceParser := mockParser(t, WithPathContextNames[any]([]string{"resource"}))

	failingConverter := func(
		_ *ParserCollection[any],
		_ StatementsGetter,
		_ []*Statement[any],
	) (any, error) {
		return nil, errors.New(`invalid context inferred, got: "resource", expected: "metric"`)
	}

	pc, err := NewParserCollection(
		componenttest.NewNopTelemetrySettings(),
		WithParserCollectionContext("metric", metricParser, WithStatementConverter(newNopParsedStatementsConverter[any]())),
		WithParserCollectionContext("resource", resourceParser, WithStatementConverter(failingConverter)),
	)

	require.NoError(t, err)

	result, err := pc.ParseStatements(
		NewStatementsGetter([]string{`set(resource.attributes["bar"], "bar")`}),
		WithContextInferenceConditions([]string{`metric.attributes["foo"] == "foo"`}),
	)

	require.NoError(t, err)
	require.NotNil(t, result)
}

func Test_ParseStatementsWithContext_UnknownContextError(t *testing.T) {
	pc, err := NewParserCollection[any](componenttest.NewNopTelemetrySettings())
	require.NoError(t, err)

	statements := mockGetter{[]string{`set(attributes["bar"], "bar")`}}
	_, err = pc.ParseStatementsWithContext("bar", statements, false)

	assert.ErrorContains(t, err, `unknown context "bar"`)
}

func Test_ParseStatementsWithContext_PrependPathContext(t *testing.T) {
	ps := mockParser(t, WithPathContextNames[any]([]string{"dummy"}))
	pc, err := NewParserCollection(
		componenttest.NewNopTelemetrySettings(),
		WithParserCollectionContext("dummy", ps, WithStatementConverter(newNopParsedStatementsConverter[any]())),
	)
	require.NoError(t, err)

	result, err := pc.ParseStatementsWithContext(
		"dummy",
		mockGetter{[]string{
			`set(attributes["foo"], "foo")`,
			`set(attributes["bar"], "bar")`,
		}},
		true,
	)

	require.NoError(t, err)
	require.Len(t, result, 2)
	parsedStatements := result.([]*Statement[any])
	assert.Equal(t, `set(dummy.attributes["foo"], "foo")`, parsedStatements[0].origText)
	assert.Equal(t, `set(dummy.attributes["bar"], "bar")`, parsedStatements[1].origText)
}

func Test_NewStatementsGetter(t *testing.T) {
	statements := []string{`set(foo, "bar")`, `set(bar, "foo")`}
	statementsGetter := NewStatementsGetter(statements)
	assert.Implements(t, (*StatementsGetter)(nil), statementsGetter)
	assert.Equal(t, statements, statementsGetter.GetStatements())
}

func Test_ParseConditions_Success(t *testing.T) {
	ps := mockParser(t, WithPathContextNames[any]([]string{"foo"}))

	pc, err := NewParserCollection(
		componenttest.NewNopTelemetrySettings(),
		WithParserCollectionContext("foo", ps, WithConditionConverter(newNopParsedConditionsConverter[any]())),
	)
	require.NoError(t, err)

	conditions := mockGetter{values: []string{`foo.attributes["bar"] == "foo"`, `foo.attributes["bar"] == "bar"`}}
	result, err := pc.ParseConditions(conditions)
	require.NoError(t, err)

	assert.IsType(t, []*Condition[any]{}, result)
	assert.Len(t, result.([]*Condition[any]), 2)
	assert.NotNil(t, result)
}

func Test_ParseConditions_MultipleContexts_Success(t *testing.T) {
	fooParser := mockParser(t, WithPathContextNames[any]([]string{"foo"}))
	barParser := mockParser(t, WithPathContextNames[any]([]string{"bar"}))
	failingConverter := func(
		_ *ParserCollection[any],
		_ ConditionsGetter,
		_ []*Condition[any],
	) (any, error) {
		return nil, errors.New("failing converter")
	}

	pc, err := NewParserCollection(
		componenttest.NewNopTelemetrySettings(),
		WithParserCollectionContext("foo", fooParser, WithConditionConverter(failingConverter)),
		WithParserCollectionContext("bar", barParser, WithConditionConverter(newNopParsedConditionsConverter[any]())),
	)
	require.NoError(t, err)

	// The `foo` context is never used, so these conditions will successfully parse.
	conditions := mockGetter{values: []string{`bar.attributes["bar"] == "foo"`, `bar.attributes["bar"] == "bar"`}}
	result, err := pc.ParseConditions(conditions)
	require.NoError(t, err)

	assert.IsType(t, []*Condition[any]{}, result)
	assert.Len(t, result.([]*Condition[any]), 2)
	assert.NotNil(t, result)
}

func Test_ParseConditions_NoContextInferredError(t *testing.T) {
	pc, err := NewParserCollection[any](componenttest.NewNopTelemetrySettings())
	require.NoError(t, err)
	pc.contextInferrer = &mockStaticContextInferrer{""}

	conditions := mockGetter{values: []string{`bar.attributes["bar"] == "foo"`}}
	_, err = pc.ParseConditions(conditions)

	assert.ErrorContains(t, err, "unable to infer context from conditions")
}

func Test_ParseConditions_ContextInferenceError(t *testing.T) {
	pc, err := NewParserCollection[any](componenttest.NewNopTelemetrySettings())
	require.NoError(t, err)
	pc.contextInferrer = &mockFailingContextInferrer{err: errors.New("inference error")}

	conditions := mockGetter{values: []string{`bar.attributes["bar"] == "foo"`}}
	_, err = pc.ParseConditions(conditions)

	assert.EqualError(t, err, "inference error")
}

func Test_ParseConditions_UnknownContextError(t *testing.T) {
	pc, err := NewParserCollection[any](componenttest.NewNopTelemetrySettings(),
		WithParserCollectionContext("bar", mockParser(t, WithPathContextNames[any]([]string{"bar"})), WithConditionConverter(newNopParsedConditionsConverter[any]())),
		WithParserCollectionContext("te", mockParser(t, WithPathContextNames[any]([]string{"te"})), WithConditionConverter(newNopParsedConditionsConverter[any]())),
	)
	require.NoError(t, err)

	conditions := mockGetter{values: []string{`foo.attributes["bar"] == "foo"`}}
	_, err = pc.ParseConditions(conditions)

	assert.ErrorContains(t, err, `context "foo" inferred from the conditions`)
	assert.ErrorContains(t, err, "is not a supported context")
}

func Test_ParseConditions_ParseConditionsError(t *testing.T) {
	ps := mockParser(t, WithPathContextNames[any]([]string{"foo"}))
	ps.pathParser = func(_ Path[any]) (GetSetter[any], error) {
		return nil, errors.New("parse conditions error")
	}

	pc, err := NewParserCollection(
		componenttest.NewNopTelemetrySettings(),
		WithParserCollectionContext("foo", ps, WithConditionConverter(newNopParsedConditionsConverter[any]())),
	)
	require.NoError(t, err)

	conditions := mockGetter{values: []string{`foo.attributes["bar"] == "foo"`}}
	_, err = pc.ParseConditions(conditions)
	assert.ErrorContains(t, err, "parse conditions error")
}

func Test_ParseConditions_ConverterError(t *testing.T) {
	ps := mockParser(t, WithPathContextNames[any]([]string{"dummy"}))
	conv := func(_ *ParserCollection[any], _ ConditionsGetter, _ []*Condition[any]) (any, error) {
		return nil, errors.New("converter error")
	}

	pc, err := NewParserCollection(
		componenttest.NewNopTelemetrySettings(),
		WithParserCollectionContext("dummy", ps, WithConditionConverter(conv)),
	)
	require.NoError(t, err)

	conditions := mockGetter{values: []string{`dummy.attributes["bar"] == "foo"`}}
	_, err = pc.ParseConditions(conditions)

	assert.EqualError(t, err, "converter error")
}

func Test_ParseConditions_ConverterNilReturn(t *testing.T) {
	ps := mockParser(t, WithPathContextNames[any]([]string{"dummy"}))
	conv := func(_ *ParserCollection[any], _ ConditionsGetter, _ []*Condition[any]) (any, error) {
		return nil, nil
	}

	pc, err := NewParserCollection(
		componenttest.NewNopTelemetrySettings(),
		WithParserCollectionContext("dummy", ps, WithConditionConverter(conv)),
	)
	require.NoError(t, err)

	conditions := mockGetter{values: []string{`dummy.attributes["bar"] == "foo"`}}
	result, err := pc.ParseConditions(conditions)
	assert.NoError(t, err)
	assert.Nil(t, result)
}

func Test_ParseConditions_ConditionsConverterGetterType(t *testing.T) {
	ps := mockParser(t, WithPathContextNames[any]([]string{"dummy"}))
	conditions := mockGetter{values: []string{`dummy.attributes["bar"] == "foo"`}}
	conv := func(_ *ParserCollection[any], conditionsGetter ConditionsGetter, _ []*Condition[any]) (any, error) {
		switch conditionsGetter.(type) {
		case mockGetter:
			return conditions, nil
		default:
			return nil, fmt.Errorf("invalid ConditionsGetter type, expected: mockGetter, got: %T", conditionsGetter)
		}
	}

	pc, err := NewParserCollection(componenttest.NewNopTelemetrySettings(), WithParserCollectionContext("dummy", ps, WithConditionConverter(conv)))
	require.NoError(t, err)

	_, err = pc.ParseConditions(conditions)
	require.NoError(t, err)
}

func Test_ParseConditionsWithContext_UnknownContextError(t *testing.T) {
	pc, err := NewParserCollection[any](componenttest.NewNopTelemetrySettings())
	require.NoError(t, err)

	conditions := mockGetter{[]string{`attributes["bar"] == "bar"`}}
	_, err = pc.ParseConditionsWithContext("bar", conditions, false)

	assert.ErrorContains(t, err, `unknown context "bar"`)
}

func Test_ParseConditionsWithContext_PrependPathContext(t *testing.T) {
	ps := mockParser(t, WithPathContextNames[any]([]string{"dummy"}))
	pc, err := NewParserCollection(
		componenttest.NewNopTelemetrySettings(),
		WithParserCollectionContext("dummy", ps, WithConditionConverter(newNopParsedConditionsConverter[any]())),
	)
	require.NoError(t, err)

	result, err := pc.ParseConditionsWithContext(
		"dummy",
		mockGetter{[]string{
			`attributes["foo"] == "foo"`,
			`attributes["bar"] == "bar"`,
		}},
		true,
	)

	require.NoError(t, err)
	require.Len(t, result, 2)
	parsedConditions := result.([]*Condition[any])
	assert.Equal(t, `dummy.attributes["foo"] == "foo"`, parsedConditions[0].origText)
	assert.Equal(t, `dummy.attributes["bar"] == "bar"`, parsedConditions[1].origText)
}

func Test_NewConditionsGetter(t *testing.T) {
	conditions := []string{`foo == "bar"`, `bar == "foo"`}
	conditionsGetter := NewConditionsGetter(conditions)
	assert.Implements(t, (*ConditionsGetter)(nil), conditionsGetter)
	assert.Equal(t, conditions, conditionsGetter.GetConditions())
}

func Test_ParseValueExpressions_Success(t *testing.T) {
	ps := mockParser(t, WithPathContextNames[any]([]string{"foo"}))

	pc, err := NewParserCollection(
		componenttest.NewNopTelemetrySettings(),
		WithParserCollectionContext("foo", ps, WithValueExpressionConverter(newNopParsedValueExpressionsConverter[any]())),
	)
	require.NoError(t, err)

	expressions := mockGetter{values: []string{`foo.attributes["bar"]`, `foo.attributes["bar"]`}}
	result, err := pc.ParseValueExpressions(expressions)
	require.NoError(t, err)

	assert.IsType(t, []*ValueExpression[any]{}, result)
	assert.Len(t, result.([]*ValueExpression[any]), 2)
	assert.NotNil(t, result)
}

func Test_ParseValueExpressions_MultipleContexts_Success(t *testing.T) {
	fooParser := mockParser(t, WithPathContextNames[any]([]string{"foo"}))
	barParser := mockParser(t, WithPathContextNames[any]([]string{"bar"}))
	failingConverter := func(
		_ *ParserCollection[any],
		_ ValueExpressionsGetter,
		_ []*ValueExpression[any],
	) (any, error) {
		return nil, errors.New("failing converter")
	}

	pc, err := NewParserCollection(
		componenttest.NewNopTelemetrySettings(),
		WithParserCollectionContext("foo", fooParser, WithValueExpressionConverter(failingConverter)),
		WithParserCollectionContext("bar", barParser, WithValueExpressionConverter(newNopParsedValueExpressionsConverter[any]())),
	)
	require.NoError(t, err)

	// The `foo` context is never used, so these expressions will successfully parse.
	expressions := mockGetter{values: []string{`bar.attributes["foo"]`, `bar.attributes["bar"]`}}
	result, err := pc.ParseValueExpressions(expressions)
	require.NoError(t, err)

	assert.IsType(t, []*ValueExpression[any]{}, result)
	assert.Len(t, result.([]*ValueExpression[any]), 2)
	assert.NotNil(t, result)
}

func Test_ParseValueExpressions_NoContextInferredError(t *testing.T) {
	pc, err := NewParserCollection[any](componenttest.NewNopTelemetrySettings())
	require.NoError(t, err)
	pc.contextInferrer = &mockStaticContextInferrer{""}

	expressions := mockGetter{values: []string{`bar.attributes["foo"]`}}
	_, err = pc.ParseValueExpressions(expressions)

	assert.ErrorContains(t, err, "unable to infer context from expressions")
}

func Test_ParseValueExpressions_ContextInferenceError(t *testing.T) {
	pc, err := NewParserCollection[any](componenttest.NewNopTelemetrySettings())
	require.NoError(t, err)
	pc.contextInferrer = &mockFailingContextInferrer{err: errors.New("inference error")}

	expressions := mockGetter{values: []string{`bar.attributes["foo"]`}}
	_, err = pc.ParseValueExpressions(expressions)

	assert.EqualError(t, err, "inference error")
}

func Test_ParseValueExpressions_UnknownContextError(t *testing.T) {
	pc, err := NewParserCollection[any](componenttest.NewNopTelemetrySettings(),
		WithParserCollectionContext("bar", mockParser(t, WithPathContextNames[any]([]string{"bar"})), WithValueExpressionConverter(newNopParsedValueExpressionsConverter[any]())),
		WithParserCollectionContext("te", mockParser(t, WithPathContextNames[any]([]string{"te"})), WithValueExpressionConverter(newNopParsedValueExpressionsConverter[any]())),
	)
	require.NoError(t, err)

	expressions := mockGetter{values: []string{`foo.attributes["foo"]`}}
	_, err = pc.ParseValueExpressions(expressions)

	assert.ErrorContains(t, err, `context "foo" inferred from the expressions`)
	assert.ErrorContains(t, err, "is not a supported context")
}

func Test_ParseValueExpressions_ParseValueExpressionsError(t *testing.T) {
	ps := mockParser(t, WithPathContextNames[any]([]string{"foo"}))
	ps.pathParser = func(_ Path[any]) (GetSetter[any], error) {
		return nil, errors.New("parse expressions error")
	}

	pc, err := NewParserCollection(
		componenttest.NewNopTelemetrySettings(),
		WithParserCollectionContext("foo", ps, WithValueExpressionConverter(newNopParsedValueExpressionsConverter[any]())),
	)
	require.NoError(t, err)

	expressions := mockGetter{values: []string{`foo.attributes["bar"]`}}
	_, err = pc.ParseValueExpressions(expressions)
	assert.ErrorContains(t, err, "parse expressions error")
}

func Test_ParseValueExpressions_ConverterError(t *testing.T) {
	ps := mockParser(t, WithPathContextNames[any]([]string{"dummy"}))
	conv := func(_ *ParserCollection[any], _ ValueExpressionsGetter, _ []*ValueExpression[any]) (any, error) {
		return nil, errors.New("converter error")
	}

	pc, err := NewParserCollection(
		componenttest.NewNopTelemetrySettings(),
		WithParserCollectionContext("dummy", ps, WithValueExpressionConverter(conv)),
	)
	require.NoError(t, err)

	expressions := mockGetter{values: []string{`dummy.attributes["bar"]`}}
	_, err = pc.ParseValueExpressions(expressions)

	assert.EqualError(t, err, "converter error")
}

func Test_ParseValueExpressions_ConverterNilReturn(t *testing.T) {
	ps := mockParser(t, WithPathContextNames[any]([]string{"dummy"}))
	conv := func(_ *ParserCollection[any], _ ValueExpressionsGetter, _ []*ValueExpression[any]) (any, error) {
		return nil, nil
	}

	pc, err := NewParserCollection(
		componenttest.NewNopTelemetrySettings(),
		WithParserCollectionContext("dummy", ps, WithValueExpressionConverter(conv)),
	)
	require.NoError(t, err)

	expressions := mockGetter{values: []string{`dummy.attributes["bar"]`}}
	result, err := pc.ParseValueExpressions(expressions)
	assert.NoError(t, err)
	assert.Nil(t, result)
}

func Test_ParseValueExpressions_ValueExpressionsConverterGetterType(t *testing.T) {
	ps := mockParser(t, WithPathContextNames[any]([]string{"dummy"}))
	expressions := mockGetter{values: []string{`dummy.attributes["bar"]`}}
	conv := func(_ *ParserCollection[any], expressionsGetter ValueExpressionsGetter, _ []*ValueExpression[any]) (any, error) {
		switch expressionsGetter.(type) {
		case mockGetter:
			return expressions, nil
		default:
			return nil, fmt.Errorf("invalid ValueExpressionsGetter type, expected: mockGetter, got: %T", expressionsGetter)
		}
	}

	pc, err := NewParserCollection(componenttest.NewNopTelemetrySettings(), WithParserCollectionContext("dummy", ps, WithValueExpressionConverter(conv)))
	require.NoError(t, err)

	_, err = pc.ParseValueExpressions(expressions)
	require.NoError(t, err)
}

func Test_ParseValueExpressions_WithContextInferenceConditions(t *testing.T) {
	metricParser := mockParser(t, WithPathContextNames[any]([]string{"metric", "resource"}))
	resourceParser := mockParser(t, WithPathContextNames[any]([]string{"resource"}))

	failingConverter := func(
		_ *ParserCollection[any],
		_ ValueExpressionsGetter,
		_ []*ValueExpression[any],
	) (any, error) {
		return nil, errors.New(`invalid context inferred, got: "resource", expected: "metric"`)
	}

	pc, err := NewParserCollection(
		componenttest.NewNopTelemetrySettings(),
		WithParserCollectionContext("metric", metricParser, WithValueExpressionConverter(newNopParsedValueExpressionsConverter[any]())),
		WithParserCollectionContext("resource", resourceParser, WithValueExpressionConverter(failingConverter)),
	)

	require.NoError(t, err)

	result, err := pc.ParseValueExpressions(
		NewValueExpressionsGetter([]string{`resource.attributes["bar"]`}),
		WithContextInferenceConditions([]string{`metric.attributes["foo"] == "foo"`}),
	)

	require.NoError(t, err)
	require.NotNil(t, result)
}

func Test_ParseValueExpressionsWithContext_UnknownContextError(t *testing.T) {
	pc, err := NewParserCollection[any](componenttest.NewNopTelemetrySettings())
	require.NoError(t, err)

	expressions := mockGetter{[]string{`attributes["bar"]`}}
	_, err = pc.ParseValueExpressionsWithContext("bar", expressions, false)

	assert.ErrorContains(t, err, `unknown context "bar"`)
}

func Test_ParseValueExpressionsWithContext_PrependPathContext(t *testing.T) {
	ps := mockParser(t, WithPathContextNames[any]([]string{"dummy"}))
	pc, err := NewParserCollection(
		componenttest.NewNopTelemetrySettings(),
		WithParserCollectionContext("dummy", ps, WithValueExpressionConverter(newNopParsedValueExpressionsConverter[any]())),
	)
	require.NoError(t, err)

	result, err := pc.ParseValueExpressionsWithContext(
		"dummy",
		mockGetter{[]string{
			`attributes["foo"]`,
			`attributes["bar"]`,
		}},
		true,
	)

	require.NoError(t, err)
	require.Len(t, result, 2)
	parsedValueExpressions := result.([]*ValueExpression[any])
	assert.Equal(t, `dummy.attributes["foo"]`, parsedValueExpressions[0].origText)
	assert.Equal(t, `dummy.attributes["bar"]`, parsedValueExpressions[1].origText)
}

func Test_NewValueExpressionsGetter(t *testing.T) {
	expressions := []string{`foo`, `bar`}
	expressionsGetter := NewValueExpressionsGetter(expressions)
	assert.Implements(t, (*ValueExpressionsGetter)(nil), expressionsGetter)
	assert.Equal(t, expressions, expressionsGetter.GetValueExpressions())
}

func mockParser(t *testing.T, options ...Option[any]) *Parser[any] {
	mockSetFactory := NewFactory("set", &mockSetArguments[any]{},
		func(_ FunctionContext, _ Arguments) (ExprFunc[any], error) {
			return func(context.Context, any) (any, error) {
				return nil, nil
			}, nil
		})

	ps, err := NewParser(
		CreateFactoryMap[any](mockSetFactory),
		testParsePath[any],
		componenttest.NewNopTelemetrySettings(),
		append([]Option[any]{
			WithEnumParser[any](testParseEnum),
		}, options...)...,
	)

	require.NoError(t, err)
	return &ps
}
