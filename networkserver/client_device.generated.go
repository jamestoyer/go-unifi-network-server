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

type ClientDevice struct {
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
	Blocked *bool `json:"blocked,omitempty"`
	// FixedAPEnabled has been auto generated from the Unifi Network Server API specification
	//
	// Validation: false|true
	FixedAPEnabled *bool `json:"fixed_ap_enabled,omitempty"`
	// FixedAPMAC has been auto generated from the Unifi Network Server API specification
	//
	// Validation: ^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$
	FixedAPMAC *string `json:"fixed_ap_mac,omitempty"`
	// FixedIP has been auto generated from the Unifi Network Server API specification
	//
	// Validation: None
	FixedIP *string `json:"fixed_ip,omitempty"`
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
	// The CurrentIP address assigned to the ClientDevice. This is read only and will only be populated when
	// a request to GetByMAC is made.
	CurrentIP *string `json:"current_ip,omitempty"`
	// The DeviceIDOverride allows for the type of client device to be overridden, which subsequently changes
	// which icon the client device is displayed with in the UI.
	DeviceIDOverride *int64 `json:"dev_id_override,omitempty"`
}

// GetID is a helper function which dereferences ID.
//
// When ID is a nil pointer it will return `""` as default.
func (s *ClientDevice) GetID() string {
	if s == nil || s.ID == nil {
		return ""
	}

	return *s.ID
}

// GetSiteID is a helper function which dereferences SiteID.
//
// When SiteID is a nil pointer it will return `""` as default.
func (s *ClientDevice) GetSiteID() string {
	if s == nil || s.SiteID == nil {
		return ""
	}

	return *s.SiteID
}

// GetHidden is a helper function which dereferences Hidden.
//
// When Hidden is a nil pointer it will return `false` as default.
func (s *ClientDevice) GetHidden() bool {
	if s == nil || s.Hidden == nil {
		return false
	}

	return *s.Hidden
}

// GetHiddenID is a helper function which dereferences HiddenID.
//
// When HiddenID is a nil pointer it will return `""` as default.
func (s *ClientDevice) GetHiddenID() string {
	if s == nil || s.HiddenID == nil {
		return ""
	}

	return *s.HiddenID
}

// GetNoDelete is a helper function which dereferences NoDelete.
//
// When NoDelete is a nil pointer it will return `false` as default.
func (s *ClientDevice) GetNoDelete() bool {
	if s == nil || s.NoDelete == nil {
		return false
	}

	return *s.NoDelete
}

// GetNoEdit is a helper function which dereferences NoEdit.
//
// When NoEdit is a nil pointer it will return `false` as default.
func (s *ClientDevice) GetNoEdit() bool {
	if s == nil || s.NoEdit == nil {
		return false
	}

	return *s.NoEdit
}

// GetBlocked is a helper function which dereferences Blocked.
//
// When Blocked is a nil pointer it will return `false` as default.
func (s *ClientDevice) GetBlocked() bool {
	if s == nil || s.Blocked == nil {
		return false
	}

	return *s.Blocked
}

// GetFixedAPEnabled is a helper function which dereferences FixedAPEnabled.
//
// When FixedAPEnabled is a nil pointer it will return `false` as default.
func (s *ClientDevice) GetFixedAPEnabled() bool {
	if s == nil || s.FixedAPEnabled == nil {
		return false
	}

	return *s.FixedAPEnabled
}

// GetFixedAPMAC is a helper function which dereferences FixedAPMAC.
//
// When FixedAPMAC is a nil pointer it will return `""` as default.
func (s *ClientDevice) GetFixedAPMAC() string {
	if s == nil || s.FixedAPMAC == nil {
		return ""
	}

	return *s.FixedAPMAC
}

// GetFixedIP is a helper function which dereferences FixedIP.
//
// When FixedIP is a nil pointer it will return `""` as default.
func (s *ClientDevice) GetFixedIP() string {
	if s == nil || s.FixedIP == nil {
		return ""
	}

	return *s.FixedIP
}

// GetHostname is a helper function which dereferences Hostname.
//
// When Hostname is a nil pointer it will return `""` as default.
func (s *ClientDevice) GetHostname() string {
	if s == nil || s.Hostname == nil {
		return ""
	}

	return *s.Hostname
}

// GetLastSeen is a helper function which dereferences LastSeen.
//
// When LastSeen is a nil pointer it will return `""` as default.
func (s *ClientDevice) GetLastSeen() string {
	if s == nil || s.LastSeen == nil {
		return ""
	}

	return *s.LastSeen
}

// GetLocalDNSRecord is a helper function which dereferences LocalDNSRecord.
//
// When LocalDNSRecord is a nil pointer it will return `""` as default.
func (s *ClientDevice) GetLocalDNSRecord() string {
	if s == nil || s.LocalDNSRecord == nil {
		return ""
	}

	return *s.LocalDNSRecord
}

// GetLocalDNSRecordEnabled is a helper function which dereferences LocalDNSRecordEnabled.
//
// When LocalDNSRecordEnabled is a nil pointer it will return `false` as default.
func (s *ClientDevice) GetLocalDNSRecordEnabled() bool {
	if s == nil || s.LocalDNSRecordEnabled == nil {
		return false
	}

	return *s.LocalDNSRecordEnabled
}

