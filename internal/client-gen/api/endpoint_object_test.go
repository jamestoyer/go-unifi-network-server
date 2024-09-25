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
)

func TestNewEndpointObject(t *testing.T) {
	tests := map[string]struct {
		name       string
		namePrefix string
		values     map[string]interface{}

		error bool
		want  *EndpointObject
	}{
		"no values": {
			name: "Empty",
			want: &EndpointObject{Name: "Empty"},
		},
		"name prefix": {
			name:       "Empty",
			namePrefix: "Prefixed",
			want:       &EndpointObject{Name: "PrefixedEmpty"},
		},
		"primitive values": {
			name: "Primitive",
			values: map[string]interface{}{
				"empty_string": "",
				"string":       "string",
				"number":       `[0-9]`,
				"decimal":      `[0-9].[0-9]+`,
				"bool":         `true|false`,
				"list":         []interface{}{},
			},
			want: &EndpointObject{
				Name: "Primitive",
				Fields: []FieldDefinition{
					{
						Name:     "EmptyString",
						JSONName: "empty_string",
						Type:     String,
					},
					{
						Name:     "String",
						JSONName: "string",
						Type:     String,
					},
					{
						Name:     "Number",
						JSONName: "number",
						Type:     Number,
					},
					{
						Name:     "Decimal",
						JSONName: "decimal",
						Type:     Decimal,
					},
					{
						Name:     "Bool",
						JSONName: "bool",
						Type:     Boolean,
					},
					{
						Name:     "List",
						JSONName: "list",
						Type:     List(String),
					},
				},
			},
		},
		"object values": {
			name: "EntrypointObject",
			values: map[string]interface{}{
				"string": "string",
				"nested_object": map[string]interface{}{
					"empty_string": "",
					"number":       `[0-9]`,
				},
			},
			want: &EndpointObject{
				Name: "EntrypointObject",
				Fields: []FieldDefinition{
					{
						Name:     "String",
						JSONName: "string",
						Type:     String,
					},
					{
						Name:     "NestedObject",
						JSONName: "nested_object",
						Type:     Object("EntrypointObjectNestedObject"),
					},
				},
				NestedObjects: []*EndpointObject{
					{
						Name: "EntrypointObjectNestedObject",
						Fields: []FieldDefinition{
							{
								Name:     "EmptyString",
								JSONName: "empty_string",
								Type:     String,
							},
							{
								Name:     "Number",
								JSONName: "number",
								Type:     Number,
							},
						},
					},
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := NewEndpointObject(test.name, test.values, test.namePrefix)
			if test.error {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.want.Name, got.Name)
			assert.ElementsMatch(t, test.want.Fields, got.Fields)
			assert.ElementsMatch(t, test.want.NestedObjects, got.NestedObjects)
		})
	}
}
