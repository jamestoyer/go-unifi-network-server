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

	"github.com/iancoleman/strcase"
	"github.com/jamestoyer/go-unifi-network-server/internal/client-gen/spec"
)

type EndpointOverrides struct {
	Name         *string `yaml:"name,omitempty"`
	ResourcePath *string `yaml:"resourcePath,omitempty"`
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

	return &ep
}

type FieldOverrides struct {
	Description *string         `yaml:"description,omitempty"`
	JSONName    *string         `yaml:"jsonName,omitempty"`
	Name        *string         `yaml:"name,omitempty"`
	Type        *spec.FieldType `yaml:"type,omitempty"`
	Validation  *regexp.Regexp  `yaml:"validation,omitempty"`
}

// ApplyFieldOverrides will return a mutated spec.Field with the given FieldOverrides applied.
//
// When useDefaultDescription is true the description will automatically be updated to a default generated description,
// which is useful when updating the name or validation of the field. However, if this is true and the description is
// explicitly overridden then the override will take priority.
func ApplyFieldOverrides(field spec.Field, overrides *FieldOverrides, useDefaultDescription bool) spec.Field {
	if overrides == nil {
		if useDefaultDescription {
			field.Description = field.DefaultDescription()
		}

		return field
	}

	if overrides.Description != nil {
		field.Description = *overrides.Description
	}

	if overrides.JSONName != nil {
		field.JSONName = *overrides.JSONName
	}

	if overrides.Name != nil {
		field.Name = strcase.ToCamel(*overrides.Name)
	}

	if overrides.Type != nil {
		field.Type = *overrides.Type
	}

	if overrides.Validation != nil {
		field.Validation = overrides.Validation
	}

	if overrides.Description != nil {
		field.Description = *overrides.Description
	} else if useDefaultDescription {
		field.Description = field.DefaultDescription()
	}

	return field
}
