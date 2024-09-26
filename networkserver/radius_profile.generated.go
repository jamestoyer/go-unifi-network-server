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

type RadiusProfileAcctServers struct {
	Ip      *string  `json:"ip,omitempty"`
	Port    *float64 `json:"port,omitempty"`
	XSecret *string  `json:"x_secret,omitempty"`
}

type RadiusProfileAuthServers struct {
	Ip      *string  `json:"ip,omitempty"`
	Port    *float64 `json:"port,omitempty"`
	XSecret *string  `json:"x_secret,omitempty"`
}
