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

type User struct {
	Blocked                       *string `json:"blocked,omitempty"`
	FixedApEnabled                *bool   `json:"fixed_ap_enabled,omitempty"`
	FixedApMac                    *string `json:"fixed_ap_mac,omitempty"`
	FixedIp                       *string `json:"fixed_ip,omitempty"`
	Hostname                      *string `json:"hostname,omitempty"`
	LastSeen                      *string `json:"last_seen,omitempty"`
	LocalDnsRecord                *string `json:"local_dns_record,omitempty"`
	LocalDnsRecordEnabled         *bool   `json:"local_dns_record_enabled,omitempty"`
	Mac                           *string `json:"mac,omitempty"`
	Name                          *string `json:"name,omitempty"`
	NetworkId                     *string `json:"network_id,omitempty"`
	Note                          *string `json:"note,omitempty"`
	UseFixedip                    *bool   `json:"use_fixedip,omitempty"`
	UsergroupId                   *string `json:"usergroup_id,omitempty"`
	VirtualNetworkOverrideEnabled *bool   `json:"virtual_network_override_enabled,omitempty"`
	VirtualNetworkOverrideId      *string `json:"virtual_network_override_id,omitempty"`
}
