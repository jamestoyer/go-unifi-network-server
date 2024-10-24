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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldType_UnmarshalText(t *testing.T) {
	tests := map[string]struct {
		data    []byte
		want    FieldType
		wantErr assert.ErrorAssertionFunc
	}{
		"data is empty": {
			data:    []byte(""),
			want:    unknownType,
			wantErr: assert.Error,
		},
		"data is invalid YAML": {
			data:    []byte("'"),
			want:    unknownType,
			wantErr: assert.Error,
		},
		"data is a string": {
			data:    []byte("String"),
			want:    FieldTypeString,
			wantErr: assert.NoError,
		},
		"data is a boolean": {
			data:    []byte("Boolean"),
			want:    FieldTypeBoolean,
			wantErr: assert.NoError,
		},
		"data is a number": {
			data:    []byte("Number"),
			want:    FieldTypeNumber,
			wantErr: assert.NoError,
		},
		"data is a decimal": {
			data:    []byte("Decimal"),
			want:    FieldTypeDecimal,
			wantErr: assert.NoError,
		},
		"data is a unix time": {
			data:    []byte("UnixTime"),
			want:    FieldTypeUnixTime,
			wantErr: assert.NoError,
		},
		"data is a list": {
			data:    []byte("List(String)"),
			want:    FieldTypeList(FieldTypeString),
			wantErr: assert.NoError,
		},
		"data is an object": {
			data:    []byte("Object(AnObject)"),
			want:    FieldTypeObject("AnObject"),
			wantErr: assert.NoError,
		},
		"data is a list of object": {
			data:    []byte("List(Object(Nested))"),
			want:    FieldTypeList(FieldTypeObject("Nested")),
			wantErr: assert.NoError,
		},
		"data is missing trailing list bracket": {
			data:    []byte("List(String"),
			want:    unknownType,
			wantErr: assert.Error,
		},
		"data is missing trailing object bracket": {
			data:    []byte("Object(String"),
			want:    unknownType,
			wantErr: assert.Error,
		},
		"data is missing trailing nested bracket": {
			data:    []byte("List(Object(String)"),
			want:    unknownType,
			wantErr: assert.Error,
		},
		"data is invalid type": {
			data:    []byte("Not a real type"),
			want:    unknownType,
			wantErr: assert.Error,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := FieldType{}
			err := got.UnmarshalText(test.data)
			test.wantErr(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}
