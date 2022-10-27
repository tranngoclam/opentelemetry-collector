// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package component

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewExporterFactory(t *testing.T) {
	const typeStr = "test"
	defaultCfg := NewExporterConfigSettings(NewID(typeStr))
	factory := NewExporterFactory(
		typeStr,
		func() ExporterConfig { return &defaultCfg })
	assert.EqualValues(t, typeStr, factory.Type())
	assert.EqualValues(t, &defaultCfg, factory.CreateDefaultConfig())
	_, err := factory.CreateTracesExporter(context.Background(), ExporterCreateSettings{}, &defaultCfg)
	assert.Error(t, err)
	_, err = factory.CreateMetricsExporter(context.Background(), ExporterCreateSettings{}, &defaultCfg)
	assert.Error(t, err)
	_, err = factory.CreateLogsExporter(context.Background(), ExporterCreateSettings{}, &defaultCfg)
	assert.Error(t, err)
}

func TestNewExporterFactory_WithOptions(t *testing.T) {
	const typeStr = "test"
	defaultCfg := NewExporterConfigSettings(NewID(typeStr))
	factory := NewExporterFactory(
		typeStr,
		func() ExporterConfig { return &defaultCfg },
		WithTracesExporter(createTracesExporter, StabilityLevelInDevelopment),
		WithMetricsExporter(createMetricsExporter, StabilityLevelAlpha),
		WithLogsExporter(createLogsExporter, StabilityLevelDeprecated))
	assert.EqualValues(t, typeStr, factory.Type())
	assert.EqualValues(t, &defaultCfg, factory.CreateDefaultConfig())

	assert.Equal(t, StabilityLevelInDevelopment, factory.TracesExporterStability())
	_, err := factory.CreateTracesExporter(context.Background(), ExporterCreateSettings{}, &defaultCfg)
	assert.NoError(t, err)

	assert.Equal(t, StabilityLevelAlpha, factory.MetricsExporterStability())
	_, err = factory.CreateMetricsExporter(context.Background(), ExporterCreateSettings{}, &defaultCfg)
	assert.NoError(t, err)

	assert.Equal(t, StabilityLevelDeprecated, factory.LogsExporterStability())
	_, err = factory.CreateLogsExporter(context.Background(), ExporterCreateSettings{}, &defaultCfg)
	assert.NoError(t, err)
}

func createTracesExporter(context.Context, ExporterCreateSettings, ExporterConfig) (TracesExporter, error) {
	return nil, nil
}

func createMetricsExporter(context.Context, ExporterCreateSettings, ExporterConfig) (MetricsExporter, error) {
	return nil, nil
}

func createLogsExporter(context.Context, ExporterCreateSettings, ExporterConfig) (LogsExporter, error) {
	return nil, nil
}
