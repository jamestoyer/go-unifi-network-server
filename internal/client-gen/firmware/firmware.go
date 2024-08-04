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
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"

	"github.com/hashicorp/go-version"
)

const (
	DefaultEndpoint = "https://fw-update.ubnt.com/api/firmware-latest"

	filterChannelRelease         = "release"
	filterProductUnifiController = "unifi-controller"

	platformUnix = "unix"

	userAgent = "jamestoyer/go-unifi-network-server"
)

type Client struct {
	endpoint   string
	httpClient *http.Client
}

func NewClient(endpoint string) (*Client, error) {
	if endpoint == "" {
		endpoint = DefaultEndpoint
	}

	return &Client{
		endpoint:   endpoint,
		httpClient: &http.Client{},
	}, nil
}

func (c *Client) DownloadLatestVersion(ctx context.Context, logger *slog.Logger) (*os.File, error) {
	latestVersion, err := c.getLatestReleaseVersion(ctx, logger)
	if err != nil {
		return nil, err
	}

	if latestVersion == nil {
		return nil, errors.New("no version found")
	}

	logger.InfoContext(ctx, "Downloading latest version", "version", latestVersion.Version.String())
	downloadPath, err := c.downloadRelease(ctx, logger, latestVersion)
	if err != nil {
		// Remove the file and ignore any errors if that fails
		_ = os.Remove(downloadPath)
		return nil, err
	}

	fmt.Println(downloadPath)
	return nil, nil
}

func (c *Client) getLatestReleaseVersion(ctx context.Context, logger *slog.Logger) (*Version, error) {
	u, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid API endpoint URL: %w", err)
	}

	query := url.Values{}
	query.Add("filter", filterQuery("channel", filterChannelRelease))
	query.Add("filter", filterQuery("product", filterProductUnifiController))
	u.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", userAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("latest version request failed: %w", err)
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil && !errors.Is(err, io.EOF) {
			logger.ErrorContext(ctx, "Failed to close response body", "error", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch latest version metadata")
	}

	var data Response
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	for _, firmware := range data.Embedded.Firmware {
		if firmware.Platform == platformUnix {
			logger.DebugContext(ctx, "Found firmware", "platform", firmware.Platform, "version", firmware.Version)

			v, err := version.NewVersion(firmware.Version)
			if err != nil {
				return nil, fmt.Errorf("failed to parse firmware version: %w", err)
			}

			return &Version{
				SHA256Checksum: firmware.SHA256Checksum,
				Version:        v,
				URL:            firmware.Links.Data.Href,
			}, nil
		}
	}

	return nil, nil
}

func (c *Client) downloadRelease(ctx context.Context, logger *slog.Logger, version *Version) (string, error) {
	logger.DebugContext(ctx, "Creating temporary release file")
	f, err := os.CreateTemp("", "unifi-client-*.zip")
	if err != nil {
		return "", fmt.Errorf("failed to create temporary file: %w", err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			logger.ErrorContext(ctx, "Failed to close tempo", "error", err)
		}
	}()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, version.URL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	logger.DebugContext(ctx, "Downloading the version", "version", version.Version.String())
	req.Header.Set("User-Agent", userAgent)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("download release request failed: %w", err)
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil && !errors.Is(err, io.EOF) {
			logger.ErrorContext(ctx, "Failed to close response body", "error", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to download release")
	}

	if _, err := io.Copy(f, resp.Body); err != nil {
		return "", fmt.Errorf("failed to write release to file: %w", err)
	}

	if _, err = f.Seek(0, 0); err != nil {
		return "", fmt.Errorf("failed to seek to beginning of file: %w", err)
	}

	h := sha256.New()
	if _, err = io.Copy(h, f); err != nil {
		return "", fmt.Errorf("failed to copy file for verification: %w", err)
	}

	if hex.EncodeToString(h.Sum(nil)) != version.SHA256Checksum {
		return "", errors.New("failed to verify release signature")
	}

	return f.Name(), nil
}

func filterQuery(key, value string) string {
	return fmt.Sprintf("eq~~%s~~%s", key, value)
}

type Response struct {
	Embedded struct {
		Firmware []struct {
			Platform       string `json:"platform"`
			Version        string `json:"version"`
			SHA256Checksum string `json:"sha256_checksum"`
			Links          struct {
				Data struct {
					Href string `json:"href"`
				} `json:"data"`
			} `json:"_links"`
		} `json:"firmware"`
	} `json:"_embedded"`
}

type Version struct {
	SHA256Checksum string
	URL            string
	Version        *version.Version
}
