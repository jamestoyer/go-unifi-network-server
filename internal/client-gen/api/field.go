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
	"github.com/zclconf/go-cty/cty"
)

var (
	characterClassesRegex = regexp.MustCompile(`[\(\[\]\)\-]`)
	repetitionsRegex      = regexp.MustCompile(`\?|\+|\*|\{\d*,?\d*\}`)
)

type FieldDefinition struct {
	Name     string
	JSONName string
	Type     cty.Type
}

func NewFieldDefinition(jsonName string, value interface{}) (FieldDefinition, error) {
	switch value.(type) {
	case string:
		return definitionFromStringValue(jsonName, value)
	}

	return FieldDefinition{}, fmt.Errorf("unsupported type for field definition: %v", value)
}

func definitionFromStringValue(jsonName string, value interface{}) (FieldDefinition, error) {
	var ctyType cty.Type
	normalisedValidationString := normaliseValidationString(value.(string))
	if normalisedValidationString == "truefalse" || normalisedValidationString == "falsetrue" {
		ctyType = cty.Bool
	} else if _, err := strconv.ParseInt(normalisedValidationString, 10, 64); err == nil {
		ctyType = cty.Number
	} else if _, err := strconv.ParseFloat(normalisedValidationString, 64); err == nil {
		// TODO: (jtoyer) Handle floats differently to integers to ensure the structs are rendered correctly
		ctyType = cty.Number
	} else {
		ctyType = cty.String
	}

	return FieldDefinition{
		JSONName: jsonName,
		Name:     strcase.ToCamel(jsonName),
		Type:     ctyType,
	}, nil
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
