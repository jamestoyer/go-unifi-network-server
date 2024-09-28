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

type DynamicDns struct {
	ID            *string   `json:"_id,omitempty"`
	SiteID        *string   `json:"site_id,omitempty"`
	Hidden        *bool     `json:"attr_hidden,omitempty"`
	HiddenID      *string   `json:"attr_hidden_id,omitempty"`
	NoDelete      *bool     `json:"attr_no_delete,omitempty"`
	NoEdit        *bool     `json:"attr_no_edit,omitempty"`
	CustomService *string   `json:"custom_service,omitempty"`
	HostName      *string   `json:"host_name,omitempty"`
	Interface     *string   `json:"interface,omitempty"`
	Login         *string   `json:"login,omitempty"`
	Options       *[]string `json:"options,omitempty"`
	Server        *string   `json:"server,omitempty"`
	Service       *string   `json:"service,omitempty"`
	XPassword     *string   `json:"x_password,omitempty"`
}

func (s *DynamicDns) GetID() string {
	if s == nil {
		return ""
	}

	return *s.ID
}

func (s *DynamicDns) GetSiteID() string {
	if s == nil {
		return ""
	}

	return *s.SiteID
}

func (s *DynamicDns) GetHidden() bool {
	if s == nil {
		return false
	}

	return *s.Hidden
}

func (s *DynamicDns) GetHiddenID() string {
	if s == nil {
		return ""
	}

	return *s.HiddenID
}

func (s *DynamicDns) GetNoDelete() bool {
	if s == nil {
		return false
	}

	return *s.NoDelete
}

func (s *DynamicDns) GetNoEdit() bool {
	if s == nil {
		return false
	}

	return *s.NoEdit
}

func (s *DynamicDns) GetCustomService() string {
	if s == nil {
		return ""
	}

	return *s.CustomService
}

func (s *DynamicDns) GetHostName() string {
	if s == nil {
		return ""
	}

	return *s.HostName
}

func (s *DynamicDns) GetInterface() string {
	if s == nil {
		return ""
	}

	return *s.Interface
}

func (s *DynamicDns) GetLogin() string {
	if s == nil {
		return ""
	}

	return *s.Login
}

func (s *DynamicDns) GetOptions() []string {
	if s == nil || s.Options == nil {
		return nil
	}

	return *s.Options
}

func (s *DynamicDns) GetServer() string {
	if s == nil {
		return ""
	}

	return *s.Server
}

func (s *DynamicDns) GetService() string {
	if s == nil {
		return ""
	}

	return *s.Service
}

func (s *DynamicDns) GetXPassword() string {
	if s == nil {
		return ""
	}

	return *s.XPassword
}

type responseBodyDynamicDns struct {
	Metadata json.RawMessage `json:"meta"`
	Payload  []DynamicDns    `json:"data"`
}

func (c *Client) CreateDynamicDns(ctx context.Context, site string, data *DynamicDns) (*DynamicDns, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "dynamicdns")
	req, err := c.NewRequest(ctx, http.MethodPost, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyDynamicDns
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to create dynamicdns: %w`, err)
	}

	var item DynamicDns
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to create DynamicDns`)
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}

func (c *Client) DeleteDynamicDns(ctx context.Context, site string, id string) (*http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "dynamicdns", id)
	req, err := c.NewRequest(ctx, http.MethodDelete, endpointPath, nil)
	if err != nil {
		return nil, err
	}

	var body responseBodyDynamicDns
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return resp, fmt.Errorf(`unable to delete DynamicDns: %w`, err)
	}

	return resp, nil
}

func (c *Client) GetDynamicDns(ctx context.Context, site, id string) (*DynamicDns, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "dynamicdns", id)
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyDynamicDns
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get DynamicDns: %w`, err)
	}

	var item DynamicDns
	switch len(body.Payload) {
	case 0:
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}

func (c *Client) ListDynamicDns(ctx context.Context, site string) ([]DynamicDns, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "dynamicdns")
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyDynamicDns
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get DynamicDns: %w`, err)
	}

	return body.Payload, resp, nil
}

func (c *Client) UpdateDynamicDns(ctx context.Context, site string, data *DynamicDns) (*DynamicDns, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "dynamicdns", data.GetID())
	req, err := c.NewRequest(ctx, http.MethodPut, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyDynamicDns
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to update dynamicdns: %w`, err)
	}

	var item DynamicDns
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to update DynamicDns`)
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}
