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
)

type EndpointObject struct {
	Name   string
	Fields []FieldDefinition
}

func NewEndpointObject(name string, values map[string]interface{}, namePrefix string) (*EndpointObject, error) {
	var fields []FieldDefinition
	var errs []error
	for fieldName, value := range values {
		definition, err := NewFieldDefinition(fieldName, value)
		if err != nil {
			errs = append(errs, err)
		}

		fields = append(fields, definition)
	}

	if len(errs) > 0 {
		return nil, fmt.Errorf("invalid field definitions for endpoint object: %w", errors.Join(errs...))
	}

	return &EndpointObject{
		Name:   namePrefix + name,
		Fields: fields,
	}, nil
}
