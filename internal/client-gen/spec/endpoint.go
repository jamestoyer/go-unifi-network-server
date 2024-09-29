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
	"errors"
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
)

type Endpoint struct {
	Name string

	Filename     string
	Objects      []*Object
	ResourcePath string
}

func EndpointFromAPISpec(name, filename string, values map[string]interface{}) (*Endpoint, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("invalid endpoint name")
	}

	// Make the API name consistent. This is going to be camel case to broadly get it to fit the Golang standards. This
	// can be overridden later
	name = strcase.ToCamel(name)

	rootObject, fieldObjects, err := objectFromAPISpec(name, values, true)
	if err != nil {
		return nil, fmt.Errorf("invalid endpoint root object: %w", err)
	}

	allObjects := []*Object{rootObject}
	o, err := endpointSubObjects(fieldObjects)
	if err != nil {
		return nil, fmt.Errorf("invalid endpoint field objects: %w", err)
	}

	allObjects = append(allObjects, o...)

	return &Endpoint{
		Name:         name,
		Filename:     filename,
		Objects:      allObjects,
		ResourcePath: strings.ToLower(name),
	}, nil
}

func endpointSubObjects(fieldObjects []*fieldObject) ([]*Object, error) {
	var objects []*Object
	for _, object := range fieldObjects {
		o, fo, err := objectFromAPISpec(object.Name, object.Value, false)
		if err != nil {
			return nil, err
		}

		objects = append(objects, o)

		subObjects, err := endpointSubObjects(fo)
		if err != nil {
			return nil, err
		}

		objects = append(objects, subObjects...)
	}

	return objects, nil
}
