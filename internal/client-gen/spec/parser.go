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

package spec

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

	"github.com/jamestoyer/go-unifi-network-server/internal/client-gen/api"
	"github.com/jamestoyer/go-unifi-network-server/internal/logging"
)

const (
	parserLogGroup = "parser"
)

type ParserConfig struct {
	Endpoints    map[string]ParserEndpointConfig `yaml:"endpoints"`
	SkippedFiles []string                        `yaml:"skippedFiles"`
}

func (p *ParserConfig) GetEndpoint(name string) ParserEndpointConfig {
	if p == nil || p.Endpoints == nil {
		return ParserEndpointConfig{}
	}

	endpointConfig, exists := p.Endpoints[name]
	if !exists {
		return ParserEndpointConfig{}
	}

	return endpointConfig
}

type ParserEndpointConfig struct {
	Skip bool `yaml:"skip"`
}

type Parser struct {
	config ParserConfig
}

func NewParser(config ParserConfig) *Parser {
	return &Parser{config: config}
}

func (p *Parser) ParseDir(ctx context.Context, dir string) ([]*api.Endpoint, error) {
	ctx = logging.CtxWithValues(ctx, slog.Group(parserLogGroup, slog.String("specDir", dir)))
	var endpoints []*api.Endpoint
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() && path != dir {
			slog.DebugContext(ctx, "Unexpected directory skipped", slog.Group(parserLogGroup, slog.String("currentPath", dir)))
			return filepath.SkipDir
		}

		if filepath.Ext(path) != ".json" {
			slog.DebugContext(ctx, "Unsupported file extension skipped", slog.Group(parserLogGroup, slog.String("currentPath", dir)))
			return nil
		}

		if slices.Contains(p.config.SkippedFiles, filepath.Base(path)) {
			slog.DebugContext(ctx, "File explicitly skipped in configuration", slog.Group(parserLogGroup, slog.String("currentPath", dir)))
			return nil
		}

		eps, err := p.ParseFile(ctx, path)
		if err != nil {
			return err
		}

		endpoints = append(endpoints, eps...)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse files in directory: %w", err)
	}

	// Sort the endpoints so that they are always in the same order.
	slices.SortFunc(endpoints, func(a, b *api.Endpoint) int {
		return strings.Compare(a.Name, b.Name)
	})

	return endpoints, nil
}

func (p *Parser) ParseFile(ctx context.Context, filename string) ([]*api.Endpoint, error) {
	ctx = logging.CtxWithValues(ctx, slog.Group(parserLogGroup, slog.String("filename", filename)))
	slog.DebugContext(ctx, "Parsing specification file for endpoints")

	if filepath.Ext(filename) != ".json" {
		return nil, errors.New("unsupported specification file type")
	}

	_, err := os.Stat(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, errors.New("file does not exist")
		}

		return nil, fmt.Errorf("invalid file: %w", err)
	}

	// TODO: (jtoyer) Add support for multiple endpoints in one file, e.g. settings.json
	contents, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read specification file: %w", err)
	}

	var fields map[string]interface{}
	if err = json.Unmarshal(contents, &fields); err != nil {
		return nil, fmt.Errorf("failed to unmarshal specification file: %w", err)
	}

	endpointName := strings.TrimSuffix(filepath.Base(filename), filepath.Ext(filename))
	ctx = logging.CtxWithValues(ctx, slog.Group(parserLogGroup, slog.String("endpoint", endpointName)))
	endpointConfig := p.config.GetEndpoint(endpointName)
	if endpointConfig.Skip {
		slog.InfoContext(ctx, "Endpoint explicitly skipped in configuration")
		return nil, nil
	}

	endpoint, err := api.NewEndpoint(endpointName, filename, fields)
	if err != nil {
		return nil, fmt.Errorf("failed to parse endpoint: %w", err)
	}

	return []*api.Endpoint{endpoint}, nil
}
