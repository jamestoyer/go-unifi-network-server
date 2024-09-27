// Code generated from Unifi Network Server API definitions.
// DO NOT EDIT.

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

package networkserver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"path"
)

type DpiApp struct {
	ID     *string `json:"_id,omitempty"`
	SiteID *string `json:"site_id,omitempty"`

	Hidden   *bool   `json:"attr_hidden,omitempty"`
	HiddenID *string `json:"attr_hidden_id,omitempty"`
	NoDelete *bool   `json:"attr_no_delete,omitempty"`
	NoEdit   *bool   `json:"attr_no_edit,omitempty"`

	Apps           *[]int64 `json:"apps,omitempty"`
	Blocked        *bool    `json:"blocked,omitempty"`
	Cats           *[]int64 `json:"cats,omitempty"`
	Enabled        *bool    `json:"enabled,omitempty"`
	Log            *bool    `json:"log,omitempty"`
	Name           *string  `json:"name,omitempty"`
	QosRateMaxDown *float64 `json:"qos_rate_max_down,omitempty"`
	QosRateMaxUp   *float64 `json:"qos_rate_max_up,omitempty"`
}

func (s *DpiApp) GetApps() []int64 {
	if s == nil || s.Apps == nil {
		return nil
	}

	return *s.Apps
}

func (s *DpiApp) GetBlocked() bool {
	if s == nil {
		return false
	}

	return *s.Blocked
}

func (s *DpiApp) GetCats() []int64 {
	if s == nil || s.Cats == nil {
		return nil
	}

	return *s.Cats
}

func (s *DpiApp) GetEnabled() bool {
	if s == nil {
		return false
	}

	return *s.Enabled
}

func (s *DpiApp) GetLog() bool {
	if s == nil {
		return false
	}

	return *s.Log
}

func (s *DpiApp) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *DpiApp) GetQosRateMaxDown() float64 {
	if s == nil {
		return 0
	}

	return *s.QosRateMaxDown
}

func (s *DpiApp) GetQosRateMaxUp() float64 {
	if s == nil {
		return 0
	}

	return *s.QosRateMaxUp
}

type responseBodyDpiApp struct {
	Metadata json.RawMessage `json:"meta"`
	Payload  []DpiApp        `json:"data"`
}

func (c *Client) CreateDpiApp(ctx context.Context, site string, data *DpiApp) (*DpiApp, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "dpiapp")
	req, err := c.NewRequest(ctx, http.MethodPost, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyDpiApp
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to create dpiapp: %w`, err)
	}

	var item DpiApp
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to create DpiApp`)
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}

func (c *Client) DeleteDpiApp(ctx context.Context, site string, id string) (*http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "dpiapp", id)
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, err
	}

	var body responseBodyDpiApp
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return resp, fmt.Errorf(`unable to delete DpiApp: %w`, err)
	}

	return nil, nil
}

func (c *Client) GetDpiApp(ctx context.Context, site, id string) (*DpiApp, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "dpiapp", id)
	req, err := c.NewRequest(ctx, http.MethodDelete, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyDpiApp
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get DpiApp: %w`, err)
	}

	var item DpiApp
	switch len(body.Payload) {
	case 0:
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}

func (c *Client) ListDpiApp(ctx context.Context, site string) ([]DpiApp, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "dpiapp")
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyDpiApp
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get DpiApp: %w`, err)
	}

	return body.Payload, resp, nil
}

func (c *Client) UpdateDpiApp(ctx context.Context, site string, data *DpiApp) (*DpiApp, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "dpiapp", *data.ID)
	req, err := c.NewRequest(ctx, http.MethodPut, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyDpiApp
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to update dpiapp: %w`, err)
	}

	var item DpiApp
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to update DpiApp`)
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}
