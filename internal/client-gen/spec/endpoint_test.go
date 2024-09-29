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

func TestEndpointFromAPISpec(t *testing.T) {
	tests := map[string]struct {
		name     string
		filename string
		values   map[string]interface{}
		want     *Endpoint
		wantErr  assert.ErrorAssertionFunc
	}{
		"invalid endpoint name": {
			name:     "   ",
			filename: "invalid.json",
			values:   nil,
			want:     nil,
			wantErr:  assert.Error,
		},
		"values nil": {
			name:     "EmptyEndpoint",
			filename: "emptyEndpoint.json",
			values:   nil,
			want: &Endpoint{
				Name:         "EmptyEndpoint",
				Filename:     "emptyEndpoint.json",
				ResourcePath: "emptyendpoint",
				Objects: []*Object{
					{
						Name:   "EmptyEndpoint",
						Fields: rootObjectFields,
					},
				},
			},
			wantErr: assert.NoError,
		},
		"values empty": {
			name:     "EmptyEndpoint",
			filename: "emptyEndpoint.json",
			values:   map[string]interface{}{},
			want: &Endpoint{
				Name:         "EmptyEndpoint",
				Filename:     "emptyEndpoint.json",
				ResourcePath: "emptyendpoint",
				Objects: []*Object{
					{
						Name:   "EmptyEndpoint",
						Fields: rootObjectFields,
					},
				},
			},
			wantErr: assert.NoError,
		},
		"name is not camel case": {
			name:     "not camel Case-endpoint",
			filename: "endpoint.json",
			values:   nil,
			want: &Endpoint{
				Name:         "NotCamelCaseEndpoint",
				Filename:     "endpoint.json",
				ResourcePath: "notcamelcaseendpoint",
				Objects: []*Object{
					{
						Name:   "NotCamelCaseEndpoint",
						Fields: rootObjectFields,
					},
				},
			},
			wantErr: assert.NoError,
		},
		"endpoint has one object": {
			name:     "SingleObject",
			filename: "singleObject.json",
			values: map[string]interface{}{
				"primitive":     `[0-9].[0-9]+`,
				"listPrimitive": []interface{}{},
			},
			want: &Endpoint{
				Name:         "SingleObject",
				Filename:     "singleObject.json",
				ResourcePath: "singleobject",
				Objects: []*Object{
					{
						Name: "SingleObject",
						Fields: append(rootObjectFields,
							[]Field{
								{
									Name: "ListPrimitive",
									Description: `ListPrimitive has been auto generated from the Unifi Network Server API specification

Element Validation: None`,
									JSONName: "listPrimitive",
									Type:     FieldTypeList(FieldTypeString),
								},
								{
									Name: "Primitive",
									Description: `Primitive has been auto generated from the Unifi Network Server API specification

Validation: [0-9].[0-9]+`,
									JSONName:   "primitive",
									Type:       FieldTypeDecimal,
									Validation: regexp.MustCompile(`[0-9].[0-9]+`),
								},
							}...,
						),
					},
				},
			},
			wantErr: assert.NoError,
		},
		"endpoint has multiple objects": {
			name:     "MultipleObjects",
			filename: "multipleObjects.json",
			values: map[string]interface{}{
				"primitive": `[0-9].[0-9]+`,
				"object": map[string]interface{}{
					"decimal":       `[0-9].[0-9]+`,
					"listPrimitive": []interface{}{},
				},
				"list": []interface{}{
					map[string]interface{}{
						"string": `\w+`,
						"object": map[string]interface{}{
							"number": `[0-9]`,
						},
					},
				},
			},
			want: &Endpoint{
				Name:         "MultipleObjects",
				Filename:     "multipleObjects.json",
				ResourcePath: "multipleobjects",
				Objects: []*Object{
					{
						Name: "MultipleObjects",
						Fields: append(rootObjectFields,
							[]Field{
								{
									Name: "List",
									Description: `List has been auto generated from the Unifi Network Server API specification

Element Validation: None`,
									JSONName: "list",
									Type:     FieldTypeList(FieldTypeObject("MultipleObjectsList")),
								},
								{
									Name: "Object",
									Description: `Object has been auto generated from the Unifi Network Server API specification

Validation: None`,
									JSONName: "object",
									Type:     FieldTypeObject("MultipleObjectsObject"),
								},
								{
									Name: "Primitive",
									Description: `Primitive has been auto generated from the Unifi Network Server API specification

Validation: [0-9].[0-9]+`,
									JSONName:   "primitive",
									Type:       FieldTypeDecimal,
									Validation: regexp.MustCompile(`[0-9].[0-9]+`),
								},
							}...,
						),
					},
					{
						Name: "MultipleObjectsList",
						Fields: []Field{
							{
								Name: "Object",
								Description: `Object has been auto generated from the Unifi Network Server API specification

Validation: None`,
								JSONName: "object",
								Type:     FieldTypeObject("MultipleObjectsListObject"),
							},
							{
								Name: "String",
								Description: `String has been auto generated from the Unifi Network Server API specification

Validation: \w+`,
								JSONName:   "string",
								Type:       FieldTypeString,
								Validation: regexp.MustCompile(`\w+`),
							},
						},
					},
					{
						Name: "MultipleObjectsListObject",
						Fields: []Field{
							{
								Name: "Number",
								Description: `Number has been auto generated from the Unifi Network Server API specification

Validation: [0-9]`,
								JSONName:   "number",
								Type:       FieldTypeNumber,
								Validation: regexp.MustCompile(`[0-9]`),
							},
						},
					},
					{
						Name: "MultipleObjectsObject",
						Fields: []Field{
							{
								Name: "Decimal",
								Description: `Decimal has been auto generated from the Unifi Network Server API specification

Validation: [0-9].[0-9]+`,
								JSONName:   "decimal",
								Type:       FieldTypeDecimal,
								Validation: regexp.MustCompile(`[0-9].[0-9]+`),
							},
							{
								Name: "ListPrimitive",
								Description: `ListPrimitive has been auto generated from the Unifi Network Server API specification

Element Validation: None`,
								JSONName: "listPrimitive",
								Type:     FieldTypeList(FieldTypeString),
							},
						},
					},
				},
			},
			wantErr: assert.NoError,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := EndpointFromAPISpec(test.name, test.filename, test.values)
			test.wantErr(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func Test_endpointSubObjects(t *testing.T) {
	tests := map[string]struct {
		fieldObjects []*fieldObject
		want         []*Object
		wantErr      assert.ErrorAssertionFunc
	}{
		"object with no field objects": {
			fieldObjects: []*fieldObject{},
			want:         nil,
			wantErr:      assert.NoError,
		},
		"object with no nested objects on field objects": {
			fieldObjects: []*fieldObject{
				{
					Name:  "A",
					Value: map[string]interface{}{"string": `string`},
				},
				{
					Name:  "B",
					Value: map[string]interface{}{"number": `[0-9]`},
				},
			},
			want: []*Object{
				{
					Name: "A",
					Fields: []Field{
						{
							Name: "String",
							Description: `String has been auto generated from the Unifi Network Server API specification

Validation: string`,
							JSONName:   "string",
							Type:       FieldTypeString,
							Validation: regexp.MustCompile(`string`),
						},
					},
				},
				{
					Name: "B",
					Fields: []Field{
						{
							Name: "Number",
							Description: `Number has been auto generated from the Unifi Network Server API specification

Validation: [0-9]`,
							JSONName:   "number",
							Type:       FieldTypeNumber,
							Validation: regexp.MustCompile(`[0-9]`),
						},
					},
				},
			},
			wantErr: assert.NoError,
		},
		"object with nested field objects": {
			fieldObjects: []*fieldObject{
				{
					Name: "rootField",
					Value: map[string]interface{}{
						"subField": map[string]interface{}{
							"string": "string",
						},
					},
				},
			},
			want: []*Object{
				{
					Name: "RootField",
					Fields: []Field{
						{
							Name: "SubField",
							Description: `SubField has been auto generated from the Unifi Network Server API specification

Validation: None`,
							JSONName: "subField",
							Type:     FieldTypeObject("RootFieldSubField"),
						},
					},
				},
				{
					Name: "RootFieldSubField",
					Fields: []Field{
						{
							Name: "String",
							Description: `String has been auto generated from the Unifi Network Server API specification

Validation: string`,
							JSONName:   "string",
							Type:       FieldTypeString,
							Validation: regexp.MustCompile(`string`),
						},
					},
				},
			},
			wantErr: assert.NoError,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := endpointSubObjects(test.fieldObjects)
			test.wantErr(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}
