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
)

// The FieldType represents information about the type of field derived from the API spec.
type FieldType interface {
	// GoType is the Golang type definition.
	GoType() string
}

// fieldTypeImpl is a help struct that implements FieldType for easy variable construction.
type fieldTypeImpl string

func (t fieldTypeImpl) GoType() string {
	return string(t)
}

const (
	Boolean fieldTypeImpl = "bool"
	Decimal fieldTypeImpl = "float64"
	Number  fieldTypeImpl = "int64"
	String  fieldTypeImpl = "string"

	UnknownType fieldTypeImpl = "Field Type Unknown"
)

func List(element FieldType) FieldType {
	return fieldTypeImpl(fmt.Sprintf("[]%s", element))
}
