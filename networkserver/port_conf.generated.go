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

type PortConf struct {
	Autoneg                       *bool               `json:"autoneg,omitempty"`
	Dot1XCtrl                     *string             `json:"dot1x_ctrl,omitempty"`
	Dot1XIdleTimeout              *float64            `json:"dot1x_idle_timeout,omitempty"`
	EgressRateLimitKbps           *int64              `json:"egress_rate_limit_kbps,omitempty"`
	EgressRateLimitKbpsEnabled    *bool               `json:"egress_rate_limit_kbps_enabled,omitempty"`
	ExcludedNetworkconfIds        *[]string           `json:"excluded_networkconf_ids,omitempty"`
	FecMode                       *string             `json:"fec_mode,omitempty"`
	Forward                       *string             `json:"forward,omitempty"`
	FullDuplex                    *bool               `json:"full_duplex,omitempty"`
	Isolation                     *bool               `json:"isolation,omitempty"`
	LldpmedEnabled                *bool               `json:"lldpmed_enabled,omitempty"`
	LldpmedNotifyEnabled          *bool               `json:"lldpmed_notify_enabled,omitempty"`
	MulticastRouterNetworkconfIds *[]string           `json:"multicast_router_networkconf_ids,omitempty"`
	Name                          *string             `json:"name,omitempty"`
	NativeNetworkconfId           *string             `json:"native_networkconf_id,omitempty"`
	OpMode                        *string             `json:"op_mode,omitempty"`
	PoeMode                       *string             `json:"poe_mode,omitempty"`
	PortKeepaliveEnabled          *bool               `json:"port_keepalive_enabled,omitempty"`
	PortSecurityEnabled           *bool               `json:"port_security_enabled,omitempty"`
	PortSecurityMacAddress        *[]string           `json:"port_security_mac_address,omitempty"`
	PriorityQueue1Level           *int64              `json:"priority_queue1_level,omitempty"`
	PriorityQueue2Level           *int64              `json:"priority_queue2_level,omitempty"`
	PriorityQueue3Level           *int64              `json:"priority_queue3_level,omitempty"`
	PriorityQueue4Level           *int64              `json:"priority_queue4_level,omitempty"`
	QosProfile                    *PortConfQosProfile `json:"qos_profile,omitempty"`
	SettingPreference             *string             `json:"setting_preference,omitempty"`
	Speed                         *float64            `json:"speed,omitempty"`
	StormctrlBcastEnabled         *bool               `json:"stormctrl_bcast_enabled,omitempty"`
	StormctrlBcastLevel           *int64              `json:"stormctrl_bcast_level,omitempty"`
	StormctrlBcastRate            *float64            `json:"stormctrl_bcast_rate,omitempty"`
	StormctrlMcastEnabled         *bool               `json:"stormctrl_mcast_enabled,omitempty"`
	StormctrlMcastLevel           *int64              `json:"stormctrl_mcast_level,omitempty"`
	StormctrlMcastRate            *float64            `json:"stormctrl_mcast_rate,omitempty"`
	StormctrlType                 *string             `json:"stormctrl_type,omitempty"`
	StormctrlUcastEnabled         *bool               `json:"stormctrl_ucast_enabled,omitempty"`
	StormctrlUcastLevel           *int64              `json:"stormctrl_ucast_level,omitempty"`
	StormctrlUcastRate            *float64            `json:"stormctrl_ucast_rate,omitempty"`
	StpPortMode                   *bool               `json:"stp_port_mode,omitempty"`
	TaggedVlanMgmt                *string             `json:"tagged_vlan_mgmt,omitempty"`
	VoiceNetworkconfId            *string             `json:"voice_networkconf_id,omitempty"`
}

type PortConfQosProfile struct {
	QosPolicies    *[]PortConfQosProfileQosPolicies `json:"qos_policies,omitempty"`
	QosProfileMode *string                          `json:"qos_profile_mode,omitempty"`
}

type PortConfQosProfileQosPolicies struct {
	QosMarking  *PortConfQosProfileQosPoliciesQosMarking  `json:"qos_marking,omitempty"`
	QosMatching *PortConfQosProfileQosPoliciesQosMatching `json:"qos_matching,omitempty"`
}

type PortConfQosProfileQosPoliciesQosMarking struct {
	CosCode          *int64   `json:"cos_code,omitempty"`
	DscpCode         *float64 `json:"dscp_code,omitempty"`
	IpPrecedenceCode *int64   `json:"ip_precedence_code,omitempty"`
	Queue            *int64   `json:"queue,omitempty"`
}

type PortConfQosProfileQosPoliciesQosMatching struct {
	CosCode          *int64   `json:"cos_code,omitempty"`
	DscpCode         *int64   `json:"dscp_code,omitempty"`
	DstPort          *float64 `json:"dst_port,omitempty"`
	IpPrecedenceCode *int64   `json:"ip_precedence_code,omitempty"`
	Protocol         *string  `json:"protocol,omitempty"`
	SrcPort          *float64 `json:"src_port,omitempty"`
}
