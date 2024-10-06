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
	// ID has been auto generated from the Unifi Network Server API specification
	//
	// Validation: None
	ID *string `json:"_id,omitempty"`
	// SiteID has been auto generated from the Unifi Network Server API specification
	//
	// Validation: None
	SiteID *string `json:"site_id,omitempty"`
	// Hidden has been auto generated from the Unifi Network Server API specification
	//
	// Validation: None
	Hidden *bool `json:"attr_hidden,omitempty"`
	// HiddenID has been auto generated from the Unifi Network Server API specification
	//
	// Validation: None
	HiddenID *string `json:"attr_hidden_id,omitempty"`
	// NoDelete has been auto generated from the Unifi Network Server API specification
	//
	// Validation: None
	NoDelete *bool `json:"attr_no_delete,omitempty"`
	// NoEdit has been auto generated from the Unifi Network Server API specification
	//
	// Validation: None
	NoEdit *bool `json:"attr_no_edit,omitempty"`
	// Blocked has been auto generated from the Unifi Network Server API specification
	//
	// Validation: None
	Blocked *string `json:"blocked,omitempty"`
	// FixedAPEnabled has been auto generated from the Unifi Network Server API specification
	//
	// Validation: false|true
	FixedAPEnabled *bool `json:"fixed_ap_enabled,omitempty"`
	// FixedAPMAC has been auto generated from the Unifi Network Server API specification
	//
	// Validation: ^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$
	FixedAPMAC *string `json:"fixed_ap_mac,omitempty"`
	// FixedAPIP has been auto generated from the Unifi Network Server API specification
	//
	// Validation: None
	FixedAPIP *string `json:"fixed_ip,omitempty"`
	// Hostname has been auto generated from the Unifi Network Server API specification
	//
	// Validation: None
	Hostname *string `json:"hostname,omitempty"`
	// LastSeen has been auto generated from the Unifi Network Server API specification
	//
	// Validation: None
	LastSeen *string `json:"last_seen,omitempty"`
	// LocalDNSRecord has been auto generated from the Unifi Network Server API specification
	//
	// Validation: None
	LocalDNSRecord *string `json:"local_dns_record,omitempty"`
	// LocalDNSRecordEnabled has been auto generated from the Unifi Network Server API specification
	//
	// Validation: false|true
	LocalDNSRecordEnabled *bool `json:"local_dns_record_enabled,omitempty"`
	// MAC has been auto generated from the Unifi Network Server API specification
	//
	// Validation: ^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$
	MAC *string `json:"mac,omitempty"`
	// Name has been auto generated from the Unifi Network Server API specification
	//
	// Validation: None
	Name *string `json:"name,omitempty"`
	// NetworkID has been auto generated from the Unifi Network Server API specification
	//
	// Validation: None
	NetworkID *string `json:"network_id,omitempty"`
	// Note has been auto generated from the Unifi Network Server API specification
	//
	// Validation: None
	Note *string `json:"note,omitempty"`
	// UseFixedIP has been auto generated from the Unifi Network Server API specification
	//
	// Validation: false|true
	UseFixedIP *bool `json:"use_fixedip,omitempty"`
	// UserGroupID has been auto generated from the Unifi Network Server API specification
	//
	// Validation: None
	UserGroupID *string `json:"usergroup_id,omitempty"`
	// VirtualNetworkOverrideEnabled has been auto generated from the Unifi Network Server API specification
	//
	// Validation: false|true
	VirtualNetworkOverrideEnabled *bool `json:"virtual_network_override_enabled,omitempty"`
	// VirtualNetworkOverrideID has been auto generated from the Unifi Network Server API specification
	//
	// Validation: None
	VirtualNetworkOverrideID *string `json:"virtual_network_override_id,omitempty"`
}

// GetID is a helper function which dereferences ID.
//
// When ID is a nil pointer it will return `""` as default.
func (s *User) GetID() string {
	if s == nil || s.ID == nil {
		return ""
	}

	return *s.ID
}

