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

type DhcpOption struct {
	ID     *string `json:"_id,omitempty"`
	SiteID *string `json:"site_id,omitempty"`

	Hidden   *bool   `json:"attr_hidden,omitempty"`
	HiddenID *string `json:"attr_hidden_id,omitempty"`
	NoDelete *bool   `json:"attr_no_delete,omitempty"`
	NoEdit   *bool   `json:"attr_no_edit,omitempty"`

	Code   *string `json:"code,omitempty"`
	Name   *string `json:"name,omitempty"`
	Signed *bool   `json:"signed,omitempty"`
	Type   *string `json:"type,omitempty"`
	Width  *int64  `json:"width,omitempty"`
}

func (s *DhcpOption) GetCode() string {
	if s == nil {
		return ""
	}

	return *s.Code
}

func (s *DhcpOption) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *DhcpOption) GetSigned() bool {
	if s == nil {
		return false
	}

	return *s.Signed
}

func (s *DhcpOption) GetType() string {
	if s == nil {
		return ""
	}

	return *s.Type
}

func (s *DhcpOption) GetWidth() int64 {
	if s == nil {
		return 0
	}

	return *s.Width
}

type responseBodyDhcpOption struct {
	Metadata json.RawMessage `json:"meta"`
	Payload  []DhcpOption    `json:"data"`
}

func (c *Client) CreateDhcpOption(ctx context.Context, site string, data *DhcpOption) (*DhcpOption, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "dhcpoption")
	req, err := c.NewRequest(ctx, http.MethodPost, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyDhcpOption
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to create dhcpoption: %w`, err)
	}

	var item DhcpOption
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to create DhcpOption`)
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}

func (c *Client) DeleteDhcpOption(ctx context.Context, site string, id string) (*http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "dhcpoption", id)
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, err
	}

	var body responseBodyDhcpOption
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return resp, fmt.Errorf(`unable to delete DhcpOption: %w`, err)
	}

	return nil, nil
}

func (c *Client) GetDhcpOption(ctx context.Context, site, id string) (*DhcpOption, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "dhcpoption", id)
	req, err := c.NewRequest(ctx, http.MethodDelete, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyDhcpOption
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get DhcpOption: %w`, err)
	}

	var item DhcpOption
	switch len(body.Payload) {
	case 0:
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}

func (c *Client) ListDhcpOption(ctx context.Context, site string) ([]DhcpOption, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "dhcpoption")
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyDhcpOption
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get DhcpOption: %w`, err)
	}

	return body.Payload, resp, nil
}

func (c *Client) UpdateDhcpOption(ctx context.Context, site string, data *DhcpOption) (*DhcpOption, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "dhcpoption", *data.ID)
	req, err := c.NewRequest(ctx, http.MethodPut, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyDhcpOption
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to update dhcpoption: %w`, err)
	}

	var item DhcpOption
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to update DhcpOption`)
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}
