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

func TestEndpoint_Objects(t *testing.T) {
	tests := map[string]struct {
		name     string
		filename string
		spec     map[string]interface{}
		want     []*EndpointObject
	}{
		"no nested objects": {
			name:     "Simple",
			filename: "simple.json",
			spec: map[string]interface{}{
				"empty_string": "",
				"string":       "string",
				"number":       `[0-9]`,
				"decimal":      `[0-9].[0-9]+`,
				"bool":         `true|false`,
				"list":         []interface{}{},
			},
			want: []*EndpointObject{
				{
					Name: "Simple",
					Fields: []FieldDefinition{
						{
							Name:     "Bool",
							JSONName: "bool",
							Type:     Boolean,
						},
						{
							Name:     "Decimal",
							JSONName: "decimal",
							Type:     Decimal,
						},
						{
							Name:     "EmptyString",
							JSONName: "empty_string",
							Type:     String,
						},
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
		"nested objects": {
			name:     "Nested",
			filename: "nested.json",
			spec: map[string]interface{}{
				"object": map[string]interface{}{
					"decimal": `[0-9].[0-9]+`,
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
			want: []*EndpointObject{
				{
					Name: "Nested",
					Fields: []FieldDefinition{
						{
							Name:     "List",
							JSONName: "list",
							Type:     List(Object("NestedList")),
							Object: &EndpointObject{
								Name: "NestedList",
								Fields: []FieldDefinition{
									{
										Name:     "Object",
										JSONName: "object",
										Type:     Object("NestedListObject"),
										Object: &EndpointObject{
											Name: "NestedListObject",
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
						{
							Name:     "Object",
							JSONName: "object",
							Type:     Object("NestedObject"),
							Object: &EndpointObject{
								Name: "NestedObject",
								Fields: []FieldDefinition{
									{
										Name:     "Decimal",
										JSONName: "decimal",
										Type:     Decimal,
									},
								},
							},
						},
					},
				},
				{
					Name: "NestedList",
					Fields: []FieldDefinition{
						{
							Name:     "Object",
							JSONName: "object",
							Type:     Object("NestedListObject"),
							Object: &EndpointObject{
								Name: "NestedListObject",
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
				{
					Name: "NestedListObject",
					Fields: []FieldDefinition{
						{
							Name:     "Number",
							JSONName: "number",
							Type:     Number,
						},
					},
				},
				{
					Name: "NestedObject",
					Fields: []FieldDefinition{
						{
							Name:     "Decimal",
							JSONName: "decimal",
							Type:     Decimal,
						},
					},
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			endpoint, err := NewEndpoint(test.name, test.filename, test.spec)
			require.NoError(t, err)

			assert.Equal(t, test.want, endpoint.Objects())
		})
	}
}