// GetSiteID is a helper function which dereferences SiteID.
//
// When SiteID is a nil pointer it will return `""` as default.
func (s *User) GetSiteID() string {
	if s == nil || s.SiteID == nil {
		return ""
	}

	return *s.SiteID
}

// GetHidden is a helper function which dereferences Hidden.
//
// When Hidden is a nil pointer it will return `false` as default.
func (s *User) GetHidden() bool {
	if s == nil || s.Hidden == nil {
		return false
	}

	return *s.Hidden
}

// GetHiddenID is a helper function which dereferences HiddenID.
//
// When HiddenID is a nil pointer it will return `""` as default.
func (s *User) GetHiddenID() string {
	if s == nil || s.HiddenID == nil {
		return ""
	}

	return *s.HiddenID
}

// GetNoDelete is a helper function which dereferences NoDelete.
//
// When NoDelete is a nil pointer it will return `false` as default.
func (s *User) GetNoDelete() bool {
	if s == nil || s.NoDelete == nil {
		return false
	}

	return *s.NoDelete
}

// GetNoEdit is a helper function which dereferences NoEdit.
//
// When NoEdit is a nil pointer it will return `false` as default.
func (s *User) GetNoEdit() bool {
	if s == nil || s.NoEdit == nil {
		return false
	}

	return *s.NoEdit
}

// GetBlocked is a helper function which dereferences Blocked.
//
// When Blocked is a nil pointer it will return `""` as default.
func (s *User) GetBlocked() string {
	if s == nil || s.Blocked == nil {
		return ""
	}

	return *s.Blocked
}

// GetFixedAPEnabled is a helper function which dereferences FixedAPEnabled.
//
// When FixedAPEnabled is a nil pointer it will return `false` as default.
func (s *User) GetFixedAPEnabled() bool {
	if s == nil || s.FixedAPEnabled == nil {
		return false
	}

	return *s.FixedAPEnabled
}

// GetFixedAPMAC is a helper function which dereferences FixedAPMAC.
//
// When FixedAPMAC is a nil pointer it will return `""` as default.
func (s *User) GetFixedAPMAC() string {
	if s == nil || s.FixedAPMAC == nil {
		return ""
	}

	return *s.FixedAPMAC
}

// GetFixedAPIP is a helper function which dereferences FixedAPIP.
//
// When FixedAPIP is a nil pointer it will return `""` as default.
func (s *User) GetFixedAPIP() string {
	if s == nil || s.FixedAPIP == nil {
		return ""
	}

	return *s.FixedAPIP
}

// GetHostname is a helper function which dereferences Hostname.
//
// When Hostname is a nil pointer it will return `""` as default.
func (s *User) GetHostname() string {
	if s == nil || s.Hostname == nil {
		return ""
	}

	return *s.Hostname
}

// GetLastSeen is a helper function which dereferences LastSeen.
//
// When LastSeen is a nil pointer it will return `""` as default.
func (s *User) GetLastSeen() string {
	if s == nil || s.LastSeen == nil {
		return ""
	}

	return *s.LastSeen
}

// GetLocalDNSRecord is a helper function which dereferences LocalDNSRecord.
//
// When LocalDNSRecord is a nil pointer it will return `""` as default.
func (s *User) GetLocalDNSRecord() string {
	if s == nil || s.LocalDNSRecord == nil {
		return ""
	}

	return *s.LocalDNSRecord
}

// GetLocalDNSRecordEnabled is a helper function which dereferences LocalDNSRecordEnabled.
//
// When LocalDNSRecordEnabled is a nil pointer it will return `false` as default.
func (s *User) GetLocalDNSRecordEnabled() bool {
	if s == nil || s.LocalDNSRecordEnabled == nil {
		return false
	}

	return *s.LocalDNSRecordEnabled
}

