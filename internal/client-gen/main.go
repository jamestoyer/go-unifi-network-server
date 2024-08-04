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
	"log/slog"
	"os"
	"path/filepath"

	"github.com/jamestoyer/go-unifi-network-server/internal/client-gen/firmware"
)

var (
	verbose = flag.Bool("v", false, "Print verbose logs")

	logger *slog.Logger
)

func main() {
	flag.Parse()

	ctx := context.Background()

	logLevel := slog.LevelInfo
	if verbose != nil && *verbose {
		logLevel = slog.LevelDebug
	}

	logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))

	logger.InfoContext(ctx, "Downloading Unifi Network Server")
	firmwareClient, err := firmware.NewClient("")
	if err != nil {
		logger.ErrorContext(ctx, "Failed to create firmware client", "error", err)
		os.Exit(1)
	}

	archiveFile, version, err := firmwareClient.DownloadLatestVersion(ctx, logger)
	if err != nil {
		logger.ErrorContext(ctx, "Failed to download Unifi Network Server", "error", err)
		os.Exit(1)
	}

	defer func() {
		// Remove the downloaded archive file after we're done getting the API files out
		_ = os.Remove(archiveFile)
	}()

	logger.InfoContext(ctx, "Extracting API files")
	wd, err := os.Getwd()
	if err != nil {
		logger.ErrorContext(ctx, "Failed to get current working directory", "error", err)
		os.Exit(1)
	}

	apiDestination := filepath.Join(wd, "build", "api", version.Version.Core().String())
	if err = os.MkdirAll(apiDestination, 0o755); err != nil {
		logger.ErrorContext(ctx, "Failed to create API extraction directory", "error", err)
	}

	if err = firmware.ExtractAPI(ctx, logger, archiveFile, apiDestination); err != nil {
		logger.ErrorContext(ctx, "Failed to extract Unifi Network Server APIs", "error", err)
		os.Exit(1)
	}
}
