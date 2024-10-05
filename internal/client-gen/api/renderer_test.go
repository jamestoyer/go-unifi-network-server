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
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_goDescription(t *testing.T) {
	tests := map[string]struct {
		description string
		want        string
	}{
		"empty string": {
			description: "",
			want:        "",
		},
		"single line string": {
			description: "A single line string",
			want:        "// A single line string",
		},
		"multi line string": {
			description: `A multi line string

With some blank lines in them`,
			want: `// A multi line string
//
// With some blank lines in them`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := goDescription(test.description)
			assert.Equal(t, test.want, got)
		})
	}
}
