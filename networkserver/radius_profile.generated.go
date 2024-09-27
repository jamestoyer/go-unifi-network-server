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

type RadiusProfile struct {
	ID     *string `json:"_id,omitempty"`
	SiteID *string `json:"site_id,omitempty"`

	Hidden   *bool   `json:"attr_hidden,omitempty"`
	HiddenID *string `json:"attr_hidden_id,omitempty"`
	NoDelete *bool   `json:"attr_no_delete,omitempty"`
	NoEdit   *bool   `json:"attr_no_edit,omitempty"`

	AccountingEnabled         *bool                       `json:"accounting_enabled,omitempty"`
	AcctServers               *[]RadiusProfileAcctServers `json:"acct_servers,omitempty"`
	AuthServers               *[]RadiusProfileAuthServers `json:"auth_servers,omitempty"`
	InterimUpdateEnabled      *bool                       `json:"interim_update_enabled,omitempty"`
	InterimUpdateInterval     *float64                    `json:"interim_update_interval,omitempty"`
	Name                      *string                     `json:"name,omitempty"`
	TlsEnabled                *bool                       `json:"tls_enabled,omitempty"`
	UseUsgAcctServer          *bool                       `json:"use_usg_acct_server,omitempty"`
	UseUsgAuthServer          *bool                       `json:"use_usg_auth_server,omitempty"`
	VlanEnabled               *bool                       `json:"vlan_enabled,omitempty"`
	VlanWlanMode              *string                     `json:"vlan_wlan_mode,omitempty"`
	XCaCrt                    *string                     `json:"x_ca_crt,omitempty"`
	XCaCrtFilename            *string                     `json:"x_ca_crt_filename,omitempty"`
	XClientCrt                *string                     `json:"x_client_crt,omitempty"`
	XClientCrtFilename        *string                     `json:"x_client_crt_filename,omitempty"`
	XClientPrivateKey         *string                     `json:"x_client_private_key,omitempty"`
	XClientPrivateKeyFilename *string                     `json:"x_client_private_key_filename,omitempty"`
	XClientPrivateKeyPassword *string                     `json:"x_client_private_key_password,omitempty"`
}

func (s *RadiusProfile) GetAccountingEnabled() bool {
	if s == nil {
		return false
	}

	return *s.AccountingEnabled
}

func (s *RadiusProfile) GetAcctServers() []RadiusProfileAcctServers {
	if s == nil || s.AcctServers == nil {
		return nil
	}

	return *s.AcctServers
}

func (s *RadiusProfile) GetAuthServers() []RadiusProfileAuthServers {
	if s == nil || s.AuthServers == nil {
		return nil
	}

	return *s.AuthServers
}

func (s *RadiusProfile) GetInterimUpdateEnabled() bool {
	if s == nil {
		return false
	}

	return *s.InterimUpdateEnabled
}

func (s *RadiusProfile) GetInterimUpdateInterval() float64 {
	if s == nil {
		return 0
	}

	return *s.InterimUpdateInterval
}

func (s *RadiusProfile) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *RadiusProfile) GetTlsEnabled() bool {
	if s == nil {
		return false
	}

	return *s.TlsEnabled
}

func (s *RadiusProfile) GetUseUsgAcctServer() bool {
	if s == nil {
		return false
	}

	return *s.UseUsgAcctServer
}

func (s *RadiusProfile) GetUseUsgAuthServer() bool {
	if s == nil {
		return false
	}

	return *s.UseUsgAuthServer
}

func (s *RadiusProfile) GetVlanEnabled() bool {
	if s == nil {
		return false
	}

	return *s.VlanEnabled
}

func (s *RadiusProfile) GetVlanWlanMode() string {
	if s == nil {
		return ""
	}

	return *s.VlanWlanMode
}

func (s *RadiusProfile) GetXCaCrt() string {
	if s == nil {
		return ""
	}

	return *s.XCaCrt
}

func (s *RadiusProfile) GetXCaCrtFilename() string {
	if s == nil {
		return ""
	}

	return *s.XCaCrtFilename
}

func (s *RadiusProfile) GetXClientCrt() string {
	if s == nil {
		return ""
	}

	return *s.XClientCrt
}

func (s *RadiusProfile) GetXClientCrtFilename() string {
	if s == nil {
		return ""
	}

	return *s.XClientCrtFilename
}

func (s *RadiusProfile) GetXClientPrivateKey() string {
	if s == nil {
		return ""
	}

	return *s.XClientPrivateKey
}

func (s *RadiusProfile) GetXClientPrivateKeyFilename() string {
	if s == nil {
		return ""
	}

	return *s.XClientPrivateKeyFilename
}

func (s *RadiusProfile) GetXClientPrivateKeyPassword() string {
	if s == nil {
		return ""
	}

	return *s.XClientPrivateKeyPassword
}

type RadiusProfileAcctServers struct {
	Ip      *string  `json:"ip,omitempty"`
	Port    *float64 `json:"port,omitempty"`
	XSecret *string  `json:"x_secret,omitempty"`
}

func (s *RadiusProfileAcctServers) GetIp() string {
	if s == nil {
		return ""
	}

	return *s.Ip
}

func (s *RadiusProfileAcctServers) GetPort() float64 {
	if s == nil {
		return 0
	}

	return *s.Port
}

func (s *RadiusProfileAcctServers) GetXSecret() string {
	if s == nil {
		return ""
	}

	return *s.XSecret
}

type RadiusProfileAuthServers struct {
	Ip      *string  `json:"ip,omitempty"`
	Port    *float64 `json:"port,omitempty"`
	XSecret *string  `json:"x_secret,omitempty"`
}

func (s *RadiusProfileAuthServers) GetIp() string {
	if s == nil {
		return ""
	}

	return *s.Ip
}

func (s *RadiusProfileAuthServers) GetPort() float64 {
	if s == nil {
		return 0
	}

	return *s.Port
}

func (s *RadiusProfileAuthServers) GetXSecret() string {
	if s == nil {
		return ""
	}

	return *s.XSecret
}
