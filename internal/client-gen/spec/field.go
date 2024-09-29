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
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
)

var (
	defaultField = Field{}

	characterClassesRegex = regexp.MustCompile(`[\(\[\]\)\-]`)
	repetitionsRegex      = regexp.MustCompile(`\?|\+|\*|\{\d*,?\d*\}`)
)

type FieldOverrides struct {
	Description *string        `yaml:"description,omitempty"`
	JSONName    *string        `yaml:"jsonName,omitempty"`
	Name        *string        `yaml:"name,omitempty"`
	Type        *FieldType     `yaml:"type,omitempty"`
	Validation  *regexp.Regexp `yaml:"validation,omitempty"`
}

type Field struct {
	Name        string
	Description string
	Type        FieldType

	JSONName   string
	Validation *regexp.Regexp
}

// ApplyOverrides will return a mutated Field with the given FieldOverrides applied.
//
// When useDefaultDescription is true the description will automatically be updated to a default generated description,
// which is useful when updating the name or validation of the field. However, if this is true and the description is
// explicitly overridden then the override will take priority.
func (f Field) ApplyOverrides(overrides *FieldOverrides, useDefaultDescription bool) Field {
	if overrides == nil {
		if useDefaultDescription {
			f.Description = f.defaultDescription()
		}

		return f
	}

	if overrides.Description != nil {
		f.Description = *overrides.Description
	}

	if overrides.JSONName != nil {
		f.JSONName = *overrides.JSONName
	}

	if overrides.Name != nil {
		f.Name = strcase.ToCamel(*overrides.Name)
	}

	if overrides.Type != nil {
		f.Type = *overrides.Type
	}

	if overrides.Validation != nil {
		f.Validation = overrides.Validation
	}

	if overrides.Description != nil {
		f.Description = *overrides.Description
	} else if useDefaultDescription {
		f.Description = f.defaultDescription()
	}

	return f
}

func (f Field) defaultDescription() string {
	var validationPrefix string
	var validationRegex *regexp.Regexp

	switch {
	case f.Type.IsPrimitiveType():
		validationRegex = f.Validation
	case f.Type.IsListType():
		validationPrefix = "Element "
		if f.Type.ElementType().IsPrimitiveType() {
			validationRegex = f.Validation
		}
	}

	validationString := "None"
	if validationRegex != nil && validationRegex.String() != "" {
		validationString = validationRegex.String()
	}

	return fmt.Sprintf(`%s has been auto generated from the Unifi Network Server API specification

%sValidation: %s`, f.Name, validationPrefix, validationString)
}

type fieldObject struct {
	Name  string
	Value map[string]interface{}
}

func parseFieldDefinition(name, parentObject string, value interface{}) (Field, *fieldObject, error) {
	fieldType, object, err := getFieldType(name, parentObject, value)
	if err != nil {
		return defaultField, nil, fmt.Errorf("failed to parse type for field %s: %w", name, err)
	}

	f := Field{
		Name:     strcase.ToCamel(name),
		Type:     fieldType,
		JSONName: name,
	}

	regex, err := getValidationRegex(value)
	if err != nil {
		return defaultField, nil, fmt.Errorf("invalid validation regex for field %s: %w", name, err)
	}

	f.Validation = regex
	f.Description = f.defaultDescription()

	return f, object, nil
}

func getFieldType(name, parentObject string, value interface{}) (FieldType, *fieldObject, error) {
	switch t := value.(type) {
	case string:
		return fieldTypeFromStringValue(value), nil, nil
	case []interface{}:
		return fieldTypeFromInterfaceValue(name, parentObject, t)
	case map[string]interface{}:
		ft, object := fieldTypeFromObjectValue(name, parentObject, t)
		return ft, object, nil
	}

	return unknownType, nil, fmt.Errorf("unsupported type for field definition: %v", value)
}

func getValidationRegex(value interface{}) (*regexp.Regexp, error) {
	var validationString string
	switch v := value.(type) {
	case string:
		validationString = v
	case []interface{}:
		if len(v) == 1 {
			if _, ok := v[0].(string); ok {
				validationString = v[0].(string)
			}
		}
	}

	if validationString == "" {
		return nil, nil
	}

	regex, err := regexp.Compile(validationString)
	if err != nil {
		return nil, fmt.Errorf("failed to compile validation regex: %w", err)
	}

	return regex, nil
}

func fieldTypeFromInterfaceValue(name, parentObject string, value []interface{}) (FieldType, *fieldObject, error) {
	switch len(value) {
	case 0:
		return FieldTypeList(FieldTypeString), nil, nil
	case 1:
		ft, object, err := getFieldType(name, parentObject, value[0])
		return FieldTypeList(ft), object, err
	default:
		return FieldTypeList(unknownType), nil, fmt.Errorf("unsupported validation for list: %v", value)
	}
}

func fieldTypeFromObjectValue(name, parentObject string, values map[string]interface{}) (FieldType, *fieldObject) {
	// Explicitly add a dash here to ensure correct casing for the object name
	name = strcase.ToCamel(parentObject + "-" + name)
	return FieldTypeObject(name), &fieldObject{Name: name, Value: values}
}

func fieldTypeFromStringValue(value interface{}) FieldType {
	normalisedValidationString := normaliseValidationString(value.(string))
	if normalisedValidationString == "truefalse" || normalisedValidationString == "falsetrue" {
		return FieldTypeBoolean
	}

	if _, err := strconv.ParseInt(normalisedValidationString, 10, 64); err == nil {
		return FieldTypeNumber
	}

	if _, err := strconv.ParseFloat(normalisedValidationString, 64); err == nil {
		return FieldTypeDecimal
	}

	return FieldTypeString
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

// func getFieldType(jsonName, endpointObjectName string, value interface{}) (FieldType, *EndpointObject, error) {
// 	var endpointObject *EndpointObject
// 	var fieldType FieldType
// 	switch t := value.(type) {
// 	case string:
// 		fieldType = fieldTypeFromStringValue(value)
// 	case []interface{}:
// 		ft, object, err := fieldTypeFromInterfaceValue(jsonName, endpointObjectName, t)
// 		if err != nil {
// 			return UnknownType, nil, err
// 		}
//
// 		endpointObject = object
// 		fieldType = ft
// 	case map[string]interface{}:
// 		ft, object, err := fieldTypeFromObjectValue(jsonName, endpointObjectName, t)
// 		if err != nil {
// 			return UnknownType, nil, err
// 		}
//
// 		endpointObject = object
// 		fieldType = ft
// 	default:
// 		return UnknownType, nil, fmt.Errorf("unsupported type for field definition: %v", value)
// 	}
//
// 	return fieldType, endpointObject, nil
// }
