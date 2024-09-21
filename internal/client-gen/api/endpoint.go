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
	"context"
	"log/slog"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
)

type Endpoint struct {
	Name    string
	Structs []Struct

	specPath string
}

func NewEndpoint(specPath string) *Endpoint {
	name := strings.TrimSuffix(filepath.Base(specPath), filepath.Ext(specPath))
	return &Endpoint{
		Name:     name,
		specPath: specPath,
	}
}

func (e *Endpoint) GoFilename() string {
	return strcase.ToSnake(e.Name) + ".generated.go"
}

func (e *Endpoint) UpdateStructs(ctx context.Context, logger *slog.Logger, apiFields map[string]interface{}) error {
	s := Struct{
		Name:   e.Name,
		Fields: map[string]Field{},
	}

	for name, value := range apiFields {
		switch t := value.(type) {
		case string:
			s.Fields[name] = NewField("", name, "string")
		default:
			logger.DebugContext(ctx, "Skipping field with unsupported type", slog.String("name", name), slog.Any("type", t))
			continue
		}
	}

	e.Structs = append(e.Structs, s)
	return nil
}

type Struct struct {
	Name string

	Fields map[string]Field
}
