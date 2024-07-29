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
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-version"
)

const (
	DefaultEndpoint = "https://fw-update.ubnt.com/api/firmware-latest"

	filterChannelRelease         = "release"
	filterProductUnifiController = "unifi-controller"

	platformUnix = "unix"
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

func (c *Client) LatestVersion(ctx context.Context, logger *slog.Logger) (io.ReadWriteCloser, error) {
	versionURL, err := c.getLatestReleaseVersion(ctx, logger)
	if err != nil {
		return nil, err
	}

	if versionURL == nil {
		return nil, errors.New("no version found")
	}

	fmt.Println(versionURL)
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

	req.Header.Set("User-Agent", "jamestoyer/go-unifi-network-server")

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

func (c *Client) downloadRelease(ctx context.Context, logger *slog.Logger, version *Version) error {

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
