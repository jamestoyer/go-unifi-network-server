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

func TestField_ApplyFieldOverrides(t *testing.T) {
	defaultField := spec.Field{
		Name:        "Default",
		Description: "A default description",
		Type:        spec.FieldTypeBoolean,
		JSONName:    "default",
		Validation:  regexp.MustCompile("false|true"),
	}

	tests := map[string]struct {
		defaultDescription bool
		overrides          *FieldOverrides
		want               spec.Field
	}{
		"overrides is nil": {
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
		"overridden name is not camel case": {
			overrides: &FieldOverrides{
				Name: networkserver.String("not A camel-cAse Name"),
			},
			want: spec.Field{
				Name:        "NotACamelCAseName",
				Description: "A default description",
				Type:        spec.FieldTypeBoolean,
				JSONName:    "default",
				Validation:  regexp.MustCompile("false|true"),
			},
		},
		"description is overridden": {
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
		"use default description with nil overrides": {
			defaultDescription: true,
			overrides:          nil,
			want: spec.Field{
				Name: "Default",
				Description: `Default has been auto generated from the Unifi Network Server API specification

Validation: false|true`,
				Type:       spec.FieldTypeBoolean,
				JSONName:   "default",
				Validation: regexp.MustCompile("false|true"),
			},
		},
		"use default description with empty overrides": {
			defaultDescription: true,
			overrides:          &FieldOverrides{},
			want: spec.Field{
				Name: "Default",
				Description: `Default has been auto generated from the Unifi Network Server API specification

Validation: false|true`,
				Type:       spec.FieldTypeBoolean,
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
			want: spec.Field{
				Name: "Updated",
				Description: `Updated has been auto generated from the Unifi Network Server API specification

Validation: [0-9]+`,
				Type:       spec.FieldTypeBoolean,
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
			want: spec.Field{
				Name:        "Updated",
				Description: "This will override the default description",
				Type:        spec.FieldTypeBoolean,
				JSONName:    "default",
				Validation:  regexp.MustCompile("[0-9]+"),
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := ApplyFieldOverrides(defaultField, test.overrides, test.defaultDescription)
			assert.Equal(t, test.want, got)
		})
	}
}
