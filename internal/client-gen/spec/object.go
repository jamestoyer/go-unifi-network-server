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
	"maps"
	"slices"
	"strings"

	"github.com/iancoleman/strcase"
)

var rootObjectFields []Field

type Object struct {
	Name   string
	Fields []Field
}

func objectFromAPISpec(name string, values map[string]interface{}, addDefaultFields bool) (*Object, []*fieldObject, error) {
	if strings.TrimSpace(name) == "" {
		return nil, nil, errors.New("invalid endpoint name")
	}

	// Make the Object name consistent. This is going to be camel case to broadly get it to fit the Golang standards.
	// This can be overridden later
	name = strcase.ToCamel(name)

	var fields []Field
	if addDefaultFields {
		fields = append(fields, rootObjectFields...)
	}

	var fieldObjects []*fieldObject
	sortedFieldNames := slices.Sorted(maps.Keys(values))
	for _, fieldName := range sortedFieldNames {
		value := values[fieldName]
		field, object, err := parseFieldDefinition(fieldName, name, value)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse field %s: %w", fieldName, err)
		}

		fields = append(fields, field)
		if object != nil {
			fieldObjects = append(fieldObjects, object)
		}
	}

	return &Object{
		Name:   name,
		Fields: fields,
	}, fieldObjects, nil
}

func init() {
	id := Field{Name: "ID", Type: FieldTypeString, JSONName: "_id"}
	id.Description = id.DefaultDescription()
	siteID := Field{Name: "SiteID", Type: FieldTypeString, JSONName: "site_id"}
	siteID.Description = siteID.DefaultDescription()
	hidden := Field{Name: "Hidden", Type: FieldTypeBoolean, JSONName: "attr_hidden"}
	hidden.Description = hidden.DefaultDescription()
	hiddenID := Field{Name: "HiddenID", Type: FieldTypeString, JSONName: "attr_hidden_id"}
	hiddenID.Description = hiddenID.DefaultDescription()
	noDelete := Field{Name: "NoDelete", Type: FieldTypeBoolean, JSONName: "attr_no_delete"}
	noDelete.Description = noDelete.DefaultDescription()
	noEdit := Field{Name: "NoEdit", Type: FieldTypeBoolean, JSONName: "attr_no_edit"}
	noEdit.Description = noEdit.DefaultDescription()

	rootObjectFields = []Field{
		id,
		siteID,
		hidden,
		hiddenID,
		noDelete,
		noEdit,
	}
}
