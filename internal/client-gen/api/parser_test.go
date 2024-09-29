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
	"log/slog"
	"os"
	"path/filepath"
	"testing"

	"github.com/jamestoyer/go-unifi-network-server/internal/logging"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseDir(t *testing.T) {
	logger := slog.New(logging.Handler{Handler: slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})})
	slog.SetDefault(logger)

	tests := map[string]struct {
		config ParserConfig
		dir    string

		want    []*Endpoint
		wantErr assert.ErrorAssertionFunc
	}{
		"no files in dir": {
			dir:     filepath.Join("testdata", "empty-dir"),
			want:    nil,
			wantErr: assert.NoError,
		},
		"files in dir": {
			dir: filepath.Join("testdata", "valid-specs"),
			want: []*Endpoint{
				accountEndpoint(t, filepath.Join("testdata", "valid-specs", "account.json")),
				deviceEndpoint(t, filepath.Join("testdata", "valid-specs", "Device.json")),
				userEndpoint(t, filepath.Join("testdata", "valid-specs", "User.json")),
			},
			wantErr: assert.NoError,
		},
		"skipped files": {
			config: ParserConfig{
				SkippedFiles: []string{
					"account.json",
					"Device.json",
				},
			},
			dir: filepath.Join("testdata", "valid-specs"),
			want: []*Endpoint{
				userEndpoint(t, filepath.Join("testdata", "valid-specs", "User.json")),
			},
			wantErr: assert.NoError,
		},
		"skipped endpoints": {
			config: ParserConfig{
				Endpoints: map[string]ParserEndpointConfig{
					"Device": {Skip: true},
				},
			},
			dir: filepath.Join("testdata", "valid-specs"),
			want: []*Endpoint{
				accountEndpoint(t, filepath.Join("testdata", "valid-specs", "account.json")),
				userEndpoint(t, filepath.Join("testdata", "valid-specs", "User.json")),
			},
			wantErr: assert.NoError,
		},
		"invalid dir": {
			dir:     "dir-does-not-exist",
			want:    nil,
			wantErr: assert.Error,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			parser := NewParser(test.config)
			got, err := parser.ParseDir(ctx, test.dir)

			test.wantErr(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestParser_ParseFile(t *testing.T) {
	tests := map[string]struct {
		config   ParserConfig
		filename string
		want     []*Endpoint
		wantErr  assert.ErrorAssertionFunc
	}{
		"no file found": {
			filename: "invalid-file.json",
			wantErr:  assert.Error,
		},
		"unsupported file type": {
			filename: filepath.Join("testdata", "valid-specs", "spec.yml"),
			wantErr:  assert.Error,
		},
		"path is directory": {
			filename: filepath.Join("testdata", "valid-specs"),
			wantErr:  assert.Error,
		},
		"valid spec file": {
			filename: filepath.Join("testdata", "valid-specs", "account.json"),
			want: []*Endpoint{
				accountEndpoint(t, filepath.Join("testdata", "valid-specs", "account.json")),
			},
			wantErr: assert.NoError,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			parser := NewParser(test.config)
			got, err := parser.ParseFile(ctx, test.filename)
			test.wantErr(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func accountEndpoint(t *testing.T, filename string) *Endpoint {
	t.Helper()
	endpoint, err := NewEndpoint("Account", filename, map[string]interface{}{
		"name": "^[^\"' ]+$",
	})
	require.NoError(t, err)
	return endpoint
}

func deviceEndpoint(t *testing.T, filename string) *Endpoint {
	t.Helper()
	endpoint, err := NewEndpoint("Device", filename, map[string]interface{}{
		"name": ".{0,128}",
	})
	require.NoError(t, err)
	return endpoint
}

func userEndpoint(t *testing.T, filename string) *Endpoint {
	t.Helper()
	endpoint, err := NewEndpoint("User", filename, map[string]interface{}{
		"mac": "^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$",
	})
	require.NoError(t, err)
	return endpoint
}
