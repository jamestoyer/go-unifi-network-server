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

type UserGroup struct {
	ID     *string `json:"_id,omitempty"`
	SiteID *string `json:"site_id,omitempty"`

	Hidden   *bool   `json:"attr_hidden,omitempty"`
	HiddenID *string `json:"attr_hidden_id,omitempty"`
	NoDelete *bool   `json:"attr_no_delete,omitempty"`
	NoEdit   *bool   `json:"attr_no_edit,omitempty"`

	Name           *string `json:"name,omitempty"`
	QosRateMaxDown *int64  `json:"qos_rate_max_down,omitempty"`
	QosRateMaxUp   *int64  `json:"qos_rate_max_up,omitempty"`
}

func (s *UserGroup) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *UserGroup) GetQosRateMaxDown() int64 {
	if s == nil {
		return 0
	}

	return *s.QosRateMaxDown
}

func (s *UserGroup) GetQosRateMaxUp() int64 {
	if s == nil {
		return 0
	}

	return *s.QosRateMaxUp
}

type responseBodyUserGroup struct {
	Metadata json.RawMessage `json:"meta"`
	Payload  []UserGroup     `json:"data"`
}

func (c *Client) CreateUserGroup(ctx context.Context, site string, data *UserGroup) (*UserGroup, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "usergroup")
	req, err := c.NewRequest(ctx, http.MethodPost, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyUserGroup
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to create usergroup: %w`, err)
	}

	var item UserGroup
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to create UserGroup`)
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}

func (c *Client) DeleteUserGroup(ctx context.Context, site string, id string) (*http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "usergroup", id)
	req, err := c.NewRequest(ctx, http.MethodDelete, endpointPath, nil)
	if err != nil {
		return nil, err
	}

	var body responseBodyUserGroup
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return resp, fmt.Errorf(`unable to delete UserGroup: %w`, err)
	}

	return nil, nil
}

func (c *Client) GetUserGroup(ctx context.Context, site, id string) (*UserGroup, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "usergroup", id)
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyUserGroup
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get UserGroup: %w`, err)
	}

	var item UserGroup
	switch len(body.Payload) {
	case 0:
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}

func (c *Client) ListUserGroup(ctx context.Context, site string) ([]UserGroup, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "usergroup")
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyUserGroup
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get UserGroup: %w`, err)
	}

	return body.Payload, resp, nil
}

func (c *Client) UpdateUserGroup(ctx context.Context, site string, data *UserGroup) (*UserGroup, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "usergroup", *data.ID)
	req, err := c.NewRequest(ctx, http.MethodPut, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyUserGroup
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to update usergroup: %w`, err)
	}

	var item UserGroup
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to update UserGroup`)
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}
