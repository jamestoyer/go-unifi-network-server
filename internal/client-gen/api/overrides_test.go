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
	"regexp"
	"testing"

	"github.com/jamestoyer/go-unifi-network-server/internal/client-gen/spec"
	"github.com/jamestoyer/go-unifi-network-server/networkserver"
	"github.com/stretchr/testify/assert"
)

func TestApplyEndpointOverrides(t *testing.T) {
	defaultEndpoint := &spec.Endpoint{
		Name:         "Default",
		ResourcePath: "default-path",
		Objects: []*spec.Object{
			{
				Name:   "A",
				Fields: []spec.Field{},
			},
			{
				Name:   "B",
				Fields: []spec.Field{},
			},
			{
				Name: "C",
				Fields: []spec.Field{
					{
						Name: "Default",
					},
				},
			},
			{
				Name:   "D",
				Fields: []spec.Field{},
			},
		},
	}

	tests := map[string]struct {
		endpoint  *spec.Endpoint
		overrides *EndpointOverrides
		want      *spec.Endpoint
	}{
		"endpoint is nil": {
			endpoint: nil,
			want:     nil,
		},
		"overrides is nil": {
			endpoint:  defaultEndpoint,
			overrides: nil,
			want:      defaultEndpoint,
		},
		"name is overridden": {
			endpoint: defaultEndpoint,
			overrides: &EndpointOverrides{
				Name: networkserver.String("Updated"),
			},
			want: &spec.Endpoint{
				Name:         "Updated",
				ResourcePath: "default-path",
				Objects: []*spec.Object{
					{
						Name:   "A",
						Fields: []spec.Field{},
					},
					{
						Name:   "B",
						Fields: []spec.Field{},
					},
					{
						Name: "C",
						Fields: []spec.Field{
							{
								Name: "Default",
							},
						},
					},
					{
						Name:   "D",
						Fields: []spec.Field{},
					},
				},
			},
		},
		"resource path is overridden": {
			endpoint: defaultEndpoint,
			overrides: &EndpointOverrides{
				ResourcePath: networkserver.String("new-path"),
			},
			want: &spec.Endpoint{
				Name:         "Default",
				ResourcePath: "new-path",
				Objects: []*spec.Object{
					{
						Name:   "A",
						Fields: []spec.Field{},
					},
					{
						Name:   "B",
						Fields: []spec.Field{},
					},
					{
						Name: "C",
						Fields: []spec.Field{
							{
								Name: "Default",
							},
						},
					},
					{
						Name:   "D",
						Fields: []spec.Field{},
					},
				},
			},
		},
		"objects are overridden": {
			endpoint: defaultEndpoint,
			overrides: &EndpointOverrides{
				Objects: map[string]*ObjectOverrides{
					"B": nil,
					"C": {
						Name: networkserver.String("Updated"),
						Fields: map[string]*FieldOverrides{
							"Default": {
								GenerateDefaultDescription: networkserver.Bool(false),
								JSONName:                   networkserver.String("aString"),
							},
						},
					},
					"D": {},
				},
			},
			want: &spec.Endpoint{
				Name:         "Default",
				ResourcePath: "default-path",
				Objects: []*spec.Object{
					{
						Name:   "A",
						Fields: []spec.Field{},
					},
					{
						Name:   "B",
						Fields: []spec.Field{},
					},
					{
						Name: "Updated",
						Fields: []spec.Field{
							{
								Name:     "Default",
								JSONName: "aString",
							},
						},
					},
					{
						Name:   "D",
						Fields: []spec.Field{},
					},
				},
			},
		},
		"actions are empty": {
			endpoint: &spec.Endpoint{
				Name:         "ActionsOverride",
				ResourcePath: "actions",
			},
			overrides: &EndpointOverrides{
				Actions: &EndpointActionsOverrides{},
			},
			want: &spec.Endpoint{
				Name:         "ActionsOverride",
				ResourcePath: "actions",
			},
		},
		"actions are overridden": {
			endpoint: &spec.Endpoint{
				Name:         "ActionsOverride",
				ResourcePath: "actions",
			},
			overrides: &EndpointOverrides{
				Name: networkserver.String("Updated"),
				Actions: &EndpointActionsOverrides{
					DisableCreate: networkserver.Bool(true),
					DisableDelete: networkserver.Bool(true),
					DisableGet:    networkserver.Bool(true),
					DisableList:   networkserver.Bool(true),
					DisableUpdate: networkserver.Bool(true),
				},
			},
			want: &spec.Endpoint{
				Name:         "Updated",
				ResourcePath: "actions",
				Actions: spec.EndpointActions{
					DisableCreate: true,
					DisableDelete: true,
					DisableGet:    true,
					DisableList:   true,
					DisableUpdate: true,
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := ApplyEndpointOverrides(test.endpoint, test.overrides)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestApplyObjectOverrides(t *testing.T) {
	defaultObject := &spec.Object{
		Name: "Default",
		Fields: []spec.Field{
			{
				Name: "A",
				Type: spec.FieldTypeBoolean,
			},
			{
				Name: "B",
				Type: spec.FieldTypeDecimal,
			},
			{
				Name: "C",
				Type: spec.FieldTypeNumber,
			},
			{
				Name: "D",
				Type: spec.FieldTypeString,
			},
		},
	}

	tests := map[string]struct {
		object    *spec.Object
		overrides *ObjectOverrides
		want      *spec.Object
	}{
		"object is nil": {
			object: nil,
			want:   nil,
		},
		"overrides is nil": {
			object:    defaultObject,
			overrides: nil,
			want:      defaultObject,
		},
		"overrides contains no overrides": {
			object:    defaultObject,
			overrides: &ObjectOverrides{},
			want:      defaultObject,
		},
		"name is overridden": {
			object: defaultObject,
			overrides: &ObjectOverrides{
				Name: networkserver.String("Updated"),
			},
			want: &spec.Object{
				Name: "Updated",
				Fields: []spec.Field{
					{
						Name: "A",
						Type: spec.FieldTypeBoolean,
					},
					{
						Name: "B",
						Type: spec.FieldTypeDecimal,
					},
					{
						Name: "C",
						Type: spec.FieldTypeNumber,
					},
					{
						Name: "D",
						Type: spec.FieldTypeString,
					},
				},
			},
		},
		"fields are overridden": {
			object: defaultObject,
			overrides: &ObjectOverrides{
				Fields: map[string]*FieldOverrides{
					"B": nil,
					"C": {
						Name:        networkserver.String("Updated"),
						Description: networkserver.String("Updated description"),
					},
					"D": {},
				},
			},
			want: &spec.Object{
				Name: "Default",
				Fields: []spec.Field{
					{
						Name: "A",
						Type: spec.FieldTypeBoolean,
					},
					{
						Name: "B",
						Type: spec.FieldTypeDecimal,
					},
					{
						Name:        "Updated",
						Description: "Updated description",
						Type:        spec.FieldTypeNumber,
					},
					{
						Description: `D has been auto generated from the Unifi Network Server API specification

Validation: None`,
						Name: "D",
						Type: spec.FieldTypeString,
					},
				},
			},
		},
		"additional fields": {
			object: defaultObject,
			overrides: &ObjectOverrides{
				Fields: map[string]*FieldOverrides{
					"C": {
						Name:                       networkserver.String("Updated"),
						GenerateDefaultDescription: networkserver.Bool(false),
					},
				},
				AdditionalFields: map[string]*FieldOverrides{
					"AdditionalC": {
						Name:                       networkserver.String("NewField1"),
						GenerateDefaultDescription: networkserver.Bool(false),
					},
					"AdditionalA": {
						Description: networkserver.String("A custom description"),
						JSONName:    networkserver.String("custom_json"),
						Name:        networkserver.String("OverriddenName"),
						Type:        &spec.FieldTypeDecimal,
						Validation:  regexp.MustCompile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`),
					},
					"AdditionalB": {},
				},
			},
			want: &spec.Object{
				Name: "Default",
				Fields: []spec.Field{
					{
						Name: "A",
						Type: spec.FieldTypeBoolean,
					},
					{
						Name: "B",
						Type: spec.FieldTypeDecimal,
					},
					{
						Name: "Updated",
						Type: spec.FieldTypeNumber,
					},
					{
						Name: "D",
						Type: spec.FieldTypeString,
					},
					{
						Name:        "OverriddenName",
						Description: "A custom description",
						JSONName:    "custom_json",
						Type:        spec.FieldTypeDecimal,
						Validation:  regexp.MustCompile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`),
					},
					{
						Name: "AdditionalB",
						Description: `AdditionalB has been auto generated from the Unifi Network Server API specification

Validation: None`,
						JSONName: "additional_b",
						Type:     spec.FieldTypeString,
					},
					{
						Name:     "NewField1",
						JSONName: "new_field_1",
						Type:     spec.FieldTypeString,
					},
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := ApplyObjectOverrides(test.object, test.overrides)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestField_ApplyFieldOverrides(t *testing.T) {
	defaultField := spec.Field{
		Name:        "Default",
		Description: "A default description",
		Type:        spec.FieldTypeBoolean,
		JSONName:    "default",
		Validation:  regexp.MustCompile("false|true"),
	}

	tests := map[string]struct {
		field     spec.Field
		overrides *FieldOverrides
		want      spec.Field
	}{
		"overrides is nil": {
			field:     defaultField,
			overrides: nil,
			want: spec.Field{
				Name:        "Default",
				Description: "A default description",
				Type:        spec.FieldTypeBoolean,
				JSONName:    "default",
				Validation:  regexp.MustCompile("false|true"),
			},
		},
		"overrides contains no overrides": {
			field:     defaultField,
			overrides: &FieldOverrides{},
			want: spec.Field{
				Name:        "Default",
				Description: "A default description",
				Type:        spec.FieldTypeBoolean,
				JSONName:    "default",
				Validation:  regexp.MustCompile("false|true"),
			},
		},
		"name is overridden": {
			field: defaultField,
			overrides: &FieldOverrides{
				Name: networkserver.String("OverriddenName"),
			},
			want: spec.Field{
				Name:        "OverriddenName",
				Description: "A default description",
				Type:        spec.FieldTypeBoolean,
				JSONName:    "default",
				Validation:  regexp.MustCompile("false|true"),
			},
		},
		"description is overridden": {
			field: defaultField,
			overrides: &FieldOverrides{
				Description: networkserver.String(`An updated
description field`),
			},
			want: spec.Field{
				Name: "Default",
				Description: `An updated
description field`,
				Type:       spec.FieldTypeBoolean,
				JSONName:   "default",
				Validation: regexp.MustCompile("false|true"),
			},
		},
		"type is overridden": {
			field: defaultField,
			overrides: &FieldOverrides{
				Type: &spec.FieldTypeString,
			},
			want: spec.Field{
				Name:        "Default",
				Description: "A default description",
				Type:        spec.FieldTypeString,
				JSONName:    "default",
				Validation:  regexp.MustCompile("false|true"),
			},
		},
		"json name is overridden": {
			field: defaultField,
			overrides: &FieldOverrides{
				JSONName: networkserver.String("overriddenName"),
			},
			want: spec.Field{
				Name:        "Default",
				Description: "A default description",
				Type:        spec.FieldTypeBoolean,
				JSONName:    "overriddenName",
				Validation:  regexp.MustCompile("false|true"),
			},
		},
		"validation is overridden": {
			field: defaultField,
			overrides: &FieldOverrides{
				Validation: regexp.MustCompile("[0-9]"),
			},
			want: spec.Field{
				Name:        "Default",
				Description: "A default description",
				Type:        spec.FieldTypeBoolean,
				JSONName:    "default",
				Validation:  regexp.MustCompile("[0-9]"),
			},
		},
		"generate default description is nil with default description": {
			field: spec.Field{
				Name: "Default",
				Description: `Default has been auto generated from the Unifi Network Server API specification

Validation: false|true`,
				Type:       spec.FieldTypeBoolean,
				JSONName:   "default",
				Validation: regexp.MustCompile("false|true"),
			},
			overrides: &FieldOverrides{
				Name:       networkserver.String("Updated"),
				Validation: regexp.MustCompile("[0-9]+"),
			},
			want: spec.Field{
				Name: "Updated",
				Description: `Updated has been auto generated from the Unifi Network Server API specification

Validation: [0-9]+`,
				Type:       spec.FieldTypeBoolean,
				JSONName:   "default",
				Validation: regexp.MustCompile("[0-9]+"),
			},
		},
		"generate default description is true with default description": {
			field: spec.Field{
				Name: "Default",
				Description: `Default has been auto generated from the Unifi Network Server API specification

Validation: false|true`,
				Type:       spec.FieldTypeBoolean,
				JSONName:   "default",
				Validation: regexp.MustCompile("false|true"),
			},
			overrides: &FieldOverrides{
				GenerateDefaultDescription: networkserver.Bool(true),
				Name:                       networkserver.String("Updated"),
				Validation:                 regexp.MustCompile("[0-9]+"),
			},
			want: spec.Field{
				Name: "Updated",
				Description: `Updated has been auto generated from the Unifi Network Server API specification

Validation: [0-9]+`,
				Type:       spec.FieldTypeBoolean,
				JSONName:   "default",
				Validation: regexp.MustCompile("[0-9]+"),
			},
		},
		"generate default description is false with default description": {
			field: spec.Field{
				Name: "Default",
				Description: `Default has been auto generated from the Unifi Network Server API specification

Validation: false|true`,
				Type:       spec.FieldTypeBoolean,
				JSONName:   "default",
				Validation: regexp.MustCompile("false|true"),
			},
			overrides: &FieldOverrides{
				GenerateDefaultDescription: networkserver.Bool(false),
				Name:                       networkserver.String("Updated"),
				Validation:                 regexp.MustCompile("[0-9]+"),
			},
			want: spec.Field{
				Name:        "Updated",
				Description: "",
				Type:        spec.FieldTypeBoolean,
				JSONName:    "default",
				Validation:  regexp.MustCompile("[0-9]+"),
			},
		},
		"generate default description is false with non default description": {
			field: defaultField,
			overrides: &FieldOverrides{
				GenerateDefaultDescription: networkserver.Bool(false),
				Name:                       networkserver.String("Updated"),
				Validation:                 regexp.MustCompile("[0-9]+"),
			},
			want: spec.Field{
				Name:        "Updated",
				Description: "A default description",
				Type:        spec.FieldTypeBoolean,
				JSONName:    "default",
				Validation:  regexp.MustCompile("[0-9]+"),
			},
		},
		"generate default description with empty string override": {
			field: defaultField,
			overrides: &FieldOverrides{
				Description: networkserver.String(""),
				Name:        networkserver.String("Updated"),
				Validation:  regexp.MustCompile("[0-9]+"),
			},
			want: spec.Field{
				Name:        "Updated",
				Description: "",
				Type:        spec.FieldTypeBoolean,
				JSONName:    "default",
				Validation:  regexp.MustCompile("[0-9]+"),
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := ApplyFieldOverrides(test.field, test.overrides)
			assert.Equal(t, test.want, got)
		})
	}
}