// GetMAC is a helper function which dereferences MAC.
//
// When MAC is a nil pointer it will return `""` as default.
func (s *ClientDevice) GetMAC() string {
	if s == nil || s.MAC == nil {
		return ""
	}

	return *s.MAC
}

// GetName is a helper function which dereferences Name.
//
// When Name is a nil pointer it will return `""` as default.
func (s *ClientDevice) GetName() string {
	if s == nil || s.Name == nil {
		return ""
	}

	return *s.Name
}

// GetNetworkID is a helper function which dereferences NetworkID.
//
// When NetworkID is a nil pointer it will return `""` as default.
func (s *ClientDevice) GetNetworkID() string {
	if s == nil || s.NetworkID == nil {
		return ""
	}

	return *s.NetworkID
}

// GetNote is a helper function which dereferences Note.
//
// When Note is a nil pointer it will return `""` as default.
func (s *ClientDevice) GetNote() string {
	if s == nil || s.Note == nil {
		return ""
	}

	return *s.Note
}

// GetUseFixedIP is a helper function which dereferences UseFixedIP.
//
// When UseFixedIP is a nil pointer it will return `false` as default.
func (s *ClientDevice) GetUseFixedIP() bool {
	if s == nil || s.UseFixedIP == nil {
		return false
	}

	return *s.UseFixedIP
}

// GetUserGroupID is a helper function which dereferences UserGroupID.
//
// When UserGroupID is a nil pointer it will return `""` as default.
func (s *ClientDevice) GetUserGroupID() string {
	if s == nil || s.UserGroupID == nil {
		return ""
	}

	return *s.UserGroupID
}

// GetVirtualNetworkOverrideEnabled is a helper function which dereferences VirtualNetworkOverrideEnabled.
//
// When VirtualNetworkOverrideEnabled is a nil pointer it will return `false` as default.
func (s *ClientDevice) GetVirtualNetworkOverrideEnabled() bool {
	if s == nil || s.VirtualNetworkOverrideEnabled == nil {
		return false
	}

	return *s.VirtualNetworkOverrideEnabled
}

// GetVirtualNetworkOverrideID is a helper function which dereferences VirtualNetworkOverrideID.
//
// When VirtualNetworkOverrideID is a nil pointer it will return `""` as default.
func (s *ClientDevice) GetVirtualNetworkOverrideID() string {
	if s == nil || s.VirtualNetworkOverrideID == nil {
		return ""
	}

	return *s.VirtualNetworkOverrideID
}

// GetCurrentIP is a helper function which dereferences CurrentIP.
//
// When CurrentIP is a nil pointer it will return `""` as default.
func (s *ClientDevice) GetCurrentIP() string {
	if s == nil || s.CurrentIP == nil {
		return ""
	}

	return *s.CurrentIP
}

// GetDeviceIDOverride is a helper function which dereferences DeviceIDOverride.
//
// When DeviceIDOverride is a nil pointer it will return `0` as default.
func (s *ClientDevice) GetDeviceIDOverride() int64 {
	if s == nil || s.DeviceIDOverride == nil {
		return 0
	}

	return *s.DeviceIDOverride
}

type ClientDeviceService service

type responseBodyClientDevice struct {
	Metadata json.RawMessage `json:"meta"`
	Payload  []ClientDevice  `json:"data"`
}

func (s *ClientDeviceService) Create(ctx context.Context, data *ClientDevice) (*ClientDevice, *http.Response, error) {
	req, err := s.NewRequest(http.MethodPost, s.ResourceAPIPath("user"), data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyClientDevice
	resp, err := s.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to create ClientDevice: %w`, err)
	}

	var item *ClientDevice
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to create ClientDevice`)
	case 1:
		item = &body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return item, resp, err
}

func (s *ClientDeviceService) Get(ctx context.Context, id string) (*ClientDevice, *http.Response, error) {
	endpointPath := path.Join(s.ResourceAPIPath("user"), id)
	req, err := s.NewRequest(http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyClientDevice
	resp, err := s.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get ClientDevice: %w`, err)
	}

	var item *ClientDevice
	switch len(body.Payload) {
	case 0:
	case 1:
		item = &body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return item, resp, err
}

func (s *ClientDeviceService) List(ctx context.Context) ([]ClientDevice, *http.Response, error) {
	req, err := s.NewRequest(http.MethodGet, s.ResourceAPIPath("user"), nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyClientDevice
	resp, err := s.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get ClientDevice: %w`, err)
	}

	return body.Payload, resp, nil
}

func (s *ClientDeviceService) Update(ctx context.Context, data *ClientDevice) (*ClientDevice, *http.Response, error) {
	endpointPath := path.Join(s.ResourceAPIPath("user"), data.GetID())
	req, err := s.NewRequest(http.MethodPut, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyClientDevice
	resp, err := s.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to update ClientDevice: %w`, err)
	}

	var item *ClientDevice
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to update ClientDevice`)
	case 1:
		item = &body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return item, resp, err
}
