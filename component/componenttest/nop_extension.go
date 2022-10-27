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

package componenttest // import "go.opentelemetry.io/collector/component/componenttest"

import (
	"context"

	"go.opentelemetry.io/collector/component"
)

// NewNopExtensionCreateSettings returns a new nop settings for Create*Extension functions.
func NewNopExtensionCreateSettings() component.ExtensionCreateSettings {
	return component.ExtensionCreateSettings{
		TelemetrySettings: NewNopTelemetrySettings(),
		BuildInfo:         component.NewDefaultBuildInfo(),
	}
}

type nopExtensionConfig struct {
	component.ExtensionConfigSettings `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct
}

// NewNopExtensionFactory returns a component.ExtensionFactory that constructs nop extensions.
func NewNopExtensionFactory() component.ExtensionFactory {
	return component.NewExtensionFactory(
		"nop",
		func() component.ExtensionConfig {
			return &nopExtensionConfig{
				ExtensionConfigSettings: component.NewExtensionConfigSettings(component.NewID("nop")),
			}
		},
		func(context.Context, component.ExtensionCreateSettings, component.ExtensionConfig) (component.Extension, error) {
			return nopExtensionInstance, nil
		},
		component.StabilityLevelStable)
}

var nopExtensionInstance = &nopExtension{}

// nopExtension stores consumed traces and metrics for testing purposes.
type nopExtension struct {
	nopComponent
}
