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

type FirewallGroup struct {
	ID           *string   `json:"_id,omitempty"`
	SiteID       *string   `json:"site_id,omitempty"`
	Hidden       *bool     `json:"attr_hidden,omitempty"`
	HiddenID     *string   `json:"attr_hidden_id,omitempty"`
	NoDelete     *bool     `json:"attr_no_delete,omitempty"`
	NoEdit       *bool     `json:"attr_no_edit,omitempty"`
	GroupMembers *[]string `json:"group_members,omitempty"`
	GroupType    *string   `json:"group_type,omitempty"`
	Name         *string   `json:"name,omitempty"`
}

func (s *FirewallGroup) GetID() string {
	if s == nil {
		return ""
	}

	return *s.ID
}

func (s *FirewallGroup) GetSiteID() string {
	if s == nil {
		return ""
	}

	return *s.SiteID
}

func (s *FirewallGroup) GetHidden() bool {
	if s == nil {
		return false
	}

	return *s.Hidden
}

func (s *FirewallGroup) GetHiddenID() string {
	if s == nil {
		return ""
	}

	return *s.HiddenID
}

func (s *FirewallGroup) GetNoDelete() bool {
	if s == nil {
		return false
	}

	return *s.NoDelete
}

func (s *FirewallGroup) GetNoEdit() bool {
	if s == nil {
		return false
	}

	return *s.NoEdit
}

func (s *FirewallGroup) GetGroupMembers() []string {
	if s == nil || s.GroupMembers == nil {
		return nil
	}

	return *s.GroupMembers
}

func (s *FirewallGroup) GetGroupType() string {
	if s == nil {
		return ""
	}

	return *s.GroupType
}

func (s *FirewallGroup) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

type responseBodyFirewallGroup struct {
	Metadata json.RawMessage `json:"meta"`
	Payload  []FirewallGroup `json:"data"`
}

func (c *Client) CreateFirewallGroup(ctx context.Context, site string, data *FirewallGroup) (*FirewallGroup, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "firewallgroup")
	req, err := c.NewRequest(ctx, http.MethodPost, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyFirewallGroup
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to create firewallgroup: %w`, err)
	}

	var item FirewallGroup
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to create FirewallGroup`)
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}

func (c *Client) DeleteFirewallGroup(ctx context.Context, site string, id string) (*http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "firewallgroup", id)
	req, err := c.NewRequest(ctx, http.MethodDelete, endpointPath, nil)
	if err != nil {
		return nil, err
	}

	var body responseBodyFirewallGroup
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return resp, fmt.Errorf(`unable to delete FirewallGroup: %w`, err)
	}

	return resp, nil
}

func (c *Client) GetFirewallGroup(ctx context.Context, site, id string) (*FirewallGroup, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "firewallgroup", id)
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyFirewallGroup
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get FirewallGroup: %w`, err)
	}

	var item FirewallGroup
	switch len(body.Payload) {
	case 0:
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}

func (c *Client) ListFirewallGroup(ctx context.Context, site string) ([]FirewallGroup, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "firewallgroup")
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyFirewallGroup
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get FirewallGroup: %w`, err)
	}

	return body.Payload, resp, nil
}

func (c *Client) UpdateFirewallGroup(ctx context.Context, site string, data *FirewallGroup) (*FirewallGroup, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "firewallgroup", data.GetID())
	req, err := c.NewRequest(ctx, http.MethodPut, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyFirewallGroup
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to update firewallgroup: %w`, err)
	}

	var item FirewallGroup
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to update FirewallGroup`)
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}
