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

type User struct {
	ID                            *string `json:"_id,omitempty"`
	SiteID                        *string `json:"site_id,omitempty"`
	Hidden                        *bool   `json:"attr_hidden,omitempty"`
	HiddenID                      *string `json:"attr_hidden_id,omitempty"`
	NoDelete                      *bool   `json:"attr_no_delete,omitempty"`
	NoEdit                        *bool   `json:"attr_no_edit,omitempty"`
	Blocked                       *string `json:"blocked,omitempty"`
	FixedAPEnabled                *bool   `json:"fixed_ap_enabled,omitempty"`
	FixedAPMac                    *string `json:"fixed_ap_mac,omitempty"`
	FixedAPIP                     *string `json:"fixed_ip,omitempty"`
	Hostname                      *string `json:"hostname,omitempty"`
	LastSeen                      *string `json:"last_seen,omitempty"`
	LocalDNSRecord                *string `json:"local_dns_record,omitempty"`
	LocalDNSRecordEnabled         *bool   `json:"local_dns_record_enabled,omitempty"`
	MAC                           *string `json:"mac,omitempty"`
	Name                          *string `json:"name,omitempty"`
	NetworkID                     *string `json:"network_id,omitempty"`
	Note                          *string `json:"note,omitempty"`
	UseFixedIP                    *bool   `json:"use_fixedip,omitempty"`
	UserGroupID                   *string `json:"usergroup_id,omitempty"`
	VirtualNetworkOverrideEnabled *bool   `json:"virtual_network_override_enabled,omitempty"`
	VirtualNetworkOverrideID      *string `json:"virtual_network_override_id,omitempty"`
}

func (s *User) GetID() string {
	if s == nil {
		return ""
	}

	return *s.ID
}

func (s *User) GetSiteID() string {
	if s == nil {
		return ""
	}

	return *s.SiteID
}

func (s *User) GetHidden() bool {
	if s == nil {
		return false
	}

	return *s.Hidden
}

func (s *User) GetHiddenID() string {
	if s == nil {
		return ""
	}

	return *s.HiddenID
}

func (s *User) GetNoDelete() bool {
	if s == nil {
		return false
	}

	return *s.NoDelete
}

func (s *User) GetNoEdit() bool {
	if s == nil {
		return false
	}

	return *s.NoEdit
}

func (s *User) GetBlocked() string {
	if s == nil {
		return ""
	}

	return *s.Blocked
}

func (s *User) GetFixedAPEnabled() bool {
	if s == nil {
		return false
	}

	return *s.FixedAPEnabled
}

func (s *User) GetFixedAPMac() string {
	if s == nil {
		return ""
	}

	return *s.FixedAPMac
}

func (s *User) GetFixedAPIP() string {
	if s == nil {
		return ""
	}

	return *s.FixedAPIP
}

func (s *User) GetHostname() string {
	if s == nil {
		return ""
	}

	return *s.Hostname
}

func (s *User) GetLastSeen() string {
	if s == nil {
		return ""
	}

	return *s.LastSeen
}

func (s *User) GetLocalDNSRecord() string {
	if s == nil {
		return ""
	}

	return *s.LocalDNSRecord
}

func (s *User) GetLocalDNSRecordEnabled() bool {
	if s == nil {
		return false
	}

	return *s.LocalDNSRecordEnabled
}

func (s *User) GetMAC() string {
	if s == nil {
		return ""
	}

	return *s.MAC
}

func (s *User) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *User) GetNetworkID() string {
	if s == nil {
		return ""
	}

	return *s.NetworkID
}

func (s *User) GetNote() string {
	if s == nil {
		return ""
	}

	return *s.Note
}

func (s *User) GetUseFixedIP() bool {
	if s == nil {
		return false
	}

	return *s.UseFixedIP
}

func (s *User) GetUserGroupID() string {
	if s == nil {
		return ""
	}

	return *s.UserGroupID
}

func (s *User) GetVirtualNetworkOverrideEnabled() bool {
	if s == nil {
		return false
	}

	return *s.VirtualNetworkOverrideEnabled
}

func (s *User) GetVirtualNetworkOverrideID() string {
	if s == nil {
		return ""
	}

	return *s.VirtualNetworkOverrideID
}

type responseBodyUser struct {
	Metadata json.RawMessage `json:"meta"`
	Payload  []User          `json:"data"`
}

func (c *Client) CreateUser(ctx context.Context, site string, data *User) (*User, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "user")
	req, err := c.NewRequest(ctx, http.MethodPost, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyUser
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to create user: %w`, err)
	}

	var item User
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to create User`)
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}

func (c *Client) DeleteUser(ctx context.Context, site string, id string) (*http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "user", id)
	req, err := c.NewRequest(ctx, http.MethodDelete, endpointPath, nil)
	if err != nil {
		return nil, err
	}

	var body responseBodyUser
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return resp, fmt.Errorf(`unable to delete User: %w`, err)
	}

	return resp, nil
}

func (c *Client) GetUser(ctx context.Context, site, id string) (*User, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "user", id)
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyUser
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get User: %w`, err)
	}

	var item User
	switch len(body.Payload) {
	case 0:
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}

func (c *Client) ListUser(ctx context.Context, site string) ([]User, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "user")
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyUser
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get User: %w`, err)
	}

	return body.Payload, resp, nil
}

func (c *Client) UpdateUser(ctx context.Context, site string, data *User) (*User, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "user", data.GetID())
	req, err := c.NewRequest(ctx, http.MethodPut, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyUser
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to update user: %w`, err)
	}

	var item User
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to update User`)
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}
