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
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const golangGeneratedFileSuffix = ".generated.go"

func Generate(ctx context.Context, logger *slog.Logger, endpoints []*Endpoint, outputDir string) error {
	slog.InfoContext(ctx, "Removing previously generated files")
	if err := RemoveGeneratedFiles(ctx, outputDir); err != nil {
		return err
	}

	slog.InfoContext(ctx, "Generating endpoints")
	tmpl, err := template.ParseFS(templatesFs, "templates/*")
	if err != nil {
		return fmt.Errorf("error parsing API templates: %w", err)
	}

	for _, endpoint := range endpoints {
		if err = endpoint.Render(ctx, logger, outputDir, tmpl); err != nil {
			return fmt.Errorf("failed to render API endpoint: %w", err)
		}
	}

	return nil
}

func RemoveGeneratedFiles(ctx context.Context, outputDir string) error {
	err := filepath.WalkDir(outputDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() && path != outputDir {
			return filepath.SkipDir
		}

		if strings.HasSuffix(filepath.Base(path), golangGeneratedFileSuffix) {
			slog.DebugContext(ctx, "Removing generated file", slog.String("filename", path))
			if err = os.Remove(path); err != nil {
				return fmt.Errorf("failed to remove generated file: %w", err)
			}
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("unable to remove generated files: %w", err)
	}

	return nil
}
