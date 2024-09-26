// Copyright 2024 James Toyer
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewFieldDefinition(t *testing.T) {
	tests := map[string]struct {
		name  string
		value interface{}
		want  FieldDefinition
	}{
		"empty string value": {
			name:  "hostname",
			value: "",
			want: FieldDefinition{
				Name:     "Hostname",
				JSONName: "hostname",
				Type:     String,
			},
		},
		"value has validation for a string": {
			name:  "mac",
			value: `^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$`,
			want: FieldDefinition{
				Name:     "Mac",
				JSONName: "mac",
				Type:     String,
			},
		},
		"value has false|true validation": {
			name:  "use_fixedip",
			value: `false|true`,
			want: FieldDefinition{
				Name:     "UseFixedip",
				JSONName: "use_fixedip",
				Type:     Boolean,
			},
		},
		"value has true|false validation": {
			name:  "use_fixedip",
			value: `true|false`,
			want: FieldDefinition{
				Name:     "UseFixedip",
				JSONName: "use_fixedip",
				Type:     Boolean,
			},
		},
		"value has validation for integer": {
			name:  "qos_rate_max_down",
			value: `^-1|[2-9]|[1-9][0-9]{1,4}|100000$`,
			want: FieldDefinition{
				Name:     "QosRateMaxDown",
				JSONName: "qos_rate_max_down",
				Type:     Number,
			},
		},
		"value has validation for float": {
			name:  "heightInMeters",
			value: `^([-]?[\d]+[.]?[\d]*)$`,
			want: FieldDefinition{
				Name:     "HeightInMeters",
				JSONName: "heightInMeters",
				Type:     Decimal,
			},
		},
		"value has an empty list of validations": {
			name:  "igmp_proxy_downstream_networkconf_ids",
			value: []interface{}{},
			want: FieldDefinition{
				Name:     "IgmpProxyDownstreamNetworkconfIds",
				JSONName: "igmp_proxy_downstream_networkconf_ids",
				Type:     List(String),
			},
		},
		"value has a list with empty string validation": {
			name:  "enabled_networks",
			value: []interface{}{""},
			want: FieldDefinition{
				Name:     "EnabledNetworks",
				JSONName: "enabled_networks",
				Type:     List(String),
			},
		},
		"value has a list with string validation": {
			name:  "options",
			value: []interface{}{`^[^"' ]+$`},
			want: FieldDefinition{
				Name:     "Options",
				JSONName: "options",
				Type:     List(String),
			},
		},
		"value has a list with number validation": {
			name:  "auth_ids",
			value: []interface{}{"0|1|2|3|4|5"},
			want: FieldDefinition{
				Name:     "AuthIds",
				JSONName: "auth_ids",
				Type:     List(Number),
			},
		},
		"value has a list with decimal validation": {
			name:  "a",
			value: []interface{}{`^([-]?[\d]+[.]?[\d]*)$`},
			want: FieldDefinition{
				Name:     "A",
				JSONName: "a",
				Type:     List(Decimal),
			},
		},
		"value is an object": {
			name: "config",
			value: map[string]interface{}{
				"list":   []interface{}{},
				"number": `[0-9]`,
				"string": ``,
			},
			want: FieldDefinition{
				Name:     "Config",
				JSONName: "config",
				Type:     Object("FieldTestConfig"),
				Object: &EndpointObject{
					Name: "FieldTestConfig",
					Fields: []FieldDefinition{
						{
							Name:     "List",
							JSONName: "list",
							Type:     List(String),
						},
						{
							Name:     "Number",
							JSONName: "number",
							Type:     Number,
						},
						{
							Name:     "String",
							JSONName: "string",
							Type:     String,
						},
					},
				},
			},
		},
		"value is a nested object": {
			name: "config",
			value: map[string]interface{}{
				"nested": map[string]interface{}{
					"number": `[0-9]`,
				},
				"string": ``,
			},
			want: FieldDefinition{
				Name:     "Config",
				JSONName: "config",
				Type:     Object("FieldTestConfig"),
				Object: &EndpointObject{
					Name: "FieldTestConfig",
					Fields: []FieldDefinition{
						{
							Name:     "Nested",
							JSONName: "nested",
							Type:     Object("FieldTestConfigNested"),
							Object: &EndpointObject{
								Name: "FieldTestConfigNested",
								Fields: []FieldDefinition{
									{
										Name:     "Number",
										JSONName: "number",
										Type:     Number,
									},
								},
							},
						},
						{
							Name:     "String",
							JSONName: "string",
							Type:     String,
						},
					},
				},
			},
		},
		"value has a list with object value": {
			name: "satisfaction_table",
			value: []interface{}{
				map[string]interface{}{
					"device_mac":   `^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$`,
					"satisfaction": `^[0-9]+\.?[0-9]*$`,
				},
			},
			want: FieldDefinition{
				Name:     "SatisfactionTable",
				JSONName: "satisfaction_table",
				Type:     List(Object("FieldTestSatisfactionTable")),
				Object: &EndpointObject{
					Name: "FieldTestSatisfactionTable",
					Fields: []FieldDefinition{
						{
							Name:     "DeviceMac",
							JSONName: "device_mac",
							Type:     String,
						},
						{
							Name:     "Satisfaction",
							JSONName: "satisfaction",
							Type:     String,
						},
					},
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := NewFieldDefinition(test.name, "FieldTest", test.value)
			require.NoError(t, err)

			assert.Equal(t, test.want, got)
		})
	}
}

func Test_normaliseValidationString(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"empty": {
			input: "",
			want:  "",
		},
		"simple string": {
			input: "auto|wpa1|wpa2",
			want:  "autowpa1wpa2",
		},
		"length limited wildcard string": {
			input: `.{1,128}`,
			want:  ".",
		},
		"complex string": {
			input: `(sun|mon|tue|wed|thu|fri|sat)(\-(sun|mon|tue|wed|thu|fri|sat))?\|([0-2][0-9][0-5][0-9])\-([0-2][0-9][0-5][0-9])`,
			want:  `sunmontuewedthufrisat\sunmontuewedthufrisat\02090509\02090509`,
		},
		"cidr string": {
			input: `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\/([8-9]|[1-2][0-9]|3[0-2])$|^$`,
			want:  `091909109204092505\.091909109204092505\/891209302`,
		},
		"ip string": {
			input: `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$|^$`,
			want:  `091909109204092505\.091909109204092505`,
		},
		"simple integer": {
			input: "[0-7]+",
			want:  "07",
		},
		"complex integer": {
			input: "^-1|[2-9]|[1-9][0-9]{1,4}|100000$",
			want:  "1291909100000",
		},
		"simple float": {
			input: "^([\\d]+[.]?[\\d]*)$",
			want:  "09.09",
		},
		"complex float": {
			input: `^([-]?[\d]+[.]?[\d]*)$`,
			want:  "09.09",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := normaliseValidationString(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}
