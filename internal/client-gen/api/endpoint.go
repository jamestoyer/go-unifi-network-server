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
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"text/template"

	"github.com/iancoleman/strcase"
)

type Endpoint struct {
	Name       string
	goFilePath string
	rootObject *EndpointObject
}

func NewEndpoint(name string, spec map[string]interface{}) (*Endpoint, error) {
	rootObject, err := NewEndpointObject(name, spec, "")
	if err != nil {
		return nil, fmt.Errorf("failed to generate objects for endpoint: %w", err)
	}

	return &Endpoint{
		Name:       name,
		goFilePath: strcase.ToSnake(name) + ".generated.go",
		rootObject: rootObject,
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

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create API endpoint file: %w", err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			logger.ErrorContext(ctx, "Failed to close file", "error", err)
		}
	}()

	endpointTmpl := endpointTemplate{
		PackageName: modulePath,
		Structs:     e.Objects(),
	}
	if err = tmpl.ExecuteTemplate(f, "file.gotmpl", endpointTmpl); err != nil {
		return fmt.Errorf("failed to execute generate API spec file: %w", err)
	}

	return nil
}

type endpointTemplate struct {
	PackageName string
	Structs     []*EndpointObject
}
