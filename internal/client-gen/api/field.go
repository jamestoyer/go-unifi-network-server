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
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
)

var (
	characterClassesRegex = regexp.MustCompile(`[\(\[\]\)\-]`)
	repetitionsRegex      = regexp.MustCompile(`\?|\+|\*|\{\d*,?\d*\}`)
)

type FieldDefinition struct {
	Name     string
	JSONName string
	Type     FieldType
	Object   *EndpointObject
}

func NewFieldDefinition(jsonName, endpointObjectName string, value interface{}) (FieldDefinition, error) {
	fieldDefinition := FieldDefinition{
		JSONName: jsonName,
		Name:     strcase.ToCamel(jsonName),
	}

	fieldType, endpointObject, err := getFieldType(jsonName, endpointObjectName, value)
	if err != nil {
		return fieldDefinition, fmt.Errorf("unable to get field type: %w", err)
	}

	fieldDefinition.Type = fieldType
	fieldDefinition.Object = endpointObject
	return fieldDefinition, nil
}

func getFieldType(jsonName, endpointObjectName string, value interface{}) (FieldType, *EndpointObject, error) {
	var endpointObject *EndpointObject
	var fieldType FieldType
	switch t := value.(type) {
	case string:
		fieldType = fieldTypeFromStringValue(value)
	case []interface{}:
		ft, object, err := fieldTypeFromInterfaceValue(jsonName, endpointObjectName, t)
		if err != nil {
			return UnknownType, nil, err
		}

		endpointObject = object
		fieldType = ft
	case map[string]interface{}:
		ft, object, err := fieldTypeFromObjectValue(jsonName, endpointObjectName, t)
		if err != nil {
			return UnknownType, nil, err
		}

		endpointObject = object
		fieldType = ft
	default:
		return UnknownType, nil, fmt.Errorf("unsupported type for field definition: %v", value)
	}

	return fieldType, endpointObject, nil
}

func fieldTypeFromInterfaceValue(jsonName, endpointObjectName string, value []interface{}) (FieldType, *EndpointObject, error) {
	switch len(value) {
	case 0:
		return List(String), nil, nil
	case 1:
		ft, object, err := getFieldType(jsonName, endpointObjectName, value[0])
		if err != nil {
			return UnknownType, nil, err
		}

		return List(ft), object, nil
	default:
		return UnknownType, nil, fmt.Errorf("unsupported validation for list: %v", value)
	}
}

func fieldTypeFromObjectValue(jsonName, endpointObjectName string, values map[string]interface{}) (FieldType, *EndpointObject, error) {
	name := strcase.ToCamel(jsonName)
	fieldType := Object(endpointObjectName + name)
	object, err := NewEndpointObject(name, values, endpointObjectName)
	if err != nil {
		return FieldType{}, nil, fmt.Errorf("failed to create endpoint object for field: %w", err)
	}

	return fieldType, object, err
}

func fieldTypeFromStringValue(value interface{}) FieldType {
	normalisedValidationString := normaliseValidationString(value.(string))
	if normalisedValidationString == "truefalse" || normalisedValidationString == "falsetrue" {
		return Boolean
	}

	if _, err := strconv.ParseInt(normalisedValidationString, 10, 64); err == nil {
		return Number
	}

	if _, err := strconv.ParseFloat(normalisedValidationString, 64); err == nil {
		return Decimal
	}

	return String
}

// normaliseValidationString processes the validation string such that it can be subsequently used to work out if the
// type of the field is a string, bool, integer or float.
//
// This function was heavily influenced by
// https://github.com/paultyng/go-unifi/blob/971b9fc3b09f193177cd65e7e2b2e6eabe8e67a3/fields/main.go#L606-L624.
func normaliseValidationString(value string) string {
	if value == "" {
		return value
	}
	value = strings.TrimPrefix(value, "^")
	value = strings.TrimSuffix(value, "$")

	// Character class verbose replacements
	value = strings.ReplaceAll(value, `\d`, "[0-9]")

	// Conditionals
	value = strings.ReplaceAll(value, "|", "")

	// Flags
	value = strings.ReplaceAll(value, "$", "")
	value = strings.ReplaceAll(value, "^", "")

	// Regex replacements
	value = repetitionsRegex.ReplaceAllString(value, "")
	value = characterClassesRegex.ReplaceAllString(value, "")

	return value
}
