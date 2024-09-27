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

type Account struct {
	ID     *string `json:"_id,omitempty"`
	SiteID *string `json:"site_id,omitempty"`

	Hidden   *bool   `json:"attr_hidden,omitempty"`
	HiddenID *string `json:"attr_hidden_id,omitempty"`
	NoDelete *bool   `json:"attr_no_delete,omitempty"`
	NoEdit   *bool   `json:"attr_no_edit,omitempty"`

	Ip               *string `json:"ip,omitempty"`
	Name             *string `json:"name,omitempty"`
	NetworkconfId    *string `json:"networkconf_id,omitempty"`
	TunnelConfigType *string `json:"tunnel_config_type,omitempty"`
	TunnelMediumType *int64  `json:"tunnel_medium_type,omitempty"`
	TunnelType       *int64  `json:"tunnel_type,omitempty"`
	Vlan             *int64  `json:"vlan,omitempty"`
	XPassword        *string `json:"x_password,omitempty"`
}

func (s *Account) GetIp() string {
	if s == nil {
		return ""
	}

	return *s.Ip
}

func (s *Account) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *Account) GetNetworkconfId() string {
	if s == nil {
		return ""
	}

	return *s.NetworkconfId
}

func (s *Account) GetTunnelConfigType() string {
	if s == nil {
		return ""
	}

	return *s.TunnelConfigType
}

func (s *Account) GetTunnelMediumType() int64 {
	if s == nil {
		return 0
	}

	return *s.TunnelMediumType
}

func (s *Account) GetTunnelType() int64 {
	if s == nil {
		return 0
	}

	return *s.TunnelType
}

func (s *Account) GetVlan() int64 {
	if s == nil {
		return 0
	}

	return *s.Vlan
}

func (s *Account) GetXPassword() string {
	if s == nil {
		return ""
	}

	return *s.XPassword
}
