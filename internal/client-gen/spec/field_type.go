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
)

var (
	FieldTypeBoolean = FieldType{
		fieldTypeImpl: primitiveFieldType{kind: "bool", defaultValue: "false"},
	}
	FieldTypeDecimal = FieldType{
		fieldTypeImpl: primitiveFieldType{kind: "float64", defaultValue: "0"},
	}
	FieldTypeNumber = FieldType{
		fieldTypeImpl: primitiveFieldType{kind: "int64", defaultValue: "0"},
	}
	FieldTypeString = FieldType{
		fieldTypeImpl: primitiveFieldType{kind: "string", defaultValue: `""`},
	}

	unknownType = FieldType{
		fieldTypeImpl: primitiveFieldType{kind: "Field Type Unknown"},
	}
)

// fieldTypeImpl should be implemented by anything that can be used as valid API field type.
type fieldTypeImpl interface {
	// GoType is the Golang type definition.
	GoType() string

	DefaultValue() string
}

// FieldType represents the derived type from the API specification for a field on an endpoint object.
type FieldType struct {
	fieldTypeImpl
}

// IsListType indicates whether the type is a Golang list. Use ElementType to determine the type of the elements
// expected for the list.
func (t *FieldType) IsListType() bool {
	_, ok := t.fieldTypeImpl.(listFieldType)
	return ok
}

// IsObjectType indicates whether the type is a Golang struct.
func (t *FieldType) IsObjectType() bool {
	_, ok := t.fieldTypeImpl.(objectFieldType)
	return ok
}

// IsPrimitiveType indicates whether the type is a Golang primitive.
func (t *FieldType) IsPrimitiveType() bool {
	_, ok := t.fieldTypeImpl.(primitiveFieldType)
	return ok
}

// ElementType returns the expected FieldType for all elements of a list. If this type is not a list the nil is
// returned.
func (t *FieldType) ElementType() *FieldType {
	if listType, ok := t.fieldTypeImpl.(listFieldType); ok {
		return &listType.elementType
	}

	return nil
}

func (t *FieldType) UnmarshalText(text []byte) error {
	if text == nil {
		return errors.New("text must not be nil")
	}

	value := strings.TrimSpace(string(text))

	ft := unmarshalFieldType(value)
	*t = ft
	if ft == unknownType {
		return fmt.Errorf("unknown field type: %s", value)
	}

	return nil
}

func unmarshalFieldType(value string) FieldType {
	switch value {
	case "Boolean":
		return FieldTypeBoolean
	case "Decimal":
		return FieldTypeDecimal
	case "Number":
		return FieldTypeNumber
	case "String":
		return FieldTypeString
	default:
		if strings.HasPrefix(value, "List(") && strings.HasSuffix(value, ")") {
			return unmarshalListFieldType(value)
		}

		if strings.HasPrefix(value, "Object(") && strings.HasSuffix(value, ")") {
			return unmarshalObjectFieldType(value)
		}

		return unknownType
	}
}

func unmarshalListFieldType(value string) FieldType {
	value = strings.TrimPrefix(value, "List(")
	value = strings.TrimSuffix(value, ")")

	elementType := unmarshalFieldType(value)
	if elementType == unknownType {
		return elementType
	}

	return FieldTypeList(elementType)
}

func unmarshalObjectFieldType(value string) FieldType {
	value = strings.TrimPrefix(value, "Object(")
	value = strings.TrimSuffix(value, ")")

	if strings.ContainsAny(value, "()") {
		return unknownType
	}

	return FieldTypeObject(value)
}

type primitiveFieldType struct {
	defaultValue string
	kind         string
}

func (t primitiveFieldType) GoType() string {
	return t.kind
}

func (t primitiveFieldType) DefaultValue() string {
	return t.defaultValue
}

type listFieldType struct {
	elementType FieldType
}

func (t listFieldType) GoType() string {
	return "[]" + t.elementType.GoType()
}

func (t listFieldType) DefaultValue() string {
	return "nil"
}

func FieldTypeList(element FieldType) FieldType {
	return FieldType{
		fieldTypeImpl: listFieldType{elementType: element},
	}
}

type objectFieldType struct {
	kind string
}

func (t objectFieldType) GoType() string {
	return t.kind
}

func (t objectFieldType) DefaultValue() string {
	return "nil"
}

func FieldTypeObject(name string) FieldType {
	return FieldType{
		fieldTypeImpl: objectFieldType{kind: name},
	}
}
