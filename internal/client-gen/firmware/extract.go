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

package firmware

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
)

const (
	apiJarFilename = "ace.jar"
	apiDir         = "api/fields"
)

func ExtractAPI(ctx context.Context, logger *slog.Logger, filename, destination string) error {
	logger.DebugContext(ctx, "Unzipping server download")
	zipReader, err := zip.OpenReader(filename)
	if err != nil {
		return fmt.Errorf("failed to open zip: %w", err)
	}

	var aceFile *zip.File
	for _, file := range zipReader.File {
		if filepath.Base(file.Name) == apiJarFilename {
			logger.DebugContext(ctx, "API jar found")
			aceFile = file
			break
		}
	}

	if aceFile == nil {
		return fmt.Errorf("failed to find %s in zip", apiJarFilename)
	}

	logger.DebugContext(ctx, "Opening API jar file")
	aceReader, err := aceFile.Open()
	if err != nil {
		return fmt.Errorf("failed to open file in zip: %w", err)
	}

	defer func() {
		if err := aceReader.Close(); err != nil {
			logger.ErrorContext(ctx, "Failed to close file", "error", err)
		}
	}()

	logger.DebugContext(ctx, "Reading API jar contents")
	aceContents, err := io.ReadAll(aceReader)
	if err != nil {
		return fmt.Errorf("failed to read jar: %w", err)
	}

	jarReader, err := zip.NewReader(bytes.NewReader(aceContents), int64(len(aceContents)))
	if err != nil {
		return fmt.Errorf("failed to open zip: %w", err)
	}

	for _, file := range jarReader.File {
		if filepath.Dir(file.Name) != apiDir || filepath.Ext(file.Name) == "" {
			logger.DebugContext(ctx, "Skipping non API file", "name", file.Name)
			continue
		}

		logger.InfoContext(ctx, "Extracting API definition", "name", file.Name)
		f, err := file.Open()
		if err != nil {
			return fmt.Errorf("failed to open file in zip: %w", err)
		}

		dst, err := os.Create(filepath.Join(destination, filepath.Base(file.Name)))
		if err != nil {
			_ = f.Close()
			return fmt.Errorf("unable to create API destination file: %w", err)
		}

		if _, err := io.Copy(dst, f); err != nil {
			_ = f.Close()
			_ = dst.Close()
			return fmt.Errorf("failed to copy file into API destination file: %w", err)
		}

		_ = f.Close()
		_ = dst.Close()
	}

	return nil
}
