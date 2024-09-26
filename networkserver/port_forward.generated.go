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

type PortForward struct {
	DestinationIp      *string                      `json:"destination_ip,omitempty"`
	DestinationIps     *[]PortForwardDestinationIps `json:"destination_ips,omitempty"`
	DstPort            *string                      `json:"dst_port,omitempty"`
	Enabled            *bool                        `json:"enabled,omitempty"`
	Fwd                *string                      `json:"fwd,omitempty"`
	FwdPort            *string                      `json:"fwd_port,omitempty"`
	Log                *bool                        `json:"log,omitempty"`
	Name               *string                      `json:"name,omitempty"`
	PfwdInterface      *string                      `json:"pfwd_interface,omitempty"`
	Proto              *string                      `json:"proto,omitempty"`
	Src                *string                      `json:"src,omitempty"`
	SrcFirewallGroupId *string                      `json:"src_firewall_group_id,omitempty"`
	SrcLimitingEnabled *bool                        `json:"src_limiting_enabled,omitempty"`
	SrcLimitingType    *string                      `json:"src_limiting_type,omitempty"`
}

type PortForwardDestinationIps struct {
	DestinationIp *string `json:"destination_ip,omitempty"`
	Interface     *string `json:"interface,omitempty"`
}
