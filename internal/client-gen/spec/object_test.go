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

	"github.com/stretchr/testify/assert"
)

func Test_objectFromAPISpec(t *testing.T) {
	rootFields := []Field{
		{
			Name:        "ID",
			Description: "ID has been auto generated from the Unifi Network Server API specification\n\nValidation: None",
			JSONName:    "_id",
			Type:        FieldTypeString,
		},
		{
			Name:        "SiteID",
			Description: "SiteID has been auto generated from the Unifi Network Server API specification\n\nValidation: None",
			JSONName:    "site_id",
			Type:        FieldTypeString,
		},
		{
			Name:        "Hidden",
			Description: "Hidden has been auto generated from the Unifi Network Server API specification\n\nValidation: None",
			JSONName:    "attr_hidden",
			Type:        FieldTypeBoolean,
		},
		{
			Name:        "HiddenID",
			Description: "HiddenID has been auto generated from the Unifi Network Server API specification\n\nValidation: None",
			JSONName:    "attr_hidden_id",
			Type:        FieldTypeString,
		},
		{
			Name:        "NoDelete",
			Description: "NoDelete has been auto generated from the Unifi Network Server API specification\n\nValidation: None",
			JSONName:    "attr_no_delete",
			Type:        FieldTypeBoolean,
		},
		{
			Name:        "NoEdit",
			Description: "NoEdit has been auto generated from the Unifi Network Server API specification\n\nValidation: None",
			JSONName:    "attr_no_edit",
			Type:        FieldTypeBoolean,
		},
	}

	tests := map[string]struct {
		name             string
		addDefaultFields bool
		values           map[string]interface{}
		wantObject       *Object
		wantFieldObjects []*fieldObject
		wantErr          assert.ErrorAssertionFunc
	}{
		"invalid object name": {
			name:       "     ",
			values:     map[string]interface{}{},
			wantObject: nil,
			wantErr:    assert.Error,
		},
		"values nil": {
			name:   "nilValues",
			values: nil,
			wantObject: &Object{
				Name: "NilValues",
			},
			wantErr: assert.NoError,
		},
		"values empty": {
			name:   "emptyValues",
			values: map[string]interface{}{},
			wantObject: &Object{
				Name: "EmptyValues",
			},
			wantErr: assert.NoError,
		},
		"add default fields": {
			name:             "emptyValues",
			addDefaultFields: true,
			values:           map[string]interface{}{},
			wantObject: &Object{
				Name:   "EmptyValues",
				Fields: rootFields,
			},
			wantErr: assert.NoError,
		},
		"name is not camel case": {
			name:   "not A camel Case-Object",
			values: nil,
			wantObject: &Object{
				Name: "NotACamelCaseObject",
			},
			wantErr: assert.NoError,
		},
		"values are primitives and lists of primitives": {
			name:             "NoFieldObjects",
			addDefaultFields: true,
			values: map[string]interface{}{
				"empty_string": "",
				"string":       "string",
				"number":       `[0-9]`,
				"decimal":      `[0-9].[0-9]+`,
				"bool":         `true|false`,
				"emptyList":    []interface{}{},
				"list":         []interface{}{`[0-9].[0-9]+`},
			},
			wantObject: &Object{
				Name: "NoFieldObjects",
				Fields: append(rootFields,
					[]Field{
						{
							Name: "Bool",
							Description: `Bool has been auto generated from the Unifi Network Server API specification

Validation: true|false`,
							JSONName:   "bool",
							Type:       FieldTypeBoolean,
							Validation: regexp.MustCompile(`true|false`),
						},
						{
							Name: "Decimal",
							Description: `Decimal has been auto generated from the Unifi Network Server API specification

Validation: [0-9].[0-9]+`,
							JSONName:   "decimal",
							Type:       FieldTypeDecimal,
							Validation: regexp.MustCompile(`[0-9].[0-9]+`),
						},
						{
							Name: "EmptyList",
							Description: `EmptyList has been auto generated from the Unifi Network Server API specification

Element Validation: None`,
							JSONName: "emptyList",
							Type:     FieldTypeList(FieldTypeString),
						},
						{
							Name: "EmptyString",
							Description: `EmptyString has been auto generated from the Unifi Network Server API specification

Validation: None`,
							JSONName: "empty_string",
							Type:     FieldTypeString,
						},
						{
							Name: "List",
							Description: `List has been auto generated from the Unifi Network Server API specification

Element Validation: [0-9].[0-9]+`,
							JSONName:   "list",
							Type:       FieldTypeList(FieldTypeDecimal),
							Validation: regexp.MustCompile(`[0-9].[0-9]+`),
						},
						{
							Name: "Number",
							Description: `Number has been auto generated from the Unifi Network Server API specification

Validation: [0-9]`,
							JSONName:   "number",
							Type:       FieldTypeNumber,
							Validation: regexp.MustCompile(`[0-9]`),
						},
						{
							Name: "String",
							Description: `String has been auto generated from the Unifi Network Server API specification

Validation: string`,
							JSONName:   "string",
							Type:       FieldTypeString,
							Validation: regexp.MustCompile(`string`),
						},
					}...,
				),
			},
			wantErr: assert.NoError,
		},
		"values are objects and lists of objects": {
			name:             "FieldObjectsFound",
			addDefaultFields: true,
			values: map[string]interface{}{
				"object_a": map[string]interface{}{
					"value": `\w`,
				},
				"listB": []interface{}{
					map[string]interface{}{
						"decimal": `[0-9].[0-9]+`,
					},
				},
				"nested_list_object_c": []interface{}{
					map[string]interface{}{
						"nested_list": map[string]interface{}{
							"decimal": `[0-9].[0-9]+`,
						},
					},
				},
				"nested_object_list_object_d": map[string]interface{}{
					"topObject": map[string]interface{}{
						"nestedList": []interface{}{
							map[string]interface{}{
								"decimal": `[0-9].[0-9]+`,
							},
						},
					},
				},
			},
			wantObject: &Object{
				Name: "FieldObjectsFound",
				Fields: append(
					rootFields,
					[]Field{
						{
							Name: "ListB",
							Description: `ListB has been auto generated from the Unifi Network Server API specification

Element Validation: None`,
							Type:     FieldTypeList(FieldTypeObject("FieldObjectsFoundListB")),
							JSONName: "listB",
						},
						{
							Name: "NestedListObjectC",
							Description: `NestedListObjectC has been auto generated from the Unifi Network Server API specification

Element Validation: None`,
							Type:     FieldTypeList(FieldTypeObject("FieldObjectsFoundNestedListObjectC")),
							JSONName: "nested_list_object_c",
						},
						{
							Name: "NestedObjectListObjectD",
							Description: `NestedObjectListObjectD has been auto generated from the Unifi Network Server API specification

Validation: None`,
							Type:     FieldTypeObject("FieldObjectsFoundNestedObjectListObjectD"),
							JSONName: "nested_object_list_object_d",
						},
						{
							Name: "ObjectA",
							Description: `ObjectA has been auto generated from the Unifi Network Server API specification

Validation: None`,
							Type:     FieldTypeObject("FieldObjectsFoundObjectA"),
							JSONName: "object_a",
						},
					}...,
				),
			},
			wantFieldObjects: []*fieldObject{
				{
					Name: "FieldObjectsFoundListB",
					Value: map[string]interface{}{
						"decimal": `[0-9].[0-9]+`,
					},
				},
				{
					Name: "FieldObjectsFoundNestedListObjectC",
					Value: map[string]interface{}{
						"nested_list": map[string]interface{}{
							"decimal": `[0-9].[0-9]+`,
						},
					},
				},
				{
					Name: "FieldObjectsFoundNestedObjectListObjectD",
					Value: map[string]interface{}{
						"topObject": map[string]interface{}{
							"nestedList": []interface{}{
								map[string]interface{}{
									"decimal": `[0-9].[0-9]+`,
								},
							},
						},
					},
				},
				{
					Name: "FieldObjectsFoundObjectA",
					Value: map[string]interface{}{
						"value": `\w`,
					},
				},
			},
			wantErr: assert.NoError,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			gotObject, gotFieldObjects, err := objectFromAPISpec(test.name, test.values, test.addDefaultFields)
			test.wantErr(t, err)
			assert.Equal(t, test.wantObject, gotObject)
			assert.Equal(t, test.wantFieldObjects, gotFieldObjects)
		})
	}
}
