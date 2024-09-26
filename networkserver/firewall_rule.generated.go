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

type FirewallRule struct {
	Action                *string   `json:"action,omitempty"`
	Contiguous            *bool     `json:"contiguous,omitempty"`
	DstAddress            *string   `json:"dst_address,omitempty"`
	DstAddressIpv6        *string   `json:"dst_address_ipv6,omitempty"`
	DstFirewallgroupIds   *[]string `json:"dst_firewallgroup_ids,omitempty"`
	DstNetworkconfId      *string   `json:"dst_networkconf_id,omitempty"`
	DstNetworkconfType    *string   `json:"dst_networkconf_type,omitempty"`
	DstPort               *string   `json:"dst_port,omitempty"`
	Enabled               *bool     `json:"enabled,omitempty"`
	IcmpTypename          *string   `json:"icmp_typename,omitempty"`
	Icmpv6Typename        *string   `json:"icmpv6_typename,omitempty"`
	Ipsec                 *string   `json:"ipsec,omitempty"`
	Logging               *bool     `json:"logging,omitempty"`
	Monthdays             *string   `json:"monthdays,omitempty"`
	MonthdaysNegate       *bool     `json:"monthdays_negate,omitempty"`
	Name                  *string   `json:"name,omitempty"`
	Protocol              *string   `json:"protocol,omitempty"`
	ProtocolMatchExcepted *bool     `json:"protocol_match_excepted,omitempty"`
	ProtocolV6            *string   `json:"protocol_v6,omitempty"`
	RuleIndex             *int64    `json:"rule_index,omitempty"`
	Ruleset               *string   `json:"ruleset,omitempty"`
	SettingPreference     *string   `json:"setting_preference,omitempty"`
	SrcAddress            *string   `json:"src_address,omitempty"`
	SrcAddressIpv6        *string   `json:"src_address_ipv6,omitempty"`
	SrcFirewallgroupIds   *[]string `json:"src_firewallgroup_ids,omitempty"`
	SrcMacAddress         *string   `json:"src_mac_address,omitempty"`
	SrcNetworkconfId      *string   `json:"src_networkconf_id,omitempty"`
	SrcNetworkconfType    *string   `json:"src_networkconf_type,omitempty"`
	SrcPort               *string   `json:"src_port,omitempty"`
	Startdate             *string   `json:"startdate,omitempty"`
	Starttime             *string   `json:"starttime,omitempty"`
	StateEstablished      *bool     `json:"state_established,omitempty"`
	StateInvalid          *bool     `json:"state_invalid,omitempty"`
	StateNew              *bool     `json:"state_new,omitempty"`
	StateRelated          *bool     `json:"state_related,omitempty"`
	Stopdate              *string   `json:"stopdate,omitempty"`
	Stoptime              *string   `json:"stoptime,omitempty"`
	Utc                   *bool     `json:"utc,omitempty"`
	Weekdays              *string   `json:"weekdays,omitempty"`
	WeekdaysNegate        *bool     `json:"weekdays_negate,omitempty"`
}