// GetMAC is a helper function which dereferences MAC.
//
// When MAC is a nil pointer it will return `""` as default.
func (s *User) GetMAC() string {
	if s == nil || s.MAC == nil {
		return ""
	}

	return *s.MAC
}

// GetName is a helper function which dereferences Name.
//
// When Name is a nil pointer it will return `""` as default.
func (s *User) GetName() string {
	if s == nil || s.Name == nil {
		return ""
	}

	return *s.Name
}

// GetNetworkID is a helper function which dereferences NetworkID.
//
// When NetworkID is a nil pointer it will return `""` as default.
func (s *User) GetNetworkID() string {
	if s == nil || s.NetworkID == nil {
		return ""
	}

	return *s.NetworkID
}

// GetNote is a helper function which dereferences Note.
//
// When Note is a nil pointer it will return `""` as default.
func (s *User) GetNote() string {
	if s == nil || s.Note == nil {
		return ""
	}

	return *s.Note
}

// GetUseFixedIP is a helper function which dereferences UseFixedIP.
//
// When UseFixedIP is a nil pointer it will return `false` as default.
func (s *User) GetUseFixedIP() bool {
	if s == nil || s.UseFixedIP == nil {
		return false
	}

	return *s.UseFixedIP
}

// GetUserGroupID is a helper function which dereferences UserGroupID.
//
// When UserGroupID is a nil pointer it will return `""` as default.
func (s *User) GetUserGroupID() string {
	if s == nil || s.UserGroupID == nil {
		return ""
	}

	return *s.UserGroupID
}

// GetVirtualNetworkOverrideEnabled is a helper function which dereferences VirtualNetworkOverrideEnabled.
//
// When VirtualNetworkOverrideEnabled is a nil pointer it will return `false` as default.
func (s *User) GetVirtualNetworkOverrideEnabled() bool {
	if s == nil || s.VirtualNetworkOverrideEnabled == nil {
		return false
	}

	return *s.VirtualNetworkOverrideEnabled
}

// GetVirtualNetworkOverrideID is a helper function which dereferences VirtualNetworkOverrideID.
//
// When VirtualNetworkOverrideID is a nil pointer it will return `""` as default.
func (s *User) GetVirtualNetworkOverrideID() string {
	if s == nil || s.VirtualNetworkOverrideID == nil {
		return ""
	}

	return *s.VirtualNetworkOverrideID
}

type responseBodyUser struct {
	Metadata json.RawMessage `json:"meta"`
	Payload  []User          `json:"data"`
}

func (c *Client) CreateUser(ctx context.Context, data *User) (*User, *http.Response, error) {
	req, err := c.NewRequest(ctx, http.MethodPost, c.ResourceAPIPath("user"), data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyUser
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to create User: %w`, err)
	}

	var item *User
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to create User`)
	case 1:
		item = &body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return item, resp, err
}

func (c *Client) DeleteUser(ctx context.Context, id string) (*http.Response, error) {
	endpointPath := path.Join(c.ResourceAPIPath("user"), id)
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

func (c *Client) GetUser(ctx context.Context, id string) (*User, *http.Response, error) {
	endpointPath := path.Join(c.ResourceAPIPath("user"), id)
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyUser
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get User: %w`, err)
	}

	var item *User
	switch len(body.Payload) {
	case 0:
	case 1:
		item = &body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return item, resp, err
}

func (c *Client) ListUser(ctx context.Context) ([]User, *http.Response, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, c.ResourceAPIPath("user"), nil)
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

func (c *Client) UpdateUser(ctx context.Context, data *User) (*User, *http.Response, error) {
	endpointPath := path.Join(c.ResourceAPIPath("user"), data.GetID())
	req, err := c.NewRequest(ctx, http.MethodPut, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyUser
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to update User: %w`, err)
	}

	var item *User
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to update User`)
	case 1:
		item = &body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return item, resp, err
}
