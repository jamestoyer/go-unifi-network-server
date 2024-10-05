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
	"gopkg.in/yaml.v3"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/jamestoyer/go-unifi-network-server/internal/client-gen/api"
	"github.com/jamestoyer/go-unifi-network-server/internal/client-gen/firmware"
	"github.com/jamestoyer/go-unifi-network-server/internal/client-gen/spec"
	"github.com/jamestoyer/go-unifi-network-server/internal/logging"
)

const packageName = "networkserver"

var (
	verboseFlag    = flag.Bool("v", false, "Print verbose logs")
	apiSpecDirFlag = flag.String("dir", "", "Directory containing api spec files. If not set the latest API spec will be downloaded")
	configFileFlag = flag.String("config", filepath.Join("internal", "client-gen", "config.yml"), "Path to configuration file")

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

	config, err := loadConfig(*configFileFlag)
	if err != nil {
		logger.ErrorContext(ctx, "Failed to unmarshal configuration", slog.Any("error", err))
		os.Exit(1)
	}

	slog.InfoContext(ctx, "Parsing API specification directory")
	parser := api.NewParser(config.Parser)
	endpoints, err := parser.ParseDir(ctx, apiSpecDir)
	switch {
	case err != nil:
		slog.ErrorContext(ctx, "Failed to parse API specification directory", slog.Any("error", err))
		os.Exit(1)
	case len(endpoints) == 0:
		slog.InfoContext(ctx, "No endpoints found to generate")
		os.Exit(0)
	default:
		slog.InfoContext(ctx, "Found endpoints to render", slog.Int("endpointCount", len(endpoints)))
	}

	if err = renderAPIClient(ctx, endpoints); err != nil {
		logger.ErrorContext(ctx, "Failed to generate API client", slog.Any("error", err))
		os.Exit(1)
	}

	logger.InfoContext(ctx, "API client generated")
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

func renderAPIClient(ctx context.Context, endpoints []*spec.Endpoint) error {
	if err := os.MkdirAll(packageName, 0o775); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", packageName, err)
	}

	slog.InfoContext(ctx, "Removing previously generated files")
	if err := api.RemoveGeneratedFiles(ctx, packageName); err != nil {
		return err
	}

	renderer, err := api.NewRenderer()
	if err != nil {
		return err
	}

	slog.InfoContext(ctx, "Rendering API client")
	return renderer.RenderEndpoints(ctx, endpoints, packageName, packageName)
}

type Config struct {
	Parser api.ParserConfig `yaml:"parser"`
}

func loadConfig(file string) (*Config, error) {
	contents, err := os.ReadFile(*configFileFlag)
	if err != nil {
		return nil, fmt.Errorf("failed to read configuraiton file: %w", err)
	}

	var config *Config
	if err = yaml.Unmarshal(contents, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal yaml: %w", err)
	}

	return config, nil
}
