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
	"maps"
	"regexp"
	"slices"

	"github.com/iancoleman/strcase"
	"github.com/jamestoyer/go-unifi-network-server/internal/client-gen/spec"
)

type Overrides struct {
	Endpoints map[string]*EndpointOverrides `yaml:"endpoints"`
}

func ApplyOverrides(endpoints []*spec.Endpoint, overrides Overrides) []*spec.Endpoint {
	var newEndpoints []*spec.Endpoint
	for _, ep := range endpoints {
		override := overrides.Endpoints[ep.Name]
		updated := ApplyEndpointOverrides(ep, override)
		newEndpoints = append(newEndpoints, updated)
	}

	return newEndpoints
}

type EndpointOverrides struct {
	Name         *string                     `yaml:"name,omitempty"`
	Objects      map[string]*ObjectOverrides `yaml:"objects,omitempty"`
	ResourcePath *string                     `yaml:"resourcePath,omitempty"`
}

// ApplyEndpointOverrides will return a mutated spec.Endpoint with the given EndpointOverrides applied.
func ApplyEndpointOverrides(endpoint *spec.Endpoint, overrides *EndpointOverrides) *spec.Endpoint {
	if endpoint == nil {
		return nil
	}

	// Dereference the endpoint to ensure we don't make changes to the given endpoint pointer
	ep := *endpoint
	if overrides == nil {
		return &ep
	}

	if overrides.Name != nil {
		ep.Name = *overrides.Name
	}

	if overrides.ResourcePath != nil {
		ep.ResourcePath = *overrides.ResourcePath
	}

	var objects []*spec.Object
	for _, object := range endpoint.Objects {
		override := overrides.Objects[object.Name]
		updated := ApplyObjectOverrides(object, override)
		objects = append(objects, updated)
	}

	ep.Objects = objects
	return &ep
}

type ObjectOverrides struct {
	Name             *string                    `yaml:"name,omitempty"`
	AdditionalFields map[string]*FieldOverrides `yaml:"additionalFields,omitempty"`
	Fields           map[string]*FieldOverrides `yaml:"fields,omitempty"`
}

func ApplyObjectOverrides(object *spec.Object, overrides *ObjectOverrides) *spec.Object {
	if object == nil {
		return nil
	}

	// Dereference the spec.Object to ensure we don't make changes to the given spec.Object pointer
	o := *object
	if overrides == nil {
		return &o
	}

	if overrides.Name != nil {
		o.Name = *overrides.Name
	}

	if overrides.Fields != nil {
		var overriddenFields []spec.Field
		for _, field := range o.Fields {
			fieldOverrides, ok := overrides.Fields[field.Name]
			if !ok {
				overriddenFields = append(overriddenFields, field)
				continue
			}

			f := ApplyFieldOverrides(field, fieldOverrides)
			overriddenFields = append(overriddenFields, f)
		}

		o.Fields = overriddenFields
	}

	if overrides.AdditionalFields != nil {
		var additionalFields []spec.Field
		sortedNames := slices.Sorted(maps.Keys(overrides.AdditionalFields))
		for _, name := range sortedNames {
			additionalField := overrides.AdditionalFields[name]
			additionalFields = append(additionalFields, populateAdditionalField(name, additionalField))
		}

		o.Fields = append(o.Fields, additionalFields...)
	}

	return &o
}

func populateAdditionalField(name string, overrides *FieldOverrides) spec.Field {
	field := spec.Field{
		Name:        name,
		Description: "",
		Type:        spec.FieldType{},
		JSONName:    "",
		Validation:  nil,
	}

	if overrides.Name != nil {
		field.Name = *overrides.Name
	}

	if overrides.Type != nil {
		field.Type = *overrides.Type
	} else {
		field.Type = spec.FieldTypeString
	}

	if overrides.JSONName != nil {
		field.JSONName = *overrides.JSONName
	} else {
		field.JSONName = strcase.ToSnake(field.Name)
	}

	if overrides.Validation != nil {
		field.Validation = overrides.Validation
	}

	generateDefaultDescription := true
	if overrides.GenerateDefaultDescription != nil {
		generateDefaultDescription = *overrides.GenerateDefaultDescription
	}

	if overrides.Description != nil {
		field.Description = *overrides.Description
	} else if generateDefaultDescription {
		field.Description = field.DefaultDescription()
	}

	return field
}

type FieldOverrides struct {
	Description *string `yaml:"description,omitempty"`

	// GenerateDefaultDescription indicates whether the default description for the field should be set. When this is not
	// set it defaults to `true`.
	//
	// The behaviour of this is strongly linked to Description. When this is true and Description is set, even to an
	// empty string, then the Description will override the default description. If Description is nil then this will
	// generate the default description.
	//
	// To remove the default description from the field without having to set Description to an empty string this should
	// be set to false.
	GenerateDefaultDescription *bool           `yaml:"generateDefaultDescription,omitempty"`
	JSONName                   *string         `yaml:"jsonName,omitempty"`
	Name                       *string         `yaml:"name,omitempty"`
	Type                       *spec.FieldType `yaml:"type,omitempty"`
	Validation                 *regexp.Regexp  `yaml:"validation,omitempty"`
}

// ApplyFieldOverrides will return a mutated spec.Field with the given FieldOverrides applied.
func ApplyFieldOverrides(field spec.Field, overrides *FieldOverrides) spec.Field {
	if overrides == nil {
		return field
	}

	if field.Description == field.DefaultDescription() {
		field.Description = ""
	}

	generateDefaultDescription := true
	if overrides.GenerateDefaultDescription != nil {
		generateDefaultDescription = *overrides.GenerateDefaultDescription && field.Description == ""
	}

	if overrides.Description != nil {
		field.Description = *overrides.Description
	}

	if overrides.JSONName != nil {
		field.JSONName = *overrides.JSONName
	}

	if overrides.Name != nil {
		field.Name = *overrides.Name
	}

	if overrides.Type != nil {
		field.Type = *overrides.Type
	}

	if overrides.Validation != nil {
		field.Validation = overrides.Validation
	}

	if overrides.Description != nil {
		field.Description = *overrides.Description
	} else if generateDefaultDescription && field.Description == "" {
		field.Description = field.DefaultDescription()
	}

	return field
}
