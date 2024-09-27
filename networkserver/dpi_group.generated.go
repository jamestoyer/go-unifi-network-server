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

type DpiGroup struct {
	ID     *string `json:"_id,omitempty"`
	SiteID *string `json:"site_id,omitempty"`

	Hidden   *bool   `json:"attr_hidden,omitempty"`
	HiddenID *string `json:"attr_hidden_id,omitempty"`
	NoDelete *bool   `json:"attr_no_delete,omitempty"`
	NoEdit   *bool   `json:"attr_no_edit,omitempty"`

	DpiappIds *[]string `json:"dpiapp_ids,omitempty"`
	Enabled   *bool     `json:"enabled,omitempty"`
	Name      *string   `json:"name,omitempty"`
}

func (s *DpiGroup) GetDpiappIds() []string {
	if s == nil || s.DpiappIds == nil {
		return nil
	}

	return *s.DpiappIds
}

func (s *DpiGroup) GetEnabled() bool {
	if s == nil {
		return false
	}

	return *s.Enabled
}

func (s *DpiGroup) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

type responseBodyDpiGroup struct {
	Metadata json.RawMessage `json:"meta"`
	Payload  []DpiGroup      `json:"data"`
}

func (c *Client) CreateDpiGroup(ctx context.Context, site string, data *DpiGroup) (*DpiGroup, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "dpigroup")
	req, err := c.NewRequest(ctx, http.MethodPost, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyDpiGroup
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to create dpigroup: %w`, err)
	}

	var item DpiGroup
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to create DpiGroup`)
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}

func (c *Client) DeleteDpiGroup(ctx context.Context, site string, id string) (*http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "dpigroup", id)
	req, err := c.NewRequest(ctx, http.MethodDelete, endpointPath, nil)
	if err != nil {
		return nil, err
	}

	var body responseBodyDpiGroup
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return resp, fmt.Errorf(`unable to delete DpiGroup: %w`, err)
	}

	return nil, nil
}

func (c *Client) GetDpiGroup(ctx context.Context, site, id string) (*DpiGroup, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "dpigroup", id)
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyDpiGroup
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get DpiGroup: %w`, err)
	}

	var item DpiGroup
	switch len(body.Payload) {
	case 0:
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}

func (c *Client) ListDpiGroup(ctx context.Context, site string) ([]DpiGroup, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "dpigroup")
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyDpiGroup
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get DpiGroup: %w`, err)
	}

	return body.Payload, resp, nil
}

func (c *Client) UpdateDpiGroup(ctx context.Context, site string, data *DpiGroup) (*DpiGroup, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "dpigroup", *data.ID)
	req, err := c.NewRequest(ctx, http.MethodPut, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyDpiGroup
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to update dpigroup: %w`, err)
	}

	var item DpiGroup
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to update DpiGroup`)
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}
