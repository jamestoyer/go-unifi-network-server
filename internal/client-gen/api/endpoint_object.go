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
	"errors"
	"fmt"
	"maps"
	"slices"
)

type EndpointObject struct {
	Name   string
	Fields []FieldDefinition
	Root   bool
}

func NewEndpointObject(name string, values map[string]interface{}, namePrefix string, rootObject bool) (*EndpointObject, error) {
	name = namePrefix + name
	var fields []FieldDefinition
	var errs []error

	sortedFieldNames := slices.Sorted(maps.Keys(values))
	for _, fieldName := range sortedFieldNames {
		value := values[fieldName]
		definition, err := NewFieldDefinition(fieldName, name, value)
		if err != nil {
			errs = append(errs, err)
		}

		fields = append(fields, definition)
	}

	if len(errs) > 0 {
		return nil, fmt.Errorf("invalid field definitions for endpoint object: %w", errors.Join(errs...))
	}

	return &EndpointObject{
		Name:   name,
		Fields: fields,
		Root:   rootObject,
	}, nil
}

// NestedObjects will loop over all Fields to find any that have an EndpointObject. It will then recursively check the
// fields of any found EndpointObject to also get these and will return them all in a flattened list.
func (e *EndpointObject) NestedObjects() []*EndpointObject {
	var nestedObjects []*EndpointObject
	for _, field := range e.Fields {
		if field.Object != nil {
			nestedObjects = append(nestedObjects, field.Object)
			nestedObjects = append(nestedObjects, field.Object.NestedObjects()...)
		}
	}

	return nestedObjects
}
