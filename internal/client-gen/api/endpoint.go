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
	"bytes"
	"context"
	"fmt"
	"go/format"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

type Endpoint struct {
	Name           string
	definitionFile string
	goFilePath     string
	rootObject     *EndpointObject
}

func NewEndpoint(name, file string, spec map[string]interface{}) (*Endpoint, error) {
	// Ensure the endpoint name is always camel case, which is analogous to title case for our purposes.
	name = strcase.ToCamel(name)
	rootObject, err := NewEndpointObject(name, spec, "", true)
	if err != nil {
		return nil, fmt.Errorf("failed to generate objects for endpoint: %w", err)
	}

	return &Endpoint{
		Name:           name,
		definitionFile: file,
		goFilePath:     strcase.ToSnake(name) + golangGeneratedFileSuffix,
		rootObject:     rootObject,
	}, nil
}

// Objects traverses all field definitions and returns a list of all found EndpointObject. This will include the root
// EndpointObject for an Endpoint.
func (e *Endpoint) Objects() []*EndpointObject {
	objects := []*EndpointObject{e.rootObject}
	objects = append(objects, e.rootObject.NestedObjects()...)
	return objects
}

func (e *Endpoint) Render(ctx context.Context, logger *slog.Logger, modulePath string, tmpl *template.Template) error {
	filename := filepath.Join(modulePath, e.goFilePath)
	logger.InfoContext(ctx, "Rendering endpoint", slog.String("endpoint", e.Name), slog.String("filename", filename))

	var buffer bytes.Buffer
	endpointTmpl := endpointTemplate{
		APIPath:     strings.ToLower(e.Name),
		Name:        e.Name,
		PackageName: modulePath,
		Structs:     e.Objects(),
	}
	if err := tmpl.ExecuteTemplate(&buffer, "file.gotmpl", endpointTmpl); err != nil {
		return fmt.Errorf("failed to generate API endpoint file: %w", err)
	}

	src, err := format.Source(buffer.Bytes())
	if err != nil {
		return fmt.Errorf("failed to format API endpoint: %w", err)
	}

	if err := os.WriteFile(filename, src, 0o644); err != nil {
		return fmt.Errorf("failed to write API endpoint file: %w", err)
	}

	return nil
}

type endpointTemplate struct {
	APIPath     string
	Name        string
	PackageName string
	Structs     []*EndpointObject
}
