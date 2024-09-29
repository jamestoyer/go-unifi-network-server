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

package spec

import (
	"regexp"
	"testing"

	"github.com/jamestoyer/go-unifi-network-server/networkserver"
	"github.com/stretchr/testify/assert"
)

func TestField_ApplyOverrides(t *testing.T) {
	defaultField := Field{
		Name:        "Default",
		Description: "A default description",
		Type:        FieldTypeBoolean,
		JSONName:    "default",
		Validation:  regexp.MustCompile("false|true"),
	}

	tests := map[string]struct {
		defaultDescription bool
		overrides          *FieldOverrides
		want               Field
	}{
		"overrides is nil": {
			overrides: nil,
			want: Field{
				Name:        "Default",
				Description: "A default description",
				Type:        FieldTypeBoolean,
				JSONName:    "default",
				Validation:  regexp.MustCompile("false|true"),
			},
		},
		"overrides contains no overrides": {
			overrides: &FieldOverrides{},
			want: Field{
				Name:        "Default",
				Description: "A default description",
				Type:        FieldTypeBoolean,
				JSONName:    "default",
				Validation:  regexp.MustCompile("false|true"),
			},
		},
		"name is overridden": {
			overrides: &FieldOverrides{
				Name: networkserver.String("OverriddenName"),
			},
			want: Field{
				Name:        "OverriddenName",
				Description: "A default description",
				Type:        FieldTypeBoolean,
				JSONName:    "default",
				Validation:  regexp.MustCompile("false|true"),
			},
		},
		"overridden name is not camel case": {
			overrides: &FieldOverrides{
				Name: networkserver.String("not A camel-cAse Name"),
			},
			want: Field{
				Name:        "NotACamelCAseName",
				Description: "A default description",
				Type:        FieldTypeBoolean,
				JSONName:    "default",
				Validation:  regexp.MustCompile("false|true"),
			},
		},
		"description is overridden": {
			overrides: &FieldOverrides{
				Description: networkserver.String(`An updated
description field`),
			},
			want: Field{
				Name: "Default",
				Description: `An updated
description field`,
				Type:       FieldTypeBoolean,
				JSONName:   "default",
				Validation: regexp.MustCompile("false|true"),
			},
		},
		"type is overridden": {
			overrides: &FieldOverrides{
				Type: &FieldTypeString,
			},
			want: Field{
				Name:        "Default",
				Description: "A default description",
				Type:        FieldTypeString,
				JSONName:    "default",
				Validation:  regexp.MustCompile("false|true"),
			},
		},
		"json name is overridden": {
			overrides: &FieldOverrides{
				JSONName: networkserver.String("overriddenName"),
			},
			want: Field{
				Name:        "Default",
				Description: "A default description",
				Type:        FieldTypeBoolean,
				JSONName:    "overriddenName",
				Validation:  regexp.MustCompile("false|true"),
			},
		},
		"validation is overridden": {
			overrides: &FieldOverrides{
				Validation: regexp.MustCompile("[0-9]"),
			},
			want: Field{
				Name:        "Default",
				Description: "A default description",
				Type:        FieldTypeBoolean,
				JSONName:    "default",
				Validation:  regexp.MustCompile("[0-9]"),
			},
		},
		"use default description with nil overrides": {
			defaultDescription: true,
			overrides:          nil,
			want: Field{
				Name: "Default",
				Description: `Default has been auto generated from the Unifi Network Server API specification

Validation: false|true`,
				Type:       FieldTypeBoolean,
				JSONName:   "default",
				Validation: regexp.MustCompile("false|true"),
			},
		},
		"use default description with empty overrides": {
			defaultDescription: true,
			overrides:          &FieldOverrides{},
			want: Field{
				Name: "Default",
				Description: `Default has been auto generated from the Unifi Network Server API specification

Validation: false|true`,
				Type:       FieldTypeBoolean,
				JSONName:   "default",
				Validation: regexp.MustCompile("false|true"),
			},
		},
		"use default description with no description override": {
			defaultDescription: true,
			overrides: &FieldOverrides{
				Name:       networkserver.String("Updated"),
				Validation: regexp.MustCompile("[0-9]+"),
			},
			want: Field{
				Name: "Updated",
				Description: `Updated has been auto generated from the Unifi Network Server API specification

Validation: [0-9]+`,
				Type:       FieldTypeBoolean,
				JSONName:   "default",
				Validation: regexp.MustCompile("[0-9]+"),
			},
		},
		"use default description with description override": {
			defaultDescription: true,
			overrides: &FieldOverrides{
				Description: networkserver.String("This will override the default description"),
				Name:        networkserver.String("Updated"),
				Validation:  regexp.MustCompile("[0-9]+"),
			},
			want: Field{
				Name:        "Updated",
				Description: "This will override the default description",
				Type:        FieldTypeBoolean,
				JSONName:    "default",
				Validation:  regexp.MustCompile("[0-9]+"),
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := defaultField.ApplyOverrides(test.overrides, test.defaultDescription)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestField_defaultDescription(t *testing.T) {
	tests := map[string]struct {
		field Field
		want  string
	}{
		"nil validation": {
			field: Field{
				Name:       "NilValidation",
				Type:       FieldTypeString,
				Validation: nil,
			},
			want: `NilValidation has been auto generated from the Unifi Network Server API specification

Validation: None`,
		},
		"validation is empty string": {
			field: Field{
				Name:       "EmptyRegex",
				Type:       FieldTypeString,
				Validation: regexp.MustCompile(""),
			},
			want: `EmptyRegex has been auto generated from the Unifi Network Server API specification

Validation: None`,
		},
		"validation has real regex": {
			field: Field{
				Name:       "RegexValidation",
				Type:       FieldTypeNumber,
				Validation: regexp.MustCompile("[0-9]+"),
			},
			want: `RegexValidation has been auto generated from the Unifi Network Server API specification

Validation: [0-9]+`,
		},
		"validation is list of primitive": {
			field: Field{
				Name:       "RegexValidation",
				Type:       FieldTypeList(FieldTypeNumber),
				Validation: regexp.MustCompile("[0-9]+"),
			},
			want: `RegexValidation has been auto generated from the Unifi Network Server API specification

Element Validation: [0-9]+`,
		},
		"validation is list of object": {
			field: Field{
				Name:       "ListOfObject",
				Type:       FieldTypeList(FieldTypeObject("ListObject")),
				Validation: regexp.MustCompile("[0-9]+"),
			},
			want: `ListOfObject has been auto generated from the Unifi Network Server API specification

Element Validation: None`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := test.field.defaultDescription()
			assert.Equal(t, test.want, got)
		})
	}
}

func Test_parseFieldDefinition(t *testing.T) {
	tests := map[string]struct {
		name  string
		value interface{}

		wantField  Field
		wantObject *fieldObject
		wantErr    assert.ErrorAssertionFunc
	}{
		"empty string value": {
			name:  "hostname",
			value: "",
			wantField: Field{
				Name: "Hostname",
				Description: `Hostname has been auto generated from the Unifi Network Server API specification

Validation: None`,
				Type:       FieldTypeString,
				JSONName:   "hostname",
				Validation: nil,
			},
			wantErr: assert.NoError,
		},
		"value has validation for a string": {
			name:  "mac",
			value: `^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$`,
			wantField: Field{
				Name: "Mac",
				Description: `Mac has been auto generated from the Unifi Network Server API specification

Validation: ^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$`,
				Type:       FieldTypeString,
				JSONName:   "mac",
				Validation: regexp.MustCompile(`^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$`),
			},
			wantErr: assert.NoError,
		},
		"value has false|true validation": {
			name:  "use_fixedip",
			value: `false|true`,
			wantField: Field{
				Name: "UseFixedip",
				Description: `UseFixedip has been auto generated from the Unifi Network Server API specification

Validation: false|true`,
				JSONName:   "use_fixedip",
				Type:       FieldTypeBoolean,
				Validation: regexp.MustCompile(`false|true`),
			},
			wantErr: assert.NoError,
		},
		"value has true|false validation": {
			name:  "use_fixedip",
			value: `true|false`,
			wantField: Field{
				Name: "UseFixedip",
				Description: `UseFixedip has been auto generated from the Unifi Network Server API specification

Validation: true|false`,
				JSONName:   "use_fixedip",
				Type:       FieldTypeBoolean,
				Validation: regexp.MustCompile(`true|false`),
			},
			wantErr: assert.NoError,
		},
		"value has validation for integer": {
			name:  "qos_rate_max_down",
			value: `^-1|[2-9]|[1-9][0-9]{1,4}|100000$`,
			wantField: Field{
				Name: "QosRateMaxDown",
				Description: `QosRateMaxDown has been auto generated from the Unifi Network Server API specification

Validation: ^-1|[2-9]|[1-9][0-9]{1,4}|100000$`,
				JSONName:   "qos_rate_max_down",
				Type:       FieldTypeNumber,
				Validation: regexp.MustCompile(`^-1|[2-9]|[1-9][0-9]{1,4}|100000$`),
			},
			wantErr: assert.NoError,
		},
		"value has validation for float": {
			name:  "heightInMeters",
			value: `^([-]?[\d]+[.]?[\d]*)$`,
			wantField: Field{
				Name: "HeightInMeters",
				Description: `HeightInMeters has been auto generated from the Unifi Network Server API specification

Validation: ^([-]?[\d]+[.]?[\d]*)$`,
				JSONName:   "heightInMeters",
				Type:       FieldTypeDecimal,
				Validation: regexp.MustCompile(`^([-]?[\d]+[.]?[\d]*)$`),
			},
			wantErr: assert.NoError,
		},
		"value has an empty list of validations": {
			name:  "igmp_proxy_downstream_networkconf_ids",
			value: []interface{}{},
			wantField: Field{
				Name: "IgmpProxyDownstreamNetworkconfIds",
				Description: `IgmpProxyDownstreamNetworkconfIds has been auto generated from the Unifi Network Server API specification

Element Validation: None`,
				JSONName: "igmp_proxy_downstream_networkconf_ids",
				Type:     FieldTypeList(FieldTypeString),
			},
			wantErr: assert.NoError,
		},
		"value has a list with empty string validation": {
			name:  "enabled_networks",
			value: []interface{}{""},
			wantField: Field{
				Name: "EnabledNetworks",
				Description: `EnabledNetworks has been auto generated from the Unifi Network Server API specification

Element Validation: None`,
				JSONName: "enabled_networks",
				Type:     FieldTypeList(FieldTypeString),
			},
			wantErr: assert.NoError,
		},
		"value has a list with string validation": {
			name:  "options",
			value: []interface{}{`^[^"' ]+$`},
			wantField: Field{
				Name: "Options",
				Description: `Options has been auto generated from the Unifi Network Server API specification

Element Validation: ^[^"' ]+$`,
				JSONName:   "options",
				Type:       FieldTypeList(FieldTypeString),
				Validation: regexp.MustCompile(`^[^"' ]+$`),
			},
			wantErr: assert.NoError,
		},
		"value has a list with number validation": {
			name:  "auth_ids",
			value: []interface{}{"0|1|2|3|4|5|1000"},
			wantField: Field{
				Name: "AuthIds",
				Description: `AuthIds has been auto generated from the Unifi Network Server API specification

Element Validation: 0|1|2|3|4|5|1000`,
				JSONName:   "auth_ids",
				Type:       FieldTypeList(FieldTypeNumber),
				Validation: regexp.MustCompile("0|1|2|3|4|5|1000"),
			},
			wantErr: assert.NoError,
		},
		"value has a list with decimal validation": {
			name:  "a",
			value: []interface{}{`^([-]?[\d]+[.]?[\d]*)$`},
			wantField: Field{
				Name: "A",
				Description: `A has been auto generated from the Unifi Network Server API specification

Element Validation: ^([-]?[\d]+[.]?[\d]*)$`,
				JSONName:   "a",
				Type:       FieldTypeList(FieldTypeDecimal),
				Validation: regexp.MustCompile(`^([-]?[\d]+[.]?[\d]*)$`),
			},
			wantErr: assert.NoError,
		},
		"value is an object": {
			name: "config",
			value: map[string]interface{}{
				"list":   []interface{}{},
				"number": `[0-9]`,
				"string": ``,
			},
			wantField: Field{
				Name: "Config",
				Description: `Config has been auto generated from the Unifi Network Server API specification

Validation: None`,
				JSONName: "config",
				Type:     FieldTypeObject("Config"),
			},
			wantObject: &fieldObject{
				Name: "Config",
				Value: map[string]interface{}{
					"list":   []interface{}{},
					"number": `[0-9]`,
					"string": ``,
				},
			},
			wantErr: assert.NoError,
		},
		"value is a nested object": {
			name: "config",
			value: map[string]interface{}{
				"nested": map[string]interface{}{
					"number": `[0-9]`,
				},
				"string": ``,
			},
			wantField: Field{
				Name: "Config",
				Description: `Config has been auto generated from the Unifi Network Server API specification

Validation: None`,
				JSONName: "config",
				Type:     FieldTypeObject("Config"),
			},
			wantObject: &fieldObject{
				Name: "Config",
				Value: map[string]interface{}{
					"nested": map[string]interface{}{
						"number": `[0-9]`,
					},
					"string": ``,
				},
			},
			wantErr: assert.NoError,
		},
		"value has a list with object value": {
			name: "satisfaction_table",
			value: []interface{}{
				map[string]interface{}{
					"device_mac":   `^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$`,
					"satisfaction": `^[0-9]+\.?[0-9]*$`,
				},
			},
			wantField: Field{
				Name: "SatisfactionTable",
				Description: `SatisfactionTable has been auto generated from the Unifi Network Server API specification

Element Validation: None`,
				JSONName: "satisfaction_table",
				Type:     FieldTypeList(FieldTypeObject("SatisfactionTable")),
			},
			wantObject: &fieldObject{
				Name: "SatisfactionTable",
				Value: map[string]interface{}{
					"device_mac":   `^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$`,
					"satisfaction": `^[0-9]+\.?[0-9]*$`,
				},
			},
			wantErr: assert.NoError,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			gotField, gotObjects, gotErr := parseFieldDefinition(test.name, test.value)
			test.wantErr(t, gotErr)
			assert.Equal(t, test.wantField, gotField)
			assert.Equal(t, test.wantObject, gotObjects)
		})
	}
}

func Test_getValidationRegex(t *testing.T) {
	tests := map[string]struct {
		value   interface{}
		want    *regexp.Regexp
		wantErr assert.ErrorAssertionFunc
	}{
		"empty string": {
			value:   "",
			want:    nil,
			wantErr: assert.NoError,
		},
		"string": {
			value:   `.{0,128}`,
			want:    regexp.MustCompile(`.{0,128}`),
			wantErr: assert.NoError,
		},
		"empty list": {
			value:   []interface{}{},
			want:    nil,
			wantErr: assert.NoError,
		},
		"list with empty string": {
			value:   []interface{}{""},
			want:    nil,
			wantErr: assert.NoError,
		},
		"list with string": {
			value:   []interface{}{`.{0,128}`},
			want:    regexp.MustCompile(`.{0,128}`),
			wantErr: assert.NoError,
		},
		"list with object": {
			value:   []interface{}{map[string]interface{}{"foo": "bar"}},
			want:    nil,
			wantErr: assert.NoError,
		},
		"object": {
			value:   map[string]interface{}{"foo": "bar"},
			want:    nil,
			wantErr: assert.NoError,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := getValidationRegex(test.value)
			test.wantErr(t, err)
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
