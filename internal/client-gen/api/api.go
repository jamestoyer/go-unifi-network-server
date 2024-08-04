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
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

var (
	skippedFiles = []string{
		"AuthenticationRequest.json",
		"HeatMap.json",
		"HeatMapPoint.json",
		"Map.json",
		"MediaFile.json",
		"Setting.json",
		"SpatialRecord.json",
		"VirtualDevice.json",
		"Wall.json",
	}
)

func Generate(ctx context.Context, logger *slog.Logger, apiSpecDir string) error {
	var errs []error
	err := filepath.WalkDir(apiSpecDir, func(path string, d fs.DirEntry, err error) error {
		switch {
		case err != nil:
			return err
		case d.IsDir():
			return nil
		case !strings.HasSuffix(path, ".json"):
			return nil
		case slices.Contains(skippedFiles, d.Name()):
			logger.DebugContext(ctx, "Explicitly skipping file", slog.String("file", d.Name()))
			return nil
		}

		if err := generateAPIFile(ctx, logger, path); err != nil {
			errs = append(errs, fmt.Errorf("unable to generate API client: %w", err))
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to walk API spec dir: %w", err)
	}

	return errors.Join(errs...)
}

func generateAPIFile(ctx context.Context, logger *slog.Logger, file string) error {
	contents, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("failed to read API spec file: %w", err)
	}

	var fields map[string]interface{}
	if err = json.Unmarshal(contents, &fields); err != nil {
		return fmt.Errorf("failed to unmarshal API spec file: %w", err)
	}

	for name, value := range fields {
		fmt.Println(name, value)
	}

	return nil
}
