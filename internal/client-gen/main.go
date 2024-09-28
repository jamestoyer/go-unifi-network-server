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

package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/jamestoyer/go-unifi-network-server/internal/client-gen/api"
	"github.com/jamestoyer/go-unifi-network-server/internal/client-gen/firmware"
	"github.com/jamestoyer/go-unifi-network-server/internal/logging"
)

const packageName = "networkserver"

var (
	verboseFlag    = flag.Bool("v", false, "Print verbose logs")
	apiSpecDirFlag = flag.String("dir", "", "Directory containing api spec files. If not set the latest API spec will be downloaded")

	logger *slog.Logger
)

func main() {
	flag.Parse()

	ctx := context.Background()

	logLevel := slog.LevelInfo
	if *verboseFlag {
		logLevel = slog.LevelDebug
	}

	logger = slog.New(logging.Handler{Handler: slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel})})
	slog.SetDefault(logger)
	logger.InfoContext(ctx, "Generating API client")

	var apiSpecDir string
	if *apiSpecDirFlag == "" {
		logger.DebugContext(ctx, "No API spec directory specified, downloading the latest API spec")
		downloadedAPISpecDir, err := downloadAPISpec(ctx)
		if err != nil {
			logger.ErrorContext(ctx, "Failed to download the API specification", slog.Any("error", err))
			os.Exit(1)
		}

		apiSpecDir = downloadedAPISpecDir
	} else {
		logger.DebugContext(ctx, "Using API spec directory", slog.Any("directory", *apiSpecDirFlag))
		apiSpecDir = *apiSpecDirFlag
	}

	err := generateAPIClient(ctx, apiSpecDir)
	if err != nil {
		logger.ErrorContext(ctx, "Failed to generate API client", slog.Any("error", err))
		os.Exit(1)
	}

	logger.InfoContext(ctx, "Generated API client")
}

func downloadAPISpec(ctx context.Context) (string, error) {
	logger.InfoContext(ctx, "Downloading Unifi Network Server")
	firmwareClient, err := firmware.NewClient("")
	if err != nil {
		return "", fmt.Errorf("failed to create a new firmware client: %w", err)
	}

	archiveFile, version, err := firmwareClient.DownloadLatestVersion(ctx, logger)
	if err != nil {
		return "", fmt.Errorf("failed downloaded the latest firmware version: %w", err)
	}

	defer func() {
		// Remove the downloaded archive file after we're done getting the API files out
		_ = os.Remove(archiveFile)
	}()

	logger.InfoContext(ctx, "Extracting API files")
	wd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current working directory: %w", err)
	}

	apiDestination := filepath.Join(wd, "build", "api", version.Version.Core().String())
	if err = os.MkdirAll(apiDestination, 0o755); err != nil {
		return "", fmt.Errorf("failed to create directory %s: %w", apiDestination, err)
	}

	if err = firmware.ExtractAPI(ctx, logger, archiveFile, apiDestination); err != nil {
		return "", fmt.Errorf("failed to extract Unifi Network Server APIs: %w", err)
	}

	return apiDestination, nil
}

func generateAPIClient(ctx context.Context, apiSpecDir string) error {
	if err := os.MkdirAll(packageName, 0o775); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", packageName, err)
	}

	return api.Generate(ctx, logger, apiSpecDir, packageName)
}
