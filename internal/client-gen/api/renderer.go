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
	"bufio"
	"bytes"
	"context"
	"embed"
	"fmt"
	"go/format"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/jamestoyer/go-unifi-network-server/internal/client-gen/spec"
	"github.com/jamestoyer/go-unifi-network-server/internal/logging"
)

//go:embed templates
var templatesFs embed.FS

type Renderer struct {
	templates *template.Template
}

func NewRenderer() (*Renderer, error) {
	tmpl := template.New("default").Funcs(template.FuncMap{
		"goDescription": goDescription,
	})
	tmpl, err := tmpl.ParseFS(templatesFs, "templates/*")
	if err != nil {
		return nil, fmt.Errorf("failed to parse API templates: %w", err)
	}

	return &Renderer{
		templates: tmpl,
	}, nil
}

func (r *Renderer) RenderEndpoint(ctx context.Context, endpoint *spec.Endpoint, targetDir, packageName string) error {
	filename := filepath.Join(targetDir, generatedFilename(endpoint.Name))
	ctx = logging.CtxWithValues(ctx, slog.String("endpoint", endpoint.Name), slog.String("filename", filename))
	slog.DebugContext(ctx, "Rendering endpoint")

	var buffer bytes.Buffer
	templateContent := struct {
		Actions      spec.EndpointActions
		Name         string
		Objects      []*spec.Object
		PackageName  string
		ResourcePath string
	}{
		Actions:      endpoint.Actions,
		Name:         endpoint.Name,
		Objects:      endpoint.Objects,
		PackageName:  packageName,
		ResourcePath: endpoint.ResourcePath,
	}
	if err := r.templates.ExecuteTemplate(&buffer, "service.gotmpl", templateContent); err != nil {
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

func (r *Renderer) RenderEndpoints(ctx context.Context, endpoints []*spec.Endpoint, targetDir, packageName string) error {
	slog.DebugContext(ctx, "Rendering endpoints", slog.Int("endpointCount", len(endpoints)))
	for _, endpoint := range endpoints {
		if err := r.RenderEndpoint(ctx, endpoint, targetDir, packageName); err != nil {
			return err
		}
	}

	return nil
}

func (r *Renderer) RenderClient(ctx context.Context, endpoints []*spec.Endpoint, targetDir, packageName string) error {
	slog.DebugContext(ctx, "Rendering client", slog.Int("endpointCount", len(endpoints)))
	if err := r.RenderEndpoints(ctx, endpoints, targetDir, packageName); err != nil {
		return fmt.Errorf("failed to render client endpoints: %w", err)
	}

	var buffer bytes.Buffer
	templateContent := struct {
		PackageName string
		Endpoints   []*spec.Endpoint
	}{
		Endpoints:   endpoints,
		PackageName: packageName,
	}
	if err := r.templates.ExecuteTemplate(&buffer, "client.gotmpl", templateContent); err != nil {
		return fmt.Errorf("failed to generate API client file: %w", err)
	}

	src, err := format.Source(buffer.Bytes())
	if err != nil {
		return fmt.Errorf("failed to format API endpoint: %w", err)
	}

	filename := filepath.Join(targetDir, generatedFilename("client"))
	if err := os.WriteFile(filename, src, 0o644); err != nil {
		return fmt.Errorf("failed to write API endpoint file: %w", err)
	}

	return nil
}

func generatedFilename(name string) string {
	return strcase.ToSnake(name) + golangGeneratedFileSuffix
}

func goDescription(description string) string {
	if description == "" {
		return description
	}

	builder := strings.Builder{}
	scanner := bufio.NewScanner(strings.NewReader(description))
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			builder.WriteString("//")
		} else {
			builder.WriteString("// " + scanner.Text())
		}

		builder.WriteString(fmt.Sprintln())
	}

	d := strings.TrimSpace(builder.String())

	return d
}
