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

var (
	Boolean = FieldType{
		fieldTypeImpl: primitiveFieldType{kind: "bool"},
	}
	Decimal = FieldType{
		fieldTypeImpl: primitiveFieldType{kind: "float64"},
	}
	Number = FieldType{
		fieldTypeImpl: primitiveFieldType{kind: "int64"},
	}
	String = FieldType{
		fieldTypeImpl: primitiveFieldType{kind: "string"},
	}

	UnknownType = FieldType{
		fieldTypeImpl: primitiveFieldType{kind: "Field Type Unknown"},
	}
)

// fieldTypeImpl should be implemented by anything that can be used as valid API field type.
type fieldTypeImpl interface {
	// GoType is the Golang type definition.
	GoType() string

	// // IsObjectType indicates whether the type is a Golang struct.
	// // TODO: (jtoyer) Add information about how to get the object definition
	// IsObjectType() bool
}

// FieldType represents the derived type from the API specification for a field on an endpoint object.
type FieldType struct {
	fieldTypeImpl
}

// IsListType indicates whether the type is a Golang list. Use ElementType to determine the type of the elements
// expected for the list.
func (t FieldType) IsListType() bool {
	_, ok := t.fieldTypeImpl.(listFieldType)
	return ok
}

// IsPrimitiveType indicates whether the type is a Golang primitive.
func (t FieldType) IsPrimitiveType() bool {
	_, ok := t.fieldTypeImpl.(primitiveFieldType)
	return ok
}

// ElementType returns the expected FieldType for all elements of a list. If this type is not a list the nil is
// returned.
func (t FieldType) ElementType() *FieldType {
	if listType, ok := t.fieldTypeImpl.(listFieldType); ok {
		return &listType.elementType
	}

	return nil
}

func (t FieldType) String() string {
	return t.fieldTypeImpl.GoType()
}

type primitiveFieldType struct {
	kind string
}

func (t primitiveFieldType) GoType() string {
	return t.kind
}

type listFieldType struct {
	elementType FieldType
}

func (t listFieldType) GoType() string {
	return "[]" + t.elementType.GoType()
}

func List(element FieldType) FieldType {
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

func Object(name string) FieldType {
	return FieldType{
		fieldTypeImpl: objectFieldType{kind: name},
	}
}
