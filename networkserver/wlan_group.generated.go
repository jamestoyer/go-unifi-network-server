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

type WlanGroup struct {
	ID     *string `json:"_id,omitempty"`
	SiteID *string `json:"site_id,omitempty"`

	Hidden   *bool   `json:"attr_hidden,omitempty"`
	HiddenID *string `json:"attr_hidden_id,omitempty"`
	NoDelete *bool   `json:"attr_no_delete,omitempty"`
	NoEdit   *bool   `json:"attr_no_edit,omitempty"`

	Name *string `json:"name,omitempty"`
}

func (s *WlanGroup) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

type responseBodyWlanGroup struct {
	Metadata json.RawMessage `json:"meta"`
	Payload  []WlanGroup     `json:"data"`
}

func (c *Client) CreateWlanGroup(ctx context.Context, site string, data *WlanGroup) (*WlanGroup, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "wlangroup")
	req, err := c.NewRequest(ctx, http.MethodPost, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyWlanGroup
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to create wlangroup: %w`, err)
	}

	var item WlanGroup
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to create WlanGroup`)
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}

func (c *Client) DeleteWlanGroup(ctx context.Context, site string, id string) (*http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "wlangroup", id)
	req, err := c.NewRequest(ctx, http.MethodDelete, endpointPath, nil)
	if err != nil {
		return nil, err
	}

	var body responseBodyWlanGroup
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return resp, fmt.Errorf(`unable to delete WlanGroup: %w`, err)
	}

	return nil, nil
}

func (c *Client) GetWlanGroup(ctx context.Context, site, id string) (*WlanGroup, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "wlangroup", id)
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyWlanGroup
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get WlanGroup: %w`, err)
	}

	var item WlanGroup
	switch len(body.Payload) {
	case 0:
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}

func (c *Client) ListWlanGroup(ctx context.Context, site string) ([]WlanGroup, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "wlangroup")
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyWlanGroup
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get WlanGroup: %w`, err)
	}

	return body.Payload, resp, nil
}

func (c *Client) UpdateWlanGroup(ctx context.Context, site string, data *WlanGroup) (*WlanGroup, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "wlangroup", *data.ID)
	req, err := c.NewRequest(ctx, http.MethodPut, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyWlanGroup
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to update wlangroup: %w`, err)
	}

	var item WlanGroup
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to update WlanGroup`)
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}
